package profiles

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/unicorn-sport/backend/internal/domain"
)

// ProfilesModule handles player profile viewing
type ProfilesModule struct {
	db       *gorm.DB
	s3Client *s3.Client
	s3Bucket string
}

// NewProfilesModule creates a new profiles module
func NewProfilesModule(db *gorm.DB, s3Client *s3.Client, s3Bucket string) *ProfilesModule {
	return &ProfilesModule{
		db:       db,
		s3Client: s3Client,
		s3Bucket: s3Bucket,
	}
}

// getPresignedURL converts an s3:// URL to a presigned URL
func (m *ProfilesModule) getPresignedURL(s3URL string) string {
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

	presignClient := s3.NewPresignClient(m.s3Client)
	presignedReq, err := presignClient.PresignGetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(m.s3Bucket),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(time.Hour))
	if err != nil {
		return s3URL
	}

	return presignedReq.URL
}

// PlayerListResponse is the public player list response
type PlayerListResponse struct {
	ID                uuid.UUID `json:"id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	LastNameInit      string    `json:"last_name_init"`
	Age               int       `json:"age"`
	Position          string    `json:"position"`
	Country           string    `json:"country"`
	HeightCm          *int      `json:"height_cm,omitempty"`
	PreferredFoot     *string   `json:"preferred_foot,omitempty"`
	ThumbnailURL      *string   `json:"thumbnail_url,omitempty"`
	ProfilePhotoURL   *string   `json:"profile_photo_url,omitempty"`
	VideoThumbnailURL *string   `json:"video_thumbnail_url,omitempty"`
	AcademyName       *string   `json:"academy_name,omitempty"`
	VideoCount        int       `json:"video_count"`
	IsVerified        bool      `json:"is_verified"`
}

// FeaturedPlayerResponse is the response for featured players on homepage
type FeaturedPlayerResponse struct {
	ID              uuid.UUID `json:"id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	DateOfBirth     string    `json:"date_of_birth,omitempty"`
	Age             int       `json:"age"`
	Position        string    `json:"position"`
	Country         string    `json:"country"`
	ProfilePhotoURL *string   `json:"profile_photo_url,omitempty"`
	AcademyName     *string   `json:"academy_name,omitempty"`
	IsVerified      bool      `json:"is_verified"`
}

// PlayerDetailResponse is the detailed player response
type PlayerDetailResponse struct {
	ID              uuid.UUID            `json:"id"`
	FirstName       string               `json:"first_name"`
	LastName        string               `json:"last_name"`
	LastNameInit    string               `json:"last_name_init"`
	Age             int                  `json:"age"`
	Position        string               `json:"position"`
	PreferredFoot   *string              `json:"preferred_foot,omitempty"`
	HeightCm        *int                 `json:"height_cm,omitempty"`
	WeightKg        *int                 `json:"weight_kg,omitempty"`
	Country         string               `json:"country"`
	City            *string              `json:"city,omitempty"`
	State           *string              `json:"state,omitempty"`
	SchoolName      *string              `json:"school_name,omitempty"`
	TournamentName  *string              `json:"tournament_name,omitempty"`
	AcademyID       *uuid.UUID           `json:"academy_id,omitempty"`
	AcademyName     *string              `json:"academy_name,omitempty"`
	ProfilePhotoURL *string              `json:"profile_photo_url,omitempty"`
	IsVerified      bool                 `json:"is_verified"`
	Stats           *PlayerStatsResponse `json:"stats,omitempty"`
	HighlightVideos []VideoResponse      `json:"highlight_videos"`
	FullMatchVideos []VideoResponse      `json:"full_match_videos"`
}

// PlayerStatsResponse contains aggregated player statistics
// All stats are derived from highlight uploads (1 highlight = 1 stat)
// This means uploading a "goal" highlight automatically counts as 1 goal
type PlayerStatsResponse struct {
	Season         string `json:"season"`
	MatchesPlayed  int    `json:"matches_played"`
	MatchesStarted int    `json:"matches_started"`
	MinutesPlayed  int    `json:"minutes_played"`
	// Offensive stats
	Goals     int `json:"goals"`
	Assists   int `json:"assists"`
	Shooting  int `json:"shooting"`
	Dribbling int `json:"dribbling"`
	// Playmaking stats
	Passing int `json:"passing"`
	Heading int `json:"heading"`
	Speed   int `json:"speed"`
	// Defensive stats
	Defending int `json:"defending"`
	Tackling  int `json:"tackling"`
	Saves     int `json:"saves"` // For goalkeepers
	// Total highlights count
	TotalHighlights int `json:"total_highlights"`
}

