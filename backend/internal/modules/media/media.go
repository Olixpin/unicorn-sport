package media

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/unicorn-sport/backend/internal/domain"
)

// MediaModule handles video and media operations
type MediaModule struct {
	db            *gorm.DB
	s3Client      *s3.Client
	s3Bucket      string
	cloudFrontURL string
	awsRegion     string
}

// NewMediaModule creates a new media module
func NewMediaModule(db *gorm.DB, awsRegion, accessKeyID, secretAccessKey, s3Bucket, cloudFrontURL string) *MediaModule {
	var s3Client *s3.Client

	if accessKeyID != "" && secretAccessKey != "" {
		cfg, err := config.LoadDefaultConfig(context.Background(),
			config.WithRegion(awsRegion),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")),
		)
		if err == nil {
			s3Client = s3.NewFromConfig(cfg)
		}
	}

	return &MediaModule{
		db:            db,
		s3Client:      s3Client,
		s3Bucket:      s3Bucket,
		cloudFrontURL: cloudFrontURL,
		awsRegion:     awsRegion,
	}
}

// GetS3Client returns the S3 client for use by other modules
func (m *MediaModule) GetS3Client() *s3.Client {
	return m.s3Client
}

// =============================================================================
// UPLOAD ENDPOINTS
// =============================================================================

// InitUploadRequest initiates a new upload session
type InitUploadRequest struct {
	UploadType  string `json:"upload_type" binding:"required,oneof=highlight full_match thumbnail"`
	FileName    string `json:"file_name" binding:"required"`
	ContentType string `json:"content_type" binding:"required"`
	FileSize    int64  `json:"file_size" binding:"required"`
}

// InitUploadResponse returns upload session details
type InitUploadResponse struct {
	SessionID    uuid.UUID `json:"session_id"`
	UploadURL    string    `json:"upload_url"`
	UploadMethod string    `json:"upload_method"` // "direct" or "multipart"
	S3Key        string    `json:"s3_key"`
	ExpiresIn    int       `json:"expires_in"`
	MaxFileSize  int64     `json:"max_file_size"`
}

