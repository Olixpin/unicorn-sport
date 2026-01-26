package subscriptions

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stripe/stripe-go/v76"
	portalsession "github.com/stripe/stripe-go/v76/billingportal/session"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/stripe/stripe-go/v76/customer"
	"github.com/stripe/stripe-go/v76/webhook"
	"gorm.io/gorm"

	"github.com/unicorn-sport/backend/internal/domain"
)

// SubscriptionModule handles subscription operations
type SubscriptionModule struct {
	db            *gorm.DB
	stripeKey     string
	webhookSecret string
	priceIDs      map[string]string
	successURL    string
	cancelURL     string
}

// NewSubscriptionModule creates a new subscription module
func NewSubscriptionModule(db *gorm.DB, stripeKey, webhookSecret string, priceIDs map[string]string, successURL, cancelURL string) *SubscriptionModule {
	if stripeKey != "" {
		stripe.Key = stripeKey
	}
	return &SubscriptionModule{
		db:            db,
		stripeKey:     stripeKey,
		webhookSecret: webhookSecret,
		priceIDs:      priceIDs,
		successURL:    successURL,
		cancelURL:     cancelURL,
	}
}

// TierInfo describes a subscription tier
type TierInfo struct {
	Tier        string   `json:"tier"`
	Name        string   `json:"name"`
	Price       int      `json:"price"`
	Currency    string   `json:"currency"`
	Interval    string   `json:"interval"`
	Features    []string `json:"features"`
	Recommended bool     `json:"recommended,omitempty"`
}

// GetTiers returns available subscription tiers
func (m *SubscriptionModule) GetTiers(c *gin.Context) {
	tiers := []TierInfo{
		{
			Tier:     "free",
			Name:     "Free",
			Price:    0,
			Currency: "usd",
			Interval: "month",
			Features: []string{
				"View player highlights",
				"Browse player profiles",
				"Basic search",
			},
		},
		{
			Tier:        "scout",
			Name:        "Scout",
			Price:       99,
			Currency:    "usd",
			Interval:    "month",
			Recommended: true,
			Features: []string{
				"All Free features",
				"Full match access",
				"Save favorite players",
				"Advanced search filters",
			},
		},
		{
			Tier:     "pro",
			Name:     "Pro",
			Price:    299,
			Currency: "usd",
			Interval: "month",
			Features: []string{
				"All Scout features",
				"Contact players directly",
				"Player comparison tools",
				"Export player data",
				"Priority support",
			},
		},
		{
			Tier:     "club",
			Name:     "Club / Enterprise",
			Price:    999,
			Currency: "usd",
			Interval: "month",
			Features: []string{
				"All Pro features",
				"API access",
				"Multiple team members",
				"Custom integrations",
				"Dedicated account manager",
			},
		},
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"tiers": tiers}})
}

// GetCurrentSubscription returns the user's current subscription
func (m *SubscriptionModule) GetCurrentSubscription(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var sub domain.Subscription
	if err := m.db.Where("user_id = ?", userID).First(&sub).Error; err != nil {
		// Return free tier if no subscription exists
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"tier":   "free",
				"status": "active",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"id":                   sub.ID,
			"tier":                 sub.Tier,
			"status":               sub.Status,
			"current_period_start": sub.CurrentPeriodStart,
			"current_period_end":   sub.CurrentPeriodEnd,
			"cancel_at_period_end": sub.CancelAtPeriodEnd,
		},
	})
}

// CreateCheckoutRequest represents a checkout session request
type CreateCheckoutRequest struct {
	Tier string `json:"tier" binding:"required,oneof=scout pro club"`
}

// CreateCheckout creates a Stripe checkout session
func (m *SubscriptionModule) CreateCheckout(c *gin.Context) {
	var req CreateCheckoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()}})
		return
	}

	userID, _ := c.Get("user_id")
	userEmail, _ := c.Get("user_email")

	// Check if Stripe is configured
	if m.stripeKey == "" {
		c.JSON(http.StatusServiceUnavailable, gin.H{"success": false, "error": gin.H{"code": "STRIPE_NOT_CONFIGURED", "message": "Payment processing is not configured"}})
		return
	}

	// Get or create Stripe customer
	var sub domain.Subscription
	m.db.Where("user_id = ?", userID).First(&sub)

	var customerID string
	if sub.StripeCustomerID != nil {
		customerID = *sub.StripeCustomerID
	} else {
		// Create new Stripe customer
		params := &stripe.CustomerParams{
			Email: stripe.String(userEmail.(string)),
			Metadata: map[string]string{
				"user_id": userID.(uuid.UUID).String(),
			},
		}
		cust, err := customer.New(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "CUSTOMER_CREATE_FAILED", "message": "Failed to create customer"}})
			return
		}
		customerID = cust.ID
	}

	// Get price ID for tier
	priceID, ok := m.priceIDs[req.Tier]
	if !ok || priceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "INVALID_TIER", "message": "Invalid subscription tier"}})
		return
	}

	// Create checkout session
	params := &stripe.CheckoutSessionParams{
		Customer: stripe.String(customerID),
		Mode:     stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(priceID),
				Quantity: stripe.Int64(1),
			},
		},
		SuccessURL: stripe.String(m.successURL + "?session_id={CHECKOUT_SESSION_ID}"),
		CancelURL:  stripe.String(m.cancelURL),
		Metadata: map[string]string{
			"user_id": userID.(uuid.UUID).String(),
			"tier":    req.Tier,
		},
	}

	sess, err := session.New(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "CHECKOUT_FAILED", "message": "Failed to create checkout session"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"checkout_url": sess.URL,
			"session_id":   sess.ID,
		},
	})
}

