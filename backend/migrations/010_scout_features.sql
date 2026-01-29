-- Migration 010: Enhanced Scout Features
-- Adds: Tags for saved players, contact request tracking, player stats aggregation

-- 1. Add tags to saved_players
ALTER TABLE saved_players ADD COLUMN IF NOT EXISTS tags TEXT[];
ALTER TABLE saved_players ADD COLUMN IF NOT EXISTS priority VARCHAR(20) DEFAULT 'medium';
ALTER TABLE saved_players ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP DEFAULT NOW();

COMMENT ON COLUMN saved_players.tags IS 'Array of scout-defined tags like "U18 CB", "Priority", etc.';
COMMENT ON COLUMN saved_players.priority IS 'Priority level: high, medium, low';

-- Create index for tag searches
CREATE INDEX IF NOT EXISTS idx_saved_players_tags ON saved_players USING GIN(tags);

-- 2. Enhance contact_requests with full tracking
ALTER TABLE contact_requests ADD COLUMN IF NOT EXISTS academy_response TEXT;
ALTER TABLE contact_requests ADD COLUMN IF NOT EXISTS academy_responded_at TIMESTAMP;
ALTER TABLE contact_requests ADD COLUMN IF NOT EXISTS scout_read_at TIMESTAMP;
ALTER TABLE contact_requests ADD COLUMN IF NOT EXISTS follow_up_reminder_at TIMESTAMP;
ALTER TABLE contact_requests ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP DEFAULT NOW();

-- Contact request statuses: pending, sent_to_academy, academy_responded, contact_shared, expired, cancelled

COMMENT ON COLUMN contact_requests.academy_response IS 'Response from academy/player representative';
COMMENT ON COLUMN contact_requests.academy_responded_at IS 'When academy responded';
COMMENT ON COLUMN contact_requests.scout_read_at IS 'When scout read the response';
COMMENT ON COLUMN contact_requests.follow_up_reminder_at IS 'When to remind scout about this request';

-- 3. Create player_stats table for aggregated performance data
CREATE TABLE IF NOT EXISTS player_stats (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    player_id UUID NOT NULL REFERENCES players(id) ON DELETE CASCADE,
    season VARCHAR(20) NOT NULL, -- e.g., "2025-26" or "2026"
    
    -- Match stats
    matches_played INTEGER DEFAULT 0,
    matches_started INTEGER DEFAULT 0,
    minutes_played INTEGER DEFAULT 0,
    
    -- Goals/Assists
    goals INTEGER DEFAULT 0,
    assists INTEGER DEFAULT 0,
    
    -- Discipline (future expansion)
    yellow_cards INTEGER DEFAULT 0,
    red_cards INTEGER DEFAULT 0,
    
    -- Tournament/Competition context
    tournament_id UUID REFERENCES tournaments(id) ON DELETE SET NULL,
    competition_name VARCHAR(255),
    
    -- Meta
    last_calculated_at TIMESTAMP DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    
    UNIQUE(player_id, season, tournament_id)
);

CREATE INDEX IF NOT EXISTS idx_player_stats_player_id ON player_stats(player_id);
CREATE INDEX IF NOT EXISTS idx_player_stats_season ON player_stats(season);

COMMENT ON TABLE player_stats IS 'Aggregated player performance statistics per season/tournament';

-- 4. Create scout_tags table for reusable tags
CREATE TABLE IF NOT EXISTS scout_tags (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(50) NOT NULL,
    color VARCHAR(7) DEFAULT '#6366f1', -- Hex color
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(user_id, name)
);

CREATE INDEX IF NOT EXISTS idx_scout_tags_user_id ON scout_tags(user_id);

COMMENT ON TABLE scout_tags IS 'Scouts custom tags for organizing saved players';

-- 5. Add player_id index to tournaments for "players from this tournament"
-- (Already exists but ensuring it's optimized)

-- 6. Create tournament_standings for public viewing (future)
CREATE TABLE IF NOT EXISTS tournament_standings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tournament_id UUID NOT NULL REFERENCES tournaments(id) ON DELETE CASCADE,
    team_name VARCHAR(255) NOT NULL,
    group_name VARCHAR(50),
    played INTEGER DEFAULT 0,
    won INTEGER DEFAULT 0,
    drawn INTEGER DEFAULT 0,
    lost INTEGER DEFAULT 0,
    goals_for INTEGER DEFAULT 0,
    goals_against INTEGER DEFAULT 0,
    points INTEGER DEFAULT 0,
    position INTEGER,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_tournament_standings_tournament ON tournament_standings(tournament_id);

-- 7. Add public visibility flag to tournaments
ALTER TABLE tournaments ADD COLUMN IF NOT EXISTS is_public BOOLEAN DEFAULT FALSE;
ALTER TABLE tournaments ADD COLUMN IF NOT EXISTS cover_image_url TEXT;
ALTER TABLE tournaments ADD COLUMN IF NOT EXISTS featured BOOLEAN DEFAULT FALSE;

COMMENT ON COLUMN tournaments.is_public IS 'Whether this tournament is publicly browsable';
COMMENT ON COLUMN tournaments.featured IS 'Whether to feature on public tournaments page';
