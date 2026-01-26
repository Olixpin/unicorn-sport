# Unicorn Sport - Subscription Service

## 🚨 Priority: P0 - CRITICAL BLOCKER

> **Without this, you have ZERO revenue.**
> **Scouts cannot pay, and highlights remain inaccessible.**
> **This must be built alongside the Admin Portal.**

---

## 🎯 Overview

The Subscription Service manages scout payments, access control, and billing. It's the **monetization engine** of Unicorn Sport.

### Core Responsibilities

1. **Subscription Management** - Plans, tiers, upgrades/downgrades
2. **Payment Processing** - Stripe integration for recurring billing
3. **Access Control** - Gate video content behind subscription check
4. **Billing** - Invoices, receipts, payment history
5. **Analytics** - MRR, churn, conversion tracking

---

## 💳 Subscription Tiers

### Pricing Strategy

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         SUBSCRIPTION PLANS                                   │
└─────────────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────────────┐
│  🆓 FREE: $0 (No account required)                                          │
│  ══════════════════════════════════                                          │
│  Target: Lead generation, attract scouts                                    │
│                                                                              │
│  Features:                                                                   │
│  ✓ Browse basic player profiles                                             │
│  ✓ Watch HIGHLIGHTS (FREE but can be DECEPTIVE!)                            │
│  ✓ View thumbnails                                                          │
│  ✓ Basic search by position/age/location                                    │
│  ✗ Cannot see full player profiles                                          │
│  ✗ Cannot watch full matches (the REAL performance)                         │
│  ✗ Cannot contact players                                                   │
│  ✗ No analytics                                                             │
│                                                                              │
│  ⚠️ Highlights show flashy moments but hide context!                       │
└─────────────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────────────┐
│  🥉 SCOUT: $99/month                                                         │
│  ═════════════════════                                                       │
│  Target: Individual scouts, small agencies                                  │
│                                                                              │
│  Features:                                                                   │
│  ✓ Everything FREE, plus:                                                   │
│  ✓ Watch FULL MATCH recordings (see REAL performance!)                      │
│  ✓ Full player profiles                                                     │
│  ✓ Basic analytics                                                          │
│  ✓ Save favorites                                                           │
│  ✗ Cannot contact players (Pro required)                                    │
│  ✗ No advanced analytics                                                    │
│  ✗ No API access                                                            │
│                                                                              │
│  💡 See how players REALLY perform over 90 minutes                          │
└─────────────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────────────┐
│  🥈 PRO: $299/month                                                          │
│  ══════════════════                                                          │
│  Target: Professional scouts, academy recruiters                            │
│                                                                              │
│  Features:                                                                   │
│  ✓ Everything in Scout                                                      │
│  ✓ CONTACT players ("Write to us")                                          │
│  ✓ Advanced analytics & performance stats                                   │
│  ✓ Save shortlists (up to 100 players)                                      │
│  ✓ Advanced search filters                                                  │
│  ✓ Email alerts for new talent matching preferences                         │
│  ✓ Export scouting reports                                                  │
│  ✗ No API access                                                            │                                                            │
└─────────────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────────────┐
│  🥇 ENTERPRISE: Custom pricing (from $999/month)                             │
│  ═══════════════════════════════════════════════                             │
│  Target: Clubs, large agencies, federations                                 │
│                                                                              │
│  Features:                                                                   │
│  ✓ Everything in Pro                                                        │
│  ✓ API access for integration                                               │
│  ✓ Unlimited player contacts                                                │
│  ✓ Unlimited shortlists                                                     │
│  ✓ Bulk data exports                                                        │
│  ✓ Priority support                                                         │
│  ✓ Custom integrations                                                      │
│  ✓ Multiple user seats                                                      │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Free Tier (No Subscription)

```
What FREE users can access:
├── PLAYER HIGHLIGHTS (flashy moments, dribbles, goals)
├── Thumbnail images
├── Search results
└── Limited player info

⚠️ BUT highlights can be DECEPTIVE:
├── Player looks amazing vs weak team
├── Cherry-picked best moments only
├── No context about consistency or real performance
```

### What PAID Subscribers Get

```
PAID subscribers unlock THE TRUTH:
├── Browse full PLAYER PROFILES
├── Watch FULL MATCHES (see real context & performance)
├── Access ANALYTICS (detailed performance stats)
├── CONTACT players ("Write to us if interested")
├── Shortlists & saved searches
└── API access (Enterprise tier)

💡 VALUE: "Don't get deceived by highlights - see the REAL performance"
```

> **Strategy**: Highlights are FREE to attract scouts. 
> Full matches + profiles + contact = PAID to verify if player is real deal.

---

## 🔄 User Flow

### Scout Subscription Journey

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     SCOUT SUBSCRIPTION JOURNEY                               │
└─────────────────────────────────────────────────────────────────────────────┘

