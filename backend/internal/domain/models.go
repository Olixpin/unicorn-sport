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
	AcademyID           *uuid.UUID `json:"academy_id,omitempty" gorm:"type:uuid;index"`
	FirstName           string     `json:"first_name" gorm:"not null"`
	LastName            string     `json:"last_name" gorm:"not null"`
	DateOfBirth         time.Time  `json:"date_of_birth" gorm:"not null"`
	Position            string     `json:"position" gorm:"not null;index"` // Primary/natural position
	SecondaryPositions  *string    `json:"secondary_positions,omitempty"`  // JSON array of other positions
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

	Tournament     *Tournament       `json:"tournament,omitempty" gorm:"foreignKey:TournamentID"`
	Academy        *Academy          `json:"academy,omitempty" gorm:"foreignKey:AcademyID"`
	Videos         []Video           `json:"videos,omitempty" gorm:"many2many:player_videos;"`
	Highlights     []PlayerHighlight `json:"highlights,omitempty" gorm:"foreignKey:PlayerID"`
	AcademyPlayers []AcademyPlayer   `json:"academy_memberships,omitempty" gorm:"foreignKey:PlayerID"`
}

// AcademyPlayer represents a player's membership in an academy with their squad role
// This allows the same player to have different roles at different academies
// and tracks their position as assigned by the academy (which may differ from natural position)
type AcademyPlayer struct {
	ID           uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	AcademyID    uuid.UUID  `json:"academy_id" gorm:"type:uuid;not null;index"`
	PlayerID     uuid.UUID  `json:"player_id" gorm:"type:uuid;not null;index"`
	SquadRole    *string    `json:"squad_role,omitempty"`                       // Position as used by this academy
	JerseyNumber *int       `json:"jersey_number,omitempty"`                    // Squad number at this academy
	SquadStatus  string     `json:"squad_status" gorm:"default:'active';index"` // active, injured, loan, departed
	JoinedAt     *time.Time `json:"joined_at,omitempty"`
	DepartedAt   *time.Time `json:"departed_at,omitempty"`
	Notes        *string    `json:"notes,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`

	Academy *Academy `json:"academy,omitempty" gorm:"foreignKey:AcademyID"`
	Player  *Player  `json:"player,omitempty" gorm:"foreignKey:PlayerID"`
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

// Video represents video content (highlights and clips uploaded by admin)
type Video struct {
	ID              uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	VideoType       string    `json:"video_type" gorm:"not null;index"` // highlight, full_match
	Title           string    `json:"title" gorm:"not null"`
	Description     *string   `json:"description,omitempty"`
	BlobURL         string    `json:"-" gorm:"not null"`
	ThumbnailURL    *string   `json:"thumbnail_url,omitempty"`
	DurationSeconds *int      `json:"duration_seconds,omitempty"`
	FileSizeBytes   *int64    `json:"file_size_bytes,omitempty"`

	// Review workflow
	Status      string     `json:"status" gorm:"default:'approved';index"` // pending, approved, rejected
	ReviewNotes *string    `json:"review_notes,omitempty"`
	ReviewedBy  *uuid.UUID `json:"-" gorm:"type:uuid"`
	ReviewedAt  *time.Time `json:"-"`

	// Associations
	TournamentID *uuid.UUID `json:"tournament_id,omitempty" gorm:"type:uuid;index"`
	MatchID      *uuid.UUID `json:"match_id,omitempty" gorm:"type:uuid;index"` // Source match for highlight
	UploadedBy   uuid.UUID  `json:"-" gorm:"type:uuid;not null"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`

	Tournament *Tournament `json:"tournament,omitempty" gorm:"foreignKey:TournamentID"`
	Match      *Match      `json:"match,omitempty" gorm:"foreignKey:MatchID"`
	Players    []Player    `json:"players,omitempty" gorm:"many2many:player_videos;"`
}

// Match represents a game/fixture in a tournament
type Match struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Title       string    `json:"title" gorm:"not null"`
	Description *string   `json:"description,omitempty"`
	MatchDate   time.Time `json:"match_date" gorm:"not null;index"`
	Location    *string   `json:"location,omitempty"`

	// Match details
	Stage       *string `json:"stage,omitempty"`        // group, round_of_16, quarterfinal, semifinal, final
	HomeTeam    *string `json:"home_team,omitempty"`    // Team A name
	AwayTeam    *string `json:"away_team,omitempty"`    // Team B name
	HomeScore   *int    `json:"home_score,omitempty"`   // Final score
	AwayScore   *int    `json:"away_score,omitempty"`   // Final score
	MatchNumber *int    `json:"match_number,omitempty"` // Match # in tournament

	// Status
	Status string `json:"status" gorm:"default:'scheduled';index"` // scheduled, in_progress, completed, cancelled

	// Associations
	TournamentID *uuid.UUID `json:"tournament_id,omitempty" gorm:"type:uuid;index"`
	AcademyID    *uuid.UUID `json:"academy_id,omitempty" gorm:"type:uuid;index"`

	// Metadata
	CreatedBy uuid.UUID `json:"-" gorm:"type:uuid;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relationships
	Tournament *Tournament `json:"tournament,omitempty" gorm:"foreignKey:TournamentID"`
	Academy    *Academy    `json:"academy,omitempty" gorm:"foreignKey:AcademyID"`
	Players    []Player    `json:"players,omitempty" gorm:"many2many:match_players;"`
	Video      *MatchVideo `json:"video,omitempty" gorm:"foreignKey:MatchID"`
}

// MatchVideo represents the full match video (PAID content)
type MatchVideo struct {
	ID      uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	MatchID uuid.UUID `json:"match_id" gorm:"type:uuid;uniqueIndex;not null"`

	// Video storage (S3)
	VideoURL        string  `json:"-" gorm:"not null"` // S3 key - hidden from JSON
	ThumbnailURL    *string `json:"thumbnail_url,omitempty"`
	DurationSeconds *int    `json:"duration_seconds,omitempty"`
	FileSizeBytes   *int64  `json:"file_size_bytes,omitempty"`

	// Processing status
	Status          string  `json:"status" gorm:"default:'processing';index"` // processing, ready, failed, archived
	ProcessingError *string `json:"-"`

	// Pricing for pay-per-view
	PriceCents int    `json:"price_cents" gorm:"default:999"` // Default $9.99
	Currency   string `json:"currency" gorm:"default:'USD'"`

	// Stats
	ViewCount     int `json:"view_count" gorm:"default:0"`
	PurchaseCount int `json:"purchase_count" gorm:"default:0"`

	// Metadata
	UploadedBy uuid.UUID `json:"-" gorm:"type:uuid;not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	// Relationships
	Match *Match `json:"match,omitempty" gorm:"foreignKey:MatchID"`
}

