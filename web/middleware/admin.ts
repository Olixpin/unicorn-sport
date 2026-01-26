import { useAuthStore } from '~/stores/auth'

export default defineNuxtRouteMiddleware(async () => {
  const authStore = useAuthStore()
  
  // Initialize auth from cookies/storage
  if (!authStore.initialized) {
    await authStore.initFromStorage()
  }
  
  // Check if user is authenticated and is an admin
  if (!authStore.isAuthenticated) {
    return navigateTo('/auth/login')
  }
  
  // Check for admin role
  if (authStore.user?.role !== 'admin') {
    return navigateTo('/dashboard')
  }
})
