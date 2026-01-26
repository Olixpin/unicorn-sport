package contact

import (
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/unicorn-sport/backend/internal/domain"
)

// ContactModule handles public contact form submissions
type ContactModule struct {
	db *gorm.DB
}

// NewContactModule creates a new contact module
func NewContactModule(db *gorm.DB) *ContactModule {
	return &ContactModule{db: db}
}

// ContactRequest represents a public contact form submission
type ContactRequest struct {
	Name    string  `json:"name" binding:"required,min=2,max=100"`
	Email   string  `json:"email" binding:"required,email"`
	Subject *string `json:"subject,omitempty"`
	Message string  `json:"message" binding:"required,min=10,max=5000"`
	Type    *string `json:"type,omitempty"` // general, partnership, media, support
}

// SubmitContact handles public contact form submissions
// UX Psychology: Immediate feedback, clear validation, spam protection
func (m *ContactModule) SubmitContact(c *gin.Context) {
	var req ContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "VALIDATION_ERROR",
				"message": formatValidationError(err.Error()),
			},
		})
		return
	}

	// Sanitize inputs
	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	req.Message = strings.TrimSpace(req.Message)

	// Basic spam detection (can be enhanced with reCAPTCHA in production)
	if isSpamMessage(req.Message) {
		log.Printf("⚠️ Potential spam detected from %s", req.Email)
		// Still return success to not reveal spam detection
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Thank you for your message. We'll get back to you soon!",
		})
		return
	}

	// Validate contact type
	contactType := "general"
	if req.Type != nil {
		validTypes := map[string]bool{"general": true, "partnership": true, "media": true, "support": true}
		if validTypes[*req.Type] {
			contactType = *req.Type
		}
	}

	// Rate limiting check: max 3 submissions per email per hour
	var recentCount int64
	oneHourAgo := time.Now().Add(-1 * time.Hour)
	m.db.Model(&domain.GeneralContactRequest{}).
		Where("email = ? AND created_at > ?", req.Email, oneHourAgo).
		Count(&recentCount)

	if recentCount >= 3 {
		c.JSON(http.StatusTooManyRequests, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "RATE_LIMITED",
				"message": "You've submitted too many messages recently. Please try again later.",
			},
		})
		return
	}

	// Get client info for audit
	ipAddress := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	// Create contact request
	contact := domain.GeneralContactRequest{
		Name:      req.Name,
		Email:     req.Email,
		Subject:   req.Subject,
		Message:   req.Message,
		Type:      contactType,
		Status:    "new",
		IPAddress: &ipAddress,
		UserAgent: &userAgent,
		CreatedAt: time.Now(),
	}

	if err := m.db.Create(&contact).Error; err != nil {
		log.Printf("Failed to save contact request: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "SAVE_FAILED",
				"message": "Failed to submit your message. Please try again.",
			},
		})
		return
	}

	log.Printf("📬 New contact request from %s <%s> (type: %s)", req.Name, req.Email, contactType)

	// UX Best Practice: Warm, human response with clear expectations
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Thank you for reaching out! We've received your message and will respond within 24-48 hours.",
		"data": gin.H{
			"id":         contact.ID,
			"name":       contact.Name,
			"email":      contact.Email,
			"type":       contact.Type,
			"created_at": contact.CreatedAt,
		},
	})
}

// formatValidationError converts binding errors to user-friendly messages
func formatValidationError(err string) string {
	switch {
	case strings.Contains(err, "'Name'") && strings.Contains(err, "required"):
		return "Please provide your name"
	case strings.Contains(err, "'Name'") && strings.Contains(err, "min"):
		return "Name must be at least 2 characters"
	case strings.Contains(err, "'Email'") && strings.Contains(err, "required"):
		return "Please provide your email address"
	case strings.Contains(err, "'Email'") && strings.Contains(err, "email"):
		return "Please provide a valid email address"
	case strings.Contains(err, "'Message'") && strings.Contains(err, "required"):
		return "Please provide a message"
	case strings.Contains(err, "'Message'") && strings.Contains(err, "min"):
		return "Message must be at least 10 characters"
	case strings.Contains(err, "'Message'") && strings.Contains(err, "max"):
		return "Message is too long (max 5000 characters)"
	default:
		return "Please check your input and try again"
	}
}

// isSpamMessage performs basic spam detection
func isSpamMessage(message string) bool {
	lowerMsg := strings.ToLower(message)

	// Common spam patterns
	spamPatterns := []string{
		"click here",
		"buy now",
		"limited offer",
		"act now",
		"free money",
		"bitcoin",
		"crypto invest",
		"work from home",
		"make money fast",
		"casino",
		"viagra",
		"lottery winner",
	}

	for _, pattern := range spamPatterns {
		if strings.Contains(lowerMsg, pattern) {
			return true
		}
	}

	// Check for excessive URLs
	urlPattern := regexp.MustCompile(`https?://`)
	if len(urlPattern.FindAllString(message, -1)) > 3 {
		return true
	}

	return false
}
