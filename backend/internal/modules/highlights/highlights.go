package highlights

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/unicorn-sport/backend/internal/domain"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Module holds dependencies for the highlights module
type Module struct {
	DB       *gorm.DB
	S3Client *s3.Client
	S3Bucket string
	CDNHost  string
}

// NewModule creates a new highlights module
func NewModule(db *gorm.DB, s3Client *s3.Client, bucket, cdnHost string) *Module {
	return &Module{
		DB:       db,
		S3Client: s3Client,
		S3Bucket: bucket,
		CDNHost:  cdnHost,
	}
}

// ==================== HIGHLIGHT CRUD ====================

// CreateHighlightRequest is the request for creating a highlight
type CreateHighlightRequest struct {
	PlayerID         string `json:"player_id" binding:"required"`
	MatchID          string `json:"match_id"` // Optional
	HighlightType    string `json:"highlight_type" binding:"required"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	TimestampInMatch int    `json:"timestamp_in_match"`
}

// InitHighlightUpload initializes upload for a player highlight
func (m *Module) InitHighlightUpload(c *gin.Context) {
	var req struct {
		PlayerID    string `json:"player_id" binding:"required"`
		MatchID     string `json:"match_id"` // Optional
		FileName    string `json:"file_name" binding:"required"`
		FileSize    int64  `json:"file_size" binding:"required"`
		ContentType string `json:"content_type" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request", "error": err.Error()})
		return
	}

	playerID, err := uuid.Parse(req.PlayerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid player ID"})
		return
	}

	// Verify player exists
	var player domain.Player
	if err := m.DB.First(&player, "id = ?", playerID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Player not found"})
		return
	}

	// Validate file size (max 1GB for highlights)
	maxSize := int64(1 * 1024 * 1024 * 1024)
	if req.FileSize > maxSize {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "File too large. Maximum 1GB for highlights."})
		return
	}

	userID, _ := c.Get("user_id")

	// Generate S3 key
	s3Key := fmt.Sprintf("highlights/%s/%s/%s", playerID.String(), uuid.New().String()[:8], req.FileName)

	// Generate presigned URL
	presigner := s3.NewPresignClient(m.S3Client)
	presignResult, err := presigner.PresignPutObject(c.Request.Context(), &s3.PutObjectInput{
		Bucket:      aws.String(m.S3Bucket),
		Key:         aws.String(s3Key),
		ContentType: aws.String(req.ContentType),
	}, s3.WithPresignExpires(1*time.Hour))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to generate upload URL"})
		return
	}

	// Create session
	var matchID *uuid.UUID
	if req.MatchID != "" {
		if mid, err := uuid.Parse(req.MatchID); err == nil {
			matchID = &mid
		}
	}

	session := domain.UploadSession{
		UploadType:  "highlight",
		ContentType: req.ContentType,
		FileName:    req.FileName,
		FileSize:    req.FileSize,
		S3Key:       s3Key,
		Status:      "pending",
		EntityType:  stringPtr("highlight"),
		EntityID:    &playerID,
		UploadedBy:  userID.(uuid.UUID),
		ExpiresAt:   time.Now().Add(1 * time.Hour),
	}

	if err := m.DB.Create(&session).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to create session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"session_id": session.ID,
			"upload_url": presignResult.URL,
			"s3_key":     s3Key,
			"expires_in": 3600,
			"match_id":   matchID,
		},
	})
}

