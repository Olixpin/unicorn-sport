import { useAuthStore } from '~/stores/auth'

/**
 * Redirect authenticated users away from marketing pages
 * to the main app experience (Discover page)
 */
export default defineNuxtRouteMiddleware(async () => {
  const authStore = useAuthStore()
  
  // Initialize auth from cookies/storage if not already done
  if (!authStore.initialized) {
    await authStore.initFromStorage()
  }
  
  // If user is authenticated, redirect to discover page
  if (authStore.isAuthenticated) {
    return navigateTo('/discover')
  }
})