// InitUpload starts a new upload session and returns presigned URL
func (m *MediaModule) InitUpload(c *gin.Context) {
	var req InitUploadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
		return
	}

	// Validate content type based on upload type
	var allowedTypes map[string]bool
	var maxSize int64

	switch req.UploadType {
	case "highlight", "full_match":
		allowedTypes = map[string]bool{
			"video/mp4":        true,
			"video/quicktime":  true,
			"video/x-msvideo":  true,
			"video/webm":       true,
			"video/x-matroska": true,
		}
		if req.UploadType == "highlight" {
			maxSize = 1 * 1024 * 1024 * 1024 // 1GB for highlights
		} else {
			maxSize = 25 * 1024 * 1024 * 1024 // 25GB for full matches
		}
	case "thumbnail":
		allowedTypes = map[string]bool{
			"image/jpeg": true,
			"image/png":  true,
			"image/webp": true,
		}
		maxSize = 10 * 1024 * 1024 // 10MB for thumbnails
	default:
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_UPLOAD_TYPE", "message": "Invalid upload type"}})
		return
	}

	if !allowedTypes[req.ContentType] {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_CONTENT_TYPE", "message": "File type not allowed for this upload type"}})
		return
	}

	if req.FileSize > maxSize {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "FILE_TOO_LARGE", "message": fmt.Sprintf("Maximum file size is %d MB", maxSize/(1024*1024))}})
		return
	}

	adminID := c.MustGet("user_id").(uuid.UUID)

	// Generate unique session ID and S3 key
	sessionID := uuid.New()
	var s3Key string

	switch req.UploadType {
	case "highlight":
		s3Key = fmt.Sprintf("highlights/%s/%s", sessionID.String(), sanitizeFileName(req.FileName))
	case "full_match":
		now := time.Now()
		s3Key = fmt.Sprintf("matches/%d/%02d/%s/%s", now.Year(), now.Month(), sessionID.String(), sanitizeFileName(req.FileName))
	case "thumbnail":
		s3Key = fmt.Sprintf("thumbnails/%s/%s", sessionID.String(), sanitizeFileName(req.FileName))
	}

	expiresIn := 3600 // 1 hour
	expiresAt := time.Now().Add(time.Duration(expiresIn) * time.Second)

	// Create upload session record
	session := domain.UploadSession{
		ID:          sessionID,
		UploadType:  req.UploadType,
		ContentType: req.ContentType,
		FileName:    req.FileName,
		FileSize:    req.FileSize,
		S3Key:       s3Key,
		Status:      "pending",
		UploadedBy:  adminID,
		ExpiresAt:   expiresAt,
	}

	if err := m.db.Create(&session).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "SESSION_CREATE_FAILED", "message": "Failed to create upload session"}})
		return
	}

	// Generate presigned URL
	var uploadURL string
	uploadMethod := "direct"

	if m.s3Client == nil {
		// Development mode - return mock URL
		uploadURL = fmt.Sprintf("http://localhost:9000/%s/%s", m.s3Bucket, s3Key)
	} else {
		// For large files (>100MB), use multipart upload
		if req.FileSize > 100*1024*1024 {
			uploadMethod = "multipart"
			// For multipart, the frontend will request part URLs separately
			uploadURL = "" // Frontend will call GetUploadPartURL for each part
		} else {
			// Direct upload with presigned PUT
			presignClient := s3.NewPresignClient(m.s3Client)
			presignedReq, err := presignClient.PresignPutObject(context.Background(), &s3.PutObjectInput{
				Bucket:        aws.String(m.s3Bucket),
				Key:           aws.String(s3Key),
				ContentType:   aws.String(req.ContentType),
				ContentLength: aws.Int64(req.FileSize),
			}, s3.WithPresignExpires(time.Duration(expiresIn)*time.Second))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "PRESIGN_FAILED", "message": "Failed to generate upload URL"}})
				return
			}
			uploadURL = presignedReq.URL
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": InitUploadResponse{
			SessionID:    sessionID,
			UploadURL:    uploadURL,
			UploadMethod: uploadMethod,
			S3Key:        s3Key,
			ExpiresIn:    expiresIn,
			MaxFileSize:  maxSize,
		},
	})
}

// InitMultipartRequest starts multipart upload for large files
type InitMultipartRequest struct {
	SessionID string `json:"session_id" binding:"required"`
}

// InitMultipart initiates S3 multipart upload
func (m *MediaModule) InitMultipart(c *gin.Context) {
	var req InitMultipartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
		return
	}

	sessionID, err := uuid.Parse(req.SessionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_SESSION_ID", "message": "Invalid session ID"}})
		return
	}

	var session domain.UploadSession
	if err := m.db.First(&session, "id = ?", sessionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "SESSION_NOT_FOUND", "message": "Upload session not found"}})
		return
	}

	if session.Status != "pending" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_SESSION_STATE", "message": "Upload already started or completed"}})
		return
	}

	if m.s3Client == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"success": false, "error": gin.H{"code": "S3_NOT_CONFIGURED", "message": "S3 is not configured"}})
		return
	}

	// Create multipart upload
	createResp, err := m.s3Client.CreateMultipartUpload(context.Background(), &s3.CreateMultipartUploadInput{
		Bucket:      aws.String(m.s3Bucket),
		Key:         aws.String(session.S3Key),
		ContentType: aws.String(session.ContentType),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "MULTIPART_INIT_FAILED", "message": "Failed to initialize multipart upload"}})
		return
	}

	// Calculate number of parts (100MB each, except last part)
	partSize := int64(100 * 1024 * 1024) // 100MB
	totalParts := int((session.FileSize + partSize - 1) / partSize)

	// Update session
	uploadID := *createResp.UploadId
	session.S3UploadID = &uploadID
	session.PartsTotal = &totalParts
	session.Status = "uploading"
	m.db.Save(&session)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"upload_id":   uploadID,
			"total_parts": totalParts,
			"part_size":   partSize,
		},
	})
}

// GetPartURLRequest gets presigned URL for a single part
type GetPartURLRequest struct {
	SessionID  string `json:"session_id" binding:"required"`
	PartNumber int    `json:"part_number" binding:"required,min=1,max=10000"`
}

