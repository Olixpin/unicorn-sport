# Video Architecture Design

## Overview

This document outlines the video content architecture for Unicorn Sport, designed to monetize full match recordings while using player highlights as free marketing content to attract scouts.

---

## Business Model

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                           CONTENT MONETIZATION                              │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   PLAYER HIGHLIGHTS (FREE) 🆓                                               │
│   ─────────────────────────                                                 │
│   • Purpose: Marketing & Discovery                                          │
│   • Length: 30 seconds - 3 minutes                                          │
│   • Size: 50-500 MB                                                         │
│   • Access: Anyone (including non-registered users)                         │
│   • Goal: Attract scouts, showcase player skills                            │
│                                                                             │
│   FULL MATCH VIDEOS (PAID) 💰                                               │
│   ───────────────────────────                                               │
│   • Purpose: Premium content for serious scouts                             │
│   • Length: 60-90 minutes                                                   │
│   • Size: 5-20 GB                                                           │
│   • Access: Scout+ subscription OR pay-per-view                             │
│   • Goal: Revenue generation, deep player evaluation                        │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Data Model Hierarchy

```
TOURNAMENT
│   Example: "Lagos Youth Cup 2024"
│   Fields: name, year, location, dates, description
│
└── MATCHES (1:many)
        │   Example: "Group A - Academy FC vs United Stars"
        │   Fields: title, match_date, location, teams, stage
        │
        ├── FULL MATCH VIDEO (1:1) ──────────────────────► PAID 💰
        │       Fields: video_url, duration, file_size, status
        │
        └── MATCH PLAYERS (many:many with Player)
                │   Fields: position_played, minutes, goals, assists
                │
                └── PLAYER HIGHLIGHTS (1:many per player per match) ──► FREE 🆓
                        Fields: video_url, highlight_type, duration
                        Types: dribbling, defending, shooting, goal, assist, save, etc.
```

---

## Entity Relationships

```sql
-- A tournament has many matches
Tournament 1 ──── * Match

-- A match has one full video (optional - may not be uploaded yet)
Match 1 ──── 0..1 MatchVideo

-- A match has many players (through match_players join table)
Match * ──── * Player (via MatchPlayer)

-- Each player in a match can have many highlights from that match
MatchPlayer 1 ──── * PlayerHighlight

-- A player can have highlights across multiple matches
Player 1 ──── * PlayerHighlight
```

---

## Revised Schema

### New Tables

```sql
-- Full match videos (PAID content)
CREATE TABLE match_videos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    match_id UUID NOT NULL UNIQUE REFERENCES matches(id) ON DELETE CASCADE,
    
    -- Video storage
    video_url TEXT NOT NULL,           -- S3 URL (private, streamed via presigned URL)
    thumbnail_url TEXT,
    duration_seconds INTEGER,
    file_size_bytes BIGINT,
    
    -- Processing status
    status VARCHAR(20) DEFAULT 'processing',  -- processing, ready, failed
    
    -- Pricing (for pay-per-view option)
    price_cents INTEGER DEFAULT 999,          -- $9.99 default
    
    -- Metadata
    uploaded_by UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Player highlights (FREE content)
CREATE TABLE player_highlights (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    player_id UUID NOT NULL REFERENCES players(id) ON DELETE CASCADE,
    match_id UUID REFERENCES matches(id) ON DELETE SET NULL,  -- Which match this highlight is from
    
    -- Highlight categorization
    highlight_type VARCHAR(30) NOT NULL,  -- dribbling, defending, shooting, goal, assist, save, tackling, heading, speed
    
    -- Video storage
    video_url TEXT NOT NULL,
    thumbnail_url TEXT,
    duration_seconds INTEGER,
    file_size_bytes BIGINT,
    
    -- Metadata
    title TEXT,
    description TEXT,
    timestamp_in_match INTEGER,  -- When in the match this happened (seconds)
    
    -- Status
    status VARCHAR(20) DEFAULT 'approved',  -- approved, pending, rejected
    
    uploaded_by UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    
    -- Index for quick lookups
    CONSTRAINT player_highlights_type_check CHECK (
        highlight_type IN ('dribbling', 'defending', 'shooting', 'passing', 'goal', 'assist', 'save', 'tackling', 'heading', 'speed', 'other')
    )
);

CREATE INDEX idx_highlights_player ON player_highlights(player_id);
CREATE INDEX idx_highlights_match ON player_highlights(match_id);
CREATE INDEX idx_highlights_type ON player_highlights(highlight_type);
```

