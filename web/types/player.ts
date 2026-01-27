export interface Player {
  id: string
  user_id?: string
  first_name: string
  last_name: string
  date_of_birth: string
  position: string
  preferred_foot?: 'left' | 'right' | 'both'
  height_cm?: number
  weight_kg?: number
  country: string
  state?: string
  city?: string
  school_name?: string
  verification_status: 'pending' | 'verified' | 'rejected'
  is_verified?: boolean
  profile_photo_url?: string
  thumbnail_url?: string
  tournament_id?: string
  tournament_year?: number
  academy_id?: string
  academy_name?: string
  video_count?: number
  created_at: string
  updated_at: string
  tournament?: Tournament
  academy?: Academy
  videos?: Video[]
}

export interface Academy {
  id: string
  name: string
  country: string
  city?: string
  is_verified: boolean
}

export interface PlayerFilters {
  position?: string
  country?: string
  age_min?: number
  age_max?: number
  verified_only?: boolean
  tournament_id?: string
  page?: number
  limit?: number
}

export interface Tournament {
  id: string
  name: string
  year: number
  location?: string
  country?: string
  start_date?: string
  end_date?: string
  description?: string
  status: 'upcoming' | 'ongoing' | 'completed'
  created_at: string
}

export interface Video {
  id: string
  video_type: 'highlight' | 'full_match'
  title: string
  description?: string
  thumbnail_url?: string
  duration_seconds?: number
  file_size_bytes?: number
  tournament_id?: string
  created_at: string
  tournament?: Tournament
}

export interface SavedPlayer {
  id: string
  user_id: string
  player_id: string
  notes?: string
  created_at: string
  player?: Player
}

export interface ContactRequest {
  id: string
  user_id: string
  player_id: string
  message: string
  status: 'pending' | 'forwarded' | 'responded'
  created_at: string
  player?: Player
}
