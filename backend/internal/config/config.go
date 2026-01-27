package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/unicorn-sport/backend/internal/domain"
)

// Config holds all configuration for the application
type Config struct {
	Environment string
	Port        string
	Database    DatabaseConfig
	JWT         JWTConfig
	AWS         AWSConfig
	Stripe      StripeConfig
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	Secret          string
	AccessTokenTTL  int // in minutes
	RefreshTokenTTL int // in days
}

// AWSConfig holds AWS configuration
type AWSConfig struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string
	S3Bucket        string
	CloudFrontURL   string
}

// StripeConfig holds Stripe configuration
type StripeConfig struct {
	SecretKey     string
	WebhookSecret string
	PriceIDs      map[string]string // tier -> price ID
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if it exists (for local development)
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	config := &Config{
		Environment: getEnv("ENVIRONMENT", "development"),
		Port:        getEnv("PORT", "8080"),
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "unicorn_sport"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret:          getEnv("JWT_SECRET", "your-secret-key"),
			AccessTokenTTL:  getEnvAsInt("JWT_ACCESS_TTL_MINUTES", 15),
			RefreshTokenTTL: getEnvAsInt("JWT_REFRESH_TTL_DAYS", 7),
		},
		AWS: AWSConfig{
			Region:          getEnv("AWS_REGION", "us-east-1"),
			AccessKeyID:     getEnv("AWS_ACCESS_KEY_ID", ""),
			SecretAccessKey: getEnv("AWS_SECRET_ACCESS_KEY", ""),
			S3Bucket:        getEnv("AWS_S3_BUCKET", "unicorn-sport-media"),
			CloudFrontURL:   getEnv("AWS_CLOUDFRONT_URL", ""),
		},
		Stripe: StripeConfig{
			SecretKey:     getEnv("STRIPE_SECRET_KEY", ""),
			WebhookSecret: getEnv("STRIPE_WEBHOOK_SECRET", ""),
			PriceIDs: map[string]string{
				"scout":      getEnv("STRIPE_PRICE_SCOUT", ""),
				"pro":        getEnv("STRIPE_PRICE_PRO", ""),
				"enterprise": getEnv("STRIPE_PRICE_ENTERPRISE", ""),
			},
		},
	}

	return config, nil
}

// InitDB initializes the database connection
func (c *Config) InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host, c.Database.Port, c.Database.User, c.Database.Password, c.Database.DBName, c.Database.SSLMode)

	var gormLogger logger.Interface
	if c.Environment == "development" {
		gormLogger = logger.Default.LogMode(logger.Info)
	} else {
		gormLogger = logger.Default.LogMode(logger.Error)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto-migrate the schema
	if err := db.AutoMigrate(
		&domain.User{},
		&domain.RefreshToken{},
		&domain.PasswordResetToken{},
		&domain.EmailVerificationToken{},
		&domain.GeneralContactRequest{},
		&domain.Scout{},
		&domain.Tournament{},
		&domain.Player{},
		&domain.Subscription{},
		&domain.Video{},
		&domain.PlayerVideo{},
		&domain.Match{},
		&domain.MatchPlayer{},
		&domain.UploadSession{},
		&domain.ContactRequest{},
		&domain.SavedPlayer{},
		&domain.VideoView{},
		&domain.AuditLog{},
		&domain.Academy{},
		// New video architecture models
		&domain.MatchVideo{},
		&domain.PlayerHighlight{},
		&domain.MatchPurchase{},
		&domain.HighlightView{},
		&domain.MatchVideoView{},
	); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return db, nil
}

// Helper functions
// GetEnv returns environment variable value or default
func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnv(key, defaultValue string) string {
	return GetEnv(key, defaultValue)
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
