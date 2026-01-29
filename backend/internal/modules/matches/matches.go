package matches

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/unicorn-sport/backend/internal/domain"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Module holds dependencies for the matches module
type Module struct {
	DB       *gorm.DB
	S3Client *s3.Client
	S3Bucket string
	CDNHost  string // CloudFront or S3 URL for serving
}

// NewModule creates a new matches module
func NewModule(db *gorm.DB, s3Client *s3.Client, bucket, cdnHost string) *Module {
	return &Module{
		DB:       db,
		S3Client: s3Client,
		S3Bucket: bucket,
		CDNHost:  cdnHost,
	}
}

// getPresignedURL converts an s3:// URL to a presigned URL
func (m *Module) getPresignedURL(s3URL string) string {
	if s3URL == "" || !strings.HasPrefix(s3URL, "s3://") {
		return s3URL
	}

	// Parse s3://bucket/key format
	withoutPrefix := strings.TrimPrefix(s3URL, "s3://")
	parts := strings.SplitN(withoutPrefix, "/", 2)
	if len(parts) != 2 {
		return s3URL
	}

	key := parts[1]

	presignClient := s3.NewPresignClient(m.S3Client)
	presignedReq, err := presignClient.PresignGetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(m.S3Bucket),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(time.Hour))
	if err != nil {
		return s3URL
	}

	return presignedReq.URL
}

// ==================== MATCH CRUD ====================

// CreateMatchRequest is the request body for creating a match
type CreateMatchRequest struct {
	TournamentID string `json:"tournament_id" binding:"required"`
	Title        string `json:"title" binding:"required"`
	Description  string `json:"description"`
	MatchDate    string `json:"match_date" binding:"required"` // ISO date
	Location     string `json:"location"`
	Stage        string `json:"stage"` // group, quarterfinal, semifinal, final
	HomeTeam     string `json:"home_team"`
	AwayTeam     string `json:"away_team"`
	MatchNumber  int    `json:"match_number"`
}

// CreateMatch creates a new match in a tournament
func (m *Module) CreateMatch(c *gin.Context) {
	var req CreateMatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request", "error": err.Error()})
		return
	}

	tournamentID, err := uuid.Parse(req.TournamentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid tournament ID"})
		return
	}

	// Verify tournament exists
	var tournament domain.Tournament
	if err := m.DB.First(&tournament, "id = ?", tournamentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Tournament not found"})
		return
	}

	matchDate, err := time.Parse("2006-01-02", req.MatchDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid match date format (use YYYY-MM-DD)"})
		return
	}

	userID, _ := c.Get("user_id")

	match := domain.Match{
		TournamentID: &tournamentID,
		Title:        req.Title,
		Description:  stringPtr(req.Description),
		MatchDate:    matchDate,
		Location:     stringPtr(req.Location),
		Stage:        stringPtr(req.Stage),
		HomeTeam:     stringPtr(req.HomeTeam),
		AwayTeam:     stringPtr(req.AwayTeam),
		MatchNumber:  intPtr(req.MatchNumber),
		Status:       "scheduled",
		CreatedBy:    userID.(uuid.UUID),
	}

	if err := m.DB.Create(&match).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to create match"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Match created successfully",
		"data":    match,
	})
}

