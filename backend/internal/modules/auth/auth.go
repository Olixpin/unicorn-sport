package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/unicorn-sport/backend/internal/domain"
)

// AuthModule handles authentication
type AuthModule struct {
	db              *gorm.DB
	jwtSecret       string
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

// NewAuthModule creates a new auth module
func NewAuthModule(db *gorm.DB, jwtSecret string, accessTTLMinutes, refreshTTLDays int) *AuthModule {
	return &AuthModule{
		db:              db,
		jwtSecret:       jwtSecret,
		accessTokenTTL:  time.Duration(accessTTLMinutes) * time.Minute,
		refreshTokenTTL: time.Duration(refreshTTLDays) * 24 * time.Hour,
	}
}

// Claims represents JWT claims
type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	Role   string    `json:"role"`
	jwt.RegisteredClaims
}

// --- Request/Response Types ---

// RegisterRequest represents scout registration request
type RegisterRequest struct {
	Email            string  `json:"email" binding:"required,email"`
	Password         string  `json:"password" binding:"required,min=8"`
	FirstName        string  `json:"first_name" binding:"required"`
	LastName         string  `json:"last_name" binding:"required"`
	OrganizationName *string `json:"organization_name,omitempty"`
	OrganizationType *string `json:"organization_type,omitempty"`
	Country          *string `json:"country,omitempty"`
}

// LoginRequest represents login request
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// AuthResponse represents authentication response
type AuthResponse struct {
	User         UserResponse          `json:"user"`
	Subscription *SubscriptionResponse `json:"subscription,omitempty"`
	AccessToken  string                `json:"access_token"`
	RefreshToken string                `json:"refresh_token"`
	ExpiresIn    int                   `json:"expires_in"`
}

// UserResponse represents user info in response
type UserResponse struct {
	ID            uuid.UUID `json:"id"`
	Email         string    `json:"email"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Role          string    `json:"role"`
	EmailVerified bool      `json:"email_verified"`
	CreatedAt     time.Time `json:"created_at"`
}

// SubscriptionResponse represents subscription info
type SubscriptionResponse struct {
	Tier      string     `json:"tier"`
	Status    string     `json:"status"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// --- Auth Endpoints ---

// Register handles scout self-registration
// @Summary Register a new scout
// @Description Create a new scout account. Only scouts can self-register; players are created by admin.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Registration details"
// @Success 201 {object} map[string]interface{} "User registered successfully"
// @Failure 400 {object} map[string]interface{} "Validation error"
// @Failure 409 {object} map[string]interface{} "Email already exists"
// @Router /auth/register [post]
func (a *AuthModule) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
		return
	}

	// Check if user already exists
	var existingUser domain.User
	if err := a.db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"success": false, "error": gin.H{"code": "EMAIL_EXISTS", "message": "Email already registered"}})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "HASH_FAILED", "message": "Failed to process password"}})
		return
	}

	// Create user
	user := domain.User{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Role:         "scout",
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := a.db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "CREATE_FAILED", "message": "Failed to create user"}})
		return
	}

	// Create scout profile
	scout := domain.Scout{
		UserID:           user.ID,
		OrganizationName: req.OrganizationName,
		OrganizationType: req.OrganizationType,
		Country:          req.Country,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
	a.db.Create(&scout)

	// Create free subscription
	subscription := domain.Subscription{
		UserID:    user.ID,
		Tier:      "free",
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	a.db.Create(&subscription)

	// Generate tokens
	accessToken, refreshToken, err := a.generateTokens(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "TOKEN_FAILED", "message": "Failed to generate tokens"}})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data": AuthResponse{
			User: UserResponse{
				ID:            user.ID,
				Email:         user.Email,
				FirstName:     user.FirstName,
				LastName:      user.LastName,
				Role:          user.Role,
				EmailVerified: user.EmailVerified,
				CreatedAt:     user.CreatedAt,
			},
			Subscription: &SubscriptionResponse{
				Tier:   "free",
				Status: "active",
			},
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiresIn:    int(a.accessTokenTTL.Seconds()),
		},
	})
}

