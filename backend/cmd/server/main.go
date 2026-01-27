package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/unicorn-sport/backend/internal/config"
	"github.com/unicorn-sport/backend/internal/middleware"
	"github.com/unicorn-sport/backend/internal/modules/admin"
	"github.com/unicorn-sport/backend/internal/modules/auth"
	"github.com/unicorn-sport/backend/internal/modules/contact"
	"github.com/unicorn-sport/backend/internal/modules/highlights"
	"github.com/unicorn-sport/backend/internal/modules/matches"
	"github.com/unicorn-sport/backend/internal/modules/media"
	"github.com/unicorn-sport/backend/internal/modules/profiles"
	"github.com/unicorn-sport/backend/internal/modules/search"
	"github.com/unicorn-sport/backend/internal/modules/subscriptions"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	db, err := cfg.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize modules
	authModule := auth.NewAuthModule(db, cfg.JWT.Secret, cfg.JWT.AccessTokenTTL, cfg.JWT.RefreshTokenTTL)
	adminModule := admin.NewAdminModule(db)
	mediaModule := media.NewMediaModule(db, cfg.AWS.Region, cfg.AWS.AccessKeyID, cfg.AWS.SecretAccessKey, cfg.AWS.S3Bucket, cfg.AWS.CloudFrontURL)
	profilesModule := profiles.NewProfilesModule(db)
	searchModule := search.NewSearchModule(db)
	contactModule := contact.NewContactModule(db)

	// Initialize S3 client for matches/highlights modules
	s3Client := mediaModule.GetS3Client()
	matchesModule := matches.NewModule(db, s3Client, cfg.AWS.S3Bucket, cfg.AWS.CloudFrontURL)
	highlightsModule := highlights.NewModule(db, s3Client, cfg.AWS.S3Bucket, cfg.AWS.CloudFrontURL)

	// Subscription URLs
	successURL := os.Getenv("STRIPE_SUCCESS_URL")
	if successURL == "" {
		successURL = "http://localhost:3000/subscription/success"
	}
	cancelURL := os.Getenv("STRIPE_CANCEL_URL")
	if cancelURL == "" {
		cancelURL = "http://localhost:3000/subscription/cancel"
	}
	subscriptionsModule := subscriptions.NewSubscriptionModule(db, cfg.Stripe.SecretKey, cfg.Stripe.WebhookSecret, cfg.Stripe.PriceIDs, successURL, cancelURL)

	// Setup router
	r := setupRouter(cfg, authModule, adminModule, mediaModule, profilesModule, searchModule, subscriptionsModule, contactModule, matchesModule, highlightsModule)

	// Start server
	log.Printf("🚀 Unicorn Sport API starting on port %s", cfg.Port)
	log.Printf("📍 Environment: %s", cfg.Environment)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupRouter(
	cfg *config.Config,
	authModule *auth.AuthModule,
	adminModule *admin.AdminModule,
	mediaModule *media.MediaModule,
	profilesModule *profiles.ProfilesModule,
	searchModule *search.SearchModule,
	subscriptionsModule *subscriptions.SubscriptionModule,
	contactModule *contact.ContactModule,
	matchesModule *matches.Module,
	highlightsModule *highlights.Module,
) *gin.Engine {
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "version": "1.0.0"})
	})

	// Serve OpenAPI spec
	r.GET("/openapi.yaml", func(c *gin.Context) {
		c.File("docs/openapi.yaml")
	})

	// Scalar API Documentation
	r.GET("/docs", func(c *gin.Context) {
		html := `<!DOCTYPE html>
<html>
<head>
    <title>Unicorn Sport API Documentation</title>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
</head>
<body>
    <script id="api-reference" data-url="/openapi.yaml"></script>
    <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
</body>
</html>`
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
	})

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// ==================
		// AUTH ROUTES
		// ==================
		authRoutes := v1.Group("/auth")
		{
			authRoutes.POST("/register", authModule.Register)
			authRoutes.POST("/login", authModule.Login)
			authRoutes.POST("/refresh", authModule.Refresh)
			authRoutes.POST("/forgot-password", authModule.ForgotPassword)
			authRoutes.POST("/reset-password", authModule.ResetPassword)

			// Protected auth routes
			authProtected := authRoutes.Group("")
			authProtected.Use(middleware.JWTMiddleware(cfg.JWT.Secret))
			{
				authProtected.POST("/logout", authModule.Logout)
				authProtected.GET("/me", authModule.GetMe)
				authProtected.POST("/change-password", authModule.ChangePassword)
				authProtected.POST("/verify-email", authModule.VerifyEmail)
				authProtected.POST("/send-verification", authModule.SendVerificationEmail)
			}
		}

		// ==================
		// PUBLIC ROUTES
		// ==================
		// Players - public listing
		v1.GET("/players", profilesModule.ListPlayers)
		v1.GET("/players/featured", profilesModule.GetFeaturedPlayers)
		v1.GET("/players/:id", optionalAuth(cfg.JWT.Secret, profilesModule.GetPlayer))

		// Player highlights (FREE - public)
		v1.GET("/players/:id/highlights", highlightsModule.GetPlayerHighlightsPublic)
		v1.GET("/players/:id/tournaments", highlightsModule.GetPlayerTournamentAppearances)
		v1.GET("/highlights/:id", highlightsModule.GetHighlight)
		v1.GET("/highlight-types", highlightsModule.GetHighlightTypes)

		// Videos - public highlights
		v1.GET("/videos/highlights", mediaModule.ListHighlights)

		// Search
		v1.GET("/search", searchModule.SearchPlayers) // Alias for /search/players
		v1.GET("/search/players", searchModule.SearchPlayers)
		v1.GET("/search/filters", searchModule.GetFilterOptions)
		v1.GET("/stats", searchModule.GetStats)

		// Public contact form (landing page inquiries)
		v1.POST("/contact", contactModule.SubmitContact)

		// Subscription tiers (public info)
		v1.GET("/subscriptions/tiers", subscriptionsModule.GetTiers)

		// Stripe webhook (no auth)
		v1.POST("/webhooks/stripe", subscriptionsModule.HandleWebhook)

		// ==================
		// AUTHENTICATED ROUTES
		// ==================
		protected := v1.Group("")
		protected.Use(middleware.JWTMiddleware(cfg.JWT.Secret))
		{
			// Video streaming (checks subscription for full matches)
			protected.GET("/videos/:id/stream", mediaModule.GetVideoStreamURL)
			protected.GET("/videos/full-matches", mediaModule.ListFullMatches)

			// Subscription management
			protected.GET("/subscriptions/me", subscriptionsModule.GetCurrentSubscription)
			protected.GET("/subscriptions/current", subscriptionsModule.GetCurrentSubscription) // alias
			protected.POST("/subscriptions/checkout", subscriptionsModule.CreateCheckout)
			protected.POST("/subscriptions/portal", subscriptionsModule.CreatePortalSession)
			protected.POST("/subscriptions/cancel", subscriptionsModule.CancelSubscription)

			// ==================
			// SCOUT FEATURES (Scout+ tier) - Saved players
			// ==================
			scout := protected.Group("")
			{
				scout.GET("/saved-players", profilesModule.GetSavedPlayers)
				scout.POST("/players/:id/save", profilesModule.SavePlayer)
				scout.DELETE("/players/:id/save", profilesModule.UnsavePlayer)
			}

			// ==================
			// PRO FEATURES (Pro+ tier) - Contact players
			// ==================
			pro := protected.Group("")
			{
				pro.POST("/players/:id/contact", profilesModule.CreateContactRequest)
				pro.GET("/contact-requests", profilesModule.GetMyContactRequests)
			}

			// ==================
			// ADMIN ROUTES
			// ==================
			adminRoutes := protected.Group("/admin")
			adminRoutes.Use(middleware.AdminMiddleware())
			{
				// Dashboard stats
				adminRoutes.GET("/stats", adminModule.GetStats)

				// Academy management
				adminRoutes.GET("/academies", adminModule.ListAcademies)
				adminRoutes.POST("/academies", adminModule.CreateAcademy)
				adminRoutes.GET("/academies/:id", adminModule.GetAcademy)
				adminRoutes.PUT("/academies/:id", adminModule.UpdateAcademy)
				adminRoutes.DELETE("/academies/:id", adminModule.DeleteAcademy)

				// Contact request management
				adminRoutes.GET("/contact-requests", adminModule.ListContactRequests)
				adminRoutes.PUT("/contact-requests/:id/approve", adminModule.ApproveContactRequest)
				adminRoutes.PUT("/contact-requests/:id/reject", adminModule.RejectContactRequest)

				// Player management
				adminRoutes.GET("/players", adminModule.ListPlayers)
				adminRoutes.POST("/players", adminModule.CreatePlayer)
				adminRoutes.GET("/players/:id", adminModule.GetPlayer)
				adminRoutes.PUT("/players/:id", adminModule.UpdatePlayer)
				adminRoutes.DELETE("/players/:id", adminModule.DeletePlayer)

				// Tournament management
				adminRoutes.GET("/events", adminModule.ListTournaments)
				adminRoutes.POST("/events", adminModule.CreateTournament)
				adminRoutes.PUT("/events/:id", adminModule.UpdateTournament)

				// Video management
				adminRoutes.GET("/videos", mediaModule.ListVideos)
				adminRoutes.GET("/videos/stats", mediaModule.GetVideoStats)
				adminRoutes.GET("/videos/:id", mediaModule.GetVideo)
				adminRoutes.POST("/videos", mediaModule.CreateVideo)
				adminRoutes.PUT("/videos/:id", mediaModule.UpdateVideo)
				adminRoutes.DELETE("/videos/:id", mediaModule.DeleteVideo)
				adminRoutes.POST("/videos/:id/players", mediaModule.LinkPlayerToVideo)
				adminRoutes.POST("/videos/:id/approve", mediaModule.ApproveVideo)
				adminRoutes.POST("/videos/:id/reject", mediaModule.RejectVideo)

				// Upload workflow
				adminRoutes.POST("/upload/init", mediaModule.InitUpload)
				adminRoutes.POST("/upload/multipart/init", mediaModule.InitMultipart)
				adminRoutes.POST("/upload/multipart/part-url", mediaModule.GetUploadPartURL)
				adminRoutes.POST("/upload/multipart/complete", mediaModule.CompleteMultipart)

				// Legacy upload endpoints (for backward compatibility)
				adminRoutes.POST("/videos/upload", mediaModule.GetUploadURL)
				adminRoutes.POST("/videos/:id/confirm", mediaModule.ConfirmUpload)

				// ==================
				// MATCH MANAGEMENT (New Video Architecture)
				// ==================
				// Tournament matches
				adminRoutes.GET("/tournaments/:tournamentId/matches", matchesModule.ListMatches)
				adminRoutes.POST("/tournaments/:tournamentId/matches", matchesModule.CreateMatch)

				// Match CRUD
				adminRoutes.GET("/matches/:id", matchesModule.GetMatch)
				adminRoutes.PUT("/matches/:id", matchesModule.UpdateMatch)
				adminRoutes.DELETE("/matches/:id", matchesModule.DeleteMatch)

				// Match players
				adminRoutes.POST("/matches/:id/players", matchesModule.AddPlayerToMatch)
				adminRoutes.PUT("/matches/:id/players/:playerId", matchesModule.UpdateMatchPlayer)
				adminRoutes.DELETE("/matches/:id/players/:playerId", matchesModule.RemovePlayerFromMatch)

				// Match video (full match - PAID content)
				adminRoutes.POST("/matches/:id/video/upload", matchesModule.InitMatchVideoUpload)
				adminRoutes.POST("/matches/:id/video", matchesModule.SaveMatchVideo)
				adminRoutes.DELETE("/matches/:id/video", matchesModule.DeleteMatchVideo)

				// Match video thumbnail
				adminRoutes.POST("/matches/:id/video/thumbnail/upload", matchesModule.InitThumbnailUpload)
				adminRoutes.PUT("/matches/:id/video/thumbnail", matchesModule.UpdateThumbnail)

				// Multipart upload for large videos
				adminRoutes.POST("/matches/upload/part-url", matchesModule.GetMultipartPartURL)
				adminRoutes.POST("/matches/upload/complete", matchesModule.CompleteMultipartUpload)

				// ==================
				// HIGHLIGHT MANAGEMENT (FREE content)
				// ==================
				// Highlight upload
				adminRoutes.POST("/highlights/upload", highlightsModule.InitHighlightUpload)
				adminRoutes.POST("/highlights/upload/init", highlightsModule.InitHighlightUpload) // alias
				adminRoutes.POST("/highlights", highlightsModule.CreateHighlight)
				adminRoutes.POST("/highlights/:id/thumbnail/upload", highlightsModule.InitThumbnailUpload)
				adminRoutes.PUT("/highlights/:id/thumbnail", highlightsModule.UpdateThumbnail)

				// Highlight CRUD
				adminRoutes.GET("/players/:id/highlights", highlightsModule.ListPlayerHighlights)
				adminRoutes.GET("/matches/:id/highlights", highlightsModule.ListMatchHighlights)
				adminRoutes.PUT("/highlights/:id", highlightsModule.UpdateHighlight)
				adminRoutes.DELETE("/highlights/:id", highlightsModule.DeleteHighlight)

				// User management
				adminRoutes.GET("/users", adminModule.ListUsers)
				adminRoutes.PUT("/users/:id", adminModule.UpdateUser)
			}
		}
	}

	return r
}

// optionalAuth middleware that sets user context if authenticated, but doesn't require it
func optionalAuth(jwtSecret string, handler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Try to authenticate, but don't fail if not authenticated
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" && len(authHeader) > 7 {
			tokenString := authHeader[7:] // Remove "Bearer "
			middleware.ParseAndSetClaims(c, tokenString, jwtSecret)
		}
		handler(c)
	}
}
