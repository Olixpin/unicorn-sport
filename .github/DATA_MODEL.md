# Unicorn Sport - Data Model

## 📊 Overview

This document describes the complete data model for Unicorn Sport. The MVP uses a **single PostgreSQL database** with all tables in one schema for simplicity and ACID compliance.

> **Design Principle:** Academy creates all player profiles and content. 
> Players cannot self-register or upload. Scouts pay to access full matches.

---

## 🗂️ Entity Relationship Diagram

```
┌─────────────────────────────────────────────────────────────────────────────────────────┐
│                    UNICORN SPORT DATA MODEL (Single Database)                            │
└─────────────────────────────────────────────────────────────────────────────────────────┘

                                    USERS & AUTH
┌─────────────────────────────────────────────────────────────────────────────────────────┐
│                                                                                          │
│    ┌──────────────────┐         ┌──────────────────┐                                    │
│    │      users       │         │  refresh_tokens  │                                    │
│    ├──────────────────┤         ├──────────────────┤                                    │
│    │ id (PK)          │────────▶│ id (PK)          │                                    │
│    │ email            │         │ user_id (FK)     │                                    │
│    │ password_hash    │         │ token_hash       │                                    │
│    │ role             │         │ expires_at       │                                    │
│    │ email_verified   │         │ created_at       │                                    │
│    │ created_at       │         └──────────────────┘                                    │
│    │ updated_at       │                                                                  │
│    └────────┬─────────┘                                                                  │
│             │                                                                            │
│             │ Roles: 'admin', 'scout', 'player'                                          │
│             │ • admin: Full access, creates players                                      │
│             │ • scout: Subscribe to view content                                         │
│             │ • player: View own profile only                                            │
│             │                                                                            │
└─────────────┼────────────────────────────────────────────────────────────────────────────┘
              │
              ▼
                                    PLAYERS (Created by Admin)
┌─────────────────────────────────────────────────────────────────────────────────────────┐
│                                                                                          │
│    ┌──────────────────────────────────────────────────────────────────────────────┐     │
│    │                              players                                          │     │
│    ├──────────────────────────────────────────────────────────────────────────────┤     │
│    │ id (PK)              │ user_id (FK, nullable)  │ created_by (FK → users)     │     │
│    ├──────────────────────┼─────────────────────────┼─────────────────────────────┤     │
│    │ first_name           │ last_name               │ date_of_birth               │     │
│    │ position             │ preferred_foot          │ height_cm, weight_kg        │     │
│    │ country              │ state                   │ city                        │     │
│    │ school_name          │ tournament_id (FK)      │ tournament_year             │     │
│    ├──────────────────────┼─────────────────────────┼─────────────────────────────┤     │
│    │ verification_status  │ verification_doc_type   │ verification_doc_url        │     │
│    │ verified_at          │ verified_by (FK)        │                             │     │
│    ├──────────────────────┼─────────────────────────┼─────────────────────────────┤     │
│    │ profile_photo_url    │ thumbnail_url           │ search_vector (TSVECTOR)    │     │
│    │ created_at           │ updated_at              │                             │     │
│    └──────────────────────┴─────────────────────────┴─────────────────────────────┘     │
│                                                                                          │
│    NOTE: Players are created by ADMIN, not self-registered.                              │
│          user_id is NULL until admin issues login credentials.                           │
│          verification_doc_url is NEVER exposed to API (admin-only).                      │
│                                                                                          │
└─────────────────────────────────────────────────────────────────────────────────────────┘
              │
              ▼
                                    SCOUTS & SUBSCRIPTIONS
┌─────────────────────────────────────────────────────────────────────────────────────────┐
│                                                                                          │
│    ┌──────────────────┐         ┌──────────────────┐                                    │
│    │      scouts      │         │  subscriptions   │                                    │
│    ├──────────────────┤         ├──────────────────┤                                    │
│    │ id (PK)          │         │ id (PK)          │                                    │
│    │ user_id (FK)     │────────▶│ user_id (FK)     │                                    │
│    │ organization_name│         │ tier             │  Tiers:                            │
│    │ organization_type│         │ status           │  • 'free' - highlights only        │
│    │ country          │         │ stripe_customer  │  • 'scout' - + full matches        │
│    │ is_verified      │         │ stripe_sub_id    │  • 'pro' - + contact players       │
│    │ created_at       │         │ period_start     │  • 'club' - + API access           │
│    └──────────────────┘         │ period_end       │                                    │
│                                 │ created_at       │                                    │
│                                 └──────────────────┘                                    │
│                                                                                          │
└─────────────────────────────────────────────────────────────────────────────────────────┘
              │
              ▼
                                    VIDEOS & CONTENT
┌─────────────────────────────────────────────────────────────────────────────────────────┐
│                                                                                          │
│    ┌──────────────────┐         ┌──────────────────┐         ┌──────────────────┐       │
│    │      videos      │         │  player_videos   │         │    tournaments   │       │
│    ├──────────────────┤         ├──────────────────┤         ├──────────────────┤       │
│    │ id (PK)          │◀───────▶│ player_id (FK)   │         │ id (PK)          │       │
│    │ video_type       │         │ video_id (FK)    │         │ name             │       │
│    │ title            │         │ is_primary       │         │ year             │       │
│    │ description      │         └──────────────────┘         │ location         │       │
│    │ blob_url         │                                      │ created_at       │       │
│    │ thumbnail_url    │         video_type:                  └──────────────────┘       │
│    │ duration_seconds │         • 'highlight' - FREE                                    │
│    │ file_size_bytes  │         • 'full_match' - PAID                                   │
│    │ tournament_id    │                                                                  │
│    │ uploaded_by (FK) │         All content uploaded by ADMIN.                          │
│    │ created_at       │         No user-generated content.                              │
│    └──────────────────┘                                                                  │
│                                                                                          │
└─────────────────────────────────────────────────────────────────────────────────────────┘
              │
              ▼
                                    CONTACT & ANALYTICS
┌─────────────────────────────────────────────────────────────────────────────────────────┐
│                                                                                          │
│    ┌──────────────────┐         ┌──────────────────┐                                    │
│    │ contact_requests │         │   video_views    │                                    │
│    ├──────────────────┤         ├──────────────────┤                                    │
│    │ id (PK)          │         │ id (PK)          │                                    │
│    │ scout_id (FK)    │         │ video_id (FK)    │                                    │
│    │ player_id (FK)   │         │ viewer_id (FK)   │  "Write to us" feature             │
│    │ message          │         │ player_id (FK)   │  Scouts contact via admin,         │
│    │ status           │         │ created_at       │  not directly to players.          │
│    │ handled_by (FK)  │         └──────────────────┘                                    │
│    │ handled_at       │                                                                  │
│    │ admin_notes      │         status: 'pending', 'forwarded', 'responded'              │
│    │ created_at       │                                                                  │
│    └──────────────────┘                                                                  │
│                                                                                          │
└─────────────────────────────────────────────────────────────────────────────────────────┘
```

