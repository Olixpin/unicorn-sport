# Unicorn Sport - Technical Architecture

## 🏗️ System Overview

Unicorn Sport uses a **containerized modular monolith** architecture optimized for MVP speed and simplicity. The application runs as a single Go binary with logically separated modules, backed by a single PostgreSQL database.

> **Architecture Philosophy:** Start simple, extract services when you have the team size 
> and scale to justify the complexity. A well-structured monolith can be split later.

---

## 📊 MVP Architecture (Current)

```
┌─────────────────────────────────────────────────────────────────────────────────┐
│                                   CLIENTS                                        │
│  ┌─────────────────────────────────┐  ┌─────────────────────────────────┐       │
│  │         Scout/Player Web        │  │         Admin Dashboard          │       │
│  │           (Nuxt 3 SSR)          │  │          (Nuxt 3 SPA)           │       │
│  └────────────────┬────────────────┘  └────────────────┬────────────────┘       │
└───────────────────┼────────────────────────────────────┼────────────────────────┘
                    │                                    │
                    └──────────────┬─────────────────────┘
                                   │
                    ┌──────────────▼──────────────┐
                    │     Cloudflare (CDN/SSL)    │
                    │   • Static asset caching    │
                    │   • DDoS protection         │
                    │   • SSL termination         │
                    └──────────────┬──────────────┘
                                   │
                    ┌──────────────▼──────────────┐
                    │    CONTAINERIZED API        │
                    │   (Single Go Binary)        │
                    │         :8080               │
                    │                             │
                    │  ┌───────────────────────┐  │
                    │  │      MODULES          │  │
                    │  ├───────────────────────┤  │
                    │  │ • Auth (JWT, login)   │  │
                    │  │ • Profiles (players)  │  │
                    │  │ • Media (videos)      │  │
                    │  │ • Admin (CRUD, verify)│  │
                    │  │ • Subscriptions       │  │
                    │  │ • Search (discovery)  │  │
                    │  │ • Contact (inquiries) │  │
                    │  └───────────────────────┘  │
                    │                             │
                    └──────────────┬──────────────┘
                                   │
                    ┌──────────────┴──────────────┐
                    │                             │
          ┌─────────▼─────────┐       ┌──────────▼──────────┐
          │    PostgreSQL     │       │       AWS S3        │
          │   (Single DB)     │       │                     │
          │                   │       │ • Videos (public)   │
          │ • All tables      │       │ • Thumbnails        │
          │ • Full-text search│       │ • Documents (private)│
          │ • ACID compliance │       │                     │
          └───────────────────┘       └─────────────────────┘
```

### Why This Architecture?

| Decision | Rationale |
|----------|-----------|
| **Single Go binary** | Faster development, simpler deployment, easy debugging |
| **Single database** | No cross-service queries, simpler migrations, ACID transactions |
| **No message bus** | Direct function calls for MVP; add NATS when async workflows needed |
| **No Redis** | PostgreSQL is fast enough for MVP scale; add caching at 10k+ users |
| **Containerized** | Consistent environments, easy deployment, production-ready |

### Module Structure

```
backend/
├── cmd/server/main.go           # Single entry point
├── internal/
│   ├── config/                  # Configuration
│   ├── domain/                  # Shared domain models
│   ├── modules/                 # Feature modules
│   │   ├── auth/               # Authentication & JWT
│   │   ├── profiles/           # Player profiles
│   │   ├── media/              # Video management
│   │   ├── admin/              # Admin operations (P0)
│   │   ├── subscriptions/      # Stripe payments (P0)
│   │   ├── search/             # Player discovery
│   │   └── contact/            # Scout inquiries
│   ├── middleware/             # Auth, subscription checks
│   └── pkg/                    # Shared utilities
├── migrations/                  # Single migration folder
├── Dockerfile
└── docker-compose.yml
```

---

## 🔮 Target Architecture (Post-PMF)

> **When to migrate:** 5+ engineers, 50k+ users, or clear service boundaries needed.

