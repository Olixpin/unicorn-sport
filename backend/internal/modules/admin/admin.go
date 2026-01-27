package admin

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/unicorn-sport/backend/internal/domain"
)

// AdminModule handles admin operations
type AdminModule struct {
	db       *gorm.DB
	s3Client *s3.Client
	s3Bucket string
}

// NewAdminModule creates a new admin module
func NewAdminModule(db *gorm.DB, s3Client *s3.Client, s3Bucket string) *AdminModule {
	return &AdminModule{
		db:       db,
		s3Client: s3Client,
		s3Bucket: s3Bucket,
	}
}

// --- Dashboard Stats ---

// GetStats returns dashboard statistics for admin
func (m *AdminModule) GetStats(c *gin.Context) {
	var totalPlayers int64
	var verifiedPlayers int64
	var pendingPlayers int64
	var totalAcademies int64
	var verifiedAcademies int64
	var totalScouts int64
	var totalVideos int64
	var totalTournaments int64

	// Current counts
	m.db.Model(&domain.Player{}).Count(&totalPlayers)
	m.db.Model(&domain.Player{}).Where("verification_status = ?", "verified").Count(&verifiedPlayers)
	m.db.Model(&domain.Player{}).Where("verification_status = ?", "pending").Count(&pendingPlayers)
	m.db.Model(&domain.Academy{}).Count(&totalAcademies)
	m.db.Model(&domain.Academy{}).Where("is_verified = ?", true).Count(&verifiedAcademies)
	m.db.Model(&domain.User{}).Where("role = ?", "scout").Count(&totalScouts)
	m.db.Model(&domain.Video{}).Count(&totalVideos)
	m.db.Model(&domain.Tournament{}).Count(&totalTournaments)

	// Calculate growth (last 30 days vs previous 30 days)
	now := time.Now()
	thirtyDaysAgo := now.AddDate(0, 0, -30)
	sixtyDaysAgo := now.AddDate(0, 0, -60)

	var playersLast30 int64
	var playersPrev30 int64
	var academiesLast30 int64
	var academiesPrev30 int64
	var scoutsLast30 int64
	var scoutsPrev30 int64

	m.db.Model(&domain.Player{}).Where("created_at >= ?", thirtyDaysAgo).Count(&playersLast30)
	m.db.Model(&domain.Player{}).Where("created_at >= ? AND created_at < ?", sixtyDaysAgo, thirtyDaysAgo).Count(&playersPrev30)
	m.db.Model(&domain.Academy{}).Where("created_at >= ?", thirtyDaysAgo).Count(&academiesLast30)
	m.db.Model(&domain.Academy{}).Where("created_at >= ? AND created_at < ?", sixtyDaysAgo, thirtyDaysAgo).Count(&academiesPrev30)
	m.db.Model(&domain.User{}).Where("role = ? AND created_at >= ?", "scout", thirtyDaysAgo).Count(&scoutsLast30)
	m.db.Model(&domain.User{}).Where("role = ? AND created_at >= ? AND created_at < ?", "scout", sixtyDaysAgo, thirtyDaysAgo).Count(&scoutsPrev30)

	// Calculate percentage growth
	calcGrowth := func(current, previous int64) float64 {
		if previous == 0 {
			if current > 0 {
				return 100.0
			}
			return 0.0
		}
		return float64(current-previous) / float64(previous) * 100
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"total_players":      totalPlayers,
			"verified_players":   verifiedPlayers,
			"pending_players":    pendingPlayers,
			"total_academies":    totalAcademies,
			"verified_academies": verifiedAcademies,
			"total_scouts":       totalScouts,
			"total_videos":       totalVideos,
			"total_events":       totalTournaments,
			"growth": gin.H{
				"players":   calcGrowth(playersLast30, playersPrev30),
				"academies": calcGrowth(academiesLast30, academiesPrev30),
				"scouts":    calcGrowth(scoutsLast30, scoutsPrev30),
			},
		},
	})
}

// --- Academy Management ---

// CreateAcademyRequest represents the request to create an academy
type CreateAcademyRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description,omitempty"`
	Country     string  `json:"country" binding:"required"`
	State       *string `json:"state,omitempty"`
	City        *string `json:"city,omitempty"`
	Address     *string `json:"address,omitempty"`
	Phone       *string `json:"phone,omitempty"`
	Email       *string `json:"email,omitempty"`
	Website     *string `json:"website,omitempty"`
	LogoURL     *string `json:"logo_url,omitempty"`
	FoundedYear *int    `json:"founded_year,omitempty"`
	IsVerified  bool    `json:"is_verified"`
}