Step 1: Discovery
┌─────────────────────────────────────────────────────────────────────────────┐
│  Scout finds Unicorn Sport via:                                              │
│  • YouTube sample match with link to platform                               │
│  • Google search "verified african football talent"                         │
│  • Industry referral                                                         │
│  • Social media / marketing                                                  │
│                                                                              │
│  Lands on platform, watches highlights (FREE)                                 │
│  Sees highlight thumbnails but cannot play videos                           │
└─────────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
Step 2: Paywall Hit
┌─────────────────────────────────────────────────────────────────────────────┐
│  Scout clicks "Watch Highlight"                                              │
│                                                                              │
│  ┌───────────────────────────────────────────────────────────────────────┐  │
│  │                                                                        │  │
│  │    🔒 Subscribe to Watch This Highlight                               │  │
│  │                                                                        │  │
│  │    This is a verified player highlight from Lagos Regional Showcase.  │  │
│  │    Subscribe to access all highlights and connect with talent.        │  │
│  │                                                                        │  │
│  │    [Start Free Trial - 7 Days]        [View Plans]                    │  │
│  │                                                                        │  │
│  └───────────────────────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
Step 3: Checkout
┌─────────────────────────────────────────────────────────────────────────────┐
│  Scout selects plan and enters payment (Stripe Checkout)                    │
│                                                                              │
│  • Credit/Debit card                                                        │
│  • Billing address                                                          │
│  • Tax calculation (if applicable)                                          │
│                                                                              │
│  Stripe handles:                                                             │
│  • PCI compliance                                                           │
│  • Recurring billing                                                        │
│  • Failed payment retries                                                   │
└─────────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
Step 4: Access Granted
┌─────────────────────────────────────────────────────────────────────────────┐
│  Subscription active:                                                        │
│                                                                              │
│  • All highlights now playable                                              │
│  • Contact buttons enabled (Pro+)                                           │
│  • Shortlist feature unlocked (Pro+)                                        │
│  • Monthly usage tracking begins                                            │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 📋 Data Model

### Database Schema

```sql
-- Subscription plans
CREATE TABLE plans (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50) NOT NULL,              -- 'basic', 'pro', 'enterprise'
    display_name VARCHAR(100) NOT NULL,     -- 'Basic', 'Pro', 'Enterprise'
    price_cents INTEGER NOT NULL,           -- 9900, 29900, etc.
    currency VARCHAR(3) DEFAULT 'USD',
    billing_period VARCHAR(20) DEFAULT 'monthly', -- 'monthly', 'yearly'
    stripe_price_id VARCHAR(100),           -- Stripe Price ID
    features JSONB,                         -- Feature flags
    highlight_limit INTEGER,                -- NULL = unlimited
    contact_limit INTEGER,                  -- NULL = unlimited
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT NOW()
);

-- User subscriptions
CREATE TABLE subscriptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    plan_id UUID NOT NULL REFERENCES plans(id),
    stripe_subscription_id VARCHAR(100),
    stripe_customer_id VARCHAR(100),
    status VARCHAR(20) NOT NULL,            -- 'active', 'past_due', 'cancelled', 'trialing'
    current_period_start TIMESTAMP,
    current_period_end TIMESTAMP,
    cancel_at_period_end BOOLEAN DEFAULT false,
    cancelled_at TIMESTAMP,
    trial_end TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Usage tracking
CREATE TABLE usage_records (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    subscription_id UUID REFERENCES subscriptions(id),
    user_id UUID NOT NULL REFERENCES users(id),
    usage_type VARCHAR(50) NOT NULL,        -- 'highlight_view', 'player_contact'
    resource_id UUID,                       -- video_id or profile_id
    period_start DATE NOT NULL,             -- Billing period
    count INTEGER DEFAULT 1,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Payment history
CREATE TABLE payments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    subscription_id UUID REFERENCES subscriptions(id),
    stripe_payment_intent_id VARCHAR(100),
    stripe_invoice_id VARCHAR(100),
    amount_cents INTEGER NOT NULL,
    currency VARCHAR(3) DEFAULT 'USD',
    status VARCHAR(20) NOT NULL,            -- 'succeeded', 'failed', 'pending'
    paid_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Indexes
CREATE INDEX idx_subscriptions_user ON subscriptions(user_id);
CREATE INDEX idx_subscriptions_status ON subscriptions(status);
CREATE INDEX idx_usage_subscription_period ON usage_records(subscription_id, period_start);
```

---

## 📡 API Endpoints

### Public Endpoints

