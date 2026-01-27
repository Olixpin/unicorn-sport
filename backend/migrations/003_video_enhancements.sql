-- Video Management Enhancements
-- Adds status workflow, review tracking, and match recordings

-- Add status and review fields to videos table
ALTER TABLE videos ADD COLUMN IF NOT EXISTS status VARCHAR(50) DEFAULT 'approved';
ALTER TABLE videos ADD COLUMN IF NOT EXISTS review_notes TEXT;
ALTER TABLE videos ADD COLUMN IF NOT EXISTS reviewed_by UUID REFERENCES users(id);
ALTER TABLE videos ADD COLUMN IF NOT EXISTS reviewed_at TIMESTAMP;
ALTER TABLE videos ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP DEFAULT NOW();

-- Add index for status filtering
CREATE INDEX IF NOT EXISTS idx_videos_status ON videos(status);
CREATE INDEX IF NOT EXISTS idx_videos_uploaded_by ON videos(uploaded_by);
CREATE INDEX IF NOT EXISTS idx_videos_created_at ON videos(created_at DESC);

-- Update existing videos to approved status (admin-uploaded = auto-approved)
UPDATE videos SET status = 'approved' WHERE status IS NULL;

-- Matches table for full tournament recordings
CREATE TABLE IF NOT EXISTS matches (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    match_date DATE NOT NULL,
    location VARCHAR(255),
    event_name VARCHAR(255),  -- e.g., "Talent Day March 2026"
    
    -- Video info (stored in S3)
    video_url TEXT NOT NULL,
    thumbnail_url TEXT,
    duration_seconds INTEGER,
    file_size_bytes BIGINT,
    
    -- Status
    status VARCHAR(50) DEFAULT 'active',  -- active, archived, processing
    
    -- Associations
    tournament_id UUID REFERENCES tournaments(id) ON DELETE SET NULL,
    academy_id UUID REFERENCES academies(id) ON DELETE SET NULL,  -- Which academy's match
    
    -- Metadata
    uploaded_by UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_matches_match_date ON matches(match_date DESC);
CREATE INDEX IF NOT EXISTS idx_matches_tournament_id ON matches(tournament_id);
CREATE INDEX IF NOT EXISTS idx_matches_academy_id ON matches(academy_id);
CREATE INDEX IF NOT EXISTS idx_matches_status ON matches(status);
CREATE INDEX IF NOT EXISTS idx_matches_event_name ON matches(event_name);

-- Match-Player relationship (players who participated in a match)
CREATE TABLE IF NOT EXISTS match_players (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    match_id UUID NOT NULL REFERENCES matches(id) ON DELETE CASCADE,
    player_id UUID NOT NULL REFERENCES players(id) ON DELETE CASCADE,
    
    -- Performance data (optional)
    position_played VARCHAR(50),
    minutes_played INTEGER,
    goals INTEGER DEFAULT 0,
    assists INTEGER DEFAULT 0,
    notes TEXT,
    
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    
    UNIQUE(match_id, player_id)
);

CREATE INDEX IF NOT EXISTS idx_match_players_match_id ON match_players(match_id);
CREATE INDEX IF NOT EXISTS idx_match_players_player_id ON match_players(player_id);

-- Upload sessions for tracking multipart uploads
CREATE TABLE IF NOT EXISTS upload_sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- What's being uploaded
    upload_type VARCHAR(50) NOT NULL,  -- video, match, thumbnail, document
    content_type VARCHAR(100) NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    file_size BIGINT NOT NULL,
    
    -- S3 multipart upload info
    s3_upload_id VARCHAR(255),
    s3_key TEXT NOT NULL,
    
    -- Status
    status VARCHAR(50) DEFAULT 'pending',  -- pending, uploading, completed, failed, expired
    parts_uploaded INTEGER DEFAULT 0,
    total_parts INTEGER,
    
    -- Reference to the entity being created
    entity_type VARCHAR(50),  -- video, match
    entity_id UUID,
    
    -- Metadata
    uploaded_by UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMP NOT NULL,
    completed_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_upload_sessions_status ON upload_sessions(status);
CREATE INDEX IF NOT EXISTS idx_upload_sessions_uploaded_by ON upload_sessions(uploaded_by);
CREATE INDEX IF NOT EXISTS idx_upload_sessions_expires_at ON upload_sessions(expires_at);
