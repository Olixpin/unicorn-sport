import { useAuthStore } from '~/stores/auth'
import { useSubscriptionStore } from '~/stores/subscription'
import type { RouteLocationNormalized } from 'vue-router'

export default defineNuxtRouteMiddleware(async (to: RouteLocationNormalized) => {
  const authStore = useAuthStore()
  const subscriptionStore = useSubscriptionStore()
  
  // Initialize auth from cookies/storage
  if (!authStore.initialized) {
    await authStore.initFromStorage()
  }
  
  // Must be authenticated first
  if (!authStore.isAuthenticated) {
    return navigateTo('/auth/login')
  }
  
  // Check route meta for required subscription tier
  const requiredTier = to.meta.requiredTier as string | undefined
  
  if (requiredTier) {
    const tiers = ['free', 'scout', 'pro', 'club']
    const userTierIndex = tiers.indexOf(subscriptionStore.tier)
    const requiredTierIndex = tiers.indexOf(requiredTier)
    
    if (userTierIndex < requiredTierIndex) {
      return navigateTo('/pricing')
    }
  }
})