// SimilarPlayerResponse for recommendations
type SimilarPlayerResponse struct {
	ID              uuid.UUID `json:"id"`
	FirstName       string    `json:"first_name"`
	LastNameInit    string    `json:"last_name_init"`
	Age             int       `json:"age"`
	Position        string    `json:"position"`
	Country         string    `json:"country"`
	ProfilePhotoURL *string   `json:"profile_photo_url,omitempty"`
	AcademyName     *string   `json:"academy_name,omitempty"`
	IsVerified      bool      `json:"is_verified"`
	MatchReason     string    `json:"match_reason"` // Why this player was recommended
}

// VideoResponse is the video response in player detail
type VideoResponse struct {
	ID              uuid.UUID `json:"id"`
	Title           string    `json:"title"`
	ThumbnailURL    *string   `json:"thumbnail_url,omitempty"`
	DurationSeconds *int      `json:"duration_seconds,omitempty"`
}

// ListPlayers returns public player list with pagination
// @Summary List players
// @Description Get a paginated list of verified players with optional filters
// @Tags Players
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(20) maximum(100)
// @Param position query string false "Filter by position"
// @Param country query string false "Filter by country"
// @Param age_min query int false "Minimum age"
// @Param age_max query int false "Maximum age"
// @Success 200 {object} map[string]interface{} "List of players with pagination"
// @Router /players [get]
func (m *ProfilesModule) ListPlayers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if limit > 100 {
		limit = 100
	}
	offset := (page - 1) * limit

	var players []domain.Player
	var total int64

	query := m.db.Model(&domain.Player{}).
		Preload("Academy").
		Preload("Highlights").
		Where("deleted_at IS NULL").
		Where("verification_status = ?", "verified") // Only show verified players publicly

	// Apply filters
	if position := c.Query("position"); position != "" {
		query = query.Where("position ILIKE ?", "%"+position+"%")
	}
	if country := c.Query("country"); country != "" {
		query = query.Where("country = ?", country)
	}
	if ageMin := c.Query("age_min"); ageMin != "" {
		if minAge, err := strconv.Atoi(ageMin); err == nil {
			maxBirthDate := time.Now().AddDate(-minAge, 0, 0)
			query = query.Where("date_of_birth <= ?", maxBirthDate)
		}
	}
	if ageMax := c.Query("age_max"); ageMax != "" {
		if maxAge, err := strconv.Atoi(ageMax); err == nil {
			minBirthDate := time.Now().AddDate(-maxAge-1, 0, 0)
			query = query.Where("date_of_birth > ?", minBirthDate)
		}
	}
	if tournamentID := c.Query("tournament_id"); tournamentID != "" {
		query = query.Where("tournament_id = ?", tournamentID)
	}
	if academyID := c.Query("academy_id"); academyID != "" {
		query = query.Where("academy_id = ?", academyID)
	}

	query.Count(&total)
	query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&players)

	// Convert to public response
	response := make([]PlayerListResponse, len(players))
	for i, p := range players {
		// Convert thumbnail URL to presigned URL if it's an S3 URL
		var thumbnailURL *string
		if p.ThumbnailURL != nil && *p.ThumbnailURL != "" {
			url := m.getPresignedURL(*p.ThumbnailURL)
			thumbnailURL = &url
		}

		// Convert profile photo URL to presigned URL if it's an S3 URL
		var profilePhotoURL *string
		if p.ProfilePhotoURL != nil && *p.ProfilePhotoURL != "" {
			url := m.getPresignedURL(*p.ProfilePhotoURL)
			profilePhotoURL = &url
		}

		// Get video thumbnail from first highlight if available
		var videoThumbnailURL *string
		if len(p.Highlights) > 0 && p.Highlights[0].ThumbnailURL != nil && *p.Highlights[0].ThumbnailURL != "" {
			url := m.getPresignedURL(*p.Highlights[0].ThumbnailURL)
			videoThumbnailURL = &url
		}

		// Get academy name
		var academyName *string
		if p.Academy != nil {
			academyName = &p.Academy.Name
		}

		response[i] = PlayerListResponse{
			ID:                p.ID,
			FirstName:         p.FirstName,
			LastName:          p.LastName,
			LastNameInit:      p.GetLastNameInit(),
			Age:               p.GetAge(),
			Position:          p.Position,
			Country:           p.Country,
			HeightCm:          p.HeightCm,
			PreferredFoot:     p.PreferredFoot,
			ThumbnailURL:      thumbnailURL,
			ProfilePhotoURL:   profilePhotoURL,
			VideoThumbnailURL: videoThumbnailURL,
			AcademyName:       academyName,
			VideoCount:        len(p.Highlights),
			IsVerified:        p.IsVerified(),
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"players": response,
			"pagination": gin.H{
				"page":        page,
				"limit":       limit,
				"total":       total,
				"total_pages": (total + int64(limit) - 1) / int64(limit),
			},
		},
	})
}