// CreateHighlight creates a highlight record after upload
func (m *Module) CreateHighlight(c *gin.Context) {
	var req struct {
		PlayerID         string `json:"player_id" binding:"required"`
		MatchID          string `json:"match_id"`
		HighlightType    string `json:"highlight_type" binding:"required"`
		S3Key            string `json:"s3_key" binding:"required"`
		ThumbnailURL     string `json:"thumbnail_url"`
		DurationSeconds  int    `json:"duration_seconds"`
		FileSizeBytes    int64  `json:"file_size_bytes"`
		Title            string `json:"title"`
		Description      string `json:"description"`
		TimestampInMatch int    `json:"timestamp_in_match"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request", "error": err.Error()})
		return
	}

	// Validate highlight type
	if !domain.IsValidHighlightType(req.HighlightType) {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":     false,
			"message":     "Invalid highlight type",
			"valid_types": domain.ValidHighlightTypes(),
		})
		return
	}

	playerID, err := uuid.Parse(req.PlayerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid player ID"})
		return
	}

	// Verify player exists
	var player domain.Player
	if err := m.DB.First(&player, "id = ?", playerID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Player not found"})
		return
	}

	var matchID *uuid.UUID
	if req.MatchID != "" {
		if mid, err := uuid.Parse(req.MatchID); err == nil {
			matchID = &mid
		}
	}

	userID, _ := c.Get("user_id")

	highlight := domain.PlayerHighlight{
		PlayerID:         playerID,
		MatchID:          matchID,
		HighlightType:    req.HighlightType,
		VideoURL:         req.S3Key,
		ThumbnailURL:     stringPtr(req.ThumbnailURL),
		DurationSeconds:  intPtr(req.DurationSeconds),
		FileSizeBytes:    int64Ptr(req.FileSizeBytes),
		Title:            stringPtr(req.Title),
		Description:      stringPtr(req.Description),
		TimestampInMatch: intPtr(req.TimestampInMatch),
		Status:           "approved", // Auto-approve for admin uploads
		UploadedBy:       userID.(uuid.UUID),
	}

	if err := m.DB.Create(&highlight).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to create highlight"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Highlight created successfully",
		"data":    highlight,
	})
}

// ListPlayerHighlights lists all highlights for a player
func (m *Module) ListPlayerHighlights(c *gin.Context) {
	playerID := c.Param("id")
	pid, err := uuid.Parse(playerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid player ID"})
		return
	}

	// Filters
	highlightType := c.Query("type")
	matchID := c.Query("match_id")

	query := m.DB.Where("player_id = ?", pid).Preload("Match")

	if highlightType != "" {
		query = query.Where("highlight_type = ?", highlightType)
	}
	if matchID != "" {
		if mid, err := uuid.Parse(matchID); err == nil {
			query = query.Where("match_id = ?", mid)
		}
	}

	var highlights []domain.PlayerHighlight
	if err := query.Order("created_at DESC").Find(&highlights).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to fetch highlights"})
		return
	}

	// Add CDN URLs for streaming
	type highlightResponse struct {
		domain.PlayerHighlight
		StreamURL string `json:"stream_url"`
	}

	response := make([]highlightResponse, len(highlights))
	for i, h := range highlights {
		response[i] = highlightResponse{
			PlayerHighlight: h,
			StreamURL:       m.getStreamURL(h.VideoURL),
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})
}

// ListMatchHighlights lists all highlights for a match
func (m *Module) ListMatchHighlights(c *gin.Context) {
	matchID := c.Param("id")
	mid, err := uuid.Parse(matchID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid match ID"})
		return
	}

	// Optional player filter
	playerID := c.Query("player_id")

	query := m.DB.Where("match_id = ?", mid).Preload("Player")

	if playerID != "" {
		if pid, err := uuid.Parse(playerID); err == nil {
			query = query.Where("player_id = ?", pid)
		}
	}

	var highlights []domain.PlayerHighlight
	if err := query.Order("timestamp_in_match ASC, created_at DESC").Find(&highlights).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to fetch highlights"})
		return
	}

	// Group by player
	type playerHighlights struct {
		PlayerID        uuid.UUID `json:"player_id"`
		PlayerName      string    `json:"player_name"`
		ProfilePhotoURL *string   `json:"profile_photo_url"`
		Highlights      []struct {
			ID               uuid.UUID `json:"id"`
			HighlightType    string    `json:"highlight_type"`
			Title            *string   `json:"title"`
			DurationSeconds  *int      `json:"duration_seconds"`
			ThumbnailURL     *string   `json:"thumbnail_url"`
			StreamURL        string    `json:"stream_url"`
			TimestampInMatch *int      `json:"timestamp_in_match"`
			ViewCount        int       `json:"view_count"`
		} `json:"highlights"`
	}

	playerMap := make(map[uuid.UUID]*playerHighlights)
	for _, h := range highlights {
		if _, exists := playerMap[h.PlayerID]; !exists {
			playerMap[h.PlayerID] = &playerHighlights{
				PlayerID:        h.PlayerID,
				PlayerName:      fmt.Sprintf("%s %s", h.Player.FirstName, h.Player.LastName),
				ProfilePhotoURL: h.Player.ProfilePhotoURL,
				Highlights: []struct {
					ID               uuid.UUID `json:"id"`
					HighlightType    string    `json:"highlight_type"`
					Title            *string   `json:"title"`
					DurationSeconds  *int      `json:"duration_seconds"`
					ThumbnailURL     *string   `json:"thumbnail_url"`
					StreamURL        string    `json:"stream_url"`
					TimestampInMatch *int      `json:"timestamp_in_match"`
					ViewCount        int       `json:"view_count"`
				}{},
			}
		}
		playerMap[h.PlayerID].Highlights = append(playerMap[h.PlayerID].Highlights, struct {
			ID               uuid.UUID `json:"id"`
			HighlightType    string    `json:"highlight_type"`
			Title            *string   `json:"title"`
			DurationSeconds  *int      `json:"duration_seconds"`
			ThumbnailURL     *string   `json:"thumbnail_url"`
			StreamURL        string    `json:"stream_url"`
			TimestampInMatch *int      `json:"timestamp_in_match"`
			ViewCount        int       `json:"view_count"`
		}{
			ID:               h.ID,
			HighlightType:    h.HighlightType,
			Title:            h.Title,
			DurationSeconds:  h.DurationSeconds,
			ThumbnailURL:     h.ThumbnailURL,
			StreamURL:        m.getStreamURL(h.VideoURL),
			TimestampInMatch: h.TimestampInMatch,
			ViewCount:        h.ViewCount,
		})
	}

	// Convert map to slice
	response := make([]*playerHighlights, 0, len(playerMap))
	for _, ph := range playerMap {
		response = append(response, ph)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})
}

// GetHighlight retrieves a single highlight
func (m *Module) GetHighlight(c *gin.Context) {
	highlightID := c.Param("id")
	hid, err := uuid.Parse(highlightID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid highlight ID"})
		return
	}

	var highlight domain.PlayerHighlight
	if err := m.DB.Preload("Player").Preload("Match").First(&highlight, "id = ?", hid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Highlight not found"})
		return
	}

	// Record view (if not admin viewing)
	userID, hasUser := c.Get("user_id")
	source := c.Query("source")

	view := domain.HighlightView{
		HighlightID: hid,
		Source:      stringPtr(source),
	}
	if hasUser {
		uid := userID.(uuid.UUID)
		view.ViewerID = &uid
	}
	m.DB.Create(&view) // Fire and forget

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"highlight":  highlight,
			"stream_url": m.getStreamURL(highlight.VideoURL),
		},
	})
}

// UpdateHighlight updates a highlight
func (m *Module) UpdateHighlight(c *gin.Context) {
	highlightID := c.Param("id")
	hid, err := uuid.Parse(highlightID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid highlight ID"})
		return
	}

	var highlight domain.PlayerHighlight
	if err := m.DB.First(&highlight, "id = ?", hid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Highlight not found"})
		return
	}

	var req struct {
		HighlightType    string `json:"highlight_type"`
		Title            string `json:"title"`
		Description      string `json:"description"`
		TimestampInMatch *int   `json:"timestamp_in_match"`
		Status           string `json:"status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request"})
		return
	}

	updates := make(map[string]interface{})

	if req.HighlightType != "" {
		if !domain.IsValidHighlightType(req.HighlightType) {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid highlight type"})
			return
		}
		updates["highlight_type"] = req.HighlightType
	}
	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.TimestampInMatch != nil {
		updates["timestamp_in_match"] = *req.TimestampInMatch
	}
	if req.Status != "" {
		updates["status"] = req.Status
	}

	if err := m.DB.Model(&highlight).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to update"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Highlight updated",
		"data":    highlight,
	})
}