```
┌─────────────────────────────────────────────────────────────────────────────────┐
│                                   CLIENTS                                        │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐             │
│  │   Web App   │  │  iOS App    │  │ Android App │  │  School API │             │
│  │  (Nuxt 3)   │  │  (Future)   │  │  (Future)   │  │  (Future)   │             │
│  └──────┬──────┘  └──────┬──────┘  └──────┬──────┘  └──────┬──────┘             │
└─────────┼────────────────┼────────────────┼────────────────┼────────────────────┘
          │                │                │                │
          ▼                ▼                ▼                ▼
┌─────────────────────────────────────────────────────────────────────────────────┐
│                              API GATEWAY                                         │
│                         Load Balancer / Rate Limiting                            │
└─────────────────────────────────────────────────────────────────────────────────┘
          │
          ▼
┌─────────────────────────────────────────────────────────────────────────────────┐
│                              MICROSERVICES LAYER                                 │
│                                                                                  │
│  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐                  │
│  │  Auth Service   │  │ Profile Service │  │  Media Service  │                  │
│  │     :8080       │  │     :8081       │  │     :8082       │                  │
│  └────────┬────────┘  └────────┬────────┘  └────────┬────────┘                  │
│           │                    │                    │                            │
│  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐                  │
│  │Discovery Service│  │ Subscription    │  │Verification Svc │                  │
│  │     :8083       │  │    Service      │  │    :8085        │                  │
│  └────────┬────────┘  └────────┬────────┘  └────────┬────────┘                  │
└───────────┼────────────────────┼────────────────────┼────────────────────────────┘
            │                    │                    │
            ▼                    ▼                    ▼
┌─────────────────────────────────────────────────────────────────────────────────┐
│                              MESSAGE BUS (NATS)                                  │
└─────────────────────────────────────────────────────────────────────────────────┘
            │                    │                    │
            ▼                    ▼                    ▼
┌─────────────────────────────────────────────────────────────────────────────────┐
│                              DATA LAYER                                          │
│  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐                  │
│  │   PostgreSQL    │  │     Redis       │  │    AWS S3       │                  │
│  │   (Per Service) │  │    (Caching)    │  │   (Media)       │                  │
│  └─────────────────┘  └─────────────────┘  └─────────────────┘                  │
└─────────────────────────────────────────────────────────────────────────────────┘
```

---

## 🔧 Service Details

### Auth Service (Port 8080)

**Purpose:** Handle all authentication and authorization concerns.

**Responsibilities:**
- User registration (email + password)
- Login/logout with JWT tokens
- Refresh token management
- Password reset flow
- Email/phone verification
- Role-based access control

**Database:** `auth_db`

**Tables:**
```sql
users           -- Core user accounts
refresh_tokens  -- Token management
verification_codes -- Email/SMS verification
```

**API Endpoints:**
```
POST   /api/v1/auth/register     -- Create new account
POST   /api/v1/auth/login        -- Login, get JWT
POST   /api/v1/auth/logout       -- Invalidate tokens
POST   /api/v1/auth/refresh      -- Refresh access token
POST   /api/v1/auth/verify-email -- Verify email address
POST   /api/v1/auth/forgot-password
POST   /api/v1/auth/reset-password
GET    /api/v1/auth/me           -- Get current user
```

---

### Profile Service (Port 8081)

**Purpose:** Manage user profiles for all user types (players, scouts, academies).

**Responsibilities:**
- Profile CRUD operations
- Player-specific details (position, physical stats)
- Scout-specific details (organization, regions)
- Academy profiles
- Trust level management
- Profile completion scoring

**Database:** `profile_db`

**Tables:**
```sql
profiles        -- Base profile for all users
player_details  -- Player-specific info (position, DOB, height)
scout_details   -- Scout-specific info (organization, regions)
```

**API Endpoints:**
```
GET    /api/v1/profiles/:id          -- Get profile
POST   /api/v1/profiles              -- Create profile
PUT    /api/v1/profiles/:id          -- Update profile
DELETE /api/v1/profiles/:id          -- Delete profile
GET    /api/v1/profiles/:id/player   -- Get player details
PUT    /api/v1/profiles/:id/player   -- Update player details
```

---

### Media Service (Port 8082)

**Purpose:** Handle all video upload, storage, and delivery.

**Responsibilities:**
- Video upload (chunked for large files)
- AWS S3 integration
- Thumbnail generation
- Video transcoding (future)
- Streaming URLs
- Upload progress tracking

**Database:** `media_db`

**Tables:**
```sql
videos   -- Video metadata
uploads  -- Upload progress tracking
```