// PlayerHighlight represents a short highlight clip (FREE content)
type PlayerHighlight struct {
	ID       uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	PlayerID uuid.UUID  `json:"player_id" gorm:"type:uuid;not null;index"`
	MatchID  *uuid.UUID `json:"match_id,omitempty" gorm:"type:uuid;index"` // Optional - which match this is from

	// Highlight categorization
	HighlightType string `json:"highlight_type" gorm:"not null;index"` // goal, assist, dribbling, defending, etc.

	// Video storage (S3)
	VideoURL        string  `json:"-" gorm:"not null"` // S3 key - hidden from JSON
	ThumbnailURL    *string `json:"thumbnail_url,omitempty"`
	DurationSeconds *int    `json:"duration_seconds,omitempty"`
	FileSizeBytes   *int64  `json:"file_size_bytes,omitempty"`

	// Context
	Title            *string `json:"title,omitempty"`
	Description      *string `json:"description,omitempty"`
	TimestampInMatch *int    `json:"timestamp_in_match,omitempty"` // When in match (seconds)

	// Status
	Status string `json:"status" gorm:"default:'approved';index"` // pending, approved, rejected, archived

	// Stats
	ViewCount int `json:"view_count" gorm:"default:0"`

	// Metadata
	UploadedBy uuid.UUID `json:"-" gorm:"type:uuid;not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	// Relationships
	Player *Player `json:"player,omitempty" gorm:"foreignKey:PlayerID"`
	Match  *Match  `json:"match,omitempty" gorm:"foreignKey:MatchID"`
}

// MatchPurchase tracks pay-per-view purchases
type MatchPurchase struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID       uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index;uniqueIndex:idx_user_match_purchase"`
	MatchVideoID uuid.UUID `json:"match_video_id" gorm:"type:uuid;not null;index;uniqueIndex:idx_user_match_purchase"`

	// Payment info
	AmountCents           int     `json:"amount_cents" gorm:"not null"`
	Currency              string  `json:"currency" gorm:"default:'USD'"`
	StripePaymentIntentID *string `json:"-"`
	StripeChargeID        *string `json:"-"`

	// Status
	Status string `json:"status" gorm:"default:'completed';index"` // pending, completed, refunded, failed

	// Access tracking
	FirstViewedAt *time.Time `json:"first_viewed_at,omitempty"`
	LastViewedAt  *time.Time `json:"last_viewed_at,omitempty"`
	ViewCount     int        `json:"view_count" gorm:"default:0"`

	// Metadata
	CreatedAt time.Time `json:"created_at"`

	// Relationships
	User       *User       `json:"user,omitempty" gorm:"foreignKey:UserID"`
	MatchVideo *MatchVideo `json:"match_video,omitempty" gorm:"foreignKey:MatchVideoID"`
}