---

## Admin Upload Workflow

### Flow 1: Tournament-Centric Upload

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                        ADMIN UPLOAD WORKFLOW                                │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   1. NAVIGATE TO TOURNAMENT                                                 │
│      Admin goes to: /admin/tournaments/:id                                  │
│      ↓                                                                      │
│   2. VIEW MATCHES IN TOURNAMENT                                             │
│      See list of matches (or create new match)                              │
│      ↓                                                                      │
│   3. SELECT A MATCH                                                         │
│      Go to: /admin/tournaments/:id/matches/:matchId                         │
│      ↓                                                                      │
│   4. UPLOAD FULL MATCH VIDEO (if not uploaded)                              │
│      • Single video upload                                                  │
│      • Multipart for large files                                            │
│      • Auto-generate thumbnail                                              │
│      ↓                                                                      │
│   5. MANAGE MATCH PLAYERS                                                   │
│      • Add players who participated                                         │
│      • Set position, minutes, goals, assists                                │
│      ↓                                                                      │
│   6. FOR EACH PLAYER - ADD HIGHLIGHTS                                       │
│      • Select player from match roster                                      │
│      • Upload highlight clip                                                │
│      • Tag with highlight type (dribbling, goal, etc.)                      │
│      • Set timestamp in match (optional)                                    │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### UI Layout

