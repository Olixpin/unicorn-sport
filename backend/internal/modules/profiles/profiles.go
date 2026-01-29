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
	ID            uuid.UUID `json:"id"`
	FirstName     string    `json:"first_name"`
	LastNameInit  string    `json:"last_name_init"`
	Age           int       `json:"age"`
	Position      string    `json:"position"`
	Country       string    `json:"country"`
	HeightCm      *int      `json:"height_cm,omitempty"`
	PreferredFoot *string   `json:"preferred_foot,omitempty"`
	ThumbnailURL  *string   `json:"thumbnail_url,omitempty"`
	IsVerified    bool      `json:"is_verified"`
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
	ID              uuid.UUID       `json:"id"`
	FirstName       string          `json:"first_name"`
	LastNameInit    string          `json:"last_name_init"`
	Age             int             `json:"age"`
	Position        string          `json:"position"`
	PreferredFoot   *string         `json:"preferred_foot,omitempty"`
	HeightCm        *int            `json:"height_cm,omitempty"`
	WeightKg        *int            `json:"weight_kg,omitempty"`
	Country         string          `json:"country"`
	State           *string         `json:"state,omitempty"`
	SchoolName      *string         `json:"school_name,omitempty"`
	TournamentName  *string         `json:"tournament_name,omitempty"`
	ProfilePhotoURL *string         `json:"profile_photo_url,omitempty"`
	IsVerified      bool            `json:"is_verified"`
	HighlightVideos []VideoResponse `json:"highlight_videos"`
	FullMatchVideos []VideoResponse `json:"full_match_videos"`
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

		response[i] = PlayerListResponse{
			ID:            p.ID,
			FirstName:     p.FirstName,
			LastNameInit:  p.GetLastNameInit(),
			Age:           p.GetAge(),
			Position:      p.Position,
			Country:       p.Country,
			HeightCm:      p.HeightCm,
			PreferredFoot: p.PreferredFoot,
			ThumbnailURL:  thumbnailURL,
			IsVerified:    p.IsVerified(),
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
	if err := m.db.Preload("Tournament").First(&player, "id = ? AND deleted_at IS NULL", pid).Error; err != nil {
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
		LastNameInit:    player.GetLastNameInit(),
		Age:             player.GetAge(),
		Position:        player.Position,
		PreferredFoot:   player.PreferredFoot,
		HeightCm:        player.HeightCm,
		WeightKg:        player.WeightKg,
		Country:         player.Country,
		State:           player.State,
		SchoolName:      player.SchoolName,
		ProfilePhotoURL: profilePhotoURL,
		IsVerified:      player.IsVerified(),
		HighlightVideos: make([]VideoResponse, len(highlightVideos)),
		FullMatchVideos: make([]VideoResponse, len(fullMatchVideos)),
	}

	if player.Tournament != nil {
		response.TournamentName = &player.Tournament.Name
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
		Notes *string `json:"notes,omitempty"`
	}
	_ = c.ShouldBindJSON(&req) // Ignore error - notes are optional

	saved := domain.SavedPlayer{
		UserID:    userID.(uuid.UUID),
		PlayerID:  pid,
		Notes:     req.Notes,
		CreatedAt: time.Now(),
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

	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "Player saved"})
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

// GetSavedPlayers returns user's saved players
func (m *ProfilesModule) GetSavedPlayers(c *gin.Context) {
	userID, _ := c.Get("user_id")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset := (page - 1) * limit

	var saved []domain.SavedPlayer
	var total int64

	m.db.Model(&domain.SavedPlayer{}).Where("user_id = ?", userID).Count(&total)
	m.db.Preload("Player").Where("user_id = ?", userID).
		Offset(offset).Limit(limit).Order("created_at DESC").Find(&saved)

	// Convert to response
	response := make([]gin.H, len(saved))
	for i, s := range saved {
		response[i] = gin.H{
			"id":        s.ID,
			"player_id": s.PlayerID,
			"notes":     s.Notes,
			"saved_at":  s.CreatedAt,
			"player": gin.H{
				"first_name":     s.Player.FirstName,
				"last_name_init": s.Player.GetLastNameInit(),
				"age":            s.Player.GetAge(),
				"position":       s.Player.Position,
				"country":        s.Player.Country,
				"thumbnail_url":  s.Player.ThumbnailURL,
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

	contact := domain.ContactRequest{
		UserID:    userID.(uuid.UUID),
		PlayerID:  pid,
		Message:   req.Message,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	if err := m.db.Create(&contact).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "CREATE_FAILED", "message": "Failed to create contact request"}})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data": gin.H{
			"id":      contact.ID,
			"status":  contact.Status,
			"message": "Contact request submitted. The academy will review and forward to the player.",
		},
	})
}

// GetMyContactRequests returns user's contact requests
func (m *ProfilesModule) GetMyContactRequests(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var requests []domain.ContactRequest
	m.db.Preload("Player").Where("user_id = ?", userID).Order("created_at DESC").Find(&requests)

	response := make([]gin.H, len(requests))
	for i, r := range requests {
		response[i] = gin.H{
			"id":         r.ID,
			"player_id":  r.PlayerID,
			"message":    r.Message,
			"status":     r.Status,
			"created_at": r.CreatedAt,
			"player": gin.H{
				"first_name":     r.Player.FirstName,
				"last_name_init": r.Player.GetLastNameInit(),
			},
		}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"contact_requests": response}})
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
