import { defineStore } from 'pinia'
import type { Subscription, SubscriptionTier, ApiResponse } from '~/types/index'
import { useAuthStore } from './auth'

interface SubscriptionState {
  subscription: Subscription | null
  loading: boolean
}

export const useSubscriptionStore = defineStore('subscription', {
  state: (): SubscriptionState => ({
    subscription: null,
    loading: false,
  }),

  getters: {
    tier: (state): SubscriptionTier => state.subscription?.tier ?? 'free',
    isActive: (state) => state.subscription?.status === 'active',
    canAccessFullMatches: (state) => {
      const tier = state.subscription?.tier
      return tier === 'scout' || tier === 'pro' || tier === 'club'
    },
    canSavePlayers: (state) => {
      const tier = state.subscription?.tier
      return tier === 'scout' || tier === 'pro' || tier === 'club'
    },
    canRequestContact: (state) => {
      const tier = state.subscription?.tier
      return tier === 'pro' || tier === 'club'
    },
    canContactPlayers: (state) => {
      const tier = state.subscription?.tier
      return tier === 'pro' || tier === 'club'
    },
    hasTier: (state) => {
      return (requiredTier: SubscriptionTier): boolean => {
        const tierOrder: SubscriptionTier[] = ['free', 'scout', 'pro', 'club']
        const currentTierIndex = tierOrder.indexOf(state.subscription?.tier ?? 'free')
        const requiredTierIndex = tierOrder.indexOf(requiredTier)
        return currentTierIndex >= requiredTierIndex
      }
    },
  },

  actions: {
    async fetchSubscription(): Promise<void> {
      const authStore = useAuthStore()
      if (!authStore.isAuthenticated) return

      this.loading = true
      try {
        const config = useRuntimeConfig()
        const response = await $fetch<ApiResponse<Subscription>>('/subscriptions/me', {
          baseURL: config.public.apiBase,
          method: 'GET',
          headers: {
            Authorization: `Bearer ${authStore.accessToken}`,
          },
        })
        if (response.success && response.data) {
          this.subscription = response.data
        }
      } catch (error) {
        console.error('Failed to fetch subscription:', error)
      } finally {
        this.loading = false
      }
    },

    async createCheckoutSession(tier: SubscriptionTier): Promise<string | null> {
      const authStore = useAuthStore()
      try {
        const config = useRuntimeConfig()
        const response = await $fetch<ApiResponse<{ url: string }>>('/subscriptions/checkout', {
          baseURL: config.public.apiBase,
          method: 'POST',
          body: { tier },
          headers: {
            Authorization: `Bearer ${authStore.accessToken}`,
          },
        })
        if (response.success && response.data?.url) {
          return response.data.url
        }
        return null
      } catch (error) {
        console.error('Failed to create checkout session:', error)
        return null
      }
    },

    reset() {
      this.subscription = null
      this.loading = false
    },
  },
})
