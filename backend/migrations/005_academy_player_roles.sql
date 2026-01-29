-- Migration: Add secondary positions and academy player roles
-- This migration adds support for:
-- 1. Players having multiple positions (natural + secondary)
-- 2. Academy-specific squad roles (where a player's role at academy may differ from natural position)

-- Add secondary_positions to players table
ALTER TABLE players ADD COLUMN IF NOT EXISTS secondary_positions TEXT;
COMMENT ON COLUMN players.secondary_positions IS 'JSON array of secondary positions the player can play';
COMMENT ON COLUMN players.position IS 'Primary/natural position of the player';

-- Create academy_players junction table for proper player-academy relationships
-- This replaces the simple academy_id FK with a rich relationship
CREATE TABLE IF NOT EXISTS academy_players (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    academy_id UUID NOT NULL REFERENCES academies(id) ON DELETE CASCADE,
    player_id UUID NOT NULL REFERENCES players(id) ON DELETE CASCADE,
    squad_role VARCHAR(100),  -- Position as used by this academy (may differ from player.position)
    jersey_number INTEGER,
    squad_status VARCHAR(50) DEFAULT 'active' CHECK (squad_status IN ('active', 'injured', 'loan', 'departed')),
    joined_at TIMESTAMPTZ,
    departed_at TIMESTAMPTZ,
    notes TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    
    -- Prevent duplicate memberships
    UNIQUE(academy_id, player_id)
);

-- Create indexes for performance
CREATE INDEX IF NOT EXISTS idx_academy_players_academy ON academy_players(academy_id);
CREATE INDEX IF NOT EXISTS idx_academy_players_player ON academy_players(player_id);
CREATE INDEX IF NOT EXISTS idx_academy_players_status ON academy_players(squad_status);

-- Migrate existing player-academy relationships to the new table
-- This preserves the current academy_id relationships
INSERT INTO academy_players (academy_id, player_id, squad_role, squad_status, created_at)
SELECT 
    academy_id,
    id as player_id,
    position as squad_role,  -- Initially set squad_role to player's current position
    'active' as squad_status,
    created_at
FROM players 
WHERE academy_id IS NOT NULL
ON CONFLICT (academy_id, player_id) DO NOTHING;

-- Add comments for documentation
COMMENT ON TABLE academy_players IS 'Tracks player memberships in academies with their squad-specific roles';
COMMENT ON COLUMN academy_players.squad_role IS 'Position as assigned by the academy - may differ from player natural position';
COMMENT ON COLUMN academy_players.jersey_number IS 'Squad number assigned by this academy';
COMMENT ON COLUMN academy_players.squad_status IS 'Current status: active, injured, loan, departed';