// GetUploadPartURL generates presigned URL for a multipart upload part
func (m *MediaModule) GetUploadPartURL(c *gin.Context) {
	var req GetPartURLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
		return
	}

	sessionID, _ := uuid.Parse(req.SessionID)
	var session domain.UploadSession
	if err := m.db.First(&session, "id = ?", sessionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "SESSION_NOT_FOUND", "message": "Upload session not found"}})
		return
	}

	if session.S3UploadID == nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "MULTIPART_NOT_STARTED", "message": "Multipart upload not initialized"}})
		return
	}

	presignClient := s3.NewPresignClient(m.s3Client)
	presignedReq, err := presignClient.PresignUploadPart(context.Background(), &s3.UploadPartInput{
		Bucket:     aws.String(m.s3Bucket),
		Key:        aws.String(session.S3Key),
		UploadId:   session.S3UploadID,
		PartNumber: aws.Int32(int32(req.PartNumber)),
	}, s3.WithPresignExpires(15*time.Minute))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "PRESIGN_FAILED", "message": "Failed to generate part URL"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"upload_url":  presignedReq.URL,
			"part_number": req.PartNumber,
			"expires_in":  900, // 15 minutes
		},
	})
}

// CompleteMultipartRequest completes a multipart upload
type CompleteMultipartRequest struct {
	SessionID string              `json:"session_id" binding:"required"`
	Parts     []CompletedPartInfo `json:"parts" binding:"required"`
}

type CompletedPartInfo struct {
	PartNumber int    `json:"part_number"`
	ETag       string `json:"etag"`
}

// CompleteMultipart completes the multipart upload
func (m *MediaModule) CompleteMultipart(c *gin.Context) {
	var req CompleteMultipartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
		return
	}

	sessionID, _ := uuid.Parse(req.SessionID)
	var session domain.UploadSession
	if err := m.db.First(&session, "id = ?", sessionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "SESSION_NOT_FOUND", "message": "Upload session not found"}})
		return
	}

	// Build completed parts list
	var completedParts []types.CompletedPart
	for _, p := range req.Parts {
		completedParts = append(completedParts, types.CompletedPart{
			PartNumber: aws.Int32(int32(p.PartNumber)),
			ETag:       aws.String(p.ETag),
		})
	}

	// Complete multipart upload
	_, err := m.s3Client.CompleteMultipartUpload(context.Background(), &s3.CompleteMultipartUploadInput{
		Bucket:   aws.String(m.s3Bucket),
		Key:      aws.String(session.S3Key),
		UploadId: session.S3UploadID,
		MultipartUpload: &types.CompletedMultipartUpload{
			Parts: completedParts,
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "COMPLETE_FAILED", "message": "Failed to complete upload"}})
		return
	}

	// Update session
	now := time.Now()
	session.Status = "completed"
	session.CompletedAt = &now
	m.db.Save(&session)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"session_id": session.ID,
			"s3_key":     session.S3Key,
			"s3_url":     fmt.Sprintf("s3://%s/%s", m.s3Bucket, session.S3Key),
		},
	})
}

// =============================================================================
// VIDEO CRUD ENDPOINTS
// =============================================================================

// CreateVideoRequest creates a new video record
type CreateVideoRequest struct {
	VideoType       string   `json:"video_type" binding:"required,oneof=highlight full_match"`
	Title           string   `json:"title" binding:"required"`
	Description     *string  `json:"description,omitempty"`
	SessionID       *string  `json:"session_id,omitempty"` // Upload session ID
	BlobURL         *string  `json:"blob_url,omitempty"`   // Direct S3 URL (if not using session)
	ThumbnailURL    *string  `json:"thumbnail_url,omitempty"`
	DurationSeconds *int     `json:"duration_seconds,omitempty"`
	TournamentID    *string  `json:"tournament_id,omitempty"`
	MatchID         *string  `json:"match_id,omitempty"` // Source match for highlight
	PlayerIDs       []string `json:"player_ids,omitempty"`
}

