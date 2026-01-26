# Unicorn Sport - Security Documentation

## 🔐 Overview

Security is paramount for Unicorn Sport given we handle:

- **Minor children's data** (players aged 10-21)
- **Identity verification documents** (NIN, Passport)
- **Financial transactions** (scout subscriptions via Stripe)
- **Sensitive personal information** (location, contact details)

> **Key Principle:** Academy creates all content. No user-generated uploads = reduced attack surface.

---

## 🏛️ Security Architecture

### MVP Architecture (Scalable Monolith)

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                    SECURITY LAYERS (MVP)                                     │
└─────────────────────────────────────────────────────────────────────────────┘

                        ┌─────────────────────────┐
                        │      CLOUDFLARE         │
                        │  • DDoS Protection      │
                        │  • WAF Rules            │
                        │  • SSL Termination      │
                        │  • Rate Limiting        │
                        └───────────┬─────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                      SINGLE GO BINARY                                        │
│  ┌────────────────────────────────────────────────────────────────────┐     │
│  │                    MIDDLEWARE CHAIN                                 │     │
│  │  ┌─────────┐  ┌─────────┐  ┌─────────┐  ┌─────────┐  ┌─────────┐  │     │
│  │  │  CORS   │→ │ Logger  │→ │JWT Auth │→ │  RBAC   │→ │Rate Lim │  │     │
│  │  └─────────┘  └─────────┘  └─────────┘  └─────────┘  └─────────┘  │     │
│  └────────────────────────────────────────────────────────────────────┘     │
│                                    │                                         │
│  ┌─────────────────────────────────┼─────────────────────────────────┐      │
│  │                         MODULES (Each has own security context)    │      │
│  │  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────┐          │      │
│  │  │   Auth   │  │ Players  │  │  Media   │  │  Admin   │          │      │
│  │  │ Module   │  │  Module  │  │  Module  │  │  Module  │          │      │
│  │  └────┬─────┘  └────┬─────┘  └────┬─────┘  └────┬─────┘          │      │
│  │       │             │             │             │                 │      │
│  │       └─────────────┴─────────────┴─────────────┘                 │      │
│  │                            │                                       │      │
│  └────────────────────────────┼───────────────────────────────────────┘      │
│                               │                                              │
│  ┌────────────────────────────┼───────────────────────────────────────┐      │
│  │           SHARED DATA ACCESS LAYER (Repository Interfaces)         │      │
│  │  • Each module accesses ONLY its own tables                        │      │
│  │  • Cross-module communication via Go interfaces (not DB joins)     │      │
│  │  • Easy to extract to separate service later                       │      │
│  └────────────────────────────┼───────────────────────────────────────┘      │
└───────────────────────────────┼──────────────────────────────────────────────┘
                                │
                                ▼
                    ┌───────────────────────┐
                    │   PostgreSQL 16       │
                    │  Single DB, but:      │
                    │  • Table ownership    │
                    │  • Foreign keys only  │
                    │    within module      │
                    └───────────────────────┘
```

### Module Isolation Pattern (Scalable Design)

```go
// Each module owns its tables and exposes interfaces, NOT direct DB access
// This makes future microservice extraction trivial

// internal/players/repository.go - PLAYERS module owns these tables
type PlayerRepository interface {
    Create(ctx context.Context, player Player) error
    GetByID(ctx context.Context, id string) (*Player, error)
    Search(ctx context.Context, filters SearchFilters) ([]Player, error)
    // Players module NEVER queries videos table directly
}

// internal/media/repository.go - MEDIA module owns these tables  
type VideoRepository interface {
    Create(ctx context.Context, video Video) error
    GetByPlayerID(ctx context.Context, playerID string) ([]Video, error)
    // Media module NEVER queries players table directly
}

// Cross-module communication via interfaces (not DB joins!)
type PlayerService interface {
    GetPlayerForVideo(ctx context.Context, playerID string) (*PlayerSummary, error)
}