// ListAcademies returns all academies with optional filtering and pagination
func (m *AdminModule) ListAcademies(c *gin.Context) {
	var academies []domain.Academy

	// Pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "12"))
	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 12
	}
	offset := (page - 1) * perPage

	// Build query
	query := m.db.Model(&domain.Academy{})

	// Search filter (name, city, or country)
	if search := c.Query("q"); search != "" {
		searchPattern := "%" + strings.ToLower(search) + "%"
		query = query.Where("LOWER(name) LIKE ? OR LOWER(city) LIKE ? OR LOWER(country) LIKE ?", searchPattern, searchPattern, searchPattern)
	}

	// Country filter (case-insensitive)
	if country := c.Query("country"); country != "" {
		query = query.Where("LOWER(country) = LOWER(?)", country)
	}

	// Status filter
	if status := c.Query("status"); status != "" {
		switch status {
		case "verified":
			query = query.Where("is_verified = ?", true)
		case "pending":
			query = query.Where("is_verified = ?", false)
		}
	}

	// Get total count before pagination
	var total int64
	if err := query.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to count academies",
		})
		return
	}

	// Fetch academies with pagination
	if err := query.Order("created_at DESC").Offset(offset).Limit(perPage).Find(&academies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to fetch academies",
		})
		return
	}

	// Get global stats (unfiltered) for dashboard cards
	var totalAcademies int64
	var verifiedCount int64
	var pendingCount int64
	var totalPlayers int64

	m.db.Model(&domain.Academy{}).Count(&totalAcademies)
	m.db.Model(&domain.Academy{}).Where("is_verified = ?", true).Count(&verifiedCount)
	m.db.Model(&domain.Academy{}).Where("is_verified = ?", false).Count(&pendingCount)

	// Sum player counts from all academies
	type PlayerCountResult struct {
		Total int64
	}
	var playerResult PlayerCountResult
	m.db.Model(&domain.Academy{}).Select("COALESCE(SUM(player_count), 0) as total").Scan(&playerResult)
	totalPlayers = playerResult.Total

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"academies": academies,
			"total":     total,
			"page":      page,
			"per_page":  perPage,
			"stats": gin.H{
				"total_academies": totalAcademies,
				"verified_count":  verifiedCount,
				"pending_count":   pendingCount,
				"total_players":   totalPlayers,
			},
		},
	})
}

// CreateAcademy creates a new academy
func (m *AdminModule) CreateAcademy(c *gin.Context) {
	var req CreateAcademyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request: " + err.Error(),
		})
		return
	}

	academy := domain.Academy{
		Name:        req.Name,
		Description: req.Description,
		Country:     req.Country,
		State:       req.State,
		City:        req.City,
		Address:     req.Address,
		Phone:       req.Phone,
		Email:       req.Email,
		Website:     req.Website,
		LogoURL:     req.LogoURL,
		FoundedYear: req.FoundedYear,
		IsVerified:  req.IsVerified,
	}

	if err := m.db.Create(&academy).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to create academy: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data": gin.H{
			"academy": academy,
		},
	})
}

// GetAcademy returns a single academy by ID
func (m *AdminModule) GetAcademy(c *gin.Context) {
	id := c.Param("id")

	var academy domain.Academy
	if err := m.db.First(&academy, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Academy not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    academy,
	})
}

// UpdateAcademy updates an existing academy
func (m *AdminModule) UpdateAcademy(c *gin.Context) {
	id := c.Param("id")

	var academy domain.Academy
	if err := m.db.First(&academy, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Academy not found",
		})
		return
	}

	var req CreateAcademyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request: " + err.Error(),
		})
		return
	}

	academy.Name = req.Name
	academy.Description = req.Description
	academy.Country = req.Country
	academy.State = req.State
	academy.City = req.City
	academy.Address = req.Address
	academy.Phone = req.Phone
	academy.Email = req.Email
	academy.Website = req.Website
	academy.LogoURL = req.LogoURL
	academy.FoundedYear = req.FoundedYear
	academy.IsVerified = req.IsVerified

	if err := m.db.Save(&academy).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to update academy: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    academy,
	})
}

// DeleteAcademy deletes an academy
func (m *AdminModule) DeleteAcademy(c *gin.Context) {
	id := c.Param("id")

	result := m.db.Delete(&domain.Academy{}, "id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to delete academy",
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Academy not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Academy deleted successfully",
	})
}

// --- Contact Request Management ---

// ListContactRequests returns contact requests filtered by status
func (m *AdminModule) ListContactRequests(c *gin.Context) {
	status := c.Query("status")

	var requests []domain.ContactRequest
	query := m.db.Preload("Player")

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Order("created_at DESC").Find(&requests).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to fetch contact requests",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"requests": requests,
		},
	})
}

// ApproveContactRequest approves a contact request
func (m *AdminModule) ApproveContactRequest(c *gin.Context) {
	id := c.Param("id")

	var request domain.ContactRequest
	if err := m.db.First(&request, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Contact request not found",
		})
		return
	}

	now := time.Now()
	request.Status = "approved"
	request.HandledAt = &now

	if err := m.db.Save(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to approve contact request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    request,
	})
}

// RejectContactRequest rejects a contact request
func (m *AdminModule) RejectContactRequest(c *gin.Context) {
	id := c.Param("id")

	var request domain.ContactRequest
	if err := m.db.First(&request, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Contact request not found",
		})
		return
	}

	now := time.Now()
	request.Status = "rejected"
	request.HandledAt = &now

	if err := m.db.Save(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to reject contact request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    request,
	})
}

// --- Player Management ---

// CreatePlayerRequest represents the request to create a player
type CreatePlayerRequest struct {
	FirstName           string  `json:"first_name" binding:"required"`
	LastName            string  `json:"last_name" binding:"required"`
	DateOfBirth         string  `json:"date_of_birth" binding:"required"`
	Position            string  `json:"position" binding:"required"`
	PreferredFoot       *string `json:"preferred_foot,omitempty"`
	HeightCm            *int    `json:"height_cm,omitempty"`
	WeightKg            *int    `json:"weight_kg,omitempty"`
	Country             string  `json:"country" binding:"required"`
	State               *string `json:"state,omitempty"`
	City                *string `json:"city,omitempty"`
	AcademyID           *string `json:"academy_id,omitempty"`
	SchoolName          *string `json:"school_name,omitempty"`
	TournamentID        *string `json:"tournament_id,omitempty"`
	VerificationStatus  string  `json:"verification_status,omitempty"`
	VerificationDocType *string `json:"verification_doc_type,omitempty"`
}