// GetPlayer returns detailed player info
// @Summary Get player by ID
// @Description Get detailed information about a specific player including videos
// @Tags Players
// @Accept json
// @Produce json
// @Param id path string true "Player ID (UUID)"
// @Success 200 {object} map[string]interface{} "Player details"
// @Failure 400 {object} map[string]interface{} "Invalid player ID"
// @Failure 404 {object} map[string]interface{} "Player not found"
// @Router /players/{id} [get]
func (m *ProfilesModule) GetPlayer(c *gin.Context) {
	playerID := c.Param("id")
	pid, err := uuid.Parse(playerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_UUID", "message": "Invalid player ID"}})
		return
	}

	var player domain.Player
	if err := m.db.Preload("Tournament").Preload("Academy").First(&player, "id = ? AND deleted_at IS NULL", pid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "Player not found"}})
		return
	}

	// Get subscription tier to determine video access
	subscriptionTier := "free"
	if tier, exists := c.Get("subscription_tier"); exists {
		subscriptionTier = tier.(string)
	}

	// Get highlight videos (always visible)
	var highlightVideos []domain.Video
	m.db.Joins("JOIN player_videos ON player_videos.video_id = videos.id").
		Where("player_videos.player_id = ? AND videos.video_type = ?", pid, "highlight").
		Find(&highlightVideos)

	// Get full match videos (only for paid subscribers)
	var fullMatchVideos []domain.Video
	if subscriptionTier != "free" {
		m.db.Joins("JOIN player_videos ON player_videos.video_id = videos.id").
			Where("player_videos.player_id = ? AND videos.video_type = ?", pid, "full_match").
			Find(&fullMatchVideos)
	}

	// Get player stats - aggregate from match_players table
	stats := m.getPlayerStats(pid)

	// Convert profile photo URL to presigned URL
	var profilePhotoURL *string
	if player.ProfilePhotoURL != nil && *player.ProfilePhotoURL != "" {
		url := m.getPresignedURL(*player.ProfilePhotoURL)
		profilePhotoURL = &url
	}

	// Build response
	response := PlayerDetailResponse{
		ID:              player.ID,
		FirstName:       player.FirstName,
		LastName:        player.LastName,
		LastNameInit:    player.GetLastNameInit(),
		Age:             player.GetAge(),
		Position:        player.Position,
		PreferredFoot:   player.PreferredFoot,
		HeightCm:        player.HeightCm,
		WeightKg:        player.WeightKg,
		Country:         player.Country,
		City:            player.City,
		State:           player.State,
		SchoolName:      player.SchoolName,
		ProfilePhotoURL: profilePhotoURL,
		IsVerified:      player.IsVerified(),
		Stats:           stats,
		HighlightVideos: make([]VideoResponse, len(highlightVideos)),
		FullMatchVideos: make([]VideoResponse, len(fullMatchVideos)),
	}

	if player.Tournament != nil {
		response.TournamentName = &player.Tournament.Name
	}

	if player.Academy != nil && player.Academy.Name != "" {
		response.AcademyID = &player.Academy.ID
		response.AcademyName = &player.Academy.Name
	}

	for i, v := range highlightVideos {
		response.HighlightVideos[i] = VideoResponse{
			ID:              v.ID,
			Title:           v.Title,
			ThumbnailURL:    v.ThumbnailURL,
			DurationSeconds: v.DurationSeconds,
		}
	}

	for i, v := range fullMatchVideos {
		response.FullMatchVideos[i] = VideoResponse{
			ID:              v.ID,
			Title:           v.Title,
			ThumbnailURL:    v.ThumbnailURL,
			DurationSeconds: v.DurationSeconds,
		}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": response})
}