// HighlightView tracks views on highlights for analytics
type HighlightView struct {
	ID          uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	HighlightID uuid.UUID  `json:"highlight_id" gorm:"type:uuid;not null;index"`
	ViewerID    *uuid.UUID `json:"viewer_id,omitempty" gorm:"type:uuid;index"` // NULL for anonymous
	Source      *string    `json:"source,omitempty"`                           // player_profile, search, direct
	IPHash      *string    `json:"-"`                                          // For unique view counting
	CreatedAt   time.Time  `json:"created_at"`
}

// MatchVideoView tracks full match video views
type MatchVideoView struct {
	ID                   uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	MatchVideoID         uuid.UUID `json:"match_video_id" gorm:"type:uuid;not null;index"`
	ViewerID             uuid.UUID `json:"viewer_id" gorm:"type:uuid;not null;index"`
	WatchDurationSeconds *int      `json:"watch_duration_seconds,omitempty"`
	Completed            bool      `json:"completed" gorm:"default:false"`
	CreatedAt            time.Time `json:"created_at"`
}

// Highlight type constants
const (
	HighlightTypeGoal         = "goal"
	HighlightTypeAssist       = "assist"
	HighlightTypeDribbling    = "dribbling"
	HighlightTypeDefending    = "defending"
	HighlightTypeTackling     = "tackling"
	HighlightTypePassing      = "passing"
	HighlightTypeShooting     = "shooting"
	HighlightTypeHeading      = "heading"
	HighlightTypeSpeed        = "speed"
	HighlightTypeSave         = "save"
	HighlightTypeDistribution = "distribution"
	HighlightTypePositioning  = "positioning"
	HighlightTypeVision       = "vision"
	HighlightTypeOther        = "other"
)

// ValidHighlightTypes returns all valid highlight types
func ValidHighlightTypes() []string {
	return []string{
		HighlightTypeGoal, HighlightTypeAssist, HighlightTypeDribbling,
		HighlightTypeDefending, HighlightTypeTackling, HighlightTypePassing,
		HighlightTypeShooting, HighlightTypeHeading, HighlightTypeSpeed,
		HighlightTypeSave, HighlightTypeDistribution, HighlightTypePositioning,
		HighlightTypeVision, HighlightTypeOther,
	}
}

// IsValidHighlightType checks if a highlight type is valid
func IsValidHighlightType(t string) bool {
	for _, valid := range ValidHighlightTypes() {
		if t == valid {
			return true
		}
	}
	return false
}

