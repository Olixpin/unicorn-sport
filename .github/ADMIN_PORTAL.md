# Unicorn Sport - Admin Portal & Media Team Workflow

## 🚨 Priority: P0 - CRITICAL BLOCKER

> **Without this, you CANNOT create player profiles or upload content.**
> **This must be built FIRST before any other feature.**

---

## 🎯 Overview

The Admin Portal is the **complete management hub** for Unicorn Sport. The academy:

- **Creates ALL player profiles** (players do not self-register)
- **Uploads ALL video content** (full matches + highlights)
- **Controls ALL player information** (players cannot edit their profiles)

> **Key Principle**: Players DO NOT register themselves or upload content. Academy staff creates accounts, builds profiles, and uploads all media after tournaments.

---

## 👥 Admin User Roles

| Role | Description | Permissions |
|------|-------------|-------------|
| **Super Admin** | Platform owners | All operations, user management, financials |
| **Admin** | Academy staff | Create player profiles, manage tournaments |
| **Media Manager** | Head of media team | Upload/manage all content, analytics |
| **Media Editor** | Video editor | Upload content, link to players |
| **Support** | Customer support | View-only, handle inquiries |

---

## 📋 Complete Admin Workflow

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                        COMPLETE ADMIN WORKFLOW                               │
└─────────────────────────────────────────────────────────────────────────────┘

PRE-TOURNAMENT
┌─────────────────────────────────────────────────────────────────────────────┐
│  1. Create Tournament/Event in system                                        │
│  2. Register participants (schools/individuals)                             │
│  3. Verify age via NIN or International Passport                            │
│  4. Approve verified players for participation                              │
└─────────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
DURING TOURNAMENT
┌─────────────────────────────────────────────────────────────────────────────┐
│  5. Media team records full matches                                          │
│  6. Collect player data (personal info, performance metrics)                │
└─────────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
POST-TOURNAMENT
┌─────────────────────────────────────────────────────────────────────────────┐
│  7. CREATE PLAYER PROFILES (name, DOB, stats, school, photo)                │
│  8. Generate login credentials for each player                              │
│  9. Upload sample matches to YouTube (marketing only + link to platform)    │
│  10. Create individual player highlights                                    │
│  11. Upload HIGHLIGHTS to Platform (FREE - attracts scouts)                 │
│  12. Upload FULL MATCHES to Platform (PAID - premium product!)              │
│  13. Link content to player profiles                                        │
│  14. Send login credentials to players                                      │
│                                                                              │
│  ⚠️ KEY INSIGHT: Highlights are FREE but can be DECEPTIVE!                 │
│     Full matches show REAL performance → scouts PAY to verify.              │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 📹 Content Upload Workflow

### The Complete Pipeline

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     CONTENT CREATION & UPLOAD PIPELINE                       │
└─────────────────────────────────────────────────────────────────────────────┘

  TALENT DAY                    POST-PRODUCTION                  PLATFORM
  (Field Work)                  (Office Work)                    (Upload)
       │                              │                              │
       ▼                              ▼                              ▼
┌─────────────────┐          ┌─────────────────┐          ┌─────────────────┐
│ 1. RECORDING    │          │ 3. EDITING      │          │ 5. UPLOAD       │
│                 │          │                 │          │                 │
│ • Set up        │          │ • Review full   │          │ • Login to      │
│   cameras       │          │   match footage │          │   admin portal  │
│ • Record full   │─────────▶│ • Extract       │─────────▶│ • Create match  │
│   matches       │          │   highlights    │          │   record        │
│ • Note player   │          │ • Create        │          │ • Upload video  │
│   positions     │          │   thumbnails    │          │   files         │
│                 │          │                 │          │                 │
└─────────────────┘          └─────────────────┘          └─────────────────┘
       │                              │                              │
       ▼                              ▼                              ▼
