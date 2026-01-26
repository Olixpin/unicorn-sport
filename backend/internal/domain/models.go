package domain

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user account
type User struct {
	ID            uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Email         string     `json:"email" gorm:"uniqueIndex;not null"`
	PasswordHash  string     `json:"-" gorm:"not null"`
	FirstName     string     `json:"first_name" gorm:"not null"`
	LastName      string     `json:"last_name" gorm:"not null"`
	Role          string     `json:"role" gorm:"not null;check:role IN ('admin', 'scout', 'player')"`
	EmailVerified bool       `json:"email_verified" gorm:"default:false"`
	IsActive      bool       `json:"is_active" gorm:"default:true"`
	LastLoginAt   *time.Time `json:"last_login_at,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

// RefreshToken stores JWT refresh tokens
type RefreshToken struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	TokenHash string    `json:"-" gorm:"not null"`
	ExpiresAt time.Time `json:"expires_at" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
}

// PasswordResetToken stores password reset tokens (secure, time-limited)
type PasswordResetToken struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID    uuid.UUID  `json:"user_id" gorm:"type:uuid;not null;index"`
	TokenHash string     `json:"-" gorm:"not null;uniqueIndex"`
	ExpiresAt time.Time  `json:"expires_at" gorm:"not null"`
	UsedAt    *time.Time `json:"used_at,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
}

// EmailVerificationToken stores email verification codes
type EmailVerificationToken struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID    uuid.UUID  `json:"user_id" gorm:"type:uuid;not null;index"`
	Code      string     `json:"-" gorm:"not null"` // 6-digit code
	ExpiresAt time.Time  `json:"expires_at" gorm:"not null"`
	UsedAt    *time.Time `json:"used_at,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
}

// GeneralContactRequest stores public contact form submissions (landing page inquiries)
type GeneralContactRequest struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name      string     `json:"name" gorm:"not null"`
	Email     string     `json:"email" gorm:"not null;index"`
	Subject   *string    `json:"subject,omitempty"`
	Message   string     `json:"message" gorm:"not null"`
	Type      string     `json:"type" gorm:"default:'general';index"` // general, partnership, media, support
	Status    string     `json:"status" gorm:"default:'new';index"`   // new, read, responded, archived
	IPAddress *string    `json:"-"`
	UserAgent *string    `json:"-"`
	ReadAt    *time.Time `json:"read_at,omitempty"`
	ReadBy    *uuid.UUID `json:"read_by,omitempty" gorm:"type:uuid"`
	CreatedAt time.Time  `json:"created_at"`
}

