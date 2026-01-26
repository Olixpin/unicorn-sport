package media

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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
	}
}

// --- Upload Requests ---

// UploadRequest represents a request for presigned upload URL
type UploadRequest struct {
	FileName    string `json:"file_name" binding:"required"`
	ContentType string `json:"content_type" binding:"required"`
	FileSize    int64  `json:"file_size" binding:"required"`
}

// UploadResponse contains the presigned URL and video ID
type UploadResponse struct {
	UploadURL string    `json:"upload_url"`
	VideoID   uuid.UUID `json:"video_id"`
	ExpiresIn int       `json:"expires_in"`
}

// GetUploadURL generates a presigned S3 URL for video upload
func (m *MediaModule) GetUploadURL(c *gin.Context) {
	var req UploadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
		return
	}

	// Validate content type
	allowedTypes := map[string]bool{
		"video/mp4":       true,
		"video/quicktime": true,
		"video/x-msvideo": true,
		"video/webm":      true,
	}
	if !allowedTypes[req.ContentType] {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_CONTENT_TYPE", "message": "Only video files are allowed"}})
		return
	}

	// Validate file size (max 5GB)
	maxSize := int64(5 * 1024 * 1024 * 1024)
	if req.FileSize > maxSize {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "FILE_TOO_LARGE", "message": "Maximum file size is 5GB"}})
		return
	}

	adminID, _ := c.Get("user_id")

	// Generate unique video ID
	videoID := uuid.New()
	objectKey := fmt.Sprintf("videos/%s/%s", videoID.String(), req.FileName)

	// If S3 is not configured, return mock URL for development
	if m.s3Client == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": UploadResponse{
				UploadURL: fmt.Sprintf("http://localhost:9000/%s/%s", m.s3Bucket, objectKey),
				VideoID:   videoID,
				ExpiresIn: 3600,
			},
		})
		return
	}

	// Generate presigned URL
	presignClient := s3.NewPresignClient(m.s3Client)
	presignedReq, err := presignClient.PresignPutObject(context.Background(), &s3.PutObjectInput{
		Bucket:        aws.String(m.s3Bucket),
		Key:           aws.String(objectKey),
		ContentType:   aws.String(req.ContentType),
		ContentLength: aws.Int64(req.FileSize),
	}, s3.WithPresignExpires(time.Hour))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "PRESIGN_FAILED", "message": "Failed to generate upload URL"}})
		return
	}

	// Create pending video record
	blobURL := fmt.Sprintf("s3://%s/%s", m.s3Bucket, objectKey)
	video := domain.Video{
		ID:            videoID,
		VideoType:     "pending", // Will be updated when confirmed
		Title:         req.FileName,
		BlobURL:       blobURL,
		FileSizeBytes: &req.FileSize,
		UploadedBy:    adminID.(uuid.UUID),
		CreatedAt:     time.Now(),
	}

	if err := m.db.Create(&video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "CREATE_FAILED", "message": "Failed to create video record"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": UploadResponse{
			UploadURL: presignedReq.URL,
			VideoID:   videoID,
			ExpiresIn: 3600,
		},
	})
}

// ConfirmUploadRequest confirms a video upload and sets metadata
type ConfirmUploadRequest struct {
	VideoType       string   `json:"video_type" binding:"required,oneof=highlight full_match"`
	Title           string   `json:"title" binding:"required"`
	Description     *string  `json:"description,omitempty"`
	TournamentID    *string  `json:"tournament_id,omitempty"`
	PlayerIDs       []string `json:"player_ids,omitempty"`
	DurationSeconds *int     `json:"duration_seconds,omitempty"`
	ThumbnailURL    *string  `json:"thumbnail_url,omitempty"`
}

// ConfirmUpload confirms the upload and updates video metadata
func (m *MediaModule) ConfirmUpload(c *gin.Context) {
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

	var req ConfirmUploadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
		return
	}

	// Parse tournament ID
	var tournamentID *uuid.UUID
	if req.TournamentID != nil && *req.TournamentID != "" {
		tid, err := uuid.Parse(*req.TournamentID)
		if err == nil {
			tournamentID = &tid
		}
	}

	// Update video
	updates := map[string]interface{}{
		"video_type":    req.VideoType,
		"title":         req.Title,
		"description":   req.Description,
		"tournament_id": tournamentID,
	}
	if req.DurationSeconds != nil {
		updates["duration_seconds"] = *req.DurationSeconds
	}
	if req.ThumbnailURL != nil {
		updates["thumbnail_url"] = *req.ThumbnailURL
	}

	if err := m.db.Model(&video).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "UPDATE_FAILED", "message": "Failed to update video"}})
		return
	}

	// Link players to video
	if len(req.PlayerIDs) > 0 {
		for i, pidStr := range req.PlayerIDs {
			if pid, err := uuid.Parse(pidStr); err == nil {
				pv := domain.PlayerVideo{
					PlayerID:  pid,
					VideoID:   vid,
					IsPrimary: i == 0, // First player is primary
				}
				m.db.Create(&pv)
			}
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

// --- Video CRUD ---

// CreateVideoRequest for creating video without upload
type CreateVideoRequest struct {
	VideoType       string   `json:"video_type" binding:"required,oneof=highlight full_match"`
	Title           string   `json:"title" binding:"required"`
	Description     *string  `json:"description,omitempty"`
	BlobURL         string   `json:"blob_url" binding:"required"`
	ThumbnailURL    *string  `json:"thumbnail_url,omitempty"`
	DurationSeconds *int     `json:"duration_seconds,omitempty"`
	TournamentID    *string  `json:"tournament_id,omitempty"`
	PlayerIDs       []string `json:"player_ids,omitempty"`
}

// CreateVideo creates a video record (for already-uploaded content)
func (m *MediaModule) CreateVideo(c *gin.Context) {
	var req CreateVideoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
		return
	}

	adminID, _ := c.Get("user_id")

	var tournamentID *uuid.UUID
	if req.TournamentID != nil && *req.TournamentID != "" {
		tid, err := uuid.Parse(*req.TournamentID)
		if err == nil {
			tournamentID = &tid
		}
	}

	video := domain.Video{
		VideoType:       req.VideoType,
		Title:           req.Title,
		Description:     req.Description,
		BlobURL:         req.BlobURL,
		ThumbnailURL:    req.ThumbnailURL,
		DurationSeconds: req.DurationSeconds,
		TournamentID:    tournamentID,
		UploadedBy:      adminID.(uuid.UUID),
		CreatedAt:       time.Now(),
	}

	if err := m.db.Create(&video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "CREATE_FAILED", "message": "Failed to create video"}})
		return
	}

	// Link players
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

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data": gin.H{
			"id":         video.ID,
			"title":      video.Title,
			"video_type": video.VideoType,
			"created_at": video.CreatedAt,
		},
	})
}