// CreatePlayerResponse includes generated credentials
type CreatePlayerResponse struct {
	ID          uuid.UUID    `json:"id"`
	FirstName   string       `json:"first_name"`
	LastName    string       `json:"last_name"`
	Credentials *Credentials `json:"credentials,omitempty"`
	CreatedAt   time.Time    `json:"created_at"`
}

// Credentials for player login
type Credentials struct {
	Email        string `json:"email"`
	TempPassword string `json:"temp_password"`
}

// CreatePlayer creates a new player profile (admin only)
func (m *AdminModule) CreatePlayer(c *gin.Context) {
	var req CreatePlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
		return
	}

	// Parse date of birth
	dob, err := time.Parse("2006-01-02", req.DateOfBirth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_DATE", "message": "Invalid date_of_birth format, use YYYY-MM-DD"}})
		return
	}

	// Get admin user ID from context
	adminID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": gin.H{"code": "UNAUTHORIZED", "message": "User not found"}})
		return
	}

	// Parse tournament ID if provided
	var tournamentID *uuid.UUID
	var tournamentYear *int
	if req.TournamentID != nil && *req.TournamentID != "" {
		tid, err := uuid.Parse(*req.TournamentID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_UUID", "message": "Invalid tournament_id"}})
			return
		}
		tournamentID = &tid

		// Get tournament year
		var tournament domain.Tournament
		if err := m.db.First(&tournament, "id = ?", tid).Error; err == nil {
			tournamentYear = &tournament.Year
		}
	}

	// Parse academy ID if provided
	var academyID *uuid.UUID
	if req.AcademyID != nil && *req.AcademyID != "" {
		aid, err := uuid.Parse(*req.AcademyID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_UUID", "message": "Invalid academy_id"}})
			return
		}
		// Verify academy exists
		var academy domain.Academy
		if err := m.db.First(&academy, "id = ?", aid).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "Academy not found"}})
			return
		}
		academyID = &aid
	}

	// Set verification status
	verificationStatus := "pending"
	if req.VerificationStatus != "" {
		verificationStatus = req.VerificationStatus
	}

	// Create player
	player := domain.Player{
		FirstName:           req.FirstName,
		LastName:            req.LastName,
		DateOfBirth:         dob,
		Position:            req.Position,
		PreferredFoot:       req.PreferredFoot,
		HeightCm:            req.HeightCm,
		WeightKg:            req.WeightKg,
		Country:             req.Country,
		State:               req.State,
		City:                req.City,
		AcademyID:           academyID,
		SchoolName:          req.SchoolName,
		TournamentID:        tournamentID,
		TournamentYear:      tournamentYear,
		VerificationStatus:  verificationStatus,
		VerificationDocType: req.VerificationDocType,
		CreatedBy:           adminID.(uuid.UUID),
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}

	// If verified, set verification metadata
	if verificationStatus == "verified" {
		now := time.Now()
		aid := adminID.(uuid.UUID)
		player.VerifiedAt = &now
		player.VerifiedBy = &aid
	}

	if err := m.db.Create(&player).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "CREATE_FAILED", "message": "Failed to create player"}})
		return
	}

	// Generate credentials for the player
	credentials, err := m.generatePlayerCredentials(&player)
	if err != nil {
		// Player created but credentials failed - log but don't fail request
		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"data": CreatePlayerResponse{
				ID:        player.ID,
				FirstName: player.FirstName,
				LastName:  player.LastName,
				CreatedAt: player.CreatedAt,
			},
		})
		return
	}

	// Log audit
	m.logAudit(c, "create_player", "player", &player.ID, nil)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data": CreatePlayerResponse{
			ID:          player.ID,
			FirstName:   player.FirstName,
			LastName:    player.LastName,
			Credentials: credentials,
			CreatedAt:   player.CreatedAt,
		},
	})
}