// ListMatches lists matches in a tournament
func (m *Module) ListMatches(c *gin.Context) {
	tournamentID := c.Param("tournamentId")
	if tournamentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Tournament ID required"})
		return
	}

	tid, err := uuid.Parse(tournamentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid tournament ID"})
		return
	}

	var matches []domain.Match
	query := m.DB.Where("tournament_id = ?", tid).
		Preload("Video").
		Order("match_date ASC, match_number ASC")

	if err := query.Find(&matches).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to fetch matches"})
		return
	}

	// Get player counts for each match
	type matchPlayerCount struct {
		MatchID     uuid.UUID
		PlayerCount int
	}
	var playerCounts []matchPlayerCount
	m.DB.Model(&domain.MatchPlayer{}).
		Select("match_id, COUNT(*) as player_count").
		Where("match_id IN ?", getMatchIDs(matches)).
		Group("match_id").
		Scan(&playerCounts)

	countMap := make(map[uuid.UUID]int)
	for _, pc := range playerCounts {
		countMap[pc.MatchID] = pc.PlayerCount
	}

	// Get highlight counts for each match
	type matchHighlightCount struct {
		MatchID        uuid.UUID
		HighlightCount int
	}
	var highlightCounts []matchHighlightCount
	m.DB.Model(&domain.PlayerHighlight{}).
		Select("match_id, COUNT(*) as highlight_count").
		Where("match_id IN ?", getMatchIDs(matches)).
		Group("match_id").
		Scan(&highlightCounts)

	highlightMap := make(map[uuid.UUID]int)
	for _, hc := range highlightCounts {
		highlightMap[hc.MatchID] = hc.HighlightCount
	}

	// Build response
	type matchResponse struct {
		domain.Match
		PlayerCount    int  `json:"player_count"`
		HighlightCount int  `json:"highlight_count"`
		HasVideo       bool `json:"has_video"`
	}

	response := make([]matchResponse, len(matches))
	for i, match := range matches {
		response[i] = matchResponse{
			Match:          match,
			PlayerCount:    countMap[match.ID],
			HighlightCount: highlightMap[match.ID],
			HasVideo:       match.Video != nil,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})
}