// When extracting to microservices, just swap interface implementation:
// - In monolith: Direct function call
// - In microservice: HTTP/gRPC client
```

---

## 🔑 Authentication

### JWT Token Structure (Simple & Clean)

```json
{
  "header": {
    "alg": "HS256",
    "typ": "JWT"
  },
  "payload": {
    "sub": "user-uuid",
    "role": "scout",
    "tier": "pro",
    "iat": 1706180400,
    "exp": 1706181300
  }
}
```

> **Note:** No `permissions` array in JWT. Permissions derived from `role` + `tier` at runtime.
> This keeps tokens small and allows permission changes without re-issuing tokens.

### Token Security

| Aspect | Implementation |
|--------|----------------|
| **Signing Algorithm** | HS256 (HMAC-SHA256) |
| **Secret Key** | 256-bit, from environment variable |
| **Access Token TTL** | 15 minutes |
| **Refresh Token TTL** | 7 days |
| **Token Storage** | HTTP-only secure cookies |
| **Revocation** | Refresh token hashed in DB |

### Password Requirements (MVP - Keep Simple)

```
Minimum: 8 characters
At least: 1 uppercase, 1 lowercase, 1 number

// Bcrypt cost 12
hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
```

---

## 👮 Authorization (RBAC)

### MVP Roles (3 Only)

| Role | Description | Who Creates |
|------|-------------|-------------|
| `admin` | Platform operators | Seeded in DB |
| `scout` | Talent scouts/clubs | Self-register |
| `player` | Young athletes | **Admin creates** (NOT self-register) |

> **Critical:** Players do NOT self-register. Admin creates profiles with NIN/Passport verification 
> BEFORE tournament. Player may later receive login credentials to view their own profile.

### Permission Matrix (Simplified)

```
┌─────────────────────┬─────────┬───────────────────────────────┬─────────┐
│ Resource            │ Admin   │ Scout (by tier)               │ Player  │
│                     │         │ Free    Scout   Pro    Club   │         │
├─────────────────────┼─────────┼───────────────────────────────┼─────────┤
│ Players - Create    │ ✅      │ ❌      ❌      ❌     ❌     │ ❌      │
│ Players - Read List │ ✅      │ ✅      ✅      ✅     ✅     │ ❌      │
│ Players - Read Full │ ✅      │ ✅      ✅      ✅     ✅     │ Own ✅  │
│                     │         │                               │         │
│ Highlights (FREE)   │ ✅      │ ✅      ✅      ✅     ✅     │ Own ✅  │
│ Full Matches (PAID) │ ✅      │ ❌      ✅      ✅     ✅     │ Own ✅  │
│                     │         │                               │         │
│ Videos - Upload     │ ✅      │ ❌      ❌      ❌     ❌     │ ❌      │
│ Videos - Delete     │ ✅      │ ❌      ❌      ❌     ❌     │ ❌      │
│                     │         │                               │         │
│ Contact Request     │ ✅      │ ❌      ❌      ✅     ✅     │ ❌      │
│ API Access          │ ✅      │ ❌      ❌      ❌     ✅     │ ❌      │
│                     │         │                               │         │
│ Subscriptions       │ CRUD    │ Own     Own     Own    Own    │ ❌      │
│ User Management     │ CRUD    │ ❌      ❌      ❌     ❌     │ ❌      │
└─────────────────────┴─────────┴───────────────────────────────┴─────────┘
```

### Authorization Middleware (Clean Implementation)

```go
// internal/middleware/auth.go

type AuthMiddleware struct {
    jwtSecret []byte
}

func (m *AuthMiddleware) RequireRole(allowedRoles ...string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            claims := r.Context().Value("claims").(*Claims)
            
            for _, role := range allowedRoles {
                if claims.Role == role {
                    next.ServeHTTP(w, r)
                    return
                }
            }
            
            http.Error(w, "Forbidden", http.StatusForbidden)
        })
    }
}

func (m *AuthMiddleware) RequireTier(minTier string) func(http.Handler) http.Handler {
    tierLevel := map[string]int{"free": 0, "scout": 1, "pro": 2, "club": 3}
    
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            claims := r.Context().Value("claims").(*Claims)
            
            if tierLevel[claims.Tier] >= tierLevel[minTier] {
                next.ServeHTTP(w, r)
                return
            }
            
            // Return 402 Payment Required for upgrade prompt
            http.Error(w, "Subscription upgrade required", http.StatusPaymentRequired)
        })
    }
}
```

### Route Protection Example

```go
// internal/router/router.go