// generatePlayerCredentials creates user account and credentials for a player
func (m *AdminModule) generatePlayerCredentials(player *domain.Player) (*Credentials, error) {
	// Generate email: firstname.lastname.year@unicornsport.africa
	year := player.DateOfBirth.Year()
	email := fmt.Sprintf("%s.%s.%d@unicornsport.africa",
		strings.ToLower(player.FirstName),
		strings.ToLower(player.LastName),
		year)

	// Check if email already exists, append number if needed
	var count int64
	m.db.Model(&domain.User{}).Where("email LIKE ?", strings.TrimSuffix(email, "@unicornsport.africa")+"%").Count(&count)
	if count > 0 {
		email = fmt.Sprintf("%s.%s.%d.%d@unicornsport.africa",
			strings.ToLower(player.FirstName),
			strings.ToLower(player.LastName),
			year, count+1)
	}

	// Generate temporary password
	tempPassword := generateTempPassword()

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(tempPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user account
	user := domain.User{
		Email:        email,
		PasswordHash: string(hashedPassword),
		Role:         "player",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := m.db.Create(&user).Error; err != nil {
		return nil, err
	}

	// Link player to user
	player.UserID = &user.ID
	if err := m.db.Save(player).Error; err != nil {
		return nil, err
	}

	return &Credentials{
		Email:        email,
		TempPassword: tempPassword,
	}, nil
}

// UpdatePlayerRequest represents the request to update a player
type UpdatePlayerRequest struct {
	FirstName          *string `json:"first_name,omitempty"`
	LastName           *string `json:"last_name,omitempty"`
	Position           *string `json:"position,omitempty"`
	PreferredFoot      *string `json:"preferred_foot,omitempty"`
	HeightCm           *int    `json:"height_cm,omitempty"`
	WeightKg           *int    `json:"weight_kg,omitempty"`
	Country            *string `json:"country,omitempty"`
	State              *string `json:"state,omitempty"`
	City               *string `json:"city,omitempty"`
	AcademyID          *string `json:"academy_id,omitempty"`
	SchoolName         *string `json:"school_name,omitempty"`
	VerificationStatus *string `json:"verification_status,omitempty"`
	ProfilePhotoURL    *string `json:"profile_photo_url,omitempty"`
	ThumbnailURL       *string `json:"thumbnail_url,omitempty"`
}

// getPresignedURL converts an s3:// URL to a presigned URL for direct access
func (m *AdminModule) getPresignedURL(s3Url string) string {
	if s3Url == "" || m.s3Client == nil {
		return s3Url
	}

	// Handle s3://bucket/key format
	if strings.HasPrefix(s3Url, "s3://") {
		parts := strings.SplitN(strings.TrimPrefix(s3Url, "s3://"), "/", 2)
		if len(parts) != 2 {
			return s3Url
		}
		bucket := parts[0]
		key := parts[1]

		presigner := s3.NewPresignClient(m.s3Client)
		presignedReq, err := presigner.PresignGetObject(context.Background(), &s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		}, s3.WithPresignExpires(time.Hour))
		if err != nil {
			return s3Url
		}
		return presignedReq.URL
	}

	// Already a direct URL or CloudFront URL
	return s3Url
}

// GetPlayer returns a single player by ID
func (m *AdminModule) GetPlayer(c *gin.Context) {
	playerID := c.Param("id")
	pid, err := uuid.Parse(playerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_UUID", "message": "Invalid player ID"}})
		return
	}

	var player domain.Player
	if err := m.db.Preload("Tournament").Preload("Academy").Where("deleted_at IS NULL").First(&player, "id = ?", pid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "Player not found"}})
		return
	}

	// Convert profile photo URL to presigned URL
	var profilePhotoURL *string
	if player.ProfilePhotoURL != nil && *player.ProfilePhotoURL != "" {
		url := m.getPresignedURL(*player.ProfilePhotoURL)
		profilePhotoURL = &url
	}

	// Build response with presigned URL
	playerResponse := gin.H{
		"id":                  player.ID,
		"first_name":          player.FirstName,
		"last_name":           player.LastName,
		"date_of_birth":       player.DateOfBirth,
		"position":            player.Position,
		"preferred_foot":      player.PreferredFoot,
		"height_cm":           player.HeightCm,
		"weight_kg":           player.WeightKg,
		"country":             player.Country,
		"state":               player.State,
		"city":                player.City,
		"school_name":         player.SchoolName,
		"verification_status": player.VerificationStatus,
		"profile_photo_url":   profilePhotoURL,
		"thumbnail_url":       player.ThumbnailURL,
		"academy_id":          player.AcademyID,
		"tournament_id":       player.TournamentID,
		"tournament":          player.Tournament,
		"academy":             player.Academy,
		"created_at":          player.CreatedAt,
		"updated_at":          player.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"player": playerResponse,
		},
	})
}

// UpdatePlayer updates a player profile
func (m *AdminModule) UpdatePlayer(c *gin.Context) {
	playerID := c.Param("id")
	pid, err := uuid.Parse(playerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_UUID", "message": "Invalid player ID"}})
		return
	}

	var player domain.Player
	if err := m.db.First(&player, "id = ?", pid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "Player not found"}})
		return
	}

	var req UpdatePlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
		return
	}

	// Build updates map
	updates := make(map[string]interface{})
	if req.FirstName != nil {
		updates["first_name"] = *req.FirstName
	}
	if req.LastName != nil {
		updates["last_name"] = *req.LastName
	}
	if req.Position != nil {
		updates["position"] = *req.Position
	}
	if req.PreferredFoot != nil {
		updates["preferred_foot"] = *req.PreferredFoot
	}
	if req.HeightCm != nil {
		updates["height_cm"] = *req.HeightCm
	}
	if req.WeightKg != nil {
		updates["weight_kg"] = *req.WeightKg
	}
	if req.Country != nil {
		updates["country"] = *req.Country
	}
	if req.State != nil {
		updates["state"] = *req.State
	}
	if req.City != nil {
		updates["city"] = *req.City
	}
	if req.AcademyID != nil {
		if *req.AcademyID == "" {
			// Allow clearing the academy
			updates["academy_id"] = nil
		} else {
			aid, err := uuid.Parse(*req.AcademyID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_UUID", "message": "Invalid academy_id"}})
				return
			}
			// Verify academy exists
			var academy domain.Academy
			if err := m.db.First(&academy, "id = ?", aid).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "Academy not found"}})
				return
			}
			updates["academy_id"] = aid
		}
	}
	if req.SchoolName != nil {
		updates["school_name"] = *req.SchoolName
	}
	if req.ProfilePhotoURL != nil {
		updates["profile_photo_url"] = *req.ProfilePhotoURL
	}
	if req.ThumbnailURL != nil {
		updates["thumbnail_url"] = *req.ThumbnailURL
	}
	if req.VerificationStatus != nil {
		updates["verification_status"] = *req.VerificationStatus
		if *req.VerificationStatus == "verified" {
			adminID, _ := c.Get("user_id")
			now := time.Now()
			updates["verified_at"] = now
			updates["verified_by"] = adminID
		}
	}
	updates["updated_at"] = time.Now()

	if err := m.db.Model(&player).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "UPDATE_FAILED", "message": "Failed to update player"}})
		return
	}

	// Reload the player to get updated values
	m.db.First(&player, "id = ?", pid)

	// Convert profile photo URL to presigned URL
	var profilePhotoURL *string
	if player.ProfilePhotoURL != nil && *player.ProfilePhotoURL != "" {
		url := m.getPresignedURL(*player.ProfilePhotoURL)
		profilePhotoURL = &url
	}

	m.logAudit(c, "update_player", "player", &pid, nil)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"player": gin.H{
				"id":                player.ID,
				"first_name":        player.FirstName,
				"last_name":         player.LastName,
				"profile_photo_url": profilePhotoURL,
				"updated_at":        player.UpdatedAt,
			},
		},
	})
}

