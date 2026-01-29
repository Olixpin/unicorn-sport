package search

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/unicorn-sport/backend/internal/domain"
)

// SearchModule handles search operations
type SearchModule struct {
	db       *gorm.DB
	s3Client *s3.Client
	s3Bucket string
}

// NewSearchModule creates a new search module
func NewSearchModule(db *gorm.DB, s3Client *s3.Client, s3Bucket string) *SearchModule {
	return &SearchModule{
		db:       db,
		s3Client: s3Client,
		s3Bucket: s3Bucket,
	}
}

// getPresignedURL converts an s3:// URL to a presigned URL
func (m *SearchModule) getPresignedURL(s3URL string) string {
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

// PlayerSearchResult represents a search result
type PlayerSearchResult struct {
	ID                string  `json:"id"`
	FirstName         string  `json:"first_name"`
	LastName          string  `json:"last_name"`
	LastNameInit      string  `json:"last_name_init"`
	Age               int     `json:"age"`
	Position          string  `json:"position"`
	Country           string  `json:"country"`
	State             *string `json:"state,omitempty"`
	HeightCm          *int    `json:"height_cm,omitempty"`
	PreferredFoot     *string `json:"preferred_foot,omitempty"`
	ThumbnailURL      *string `json:"thumbnail_url,omitempty"`
	ProfilePhotoURL   *string `json:"profile_photo_url,omitempty"`
	VideoThumbnailURL *string `json:"video_thumbnail_url,omitempty"`
	IsVerified        bool    `json:"is_verified"`
	TournamentName    *string `json:"tournament_name,omitempty"`
	AcademyName       *string `json:"academy_name,omitempty"`
	VideoCount        int     `json:"video_count"`
}

// SearchPlayers performs full-text search on players
func (m *SearchModule) SearchPlayers(c *gin.Context) {
	query := c.Query("q")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if limit > 100 {
		limit = 100
	}
	offset := (page - 1) * limit

	var players []domain.Player
	var total int64

	dbQuery := m.db.Model(&domain.Player{}).
		Preload("Tournament").
		Preload("Academy").
		Preload("Highlights").
		Where("deleted_at IS NULL").
		Where("verification_status = ?", "verified")

	// Full-text search if query provided
	if query != "" {
		// Use ILIKE for simple text matching (PostgreSQL specific)
		searchPattern := "%" + query + "%"
		dbQuery = dbQuery.Where(
			"first_name ILIKE ? OR last_name ILIKE ? OR position ILIKE ? OR country ILIKE ? OR school_name ILIKE ?",
			searchPattern, searchPattern, searchPattern, searchPattern, searchPattern,
		)
	}

	// Apply filters
	if position := c.Query("position"); position != "" {
		dbQuery = dbQuery.Where("position ILIKE ?", "%"+position+"%")
	}
	if country := c.Query("country"); country != "" {
		dbQuery = dbQuery.Where("country = ?", country)
	}
	if state := c.Query("state"); state != "" {
		dbQuery = dbQuery.Where("state = ?", state)
	}
	if foot := c.Query("preferred_foot"); foot != "" {
		dbQuery = dbQuery.Where("preferred_foot = ?", foot)
	}
	if ageMin := c.Query("age_min"); ageMin != "" {
		if minAge, err := strconv.Atoi(ageMin); err == nil {
			maxBirthDate := time.Now().AddDate(-minAge, 0, 0)
			dbQuery = dbQuery.Where("date_of_birth <= ?", maxBirthDate)
		}
	}
	if ageMax := c.Query("age_max"); ageMax != "" {
		if maxAge, err := strconv.Atoi(ageMax); err == nil {
			minBirthDate := time.Now().AddDate(-maxAge-1, 0, 0)
			dbQuery = dbQuery.Where("date_of_birth > ?", minBirthDate)
		}
	}
	if tournamentID := c.Query("tournament_id"); tournamentID != "" {
		dbQuery = dbQuery.Where("tournament_id = ?", tournamentID)
	}
	if heightMin := c.Query("height_min"); heightMin != "" {
		if h, err := strconv.Atoi(heightMin); err == nil {
			dbQuery = dbQuery.Where("height_cm >= ?", h)
		}
	}
	if heightMax := c.Query("height_max"); heightMax != "" {
		if h, err := strconv.Atoi(heightMax); err == nil {
			dbQuery = dbQuery.Where("height_cm <= ?", h)
		}
	}

	// Get total count
	dbQuery.Count(&total)

	// Apply sorting
	sortBy := c.DefaultQuery("sort", "created_at")
	sortOrder := c.DefaultQuery("order", "desc")
	allowedSorts := map[string]bool{
		"created_at":    true,
		"first_name":    true,
		"position":      true,
		"country":       true,
		"date_of_birth": true,
	}
	if !allowedSorts[sortBy] {
		sortBy = "created_at"
	}
	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "desc"
	}
	dbQuery = dbQuery.Order(sortBy + " " + sortOrder)

	// Fetch results
	dbQuery.Offset(offset).Limit(limit).Find(&players)

	// Convert to search results
	results := make([]PlayerSearchResult, len(players))
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

		// Get video thumbnail from first highlight if available (action shot > static photo)
		var videoThumbnailURL *string
		if len(p.Highlights) > 0 && p.Highlights[0].ThumbnailURL != nil && *p.Highlights[0].ThumbnailURL != "" {
			url := m.getPresignedURL(*p.Highlights[0].ThumbnailURL)
			videoThumbnailURL = &url
		}

		results[i] = PlayerSearchResult{
			ID:                p.ID.String(),
			FirstName:         p.FirstName,
			LastName:          p.LastName,
			LastNameInit:      p.GetLastNameInit(),
			Age:               p.GetAge(),
			Position:          p.Position,
			Country:           p.Country,
			State:             p.State,
			HeightCm:          p.HeightCm,
			PreferredFoot:     p.PreferredFoot,
			ThumbnailURL:      thumbnailURL,
			ProfilePhotoURL:   profilePhotoURL,
			VideoThumbnailURL: videoThumbnailURL,
			IsVerified:        p.IsVerified(),
			VideoCount:        len(p.Highlights),
		}
		if p.Tournament != nil {
			results[i].TournamentName = &p.Tournament.Name
		}
		if p.Academy != nil {
			results[i].AcademyName = &p.Academy.Name
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"players": results,
			"filters_applied": gin.H{
				"query":    query,
				"position": c.Query("position"),
				"country":  c.Query("country"),
				"age_min":  c.Query("age_min"),
				"age_max":  c.Query("age_max"),
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

// GetFilterOptions returns available filter options
func (m *SearchModule) GetFilterOptions(c *gin.Context) {
	// Get distinct countries
	var countries []string
	m.db.Model(&domain.Player{}).
		Where("deleted_at IS NULL AND verification_status = ?", "verified").
		Distinct("country").
		Pluck("country", &countries)

	// Get distinct positions
	var positions []string
	m.db.Model(&domain.Player{}).
		Where("deleted_at IS NULL AND verification_status = ?", "verified").
		Distinct("position").
		Pluck("position", &positions)

	// Get tournaments
	var tournaments []domain.Tournament
	m.db.Order("year DESC").Find(&tournaments)

	tournamentOptions := make([]gin.H, len(tournaments))
	for i, t := range tournaments {
		tournamentOptions[i] = gin.H{
			"id":   t.ID,
			"name": t.Name,
			"year": t.Year,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"countries":      countries,
			"positions":      positions,
			"tournaments":    tournamentOptions,
			"preferred_foot": []string{"left", "right", "both"},
		},
	})
}

// GetStats returns platform statistics
func (m *SearchModule) GetStats(c *gin.Context) {
	var playerCount int64
	var verifiedCount int64
	var tournamentCount int64
	var videoCount int64

	m.db.Model(&domain.Player{}).Where("deleted_at IS NULL").Count(&playerCount)
	m.db.Model(&domain.Player{}).Where("deleted_at IS NULL AND verification_status = ?", "verified").Count(&verifiedCount)
	m.db.Model(&domain.Tournament{}).Count(&tournamentCount)
	m.db.Model(&domain.Video{}).Where("video_type != ?", "pending").Count(&videoCount)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"total_players":    playerCount,
			"verified_players": verifiedCount,
			"tournaments":      tournamentCount,
			"videos":           videoCount,
		},
	})
}