// GetMatch retrieves a single match with full details
func (m *Module) GetMatch(c *gin.Context) {
	matchID := c.Param("id")
	mid, err := uuid.Parse(matchID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid match ID"})
		return
	}

	var match domain.Match
	if err := m.DB.Preload("Tournament").Preload("Video").Preload("Players").First(&match, "id = ?", mid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Match not found"})
		return
	}

	// Get match players with details
	var matchPlayers []struct {
		domain.MatchPlayer
		Player domain.Player `gorm:"embedded;embeddedPrefix:player_"`
	}
	m.DB.Model(&domain.MatchPlayer{}).
		Select("match_players.*, players.id as player_id, players.first_name as player_first_name, players.last_name as player_last_name, players.position as player_position, players.profile_photo_url as player_profile_photo_url").
		Joins("JOIN players ON players.id = match_players.player_id").
		Where("match_players.match_id = ?", mid).
		Scan(&matchPlayers)

	// Get highlight counts per player
	type playerHighlightCount struct {
		PlayerID       uuid.UUID
		HighlightCount int
	}
	var highlightCounts []playerHighlightCount
	m.DB.Model(&domain.PlayerHighlight{}).
		Select("player_id, COUNT(*) as highlight_count").
		Where("match_id = ?", mid).
		Group("player_id").
		Scan(&highlightCounts)

	highlightMap := make(map[uuid.UUID]int)
	for _, hc := range highlightCounts {
		highlightMap[hc.PlayerID] = hc.HighlightCount
	}

	// Build players response
	type playerInMatch struct {
		ID              uuid.UUID  `json:"id"`
		PlayerID        uuid.UUID  `json:"player_id"`
		FirstName       string     `json:"first_name"`
		LastName        string     `json:"last_name"`
		Position        string     `json:"position"`
		ProfilePhotoURL *string    `json:"profile_photo_url"`
		PositionPlayed  *string    `json:"position_played"`
		MinutesPlayed   *int       `json:"minutes_played"`
		Goals           int        `json:"goals"`
		Assists         int        `json:"assists"`
		HighlightCount  int        `json:"highlight_count"`
		IsStarter       bool       `json:"is_starter"`
		JerseyNumber    *int       `json:"jersey_number"`
		FormationX      *float64   `json:"formation_x"`
		FormationY      *float64   `json:"formation_y"`
		SubbedInFor     *uuid.UUID `json:"subbed_in_for"`
		SubbedInAt      *int       `json:"subbed_in_at"`
		SubbedOutAt     *int       `json:"subbed_out_at"`
	}

	players := make([]playerInMatch, len(matchPlayers))
	for i, mp := range matchPlayers {
		// Convert profile photo URL to presigned URL
		var profilePhotoURL *string
		if mp.Player.ProfilePhotoURL != nil && *mp.Player.ProfilePhotoURL != "" {
			url := m.getPresignedURL(*mp.Player.ProfilePhotoURL)
			profilePhotoURL = &url
		}

		players[i] = playerInMatch{
			ID:              mp.ID,
			PlayerID:        mp.PlayerID,
			FirstName:       mp.Player.FirstName,
			LastName:        mp.Player.LastName,
			Position:        mp.Player.Position,
			ProfilePhotoURL: profilePhotoURL,
			PositionPlayed:  mp.PositionPlayed,
			MinutesPlayed:   mp.MinutesPlayed,
			Goals:           mp.Goals,
			Assists:         mp.Assists,
			HighlightCount:  highlightMap[mp.PlayerID],
			IsStarter:       mp.IsStarter,
			JerseyNumber:    mp.JerseyNumber,
			FormationX:      mp.FormationX,
			FormationY:      mp.FormationY,
			SubbedInFor:     mp.SubbedInFor,
			SubbedInAt:      mp.SubbedInAt,
			SubbedOutAt:     mp.SubbedOutAt,
		}
	}

	// Build match response with video URL
	matchResponse := gin.H{
		"id":            match.ID,
		"title":         match.Title,
		"description":   match.Description,
		"match_date":    match.MatchDate,
		"location":      match.Location,
		"stage":         match.Stage,
		"home_team":     match.HomeTeam,
		"away_team":     match.AwayTeam,
		"home_score":    match.HomeScore,
		"away_score":    match.AwayScore,
		"status":        match.Status,
		"tournament_id": match.TournamentID,
		"created_at":    match.CreatedAt,
		"updated_at":    match.UpdatedAt,
	}

	// Add video with playable URL if exists
	if match.Video != nil {
		var videoURL string
		// Generate presigned URL for S3 (valid for 1 hour)
		if m.CDNHost != "" {
			// Use CloudFront if configured
			videoURL = fmt.Sprintf("%s/%s", m.CDNHost, match.Video.VideoURL)
		} else {
			// Generate presigned S3 URL
			presigner := s3.NewPresignClient(m.S3Client)
			presignedReq, err := presigner.PresignGetObject(c.Request.Context(), &s3.GetObjectInput{
				Bucket: aws.String(m.S3Bucket),
				Key:    aws.String(match.Video.VideoURL),
			}, s3.WithPresignExpires(1*time.Hour))
			if err != nil {
				fmt.Printf("Warning: Failed to generate presigned URL: %v\n", err)
				videoURL = "" // Will show error in frontend
			} else {
				videoURL = presignedReq.URL
			}
		}

		matchResponse["video"] = gin.H{
			"id":               match.Video.ID,
			"match_id":         match.Video.MatchID,
			"video_url":        videoURL,
			"thumbnail_url":    match.Video.ThumbnailURL,
			"duration_seconds": match.Video.DurationSeconds,
			"file_size_bytes":  match.Video.FileSizeBytes,
			"status":           match.Video.Status,
			"price_cents":      match.Video.PriceCents,
			"currency":         match.Video.Currency,
			"view_count":       match.Video.ViewCount,
			"purchase_count":   match.Video.PurchaseCount,
			"created_at":       match.Video.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"match":   matchResponse,
			"players": players,
		},
	})
}

