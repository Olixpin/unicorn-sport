import { useAuthStore } from '~/stores/auth'

export default defineNuxtRouteMiddleware(() => {
  const authStore = useAuthStore()
  
  // Check if user is authenticated and is an admin
  if (!authStore.isAuthenticated) {
    return navigateTo('/auth/login')
  }
  
  // Check for admin role
  if (authStore.user?.role !== 'admin') {
    return navigateTo('/dashboard')
  }
})