---

## 📋 Table Definitions

### users
Core user accounts for all roles.

| Column | Type | Nullable | Description |
|--------|------|----------|-------------|
| id | UUID | No | Primary key, auto-generated |
| email | VARCHAR(255) | No | Unique email address |
| password_hash | VARCHAR(255) | No | Bcrypt hash (cost 12) |
| role | VARCHAR(20) | No | 'admin', 'scout', or 'player' |
| email_verified | BOOLEAN | No | Default false |
| created_at | TIMESTAMP | No | Account creation |
| updated_at | TIMESTAMP | No | Last update |

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL CHECK (role IN ('admin', 'scout', 'player')),
    email_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role ON users(role);
```

---

### players
Player profiles created by admin.

| Column | Type | Nullable | Description |
|--------|------|----------|-------------|
| id | UUID | No | Primary key |
| user_id | UUID | Yes | NULL until credentials issued |
| first_name | VARCHAR(100) | No | Player's first name |
| last_name | VARCHAR(100) | No | Player's last name |
| date_of_birth | DATE | No | For age calculation |
| position | VARCHAR(50) | No | Playing position |
| preferred_foot | VARCHAR(10) | Yes | left/right/both |
| height_cm | INTEGER | Yes | Height in cm |
| weight_kg | INTEGER | Yes | Weight in kg |
| country | VARCHAR(100) | No | Required |
| state | VARCHAR(100) | Yes | State/region |
| city | VARCHAR(100) | Yes | City |
| school_name | VARCHAR(200) | Yes | Current school |
| verification_status | VARCHAR(50) | No | pending/verified/rejected |
| verification_doc_type | VARCHAR(50) | Yes | nin/passport/birth_certificate |
| verification_doc_url | TEXT | Yes | S3 URL (admin-only!) |
| verified_at | TIMESTAMP | Yes | When verified |
| verified_by | UUID | Yes | Admin who verified |
| profile_photo_url | TEXT | Yes | Profile photo |
| thumbnail_url | TEXT | Yes | Card thumbnail |
| tournament_id | UUID | Yes | FK to tournaments |
| tournament_year | INTEGER | Yes | Year participated |
| created_by | UUID | No | Admin who created |
| created_at | TIMESTAMP | No | Creation time |
| updated_at | TIMESTAMP | No | Last update |
| search_vector | TSVECTOR | No | Full-text search (generated) |

```sql
CREATE TABLE players (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    
    -- Basic Info
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    date_of_birth DATE NOT NULL,
    
    -- Football Info  
    position VARCHAR(50) NOT NULL,
    preferred_foot VARCHAR(10),
    height_cm INTEGER,
    weight_kg INTEGER,
    
    -- Location
    country VARCHAR(100) NOT NULL,
    state VARCHAR(100),
    city VARCHAR(100),
    school_name VARCHAR(200),
    
    -- Verification (NIN/Passport)
    verification_status VARCHAR(50) DEFAULT 'pending' 
        CHECK (verification_status IN ('pending', 'verified', 'rejected')),
    verification_doc_type VARCHAR(50),
    verification_doc_url TEXT,  -- NEVER expose in API
    verified_at TIMESTAMP,
    verified_by UUID REFERENCES users(id),
    
    -- Media
    profile_photo_url TEXT,
    thumbnail_url TEXT,
    
    -- Tournament
    tournament_id UUID REFERENCES tournaments(id),
    tournament_year INTEGER,
    
    -- Metadata
    created_by UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    
    -- Full-text search
    search_vector TSVECTOR GENERATED ALWAYS AS (
        setweight(to_tsvector('english', coalesce(first_name, '')), 'A') ||
        setweight(to_tsvector('english', coalesce(last_name, '')), 'A') ||
        setweight(to_tsvector('english', coalesce(position, '')), 'B') ||
        setweight(to_tsvector('english', coalesce(country, '')), 'B') ||
        setweight(to_tsvector('english', coalesce(school_name, '')), 'C')
    ) STORED
);

