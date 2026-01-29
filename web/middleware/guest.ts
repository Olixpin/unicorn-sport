import { useAuthStore } from '~/stores/auth'

export default defineNuxtRouteMiddleware(() => {
  const authStore = useAuthStore()
  
  // If already authenticated, redirect to discover
  if (authStore.isAuthenticated) {
    return navigateTo('/discover')
  }
})