```
/admin/tournaments/:id/matches/:matchId
┌─────────────────────────────────────────────────────────────────────────────┐
│  ← Back to Tournament          Match: Academy FC vs United Stars            │
│                                Group Stage - March 15, 2024                 │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │ FULL MATCH VIDEO                                         💰 PAID    │   │
│  ├─────────────────────────────────────────────────────────────────────┤   │
│  │  [Thumbnail]  Match Recording                                       │   │
│  │               Duration: 1:32:45                                     │   │
│  │               Size: 12.4 GB                                         │   │
│  │               Status: ✓ Ready                                       │   │
│  │               [Replace Video] [Delete]                              │   │
│  │                                                                     │   │
│  │  ─── OR if not uploaded ───                                         │   │
│  │                                                                     │   │
│  │  [Upload Full Match Video]                                          │   │
│  │  Drag & drop or click to upload (max 25GB)                         │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │ MATCH PLAYERS                                           [+ Add Player]  │
│  ├─────────────────────────────────────────────────────────────────────┤   │
│  │                                                                     │   │
│  │  ┌─────────────────────────────────────────────────────────────┐   │   │
│  │  │ [Photo] Adebayo Johnson                                      │   │   │
│  │  │         Position: Midfielder | Minutes: 78 | Goals: 1       │   │   │
│  │  │         Highlights: 3 clips                                  │   │   │
│  │  │         [+ Add Highlight] [View Highlights]                  │   │   │
│  │  └─────────────────────────────────────────────────────────────┘   │   │
│  │                                                                     │   │
│  │  ┌─────────────────────────────────────────────────────────────┐   │   │
│  │  │ [Photo] Chinedu Okafor                                       │   │   │
│  │  │         Position: Striker | Minutes: 90 | Goals: 2          │   │   │
│  │  │         Highlights: 5 clips                                  │   │   │
│  │  │         [+ Add Highlight] [View Highlights]                  │   │   │
│  │  └─────────────────────────────────────────────────────────────┘   │   │
│  │                                                                     │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Scout User Journey

### Discovery → Conversion Flow

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                          SCOUT USER JOURNEY                                 │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   1. DISCOVER PLAYERS                                                       │
│      Browse/search: "U-17 midfielders from Lagos"                          │
│      ↓                                                                      │
│   2. VIEW PLAYER PROFILE                                                    │
│      /players/:id                                                           │
│      See: stats, academy, tournaments, highlight thumbnails                │
│      ↓                                                                      │
│   3. WATCH FREE HIGHLIGHTS 🆓                                               │
│      Click any highlight → watch in modal/player                           │
│      Filter by: type (goals, dribbling), tournament                        │
│      ↓                                                                      │
│   "This player looks promising! I want to see more..."                     │
│      ↓                                                                      │
│   4. SEE MATCH APPEARANCES                                                  │
│      Section: "Tournament Appearances"                                      │
│      - Lagos Cup 2024 - Group Stage (2 highlights)                         │
│      - Lagos Cup 2024 - Final (3 highlights)                               │
│      - Abuja Youth League 2024 (4 highlights)                              │
│      Each has: [Watch Full Match 💰]                                       │
│      ↓                                                                      │
│   5. CLICK "WATCH FULL MATCH" ──────────────────► PAYWALL 💰               │
│      ↓                                                                      │
│   ┌─────────────────────────────────────────────────────────────────────┐  │
│   │                     🔒 PREMIUM CONTENT                              │  │
│   │                                                                     │  │
│   │   Watch the full match: "Academy FC vs United Stars"               │  │
│   │   Duration: 1:32:45 | Lagos Cup 2024 - Final                       │  │
│   │                                                                     │  │
│   │   ┌─────────────────────┐  ┌─────────────────────┐                 │  │
│   │   │   PAY PER VIEW      │  │   SUBSCRIBE         │                 │  │
│   │   │   $9.99 one-time    │  │   $29/month         │                 │  │
│   │   │   [Purchase]        │  │   Unlimited access  │                 │  │
│   │   │                     │  │   [Subscribe]       │                 │  │
│   │   └─────────────────────┘  └─────────────────────┘                 │  │
│   └─────────────────────────────────────────────────────────────────────┘  │
│      ↓                                                                      │
│   6. PAYMENT → WATCH FULL MATCH                                            │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Player Profile Page Design

```
/players/:id
┌─────────────────────────────────────────────────────────────────────────────┐
│  ← Back                                                   [Save Player ⭐]  │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ┌──────────┐  ADEBAYO JOHNSON                                              │
│  │          │  Midfielder | 16 years old                                    │
│  │  [Photo] │  Lagos Football Academy                                       │
│  │          │  Lagos, Nigeria                                               │
│  └──────────┘  Preferred Foot: Right | Height: 175cm                       │
│                                                                             │
├─────────────────────────────────────────────────────────────────────────────┤
│  HIGHLIGHTS                                        [Filter: All Types ▼]   │
│  ──────────                                                                 │
│                                                                             │
│  ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐              │
│  │▶ 0:45   │ │▶ 1:23   │ │▶ 0:32   │ │▶ 2:01   │ │▶ 0:55   │              │
│  │         │ │         │ │         │ │         │ │         │              │
│  │[thumb]  │ │[thumb]  │ │[thumb]  │ │[thumb]  │ │[thumb]  │              │
│  ├─────────┤ ├─────────┤ ├─────────┤ ├─────────┤ ├─────────┤              │
│  │⚽ Goal   │ │🦵Dribble│ │🛡️Defend │ │⚽ Goal   │ │🎯Assist │              │
│  │Lagos Cup│ │Lagos Cup│ │Abuja Lg │ │Lagos Cup│ │Abuja Lg │              │
│  └─────────┘ └─────────┘ └─────────┘ └─────────┘ └─────────┘              │
│                                                                             │
├─────────────────────────────────────────────────────────────────────────────┤
│  TOURNAMENT APPEARANCES                                                     │
│  ──────────────────────                                                     │
│                                                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │ 🏆 Lagos Youth Cup 2024                                              │   │
│  │    Matches: 4 | Goals: 3 | Assists: 2 | Highlights: 8               │   │
│  │                                                                     │   │
│  │    • Final vs United Stars - [3 highlights] [Watch Full Match 💰]  │   │
│  │    • Semifinal vs City FC - [2 highlights] [Watch Full Match 💰]   │   │
│  │    • Group A Match 2 - [2 highlights] [Watch Full Match 💰]        │   │
│  │    • Group A Match 1 - [1 highlight] [Watch Full Match 💰]         │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │ 🏆 Abuja Youth League 2024                                           │   │
│  │    Matches: 6 | Goals: 2 | Assists: 4 | Highlights: 12              │   │
│  │    [Expand to see matches]                                          │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Highlight Types