**API Endpoints:**
```
POST   /api/v1/videos                -- Initiate upload
PUT    /api/v1/videos/:id/upload     -- Upload video file
GET    /api/v1/videos/:id            -- Get video details
GET    /api/v1/videos/profile/:pid   -- Get videos by profile
DELETE /api/v1/videos/:id            -- Delete video
GET    /api/v1/videos/:id/stream     -- Get streaming URL
```

---

### Discovery Service (Port 8083)

**Purpose:** Enable talent search and discovery.

**Responsibilities:**
- Full-text search on profiles
- Advanced filtering (position, location, age)
- Search result ranking
- Recommendation engine (future)
- Analytics on searches (future)

**Database:** `discovery_db`

**API Endpoints:**
```
GET    /api/v1/discover/search       -- Search players
GET    /api/v1/discover/filter       -- Filter by criteria
GET    /api/v1/discover/featured     -- Featured players
GET    /api/v1/discover/trending     -- Trending profiles
```

---

## 🚨 CRITICAL SERVICES (Must Build First)

### Admin Service (P0 - BLOCKER)

> **Without this, your media team cannot upload ANY content.**

**Purpose:** Handle admin operations for content management by the media team.

**Responsibilities:**
- Talent event management (create, list, update events)
- Match recording uploads (large file handling)
- Highlight creation and player linking
- Bulk upload operations
- Player profile management
- Analytics and reporting

**Database:** `admin_db` (or extend `media_db`)

**API Endpoints:**
```
POST   /api/v1/admin/events              -- Create talent day event
GET    /api/v1/admin/events              -- List all events
POST   /api/v1/admin/matches             -- Create match record
POST   /api/v1/admin/matches/:id/upload  -- Upload match video (chunked)
POST   /api/v1/admin/highlights          -- Create highlight
POST   /api/v1/admin/highlights/bulk     -- Bulk upload highlights
PUT    /api/v1/admin/highlights/:id/link -- Link highlight to player
GET    /api/v1/admin/dashboard           -- Admin dashboard stats
```

**Event Publications:**
```
media.highlight.linked  -- Triggered when highlight linked to player
event.created           -- Triggered when new talent day created
```

---

### Subscription Service (P0 - BLOCKER)

> **Without this, you have ZERO revenue.**

**Purpose:** Manage scout subscriptions, billing, and content access control.

**Responsibilities:**
- Subscription plan management
- Stripe/payment integration
- Access control (check subscription before video access)
- Billing and invoicing
- Trial management
- Subscription analytics

**Database:** `subscription_db`

**Tables:**
```sql
plans           -- Available subscription tiers
subscriptions   -- Active subscriptions
payments        -- Payment history
access_logs     -- Track content access
```

**API Endpoints:**
```
GET    /api/v1/subscriptions/plans       -- List available plans
POST   /api/v1/subscriptions             -- Create subscription
GET    /api/v1/subscriptions/me          -- Get my subscription
POST   /api/v1/subscriptions/checkout    -- Stripe checkout session
POST   /api/v1/webhooks/stripe           -- Stripe webhook handler
GET    /api/v1/access/check/:videoId     -- Check if user can access video
```

**Subscription Tiers:**
```
┌─────────────────────────────────────────────────────────────────────────┐
│  FREE: $0 (Lead Generation)                                              │
│  • Browse basic player profiles                                          │
│  • Watch HIGHLIGHTS (FREE but can be deceptive!)                         │
│  • Basic search                                                          │
│  • ❌ Cannot see full matches (the REAL performance)                     │
│  • ❌ Cannot contact players                                             │
├─────────────────────────────────────────────────────────────────────────┤
│  SCOUT: $99/month                                                        │
│  • Watch FULL MATCHES (see real 90-minute performance!)                  │
│  • Full player profiles                                                  │
│  • Basic analytics                                                       │
│  • ❌ Cannot contact players (Pro required)                              │
├─────────────────────────────────────────────────────────────────────────┤
│  PRO: $299/month                                                         │
│  • Everything in Scout                                                   │
│  • CONTACT players ("Write to us")                                       │
│  • Advanced analytics                                                    │
│  • Save shortlists                                                       │
│  • Email alerts                                                          │
├─────────────────────────────────────────────────────────────────────────┤
│  ENTERPRISE: $999/month                                                  │
│  • Everything in Pro                                                     │
│  • API access                                                            │
│  • 5 team seats                                                          │
│  • Priority support                                                      │
│  • Priority alerts on new talent                                         │
└─────────────────────────────────────────────────────────────────────────┘

💡 KEY INSIGHT: Highlights are FREE but can be DECEPTIVE!
   Scouts pay for FULL MATCHES to see REAL performance in context.
```

