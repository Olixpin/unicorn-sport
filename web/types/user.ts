export interface User {
  id: string
  email: string
  first_name: string
  last_name: string
  role: 'admin' | 'scout' | 'player'
  email_verified: boolean
  created_at: string
  updated_at?: string
}

export interface AuthResponse {
  user: User
  access_token: string
  refresh_token: string
  expires_in: number
}

export interface LoginRequest {
  email: string
  password: string
}

export interface RegisterRequest {
  email: string
  password: string
  first_name: string
  last_name: string
  organization_name?: string
  organization_type?: string
}