| Type | Icon | Description | Positions |
|------|------|-------------|-----------|
| `goal` | ⚽ | Scoring a goal | All outfield |
| `assist` | 🎯 | Setting up a goal | All outfield |
| `dribbling` | 🦵 | Ball control, beating defenders | Wingers, Midfielders, Forwards |
| `defending` | 🛡️ | Interceptions, blocks, marking | Defenders, Defensive Midfielders |
| `tackling` | 🦶 | Clean tackles | Defenders, Midfielders |
| `passing` | ➡️ | Key passes, vision | Midfielders |
| `shooting` | 🎯 | Shots on goal (not scored) | Forwards, Midfielders |
| `heading` | 🗣️ | Aerial ability | Defenders, Forwards |
| `speed` | ⚡ | Pace, sprinting | Wingers, Fullbacks |
| `save` | 🧤 | Goalkeeper saves | Goalkeepers only |
| `distribution` | 📤 | Goalkeeper distribution | Goalkeepers only |

---

## API Endpoints

### Admin Endpoints

```
# Match Management
POST   /admin/tournaments/:tournamentId/matches           - Create match
GET    /admin/tournaments/:tournamentId/matches           - List matches in tournament
GET    /admin/matches/:id                                  - Get match details
PUT    /admin/matches/:id                                  - Update match
DELETE /admin/matches/:id                                  - Delete match

# Match Video (Full Match - PAID)
POST   /admin/matches/:id/video/upload                    - Init upload for full match
POST   /admin/matches/:id/video                           - Complete/save video
PUT    /admin/matches/:id/video                           - Update video metadata
DELETE /admin/matches/:id/video                           - Delete video

# Match Players
POST   /admin/matches/:id/players                         - Add player to match
PUT    /admin/matches/:id/players/:playerId               - Update player stats
DELETE /admin/matches/:id/players/:playerId               - Remove player from match

# Player Highlights
POST   /admin/matches/:matchId/players/:playerId/highlights    - Add highlight
GET    /admin/matches/:matchId/players/:playerId/highlights    - List player's highlights in match
PUT    /admin/highlights/:id                               - Update highlight
DELETE /admin/highlights/:id                               - Delete highlight
```

### Public/Scout Endpoints

```
# Player Highlights (FREE)
GET    /players/:id/highlights                            - Get player's highlights
GET    /players/:id/highlights?type=goal                  - Filter by type
GET    /players/:id/tournaments                           - Get player's tournament history

# Match Videos (PAID - requires subscription or purchase)
GET    /matches/:id                                       - Get match info (free)
GET    /matches/:id/stream                                - Get stream URL (PAID - checks access)

# Payments
POST   /matches/:id/purchase                              - Purchase single match access
GET    /user/purchases                                    - List purchased matches
```

---

## Access Control

```go
// Who can access what

// Highlights - FREE
func canAccessHighlights(user *User) bool {
    return true  // Anyone, even guests
}

// Full Match Videos - PAID
func canAccessFullMatch(user *User, matchID uuid.UUID) bool {
    // Check 1: Active subscription (Scout+ tier or above)
    if user.Subscription != nil && user.Subscription.CanAccessFullMatch() {
        return true
    }
    
    // Check 2: Purchased this specific match
    if hasPurchasedMatch(user.ID, matchID) {
        return true
    }
    
    return false
}
```

---

## Implementation Priority

### Phase 1: Core Infrastructure ✅ (Mostly Done)
- [x] S3 upload with presigned URLs
- [x] Multipart upload for large files
- [x] Video model with status tracking

### Phase 2: Match-Centric Refactor
- [ ] Create `match_videos` table (1:1 with matches)
- [ ] Create `player_highlights` table (with highlight types)
- [ ] Update Match model to include video relationship
- [ ] Admin UI: Match detail page with video upload

### Phase 3: Admin Upload Flow
- [ ] Match video upload UI
- [ ] Match player management UI
- [ ] Highlight upload per player UI
- [ ] Highlight type tagging

### Phase 4: Scout Experience
- [ ] Player profile with highlights gallery
- [ ] Tournament appearances section
- [ ] Match access paywall
- [ ] Pay-per-view checkout

### Phase 5: Monetization
- [ ] Match purchase model
- [ ] Stripe integration for PPV
- [ ] Purchase history tracking
- [ ] Access verification

---

## File Structure

```
/admin
  /tournaments
    /[id]
      /index.vue           - Tournament detail + matches list
      /matches
        /new.vue           - Create new match
        /[matchId]
          /index.vue       - Match detail + video + players + highlights
          /highlights
            /new.vue       - Add highlight for a player

/players
  /[id]
    /index.vue             - Player profile with highlights
    
/matches
  /[id]
    /index.vue             - Public match page (with paywall)
```