func SetupRoutes(r chi.Router, auth *AuthMiddleware) {
    // Public routes
    r.Post("/api/auth/login", authHandler.Login)
    r.Post("/api/auth/register", authHandler.RegisterScout)
    
    // Authenticated routes
    r.Group(func(r chi.Router) {
        r.Use(auth.RequireAuth)
        
        // All authenticated users
        r.Get("/api/players", playerHandler.List)
        r.Get("/api/players/{id}", playerHandler.Get)
        
        // Highlights - everyone can view
        r.Get("/api/videos/highlights", videoHandler.ListHighlights)
        
        // Full matches - paid scouts only
        r.With(auth.RequireTier("scout")).Get("/api/videos/full-matches", videoHandler.ListFullMatches)
        
        // Contact requests - pro+ only  
        r.With(auth.RequireTier("pro")).Post("/api/contact-requests", contactHandler.Create)
        
        // Admin only
        r.With(auth.RequireRole("admin")).Post("/api/players", playerHandler.Create)
        r.With(auth.RequireRole("admin")).Post("/api/videos", videoHandler.Upload)
        r.With(auth.RequireRole("admin")).Delete("/api/videos/{id}", videoHandler.Delete)
    })
}
```

---

## 🔒 Data Protection

### Encryption Strategy (MVP)

| Layer | Implementation | Provider |
|-------|----------------|----------|
| **In Transit** | TLS 1.3 via Cloudflare | Cloudflare (free) |
| **At Rest (DB)** | PostgreSQL native encryption | AWS RDS/Managed |
| **At Rest (Blobs)** | AWS S3 SSE | AWS-managed keys |
| **Passwords** | Bcrypt cost 12 | Application |

> **MVP Note:** No KMS needed initially. Use environment variables for secrets.
> Move to AWS Secrets Manager when handling >1000 users or for SOC2 compliance.

### Sensitive Data Classification

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                    DATA CLASSIFICATION                                       │
└─────────────────────────────────────────────────────────────────────────────┘

🔴 CRITICAL (Never in API responses, admin-only access):
├── verification_doc_url (NIN/Passport scans)
├── password_hash
└── stripe_customer_id, stripe_subscription_id

🟠 SENSITIVE (Masked in API, full access for owner only):
├── date_of_birth → Return as "age" (e.g., 17)
├── last_name → Abbreviated (e.g., "James O.")
└── email → Masked for non-owners (e.g., "j***@gmail.com")

🟡 INTERNAL (Not in public API):
├── admin_notes
├── verified_by (admin user ID)
└── created_by (admin user ID)

🟢 PUBLIC (Available in player cards):
├── first_name
├── position, country
├── profile_photo_url, thumbnail_url
└── highlight video URLs (FREE)
```

### API Response Masking

```go
// internal/players/dto.go

// Public response - for player cards and search results
type PlayerPublicDTO struct {
    ID           string `json:"id"`
    FirstName    string `json:"first_name"`
    LastNameInit string `json:"last_name_init"` // "O." not "Okonkwo"
    Age          int    `json:"age"`             // 17, not DOB
    Position     string `json:"position"`
    Country      string `json:"country"`
    ThumbnailURL string `json:"thumbnail_url"`
    IsVerified   bool   `json:"is_verified"`
}

// Full response - for player detail page (authenticated scouts)
type PlayerDetailDTO struct {
    PlayerPublicDTO
    State           string   `json:"state,omitempty"`
    City            string   `json:"city,omitempty"`
    HeightCM        int      `json:"height_cm,omitempty"`
    WeightKG        int      `json:"weight_kg,omitempty"`
    PreferredFoot   string   `json:"preferred_foot,omitempty"`
    SchoolName      string   `json:"school_name,omitempty"`
    TournamentName  string   `json:"tournament_name,omitempty"`
    TournamentYear  int      `json:"tournament_year,omitempty"`
    HighlightVideos []Video  `json:"highlight_videos"`        // FREE
    FullMatchVideos []Video  `json:"full_match_videos,omitempty"` // Only if paid tier
}

func ToPublicDTO(p *Player) PlayerPublicDTO {
    return PlayerPublicDTO{
        ID:           p.ID,
        FirstName:    p.FirstName,
        LastNameInit: string(p.LastName[0]) + ".",
        Age:          calculateAge(p.DateOfBirth),
        Position:     p.Position,
        Country:      p.Country,
        ThumbnailURL: p.ThumbnailURL,
        IsVerified:   p.VerificationStatus == "verified",
    }
}
```

---

## 👶 Child Safety

### Age Verification (MVP - Pre-Tournament)

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                PLAYER VERIFICATION FLOW (Admin-Controlled)                   │
└─────────────────────────────────────────────────────────────────────────────┘

