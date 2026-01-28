<template>
  <div class="space-y-4 sm:space-y-6">
    <!-- Hero Welcome + Quick Action - Light Theme -->
    <div class="relative overflow-hidden bg-white rounded-2xl sm:rounded-3xl border border-neutral-200 p-4 sm:p-6 lg:p-8">
      <!-- Subtle Background Pattern -->
      <div class="absolute top-0 right-0 w-64 h-64 bg-gradient-to-br from-primary-100 to-transparent rounded-full blur-3xl -translate-y-1/2 translate-x-1/4 opacity-60"></div>
      <div class="absolute bottom-0 left-0 w-48 h-48 bg-gradient-to-tr from-secondary-100 to-transparent rounded-full blur-2xl translate-y-1/2 -translate-x-1/4 opacity-40"></div>
      
      <div class="relative flex flex-col lg:flex-row lg:items-center lg:justify-between gap-4 sm:gap-6">
        <div class="flex items-center gap-3 sm:gap-5">
          <div class="relative">
            <div class="w-14 h-14 sm:w-16 sm:h-16 rounded-xl sm:rounded-2xl bg-gradient-to-br from-primary-500 to-emerald-500 flex items-center justify-center shadow-xl shadow-primary-500/25">
              <span class="text-xl sm:text-2xl font-bold text-white">{{ userInitial }}</span>
            </div>
            <div class="absolute -bottom-1 -right-1 w-5 h-5 bg-emerald-500 rounded-full border-2 border-white flex items-center justify-center shadow-sm">
              <svg class="w-3 h-3 text-white" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
          <div>
            <p class="text-primary-600 text-xs sm:text-sm font-medium mb-0.5 sm:mb-1">{{ getGreeting() }}</p>
            <h1 class="font-display text-xl sm:text-2xl lg:text-3xl font-bold text-neutral-900">
              {{ userName || 'Scout' }}
            </h1>
            <p class="text-neutral-500 text-sm sm:text-base mt-0.5 sm:mt-1">Ready to discover your next star?</p>
          </div>
        </div>
        
        <div class="flex flex-col sm:flex-row gap-2 sm:gap-3 w-full sm:w-auto">
          <NuxtLink to="/discover" class="flex-1 sm:flex-initial">
            <button class="w-full sm:w-auto group flex items-center justify-center gap-2 px-4 sm:px-5 py-2.5 sm:py-3 bg-gradient-to-r from-primary-500 to-emerald-500 text-white font-semibold text-sm sm:text-base rounded-xl hover:shadow-lg hover:shadow-primary-500/30 hover:scale-[1.02] transition-all duration-200">
              <svg class="w-4 h-4 sm:w-5 sm:h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
              Discover Players
              <svg class="w-4 h-4 transition-transform group-hover:translate-x-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
              </svg>
            </button>
          </NuxtLink>
          <NuxtLink to="/dashboard/saved" class="flex-1 sm:flex-initial">
            <button class="w-full sm:w-auto flex items-center justify-center gap-2 px-4 sm:px-5 py-2.5 sm:py-3 bg-neutral-100 text-neutral-700 font-medium text-sm sm:text-base rounded-xl border border-neutral-200 hover:bg-neutral-200 hover:border-neutral-300 transition-all duration-200">
              <svg class="w-4 h-4 sm:w-5 sm:h-5 text-rose-500" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/>
              </svg>
              My Shortlist
            </button>
          </NuxtLink>
        </div>
      </div>
    </div>

    <!-- Stats Grid - Engagement Focused -->
    <div class="grid grid-cols-2 lg:grid-cols-4 gap-3 sm:gap-4">
      <!-- Saved Players with Progress -->
      <NuxtLink to="/dashboard/saved" class="group">
        <div class="bg-white rounded-xl sm:rounded-2xl p-4 sm:p-5 border border-neutral-200 hover:border-primary-300 hover:shadow-lg hover:shadow-primary-500/10 transition-all duration-300">
          <div class="flex items-start justify-between mb-3 sm:mb-4">
            <div class="w-10 h-10 sm:w-11 sm:h-11 rounded-lg sm:rounded-xl bg-gradient-to-br from-pink-500 to-rose-500 flex items-center justify-center shadow-lg shadow-rose-500/30">
              <svg class="w-4 h-4 sm:w-5 sm:h-5 text-white" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/>
              </svg>
            </div>
            <svg class="w-5 h-5 text-neutral-300 group-hover:text-primary-500 transition-colors" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
            </svg>
          </div>
          <div>
            <p class="text-2xl sm:text-3xl font-bold text-neutral-900">{{ savedCount }}</p>
            <p class="text-xs sm:text-sm text-neutral-500 mt-0.5 sm:mt-1">Saved Players</p>
          </div>
          <div class="mt-2 sm:mt-3 pt-2 sm:pt-3 border-t border-neutral-100">
            <div class="flex items-center gap-1.5">
              <div class="flex-1 h-1 sm:h-1.5 bg-neutral-100 rounded-full overflow-hidden">
                <div class="h-full bg-gradient-to-r from-pink-500 to-rose-500 rounded-full transition-all duration-500" :style="{ width: `${Math.min((savedCount / 50) * 100, 100)}%` }"></div>
              </div>
              <span class="text-[10px] sm:text-xs text-neutral-400">{{ savedCount }}/50</span>
            </div>
          </div>
        </div>
      </NuxtLink>

      <!-- Contact Requests with Status -->
      <NuxtLink to="/dashboard/contacts" class="group">
        <div class="bg-white rounded-xl sm:rounded-2xl p-4 sm:p-5 border border-neutral-200 hover:border-amber-300 hover:shadow-lg hover:shadow-amber-500/10 transition-all duration-300">
          <div class="flex items-start justify-between mb-3 sm:mb-4">
            <div class="relative">
              <div class="w-10 h-10 sm:w-11 sm:h-11 rounded-lg sm:rounded-xl bg-gradient-to-br from-amber-500 to-orange-500 flex items-center justify-center shadow-lg shadow-orange-500/30">
                <svg class="w-4 h-4 sm:w-5 sm:h-5 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                </svg>
              </div>
              <div v-if="pendingContacts > 0" class="absolute -top-1 -right-1 w-4 h-4 sm:w-5 sm:h-5 bg-red-500 rounded-full flex items-center justify-center text-[9px] sm:text-[10px] font-bold text-white ring-2 ring-white">
                {{ pendingContacts }}
              </div>
            </div>
            <svg class="w-4 h-4 sm:w-5 sm:h-5 text-neutral-300 group-hover:text-amber-500 transition-colors" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
            </svg>
          </div>
          <div>
            <p class="text-2xl sm:text-3xl font-bold text-neutral-900">{{ contactsCount }}</p>
            <p class="text-xs sm:text-sm text-neutral-500 mt-0.5 sm:mt-1">Contact Requests</p>
          </div>
          <div class="mt-2 sm:mt-3 pt-2 sm:pt-3 border-t border-neutral-100">
            <span v-if="pendingContacts > 0" class="inline-flex items-center gap-1 text-[10px] sm:text-xs font-medium text-amber-600">
              <span class="w-1.5 h-1.5 bg-amber-500 rounded-full animate-pulse"></span>
              {{ pendingContacts }} pending
            </span>
            <span v-else class="text-[10px] sm:text-xs text-neutral-400">All caught up!</span>
          </div>
        </div>
      </NuxtLink>

      <!-- This Week's Activity -->
      <div class="bg-white rounded-xl sm:rounded-2xl p-4 sm:p-5 border border-neutral-200">
        <div class="flex items-start justify-between mb-3 sm:mb-4">
          <div class="w-10 h-10 sm:w-11 sm:h-11 rounded-lg sm:rounded-xl bg-gradient-to-br from-blue-500 to-indigo-500 flex items-center justify-center shadow-lg shadow-blue-500/30">
            <svg class="w-4 h-4 sm:w-5 sm:h-5 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
          </div>
        </div>
        <div>
          <p class="text-2xl sm:text-3xl font-bold text-neutral-900">{{ viewsCount }}</p>
          <p class="text-xs sm:text-sm text-neutral-500 mt-0.5 sm:mt-1">Profiles Viewed</p>
        </div>
        <div class="mt-2 sm:mt-3 pt-2 sm:pt-3 border-t border-neutral-100">
          <span class="text-[10px] sm:text-xs text-neutral-400">This week</span>
        </div>
      </div>

      <!-- Quick Discover CTA -->
      <NuxtLink to="/discover" class="group">
        <div class="relative h-full bg-gradient-to-br from-primary-500 to-emerald-500 rounded-xl sm:rounded-2xl p-4 sm:p-5 overflow-hidden hover:shadow-lg hover:shadow-primary-500/30 transition-all duration-300 hover:scale-[1.02]">
          <div class="absolute top-0 right-0 w-20 h-20 sm:w-24 sm:h-24 bg-white/10 rounded-full -translate-y-1/2 translate-x-1/2"></div>
          <div class="absolute bottom-0 left-0 w-12 h-12 sm:w-16 sm:h-16 bg-white/10 rounded-full translate-y-1/2 -translate-x-1/2"></div>
          <div class="relative h-full flex flex-col justify-between">
            <div class="w-10 h-10 sm:w-11 sm:h-11 rounded-lg sm:rounded-xl bg-white/20 flex items-center justify-center mb-3 sm:mb-4">
              <svg class="w-4 h-4 sm:w-5 sm:h-5 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
              </svg>
            </div>
            <div>
              <p class="text-white font-bold text-base sm:text-lg">Find Talent</p>
              <p class="text-white/70 text-xs sm:text-sm mt-0.5 sm:mt-1">500+ verified→</p>
            </div>
            <svg class="absolute bottom-4 right-4 sm:bottom-5 sm:right-5 w-5 h-5 sm:w-6 sm:h-6 text-white/50 group-hover:text-white group-hover:translate-x-1 transition-all" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3" />
            </svg>
          </div>
        </div>
      </NuxtLink>
    </div>

    <!-- Content Grid -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-4 sm:gap-6">
      <!-- New Players - Featured Larger -->
      <div class="lg:col-span-2 bg-white rounded-xl sm:rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
        <div class="flex items-center justify-between p-4 sm:p-5 border-b border-neutral-100">
          <div class="flex items-center gap-2 sm:gap-3">
            <div class="w-2 h-2 bg-primary-500 rounded-full animate-pulse"></div>
            <h2 class="font-semibold text-sm sm:text-base text-neutral-900">Fresh Talent This Week</h2>
            <span class="px-2 py-0.5 sm:px-2.5 sm:py-1 bg-primary-100 text-primary-700 text-[10px] sm:text-xs font-bold rounded-full">{{ newPlayers.length }} NEW</span>
          </div>
          <NuxtLink to="/discover" class="text-xs sm:text-sm text-primary-600 hover:text-primary-700 font-medium flex items-center gap-1 group">
            See all
            <svg class="w-3 h-3 sm:w-4 sm:h-4 group-hover:translate-x-0.5 transition-transform" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
            </svg>
          </NuxtLink>
        </div>

        <div v-if="loading" class="p-12 flex justify-center">
          <div class="flex flex-col items-center gap-4">
            <div class="w-12 h-12 border-4 border-primary-500 border-t-transparent rounded-full animate-spin"></div>
            <p class="text-neutral-500 text-sm">Fetching latest players...</p>
          </div>
        </div>

        <div v-else-if="newPlayers.length === 0" class="p-12 text-center">
          <div class="w-20 h-20 mx-auto bg-gradient-to-br from-primary-100 to-primary-50 rounded-3xl flex items-center justify-center mb-5">
            <svg class="w-10 h-10 text-primary-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0z" />
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-neutral-900 mb-2">New talent coming soon</h3>
          <p class="text-neutral-500 mb-6">Check back later for fresh player profiles</p>
          <NuxtLink to="/discover">
            <button class="inline-flex items-center gap-2 px-5 py-2.5 bg-primary-500 text-white font-medium rounded-xl hover:bg-primary-600 transition-colors">
              Browse All Players
            </button>
          </NuxtLink>
        </div>

        <div v-else class="grid grid-cols-1 sm:grid-cols-2 divide-y sm:divide-y-0 sm:divide-x divide-neutral-100">
          <div 
            v-for="(player, index) in newPlayers.slice(0, 4)" 
            :key="player.id" 
            class="p-4 hover:bg-neutral-50 transition-colors group"
            :class="{ 'border-t border-neutral-100': index >= 2 }"
          >
            <NuxtLink :to="`/players/${player.id}`" class="flex items-center gap-4">
              <div class="relative">
                <div class="w-14 h-14 rounded-xl bg-neutral-200 overflow-hidden ring-2 ring-primary-200 group-hover:ring-primary-400 transition-all">
                  <img 
                    v-if="player.profile_photo_url"
                    :src="player.profile_photo_url"
                    :alt="player.first_name"
                    class="w-full h-full object-cover"
                  />
                  <div v-else class="w-full h-full flex items-center justify-center bg-gradient-to-br from-primary-100 to-primary-200">
                    <span class="text-primary-600 font-bold text-lg">{{ player.first_name?.charAt(0) }}</span>
                  </div>
                </div>
                <div v-if="player.verification_status === 'verified'" class="absolute -bottom-1 -right-1 w-5 h-5 bg-primary-500 rounded-full flex items-center justify-center ring-2 ring-white">
                  <svg class="w-3 h-3 text-white" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"/>
                  </svg>
                </div>
              </div>
              <div class="flex-1 min-w-0">
                <p class="font-semibold text-neutral-900 truncate group-hover:text-primary-600 transition-colors">{{ player.first_name }} {{ getPlayerLastInitial(player) }}.</p>
                <p class="text-sm text-neutral-500">{{ player.position }}</p>
                <p class="text-xs text-neutral-400 mt-0.5">{{ player.country }}</p>
              </div>
              <svg class="w-5 h-5 text-neutral-300 group-hover:text-primary-500 transition-colors" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
              </svg>
            </NuxtLink>
          </div>
        </div>
      </div>

      <!-- Activity Feed -->
      <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
        <div class="p-5 border-b border-neutral-100">
          <h2 class="font-semibold text-neutral-900">Your Activity</h2>
        </div>

        <div v-if="loading" class="p-8 flex justify-center">
          <div class="w-8 h-8 border-3 border-primary-500 border-t-transparent rounded-full animate-spin"></div>
        </div>

        <div v-else-if="recentlyViewed.length === 0" class="p-8 text-center">
          <div class="w-14 h-14 mx-auto bg-neutral-100 rounded-2xl flex items-center justify-center mb-4">
            <svg class="w-7 h-7 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <p class="text-neutral-600 font-medium mb-1">No activity yet</p>
          <p class="text-neutral-400 text-sm mb-4">Start exploring players!</p>
          <NuxtLink to="/discover">
            <button class="text-sm text-primary-600 hover:text-primary-700 font-medium">
              Discover Players →
            </button>
          </NuxtLink>
        </div>

        <div v-else class="divide-y divide-neutral-100">
          <div 
            v-for="player in recentlyViewed.slice(0, 5)" 
            :key="player.id" 
            class="p-4 hover:bg-neutral-50 transition-colors group"
          >
            <NuxtLink :to="`/players/${player.id}`" class="flex items-center gap-3">
              <div class="w-10 h-10 rounded-lg bg-neutral-200 overflow-hidden flex-shrink-0">
                <img 
                  v-if="player.profile_photo_url"
                  :src="player.profile_photo_url"
                  :alt="player.first_name"
                  class="w-full h-full object-cover"
                />
                <div v-else class="w-full h-full flex items-center justify-center bg-neutral-200">
                  <span class="text-neutral-500 text-sm font-medium">{{ player.first_name?.charAt(0) }}</span>
                </div>
              </div>
              <div class="flex-1 min-w-0">
                <p class="text-sm font-medium text-neutral-900 truncate group-hover:text-primary-600 transition-colors">{{ player.first_name }} {{ getPlayerLastInitial(player) }}.</p>
                <p class="text-xs text-neutral-400">Viewed recently</p>
              </div>
            </NuxtLink>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Player, ApiResponse } from '~/types'