// Scout represents scout profiles
type Scout struct {
	ID               uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID           uuid.UUID `json:"user_id" gorm:"type:uuid;uniqueIndex;not null"`
	OrganizationName *string   `json:"organization_name,omitempty"`
	OrganizationType *string   `json:"organization_type,omitempty"`
	Country          *string   `json:"country,omitempty"`
	IsVerified       bool      `json:"is_verified" gorm:"default:false"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`

	User *User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// Tournament represents events/tournaments
type Tournament struct {
	ID          uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string     `json:"name" gorm:"not null"`
	Year        int        `json:"year" gorm:"not null;index"`
	Location    *string    `json:"location,omitempty"`
	Country     *string    `json:"country,omitempty"`
	StartDate   *time.Time `json:"start_date,omitempty"`
	EndDate     *time.Time `json:"end_date,omitempty"`
	Description *string    `json:"description,omitempty"`
	Status      string     `json:"status" gorm:"default:'upcoming'"`
	CreatedBy   *uuid.UUID `json:"created_by,omitempty" gorm:"type:uuid"`
	CreatedAt   time.Time  `json:"created_at"`
}

// Player represents player profiles created by admin
type Player struct {
	ID                  uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID              *uuid.UUID `json:"user_id,omitempty" gorm:"type:uuid;index"`
	FirstName           string     `json:"first_name" gorm:"not null"`
	LastName            string     `json:"last_name" gorm:"not null"`
	DateOfBirth         time.Time  `json:"date_of_birth" gorm:"not null"`
	Position            string     `json:"position" gorm:"not null;index"`
	PreferredFoot       *string    `json:"preferred_foot,omitempty"`
	HeightCm            *int       `json:"height_cm,omitempty"`
	WeightKg            *int       `json:"weight_kg,omitempty"`
	Country             string     `json:"country" gorm:"not null;index"`
	State               *string    `json:"state,omitempty"`
	City                *string    `json:"city,omitempty"`
	SchoolName          *string    `json:"school_name,omitempty"`
	VerificationStatus  string     `json:"verification_status" gorm:"default:'pending';index"`
	VerificationDocType *string    `json:"-"`
	VerificationDocURL  *string    `json:"-"`
	VerifiedAt          *time.Time `json:"-"`
	VerifiedBy          *uuid.UUID `json:"-" gorm:"type:uuid"`
	ProfilePhotoURL     *string    `json:"profile_photo_url,omitempty"`
	ThumbnailURL        *string    `json:"thumbnail_url,omitempty"`
	TournamentID        *uuid.UUID `json:"tournament_id,omitempty" gorm:"type:uuid;index"`
	TournamentYear      *int       `json:"tournament_year,omitempty"`
	CreatedBy           uuid.UUID  `json:"-" gorm:"type:uuid;not null"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	DeletedAt           *time.Time `json:"-" gorm:"index"`

	Tournament *Tournament `json:"tournament,omitempty" gorm:"foreignKey:TournamentID"`
	Videos     []Video     `json:"videos,omitempty" gorm:"many2many:player_videos;"`
}

// Subscription represents scout subscription management
type Subscription struct {
	ID                   uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID               uuid.UUID  `json:"user_id" gorm:"type:uuid;uniqueIndex;not null"`
	Tier                 string     `json:"tier" gorm:"not null;default:'free';index"`
	Status               string     `json:"status" gorm:"default:'active';index"`
	StripeCustomerID     *string    `json:"-"`
	StripeSubscriptionID *string    `json:"-"`
	CurrentPeriodStart   *time.Time `json:"current_period_start,omitempty"`
	CurrentPeriodEnd     *time.Time `json:"current_period_end,omitempty"`
	CancelAtPeriodEnd    bool       `json:"cancel_at_period_end" gorm:"default:false"`
	CancelledAt          *time.Time `json:"cancelled_at,omitempty"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
}

// Video represents video content (uploaded by admin only)
type Video struct {
	ID              uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	VideoType       string     `json:"video_type" gorm:"not null;index"`
	Title           string     `json:"title" gorm:"not null"`
	Description     *string    `json:"description,omitempty"`
	BlobURL         string     `json:"-" gorm:"not null"`
	ThumbnailURL    *string    `json:"thumbnail_url,omitempty"`
	DurationSeconds *int       `json:"duration_seconds,omitempty"`
	FileSizeBytes   *int64     `json:"file_size_bytes,omitempty"`
	TournamentID    *uuid.UUID `json:"tournament_id,omitempty" gorm:"type:uuid;index"`
	UploadedBy      uuid.UUID  `json:"-" gorm:"type:uuid;not null"`
	CreatedAt       time.Time  `json:"created_at"`

	Tournament *Tournament `json:"tournament,omitempty" gorm:"foreignKey:TournamentID"`
	Players    []Player    `json:"players,omitempty" gorm:"many2many:player_videos;"`
}

// PlayerVideo is many-to-many relationship between players and videos
type PlayerVideo struct {
	PlayerID  uuid.UUID `json:"player_id" gorm:"type:uuid;primaryKey"`
	VideoID   uuid.UUID `json:"video_id" gorm:"type:uuid;primaryKey"`
	IsPrimary bool      `json:"is_primary" gorm:"default:false"`
}

// ContactRequest represents scout inquiries about players (Pro+ tier)
type ContactRequest struct {
	ID         uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID     uuid.UUID  `json:"user_id" gorm:"type:uuid;not null;index"`
	PlayerID   uuid.UUID  `json:"player_id" gorm:"type:uuid;not null;index"`
	Message    string     `json:"message" gorm:"not null"`
	Status     string     `json:"status" gorm:"default:'pending';index"`
	HandledBy  *uuid.UUID `json:"-" gorm:"type:uuid"`
	HandledAt  *time.Time `json:"-"`
	AdminNotes *string    `json:"-"`
	CreatedAt  time.Time  `json:"created_at"`

	Player *Player `json:"player,omitempty" gorm:"foreignKey:PlayerID"`
}

// SavedPlayer represents scout's saved/favorited players (Scout+ tier)
type SavedPlayer struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index;uniqueIndex:idx_saved_unique"`
	PlayerID  uuid.UUID `json:"player_id" gorm:"type:uuid;not null;index;uniqueIndex:idx_saved_unique"`
	Notes     *string   `json:"notes,omitempty"`
	CreatedAt time.Time `json:"created_at"`

	Player *Player `json:"player,omitempty" gorm:"foreignKey:PlayerID"`
}

// VideoView tracks video analytics
type VideoView struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	VideoID   uuid.UUID  `json:"video_id" gorm:"type:uuid;not null;index"`
	ViewerID  *uuid.UUID `json:"viewer_id,omitempty" gorm:"type:uuid;index"`
	PlayerID  *uuid.UUID `json:"player_id,omitempty" gorm:"type:uuid"`
	CreatedAt time.Time  `json:"created_at" gorm:"index"`
}

// AuditLog provides audit trail for compliance
type AuditLog struct {
	ID           uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID       *uuid.UUID `json:"user_id,omitempty" gorm:"type:uuid;index"`
	Action       string     `json:"action" gorm:"not null;index"`
	ResourceType string     `json:"resource_type" gorm:"not null"`
	ResourceID   *uuid.UUID `json:"resource_id,omitempty" gorm:"type:uuid"`
	Details      *string    `json:"details,omitempty" gorm:"type:jsonb"`
	IPAddress    *string    `json:"ip_address,omitempty"`
	CreatedAt    time.Time  `json:"created_at" gorm:"index"`
}

// TableName overrides
func (PlayerVideo) TableName() string {
	return "player_videos"
}

// Helper methods

// GetAge calculates age from date of birth
func (p *Player) GetAge() int {
	now := time.Now()
	age := now.Year() - p.DateOfBirth.Year()
	if now.YearDay() < p.DateOfBirth.YearDay() {
		age--
	}
	return age
}

// GetLastNameInit returns abbreviated last name for privacy
func (p *Player) GetLastNameInit() string {
	if len(p.LastName) > 0 {
		return string(p.LastName[0]) + "."
	}
	return ""
}

// IsVerified returns true if player is verified
func (p *Player) IsVerified() bool {
	return p.VerificationStatus == "verified"
}

// CanAccessFullMatch checks if subscription tier can access full matches
func (s *Subscription) CanAccessFullMatch() bool {
	return s.Status == "active" && (s.Tier == "scout" || s.Tier == "pro" || s.Tier == "club")
}

// CanContactPlayers checks if subscription tier can contact players
func (s *Subscription) CanContactPlayers() bool {
	return s.Status == "active" && (s.Tier == "pro" || s.Tier == "club")
}

// CanSavePlayers checks if subscription tier can save players
func (s *Subscription) CanSavePlayers() bool {
	return s.Status == "active" && s.Tier != "free"
}
