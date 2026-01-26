import { defineStore } from 'pinia'
import type { User, AuthResponse, ApiResponse } from '~/types/index'

interface AuthState {
  user: User | null
  accessToken: string | null
  refreshToken: string | null
  initialized: boolean
}

export const useAuthStore = defineStore('auth', {
  state: (): AuthState => ({
    user: null,
    accessToken: null,
    refreshToken: null,
    initialized: false,
  }),

  getters: {
    isAuthenticated: (state) => !!state.accessToken && !!state.user,
    isAdmin: (state) => state.user?.role === 'admin',
    isScout: (state) => state.user?.role === 'scout',
    isPlayer: (state) => state.user?.role === 'player',
  },

  actions: {
    async login(email: string, password: string): Promise<boolean> {
      try {
        const config = useRuntimeConfig()
        const response = await $fetch<ApiResponse<AuthResponse>>('/auth/login', {
          baseURL: config.public.apiBase,
          method: 'POST',
          body: { email, password },
        })

        if (response.success && response.data) {
          this.setAuth(response.data)
          return true
        }
        return false
      } catch (error) {
        console.error('Login failed:', error)
        return false
      }
    },

    async register(data: {
      email: string
      password: string
      first_name: string
      last_name: string
      organization_name?: string
      organization_type?: string
    }): Promise<boolean> {
      try {
        const config = useRuntimeConfig()
        const response = await $fetch<ApiResponse<AuthResponse>>('/auth/register', {
          baseURL: config.public.apiBase,
          method: 'POST',
          body: data,
        })

        if (response.success && response.data) {
          this.setAuth(response.data)
          return true
        }
        return false
      } catch (error) {
        console.error('Registration failed:', error)
        return false
      }
    },

    async refreshAccessToken(): Promise<boolean> {
      if (!this.refreshToken) return false

      try {
        const config = useRuntimeConfig()
        const response = await $fetch<ApiResponse<AuthResponse>>('/auth/refresh', {
          baseURL: config.public.apiBase,
          method: 'POST',
          body: { refresh_token: this.refreshToken },
        })

        if (response.success && response.data) {
          this.setAuth(response.data)
          return true
        }
        return false
      } catch (error) {
        console.error('Token refresh failed:', error)
        this.clearAuth()
        return false
      }
    },

    async fetchCurrentUser(): Promise<void> {
      if (!this.accessToken) return

      try {
        const config = useRuntimeConfig()
        const response = await $fetch<ApiResponse<User>>('/auth/me', {
          baseURL: config.public.apiBase,
          method: 'GET',
          headers: {
            Authorization: `Bearer ${this.accessToken}`,
          },
        })

        if (response.success && response.data) {
          this.user = response.data
        }
      } catch (error) {
        console.error('Fetch user failed:', error)
        this.clearAuth()
      }
    },

    setAuth(data: AuthResponse) {
      this.user = data.user
      this.accessToken = data.access_token
      this.refreshToken = data.refresh_token
      this.initialized = true

      // Persist to cookies (works with SSR)
      const accessTokenCookie = useCookie('access_token', { 
        maxAge: 60 * 60 * 24 * 7, // 7 days
        sameSite: 'lax',
        secure: process.env.NODE_ENV === 'production',
      })
      const refreshTokenCookie = useCookie('refresh_token', { 
        maxAge: 60 * 60 * 24 * 30, // 30 days
        sameSite: 'lax',
        secure: process.env.NODE_ENV === 'production',
      })
      
      accessTokenCookie.value = data.access_token
      refreshTokenCookie.value = data.refresh_token
    },

    // Clear auth state without using useCookie (safe to call anywhere)
    clearAuth() {
      this.user = null
      this.accessToken = null
      this.refreshToken = null
    },

    // Full logout with cookie clearing (must be called in Nuxt context)
    logout() {
      this.clearAuth()

      // Clear cookies - only works in Nuxt context
      try {
        const accessTokenCookie = useCookie('access_token')
        const refreshTokenCookie = useCookie('refresh_token')
        accessTokenCookie.value = null
        refreshTokenCookie.value = null
      } catch {
        // If not in Nuxt context, clear via document.cookie
        if (typeof document !== 'undefined') {
          document.cookie = 'access_token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;'
          document.cookie = 'refresh_token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;'
        }
      }
      
      // Also clear localStorage for backwards compatibility
      if (typeof window !== 'undefined') {
        localStorage.removeItem('access_token')
        localStorage.removeItem('refresh_token')
      }
    },

    async initFromStorage(): Promise<void> {
      if (this.initialized) return
      
      // Try cookies first (works with SSR)
      const accessTokenCookie = useCookie('access_token')
      const refreshTokenCookie = useCookie('refresh_token')
      
      if (accessTokenCookie.value) {
        this.accessToken = accessTokenCookie.value
        this.refreshToken = refreshTokenCookie.value ?? null
        await this.fetchCurrentUser()
        this.initialized = true
        return
      }
      
      // Fallback to localStorage (client-side only, for migration)
      if (typeof window !== 'undefined') {
        const storedAccessToken = localStorage.getItem('access_token')
        const storedRefreshToken = localStorage.getItem('refresh_token')
        
        if (storedAccessToken) {
          this.accessToken = storedAccessToken
          this.refreshToken = storedRefreshToken
          
          // Migrate to cookies
          accessTokenCookie.value = storedAccessToken
          refreshTokenCookie.value = storedRefreshToken
          
          // Clear localStorage after migration
          localStorage.removeItem('access_token')
          localStorage.removeItem('refresh_token')
          
          await this.fetchCurrentUser()
        }
      }
      
      this.initialized = true
    },
  },
})
