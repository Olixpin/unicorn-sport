-- Add academy_id to players table
-- This migration links players to academies

-- Add academy_id column
ALTER TABLE players ADD COLUMN IF NOT EXISTS academy_id UUID REFERENCES academies(id) ON DELETE SET NULL;

-- Create index for better query performance
CREATE INDEX IF NOT EXISTS idx_players_academy_id ON players(academy_id);

-- Optional: Migrate existing school_name data to academies
-- This is a manual process - admins should review and link players to proper academies
-- The school_name field is kept for legacy/fallback purposes