---

## P1 SERVICES (Build After P0)

### School Portal Service (P2 - Referral Partnership)

> **Referral partnership - schools refer players, not verify**

**Purpose:** Portal for partner schools to refer promising students.

**Responsibilities:**
- School registration and onboarding
- Student verification queue
- Document review interface
- Enrollment confirmation
- Verification analytics

**API Endpoints:**
```
POST   /api/v1/schools/register          -- School applies to partner
GET    /api/v1/schools/students          -- List linked students
POST   /api/v1/schools/verify/:studentId -- Verify student data
GET    /api/v1/schools/queue             -- Verification request queue
PUT    /api/v1/schools/verify/:id/approve -- Approve verification
PUT    /api/v1/schools/verify/:id/reject  -- Reject with reason
```

---

### Video Player Service (P1)

> **Scouts need to actually WATCH the highlights**

**Purpose:** Secure video streaming with subscription enforcement.

**Responsibilities:**
- Generate signed streaming URLs
- Check subscription before access
- Track video views (analytics)
- Adaptive bitrate (future)

**API Endpoints:**
```
GET    /api/v1/player/stream/:videoId    -- Get signed stream URL
POST   /api/v1/player/view/:videoId      -- Log view event
GET    /api/v1/player/watch-history      -- User's watch history
```

---

## P2 SERVICES (Build After MVP)

### Contact Service (P2)

**Purpose:** Enable scouts to contact players through the platform.

### Analytics Service (P2)

**Purpose:** Track views, engagement, and platform metrics.

---

## 🔐 Security Architecture

### Authentication Flow

```
┌──────────┐                    ┌──────────────┐                  ┌──────────────┐
│  Client  │                    │ Auth Service │                  │Other Services│
└────┬─────┘                    └──────┬───────┘                  └──────┬───────┘
     │                                 │                                  │
     │  1. POST /auth/login            │                                  │
     │  (email, password)              │                                  │
     │────────────────────────────────▶│                                  │
     │                                 │                                  │
     │  2. Validate credentials        │                                  │
     │                                 │                                  │
     │  3. Return JWT + Refresh Token  │                                  │
     │◀────────────────────────────────│                                  │
     │                                 │                                  │
     │  4. Request with JWT            │                                  │
     │  Authorization: Bearer <token>  │                                  │
     │─────────────────────────────────┼─────────────────────────────────▶│
     │                                 │                                  │
     │                                 │  5. Validate JWT (shared secret) │
     │                                 │                                  │
     │  6. Response                    │                                  │
     │◀─────────────────────────────────────────────────────────────────--│
     │                                 │                                  │
```

### JWT Structure

```json
{
  "header": {
    "alg": "HS256",
    "typ": "JWT"
  },
  "payload": {
    "sub": "user-uuid",
    "email": "user@example.com",
    "role": "player",
    "iat": 1234567890,
    "exp": 1234571490
  }
}
```

---

## 📡 Event-Driven Architecture

### NATS Message Flow

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                           EVENT FLOW EXAMPLE                                 │
│                   (Admin Uploads Highlight → Player Notified)                │
└─────────────────────────────────────────────────────────────────────────────┘

1. Admin/Media team uploads highlight via Admin Portal
   │
   ▼
2. Media Service saves to AWS S3
   │
   ▼
3. Admin links highlight to player profile
   │
   ▼
4. Media Service publishes event
   ┌─────────────────────────────────────────────────────────────────────────┐
   │  Topic: media.highlight.linked                                          │
   │  Payload: {                                                             │
   │    "event_type": "highlight.linked",                                    │
   │    "highlight_id": "uuid",                                              │
   │    "profile_id": "uuid",                                                │
   │    "match_id": "uuid",                                                  │
   │    "title": "James Okoro - Skills Highlight",                           │
   │    "timestamp": 1234567890                                              │
   │  }                                                                      │
   └─────────────────────────────────────────────────────────────────────────┘
   │
   ▼
