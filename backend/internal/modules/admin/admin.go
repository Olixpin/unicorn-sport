package admin

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/unicorn-sport/backend/internal/domain"
)

// AdminModule handles admin operations
type AdminModule struct {
	db *gorm.DB
}

// NewAdminModule creates a new admin module
func NewAdminModule(db *gorm.DB) *AdminModule {
	return &AdminModule{db: db}
}

// --- Dashboard Stats ---

// GetStats returns dashboard statistics for admin
func (m *AdminModule) GetStats(c *gin.Context) {
	var totalPlayers int64
	var verifiedPlayers int64
	var totalScouts int64
	var totalVideos int64
	var totalTournaments int64

	m.db.Model(&domain.Player{}).Count(&totalPlayers)
	m.db.Model(&domain.Player{}).Where("is_verified = ?", true).Count(&verifiedPlayers)
	m.db.Model(&domain.User{}).Where("role = ?", "scout").Count(&totalScouts)
	m.db.Model(&domain.Video{}).Count(&totalVideos)
	m.db.Model(&domain.Tournament{}).Count(&totalTournaments)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"total_players":    totalPlayers,
			"verified_players": verifiedPlayers,
			"total_scouts":     totalScouts,
			"total_videos":     totalVideos,
			"total_events":     totalTournaments,
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

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"player": player,
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

	m.logAudit(c, "update_player", "player", &pid, nil)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"id":         player.ID,
			"first_name": player.FirstName,
			"last_name":  player.LastName,
			"updated_at": time.Now(),
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

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"players": players,
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

func ptrUUID(u uuid.UUID) *uuid.UUID {
	return &u
}
