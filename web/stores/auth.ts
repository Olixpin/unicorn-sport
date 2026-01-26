import { defineStore } from 'pinia'
import type { User, AuthResponse, ApiResponse } from '~/types/index'

interface AuthState {
  user: User | null
  accessToken: string | null
  refreshToken: string | null
}

export const useAuthStore = defineStore('auth', {
  state: (): AuthState => ({
    user: null,
    accessToken: null,
    refreshToken: null,
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
        this.logout()
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
        this.logout()
      }
    },

    setAuth(data: AuthResponse) {
      this.user = data.user
      this.accessToken = data.access_token
      this.refreshToken = data.refresh_token

      // Persist to localStorage
      if (typeof window !== 'undefined') {
        localStorage.setItem('access_token', data.access_token)
        localStorage.setItem('refresh_token', data.refresh_token)
      }
    },

    logout() {
      this.user = null
      this.accessToken = null
      this.refreshToken = null

      if (typeof window !== 'undefined') {
        localStorage.removeItem('access_token')
        localStorage.removeItem('refresh_token')
      }
    },

    initFromStorage() {
      if (typeof window !== 'undefined') {
        this.accessToken = localStorage.getItem('access_token')
        this.refreshToken = localStorage.getItem('refresh_token')

        if (this.accessToken) {
          this.fetchCurrentUser()
        }
      }
    },
  },
})