// ListVideos lists all videos (admin view)
func (m *MediaModule) ListVideos(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset := (page - 1) * limit

	var videos []domain.Video
	var total int64

	query := m.db.Model(&domain.Video{})

	// Apply filters
	if videoType := c.Query("video_type"); videoType != "" {
		query = query.Where("video_type = ?", videoType)
	}
	if tournamentID := c.Query("tournament_id"); tournamentID != "" {
		query = query.Where("tournament_id = ?", tournamentID)
	}
	if playerID := c.Query("player_id"); playerID != "" {
		query = query.Joins("JOIN player_videos ON player_videos.video_id = videos.id").
			Where("player_videos.player_id = ?", playerID)
	}

	query.Count(&total)
	query.Preload("Tournament").Offset(offset).Limit(limit).Order("created_at DESC").Find(&videos)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"videos": videos,
			"pagination": gin.H{
				"page":        page,
				"limit":       limit,
				"total":       total,
				"total_pages": (total + int64(limit) - 1) / int64(limit),
			},
		},
	})
}

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
		streamURL = fmt.Sprintf("%s/videos/%s", m.cloudFrontURL, vid.String())
	} else if m.s3Client != nil {
		// Generate presigned GET URL
		presignClient := s3.NewPresignClient(m.s3Client)
		// Extract key from blob URL
		key := video.BlobURL[len(fmt.Sprintf("s3://%s/", m.s3Bucket)):]
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

	// TODO: Delete from S3 (can be done async)

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Video deleted successfully"})
}

// LinkPlayerToVideo links a player to a video
func (m *MediaModule) LinkPlayerToVideo(c *gin.Context) {
	videoID := c.Param("id")
	vid, err := uuid.Parse(videoID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_UUID", "message": "Invalid video ID"}})
		return
	}

	var req struct {
		PlayerID  string `json:"player_id" binding:"required"`
		IsPrimary bool   `json:"is_primary"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
		return
	}

	pid, err := uuid.Parse(req.PlayerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_UUID", "message": "Invalid player ID"}})
		return
	}

	pv := domain.PlayerVideo{
		PlayerID:  pid,
		VideoID:   vid,
		IsPrimary: req.IsPrimary,
	}

	if err := m.db.Create(&pv).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "LINK_FAILED", "message": "Failed to link player to video"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Player linked to video"})
}

// --- Helpers ---

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

// --- Public Endpoints ---

// ListHighlights lists highlight videos (FREE for everyone)
func (m *MediaModule) ListHighlights(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset := (page - 1) * limit

	var videos []domain.Video
	var total int64

	query := m.db.Model(&domain.Video{}).Where("video_type = ?", "highlight")

	if playerID := c.Query("player_id"); playerID != "" {
		query = query.Joins("JOIN player_videos ON player_videos.video_id = videos.id").
			Where("player_videos.player_id = ?", playerID)
	}
	if tournamentID := c.Query("tournament_id"); tournamentID != "" {
		query = query.Where("tournament_id = ?", tournamentID)
	}

	query.Count(&total)
	query.Preload("Tournament").Offset(offset).Limit(limit).Order("created_at DESC").Find(&videos)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"videos": videos,
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
	offset := (page - 1) * limit

	var videos []domain.Video
	var total int64

	query := m.db.Model(&domain.Video{}).Where("video_type = ?", "full_match")

	if playerID := c.Query("player_id"); playerID != "" {
		query = query.Joins("JOIN player_videos ON player_videos.video_id = videos.id").
			Where("player_videos.player_id = ?", playerID)
	}
	if tournamentID := c.Query("tournament_id"); tournamentID != "" {
		query = query.Where("tournament_id = ?", tournamentID)
	}

	query.Count(&total)
	query.Preload("Tournament").Offset(offset).Limit(limit).Order("created_at DESC").Find(&videos)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"videos": videos,
			"pagination": gin.H{
				"page":        page,
				"limit":       limit,
				"total":       total,
				"total_pages": (total + int64(limit) - 1) / int64(limit),
			},
		},
	})
}
