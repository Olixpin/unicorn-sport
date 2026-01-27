-- Migration 004: Video Architecture Refactor
-- This migration implements the match-centric video architecture:
-- - Full match videos (PAID content) - one per match
-- - Player highlights (FREE content) - multiple per player per match
-- 
-- Run: psql -d unicorn_sport -f 004_video_architecture_refactor.sql

BEGIN;

-- ============================================================================
-- 1. CREATE MATCH_VIDEOS TABLE (Full Match Recordings - PAID)
-- ============================================================================
-- Each match can have one full video recording
-- This is the PAID content that scouts pay to access

CREATE TABLE IF NOT EXISTS match_videos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    match_id UUID NOT NULL UNIQUE REFERENCES matches(id) ON DELETE CASCADE,
    
    -- Video storage (S3)
    video_url TEXT NOT NULL,                    -- S3 key/URL
    thumbnail_url TEXT,                         -- Thumbnail image URL
    duration_seconds INTEGER,                   -- Video length
    file_size_bytes BIGINT,                     -- File size for display
    
    -- Processing status
    status VARCHAR(20) DEFAULT 'processing' CHECK (status IN ('processing', 'ready', 'failed', 'archived')),
    processing_error TEXT,                      -- Error message if failed
    
    -- Pricing for pay-per-view
    price_cents INTEGER DEFAULT 999,            -- Default $9.99 per match
    currency VARCHAR(3) DEFAULT 'USD',
    
    -- Stats
    view_count INTEGER DEFAULT 0,
    purchase_count INTEGER DEFAULT 0,
    
    -- Metadata
    uploaded_by UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_match_videos_match ON match_videos(match_id);
CREATE INDEX idx_match_videos_status ON match_videos(status);

-- ============================================================================
-- 2. CREATE PLAYER_HIGHLIGHTS TABLE (Highlight Clips - FREE)
-- ============================================================================
-- Short clips showcasing player skills
-- These are FREE content to attract scouts