5. Notification Service (future) notifies player
   │
   ▼
6. Discovery Service indexes new content
   │
   ▼
7. Highlight appears on player's profile

NOTE: No content moderation needed - all content is professionally produced
by the academy's media team. Quality control happens at the source.
```

---

## 🗄️ Database Design Principles

### Database Per Service

Each microservice owns its database:

| Service | Database | Reasoning |
|---------|----------|-----------|
| Auth Service | auth_db | Security isolation |
| Profile Service | profile_db | Profile data ownership |
| Media Service | media_db | Media metadata |
| Discovery Service | discovery_db | Search indexes |

### Cross-Service Data Access

Services reference each other by ID only:

- `profile_id` in media_db references profile in profile_db
- `user_id` in profile_db references user in auth_db
- No direct foreign keys across databases

---

## 🚀 Deployment Architecture

### Container Strategy

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                        KUBERNETES CLUSTER                                    │
│                                                                              │
│  ┌────────────────────────────────────────────────────────────────────────┐ │
│  │                           INGRESS CONTROLLER                            │ │
│  │                        (NGINX / AWS ALB)                                │ │
│  └────────────────────────────────────────────────────────────────────────┘ │
│                                    │                                         │
│         ┌──────────────────────────┼──────────────────────────┐             │
│         ▼                          ▼                          ▼             │
│  ┌─────────────┐           ┌─────────────┐           ┌─────────────┐       │
│  │  Frontend   │           │   Backend   │           │   Backend   │       │
│  │  Deployment │           │  Deployment │           │  Deployment │       │
│  │  (3 pods)   │           │   (Auth)    │           │  (Profile)  │       │
│  │             │           │  (3 pods)   │           │  (3 pods)   │       │
│  └─────────────┘           └─────────────┘           └─────────────┘       │
│                                                                              │
│  ┌─────────────┐           ┌─────────────┐           ┌─────────────┐       │
│  │   Backend   │           │   Backend   │           │  Worker     │       │
│  │  Deployment │           │  Deployment │           │  Deployment │       │
│  │   (Media)   │           │ (Discovery) │           │ (AI Mod)    │       │
│  │  (3 pods)   │           │  (3 pods)   │           │  (2 pods)   │       │
│  └─────────────┘           └─────────────┘           └─────────────┘       │
│                                                                              │
│  ┌────────────────────────────────────────────────────────────────────────┐ │
│  │                           STATEFUL SETS                                 │ │
│  │   ┌──────────────┐    ┌──────────────┐    ┌──────────────┐            │ │
│  │   │  PostgreSQL  │    │    Redis     │    │    NATS      │            │ │
│  │   │  (Primary +  │    │   (Cluster)  │    │   (Cluster)  │            │ │
│  │   │   Replica)   │    │              │    │              │            │ │
│  │   └──────────────┘    └──────────────┘    └──────────────┘            │ │
│  └────────────────────────────────────────────────────────────────────────┘ │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 🛠️ Technology Stack

| Layer | Technology | Purpose |
|-------|------------|---------|
| **Frontend** | Nuxt 3 + Vue 3 | SSR web application |
| **Styling** | Tailwind CSS | Utility-first CSS |
| **State** | Pinia | Frontend state management |
| **Backend** | Go 1.23 | Microservices |
| **API** | Echo (Go) | HTTP framework |
| **Database** | PostgreSQL 16 | Primary data store |
| **Cache** | Redis 7 | Caching layer |
| **Messaging** | NATS 2.10 | Event bus |
| **Storage** | AWS S3 | Media files |
| **Container** | Docker | Containerization |
| **Orchestration** | Kubernetes | Container orchestration |
| **IaC** | Terraform | Infrastructure as Code |

---

## 📈 Scalability Considerations

### Horizontal Scaling

- All services are stateless (state in DB/Redis)
- Kubernetes HPA for auto-scaling
- Database read replicas for query scaling

### Performance Optimization

- Redis caching for frequent queries
- CDN for media delivery
- Database connection pooling
- Async processing via NATS

### Future Considerations

- GraphQL API for flexible queries
- Elasticsearch for advanced search
- ML models for talent scoring
- Real-time updates via WebSockets