┌──────────────────────────────────────────────────────────────────────────────┐
│  BEFORE TOURNAMENT (In-Person)                                               │
│  ┌─────────────────────────────────────────────────────────────────────────┐ │
│  │  1. Parent/Guardian presents NIN or Passport                            │ │
│  │  2. Academy staff verifies age (must be under 21)                       │ │
│  │  3. Staff photographs document (stored in AWS S3, admin-only)         │ │
│  │  4. Player profile created with verified status                          │ │
│  └─────────────────────────────────────────────────────────────────────────┘ │
└──────────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌──────────────────────────────────────────────────────────────────────────────┐
│  DURING TOURNAMENT                                                           │
│  ┌─────────────────────────────────────────────────────────────────────────┐ │
│  │  • Academy media team films matches                                      │ │
│  │  • Staff creates highlight clips                                         │ │
│  │  • Videos linked to verified player profiles                             │ │
│  └─────────────────────────────────────────────────────────────────────────┘ │
└──────────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌──────────────────────────────────────────────────────────────────────────────┐
│  POST TOURNAMENT (Optional)                                                  │
│  ┌─────────────────────────────────────────────────────────────────────────┐ │
│  │  • Admin may issue login credentials to player                           │ │
│  │  • Player can view own profile (read-only)                               │ │
│  │  • Player CANNOT edit or upload anything                                  │ │
│  └─────────────────────────────────────────────────────────────────────────┘ │
└──────────────────────────────────────────────────────────────────────────────┘
```

> **Key Insight:** No self-registration = No age verification flow needed in app.
> Age is verified IN PERSON before tournament. Much simpler and more reliable.

### Minor-Specific Protections

| Protection | Implementation |
|------------|----------------|
| **No Direct Contact** | Scouts submit contact requests, admin forwards |
| **Masked Information** | Last name abbreviated, exact DOB hidden |
| **No User Uploads** | All content from academy staff |
| **Secure Documents** | NIN/Passport URLs never in API responses |
| **Deletion Requests** | Honored within 7 days (admin action) |

---

## 📝 Audit Logging (MVP)

### Simple Audit Table

```sql
CREATE TABLE audit_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    timestamp TIMESTAMP DEFAULT NOW(),
    user_id UUID,
    action VARCHAR(100) NOT NULL,
    resource_type VARCHAR(50),
    resource_id UUID,
    ip_address INET,
    user_agent TEXT,
    details JSONB,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_audit_timestamp ON audit_logs(timestamp);
CREATE INDEX idx_audit_user ON audit_logs(user_id);
CREATE INDEX idx_audit_action ON audit_logs(action);
```

### What We Log (MVP)

| Category | Events |
|----------|--------|
| **Auth** | login, logout, login_failed, password_reset |
| **Player** | player_created, player_updated, player_verified |
| **Video** | video_uploaded, video_deleted, video_viewed |
| **Subscription** | subscription_created, subscription_upgraded, subscription_canceled |
| **Contact** | contact_request_created, contact_request_forwarded |
| **Admin** | admin_action (catch-all for admin operations) |

### Simple Logging Middleware

```go
// internal/middleware/audit.go

func AuditMiddleware(db *sql.DB) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Skip GET requests (optional - log reads if needed for compliance)
            if r.Method == "GET" {
                next.ServeHTTP(w, r)
                return
            }
            
            claims, _ := r.Context().Value("claims").(*Claims)
            userID := ""
            if claims != nil {
                userID = claims.Sub
            }
            
            // Log the action
            go logAudit(db, AuditEntry{
                UserID:     userID,
                Action:     r.Method + " " + r.URL.Path,
                IPAddress:  r.RemoteAddr,
                UserAgent:  r.UserAgent(),
            })
            
            next.ServeHTTP(w, r)
        })
    }
}
```

### Log Retention (MVP)

| Log Type | Retention | Notes |
|----------|-----------|-------|
| Audit logs | 1 year | Required for compliance |
| Application logs | 30 days | Debug purposes |

---

## 🚀 Scaling Security (Future)

When ready to scale beyond MVP:

| Current (MVP) | Future (Scale) |
|---------------|----------------|
| Environment variables for secrets | AWS Secrets Manager |
| Single DB with table ownership | Database per service |
| Cloudflare free plan | Cloudflare Pro + WAF rules |
| Simple audit table | Dedicated audit service |
| Bcrypt passwords | + WebAuthn/Passkeys option |
| HTTP-only cookies | + OAuth2 for mobile apps |

### Extracting Services (When Team > 5)

```go
// Step 1: Current (MVP) - Direct function call
type AuthService struct {
    repo UserRepository
}

