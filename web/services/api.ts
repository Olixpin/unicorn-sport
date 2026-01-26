/**
 * API Service Layer
 * 
 * Centralized API calls with proper error handling, 
 * type safety, and consistent response structure.
 * 
 * This follows the Repository Pattern for frontend API calls.
 */

import type { ApiResponse, Player, PaginatedResponse } from '~/types/index'
import type { Academy } from '~/schemas/academy'

// Get runtime config and auth store
function getApiClient() {
  const config = useRuntimeConfig()
  const authStore = useAuthStore()
  
  return {
    baseURL: config.public.apiBase as string,
    headers: {
      Authorization: authStore.accessToken ? `Bearer ${authStore.accessToken}` : '',
    },
  }
}

/**
 * Players API Service
 */
export const playersApi = {
  // LIST - Get all players with pagination and filters
  async list(params?: {
    page?: number
    per_page?: number
    q?: string
    position?: string
    country?: string
    min_age?: number
    max_age?: number
  }): Promise<PaginatedResponse<Player>> {
    const { baseURL, headers } = getApiClient()
    const query = new URLSearchParams()
    
    if (params?.page) query.append('page', params.page.toString())
    if (params?.per_page) query.append('per_page', params.per_page.toString())
    if (params?.q) query.append('q', params.q)
    if (params?.position) query.append('position', params.position)
    if (params?.country) query.append('country', params.country)
    if (params?.min_age) query.append('min_age', params.min_age.toString())
    if (params?.max_age) query.append('max_age', params.max_age.toString())

    return $fetch(`/players?${query.toString()}`, { baseURL, headers })
  },

  // READ - Get single player by ID
  async get(id: string): Promise<ApiResponse<{ player: Player }>> {
    const { baseURL, headers } = getApiClient()
    return $fetch(`/players/${id}`, { baseURL, headers })
  },

  // CREATE - Create new player (admin only)
  async create(data: Partial<Player>): Promise<ApiResponse<{ player: Player }>> {
    const { baseURL, headers } = getApiClient()
    return $fetch('/admin/players', {
      baseURL,
      headers,
      method: 'POST',
      body: data,
    })
  },

  // UPDATE - Update player (admin only)
  async update(id: string, data: Partial<Player>): Promise<ApiResponse<{ player: Player }>> {
    const { baseURL, headers } = getApiClient()
    return $fetch(`/admin/players/${id}`, {
      baseURL,
      headers,
      method: 'PUT',
      body: data,
    })
  },

  // DELETE - Delete player (admin only)
  async delete(id: string): Promise<ApiResponse<null>> {
    const { baseURL, headers } = getApiClient()
    return $fetch(`/admin/players/${id}`, {
      baseURL,
      headers,
      method: 'DELETE',
    })
  },

  // VERIFY - Toggle player verification (admin only)
  async verify(id: string, verified: boolean): Promise<ApiResponse<{ player: Player }>> {
    const { baseURL, headers } = getApiClient()
    return $fetch(`/admin/players/${id}/verify`, {
      baseURL,
      headers,
      method: 'POST',
      body: { is_verified: verified },
    })
  },
}

/**
 * Academies API Service
 */
export const academiesApi = {
  // LIST
  async list(params?: {
    page?: number
    per_page?: number
    q?: string
    country?: string
  }): Promise<ApiResponse<{ academies: Academy[]; total: number }>> {
    const { baseURL, headers } = getApiClient()
    const query = new URLSearchParams()
    
    if (params?.page) query.append('page', params.page.toString())
    if (params?.per_page) query.append('per_page', params.per_page.toString())
    if (params?.q) query.append('q', params.q)
    if (params?.country) query.append('country', params.country)

    return $fetch(`/admin/academies?${query.toString()}`, { baseURL, headers })
  },

  // READ
  async get(id: string): Promise<ApiResponse<{ academy: Academy }>> {
    const { baseURL, headers } = getApiClient()
    return $fetch(`/admin/academies/${id}`, { baseURL, headers })
  },

  // CREATE
  async create(data: Partial<Academy>): Promise<ApiResponse<{ academy: Academy }>> {
    const { baseURL, headers } = getApiClient()
    return $fetch('/admin/academies', {
      baseURL,
      headers,
      method: 'POST',
      body: data,
    })
  },

  // UPDATE
  async update(id: string, data: Partial<Academy>): Promise<ApiResponse<{ academy: Academy }>> {
    const { baseURL, headers } = getApiClient()
    return $fetch(`/admin/academies/${id}`, {
      baseURL,
      headers,
      method: 'PUT',
      body: data,
    })
  },

  // DELETE
  async delete(id: string): Promise<ApiResponse<null>> {
    const { baseURL, headers } = getApiClient()
    return $fetch(`/admin/academies/${id}`, {
      baseURL,
      headers,
      method: 'DELETE',
    })
  },
}