// Login handles user login
// @Summary Login user
// @Description Authenticate a user and get access/refresh tokens
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login credentials"
// @Success 200 {object} map[string]interface{} "Login successful"
// @Failure 400 {object} map[string]interface{} "Validation error"
// @Failure 401 {object} map[string]interface{} "Invalid credentials"
// @Router /auth/login [post]
func (a *AuthModule) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
		return
	}

	// Find user
	var user domain.User
	if err := a.db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": gin.H{"code": "INVALID_CREDENTIALS", "message": "Invalid email or password"}})
		return
	}

	// Check if user is active
	if !user.IsActive {
		c.JSON(http.StatusForbidden, gin.H{"success": false, "error": gin.H{"code": "ACCOUNT_DISABLED", "message": "Account is disabled"}})
		return
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": gin.H{"code": "INVALID_CREDENTIALS", "message": "Invalid email or password"}})
		return
	}

	// Update last login
	now := time.Now()
	a.db.Model(&user).Update("last_login_at", now)

	// Get subscription
	var subscription domain.Subscription
	a.db.Where("user_id = ?", user.ID).First(&subscription)

	// Generate tokens
	accessToken, refreshToken, err := a.generateTokens(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "TOKEN_FAILED", "message": "Failed to generate tokens"}})
		return
	}

	response := AuthResponse{
		User: UserResponse{
			ID:            user.ID,
			Email:         user.Email,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			Role:          user.Role,
			EmailVerified: user.EmailVerified,
			CreatedAt:     user.CreatedAt,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int(a.accessTokenTTL.Seconds()),
	}

	if subscription.ID != uuid.Nil {
		response.Subscription = &SubscriptionResponse{
			Tier:      subscription.Tier,
			Status:    subscription.Status,
			ExpiresAt: subscription.CurrentPeriodEnd,
		}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": response})
}

// Refresh handles token refresh
func (a *AuthModule) Refresh(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
		return
	}

	// Hash the refresh token to find it in DB
	tokenHash := hashToken(req.RefreshToken)

	// Find refresh token
	var refreshToken domain.RefreshToken
	if err := a.db.Where("token_hash = ? AND expires_at > ?", tokenHash, time.Now()).First(&refreshToken).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": gin.H{"code": "INVALID_TOKEN", "message": "Invalid or expired refresh token"}})
		return
	}

	// Get user
	var user domain.User
	if err := a.db.First(&user, "id = ?", refreshToken.UserID).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": gin.H{"code": "USER_NOT_FOUND", "message": "User not found"}})
		return
	}

	// Delete old refresh token
	a.db.Delete(&refreshToken)

	// Generate new tokens
	accessToken, newRefreshToken, err := a.generateTokens(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "TOKEN_FAILED", "message": "Failed to generate tokens"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"access_token":  accessToken,
			"refresh_token": newRefreshToken,
			"expires_in":    int(a.accessTokenTTL.Seconds()),
		},
	})
}

// Logout invalidates the refresh token
func (a *AuthModule) Logout(c *gin.Context) {
	userID, _ := c.Get("user_id")

	// Delete all refresh tokens for user
	a.db.Where("user_id = ?", userID).Delete(&domain.RefreshToken{})

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Successfully logged out"})
}

// GetMe returns current user info
func (a *AuthModule) GetMe(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var user domain.User
	if err := a.db.First(&user, "id = ?", userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "User not found"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"id":             user.ID,
			"email":          user.Email,
			"role":           user.Role,
			"email_verified": user.EmailVerified,
			"is_active":      user.IsActive,
			"created_at":     user.CreatedAt,
			"last_login_at":  user.LastLoginAt,
		},
	})
}

// --- Email Verification ---

// VerifyEmailRequest represents email verification request
type VerifyEmailRequest struct {
	Code string `json:"code" binding:"required,len=6"`
}