// CreateVideo creates a video record after upload
func (m *MediaModule) CreateVideo(c *gin.Context) {
	var req CreateVideoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
		return
	}

	adminID := c.MustGet("user_id").(uuid.UUID)

	var blobURL string
	var fileSize *int64

	// Get S3 URL from session or direct URL
	if req.SessionID != nil && *req.SessionID != "" {
		sessionID, _ := uuid.Parse(*req.SessionID)
		var session domain.UploadSession
		if err := m.db.First(&session, "id = ? AND status = 'completed'", sessionID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_SESSION", "message": "Upload session not found or not completed"}})
			return
		}
		blobURL = fmt.Sprintf("s3://%s/%s", m.s3Bucket, session.S3Key)
		fileSize = &session.FileSize
	} else if req.BlobURL != nil && *req.BlobURL != "" {
		blobURL = *req.BlobURL
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "MISSING_URL", "message": "Either session_id or blob_url is required"}})
		return
	}

	// Parse optional IDs
	var tournamentID *uuid.UUID
	if req.TournamentID != nil && *req.TournamentID != "" {
		tid, _ := uuid.Parse(*req.TournamentID)
		tournamentID = &tid
	}

	var matchID *uuid.UUID
	if req.MatchID != nil && *req.MatchID != "" {
		mid, _ := uuid.Parse(*req.MatchID)
		matchID = &mid
	}

	video := domain.Video{
		VideoType:       req.VideoType,
		Title:           req.Title,
		Description:     req.Description,
		BlobURL:         blobURL,
		ThumbnailURL:    req.ThumbnailURL,
		DurationSeconds: req.DurationSeconds,
		FileSizeBytes:   fileSize,
		Status:          "approved", // Admin uploads are auto-approved
		TournamentID:    tournamentID,
		MatchID:         matchID,
		UploadedBy:      adminID,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	if err := m.db.Create(&video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "CREATE_FAILED", "message": "Failed to create video"}})
		return
	}

	// Link players to video
	for i, pidStr := range req.PlayerIDs {
		if pid, err := uuid.Parse(pidStr); err == nil {
			pv := domain.PlayerVideo{
				PlayerID:  pid,
				VideoID:   video.ID,
				IsPrimary: i == 0,
			}
			m.db.Create(&pv)
		}
	}

	// Load relationships for response
	m.db.Preload("Players").First(&video, "id = ?", video.ID)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    m.videoToResponse(video),
	})
}