// UpdateMatch updates match details
func (m *Module) UpdateMatch(c *gin.Context) {
	matchID := c.Param("id")
	mid, err := uuid.Parse(matchID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid match ID"})
		return
	}

	var match domain.Match
	if err := m.DB.First(&match, "id = ?", mid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Match not found"})
		return
	}

	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		MatchDate   string `json:"match_date"`
		Location    string `json:"location"`
		Stage       string `json:"stage"`
		HomeTeam    string `json:"home_team"`
		AwayTeam    string `json:"away_team"`
		HomeScore   *int   `json:"home_score"`
		AwayScore   *int   `json:"away_score"`
		Status      string `json:"status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request"})
		return
	}

	updates := make(map[string]interface{})
	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.MatchDate != "" {
		if matchDate, err := time.Parse("2006-01-02", req.MatchDate); err == nil {
			updates["match_date"] = matchDate
		}
	}
	if req.Location != "" {
		updates["location"] = req.Location
	}
	if req.Stage != "" {
		updates["stage"] = req.Stage
	}
	if req.HomeTeam != "" {
		updates["home_team"] = req.HomeTeam
	}
	if req.AwayTeam != "" {
		updates["away_team"] = req.AwayTeam
	}
	if req.HomeScore != nil {
		updates["home_score"] = *req.HomeScore
	}
	if req.AwayScore != nil {
		updates["away_score"] = *req.AwayScore
	}
	if req.Status != "" {
		updates["status"] = req.Status
	}

	if err := m.DB.Model(&match).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to update match"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Match updated successfully",
		"data":    match,
	})
}

// DeleteMatch deletes a match (cascades to video and highlights)
func (m *Module) DeleteMatch(c *gin.Context) {
	matchID := c.Param("id")
	mid, err := uuid.Parse(matchID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid match ID"})
		return
	}

	if err := m.DB.Delete(&domain.Match{}, "id = ?", mid).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to delete match"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Match deleted successfully"})
}

// ==================== MATCH PLAYERS ====================

// AddPlayerToMatch adds a player to a match
func (m *Module) AddPlayerToMatch(c *gin.Context) {
	matchID := c.Param("id")
	mid, err := uuid.Parse(matchID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid match ID"})
		return
	}

	var req struct {
		PlayerID       string   `json:"player_id" binding:"required"`
		PositionPlayed string   `json:"position_played"`
		MinutesPlayed  *int     `json:"minutes_played"`
		Goals          int      `json:"goals"`
		Assists        int      `json:"assists"`
		Notes          string   `json:"notes"`
		IsStarter      *bool    `json:"is_starter"`
		JerseyNumber   *int     `json:"jersey_number"`
		FormationX     *float64 `json:"formation_x"`
		FormationY     *float64 `json:"formation_y"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request"})
		return
	}

	playerID, err := uuid.Parse(req.PlayerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid player ID"})
		return
	}

	// Verify match exists
	var match domain.Match
	if err := m.DB.First(&match, "id = ?", mid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Match not found"})
		return
	}

	// Verify player exists
	var player domain.Player
	if err := m.DB.First(&player, "id = ?", playerID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Player not found"})
		return
	}

	// Default isStarter to true if not provided
	isStarter := true
	if req.IsStarter != nil {
		isStarter = *req.IsStarter
	}

	// Default position_played to player's registered position if not provided
	positionPlayed := req.PositionPlayed
	if positionPlayed == "" {
		positionPlayed = player.Position
	}

	matchPlayer := domain.MatchPlayer{
		MatchID:        mid,
		PlayerID:       playerID,
		PositionPlayed: stringPtr(positionPlayed),
		MinutesPlayed:  req.MinutesPlayed,
		Goals:          req.Goals,
		Assists:        req.Assists,
		Notes:          stringPtr(req.Notes),
		IsStarter:      isStarter,
		JerseyNumber:   req.JerseyNumber,
		FormationX:     req.FormationX,
		FormationY:     req.FormationY,
	}

	if err := m.DB.Create(&matchPlayer).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"success": false, "message": "Player already in match"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Player added to match",
		"data":    matchPlayer,
	})
}