// MatchPlayer links players to matches they participated in
type MatchPlayer struct {
	ID             uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	MatchID        uuid.UUID  `json:"match_id" gorm:"type:uuid;not null;uniqueIndex:idx_match_player_unique"`
	PlayerID       uuid.UUID  `json:"player_id" gorm:"type:uuid;not null;uniqueIndex:idx_match_player_unique"`
	PositionPlayed *string    `json:"position_played,omitempty"`
	MinutesPlayed  *int       `json:"minutes_played,omitempty"`
	Goals          int        `json:"goals" gorm:"default:0"`
	Assists        int        `json:"assists" gorm:"default:0"`
	Notes          *string    `json:"notes,omitempty"`
	IsStarter      bool       `json:"is_starter" gorm:"default:true"`
	JerseyNumber   *int       `json:"jersey_number,omitempty"`
	FormationX     *float64   `json:"formation_x,omitempty"`                    // 0-100 percentage position on pitch
	FormationY     *float64   `json:"formation_y,omitempty"`                    // 0-100 percentage position on pitch
	SubbedInFor    *uuid.UUID `json:"subbed_in_for,omitempty" gorm:"type:uuid"` // Player ID they replaced
	SubbedInAt     *int       `json:"subbed_in_at,omitempty"`                   // Minute subbed in
	SubbedOutAt    *int       `json:"subbed_out_at,omitempty"`                  // Minute subbed out
	CreatedAt      time.Time  `json:"created_at"`

	Match  *Match  `json:"match,omitempty" gorm:"foreignKey:MatchID"`
	Player *Player `json:"player,omitempty" gorm:"foreignKey:PlayerID"`
}

// UploadSession tracks multipart uploads to S3
type UploadSession struct {
	ID          uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UploadType  string     `json:"upload_type" gorm:"not null"`  // video, match, thumbnail, document
	ContentType string     `json:"content_type" gorm:"not null"` // MIME type
	FileName    string     `json:"file_name" gorm:"not null"`
	FileSize    int64      `json:"file_size" gorm:"not null"`
	S3UploadID  *string    `json:"-"`
	S3Key       string     `json:"-" gorm:"not null"`
	Status      string     `json:"status" gorm:"default:'pending';index"` // pending, uploading, completed, failed, expired
	PartsTotal  *int       `json:"parts_total,omitempty"`
	EntityType  *string    `json:"entity_type,omitempty"` // video, match
	EntityID    *uuid.UUID `json:"entity_id,omitempty" gorm:"type:uuid"`
	UploadedBy  uuid.UUID  `json:"-" gorm:"type:uuid;not null"`
	CreatedAt   time.Time  `json:"created_at"`
	ExpiresAt   time.Time  `json:"expires_at" gorm:"not null"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
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

func (MatchPlayer) TableName() string {
	return "match_players"
}

func (UploadSession) TableName() string {
	return "upload_sessions"
}

func (MatchVideo) TableName() string {
	return "match_videos"
}

func (PlayerHighlight) TableName() string {
	return "player_highlights"
}

func (MatchPurchase) TableName() string {
	return "match_purchases"
}

func (HighlightView) TableName() string {
	return "highlight_views"
}

func (MatchVideoView) TableName() string {
	return "match_video_views"
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

// Setting represents a platform configuration setting
type Setting struct {
	Key       string    `json:"key" gorm:"primaryKey"`
	Value     string    `json:"value" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Academy represents a football academy
type Academy struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string    `json:"name" gorm:"not null;index"`
	Description *string   `json:"description,omitempty"`
	Country     string    `json:"country" gorm:"not null;index"`
	State       *string   `json:"state,omitempty"`
	City        *string   `json:"city,omitempty"`
	Address     *string   `json:"address,omitempty"`
	Phone       *string   `json:"phone,omitempty"`
	Email       *string   `json:"email,omitempty"`
	Website     *string   `json:"website,omitempty"`
	LogoURL     *string   `json:"logo_url,omitempty"`
	FoundedYear *int      `json:"founded_year,omitempty"`
	IsVerified  bool      `json:"is_verified" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