// getPlayerStats aggregates stats from multiple sources:
// - Matches played/started, minutes from match_players table
// - Goals/assists derived from player_highlights (each "goal" highlight = 1 goal, etc.)
// This approach ensures stats are automatically calculated from uploaded highlight videos
func (m *ProfilesModule) getPlayerStats(playerID uuid.UUID) *PlayerStatsResponse {
	// Get match participation stats from match_players
	type matchStats struct {
		MatchesPlayed  int `json:"matches_played"`
		MatchesStarted int `json:"matches_started"`
		MinutesPlayed  int `json:"minutes_played"`
	}

	var matchResult matchStats
	m.db.Model(&domain.MatchPlayer{}).
		Select(`
			COUNT(*) as matches_played,
			SUM(CASE WHEN is_starter = true THEN 1 ELSE 0 END) as matches_started,
			COALESCE(SUM(minutes_played), 0) as minutes_played
		`).
		Where("player_id = ?", playerID).
		Scan(&matchResult)

	// Get ALL stats from highlights (each highlight = 1 stat)
	// This is derived from highlight uploads - no manual entry needed
	// Highlight types: goal, assist, dribbling, defending, tackling, passing, shooting, heading, speed, save
	type highlightStats struct {
		Goals           int `json:"goals"`
		Assists         int `json:"assists"`
		Shooting        int `json:"shooting"`
		Dribbling       int `json:"dribbling"`
		Passing         int `json:"passing"`
		Heading         int `json:"heading"`
		Speed           int `json:"speed"`
		Defending       int `json:"defending"`
		Tackling        int `json:"tackling"`
		Saves           int `json:"saves"`
		TotalHighlights int `json:"total_highlights"`
	}

	var highlightResult highlightStats
	m.db.Model(&domain.PlayerHighlight{}).
		Select(`
			SUM(CASE WHEN LOWER(highlight_type) = 'goal' THEN 1 ELSE 0 END) as goals,
			SUM(CASE WHEN LOWER(highlight_type) = 'assist' THEN 1 ELSE 0 END) as assists,
			SUM(CASE WHEN LOWER(highlight_type) = 'shooting' THEN 1 ELSE 0 END) as shooting,
			SUM(CASE WHEN LOWER(highlight_type) = 'dribbling' THEN 1 ELSE 0 END) as dribbling,
			SUM(CASE WHEN LOWER(highlight_type) = 'passing' THEN 1 ELSE 0 END) as passing,
			SUM(CASE WHEN LOWER(highlight_type) = 'heading' THEN 1 ELSE 0 END) as heading,
			SUM(CASE WHEN LOWER(highlight_type) = 'speed' THEN 1 ELSE 0 END) as speed,
			SUM(CASE WHEN LOWER(highlight_type) = 'defending' THEN 1 ELSE 0 END) as defending,
			SUM(CASE WHEN LOWER(highlight_type) = 'tackling' THEN 1 ELSE 0 END) as tackling,
			SUM(CASE WHEN LOWER(highlight_type) = 'save' THEN 1 ELSE 0 END) as saves,
			COUNT(*) as total_highlights
		`).
		Where("player_id = ? AND status = 'approved'", playerID).
		Scan(&highlightResult)

	// Return stats if player has either match data OR highlight data
	if matchResult.MatchesPlayed == 0 && highlightResult.TotalHighlights == 0 {
		return nil
	}

	// Determine current season (e.g., "2025-26")
	now := time.Now()
	year := now.Year()
	var season string
	if now.Month() >= 8 { // August onwards = new season
		season = strconv.Itoa(year) + "-" + strconv.Itoa((year+1)%100)
	} else {
		season = strconv.Itoa(year-1) + "-" + strconv.Itoa(year%100)
	}

	return &PlayerStatsResponse{
		Season:          season,
		MatchesPlayed:   matchResult.MatchesPlayed,
		MatchesStarted:  matchResult.MatchesStarted,
		MinutesPlayed:   matchResult.MinutesPlayed,
		Goals:           highlightResult.Goals,
		Assists:         highlightResult.Assists,
		Shooting:        highlightResult.Shooting,
		Dribbling:       highlightResult.Dribbling,
		Passing:         highlightResult.Passing,
		Heading:         highlightResult.Heading,
		Speed:           highlightResult.Speed,
		Defending:       highlightResult.Defending,
		Tackling:        highlightResult.Tackling,
		Saves:           highlightResult.Saves,
		TotalHighlights: highlightResult.TotalHighlights,
	}
}