// DeletePlayer soft deletes a player
func (m *AdminModule) DeletePlayer(c *gin.Context) {
	playerID := c.Param("id")
	pid, err := uuid.Parse(playerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_UUID", "message": "Invalid player ID"}})
		return
	}

	var player domain.Player
	if err := m.db.First(&player, "id = ?", pid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "Player not found"}})
		return
	}

	// Soft delete
	now := time.Now()
	if err := m.db.Model(&player).Update("deleted_at", now).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "DELETE_FAILED", "message": "Failed to delete player"}})
		return
	}

	m.logAudit(c, "delete_player", "player", &pid, nil)

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Player deleted successfully"})
}

// ListPlayers lists all players with pagination (admin view)
func (m *AdminModule) ListPlayers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if limit > 100 {
		limit = 100
	}
	offset := (page - 1) * limit

	var players []domain.Player
	var total int64

	query := m.db.Model(&domain.Player{}).Where("deleted_at IS NULL")

	// Search filter (name, position, or country)
	if search := c.Query("search"); search != "" {
		searchPattern := "%" + strings.ToLower(search) + "%"
		query = query.Where("LOWER(first_name) LIKE ? OR LOWER(last_name) LIKE ? OR LOWER(position) LIKE ? OR LOWER(country) LIKE ?", searchPattern, searchPattern, searchPattern, searchPattern)
	}

	// Apply filters
	if country := c.Query("country"); country != "" {
		query = query.Where("LOWER(country) = LOWER(?)", country)
	}
	if position := c.Query("position"); position != "" {
		query = query.Where("position = ?", position)
	}
	if verified := c.Query("verified"); verified != "" {
		if verified == "true" {
			query = query.Where("verification_status = ?", "verified")
		} else {
			query = query.Where("verification_status != ?", "verified")
		}
	}
	if tournamentID := c.Query("tournament_id"); tournamentID != "" {
		query = query.Where("tournament_id = ?", tournamentID)
	}
	if academyID := c.Query("academy_id"); academyID != "" {
		query = query.Where("academy_id = ?", academyID)
	}

	query.Count(&total)
	query.Preload("Tournament").Preload("Academy").Offset(offset).Limit(limit).Order("created_at DESC").Find(&players)

	// Get stats (unfiltered counts)
	var totalPlayers, verifiedCount, pendingCount int64
	m.db.Model(&domain.Player{}).Where("deleted_at IS NULL").Count(&totalPlayers)
	m.db.Model(&domain.Player{}).Where("deleted_at IS NULL AND verification_status = ?", "verified").Count(&verifiedCount)
	pendingCount = totalPlayers - verifiedCount

	// Convert profile photo URLs to presigned URLs
	playerResponses := make([]gin.H, len(players))
	for i, p := range players {
		var profilePhotoURL *string
		if p.ProfilePhotoURL != nil && *p.ProfilePhotoURL != "" {
			url := m.getPresignedURL(*p.ProfilePhotoURL)
			profilePhotoURL = &url
		}
		playerResponses[i] = gin.H{
			"id":                  p.ID,
			"first_name":          p.FirstName,
			"last_name":           p.LastName,
			"date_of_birth":       p.DateOfBirth,
			"position":            p.Position,
			"preferred_foot":      p.PreferredFoot,
			"height_cm":           p.HeightCm,
			"weight_kg":           p.WeightKg,
			"country":             p.Country,
			"state":               p.State,
			"city":                p.City,
			"school_name":         p.SchoolName,
			"verification_status": p.VerificationStatus,
			"profile_photo_url":   profilePhotoURL,
			"thumbnail_url":       p.ThumbnailURL,
			"academy_id":          p.AcademyID,
			"tournament_id":       p.TournamentID,
			"tournament":          p.Tournament,
			"academy":             p.Academy,
			"created_at":          p.CreatedAt,
			"updated_at":          p.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"players": playerResponses,
			"total":   total,
			"stats": gin.H{
				"total_players":  totalPlayers,
				"verified_count": verifiedCount,
				"pending_count":  pendingCount,
			},
			"pagination": gin.H{
				"page":        page,
				"limit":       limit,
				"total":       total,
				"total_pages": (total + int64(limit) - 1) / int64(limit),
			},
		},
	})
}

// --- Tournament Management ---

// CreateTournamentRequest represents the request to create a tournament
type CreateTournamentRequest struct {
	Name        string  `json:"name" binding:"required"`
	Year        int     `json:"year" binding:"required"`
	Location    *string `json:"location,omitempty"`
	Country     *string `json:"country,omitempty"`
	StartDate   *string `json:"start_date,omitempty"`
	EndDate     *string `json:"end_date,omitempty"`
	Description *string `json:"description,omitempty"`
}