CREATE TABLE IF NOT EXISTS player_highlights (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    player_id UUID NOT NULL REFERENCES players(id) ON DELETE CASCADE,
    match_id UUID REFERENCES matches(id) ON DELETE SET NULL,  -- Which match this is from
    
    -- Highlight categorization
    highlight_type VARCHAR(30) NOT NULL CHECK (
        highlight_type IN (
            'goal', 'assist', 'dribbling', 'defending', 'tackling',
            'passing', 'shooting', 'heading', 'speed', 'save',
            'distribution', 'positioning', 'vision', 'other'
        )
    ),
    
    -- Video storage (S3)
    video_url TEXT NOT NULL,
    thumbnail_url TEXT,
    duration_seconds INTEGER,
    file_size_bytes BIGINT,
    
    -- Context
    title TEXT,                                 -- Optional title
    description TEXT,                           -- Optional description
    timestamp_in_match INTEGER,                 -- When in the match (seconds from start)
    
    -- Status
    status VARCHAR(20) DEFAULT 'approved' CHECK (status IN ('pending', 'approved', 'rejected', 'archived')),
    
    -- Stats
    view_count INTEGER DEFAULT 0,
    
    -- Metadata
    uploaded_by UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_highlights_player ON player_highlights(player_id);
CREATE INDEX idx_highlights_match ON player_highlights(match_id);
CREATE INDEX idx_highlights_type ON player_highlights(highlight_type);
CREATE INDEX idx_highlights_status ON player_highlights(status);

-- ============================================================================
-- 3. CREATE MATCH_PURCHASES TABLE (Pay-per-view tracking)
-- ============================================================================
-- Track individual match purchases for pay-per-view access

CREATE TABLE IF NOT EXISTS match_purchases (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    match_video_id UUID NOT NULL REFERENCES match_videos(id) ON DELETE CASCADE,
    
    -- Payment info
    amount_cents INTEGER NOT NULL,
    currency VARCHAR(3) DEFAULT 'USD',
    stripe_payment_intent_id TEXT,
    stripe_charge_id TEXT,
    
    -- Status
    status VARCHAR(20) DEFAULT 'completed' CHECK (status IN ('pending', 'completed', 'refunded', 'failed')),
    
    -- Access tracking
    first_viewed_at TIMESTAMP WITH TIME ZONE,
    last_viewed_at TIMESTAMP WITH TIME ZONE,
    view_count INTEGER DEFAULT 0,
    
    -- Metadata
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    UNIQUE(user_id, match_video_id)  -- One purchase per user per match
);

CREATE INDEX idx_purchases_user ON match_purchases(user_id);
CREATE INDEX idx_purchases_video ON match_purchases(match_video_id);

-- ============================================================================
-- 4. UPDATE MATCHES TABLE (Add more fields)
-- ============================================================================
-- Add stage/round info and team names for better context

ALTER TABLE matches ADD COLUMN IF NOT EXISTS stage VARCHAR(50);           -- group, round_of_16, quarterfinal, semifinal, final
ALTER TABLE matches ADD COLUMN IF NOT EXISTS home_team TEXT;              -- Team A name
ALTER TABLE matches ADD COLUMN IF NOT EXISTS away_team TEXT;              -- Team B name
ALTER TABLE matches ADD COLUMN IF NOT EXISTS home_score INTEGER;          -- Final score
ALTER TABLE matches ADD COLUMN IF NOT EXISTS away_score INTEGER;
ALTER TABLE matches ADD COLUMN IF NOT EXISTS match_number INTEGER;        -- Match # in tournament

-- ============================================================================
-- 5. CREATE HIGHLIGHT_VIEWS TABLE (Analytics)
-- ============================================================================
-- Track highlight views for analytics

CREATE TABLE IF NOT EXISTS highlight_views (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    highlight_id UUID NOT NULL REFERENCES player_highlights(id) ON DELETE CASCADE,
    viewer_id UUID REFERENCES users(id) ON DELETE SET NULL,  -- NULL for anonymous
    
    -- Context
    source VARCHAR(30),                         -- player_profile, search, direct
    ip_hash TEXT,                               -- Hashed IP for unique view counting
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_highlight_views_highlight ON highlight_views(highlight_id);
CREATE INDEX idx_highlight_views_viewer ON highlight_views(viewer_id);
CREATE INDEX idx_highlight_views_created ON highlight_views(created_at);

-- ============================================================================
-- 6. CREATE MATCH_VIDEO_VIEWS TABLE (Analytics)
-- ============================================================================
-- Track full match video views

CREATE TABLE IF NOT EXISTS match_video_views (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    match_video_id UUID NOT NULL REFERENCES match_videos(id) ON DELETE CASCADE,
    viewer_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    
    -- Watch session
    watch_duration_seconds INTEGER,             -- How long they watched
    completed BOOLEAN DEFAULT FALSE,            -- Did they finish?
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_match_views_video ON match_video_views(match_video_id);
CREATE INDEX idx_match_views_viewer ON match_video_views(viewer_id);

-- ============================================================================
-- 7. UPDATE TRIGGERS
-- ============================================================================

-- Auto-update updated_at for match_videos
CREATE OR REPLACE FUNCTION update_match_videos_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trigger_match_videos_updated_at ON match_videos;
CREATE TRIGGER trigger_match_videos_updated_at
    BEFORE UPDATE ON match_videos
    FOR EACH ROW
    EXECUTE FUNCTION update_match_videos_updated_at();

-- Auto-update updated_at for player_highlights
CREATE OR REPLACE FUNCTION update_player_highlights_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trigger_player_highlights_updated_at ON player_highlights;
CREATE TRIGGER trigger_player_highlights_updated_at
    BEFORE UPDATE ON player_highlights
    FOR EACH ROW
    EXECUTE FUNCTION update_player_highlights_updated_at();

-- Increment view count on highlight view
CREATE OR REPLACE FUNCTION increment_highlight_view_count()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE player_highlights SET view_count = view_count + 1 WHERE id = NEW.highlight_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trigger_increment_highlight_views ON highlight_views;
CREATE TRIGGER trigger_increment_highlight_views
    AFTER INSERT ON highlight_views
    FOR EACH ROW
    EXECUTE FUNCTION increment_highlight_view_count();

-- Increment view count on match video view
CREATE OR REPLACE FUNCTION increment_match_video_view_count()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE match_videos SET view_count = view_count + 1 WHERE id = NEW.match_video_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trigger_increment_match_video_views ON match_video_views;
CREATE TRIGGER trigger_increment_match_video_views
    AFTER INSERT ON match_video_views
    FOR EACH ROW
    EXECUTE FUNCTION increment_match_video_view_count();

COMMIT;

-- ============================================================================
-- ROLLBACK SCRIPT (if needed)
-- ============================================================================
-- BEGIN;
-- DROP TABLE IF EXISTS match_video_views CASCADE;
-- DROP TABLE IF EXISTS highlight_views CASCADE;
-- DROP TABLE IF EXISTS match_purchases CASCADE;
-- DROP TABLE IF EXISTS player_highlights CASCADE;
-- DROP TABLE IF EXISTS match_videos CASCADE;
-- ALTER TABLE matches DROP COLUMN IF EXISTS stage;
-- ALTER TABLE matches DROP COLUMN IF EXISTS home_team;
-- ALTER TABLE matches DROP COLUMN IF EXISTS away_team;
-- ALTER TABLE matches DROP COLUMN IF EXISTS home_score;
-- ALTER TABLE matches DROP COLUMN IF EXISTS away_score;
-- ALTER TABLE matches DROP COLUMN IF EXISTS match_number;
-- COMMIT;