// UpdateMatchPlayer updates player stats in a match
func (m *Module) UpdateMatchPlayer(c *gin.Context) {
	matchID := c.Param("id")
	playerID := c.Param("playerId")

	mid, err := uuid.Parse(matchID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid match ID"})
		return
	}

	pid, err := uuid.Parse(playerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid player ID"})
		return
	}

	var matchPlayer domain.MatchPlayer
	if err := m.DB.First(&matchPlayer, "match_id = ? AND player_id = ?", mid, pid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Player not in match"})
		return
	}

	var req struct {
		PositionPlayed string   `json:"position_played"`
		MinutesPlayed  *int     `json:"minutes_played"`
		Goals          *int     `json:"goals"`
		Assists        *int     `json:"assists"`
		Notes          string   `json:"notes"`
		IsStarter      *bool    `json:"is_starter"`
		JerseyNumber   *int     `json:"jersey_number"`
		FormationX     *float64 `json:"formation_x"`
		FormationY     *float64 `json:"formation_y"`
		SubbedInAt     *int     `json:"subbed_in_at"`
		SubbedOutAt    *int     `json:"subbed_out_at"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request"})
		return
	}

	updates := make(map[string]interface{})
	if req.PositionPlayed != "" {
		updates["position_played"] = req.PositionPlayed
	}
	if req.MinutesPlayed != nil {
		updates["minutes_played"] = *req.MinutesPlayed
	}
	if req.Goals != nil {
		updates["goals"] = *req.Goals
	}
	if req.Assists != nil {
		updates["assists"] = *req.Assists
	}
	if req.Notes != "" {
		updates["notes"] = req.Notes
	}
	if req.IsStarter != nil {
		updates["is_starter"] = *req.IsStarter
	}
	if req.JerseyNumber != nil {
		updates["jersey_number"] = *req.JerseyNumber
	}
	if req.FormationX != nil {
		updates["formation_x"] = *req.FormationX
	}
	if req.FormationY != nil {
		updates["formation_y"] = *req.FormationY
	}
	if req.SubbedInAt != nil {
		updates["subbed_in_at"] = *req.SubbedInAt
	}
	if req.SubbedOutAt != nil {
		updates["subbed_out_at"] = *req.SubbedOutAt
	}

	if err := m.DB.Model(&matchPlayer).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to update"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Player updated"})
}

// RemovePlayerFromMatch removes a player from a match
func (m *Module) RemovePlayerFromMatch(c *gin.Context) {
	matchID := c.Param("id")
	playerID := c.Param("playerId")

	mid, err := uuid.Parse(matchID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid match ID"})
		return
	}

	pid, err := uuid.Parse(playerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid player ID"})
		return
	}

	if err := m.DB.Delete(&domain.MatchPlayer{}, "match_id = ? AND player_id = ?", mid, pid).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to remove player"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Player removed from match"})
}

// ==================== MATCH VIDEO (FULL MATCH - PAID) ====================

// InitMatchVideoUpload initializes upload for full match video
func (m *Module) InitMatchVideoUpload(c *gin.Context) {
	matchID := c.Param("id")
	mid, err := uuid.Parse(matchID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid match ID"})
		return
	}

	// Check match exists
	var match domain.Match
	if err := m.DB.First(&match, "id = ?", mid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Match not found"})
		return
	}

	// Check if video already exists
	var existingVideo domain.MatchVideo
	if err := m.DB.First(&existingVideo, "match_id = ?", mid).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"success": false, "message": "Video already exists for this match. Delete it first to replace."})
		return
	}

	var req struct {
		FileName    string `json:"file_name" binding:"required"`
		FileSize    int64  `json:"file_size" binding:"required"`
		ContentType string `json:"content_type" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request"})
		return
	}

	// Validate file size (max 25GB)
	maxSize := int64(25 * 1024 * 1024 * 1024)
	if req.FileSize > maxSize {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "File too large. Maximum 25GB."})
		return
	}

	userID, _ := c.Get("user_id")

	// Generate S3 key
	s3Key := fmt.Sprintf("matches/%s/video/%s", mid.String(), req.FileName)

	// Determine upload method
	multipartThreshold := int64(100 * 1024 * 1024) // 100MB
	useMultipart := req.FileSize > multipartThreshold

	if useMultipart {
		// Create multipart upload
		createResp, err := m.S3Client.CreateMultipartUpload(c.Request.Context(), &s3.CreateMultipartUploadInput{
			Bucket:      aws.String(m.S3Bucket),
			Key:         aws.String(s3Key),
			ContentType: aws.String(req.ContentType),
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to initiate upload"})
			return
		}

		// Create upload session
		session := domain.UploadSession{
			UploadType:  "match_video",
			ContentType: req.ContentType,
			FileName:    req.FileName,
			FileSize:    req.FileSize,
			S3UploadID:  createResp.UploadId,
			S3Key:       s3Key,
			Status:      "pending",
			EntityType:  stringPtr("match"),
			EntityID:    &mid,
			UploadedBy:  userID.(uuid.UUID),
			ExpiresAt:   time.Now().Add(24 * time.Hour),
		}

		if err := m.DB.Create(&session).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to create session"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"session_id":    session.ID,
				"upload_method": "multipart",
				"s3_upload_id":  *createResp.UploadId,
				"s3_key":        s3Key,
				"part_size":     10 * 1024 * 1024, // 10MB parts
			},
		})
	} else {
		// Single presigned URL
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
		session := domain.UploadSession{
			UploadType:  "match_video",
			ContentType: req.ContentType,
			FileName:    req.FileName,
			FileSize:    req.FileSize,
			S3Key:       s3Key,
			Status:      "pending",
			EntityType:  stringPtr("match"),
			EntityID:    &mid,
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
				"session_id":    session.ID,
				"upload_method": "direct",
				"upload_url":    presignResult.URL,
				"s3_key":        s3Key,
				"expires_in":    3600,
			},
		})
	}
}