```http
GET /api/v1/subscriptions/plans
```

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": "plan-basic",
      "name": "basic",
      "display_name": "Basic",
      "price": 99.00,
      "currency": "USD",
      "billing_period": "monthly",
      "features": {
        "highlight_views": 10,
        "player_contacts": 0,
        "shortlists": false,
        "api_access": false
      }
    },
    {
      "id": "plan-pro",
      "name": "pro",
      "display_name": "Pro",
      "price": 299.00,
      "currency": "USD",
      "billing_period": "monthly",
      "features": {
        "highlight_views": -1,
        "player_contacts": 20,
        "shortlists": true,
        "api_access": false
      }
    }
  ]
}
```

### Authenticated Endpoints

```http
POST /api/v1/subscriptions/checkout
Authorization: Bearer <token>
```

**Request:**
```json
{
  "plan_id": "plan-pro",
  "success_url": "https://unicornsport.africa/subscription/success",
  "cancel_url": "https://unicornsport.africa/subscription/cancelled"
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "checkout_url": "https://checkout.stripe.com/c/pay/cs_xxx",
    "session_id": "cs_xxx"
  }
}
```

---

```http
GET /api/v1/subscriptions/me
Authorization: Bearer <token>
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": "sub-uuid",
    "plan": {
      "name": "pro",
      "display_name": "Pro"
    },
    "status": "active",
    "current_period_end": "2026-02-25T00:00:00Z",
    "usage": {
      "highlight_views": 45,
      "highlight_limit": null,
      "player_contacts": 8,
      "contact_limit": 20
    }
  }
}
```

---

```http
GET /api/v1/access/check/:videoId
Authorization: Bearer <token>
```

**Response (Has Access):**
```json
{
  "success": true,
  "data": {
    "has_access": true,
    "stream_url": "https://storage.blob.core.windows.net/videos/xxx?sig=xxx&exp=xxx"
  }
}
```

**Response (No Access):**
```json
{
  "success": false,
  "error": {
    "code": "SUBSCRIPTION_REQUIRED",
    "message": "Subscribe to view this highlight",
    "upgrade_url": "/subscribe"
  }
}
```

---

## 🔗 Stripe Integration

### Webhook Events to Handle

| Event | Action |
|-------|--------|
| `checkout.session.completed` | Create subscription record |
| `customer.subscription.updated` | Update subscription status |
| `customer.subscription.deleted` | Mark subscription cancelled |
| `invoice.paid` | Record successful payment |
| `invoice.payment_failed` | Send payment failed email, update status |

### Webhook Handler

```go
func handleStripeWebhook(c echo.Context) error {
    payload, err := ioutil.ReadAll(c.Request().Body)
    if err != nil {
        return err
    }

    event, err := webhook.ConstructEvent(
        payload,
        c.Request().Header.Get("Stripe-Signature"),
        webhookSecret,
    )
    if err != nil {
        return c.JSON(400, map[string]string{"error": "Invalid signature"})
    }

    switch event.Type {
    case "checkout.session.completed":
        // Create subscription in database
    case "customer.subscription.updated":
        // Update subscription status
    case "customer.subscription.deleted":
        // Mark as cancelled
    case "invoice.paid":
        // Record payment
    case "invoice.payment_failed":
        // Handle failed payment
    }

    return c.JSON(200, map[string]string{"received": "true"})
}
```

---

## 🛡️ Access Control Flow

### Video Access Check

```
Scout clicks "Watch Highlight"
         │
         ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│  Frontend calls: GET /api/v1/access/check/:videoId                          │
└─────────────────────────────────────────────────────────────────────────────┘
         │
         ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│  Backend checks:                                                             │
│  1. Is user authenticated? (JWT valid)                                      │
│  2. Does user have active subscription?                                     │
│  3. If Basic plan: Has user exceeded monthly limit?                         │
│  4. Is video accessible at this subscription tier?                          │
└─────────────────────────────────────────────────────────────────────────────┘
         │
         ├── NO → Return paywall message
         │
         └── YES → Generate signed URL (4-hour expiry)
                   Record view in usage_records
                   Return stream URL
```

---

## 📊 Key Metrics to Track

| Metric | Description | Why It Matters |
|--------|-------------|----------------|
| **MRR** | Monthly Recurring Revenue | Core business health |
| **Conversion Rate** | Free → Paid | Marketing effectiveness |
| **Churn Rate** | Cancellations / Total | Retention health |
| **ARPU** | Average Revenue Per User | Pricing validation |
| **LTV** | Lifetime Value | Customer acquisition budget |
| **Trial → Paid** | Trial conversions | Onboarding effectiveness |

---

## 🚀 Implementation Priority

### MVP (Week 1-2)

- [ ] Stripe account setup
- [ ] Plans table with Basic/Pro
- [ ] Checkout flow (Stripe hosted)
- [ ] Webhook handler
- [ ] Access check endpoint
- [ ] Basic subscription status page

### Enhancement (Week 3-4)

- [ ] Usage tracking
- [ ] Subscription management (upgrade/downgrade)
- [ ] Cancel flow
- [ ] Payment history
- [ ] Email notifications (welcome, payment, cancelled)

### Scale (Month 2+)

- [ ] Enterprise plan with custom pricing
- [ ] Annual billing discounts
- [ ] Team/seat management
- [ ] Invoice generation
- [ ] Analytics dashboard