/**
 * Auth API Service
 */
export const authApi = {
  async login(email: string, password: string): Promise<ApiResponse<{
    user: { id: string; email: string; role: string }
    access_token: string
    refresh_token: string
  }>> {
    const config = useRuntimeConfig()
    return $fetch('/auth/login', {
      baseURL: config.public.apiBase as string,
      method: 'POST',
      body: { email, password },
    })
  },

  async register(data: {
    email: string
    password: string
    first_name: string
    last_name: string
  }): Promise<ApiResponse<{ user: { id: string; email: string } }>> {
    const config = useRuntimeConfig()
    return $fetch('/auth/register', {
      baseURL: config.public.apiBase as string,
      method: 'POST',
      body: data,
    })
  },

  async forgotPassword(email: string): Promise<ApiResponse<null>> {
    const config = useRuntimeConfig()
    return $fetch('/auth/forgot-password', {
      baseURL: config.public.apiBase as string,
      method: 'POST',
      body: { email },
    })
  },

  async refreshToken(refreshToken: string): Promise<ApiResponse<{
    access_token: string
    refresh_token: string
  }>> {
    const config = useRuntimeConfig()
    return $fetch('/auth/refresh', {
      baseURL: config.public.apiBase as string,
      method: 'POST',
      body: { refresh_token: refreshToken },
    })
  },

  async logout(): Promise<ApiResponse<null>> {
    const { baseURL, headers } = getApiClient()
    return $fetch('/auth/logout', {
      baseURL,
      headers,
      method: 'POST',
    })
  },
}

/**
 * Subscriptions API Service
 */
export const subscriptionsApi = {
  async getPlans(): Promise<ApiResponse<{ plans: Array<{
    id: string
    name: string
    price: number
    features: string[]
  }> }>> {
    const config = useRuntimeConfig()
    return $fetch('/subscriptions/plans', {
      baseURL: config.public.apiBase as string,
    })
  },

  async subscribe(planId: string): Promise<ApiResponse<{ checkout_url: string }>> {
    const { baseURL, headers } = getApiClient()
    return $fetch('/subscriptions/subscribe', {
      baseURL,
      headers,
      method: 'POST',
      body: { plan_id: planId },
    })
  },

  async cancel(): Promise<ApiResponse<null>> {
    const { baseURL, headers } = getApiClient()
    return $fetch('/subscriptions/cancel', {
      baseURL,
      headers,
      method: 'POST',
    })
  },
}

/**
 * Saved Players API Service
 */
export const savedPlayersApi = {
  async list(): Promise<ApiResponse<{ players: Player[] }>> {
    const { baseURL, headers } = getApiClient()
    return $fetch('/users/me/saved-players', { baseURL, headers })
  },

  async save(playerId: string): Promise<ApiResponse<null>> {
    const { baseURL, headers } = getApiClient()
    return $fetch('/users/me/saved-players', {
      baseURL,
      headers,
      method: 'POST',
      body: { player_id: playerId },
    })
  },

  async remove(playerId: string): Promise<ApiResponse<null>> {
    const { baseURL, headers } = getApiClient()
    return $fetch(`/users/me/saved-players/${playerId}`, {
      baseURL,
      headers,
      method: 'DELETE',
    })
  },
}

/**
 * Contact Requests API Service
 */
export const contactsApi = {
  async list(): Promise<ApiResponse<{ contacts: Array<{
    id: string
    player_id: string
    player_name: string
    status: string
    created_at: string
  }> }>> {
    const { baseURL, headers } = getApiClient()
    return $fetch('/users/me/contacts', { baseURL, headers })
  },

  async request(playerId: string, message?: string): Promise<ApiResponse<null>> {
    const { baseURL, headers } = getApiClient()
    return $fetch('/contacts', {
      baseURL,
      headers,
      method: 'POST',
      body: { player_id: playerId, message },
    })
  },
}