func (s *AuthService) ValidateToken(token string) (*Claims, error) {
    // Direct implementation
}

// Step 2: Future - Same interface, HTTP client implementation
type AuthServiceClient struct {
    baseURL string
    client  *http.Client
}

func (c *AuthServiceClient) ValidateToken(token string) (*Claims, error) {
    // HTTP call to auth microservice
    resp, _ := c.client.Post(c.baseURL+"/validate", ...)
}

// The rest of the code doesn't change - just swap the implementation
```

---

## ✅ Security Checklist (MVP Launch)

```
Pre-Launch:
□ HTTPS enabled via Cloudflare
□ JWT secret in environment variable (not in code)
□ Database password in environment variable
□ AWS S3 storage private buckets
□ verification_doc_url never in API responses
□ Password bcrypt cost = 12
□ Rate limiting on /api/auth/* endpoints
□ CORS configured for unicornsport.africa only

Post-Launch (Week 1):
□ Monitor failed login attempts
□ Review audit logs for anomalies
□ AWS S3 storage private buckets verified
□ verification_doc_url never in API responses
□ Test subscription paywall enforcement
□ Verify admin-only routes are protected

Ongoing:
□ Rotate JWT secret quarterly
□ Review access logs monthly
□ Update dependencies for security patches
```

---

## 🚨 Security Monitoring (MVP)

### Simple Alerts via Cloudflare + Application Logs

| Trigger | Action |
|---------|--------|
| 5+ failed logins | Account lockout (15 min) |
| Rate limit exceeded | IP throttle via Cloudflare |
| Admin action | Log to audit table |
| Suspicious pattern | Manual review (daily) |

> **MVP Approach:** Use Cloudflare's built-in security features + application logs.
> No need for Elasticsearch/Kibana stack until you have a dedicated DevOps person.

---

## 🛡️ Security Headers (via Cloudflare)

Configure in Cloudflare Dashboard → Rules → Transform Rules:

```
X-Frame-Options: DENY
X-Content-Type-Options: nosniff
Referrer-Policy: strict-origin-when-cross-origin
Strict-Transport-Security: max-age=31536000
```

---

## 📜 Compliance Notes

### Key Requirements for Nigeria

| Requirement | How We Address |
|-------------|----------------|
| **NDPA Compliance** | Data stored in AWS EU region, consent captured |
| **Minor Protection** | Admin-controlled profiles, no direct contact |
| **Parental Awareness** | In-person verification requires parent/guardian |
| **Data Deletion** | Admin can delete within 7 days on request |

### International Users

For scouts outside Nigeria:
- GDPR (EU): Data export and deletion on request
- Standard privacy policy covering data usage

---

## 🔐 Secret Management (MVP)

### Environment Variables

```bash
# .env (NEVER commit to git)
DATABASE_URL=postgres://user:password@localhost:5432/unicornsport
JWT_SECRET=your-256-bit-secret-here
AWS_ACCESS_KEY_ID=AKIAXXXXXXXXXXXXXXXX
AWS_SECRET_ACCESS_KEY=xxxxx
AWS_S3_BUCKET=unicornsport-media
STRIPE_SECRET_KEY=sk_live_...

# Production: Use Docker secrets or AWS Secrets Manager
```

### Secret Rotation Schedule

| Secret | Frequency | Method |
|--------|-----------|--------|
| JWT_SECRET | Quarterly | Redeploy with new value |
| Database password | On breach only | Rotate in AWS RDS |
| Stripe keys | On breach only | Regenerate in Stripe dashboard |

---

## 📋 Final Security Checklist

### Pre-Launch (Must Have)

```
✅ Cloudflare configured with SSL
✅ All secrets in environment variables
✅ JWT auth working
✅ RBAC middleware protecting routes
✅ verification_doc_url NOT in any API response
✅ Passwords hashed with bcrypt (cost 12)
✅ Rate limiting on auth endpoints
✅ CORS restricted to unicornsport.africa
✅ Audit logging enabled
```

### Post-MVP (Nice to Have)

```
□ AWS Secrets Manager for secrets
□ Automated dependency scanning
□ Security headers hardened
□ Third-party penetration test
□ SOC2 compliance (for enterprise clients)
```