CREATE INDEX idx_players_user ON players(user_id);
CREATE INDEX idx_players_country ON players(country);
CREATE INDEX idx_players_position ON players(position);
CREATE INDEX idx_players_verified ON players(verification_status);
CREATE INDEX idx_players_search ON players USING GIN(search_vector);
```

---

### subscriptions
Scout subscription management (Stripe integration).

| Column | Type | Nullable | Description |
|--------|------|----------|-------------|
| id | UUID | No | Primary key |
| user_id | UUID | No | FK to users (scout) |
| tier | VARCHAR(50) | No | free/scout/pro/club |
| status | VARCHAR(50) | No | active/past_due/canceled |
| stripe_customer_id | VARCHAR(255) | Yes | Stripe customer ID |
| stripe_subscription_id | VARCHAR(255) | Yes | Stripe subscription ID |
| current_period_start | TIMESTAMP | Yes | Billing period start |
| current_period_end | TIMESTAMP | Yes | Billing period end |
| created_at | TIMESTAMP | No | Creation time |
| updated_at | TIMESTAMP | No | Last update |

```sql
CREATE TABLE subscriptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL UNIQUE REFERENCES users(id),
    tier VARCHAR(50) NOT NULL DEFAULT 'free' 
        CHECK (tier IN ('free', 'scout', 'pro', 'club')),
    status VARCHAR(50) DEFAULT 'active' 
        CHECK (status IN ('active', 'past_due', 'canceled')),
    stripe_customer_id VARCHAR(255),
    stripe_subscription_id VARCHAR(255),
    current_period_start TIMESTAMP,
    current_period_end TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_subscriptions_user ON subscriptions(user_id);