// CreateTournament creates a new tournament
func (m *AdminModule) CreateTournament(c *gin.Context) {
	var req CreateTournamentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
		return
	}

	adminID, _ := c.Get("user_id")

	tournament := domain.Tournament{
		Name:        req.Name,
		Year:        req.Year,
		Location:    req.Location,
		Country:     req.Country,
		Description: req.Description,
		Status:      "upcoming",
		CreatedBy:   ptrUUID(adminID.(uuid.UUID)),
		CreatedAt:   time.Now(),
	}

	// Parse dates if provided
	if req.StartDate != nil {
		if t, err := time.Parse("2006-01-02", *req.StartDate); err == nil {
			tournament.StartDate = &t
		}
	}
	if req.EndDate != nil {
		if t, err := time.Parse("2006-01-02", *req.EndDate); err == nil {
			tournament.EndDate = &t
		}
	}

	if err := m.db.Create(&tournament).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "CREATE_FAILED", "message": "Failed to create tournament"}})
		return
	}

	m.logAudit(c, "create_tournament", "tournament", &tournament.ID, nil)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data": gin.H{
			"id":         tournament.ID,
			"name":       tournament.Name,
			"created_at": tournament.CreatedAt,
		},
	})
}

// ListTournaments lists all tournaments
func (m *AdminModule) ListTournaments(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset := (page - 1) * limit

	var tournaments []domain.Tournament
	var total int64

	query := m.db.Model(&domain.Tournament{})

	// Apply filters
	if year := c.Query("year"); year != "" {
		query = query.Where("year = ?", year)
	}
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	query.Count(&total)
	query.Offset(offset).Limit(limit).Order("year DESC, created_at DESC").Find(&tournaments)

	// Get player counts for each tournament
	type TournamentWithCount struct {
		domain.Tournament
		PlayerCount int64 `json:"player_count"`
	}

	results := make([]TournamentWithCount, len(tournaments))
	for i, t := range tournaments {
		var count int64
		m.db.Model(&domain.Player{}).Where("tournament_id = ? AND deleted_at IS NULL", t.ID).Count(&count)
		results[i] = TournamentWithCount{Tournament: t, PlayerCount: count}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"events": results,
			"pagination": gin.H{
				"page":        page,
				"limit":       limit,
				"total":       total,
				"total_pages": (total + int64(limit) - 1) / int64(limit),
			},
		},
	})
}

// UpdateTournament updates a tournament
func (m *AdminModule) UpdateTournament(c *gin.Context) {
	tournamentID := c.Param("id")
	tid, err := uuid.Parse(tournamentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_UUID", "message": "Invalid tournament ID"}})
		return
	}

	var tournament domain.Tournament
	if err := m.db.First(&tournament, "id = ?", tid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "Tournament not found"}})
		return
	}

	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
		return
	}

	if err := m.db.Model(&tournament).Updates(req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "UPDATE_FAILED", "message": "Failed to update tournament"}})
		return
	}

	m.logAudit(c, "update_tournament", "tournament", &tid, nil)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"id":         tournament.ID,
			"name":       tournament.Name,
			"updated_at": time.Now(),
		},
	})
}

// --- User Management ---

// ListUsers lists all users (admin only)
func (m *AdminModule) ListUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset := (page - 1) * limit

	var total int64

	query := m.db.Table("users").
		Select("users.id, users.email, users.role, users.email_verified, users.is_active, users.created_at, users.last_login_at, scouts.organization_name, subscriptions.tier as subscription_tier").
		Joins("LEFT JOIN scouts ON scouts.user_id = users.id").
		Joins("LEFT JOIN subscriptions ON subscriptions.user_id = users.id")

	// Apply filters
	if role := c.Query("role"); role != "" {
		query = query.Where("users.role = ?", role)
	}
	if tier := c.Query("subscription_tier"); tier != "" {
		query = query.Where("subscriptions.tier = ?", tier)
	}

	query.Count(&total)

	var results []map[string]interface{}
	query.Offset(offset).Limit(limit).Order("users.created_at DESC").Scan(&results)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"users": results,
			"pagination": gin.H{
				"page":        page,
				"limit":       limit,
				"total":       total,
				"total_pages": (total + int64(limit) - 1) / int64(limit),
			},
		},
	})
}

// UpdateUser updates a user's status
func (m *AdminModule) UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	uid, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_UUID", "message": "Invalid user ID"}})
		return
	}

	var user domain.User
	if err := m.db.First(&user, "id = ?", uid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "User not found"}})
		return
	}

	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
		return
	}

	// Only allow updating specific fields
	updates := make(map[string]interface{})
	if val, ok := req["is_active"]; ok {
		updates["is_active"] = val
	}
	if val, ok := req["role"]; ok {
		updates["role"] = val
	}
	updates["updated_at"] = time.Now()

	if err := m.db.Model(&user).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "UPDATE_FAILED", "message": "Failed to update user"}})
		return
	}

	m.logAudit(c, "update_user", "user", &uid, nil)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"id":         user.ID,
			"is_active":  user.IsActive,
			"updated_at": time.Now(),
		},
	})
}

// --- Helper Functions ---

func (m *AdminModule) logAudit(c *gin.Context, action, resourceType string, resourceID *uuid.UUID, details *string) {
	userID, _ := c.Get("user_id")
	uid := userID.(uuid.UUID)
	ip := c.ClientIP()

	audit := domain.AuditLog{
		UserID:       &uid,
		Action:       action,
		ResourceType: resourceType,
		ResourceID:   resourceID,
		Details:      details,
		IPAddress:    &ip,
		CreatedAt:    time.Now(),
	}
	m.db.Create(&audit)
}