// CreatePortalSession creates a Stripe billing portal session
func (m *SubscriptionModule) CreatePortalSession(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var sub domain.Subscription
	if err := m.db.Where("user_id = ?", userID).First(&sub).Error; err != nil || sub.StripeCustomerID == nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": gin.H{"code": "NO_SUBSCRIPTION", "message": "No active subscription found"}})
		return
	}

	params := &stripe.BillingPortalSessionParams{
		Customer:  sub.StripeCustomerID,
		ReturnURL: stripe.String(m.successURL),
	}

	sess, err := portalsession.New(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "PORTAL_FAILED", "message": "Failed to create portal session"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"portal_url": sess.URL,
		},
	})
}

// CancelSubscription cancels the user's subscription at period end
func (m *SubscriptionModule) CancelSubscription(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var sub domain.Subscription
	if err := m.db.Where("user_id = ?", userID).First(&sub).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": gin.H{"code": "NOT_FOUND", "message": "No subscription found"}})
		return
	}

	now := time.Now()
	updates := map[string]interface{}{
		"cancel_at_period_end": true,
		"cancelled_at":         now,
		"updated_at":           now,
	}

	if err := m.db.Model(&sub).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": gin.H{"code": "CANCEL_FAILED", "message": "Failed to cancel subscription"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Subscription will be cancelled at the end of the billing period",
		"data": gin.H{
			"cancel_at": sub.CurrentPeriodEnd,
		},
	})
}

// HandleWebhook handles Stripe webhook events
func (m *SubscriptionModule) HandleWebhook(c *gin.Context) {
	payload, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	sig := c.GetHeader("Stripe-Signature")
	event, err := webhook.ConstructEvent(payload, sig, m.webhookSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid signature"})
		return
	}

	switch event.Type {
	case "checkout.session.completed":
		m.handleCheckoutCompleted(event.Data.Raw)

	case "customer.subscription.updated":
		m.handleSubscriptionUpdated(event.Data.Raw)

	case "customer.subscription.deleted":
		m.handleSubscriptionDeleted(event.Data.Raw)

	case "invoice.payment_failed":
		m.handlePaymentFailed(event.Data.Raw)
	}

	c.JSON(http.StatusOK, gin.H{"received": true})
}

func (m *SubscriptionModule) handleCheckoutCompleted(data []byte) {
	// Parse session from raw JSON
	// In production, you'd properly unmarshal the Stripe event data
	// For now, this is a placeholder for the webhook handling logic
}

func (m *SubscriptionModule) handleSubscriptionUpdated(data []byte) {
	// Update subscription status in database
}

func (m *SubscriptionModule) handleSubscriptionDeleted(data []byte) {
	// Mark subscription as cancelled
}

func (m *SubscriptionModule) handlePaymentFailed(data []byte) {
	// Update subscription status to past_due
}

// --- Subscription Middleware ---

// RequireSubscription middleware checks for minimum subscription tier
func RequireSubscription(db *gorm.DB, minTier string) gin.HandlerFunc {
	tierLevels := map[string]int{
		"free":  0,
		"scout": 1,
		"pro":   2,
		"club":  3,
	}

	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": gin.H{"code": "AUTH_REQUIRED", "message": "Authentication required"}})
			c.Abort()
			return
		}

		var sub domain.Subscription
		if err := db.Where("user_id = ?", userID).First(&sub).Error; err != nil {
			// No subscription = free tier
			if tierLevels[minTier] > 0 {
				c.JSON(http.StatusForbidden, gin.H{
					"success": false,
					"error": gin.H{
						"code":          "SUBSCRIPTION_REQUIRED",
						"message":       "This feature requires a " + minTier + " subscription or higher",
						"required_tier": minTier,
					},
				})
				c.Abort()
				return
			}
			c.Set("subscription_tier", "free")
			c.Next()
			return
		}

		if sub.Status != "active" {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"error": gin.H{
					"code":    "SUBSCRIPTION_INACTIVE",
					"message": "Your subscription is not active",
				},
			})
			c.Abort()
			return
		}

		userLevel := tierLevels[sub.Tier]
		requiredLevel := tierLevels[minTier]

		if userLevel < requiredLevel {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"error": gin.H{
					"code":          "UPGRADE_REQUIRED",
					"message":       "This feature requires a " + minTier + " subscription or higher",
					"current_tier":  sub.Tier,
					"required_tier": minTier,
				},
			})
			c.Abort()
			return
		}

		c.Set("subscription_tier", sub.Tier)
		c.Next()
	}
}