// DeleteHighlight deletes a highlight
func (m *Module) DeleteHighlight(c *gin.Context) {
	highlightID := c.Param("id")
	hid, err := uuid.Parse(highlightID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid highlight ID"})
		return
	}

	var highlight domain.PlayerHighlight
	if err := m.DB.First(&highlight, "id = ?", hid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Highlight not found"})
		return
	}

	// Delete from S3
	_, err = m.S3Client.DeleteObject(c.Request.Context(), &s3.DeleteObjectInput{
		Bucket: aws.String(m.S3Bucket),
		Key:    aws.String(highlight.VideoURL),
	})
	if err != nil {
		fmt.Printf("Warning: Failed to delete S3 object: %v\n", err)
	}

	if err := m.DB.Delete(&highlight).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Highlight deleted"})
}

// ==================== PUBLIC ENDPOINTS ====================

// GetPlayerHighlightsPublic returns highlights for a player (public, FREE)
func (m *Module) GetPlayerHighlightsPublic(c *gin.Context) {
	playerID := c.Param("id")
	pid, err := uuid.Parse(playerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid player ID"})
		return
	}

	// Filters
	highlightType := c.Query("type")
	limit := 20

	query := m.DB.Where("player_id = ? AND status = 'approved'", pid).
		Preload("Match").
		Limit(limit)

	if highlightType != "" {
		query = query.Where("highlight_type = ?", highlightType)
	}

	var highlights []domain.PlayerHighlight
	if err := query.Order("created_at DESC").Find(&highlights).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to fetch highlights"})
		return
	}

	// Build response with stream URLs
	type highlightResponse struct {
		ID               uuid.UUID  `json:"id"`
		HighlightType    string     `json:"highlight_type"`
		Title            *string    `json:"title"`
		Description      *string    `json:"description"`
		ThumbnailURL     *string    `json:"thumbnail_url"`
		StreamURL        string     `json:"stream_url"`
		DurationSeconds  *int       `json:"duration_seconds"`
		TimestampInMatch *int       `json:"timestamp_in_match"`
		ViewCount        int        `json:"view_count"`
		MatchID          *uuid.UUID `json:"match_id,omitempty"`
		MatchTitle       *string    `json:"match_title,omitempty"`
		CreatedAt        time.Time  `json:"created_at"`
	}

	response := make([]highlightResponse, len(highlights))
	for i, h := range highlights {
		hr := highlightResponse{
			ID:               h.ID,
			HighlightType:    h.HighlightType,
			Title:            h.Title,
			Description:      h.Description,
			ThumbnailURL:     h.ThumbnailURL,
			StreamURL:        m.getStreamURL(h.VideoURL),
			DurationSeconds:  h.DurationSeconds,
			TimestampInMatch: h.TimestampInMatch,
			ViewCount:        h.ViewCount,
			MatchID:          h.MatchID,
			CreatedAt:        h.CreatedAt,
		}
		if h.Match != nil {
			hr.MatchTitle = &h.Match.Title
		}
		response[i] = hr
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})
}