func generateTempPassword() string {
	const charset = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789!@#$"
	password := make([]byte, 12)
	for i := range password {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		password[i] = charset[num.Int64()]
	}
	return string(password)
}

// ==================== AUDIT LOGS ====================

// ListAuditLogs returns paginated audit logs
func (m *AdminModule) ListAuditLogs(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	action := c.Query("action")
	resourceType := c.Query("resource_type")
	userID := c.Query("user_id")

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}
	offset := (page - 1) * limit

	query := m.db.Model(&domain.AuditLog{})

	if action != "" {
		query = query.Where("action = ?", action)
	}
	if resourceType != "" {
		query = query.Where("resource_type = ?", resourceType)
	}
	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}

	var total int64
	query.Count(&total)

	var logs []domain.AuditLog
	query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&logs)

	// Get user info for each log
	var userIDs []uuid.UUID
	for _, log := range logs {
		if log.UserID != nil {
			userIDs = append(userIDs, *log.UserID)
		}
	}

	userMap := make(map[uuid.UUID]gin.H)
	if len(userIDs) > 0 {
		var users []domain.User
		m.db.Where("id IN ?", userIDs).Find(&users)
		for _, u := range users {
			userMap[u.ID] = gin.H{
				"id":    u.ID,
				"email": u.Email,
				"name":  u.FirstName + " " + u.LastName,
			}
		}
	}

	response := make([]gin.H, len(logs))
	for i, log := range logs {
		var user gin.H
		if log.UserID != nil {
			user = userMap[*log.UserID]
		}
		response[i] = gin.H{
			"id":            log.ID,
			"user":          user,
			"action":        log.Action,
			"resource_type": log.ResourceType,
			"resource_id":   log.ResourceID,
			"details":       log.Details,
			"ip_address":    log.IPAddress,
			"created_at":    log.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"logs":  response,
			"total": total,
			"page":  page,
			"limit": limit,
		},
	})
}

// ==================== BULK OPERATIONS ====================

// BulkPlayersRequest is the request for bulk player operations
type BulkPlayersRequest struct {
	PlayerIDs []string `json:"player_ids" binding:"required,min=1"`
	Action    string   `json:"action" binding:"required,oneof=delete verify unverify"`
}

// BulkUpdatePlayers handles bulk operations on players
func (m *AdminModule) BulkUpdatePlayers(c *gin.Context) {
	var req BulkPlayersRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	// Convert string IDs to UUIDs
	var playerIDs []uuid.UUID
	for _, idStr := range req.PlayerIDs {
		id, err := uuid.Parse(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid player ID: " + idStr})
			return
		}
		playerIDs = append(playerIDs, id)
	}

	var result *gorm.DB
	var auditAction string

	switch req.Action {
	case "delete":
		result = m.db.Model(&domain.Player{}).Where("id IN ?", playerIDs).Update("deleted_at", time.Now())
		auditAction = "bulk_delete_players"
	case "verify":
		result = m.db.Model(&domain.Player{}).Where("id IN ?", playerIDs).Update("verification_status", "verified")
		auditAction = "bulk_verify_players"
	case "unverify":
		result = m.db.Model(&domain.Player{}).Where("id IN ?", playerIDs).Update("verification_status", "pending")
		auditAction = "bulk_unverify_players"
	default:
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid action"})
		return
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to update players"})
		return
	}

	details := fmt.Sprintf(`{"count": %d, "player_ids": %v}`, len(playerIDs), req.PlayerIDs)
	m.logAudit(c, auditAction, "players", nil, &details)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": fmt.Sprintf("Successfully updated %d players", result.RowsAffected),
		"data": gin.H{
			"affected_count": result.RowsAffected,
		},
	})
}

// ==================== SETTINGS ====================

// GetSettings returns all platform settings
func (m *AdminModule) GetSettings(c *gin.Context) {
	var settings []domain.Setting
	m.db.Find(&settings)

	settingsMap := make(map[string]string)
	for _, s := range settings {
		settingsMap[s.Key] = s.Value
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    settingsMap,
	})
}

// UpdateSettingsRequest is the request for updating settings
type UpdateSettingsRequest struct {
	Settings map[string]string `json:"settings" binding:"required"`
}

// UpdateSettings updates platform settings
func (m *AdminModule) UpdateSettings(c *gin.Context) {
	var req UpdateSettingsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	for key, value := range req.Settings {
		setting := domain.Setting{
			Key:       key,
			Value:     value,
			UpdatedAt: time.Now(),
		}
		m.db.Where("key = ?", key).Assign(setting).FirstOrCreate(&setting)
	}

	m.logAudit(c, "update_settings", "settings", nil, nil)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Settings updated successfully",
	})
}

// ==================== EXPORT DATA ====================

// ExportPlayers exports players as CSV
func (m *AdminModule) ExportPlayers(c *gin.Context) {
	var players []domain.Player
	m.db.Where("deleted_at IS NULL").
		Preload("Academy").
		Order("created_at DESC").
		Find(&players)

	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment; filename=players_export.csv")

	// Write CSV header
	header := "ID,First Name,Last Name,Date of Birth,Position,Country,State,City,Academy,Verification Status,Created At\n"
	c.Writer.WriteString(header)

	// Write data rows
	for _, p := range players {
		academyName := ""
		if p.Academy != nil {
			academyName = p.Academy.Name
		}
		state := ""
		if p.State != nil {
			state = *p.State
		}
		city := ""
		if p.City != nil {
			city = *p.City
		}
		row := fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s\n",
			p.ID.String(),
			escapeCsv(p.FirstName),
			escapeCsv(p.LastName),
			p.DateOfBirth.Format("2006-01-02"),
			escapeCsv(p.Position),
			escapeCsv(p.Country),
			escapeCsv(state),
			escapeCsv(city),
			escapeCsv(academyName),
			p.VerificationStatus,
			p.CreatedAt.Format("2006-01-02 15:04:05"),
		)
		c.Writer.WriteString(row)
	}

	m.logAudit(c, "export_players", "players", nil, nil)
}