definePageMeta({
  layout: 'dashboard',
  middleware: 'auth',
})

const authStore = useAuthStore()
const subscriptionStore = useSubscriptionStore()
const config = useRuntimeConfig()
const api = useApi()

const loading = ref(true)
const savedCount = ref(0)
const contactsCount = ref(0)
const pendingContacts = ref(0)
const viewsCount = ref(0)
const recentlyViewed = ref<Player[]>([])
const newPlayers = ref<Player[]>([])

const userName = computed(() => {
  if (!authStore.user) return ''
  return authStore.user.first_name || authStore.user.email.split('@')[0]
})

const userInitial = computed(() => {
  if (!authStore.user) return 'U'
  return (authStore.user.first_name?.charAt(0) || authStore.user.email.charAt(0)).toUpperCase()
})

// Time-based greeting
function getGreeting(): string {
  const hour = new Date().getHours()
  if (hour < 12) return 'Good morning'
  if (hour < 17) return 'Good afternoon'
  return 'Good evening'
}

// Helper to get player's last name initial safely
function getPlayerLastInitial(player: Player): string {
  if (player?.last_name_init) return player.last_name_init
  if (player?.last_name) return player.last_name.charAt(0)
  return ''
}

onMounted(async () => {
  // Fetch subscription status
  await subscriptionStore.fetchSubscription()
  
  try {
    // Get saved players count
    if (subscriptionStore.canSavePlayers) {
      const savedResponse = await api.get<ApiResponse<{ saved_players: Player[]; pagination: { total: number } }>>(
        '/saved-players',
        {},
        true
      )
      if (savedResponse.success && savedResponse.data) {
        savedCount.value = savedResponse.data.pagination?.total || savedResponse.data.saved_players?.length || 0
      }
    }

    // Get contact requests count
    if (subscriptionStore.canContactPlayers) {
      const contactsResponse = await api.get<ApiResponse<{ contact_requests: Array<{ status: string }> }>>(
        '/contact-requests',
        {},
        true
      )
      if (contactsResponse.success && contactsResponse.data?.contact_requests) {
        contactsCount.value = contactsResponse.data.contact_requests.length
        pendingContacts.value = contactsResponse.data.contact_requests.filter(r => r.status === 'pending').length
      }
    }

    // Get new players (public)
    const playersResponse = await $fetch<ApiResponse<{ players: Player[] }>>('/players?limit=5', {
      baseURL: config.public.apiBase,
    })
    if (playersResponse.success) {
      newPlayers.value = playersResponse.data?.players || []
    }
  } catch (error) {
    console.error('Failed to load dashboard data:', error)
  } finally {
    loading.value = false
  }
})
</script>