// GetMultipartPartURL gets presigned URL for a multipart upload part
func (m *Module) GetMultipartPartURL(c *gin.Context) {
	var req struct {
		SessionID  string `json:"session_id" binding:"required"`
		PartNumber int    `json:"part_number" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request"})
		return
	}

	sessionID, err := uuid.Parse(req.SessionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid session ID"})
		return
	}

	var session domain.UploadSession
	if err := m.DB.First(&session, "id = ?", sessionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Session not found"})
		return
	}

	if session.S3UploadID == nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Not a multipart upload"})
		return
	}

	presigner := s3.NewPresignClient(m.S3Client)
	presignResult, err := presigner.PresignUploadPart(c.Request.Context(), &s3.UploadPartInput{
		Bucket:     aws.String(m.S3Bucket),
		Key:        aws.String(session.S3Key),
		UploadId:   session.S3UploadID,
		PartNumber: aws.Int32(int32(req.PartNumber)),
	}, s3.WithPresignExpires(1*time.Hour))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to generate part URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"upload_url":  presignResult.URL,
			"part_number": req.PartNumber,
			"expires_in":  3600,
		},
	})
}

// CompleteMultipartUpload completes a multipart upload
func (m *Module) CompleteMultipartUpload(c *gin.Context) {
	var req struct {
		SessionID string `json:"session_id" binding:"required"`
		Parts     []struct {
			PartNumber int    `json:"part_number"`
			ETag       string `json:"etag"`
		} `json:"parts" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request"})
		return
	}

	sessionID, err := uuid.Parse(req.SessionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid session ID"})
		return
	}

	var session domain.UploadSession
	if err := m.DB.First(&session, "id = ?", sessionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Session not found"})
		return
	}

	if session.S3UploadID == nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Not a multipart upload"})
		return
	}

	// Build completed parts
	var completedParts []types.CompletedPart
	for _, p := range req.Parts {
		completedParts = append(completedParts, types.CompletedPart{
			PartNumber: aws.Int32(int32(p.PartNumber)),
			ETag:       aws.String(p.ETag),
		})
	}

	// Complete multipart upload
	_, err = m.S3Client.CompleteMultipartUpload(c.Request.Context(), &s3.CompleteMultipartUploadInput{
		Bucket:   aws.String(m.S3Bucket),
		Key:      aws.String(session.S3Key),
		UploadId: session.S3UploadID,
		MultipartUpload: &types.CompletedMultipartUpload{
			Parts: completedParts,
		},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to complete upload"})
		return
	}

	// Update session
	now := time.Now()
	session.Status = "completed"
	session.CompletedAt = &now
	m.DB.Save(&session)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Upload completed",
		"data": gin.H{
			"s3_key": session.S3Key,
		},
	})
}