// ExportUsers exports users as CSV
func (m *AdminModule) ExportUsers(c *gin.Context) {
	// Get users with their subscriptions
	type userWithSub struct {
		domain.User
		SubscriptionTier string
	}
	var users []userWithSub
	m.db.Table("users").
		Select("users.*, COALESCE(subscriptions.tier, 'free') as subscription_tier").
		Joins("LEFT JOIN subscriptions ON subscriptions.user_id = users.id").
		Order("users.created_at DESC").
		Scan(&users)

	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment; filename=users_export.csv")

	// Write CSV header
	header := "ID,First Name,Last Name,Email,Role,Subscription Tier,Is Active,Created At\n"
	c.Writer.WriteString(header)

	// Write data rows
	for _, u := range users {
		row := fmt.Sprintf("%s,%s,%s,%s,%s,%s,%t,%s\n",
			u.ID.String(),
			escapeCsv(u.FirstName),
			escapeCsv(u.LastName),
			escapeCsv(u.Email),
			u.Role,
			u.SubscriptionTier,
			u.IsActive,
			u.CreatedAt.Format("2006-01-02 15:04:05"),
		)
		c.Writer.WriteString(row)
	}

	m.logAudit(c, "export_users", "users", nil, nil)
}

func escapeCsv(s string) string {
	if strings.ContainsAny(s, ",\"\n") {
		return "\"" + strings.ReplaceAll(s, "\"", "\"\"") + "\""
	}
	return s
}

// ==================== ANALYTICS ====================

// GetAnalytics returns time-series analytics data
func (m *AdminModule) GetAnalytics(c *gin.Context) {
	days, _ := strconv.Atoi(c.DefaultQuery("days", "30"))
	if days < 7 {
		days = 7
	}
	if days > 365 {
		days = 365
	}

	startDate := time.Now().AddDate(0, 0, -days)

	// Signups over time
	type dailyCount struct {
		Date  string `json:"date"`
		Count int    `json:"count"`
	}

	var userSignups []dailyCount
	m.db.Model(&domain.User{}).
		Select("DATE(created_at) as date, COUNT(*) as count").
		Where("created_at >= ?", startDate).
		Group("DATE(created_at)").
		Order("date").
		Find(&userSignups)

	var playerSignups []dailyCount
	m.db.Model(&domain.Player{}).
		Select("DATE(created_at) as date, COUNT(*) as count").
		Where("created_at >= ? AND deleted_at IS NULL", startDate).
		Group("DATE(created_at)").
		Order("date").
		Find(&playerSignups)

	var videoUploads []dailyCount
	m.db.Model(&domain.Video{}).
		Select("DATE(created_at) as date, COUNT(*) as count").
		Where("created_at >= ?", startDate).
		Group("DATE(created_at)").
		Order("date").
		Find(&videoUploads)

	var highlightUploads []dailyCount
	m.db.Model(&domain.PlayerHighlight{}).
		Select("DATE(created_at) as date, COUNT(*) as count").
		Where("created_at >= ?", startDate).
		Group("DATE(created_at)").
		Order("date").
		Find(&highlightUploads)

	var contactRequests []dailyCount
	m.db.Model(&domain.ContactRequest{}).
		Select("DATE(created_at) as date, COUNT(*) as count").
		Where("created_at >= ?", startDate).
		Group("DATE(created_at)").
		Order("date").
		Find(&contactRequests)

	// Subscription breakdown - query the Subscription model
	type tierCount struct {
		Tier  string `json:"subscription_tier"`
		Count int    `json:"count"`
	}
	var subscriptionTiers []tierCount
	m.db.Model(&domain.Subscription{}).
		Select("tier, COUNT(*) as count").
		Where("tier IS NOT NULL AND tier != ''").
		Group("tier").
		Find(&subscriptionTiers)

	// Player positions breakdown
	type labelCount struct {
		Label string `json:"subscription_tier"`
		Count int    `json:"count"`
	}
	var positionCounts []labelCount
	m.db.Model(&domain.Player{}).
		Select("position as label, COUNT(*) as count").
		Where("deleted_at IS NULL AND position != ''").
		Group("position").
		Order("count DESC").
		Limit(10).
		Find(&positionCounts)

	// Player countries breakdown
	var countryCounts []labelCount
	m.db.Model(&domain.Player{}).
		Select("country as label, COUNT(*) as count").
		Where("deleted_at IS NULL AND country != ''").
		Group("country").
		Order("count DESC").
		Limit(10).
		Find(&countryCounts)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"user_signups":       userSignups,
			"player_signups":     playerSignups,
			"video_uploads":      videoUploads,
			"highlight_uploads":  highlightUploads,
			"contact_requests":   contactRequests,
			"subscription_tiers": subscriptionTiers,
			"player_positions":   positionCounts,
			"player_countries":   countryCounts,
		},
	})
}

func ptrUUID(u uuid.UUID) *uuid.UUID {
	return &u
}