CREATE INDEX idx_subscriptions_tier ON subscriptions(tier);
CREATE INDEX idx_subscriptions_status ON subscriptions(status);
```

---

### videos
All video content (uploaded by admin only).

| Column | Type | Nullable | Description |
|--------|------|----------|-------------|
| id | UUID | No | Primary key |
| video_type | VARCHAR(50) | No | 'highlight' (FREE) or 'full_match' (PAID) |
| title | VARCHAR(255) | No | Video title |
| description | TEXT | Yes | Description |
| blob_url | TEXT | No | AWS S3 URL |
| thumbnail_url | TEXT | Yes | Thumbnail image |
| duration_seconds | INTEGER | Yes | Video length |
| file_size_bytes | BIGINT | Yes | File size |
| tournament_id | UUID | Yes | FK to tournaments |
| uploaded_by | UUID | No | Admin who uploaded |
| created_at | TIMESTAMP | No | Upload time |

```sql
CREATE TABLE videos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    video_type VARCHAR(50) NOT NULL CHECK (video_type IN ('highlight', 'full_match')),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    blob_url TEXT NOT NULL,
    thumbnail_url TEXT,
    duration_seconds INTEGER,
    file_size_bytes BIGINT,
    tournament_id UUID REFERENCES tournaments(id),
    uploaded_by UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_videos_type ON videos(video_type);
CREATE INDEX idx_videos_tournament ON videos(tournament_id);
```

---

### player_videos
Many-to-many relationship between players and videos.

```sql
CREATE TABLE player_videos (
    player_id UUID NOT NULL REFERENCES players(id) ON DELETE CASCADE,
    video_id UUID NOT NULL REFERENCES videos(id) ON DELETE CASCADE,
    is_primary BOOLEAN DEFAULT FALSE,
    PRIMARY KEY (player_id, video_id)
);

CREATE INDEX idx_player_videos_player ON player_videos(player_id);
CREATE INDEX idx_player_videos_video ON player_videos(video_id);
```

---

### contact_requests
"Write to us" feature - scouts contact players via admin (Pro+ tier only).

```sql
CREATE TABLE contact_requests (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),  -- The scout making the request
    player_id UUID NOT NULL REFERENCES players(id),
    message TEXT NOT NULL,
    status VARCHAR(50) DEFAULT 'pending' 
        CHECK (status IN ('pending', 'forwarded', 'responded', 'rejected')),
    handled_by UUID REFERENCES users(id),
    handled_at TIMESTAMP,
    admin_notes TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_contact_requests_status ON contact_requests(status);