// SaveMatchVideo creates the match video record after upload
func (m *Module) SaveMatchVideo(c *gin.Context) {
	matchID := c.Param("id")
	mid, err := uuid.Parse(matchID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid match ID"})
		return
	}

	var req struct {
		S3Key           string `json:"s3_key" binding:"required"`
		ThumbnailURL    string `json:"thumbnail_url"`
		DurationSeconds int    `json:"duration_seconds"`
		FileSizeBytes   int64  `json:"file_size_bytes"`
		PriceCents      int    `json:"price_cents"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request"})
		return
	}

	// Check match exists
	var match domain.Match
	if err := m.DB.First(&match, "id = ?", mid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Match not found"})
		return
	}

	// Check if video already exists
	var existingVideo domain.MatchVideo
	if err := m.DB.First(&existingVideo, "match_id = ?", mid).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"success": false, "message": "Video already exists"})
		return
	}

	userID, _ := c.Get("user_id")

	priceCents := req.PriceCents
	if priceCents == 0 {
		priceCents = 999 // Default $9.99
	}

	video := domain.MatchVideo{
		MatchID:         mid,
		VideoURL:        req.S3Key,
		ThumbnailURL:    stringPtr(req.ThumbnailURL),
		DurationSeconds: intPtr(req.DurationSeconds),
		FileSizeBytes:   int64Ptr(req.FileSizeBytes),
		Status:          "ready",
		PriceCents:      priceCents,
		Currency:        "USD",
		UploadedBy:      userID.(uuid.UUID),
	}

	if err := m.DB.Create(&video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to save video"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Match video saved",
		"data":    video,
	})
}

// DeleteMatchVideo deletes the match video
func (m *Module) DeleteMatchVideo(c *gin.Context) {
	matchID := c.Param("id")
	mid, err := uuid.Parse(matchID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid match ID"})
		return
	}

	var video domain.MatchVideo
	if err := m.DB.First(&video, "match_id = ?", mid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Video not found"})
		return
	}

	// Delete from S3
	_, err = m.S3Client.DeleteObject(c.Request.Context(), &s3.DeleteObjectInput{
		Bucket: aws.String(m.S3Bucket),
		Key:    aws.String(video.VideoURL),
	})
	if err != nil {
		// Log but continue - file might already be deleted
		fmt.Printf("Warning: Failed to delete S3 object: %v\n", err)
	}

	if err := m.DB.Delete(&video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to delete video"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Video deleted"})
}

// InitThumbnailUpload initializes upload for video thumbnail
func (m *Module) InitThumbnailUpload(c *gin.Context) {
	matchID := c.Param("id")
	mid, err := uuid.Parse(matchID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid match ID"})
		return
	}

	var req struct {
		FileName    string `json:"file_name" binding:"required"`
		FileSize    int64  `json:"file_size"`
		ContentType string `json:"content_type"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request"})
		return
	}

	// Check video exists
	var video domain.MatchVideo
	if err := m.DB.First(&video, "match_id = ?", mid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Video not found"})
		return
	}

	// Generate S3 key for thumbnail
	s3Key := fmt.Sprintf("thumbnails/matches/%s/%s", matchID, req.FileName)

	// Generate presigned PUT URL
	presigner := s3.NewPresignClient(m.S3Client)
	presignedReq, err := presigner.PresignPutObject(c.Request.Context(), &s3.PutObjectInput{
		Bucket:      aws.String(m.S3Bucket),
		Key:         aws.String(s3Key),
		ContentType: aws.String(req.ContentType),
	}, s3.WithPresignExpires(15*time.Minute))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to generate upload URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"upload_url": presignedReq.URL,
			"s3_key":     s3Key,
		},
	})
}

// UpdateThumbnail updates the video thumbnail URL
func (m *Module) UpdateThumbnail(c *gin.Context) {
	matchID := c.Param("id")
	mid, err := uuid.Parse(matchID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid match ID"})
		return
	}

	var req struct {
		S3Key string `json:"s3_key" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request"})
		return
	}

	// Check video exists
	var video domain.MatchVideo
	if err := m.DB.First(&video, "match_id = ?", mid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Video not found"})
		return
	}

	// Generate thumbnail URL
	var thumbnailURL string
	if m.CDNHost != "" {
		thumbnailURL = fmt.Sprintf("%s/%s", m.CDNHost, req.S3Key)
	} else {
		thumbnailURL = fmt.Sprintf("https://%s.s3.amazonaws.com/%s", m.S3Bucket, req.S3Key)
	}

	// Update video record
	video.ThumbnailURL = &thumbnailURL
	if err := m.DB.Save(&video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to update thumbnail"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Thumbnail updated",
		"data": gin.H{
			"thumbnail_url": thumbnailURL,
		},
	})
}

// ==================== HELPERS ====================

func getMatchIDs(matches []domain.Match) []uuid.UUID {
	ids := make([]uuid.UUID, len(matches))
	for i, m := range matches {
		ids[i] = m.ID
	}
	return ids
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

// ParseInt parses string to int with default
func ParseInt(s string, defaultVal int) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	return defaultVal
}
