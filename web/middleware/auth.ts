import { useAuthStore } from '~/stores/auth'
import type { RouteLocationNormalized } from 'vue-router'

export default defineNuxtRouteMiddleware(async (to: RouteLocationNormalized) => {
  const authStore = useAuthStore()
  
  // Initialize auth from cookies/storage
  if (!authStore.initialized) {
    await authStore.initFromStorage()
  }
  
  // Check if user is authenticated
  if (!authStore.isAuthenticated) {
    // Save the intended destination
    const redirectPath = to.fullPath
    
    return navigateTo({
      path: '/auth/login',
      query: redirectPath !== '/dashboard' ? { redirect: redirectPath } : undefined,
    })
  }
})