CREATE INDEX idx_contact_requests_user ON contact_requests(user_id);
CREATE INDEX idx_contact_requests_player ON contact_requests(player_id);
```

---

## 📋 Access Control Summary

| Table | Admin | Scout (Free) | Scout (Paid) | Player |
|-------|-------|--------------|--------------|--------|
| **users** | CRUD | Read own | Read own | Read own |
| **players** | CRUD | Read basic | Read full | Read own |
| **videos (highlight)** | CRUD | Read ✅ | Read ✅ | Read own |
| **videos (full_match)** | CRUD | ❌ | Read ✅ | Read own |
| **subscriptions** | CRUD | Read own | Read own | ❌ |
| **contact_requests** | CRUD | ❌ | Create (Pro+) | ❌ |
| **saved_players** | ❌ | ❌ | CRUD | ❌ |

---

## 🔒 Data Protection Notes

1. **verification_doc_url** - NEVER expose in API responses
2. **password_hash** - NEVER return in any response
3. **stripe_*_id** - Internal only, not for API
4. **admin_notes** - Internal only
5. **Player last names** - Abbreviated in public API ("James O.")
6. **Date of birth** - Return age only in public API

---

## 📋 Additional MVP Tables

### scouts
Scout profile details (extends users).

```sql
CREATE TABLE scouts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL UNIQUE REFERENCES users(id),
    organization_name VARCHAR(200) NOT NULL,
    organization_type VARCHAR(50),  -- 'agency', 'club', 'independent'
    country VARCHAR(100),
    is_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_scouts_user ON scouts(user_id);
```

---

### tournaments
Tournament/event records.

```sql
CREATE TABLE tournaments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(200) NOT NULL,
    year INTEGER NOT NULL,
    location VARCHAR(200),
    country VARCHAR(100),
    start_date DATE,
    end_date DATE,
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_tournaments_year ON tournaments(year);
```

---

### saved_players
Scout's saved/favorited players (Scout+ tier).

```sql
CREATE TABLE saved_players (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    player_id UUID NOT NULL REFERENCES players(id),
    notes TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(user_id, player_id)
);

CREATE INDEX idx_saved_players_user ON saved_players(user_id);
CREATE INDEX idx_saved_players_player ON saved_players(player_id);
```

---

### video_views
Analytics tracking for video views.

```sql
CREATE TABLE video_views (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    video_id UUID NOT NULL REFERENCES videos(id),
    viewer_id UUID REFERENCES users(id),  -- NULL for anonymous
    player_id UUID REFERENCES players(id), -- Player being viewed
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_video_views_video ON video_views(video_id);
CREATE INDEX idx_video_views_viewer ON video_views(viewer_id);
CREATE INDEX idx_video_views_date ON video_views(created_at);
```

---

### audit_logs
Simple audit trail for compliance.

```sql
CREATE TABLE audit_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    action VARCHAR(100) NOT NULL,
    resource_type VARCHAR(50) NOT NULL,
    resource_id UUID,
    details JSONB,
    ip_address INET,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_audit_logs_user ON audit_logs(user_id);
CREATE INDEX idx_audit_logs_action ON audit_logs(action);
CREATE INDEX idx_audit_logs_date ON audit_logs(created_at);
```

---

## 📊 Complete Table Summary

| Table | Purpose | Key Foreign Keys |
|-------|---------|------------------|
| **users** | All user accounts | - |
| **refresh_tokens** | JWT refresh tokens | user_id |
| **scouts** | Scout profile details | user_id |
| **players** | Player profiles | user_id, tournament_id, created_by |
| **tournaments** | Tournament/events | created_by |
| **subscriptions** | Stripe subscriptions | user_id |
| **videos** | Video metadata | tournament_id, uploaded_by |
| **player_videos** | Player-video links | player_id, video_id |
| **contact_requests** | Scout inquiries | scout_id, player_id |
| **saved_players** | Scout favorites | user_id, player_id |
| **video_views** | View analytics | video_id, viewer_id |
| **audit_logs** | Audit trail | user_id |

---

## 🚀 Migration Order

Run migrations in this order to respect foreign key dependencies:

```
1. users
2. refresh_tokens
3. scouts
4. tournaments
5. players
6. subscriptions
7. videos
8. player_videos
9. contact_requests
10. saved_players
11. video_views
12. audit_logs
```
