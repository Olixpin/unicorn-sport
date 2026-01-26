import type { ApiResponse, ApiError } from '~/types/index'
import { useAuthStore } from '~/stores/auth'

export const useApi = () => {
  const config = useRuntimeConfig()
  const authStore = useAuthStore()

  const baseURL = config.public.apiBase

  const apiFetch = async <T>(
    endpoint: string,
    options: {
      method?: 'GET' | 'POST' | 'PUT' | 'DELETE' | 'PATCH'
      body?: Record<string, unknown>
      query?: Record<string, unknown>
      requiresAuth?: boolean
    } = {}
  ): Promise<T> => {
    const { method = 'GET', body, query, requiresAuth = false } = options

    const headers: Record<string, string> = {
      'Content-Type': 'application/json',
    }

    if (requiresAuth && authStore.accessToken) {
      headers['Authorization'] = `Bearer ${authStore.accessToken}`
    }

    try {
      const response = await $fetch<T>(endpoint, {
        baseURL,
        method,
        headers,
        body: body ? JSON.stringify(body) : undefined,
        query,
      })

      return response
    } catch (error: unknown) {
      // Handle 401 - token expired
      if (error && typeof error === 'object' && 'statusCode' in error && error.statusCode === 401) {
        // Try to refresh token
        const refreshed = await authStore.refreshAccessToken()
        if (refreshed) {
          // Retry the request with new token
          headers['Authorization'] = `Bearer ${authStore.accessToken}`
          return $fetch<T>(endpoint, {
            baseURL,
            method,
            headers,
            body: body ? JSON.stringify(body) : undefined,
            query,
          })
        } else {
          // Refresh failed, logout
          authStore.logout()
          navigateTo('/auth/login')
        }
      }
      throw error
    }
  }

  return {
    // GET request
    get: <T>(endpoint: string, query?: Record<string, unknown>, requiresAuth = false) =>
      apiFetch<T>(endpoint, { method: 'GET', query, requiresAuth }),

    // POST request
    post: <T>(endpoint: string, body?: Record<string, unknown>, requiresAuth = false) =>
      apiFetch<T>(endpoint, { method: 'POST', body, requiresAuth }),

    // PUT request
    put: <T>(endpoint: string, body?: Record<string, unknown>, requiresAuth = false) =>
      apiFetch<T>(endpoint, { method: 'PUT', body, requiresAuth }),

    // DELETE request
    delete: <T>(endpoint: string, requiresAuth = false) =>
      apiFetch<T>(endpoint, { method: 'DELETE', requiresAuth }),

    // PATCH request
    patch: <T>(endpoint: string, body?: Record<string, unknown>, requiresAuth = false) =>
      apiFetch<T>(endpoint, { method: 'PATCH', body, requiresAuth }),
  }
}
