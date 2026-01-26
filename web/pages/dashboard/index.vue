<template>
  <div>
    <!-- Welcome Header -->
    <div class="mb-8">
      <div class="flex items-center gap-4">
        <div class="w-14 h-14 rounded-2xl bg-gradient-to-br from-primary-500 to-primary-600 flex items-center justify-center shadow-lg shadow-primary-500/25">
          <span class="text-2xl font-bold text-white">{{ userInitial }}</span>
        </div>
        <div>
          <h1 class="font-display text-2xl lg:text-3xl font-bold text-neutral-900">
            Welcome back{{ userName ? ', ' + userName : '' }}!
          </h1>
          <p class="text-neutral-500">Here's what's happening with your scouting activity.</p>
        </div>
      </div>
    </div>

    <!-- Quick Stats -->
    <div class="grid grid-cols-1 sm:grid-cols-3 gap-5 mb-8">
      <div class="bg-white rounded-2xl p-6 border border-neutral-200 shadow-sm hover:shadow-md hover:border-neutral-300 transition-all duration-200">
        <div class="flex items-center">
          <div class="w-12 h-12 rounded-xl bg-gradient-to-br from-pink-500 to-rose-500 flex items-center justify-center shadow-lg shadow-rose-500/25">
            <svg class="w-6 h-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
            </svg>
          </div>
          <div class="ml-4 flex-1">
            <p class="text-sm text-neutral-500">Saved Players</p>
            <div class="flex items-baseline gap-2">
              <p class="text-3xl font-bold text-neutral-900">{{ savedCount }}</p>
              <span v-if="savedCount > 0" class="text-xs text-primary-600 font-medium">active</span>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-2xl p-6 border border-neutral-200 shadow-sm hover:shadow-md hover:border-neutral-300 transition-all duration-200">
        <div class="flex items-center">
          <div class="w-12 h-12 rounded-xl bg-gradient-to-br from-amber-500 to-orange-500 flex items-center justify-center shadow-lg shadow-orange-500/25">
            <svg class="w-6 h-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
            </svg>
          </div>
          <div class="ml-4 flex-1">
            <p class="text-sm text-neutral-500">Contact Requests</p>
            <div class="flex items-baseline gap-2">
              <p class="text-3xl font-bold text-neutral-900">{{ contactsCount }}</p>
              <span v-if="pendingContacts > 0" class="text-xs text-amber-600 font-medium">{{ pendingContacts }} pending</span>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-2xl p-6 border border-neutral-200 shadow-sm hover:shadow-md hover:border-neutral-300 transition-all duration-200">
        <div class="flex items-center">
          <div class="w-12 h-12 rounded-xl bg-gradient-to-br from-primary-500 to-emerald-500 flex items-center justify-center shadow-lg shadow-emerald-500/25">
            <svg class="w-6 h-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
            </svg>
          </div>
          <div class="ml-4 flex-1">
            <p class="text-sm text-neutral-500">Players Viewed</p>
            <div class="flex items-baseline gap-2">
              <p class="text-3xl font-bold text-neutral-900">{{ viewsCount }}</p>
              <span class="text-xs text-neutral-400">this month</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Recent Activity -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- Recently Viewed -->
      <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
        <div class="flex items-center justify-between p-5 border-b border-neutral-100">
          <h2 class="font-semibold text-neutral-900">Recently Viewed</h2>
          <NuxtLink to="/discover" class="text-sm text-primary-600 hover:text-primary-700 font-medium">
            View all →
          </NuxtLink>
        </div>

        <div v-if="loading" class="p-8 flex justify-center">
          <div class="w-8 h-8 border-4 border-primary-500 border-t-transparent rounded-full animate-spin"></div>
        </div>

        <div v-else-if="recentlyViewed.length === 0" class="p-8 text-center">
          <div class="w-16 h-16 mx-auto bg-neutral-100 rounded-2xl flex items-center justify-center mb-4">
            <svg class="w-8 h-8 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
            </svg>
          </div>
          <p class="text-neutral-600 mb-4">No players viewed yet</p>
          <NuxtLink to="/discover">
            <button class="px-4 py-2 text-sm font-medium text-primary-600 hover:text-primary-700 hover:bg-primary-50 rounded-lg transition-colors">
              Start Browsing
            </button>
          </NuxtLink>
        </div>

        <div v-else class="divide-y divide-neutral-100">
          <div 
            v-for="player in recentlyViewed" 
            :key="player.id" 
            class="flex items-center p-4 hover:bg-neutral-50 transition-colors"
          >
            <div class="w-12 h-12 rounded-xl bg-neutral-200 overflow-hidden flex-shrink-0">
              <img 
                v-if="player.profile_photo_url"
                :src="player.profile_photo_url"
                :alt="player.first_name"
                class="w-full h-full object-cover"
              />
              <div v-else class="w-full h-full flex items-center justify-center bg-gradient-to-br from-neutral-200 to-neutral-300">
                <span class="text-neutral-500 font-medium">{{ player.first_name?.charAt(0) }}</span>
              </div>
            </div>
            <div class="ml-4 flex-1 min-w-0">
              <p class="font-medium text-neutral-900 truncate">{{ player.first_name }} {{ player.last_name }}</p>
              <p class="text-sm text-neutral-500 truncate">{{ player.position }} • {{ player.country }}</p>
            </div>
            <NuxtLink :to="`/players/${player.id}`">
              <button class="px-4 py-2 text-sm font-medium text-neutral-600 hover:text-neutral-900 hover:bg-neutral-100 rounded-lg transition-colors">
                View
              </button>
            </NuxtLink>
          </div>
        </div>
      </div>

      <!-- New Players This Week -->
      <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
        <div class="flex items-center justify-between p-5 border-b border-neutral-100">
          <div class="flex items-center gap-2">
            <h2 class="font-semibold text-neutral-900">New This Week</h2>
            <span class="px-2 py-0.5 bg-primary-100 text-primary-700 text-xs font-medium rounded-full">Fresh</span>
          </div>
          <NuxtLink to="/discover" class="text-sm text-primary-600 hover:text-primary-700 font-medium">
            View all →
          </NuxtLink>
        </div>

        <div v-if="loading" class="p-8 flex justify-center">
          <div class="w-8 h-8 border-4 border-primary-500 border-t-transparent rounded-full animate-spin"></div>
        </div>

        <div v-else-if="newPlayers.length === 0" class="p-8 text-center">
          <div class="w-16 h-16 mx-auto bg-neutral-100 rounded-2xl flex items-center justify-center mb-4">
            <svg class="w-8 h-8 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
            </svg>
          </div>
          <p class="text-neutral-600">No new players this week</p>
        </div>

        <div v-else class="divide-y divide-neutral-100">
          <div 
            v-for="player in newPlayers" 
            :key="player.id" 
            class="flex items-center p-4 hover:bg-neutral-50 transition-colors"
          >
            <div class="w-12 h-12 rounded-xl bg-neutral-200 overflow-hidden flex-shrink-0 ring-2 ring-primary-200">
              <img 
                v-if="player.profile_photo_url"
                :src="player.profile_photo_url"
                :alt="player.first_name"
                class="w-full h-full object-cover"
              />
              <div v-else class="w-full h-full flex items-center justify-center bg-gradient-to-br from-primary-100 to-primary-200">
                <span class="text-primary-600 font-medium">{{ player.first_name?.charAt(0) }}</span>
              </div>
            </div>
            <div class="ml-4 flex-1 min-w-0">
              <p class="font-medium text-neutral-900 truncate">{{ player.first_name }} {{ player.last_name }}</p>
              <p class="text-sm text-neutral-500 truncate">{{ player.position }} • {{ player.country }}</p>
            </div>
            <NuxtLink :to="`/players/${player.id}`">
              <button class="px-4 py-2 text-sm font-medium text-primary-600 hover:text-primary-700 hover:bg-primary-50 rounded-lg transition-colors">
                View
              </button>
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
const config = useRuntimeConfig()

const loading = ref(true)
const savedCount = ref(0)
const contactsCount = ref(0)
const pendingContacts = ref(0)
const viewsCount = ref(0)
const recentlyViewed = ref<Player[]>([])
const newPlayers = ref<Player[]>([])

const userName = computed(() => {
  if (!authStore.user) return ''
  return authStore.user.email.split('@')[0]
})

const userInitial = computed(() => {
  if (!authStore.user) return 'U'
  return authStore.user.email.charAt(0).toUpperCase()
})

onMounted(async () => {
  try {
    const api = useApi()
    
    // Get saved players count
    const savedResponse = await api.get<ApiResponse<{ players: Player[] }>>('/saved-players', {}, true)
    if (savedResponse.success) {
      savedCount.value = savedResponse.data?.players?.length || 0
    }

    // Get contact requests count
    const contactsResponse = await api.get<ApiResponse<{ requests: unknown[] }>>('/contact-requests', {}, true)
    if (contactsResponse.success) {
      contactsCount.value = contactsResponse.data?.requests?.length || 0
    }

    // Get new players
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