┌─────────────────┐          ┌─────────────────┐          ┌─────────────────┐
│ 2. PLAYER LOG   │          │ 4. METADATA     │          │ 6. LINK TO      │
│                 │          │                 │          │    PLAYERS      │
│ • Record which  │          │ • Player names  │          │ • Search player │
│   players       │          │   for each clip │          │   profiles      │
│   attended      │          │ • Duration      │          │ • Attach clips  │
│ • Match teams   │          │ • Tags (skills) │          │   to profiles   │
│ • Capture       │          │ • Match date    │          │ • Notify        │
│   player IDs    │          │ • Location      │          │   players       │
│                 │          │                 │          │                 │
└─────────────────┘          └─────────────────┘          └─────────────────┘
```

---

## 🖥️ Admin Portal Screens

### Dashboard

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  UNICORN SPORT ADMIN PORTAL                           [Media Manager] [Logout]│
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  📊 QUICK STATS                                                             │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐    │
│  │ Total        │  │ Highlights   │  │ Pending      │  │ Active       │    │
│  │ Matches: 156 │  │ Created: 892 │  │ Uploads: 12  │  │ Scouts: 45   │    │
│  └──────────────┘  └──────────────┘  └──────────────┘  └──────────────┘    │
│                                                                              │
│  📅 RECENT TALENT DAYS                                                      │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │ Date       │ Location        │ Players │ Matches │ Highlights │ Status│ │
│  ├─────────────────────────────────────────────────────────────────────┤   │
│  │ 2026-01-20 │ Lagos Stadium   │ 48      │ 4       │ 36         │ ✅    │ │
│  │ 2026-01-15 │ Abuja Field     │ 32      │ 3       │ 24         │ ✅    │ │
│  │ 2026-01-10 │ Accra Academy   │ 56      │ 5       │ 0          │ ⏳    │ │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                              │
│  [+ Create New Talent Day]  [Upload Content]  [Manage Players]              │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Create Talent Day Event

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  CREATE TALENT DAY EVENT                                                     │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  Event Details                                                               │
│  ─────────────                                                               │
│  Event Name:      [Lagos Regional Showcase - January 2026          ]        │
│  Date:            [2026-01-20]                                              │
│  Location:        [Lagos National Stadium                           ]        │
│  Country:         [Nigeria       ▼]                                         │
│                                                                              │
│  Participants                                                                │
│  ────────────                                                                │
│  Expected Players: [48    ]                                                  │
│  Number of Matches: [4     ]                                                 │
│                                                                              │
│  Partner Schools (Optional)                                                  │
│  ─────────────────────────                                                   │
│  [✓] Lagos High School                                                       │
│  [✓] Victoria Island Academy                                                 │
│  [ ] Search and add more...                                                  │
│                                                                              │
│  [Cancel]                                              [Create Event]        │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Upload Match Recording

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  UPLOAD MATCH RECORDING                                                      │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  Select Event: [Lagos Regional Showcase - January 2026     ▼]                │
│                                                                              │
│  Match Details                                                               │
│  ─────────────                                                               │
│  Match Number:    [1 of 4]                                                   │
│  Team A:          [Red Team - U16                           ]                │
│  Team B:          [Blue Team - U16                          ]                │
│  Duration:        [Will be detected from file               ]                │
│                                                                              │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │                                                                      │   │
│  │     📁 Drag and drop full match video here                          │   │
│  │                                                                      │   │
│  │        or click to browse                                           │   │
│  │                                                                      │   │
│  │     Supported: MP4 (H.264/H.265) up to 25GB                         │   │
│  │                                                                      │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                              │
│  Upload Progress: [████████████░░░░░░░░] 62% - 8.2GB / 13.2GB               │
│  Estimated time remaining: 4 minutes                                        │
│                                                                              │
│  [Cancel Upload]                                         [Next: Add Players]│
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Upload Player Highlights (Bulk)

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  UPLOAD PLAYER HIGHLIGHTS                                                    │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  Source Match: Lagos Regional - Match 1 (Red vs Blue, U16)                  │
│                                                                              │
│  Bulk Upload                                                                 │
│  ───────────                                                                 │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │                                                                      │   │
│  │     📁 Drag and drop highlight clips here                           │   │
│  │                                                                      │   │
│  │     (Multiple files supported)                                      │   │
│  │                                                                      │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                              │
│  Uploaded Clips (Link to Players)                                           │
│  ─────────────────────────────────                                           │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │ Filename           │ Duration │ Player        │ Status              │   │
│  ├─────────────────────────────────────────────────────────────────────┤   │
│  │ clip_001.mp4       │ 2:34     │ [James O. ▼]  │ ✅ Linked           │   │
│  │ clip_002.mp4       │ 1:45     │ [Select   ▼]  │ ⚠️ Needs player     │   │
│  │ clip_003.mp4       │ 3:12     │ [David K. ▼]  │ ✅ Linked           │   │
│  │ clip_004.mp4       │ 2:01     │ [Select   ▼]  │ ⚠️ Needs player     │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                              │
│  Player Search: [Type name to search registered players...        ] [Add]   │
│                                                                              │
│  [Save as Draft]                                          [Publish All]     │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 📋 Data Model Extensions (Admin)

### New Tables Required

```sql
-- Talent day events
CREATE TABLE talent_events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    event_date DATE NOT NULL,
    location_name VARCHAR(255),
    location_country VARCHAR(100),
    location_city VARCHAR(100),
    expected_players INTEGER,
    status VARCHAR(50) DEFAULT 'planned', -- planned, active, completed, cancelled
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Full match recordings
CREATE TABLE matches (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_id UUID REFERENCES talent_events(id),
    match_number INTEGER NOT NULL,
    team_a_name VARCHAR(100),
    team_b_name VARCHAR(100),
    age_group VARCHAR(20),
    blob_url VARCHAR(500),
    thumbnail_url VARCHAR(500),
    duration_seconds INTEGER,
    file_size_bytes BIGINT,
    status VARCHAR(50) DEFAULT 'processing', -- processing, ready, archived
    uploaded_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Player highlights (clips extracted from matches)
CREATE TABLE highlights (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    match_id UUID REFERENCES matches(id),
    profile_id UUID REFERENCES profiles(id), -- Player this highlight belongs to
    title VARCHAR(255),
    description TEXT,
    blob_url VARCHAR(500) NOT NULL,
    thumbnail_url VARCHAR(500),
    duration_seconds INTEGER,
    file_size_bytes BIGINT,
    tags VARCHAR(255)[], -- ['dribbling', 'goal', 'assist']
    status VARCHAR(50) DEFAULT 'published', -- draft, published, archived
    uploaded_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Track which players attended which events
CREATE TABLE event_attendance (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_id UUID REFERENCES talent_events(id),
    profile_id UUID REFERENCES profiles(id),
    team VARCHAR(50), -- 'red', 'blue', etc.
    position_played VARCHAR(50),
    notes TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(event_id, profile_id)
);
```

---

## 🔒 Access Control

### Admin Permissions Matrix

```
┌────────────────────────┬─────────┬─────────┬─────────┬─────────┐
│ Action                 │ Super   │ Media   │ Media   │ Support │
│                        │ Admin   │ Manager │ Editor  │         │
├────────────────────────┼─────────┼─────────┼─────────┼─────────┤
│ Create events          │ ✅      │ ✅      │ ❌      │ ❌      │
│ Upload matches         │ ✅      │ ✅      │ ✅      │ ❌      │
│ Upload highlights      │ ✅      │ ✅      │ ✅      │ ❌      │
│ Link to players        │ ✅      │ ✅      │ ✅      │ ❌      │
│ Delete content         │ ✅      │ ✅      │ ❌      │ ❌      │
│ View analytics         │ ✅      │ ✅      │ ✅      │ ✅      │
│ Manage users           │ ✅      │ ❌      │ ❌      │ ❌      │
│ View financials        │ ✅      │ ❌      │ ❌      │ ❌      │
│ Manage subscriptions   │ ✅      │ ❌      │ ❌      │ ✅      │
└────────────────────────┴─────────┴─────────┴─────────┴─────────┘
```

---

## 📡 Admin API Endpoints

### Events

```
POST   /api/v1/admin/events              -- Create talent day event
GET    /api/v1/admin/events              -- List all events
GET    /api/v1/admin/events/:id          -- Get event details
PUT    /api/v1/admin/events/:id          -- Update event
DELETE /api/v1/admin/events/:id          -- Cancel event
```

### Matches

```
POST   /api/v1/admin/matches             -- Create match record
POST   /api/v1/admin/matches/:id/upload  -- Upload match video
GET    /api/v1/admin/matches             -- List matches
GET    /api/v1/admin/matches/:id         -- Get match details
```

### Highlights

```
POST   /api/v1/admin/highlights          -- Create highlight
POST   /api/v1/admin/highlights/bulk     -- Bulk upload highlights
PUT    /api/v1/admin/highlights/:id      -- Update highlight
PUT    /api/v1/admin/highlights/:id/link -- Link to player profile
DELETE /api/v1/admin/highlights/:id      -- Delete highlight
```

### Attendance

```
POST   /api/v1/admin/events/:id/attendance     -- Record player attendance
GET    /api/v1/admin/events/:id/attendance     -- Get attendees
DELETE /api/v1/admin/events/:id/attendance/:pid -- Remove attendee
```

---

## 📊 Analytics (Admin Dashboard)

### Key Metrics

| Metric | Description | Update Frequency |
|--------|-------------|------------------|
| Events per month | Talent days held | Real-time |
| Players recorded | Unique players with highlights | Daily |
| Highlights created | Total clips published | Real-time |
| Scout engagement | Views, contacts, shortlists | Real-time |
| Subscription revenue | MRR from scouts | Daily |
| Conversion rate | Views → Contact | Weekly |

---

## 🚀 Implementation Priority

### Phase 1 (MVP)

- [ ] Basic event creation
- [ ] Match upload (single file)
- [ ] Highlight upload & player linking
- [ ] Simple dashboard

### Phase 2 (Enhancement)

- [ ] Bulk highlight upload
- [ ] Attendance tracking
- [ ] Analytics dashboard
- [ ] Search & filters

### Phase 3 (Scale)

- [ ] Mobile app for field recording
- [ ] AI-assisted highlight extraction
- [ ] Automated player tagging
- [ ] Multi-region support