// GetPlayerTournamentAppearances returns tournaments/matches a player participated in
func (m *Module) GetPlayerTournamentAppearances(c *gin.Context) {
	playerID := c.Param("id")
	pid, err := uuid.Parse(playerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid player ID"})
		return
	}

	// Get all matches this player participated in
	var matchPlayers []domain.MatchPlayer
	if err := m.DB.Where("player_id = ?", pid).Find(&matchPlayers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to fetch data"})
		return
	}

	if len(matchPlayers) == 0 {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": []interface{}{}})
		return
	}

	// Get match IDs
	matchIDs := make([]uuid.UUID, len(matchPlayers))
	for i, mp := range matchPlayers {
		matchIDs[i] = mp.MatchID
	}

	// Get matches with tournaments and videos
	var matches []domain.Match
	m.DB.Where("id IN ?", matchIDs).
		Preload("Tournament").
		Preload("Video").
		Order("match_date DESC").
		Find(&matches)

	// Get highlight counts per match
	type matchHighlightCount struct {
		MatchID        uuid.UUID
		HighlightCount int
	}
	var highlightCounts []matchHighlightCount
	m.DB.Model(&domain.PlayerHighlight{}).
		Select("match_id, COUNT(*) as highlight_count").
		Where("player_id = ? AND match_id IN ?", pid, matchIDs).
		Group("match_id").
		Scan(&highlightCounts)

	highlightMap := make(map[uuid.UUID]int)
	for _, hc := range highlightCounts {
		highlightMap[hc.MatchID] = hc.HighlightCount
	}

	// Get player stats per match
	statsMap := make(map[uuid.UUID]domain.MatchPlayer)
	for _, mp := range matchPlayers {
		statsMap[mp.MatchID] = mp
	}

	// Group by tournament
	type matchInTournament struct {
		MatchID         uuid.UUID `json:"match_id"`
		Title           string    `json:"title"`
		MatchDate       time.Time `json:"match_date"`
		Stage           *string   `json:"stage,omitempty"`
		Goals           int       `json:"goals"`
		Assists         int       `json:"assists"`
		MinutesPlayed   *int      `json:"minutes_played,omitempty"`
		HighlightCount  int       `json:"highlight_count"`
		HasFullMatch    bool      `json:"has_full_match"`
		FullMatchLocked bool      `json:"full_match_locked"` // For paywall UI
	}

	type tournamentAppearance struct {
		TournamentID    uuid.UUID           `json:"tournament_id"`
		TournamentName  string              `json:"tournament_name"`
		Year            int                 `json:"year"`
		TotalMatches    int                 `json:"total_matches"`
		TotalGoals      int                 `json:"total_goals"`
		TotalAssists    int                 `json:"total_assists"`
		TotalHighlights int                 `json:"total_highlights"`
		Matches         []matchInTournament `json:"matches"`
	}

	tournamentMap := make(map[uuid.UUID]*tournamentAppearance)
	for _, match := range matches {
		if match.TournamentID == nil {
			continue
		}

		tid := *match.TournamentID
		if _, exists := tournamentMap[tid]; !exists {
			tournamentMap[tid] = &tournamentAppearance{
				TournamentID:   tid,
				TournamentName: match.Tournament.Name,
				Year:           match.Tournament.Year,
				Matches:        []matchInTournament{},
			}
		}

		stats := statsMap[match.ID]
		mi := matchInTournament{
			MatchID:         match.ID,
			Title:           match.Title,
			MatchDate:       match.MatchDate,
			Stage:           match.Stage,
			Goals:           stats.Goals,
			Assists:         stats.Assists,
			MinutesPlayed:   stats.MinutesPlayed,
			HighlightCount:  highlightMap[match.ID],
			HasFullMatch:    match.Video != nil && match.Video.Status == "ready",
			FullMatchLocked: true, // Will be unlocked by subscription check on frontend
		}

		tournamentMap[tid].Matches = append(tournamentMap[tid].Matches, mi)
		tournamentMap[tid].TotalMatches++
		tournamentMap[tid].TotalGoals += stats.Goals
		tournamentMap[tid].TotalAssists += stats.Assists
		tournamentMap[tid].TotalHighlights += highlightMap[match.ID]
	}

	// Convert to slice
	response := make([]*tournamentAppearance, 0, len(tournamentMap))
	for _, ta := range tournamentMap {
		response = append(response, ta)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})
}