// SendVerificationEmail sends/resends email verification code
func (a *AuthModule) SendVerificationEmail(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var user domain.User
	if err := a.db.First(&user, "id = ?", userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "User not found"}})
		return
	}

	if user.EmailVerified {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "ALREADY_VERIFIED", "message": "Email is already verified"}})
		return
	}

	// Delete existing unused codes
	a.db.Where("user_id = ? AND used_at IS NULL", userID).Delete(&domain.EmailVerificationToken{})

	// Generate 6-digit code
	code := generateVerificationCode()

	// Store code (expires in 15 minutes)
	verification := domain.EmailVerificationToken{
		UserID:    user.ID,
		Code:      code,
		ExpiresAt: time.Now().Add(15 * time.Minute),
		CreatedAt: time.Now(),
	}

	if err := a.db.Create(&verification).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "CREATE_FAILED", "message": "Failed to generate verification code"}})
		return
	}

	// TODO: Send email with verification code
	// In production, integrate with SendGrid/SES/Postmark
	log.Printf("📧 Email verification code for %s: %s", user.Email, code)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Verification code sent to your email",
		"data": gin.H{
			"expires_in": 900, // 15 minutes in seconds
		},
	})
}

// VerifyEmail verifies email with 6-digit code
func (a *AuthModule) VerifyEmail(c *gin.Context) {
	var req VerifyEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "VALIDATION_ERROR", "message": "Code must be exactly 6 digits"}})
		return
	}

	userID, _ := c.Get("user_id")

	var user domain.User
	if err := a.db.First(&user, "id = ?", userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "User not found"}})
		return
	}

	if user.EmailVerified {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "ALREADY_VERIFIED", "message": "Email is already verified"}})
		return
	}

	// Find valid, unused verification code
	var verification domain.EmailVerificationToken
	err := a.db.Where(
		"user_id = ? AND code = ? AND expires_at > ? AND used_at IS NULL",
		userID,
		req.Code,
		time.Now(),
	).First(&verification).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "INVALID_CODE",
				"message": "Invalid or expired verification code. Please request a new one.",
			},
		})
		return
	}

	// Transaction: Update user and mark code as used
	tx := a.db.Begin()

	// Mark email as verified
	if err := tx.Model(&user).Updates(map[string]interface{}{
		"email_verified": true,
		"updated_at":     time.Now(),
	}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "UPDATE_FAILED", "message": "Failed to verify email"}})
		return
	}

	// Mark code as used
	now := time.Now()
	if err := tx.Model(&verification).Update("used_at", now).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "UPDATE_FAILED", "message": "Failed to verify email"}})
		return
	}

	tx.Commit()

	log.Printf("✅ Email verified for user %s", user.Email)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Email verified successfully",
	})
}

// generateVerificationCode creates a 6-digit numeric code
func generateVerificationCode() string {
	bytes := make([]byte, 3)
	if _, err := rand.Read(bytes); err != nil {
		// Fallback to time-based code if crypto/rand fails
		return fmt.Sprintf("%06d", time.Now().UnixNano()%900000+100000)
	}
	// Convert to 6-digit number (100000-999999)
	num := int(bytes[0])<<16 | int(bytes[1])<<8 | int(bytes[2])
	code := 100000 + (num % 900000)
	return fmt.Sprintf("%06d", code)
}

// --- Token Generation ---

func (a *AuthModule) generateTokens(user domain.User) (string, string, error) {
	// Generate access token
	claims := Claims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(a.accessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   user.ID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(a.jwtSecret))
	if err != nil {
		return "", "", err
	}

	// Generate refresh token
	refreshTokenID := uuid.New()
	refreshTokenString := refreshTokenID.String()
	refreshTokenHash := hashToken(refreshTokenString)

	// Store refresh token in database
	refreshToken := domain.RefreshToken{
		UserID:    user.ID,
		TokenHash: refreshTokenHash,
		ExpiresAt: time.Now().Add(a.refreshTokenTTL),
		CreatedAt: time.Now(),
	}
	if err := a.db.Create(&refreshToken).Error; err != nil {
		return "", "", err
	}

	return accessToken, refreshTokenString, nil
}

func hashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}