// GetSimilarPlayers returns players similar to the specified player
func (m *ProfilesModule) GetSimilarPlayers(c *gin.Context) {
	playerID := c.Param("id")
	pid, err := uuid.Parse(playerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_UUID", "message": "Invalid player ID"}})
		return
	}

	// Get the reference player
	var player domain.Player
	if err := m.db.Preload("Academy").First(&player, "id = ? AND deleted_at IS NULL", pid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "Player not found"}})
		return
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	if limit > 12 {
		limit = 12
	}

	var similar []domain.Player
	var matchReasons = make(map[uuid.UUID]string)

	// Strategy: Find players with same position, similar age, same academy, or same country
	// Priority: Same Academy > Same Position + Country > Same Position + Age

	// 1. Same academy (highest priority)
	if player.AcademyID != nil {
		var academyPlayers []domain.Player
		m.db.Preload("Academy").
			Where("id != ? AND academy_id = ? AND deleted_at IS NULL AND verification_status = 'verified'", pid, player.AcademyID).
			Limit(3).
			Find(&academyPlayers)

		for _, p := range academyPlayers {
			matchReasons[p.ID] = "Same academy"
			similar = append(similar, p)
		}
	}

	// 2. Same position + country (if not enough from academy)
	if len(similar) < limit {
		remaining := limit - len(similar)
		var existingIDs []uuid.UUID
		existingIDs = append(existingIDs, pid)
		for _, p := range similar {
			existingIDs = append(existingIDs, p.ID)
		}

		var positionCountryPlayers []domain.Player
		m.db.Preload("Academy").
			Where("id NOT IN ? AND position = ? AND country = ? AND deleted_at IS NULL AND verification_status = 'verified'", existingIDs, player.Position, player.Country).
			Limit(remaining).
			Find(&positionCountryPlayers)

		for _, p := range positionCountryPlayers {
			matchReasons[p.ID] = "Same position & country"
			similar = append(similar, p)
		}
	}

	// 3. Same position + similar age (within 2 years)
	if len(similar) < limit {
		remaining := limit - len(similar)
		var existingIDs []uuid.UUID
		existingIDs = append(existingIDs, pid)
		for _, p := range similar {
			existingIDs = append(existingIDs, p.ID)
		}

		minBirth := player.DateOfBirth.AddDate(-2, 0, 0)
		maxBirth := player.DateOfBirth.AddDate(2, 0, 0)

		var ageMatchPlayers []domain.Player
		m.db.Preload("Academy").
			Where("id NOT IN ? AND position = ? AND date_of_birth BETWEEN ? AND ? AND deleted_at IS NULL AND verification_status = 'verified'", existingIDs, player.Position, minBirth, maxBirth).
			Limit(remaining).
			Find(&ageMatchPlayers)

		for _, p := range ageMatchPlayers {
			matchReasons[p.ID] = "Similar age & position"
			similar = append(similar, p)
		}
	}

	// Build response
	response := make([]SimilarPlayerResponse, len(similar))
	for i, p := range similar {
		var profilePhotoURL *string
		if p.ProfilePhotoURL != nil && *p.ProfilePhotoURL != "" {
			url := m.getPresignedURL(*p.ProfilePhotoURL)
			profilePhotoURL = &url
		}

		var academyName *string
		if p.Academy != nil {
			academyName = &p.Academy.Name
		}

		response[i] = SimilarPlayerResponse{
			ID:              p.ID,
			FirstName:       p.FirstName,
			LastNameInit:    p.GetLastNameInit(),
			Age:             p.GetAge(),
			Position:        p.Position,
			Country:         p.Country,
			ProfilePhotoURL: profilePhotoURL,
			AcademyName:     academyName,
			IsVerified:      p.IsVerified(),
			MatchReason:     matchReasons[p.ID],
		}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"players": response}})
}

// --- Saved Players (Scout+ feature) ---

// SavePlayer saves a player to user's favorites
func (m *ProfilesModule) SavePlayer(c *gin.Context) {
	playerID := c.Param("id")
	pid, err := uuid.Parse(playerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_UUID", "message": "Invalid player ID"}})
		return
	}

	userID, _ := c.Get("user_id")

	// Check if player exists
	var player domain.Player
	if err := m.db.First(&player, "id = ? AND deleted_at IS NULL", pid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "Player not found"}})
		return
	}

	var req struct {
		Notes    *string  `json:"notes,omitempty"`
		Tags     []string `json:"tags,omitempty"`
		Priority string   `json:"priority,omitempty"` // high, medium, low
	}
	_ = c.ShouldBindJSON(&req)

	// Default priority
	priority := "medium"
	if req.Priority == "high" || req.Priority == "low" {
		priority = req.Priority
	}

	saved := domain.SavedPlayer{
		UserID:    userID.(uuid.UUID),
		PlayerID:  pid,
		Notes:     req.Notes,
		Tags:      req.Tags,
		Priority:  priority,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := m.db.Create(&saved).Error; err != nil {
		// Check if already saved
		if m.db.Where("user_id = ? AND player_id = ?", userID, pid).First(&domain.SavedPlayer{}).Error == nil {
			c.JSON(http.StatusConflict, gin.H{"success": false, "error": gin.H{"code": "ALREADY_SAVED", "message": "Player already saved"}})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "SAVE_FAILED", "message": "Failed to save player"}})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "Player saved", "data": gin.H{"id": saved.ID}})
}