// GetHighlightTypes returns all valid highlight types
func (m *Module) GetHighlightTypes(c *gin.Context) {
	types := []struct {
		Value     string   `json:"value"`
		Label     string   `json:"label"`
		Icon      string   `json:"icon"`
		Positions []string `json:"positions"` // Suggested positions for this type
	}{
		{Value: "goal", Label: "Goal", Icon: "⚽", Positions: []string{"forward", "midfielder", "winger"}},
		{Value: "assist", Label: "Assist", Icon: "🎯", Positions: []string{"midfielder", "winger", "forward"}},
		{Value: "dribbling", Label: "Dribbling", Icon: "🦵", Positions: []string{"winger", "midfielder", "forward"}},
		{Value: "defending", Label: "Defending", Icon: "🛡️", Positions: []string{"defender", "defensive_midfielder"}},
		{Value: "tackling", Label: "Tackling", Icon: "🦶", Positions: []string{"defender", "midfielder"}},
		{Value: "passing", Label: "Passing", Icon: "➡️", Positions: []string{"midfielder", "defender"}},
		{Value: "shooting", Label: "Shooting", Icon: "🎯", Positions: []string{"forward", "midfielder"}},
		{Value: "heading", Label: "Heading", Icon: "🗣️", Positions: []string{"defender", "forward"}},
		{Value: "speed", Label: "Speed", Icon: "⚡", Positions: []string{"winger", "fullback", "forward"}},
		{Value: "save", Label: "Save", Icon: "🧤", Positions: []string{"goalkeeper"}},
		{Value: "distribution", Label: "Distribution", Icon: "📤", Positions: []string{"goalkeeper"}},
		{Value: "positioning", Label: "Positioning", Icon: "📍", Positions: []string{"defender", "midfielder"}},
		{Value: "vision", Label: "Vision", Icon: "👁️", Positions: []string{"midfielder", "playmaker"}},
		{Value: "other", Label: "Other", Icon: "✨", Positions: []string{}},
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    types,
	})
}

// ==================== HELPERS ====================

func (m *Module) getStreamURL(s3Key string) string {
	if m.CDNHost != "" {
		return fmt.Sprintf("%s/%s", m.CDNHost, s3Key)
	}
	// Generate presigned URL for direct S3 access
	presigner := s3.NewPresignClient(m.S3Client)
	result, err := presigner.PresignGetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(m.S3Bucket),
		Key:    aws.String(s3Key),
	}, s3.WithPresignExpires(1*time.Hour))
	if err != nil {
		return ""
	}
	return result.URL
}

func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func intPtr(i int) *int {
	if i == 0 {
		return nil
	}
	return &i
}

func int64Ptr(i int64) *int64 {
	if i == 0 {
		return nil
	}
	return &i
}