// generateSecureToken creates a cryptographically secure random token
func generateSecureToken(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// --- Password Reset Flow ---

// ForgotPasswordRequest represents forgot password request
type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// ResetPasswordRequest represents reset password request
type ResetPasswordRequest struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

// ChangePasswordRequest represents authenticated password change
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required,min=8"`
}

// ForgotPassword initiates password reset flow
// UX Psychology: Always return success to prevent email enumeration attacks
// This protects user privacy while still providing helpful feedback
func (a *AuthModule) ForgotPassword(c *gin.Context) {
	var req ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "VALIDATION_ERROR",
				"message": err.Error(),
			},
		})
		return
	}

	// Find user by email (don't reveal if user exists or not)
	var user domain.User
	userExists := a.db.Where("email = ?", req.Email).First(&user).Error == nil

	if userExists && user.IsActive {
		// Invalidate any existing reset tokens for this user
		a.db.Where("user_id = ? AND used_at IS NULL", user.ID).Delete(&domain.PasswordResetToken{})

		// Generate secure reset token (32 bytes = 256 bits of entropy)
		resetToken, err := generateSecureToken(32)
		if err != nil {
			log.Printf("Failed to generate reset token: %v", err)
		} else {
			// Store hashed token (never store plain token in DB)
			tokenHash := hashToken(resetToken)
			resetRecord := domain.PasswordResetToken{
				UserID:    user.ID,
				TokenHash: tokenHash,
				ExpiresAt: time.Now().Add(1 * time.Hour), // 1 hour expiry
				CreatedAt: time.Now(),
			}

			if err := a.db.Create(&resetRecord).Error; err != nil {
				log.Printf("Failed to save reset token: %v", err)
			} else {
				// TODO: Send email with reset link
				// In production, integrate with SendGrid/SES/Postmark
				// Email should contain: https://yourdomain.com/reset-password?token={resetToken}
				log.Printf("📧 Password reset token created for %s (token: %s)", req.Email, resetToken)
			}
		}
	}

	// UX Best Practice: Always show success message (constant-time response)
	// This prevents attackers from determining which emails are registered
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "If an account with that email exists, you will receive password reset instructions shortly.",
		"data": gin.H{
			"email":      req.Email,
			"expires_in": 3600, // 1 hour in seconds
		},
	})
}

// ResetPassword completes password reset with token validation
func (a *AuthModule) ResetPassword(c *gin.Context) {
	var req ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "VALIDATION_ERROR",
				"message": err.Error(),
			},
		})
		return
	}

	// Hash the provided token to compare with stored hash
	tokenHash := hashToken(req.Token)

	// Find valid, unused reset token
	var resetToken domain.PasswordResetToken
	err := a.db.Where(
		"token_hash = ? AND expires_at > ? AND used_at IS NULL",
		tokenHash,
		time.Now(),
	).First(&resetToken).Error

	if err != nil {
		// Generic error message to prevent token enumeration
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "INVALID_TOKEN",
				"message": "This password reset link is invalid or has expired. Please request a new one.",
			},
		})
		return
	}

	// Validate password strength
	if len(req.NewPassword) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "WEAK_PASSWORD",
				"message": "Password must be at least 8 characters long",
			},
		})
		return
	}

	// Get user
	var user domain.User
	if err := a.db.First(&user, "id = ?", resetToken.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "USER_NOT_FOUND",
				"message": "Unable to reset password. Please contact support.",
			},
		})
		return
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "HASH_FAILED",
				"message": "Failed to process password",
			},
		})
		return
	}

	// Transaction: Update password and mark token as used
	tx := a.db.Begin()

	// Update user password
	if err := tx.Model(&user).Updates(map[string]interface{}{
		"password_hash": string(hashedPassword),
		"updated_at":    time.Now(),
	}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "UPDATE_FAILED",
				"message": "Failed to update password",
			},
		})
		return
	}

	// Mark token as used
	now := time.Now()
	if err := tx.Model(&resetToken).Update("used_at", now).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "TOKEN_UPDATE_FAILED",
				"message": "Failed to complete password reset",
			},
		})
		return
	}

	// Invalidate all refresh tokens (force re-login on all devices)
	tx.Where("user_id = ?", user.ID).Delete(&domain.RefreshToken{})

	tx.Commit()

	log.Printf("✅ Password reset completed for user %s", user.Email)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Password has been reset successfully. Please log in with your new password.",
	})
}