// UpdateSavedPlayer updates notes, tags, or priority for a saved player
func (m *ProfilesModule) UpdateSavedPlayer(c *gin.Context) {
	playerID := c.Param("id")
	pid, err := uuid.Parse(playerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_UUID", "message": "Invalid player ID"}})
		return
	}

	userID, _ := c.Get("user_id")

	var saved domain.SavedPlayer
	if err := m.db.Where("user_id = ? AND player_id = ?", userID, pid).First(&saved).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "Player not in saved list"}})
		return
	}

	var req struct {
		Notes    *string  `json:"notes"`
		Tags     []string `json:"tags"`
		Priority *string  `json:"priority"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
		return
	}

	updates := map[string]interface{}{
		"updated_at": time.Now(),
	}
	if req.Notes != nil {
		updates["notes"] = *req.Notes
	}
	if req.Tags != nil {
		updates["tags"] = req.Tags
	}
	if req.Priority != nil && (*req.Priority == "high" || *req.Priority == "medium" || *req.Priority == "low") {
		updates["priority"] = *req.Priority
	}

	m.db.Model(&saved).Updates(updates)

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Saved player updated"})
}

// UnsavePlayer removes a player from user's favorites
func (m *ProfilesModule) UnsavePlayer(c *gin.Context) {
	playerID := c.Param("id")
	pid, err := uuid.Parse(playerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_UUID", "message": "Invalid player ID"}})
		return
	}

	userID, _ := c.Get("user_id")

	result := m.db.Where("user_id = ? AND player_id = ?", userID, pid).Delete(&domain.SavedPlayer{})
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "Player not in saved list"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Player removed from saved"})
}

// GetSavedPlayers returns user's saved players with enhanced data
func (m *ProfilesModule) GetSavedPlayers(c *gin.Context) {
	userID, _ := c.Get("user_id")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset := (page - 1) * limit

	// Optional filters
	tagFilter := c.Query("tag")
	priorityFilter := c.Query("priority")

	var saved []domain.SavedPlayer
	var total int64

	query := m.db.Model(&domain.SavedPlayer{}).Where("user_id = ?", userID)

	if tagFilter != "" {
		query = query.Where("? = ANY(tags)", tagFilter)
	}
	if priorityFilter != "" {
		query = query.Where("priority = ?", priorityFilter)
	}

	query.Count(&total)
	m.db.Preload("Player.Academy").Where("user_id = ?", userID).
		Offset(offset).Limit(limit).Order("priority DESC, created_at DESC").Find(&saved)

	// Convert to response
	response := make([]gin.H, len(saved))
	for i, s := range saved {
		var profilePhotoURL, thumbnailURL *string
		if s.Player.ProfilePhotoURL != nil && *s.Player.ProfilePhotoURL != "" {
			url := m.getPresignedURL(*s.Player.ProfilePhotoURL)
			profilePhotoURL = &url
		}
		if s.Player.ThumbnailURL != nil && *s.Player.ThumbnailURL != "" {
			url := m.getPresignedURL(*s.Player.ThumbnailURL)
			thumbnailURL = &url
		}

		var academyName *string
		if s.Player.Academy != nil {
			academyName = &s.Player.Academy.Name
		}

		response[i] = gin.H{
			"id":         s.ID,
			"player_id":  s.PlayerID,
			"notes":      s.Notes,
			"tags":       s.Tags,
			"priority":   s.Priority,
			"saved_at":   s.CreatedAt,
			"updated_at": s.UpdatedAt,
			"player": gin.H{
				"first_name":        s.Player.FirstName,
				"last_name":         s.Player.LastName,
				"last_name_init":    s.Player.GetLastNameInit(),
				"age":               s.Player.GetAge(),
				"position":          s.Player.Position,
				"country":           s.Player.Country,
				"academy_name":      academyName,
				"profile_photo_url": profilePhotoURL,
				"thumbnail_url":     thumbnailURL,
				"is_verified":       s.Player.IsVerified(),
			},
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"saved_players": response,
			"pagination": gin.H{
				"page":        page,
				"limit":       limit,
				"total":       total,
				"total_pages": (total + int64(limit) - 1) / int64(limit),
			},
		},
	})
}

// GetMyTags returns the scout's custom tags
func (m *ProfilesModule) GetMyTags(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var tags []domain.ScoutTag
	m.db.Where("user_id = ?", userID).Order("name ASC").Find(&tags)

	response := make([]gin.H, len(tags))
	for i, t := range tags {
		response[i] = gin.H{
			"id":    t.ID,
			"name":  t.Name,
			"color": t.Color,
		}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"tags": response}})
}

// CreateTag creates a new custom tag
func (m *ProfilesModule) CreateTag(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var req struct {
		Name  string `json:"name" binding:"required,max=50"`
		Color string `json:"color"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
		return
	}

	color := "#6366f1"
	if req.Color != "" && len(req.Color) == 7 && req.Color[0] == '#' {
		color = req.Color
	}

	tag := domain.ScoutTag{
		UserID:    userID.(uuid.UUID),
		Name:      req.Name,
		Color:     color,
		CreatedAt: time.Now(),
	}

	if err := m.db.Create(&tag).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"success": false, "error": gin.H{"code": "ALREADY_EXISTS", "message": "Tag with this name already exists"}})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": gin.H{"id": tag.ID, "name": tag.Name, "color": tag.Color}})
}