// VideoResponse is the API response format for videos
type VideoResponse struct {
	ID              uuid.UUID       `json:"id"`
	VideoType       string          `json:"video_type"`
	Title           string          `json:"title"`
	Description     *string         `json:"description,omitempty"`
	ThumbnailURL    *string         `json:"thumbnail_url,omitempty"`
	VideoURL        string          `json:"video_url,omitempty"` // Streaming URL or presigned
	DurationSeconds *int            `json:"duration_seconds,omitempty"`
	FileSizeBytes   *int64          `json:"file_size_bytes,omitempty"`
	Status          string          `json:"status"`
	ReviewNotes     *string         `json:"review_notes,omitempty"`
	TournamentID    *uuid.UUID      `json:"tournament_id,omitempty"`
	MatchID         *uuid.UUID      `json:"match_id,omitempty"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	Players         []PlayerSummary `json:"players,omitempty"`
}

type PlayerSummary struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Position  string    `json:"position"`
}

func (m *MediaModule) videoToResponse(v domain.Video) VideoResponse {
	resp := VideoResponse{
		ID:              v.ID,
		VideoType:       v.VideoType,
		Title:           v.Title,
		Description:     v.Description,
		ThumbnailURL:    v.ThumbnailURL,
		DurationSeconds: v.DurationSeconds,
		FileSizeBytes:   v.FileSizeBytes,
		Status:          v.Status,
		ReviewNotes:     v.ReviewNotes,
		TournamentID:    v.TournamentID,
		MatchID:         v.MatchID,
		CreatedAt:       v.CreatedAt,
		UpdatedAt:       v.UpdatedAt,
	}

	// Generate video URL
	if m.cloudFrontURL != "" {
		resp.VideoURL = fmt.Sprintf("%s/%s", m.cloudFrontURL, strings.TrimPrefix(v.BlobURL, fmt.Sprintf("s3://%s/", m.s3Bucket)))
	}

	// Map players
	for _, p := range v.Players {
		resp.Players = append(resp.Players, PlayerSummary{
			ID:        p.ID,
			FirstName: p.FirstName,
			LastName:  p.LastName,
			Position:  p.Position,
		})
	}

	return resp
}

// ListVideos lists all videos with filters
func (m *MediaModule) ListVideos(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if limit > 100 {
		limit = 100
	}
	offset := (page - 1) * limit

	query := m.db.Model(&domain.Video{})

	// Apply filters
	if videoType := c.Query("video_type"); videoType != "" {
		query = query.Where("video_type = ?", videoType)
	}
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}
	if tournamentID := c.Query("tournament_id"); tournamentID != "" {
		query = query.Where("tournament_id = ?", tournamentID)
	}
	if playerID := c.Query("player_id"); playerID != "" {
		query = query.Joins("JOIN player_videos ON player_videos.video_id = videos.id").
			Where("player_videos.player_id = ?", playerID)
	}
	if search := c.Query("q"); search != "" {
		query = query.Where("title ILIKE ?", "%"+search+"%")
	}

	var total int64
	query.Count(&total)

	var videos []domain.Video
	query.Preload("Players").Preload("Tournament").
		Offset(offset).Limit(limit).Order("created_at DESC").Find(&videos)

	// Convert to response format
	var responses []VideoResponse
	for _, v := range videos {
		responses = append(responses, m.videoToResponse(v))
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"videos": responses,
			"pagination": gin.H{
				"page":        page,
				"limit":       limit,
				"total":       total,
				"total_pages": (total + int64(limit) - 1) / int64(limit),
			},
		},
	})
}

// GetVideo gets a single video by ID
func (m *MediaModule) GetVideo(c *gin.Context) {
	videoID := c.Param("id")
	vid, err := uuid.Parse(videoID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_UUID", "message": "Invalid video ID"}})
		return
	}

	var video domain.Video
	if err := m.db.Preload("Players").Preload("Tournament").First(&video, "id = ?", vid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "Video not found"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    m.videoToResponse(video),
	})
}

// UpdateVideoRequest updates video metadata
type UpdateVideoRequest struct {
	Title           *string  `json:"title,omitempty"`
	Description     *string  `json:"description,omitempty"`
	ThumbnailURL    *string  `json:"thumbnail_url,omitempty"`
	DurationSeconds *int     `json:"duration_seconds,omitempty"`
	TournamentID    *string  `json:"tournament_id,omitempty"`
	PlayerIDs       []string `json:"player_ids,omitempty"`
}

// UpdateVideo updates video metadata
func (m *MediaModule) UpdateVideo(c *gin.Context) {
	videoID := c.Param("id")
	vid, err := uuid.Parse(videoID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_UUID", "message": "Invalid video ID"}})
		return
	}

	var video domain.Video
	if err := m.db.First(&video, "id = ?", vid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "Video not found"}})
		return
	}

	var req UpdateVideoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
		return
	}

	updates := map[string]interface{}{
		"updated_at": time.Now(),
	}
	if req.Title != nil {
		updates["title"] = *req.Title
	}
	if req.Description != nil {
		updates["description"] = req.Description
	}
	if req.ThumbnailURL != nil {
		updates["thumbnail_url"] = req.ThumbnailURL
	}
	if req.DurationSeconds != nil {
		updates["duration_seconds"] = *req.DurationSeconds
	}
	if req.TournamentID != nil {
		if *req.TournamentID == "" {
			updates["tournament_id"] = nil
		} else {
			tid, _ := uuid.Parse(*req.TournamentID)
			updates["tournament_id"] = tid
		}
	}

	if err := m.db.Model(&video).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "UPDATE_FAILED", "message": "Failed to update video"}})
		return
	}

	// Update player links if provided
	if req.PlayerIDs != nil {
		// Remove existing links
		m.db.Where("video_id = ?", vid).Delete(&domain.PlayerVideo{})
		// Add new links
		for i, pidStr := range req.PlayerIDs {
			if pid, err := uuid.Parse(pidStr); err == nil {
				pv := domain.PlayerVideo{
					PlayerID:  pid,
					VideoID:   vid,
					IsPrimary: i == 0,
				}
				m.db.Create(&pv)
			}
		}
	}

	// Reload for response
	m.db.Preload("Players").First(&video, "id = ?", vid)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    m.videoToResponse(video),
	})
}

// DeleteVideo deletes a video
func (m *MediaModule) DeleteVideo(c *gin.Context) {
	videoID := c.Param("id")
	vid, err := uuid.Parse(videoID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_UUID", "message": "Invalid video ID"}})
		return
	}

	var video domain.Video
	if err := m.db.First(&video, "id = ?", vid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "Video not found"}})
		return
	}

	// Delete player-video links
	m.db.Where("video_id = ?", vid).Delete(&domain.PlayerVideo{})

	// Delete video record
	if err := m.db.Delete(&video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "DELETE_FAILED", "message": "Failed to delete video"}})
		return
	}

	// TODO: Delete from S3 asynchronously
	// go m.deleteFromS3(video.BlobURL)

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Video deleted successfully"})
}

// =============================================================================
// VIDEO REVIEW ENDPOINTS
// =============================================================================

// ApproveVideo approves a video
func (m *MediaModule) ApproveVideo(c *gin.Context) {
	videoID := c.Param("id")
	vid, err := uuid.Parse(videoID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_UUID", "message": "Invalid video ID"}})
		return
	}

	adminID := c.MustGet("user_id").(uuid.UUID)
	now := time.Now()

	result := m.db.Model(&domain.Video{}).Where("id = ?", vid).Updates(map[string]interface{}{
		"status":      "approved",
		"reviewed_by": adminID,
		"reviewed_at": now,
		"updated_at":  now,
	})

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "Video not found"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Video approved"})
}

// RejectVideoRequest contains rejection reason
type RejectVideoRequest struct {
	Reason string `json:"reason" binding:"required"`
}

// RejectVideo rejects a video
func (m *MediaModule) RejectVideo(c *gin.Context) {
	videoID := c.Param("id")
	vid, err := uuid.Parse(videoID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_UUID", "message": "Invalid video ID"}})
		return
	}

	var req RejectVideoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// Allow rejection without reason
		req.Reason = ""
	}

	adminID := c.MustGet("user_id").(uuid.UUID)
	now := time.Now()

	updates := map[string]interface{}{
		"status":      "rejected",
		"reviewed_by": adminID,
		"reviewed_at": now,
		"updated_at":  now,
	}
	if req.Reason != "" {
		updates["review_notes"] = req.Reason
	}

	result := m.db.Model(&domain.Video{}).Where("id = ?", vid).Updates(updates)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "Video not found"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Video rejected"})
}

// GetVideoStats returns video statistics
func (m *MediaModule) GetVideoStats(c *gin.Context) {
	var total, pending, approved, rejected, highlights, fullMatches int64

	m.db.Model(&domain.Video{}).Count(&total)
	m.db.Model(&domain.Video{}).Where("status = ?", "pending").Count(&pending)
	m.db.Model(&domain.Video{}).Where("status = ?", "approved").Count(&approved)
	m.db.Model(&domain.Video{}).Where("status = ?", "rejected").Count(&rejected)
	m.db.Model(&domain.Video{}).Where("video_type = ?", "highlight").Count(&highlights)
	m.db.Model(&domain.Video{}).Where("video_type = ?", "full_match").Count(&fullMatches)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"total":        total,
			"pending":      pending,
			"approved":     approved,
			"rejected":     rejected,
			"highlights":   highlights,
			"full_matches": fullMatches,
		},
	})
}

// =============================================================================
// STREAMING ENDPOINTS
// =============================================================================

// GetVideoStreamURL gets a streaming URL for a video
func (m *MediaModule) GetVideoStreamURL(c *gin.Context) {
	videoID := c.Param("id")
	vid, err := uuid.Parse(videoID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_UUID", "message": "Invalid video ID"}})
		return
	}

	var video domain.Video
	if err := m.db.First(&video, "id = ?", vid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "Video not found"}})
		return
	}

	// Check subscription for full match access
	if video.VideoType == "full_match" {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": gin.H{"code": "AUTH_REQUIRED", "message": "Authentication required for full matches"}})
			return
		}

		var sub domain.Subscription
		if err := m.db.Where("user_id = ?", userID).First(&sub).Error; err != nil || !sub.CanAccessFullMatch() {
			c.JSON(http.StatusForbidden, gin.H{"success": false, "error": gin.H{"code": "SUBSCRIPTION_REQUIRED", "message": "Upgrade to Scout tier or above to access full matches"}})
			return
		}
	}

	// Generate streaming URL
	var streamURL string
	if m.cloudFrontURL != "" {
		// Use CloudFront
		key := strings.TrimPrefix(video.BlobURL, fmt.Sprintf("s3://%s/", m.s3Bucket))
		streamURL = fmt.Sprintf("%s/%s", m.cloudFrontURL, key)
	} else if m.s3Client != nil {
		// Generate presigned GET URL
		presignClient := s3.NewPresignClient(m.s3Client)
		key := strings.TrimPrefix(video.BlobURL, fmt.Sprintf("s3://%s/", m.s3Bucket))
		presignedReq, err := presignClient.PresignGetObject(context.Background(), &s3.GetObjectInput{
			Bucket: aws.String(m.s3Bucket),
			Key:    aws.String(key),
		}, s3.WithPresignExpires(4*time.Hour))
		if err == nil {
			streamURL = presignedReq.URL
		}
	} else {
		streamURL = video.BlobURL
	}

	// Track view
	go m.trackVideoView(vid, c)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"stream_url": streamURL,
			"expires_in": 14400, // 4 hours
		},
	})
}

// =============================================================================
// PUBLIC ENDPOINTS
// =============================================================================

// ListHighlights lists approved highlight videos (FREE for everyone)
func (m *MediaModule) ListHighlights(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if limit > 100 {
		limit = 100
	}
	offset := (page - 1) * limit

	query := m.db.Model(&domain.Video{}).
		Where("video_type = ? AND status = ?", "highlight", "approved")

	if playerID := c.Query("player_id"); playerID != "" {
		query = query.Joins("JOIN player_videos ON player_videos.video_id = videos.id").
			Where("player_videos.player_id = ?", playerID)
	}
	if tournamentID := c.Query("tournament_id"); tournamentID != "" {
		query = query.Where("tournament_id = ?", tournamentID)
	}

	var total int64
	query.Count(&total)

	var videos []domain.Video
	query.Preload("Players").Preload("Tournament").
		Offset(offset).Limit(limit).Order("created_at DESC").Find(&videos)

	var responses []VideoResponse
	for _, v := range videos {
		responses = append(responses, m.videoToResponse(v))
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"videos": responses,
			"pagination": gin.H{
				"page":        page,
				"limit":       limit,
				"total":       total,
				"total_pages": (total + int64(limit) - 1) / int64(limit),
			},
		},
	})
}

// ListFullMatches lists full match videos (Scout+ tier required)
func (m *MediaModule) ListFullMatches(c *gin.Context) {
	// Check subscription tier
	userID, _ := c.Get("user_id")

	var sub domain.Subscription
	if err := m.db.Where("user_id = ?", userID).First(&sub).Error; err != nil || !sub.CanAccessFullMatch() {
		c.JSON(http.StatusPaymentRequired, gin.H{
			"success": false,
			"error": gin.H{
				"code":        "SUBSCRIPTION_REQUIRED",
				"message":     "Upgrade to Scout tier to access full matches",
				"upgrade_url": "/subscribe",
			},
		})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if limit > 100 {
		limit = 100
	}
	offset := (page - 1) * limit

	query := m.db.Model(&domain.Video{}).
		Where("video_type = ? AND status = ?", "full_match", "approved")

	if playerID := c.Query("player_id"); playerID != "" {
		query = query.Joins("JOIN player_videos ON player_videos.video_id = videos.id").
			Where("player_videos.player_id = ?", playerID)
	}
	if tournamentID := c.Query("tournament_id"); tournamentID != "" {
		query = query.Where("tournament_id = ?", tournamentID)
	}

	var total int64
	query.Count(&total)

	var videos []domain.Video
	query.Preload("Players").Preload("Tournament").
		Offset(offset).Limit(limit).Order("created_at DESC").Find(&videos)

	var responses []VideoResponse
	for _, v := range videos {
		responses = append(responses, m.videoToResponse(v))
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"videos": responses,
			"pagination": gin.H{
				"page":        page,
				"limit":       limit,
				"total":       total,
				"total_pages": (total + int64(limit) - 1) / int64(limit),
			},
		},
	})
}

// =============================================================================
// HELPER FUNCTIONS
// =============================================================================

func (m *MediaModule) trackVideoView(videoID uuid.UUID, c *gin.Context) {
	var viewerID *uuid.UUID
	if uid, exists := c.Get("user_id"); exists {
		id := uid.(uuid.UUID)
		viewerID = &id
	}

	view := domain.VideoView{
		VideoID:   videoID,
		ViewerID:  viewerID,
		CreatedAt: time.Now(),
	}
	m.db.Create(&view)
}

func sanitizeFileName(name string) string {
	// Remove any path components
	if idx := strings.LastIndex(name, "/"); idx >= 0 {
		name = name[idx+1:]
	}
	if idx := strings.LastIndex(name, "\\"); idx >= 0 {
		name = name[idx+1:]
	}
	// Replace spaces with underscores
	name = strings.ReplaceAll(name, " ", "_")
	return name
}

// Legacy compatibility - keeping old function signatures
func (m *MediaModule) GetUploadURL(c *gin.Context) {
	m.InitUpload(c)
}

func (m *MediaModule) ConfirmUpload(c *gin.Context) {
	videoID := c.Param("id")

	var req struct {
		VideoType       string   `json:"video_type" binding:"required"`
		Title           string   `json:"title" binding:"required"`
		Description     *string  `json:"description,omitempty"`
		ThumbnailURL    *string  `json:"thumbnail_url,omitempty"`
		DurationSeconds *int     `json:"duration_seconds,omitempty"`
		TournamentID    *string  `json:"tournament_id,omitempty"`
		PlayerIDs       []string `json:"player_ids,omitempty"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
		return
	}

	vid, _ := uuid.Parse(videoID)
	var video domain.Video
	if err := m.db.First(&video, "id = ?", vid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "Video not found"}})
		return
	}

	updates := map[string]interface{}{
		"video_type": req.VideoType,
		"title":      req.Title,
		"status":     "approved",
		"updated_at": time.Now(),
	}
	if req.Description != nil {
		updates["description"] = req.Description
	}
	if req.ThumbnailURL != nil {
		updates["thumbnail_url"] = req.ThumbnailURL
	}
	if req.DurationSeconds != nil {
		updates["duration_seconds"] = req.DurationSeconds
	}
	if req.TournamentID != nil && *req.TournamentID != "" {
		tid, _ := uuid.Parse(*req.TournamentID)
		updates["tournament_id"] = tid
	}

	m.db.Model(&video).Updates(updates)

	// Link players
	for i, pidStr := range req.PlayerIDs {
		if pid, err := uuid.Parse(pidStr); err == nil {
			pv := domain.PlayerVideo{
				PlayerID:  pid,
				VideoID:   vid,
				IsPrimary: i == 0,
			}
			m.db.Create(&pv)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"id":         video.ID,
			"video_type": req.VideoType,
			"title":      req.Title,
		},
	})
}

func (m *MediaModule) LinkPlayerToVideo(c *gin.Context) {
	videoID := c.Param("id")
	vid, _ := uuid.Parse(videoID)

	var req struct {
		PlayerID  string `json:"player_id" binding:"required"`
		IsPrimary bool   `json:"is_primary"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
		return
	}

	pid, _ := uuid.Parse(req.PlayerID)
	pv := domain.PlayerVideo{
		PlayerID:  pid,
		VideoID:   vid,
		IsPrimary: req.IsPrimary,
	}

	if err := m.db.Create(&pv).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "LINK_FAILED", "message": "Failed to link player"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Player linked"})
}