// ChangePassword allows authenticated users to change their password
func (a *AuthModule) ChangePassword(c *gin.Context) {
	var req ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "VALIDATION_ERROR",
				"message": err.Error(),
			},
		})
		return
	}

	userID, _ := c.Get("user_id")

	// Get user
	var user domain.User
	if err := a.db.First(&user, "id = ?", userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "NOT_FOUND",
				"message": "User not found",
			},
		})
		return
	}

	// Verify current password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.CurrentPassword)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "INVALID_PASSWORD",
				"message": "Current password is incorrect",
			},
		})
		return
	}

	// Validate new password is different
	if req.CurrentPassword == req.NewPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "SAME_PASSWORD",
				"message": "New password must be different from current password",
			},
		})
		return
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "HASH_FAILED",
				"message": "Failed to process password",
			},
		})
		return
	}

	// Update password
	if err := a.db.Model(&user).Updates(map[string]interface{}{
		"password_hash": string(hashedPassword),
		"updated_at":    time.Now(),
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "UPDATE_FAILED",
				"message": "Failed to update password",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Password changed successfully",
	})
}

// --- Profile Update ---

// UpdateProfileRequest represents profile update request
type UpdateProfileRequest struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Organization string `json:"organization"`
	Role         string `json:"role"`
}

// UpdateProfile allows authenticated users to update their profile
func (a *AuthModule) UpdateProfile(c *gin.Context) {
	var req UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "VALIDATION_ERROR",
				"message": err.Error(),
			},
		})
		return
	}

	userID, _ := c.Get("user_id")

	// Get user
	var user domain.User
	if err := a.db.First(&user, "id = ?", userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "NOT_FOUND",
				"message": "User not found",
			},
		})
		return
	}

	// Update profile fields
	updates := map[string]interface{}{
		"updated_at": time.Now(),
	}

	if req.FirstName != "" {
		updates["first_name"] = req.FirstName
	}
	if req.LastName != "" {
		updates["last_name"] = req.LastName
	}
	if req.Organization != "" {
		updates["organization_name"] = req.Organization
	}
	if req.Role != "" {
		updates["organization_type"] = req.Role
	}

	if err := a.db.Model(&user).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "UPDATE_FAILED",
				"message": "Failed to update profile",
			},
		})
		return
	}

	// Reload user
	a.db.First(&user, "id = ?", userID)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Profile updated successfully",
		"data": gin.H{
			"user": UserResponse{
				ID:            user.ID,
				Email:         user.Email,
				FirstName:     user.FirstName,
				LastName:      user.LastName,
				Role:          user.Role,
				EmailVerified: user.EmailVerified,
				CreatedAt:     user.CreatedAt,
			},
		},
	})
}

// UpdateNotificationsRequest represents notification preferences update
type UpdateNotificationsRequest struct {
	NewPlayers     bool `json:"newPlayers"`
	ContactUpdates bool `json:"contactUpdates"`
	WeeklyDigest   bool `json:"weeklyDigest"`
}

// UpdateNotifications allows users to update notification preferences
// Note: This is a stub - notification preferences would need to be stored in the database
func (a *AuthModule) UpdateNotifications(c *gin.Context) {
	var req UpdateNotificationsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": gin.H{
				"code":    "VALIDATION_ERROR",
				"message": err.Error(),
			},
		})
		return
	}

	// For now, just return success
	// In production, you'd store these preferences in the database
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Notification preferences updated",
		"data": gin.H{
			"newPlayers":     req.NewPlayers,
			"contactUpdates": req.ContactUpdates,
			"weeklyDigest":   req.WeeklyDigest,
		},
	})
}