// DeleteTag deletes a custom tag
func (m *ProfilesModule) DeleteTag(c *gin.Context) {
	tagID := c.Param("id")
	tid, err := uuid.Parse(tagID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_UUID", "message": "Invalid tag ID"}})
		return
	}

	userID, _ := c.Get("user_id")

	result := m.db.Where("id = ? AND user_id = ?", tid, userID).Delete(&domain.ScoutTag{})
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "Tag not found"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Tag deleted"})
}

// --- Contact Requests (Pro+ feature) ---

// CreateContactRequest creates a contact request for a player
func (m *ProfilesModule) CreateContactRequest(c *gin.Context) {
	playerID := c.Param("id")
	pid, err := uuid.Parse(playerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_UUID", "message": "Invalid player ID"}})
		return
	}

	userID, _ := c.Get("user_id")

	var req struct {
		Message string `json:"message" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
		return
	}

	// Check if player exists
	var player domain.Player
	if err := m.db.First(&player, "id = ? AND deleted_at IS NULL", pid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "Player not found"}})
		return
	}

	// Check for existing pending request
	var existingRequest domain.ContactRequest
	if err := m.db.Where("user_id = ? AND player_id = ? AND status IN ('pending', 'sent_to_academy')", userID, pid).First(&existingRequest).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"success": false, "error": gin.H{"code": "ALREADY_REQUESTED", "message": "You already have a pending contact request for this player"}})
		return
	}

	// Set follow-up reminder for 7 days from now
	followUpReminder := time.Now().Add(7 * 24 * time.Hour)

	contact := domain.ContactRequest{
		UserID:             userID.(uuid.UUID),
		PlayerID:           pid,
		Message:            req.Message,
		Status:             "pending",
		FollowUpReminderAt: &followUpReminder,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	if err := m.db.Create(&contact).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "CREATE_FAILED", "message": "Failed to create contact request"}})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data": gin.H{
			"id":                   contact.ID,
			"status":               contact.Status,
			"follow_up_reminder":   followUpReminder,
			"message":              "Contact request submitted. The academy will review and respond.",
			"expected_response_by": followUpReminder.Format("2006-01-02"),
		},
	})
}

// GetMyContactRequests returns user's contact requests with full tracking
func (m *ProfilesModule) GetMyContactRequests(c *gin.Context) {
	userID, _ := c.Get("user_id")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset := (page - 1) * limit
	statusFilter := c.Query("status")

	var requests []domain.ContactRequest
	var total int64

	query := m.db.Model(&domain.ContactRequest{}).Where("user_id = ?", userID)
	if statusFilter != "" {
		query = query.Where("status = ?", statusFilter)
	}
	query.Count(&total)

	m.db.Preload("Player.Academy").Where("user_id = ?", userID).Order("created_at DESC").
		Offset(offset).Limit(limit).Find(&requests)

	response := make([]gin.H, len(requests))
	for i, r := range requests {
		var profilePhotoURL *string
		if r.Player.ProfilePhotoURL != nil && *r.Player.ProfilePhotoURL != "" {
			url := m.getPresignedURL(*r.Player.ProfilePhotoURL)
			profilePhotoURL = &url
		}

		var academyName *string
		if r.Player.Academy != nil {
			academyName = &r.Player.Academy.Name
		}

		response[i] = gin.H{
			"id":                   r.ID,
			"player_id":            r.PlayerID,
			"message":              r.Message,
			"status":               r.Status,
			"academy_response":     r.AcademyResponse,
			"academy_responded_at": r.AcademyRespondedAt,
			"scout_read_at":        r.ScoutReadAt,
			"follow_up_reminder":   r.FollowUpReminderAt,
			"created_at":           r.CreatedAt,
			"updated_at":           r.UpdatedAt,
			"player": gin.H{
				"first_name":        r.Player.FirstName,
				"last_name":         r.Player.LastName,
				"last_name_init":    r.Player.GetLastNameInit(),
				"position":          r.Player.Position,
				"country":           r.Player.Country,
				"academy_name":      academyName,
				"profile_photo_url": profilePhotoURL,
			},
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"contact_requests": response,
			"pagination": gin.H{
				"page":        page,
				"limit":       limit,
				"total":       total,
				"total_pages": (total + int64(limit) - 1) / int64(limit),
			},
		},
	})
}

// MarkContactRequestRead marks a contact request response as read
func (m *ProfilesModule) MarkContactRequestRead(c *gin.Context) {
	requestID := c.Param("id")
	rid, err := uuid.Parse(requestID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_UUID", "message": "Invalid request ID"}})
		return
	}

	userID, _ := c.Get("user_id")

	var request domain.ContactRequest
	if err := m.db.Where("id = ? AND user_id = ?", rid, userID).First(&request).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "Contact request not found"}})
		return
	}

	now := time.Now()
	m.db.Model(&request).Updates(map[string]interface{}{
		"scout_read_at": now,
		"updated_at":    now,
	})

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Marked as read"})
}

// CancelContactRequest cancels a pending contact request
func (m *ProfilesModule) CancelContactRequest(c *gin.Context) {
	requestID := c.Param("id")
	rid, err := uuid.Parse(requestID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_UUID", "message": "Invalid request ID"}})
		return
	}

	userID, _ := c.Get("user_id")

	var request domain.ContactRequest
	if err := m.db.Where("id = ? AND user_id = ?", rid, userID).First(&request).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "Contact request not found"}})
		return
	}

	if request.Status != "pending" && request.Status != "sent_to_academy" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "CANNOT_CANCEL", "message": "Can only cancel pending requests"}})
		return
	}

	m.db.Model(&request).Updates(map[string]interface{}{
		"status":     "cancelled",
		"updated_at": time.Now(),
	})

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Contact request cancelled"})
}

// GetFeaturedPlayers returns featured players (for homepage)
func (m *ProfilesModule) GetFeaturedPlayers(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if limit > 20 {
		limit = 20
	}

	var players []domain.Player
	// Get verified players with photos, ordered by recent activity
	// Also include approved players if not enough verified
	m.db.Model(&domain.Player{}).
		Preload("Academy").
		Where("deleted_at IS NULL").
		Where("verification_status IN (?, ?)", "verified", "approved").
		Where("profile_photo_url IS NOT NULL AND profile_photo_url != ''").
		Order("CASE WHEN verification_status = 'verified' THEN 0 ELSE 1 END, created_at DESC").
		Limit(limit).
		Find(&players)

	response := make([]FeaturedPlayerResponse, len(players))
	for i, p := range players {
		dob := p.DateOfBirth.Format("2006-01-02")
		var academyName *string
		if p.Academy != nil {
			academyName = &p.Academy.Name
		}

		// Convert profile photo URL to presigned URL
		var profilePhotoURL *string
		if p.ProfilePhotoURL != nil && *p.ProfilePhotoURL != "" {
			url := m.getPresignedURL(*p.ProfilePhotoURL)
			profilePhotoURL = &url
		}

		response[i] = FeaturedPlayerResponse{
			ID:              p.ID,
			FirstName:       p.FirstName,
			LastName:        p.LastName,
			DateOfBirth:     dob,
			Age:             p.GetAge(),
			Position:        p.Position,
			Country:         p.Country,
			ProfilePhotoURL: profilePhotoURL,
			AcademyName:     academyName,
			IsVerified:      p.IsVerified(),
		}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"players": response}})
}

// AcademyListResponse is the public academy list response
type AcademyListResponse struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Country    string    `json:"country"`
	IsVerified bool      `json:"is_verified"`
}

// ListAcademies returns public list of academies for filtering
// @Summary List academies
// @Description Get a list of academies for filtering players
// @Tags Players
// @Produce json
// @Success 200 {object} map[string]interface{} "List of academies"
// @Router /academies [get]
func (m *ProfilesModule) ListAcademies(c *gin.Context) {
	var academies []domain.Academy

	// Get academies that have verified players
	m.db.Raw(`
		SELECT DISTINCT a.* FROM academies a
		INNER JOIN players p ON p.academy_id = a.id
		WHERE p.deleted_at IS NULL AND p.verification_status = 'verified'
		ORDER BY a.name ASC
	`).Scan(&academies)

	response := make([]AcademyListResponse, len(academies))
	for i, a := range academies {
		response[i] = AcademyListResponse{
			ID:         a.ID,
			Name:       a.Name,
			Country:    a.Country,
			IsVerified: a.IsVerified,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"academies": response,
		},
	})
}
