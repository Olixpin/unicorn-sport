<template>
  <div>
    <!-- Header -->
    <div class="mb-8">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="font-display text-2xl lg:text-3xl font-bold text-neutral-900">
            Saved Players
          </h1>
          <p class="mt-1 text-neutral-500">Players you've saved for later review</p>
        </div>
        <div v-if="savedPlayers.length > 0" class="hidden sm:flex items-center gap-3">
          <span class="px-3 py-1.5 bg-neutral-100 text-neutral-600 text-sm font-medium rounded-lg">
            {{ savedPlayers.length }} {{ savedPlayers.length === 1 ? 'player' : 'players' }}
          </span>
        </div>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex justify-center py-16">
      <div class="w-10 h-10 border-4 border-primary-500 border-t-transparent rounded-full animate-spin"></div>
    </div>

    <!-- Subscription Gate -->
    <div v-if="!subscriptionStore.canSavePlayers && !loading" class="bg-white rounded-2xl border border-neutral-200 p-12 text-center">
      <div class="w-20 h-20 mx-auto bg-gradient-to-br from-amber-100 to-orange-100 rounded-2xl flex items-center justify-center mb-6">
        <svg class="w-10 h-10 text-amber-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
        </svg>
      </div>
      <h3 class="text-xl font-semibold text-neutral-900 mb-2">Upgrade to Save Players</h3>
      <p class="text-neutral-500 mb-6 max-w-sm mx-auto">Scout tier or higher is required to save players to your shortlist.</p>
      <NuxtLink to="/pricing">
        <button class="px-6 py-3 bg-primary-500 text-white font-semibold rounded-xl hover:bg-primary-600 transition-colors shadow-lg shadow-primary-500/25">
          View Plans
        </button>
      </NuxtLink>
    </div>

    <!-- Empty State -->
    <div v-else-if="savedPlayers.length === 0 && !loading" class="bg-white rounded-2xl border border-neutral-200 p-12 text-center">
      <div class="w-20 h-20 mx-auto bg-gradient-to-br from-pink-100 to-rose-100 rounded-2xl flex items-center justify-center mb-6">
        <svg class="w-10 h-10 text-rose-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
        </svg>
      </div>
      <h3 class="text-xl font-semibold text-neutral-900 mb-2">No saved players yet</h3>
      <p class="text-neutral-500 mb-6 max-w-sm mx-auto">Start browsing to find talented players and save them to your shortlist.</p>
      <NuxtLink to="/discover">
        <button class="px-6 py-3 bg-primary-500 text-white font-semibold rounded-xl hover:bg-primary-600 transition-colors shadow-lg shadow-primary-500/25">
          Browse Players
        </button>
      </NuxtLink>
    </div>

    <!-- Saved Players Grid -->
    <div v-else-if="savedPlayers.length > 0">
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
        <div 
          v-for="savedPlayer in savedPlayers" 
          :key="savedPlayer.id" 
          class="group relative bg-white rounded-2xl border border-neutral-200 overflow-hidden shadow-sm hover:shadow-xl hover:border-neutral-300 transition-all duration-300"
        >
          <!-- Remove Button -->
          <button
            @click.stop="unsavePlayer(savedPlayer.player_id)"
            class="absolute top-3 right-3 z-10 p-2.5 bg-white/90 backdrop-blur-sm rounded-xl shadow-lg opacity-0 group-hover:opacity-100 hover:bg-red-50 transition-all duration-200"
            title="Remove from saved"
          >
            <svg class="w-5 h-5 text-red-500" fill="currentColor" viewBox="0 0 24 24">
              <path d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
            </svg>
          </button>

          <NuxtLink :to="`/players/${savedPlayer.player_id}`">
            <!-- Photo -->
            <div class="aspect-[4/5] bg-neutral-200 overflow-hidden">
              <img 
                v-if="savedPlayer.player?.profile_photo_url || savedPlayer.player?.thumbnail_url"
                :src="savedPlayer.player?.profile_photo_url || savedPlayer.player?.thumbnail_url"
                :alt="`${savedPlayer.player?.first_name}`"
                class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
              />
              <div v-else class="w-full h-full flex items-center justify-center bg-gradient-to-br from-neutral-200 to-neutral-300">
                <svg class="h-16 w-16 text-neutral-400" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z"/>
                </svg>
              </div>
            </div>

            <!-- Info -->
            <div class="p-5">
              <div class="flex items-start justify-between mb-2">
                <h3 class="font-semibold text-lg text-neutral-900 group-hover:text-primary-600 transition-colors">
                  {{ savedPlayer.player?.first_name }} {{ savedPlayer.player?.last_name_init }}.
                </h3>
                <span class="text-xs text-neutral-400">
                  Saved {{ formatDate(savedPlayer.saved_at) }}
                </span>
              </div>
              <div class="flex items-center gap-2 text-sm text-neutral-500">
                <span>{{ savedPlayer.player?.position || 'Unknown' }}</span>
                <span class="w-1 h-1 rounded-full bg-neutral-300"></span>
                <span>{{ savedPlayer.player?.country || 'Unknown' }}</span>
                <span v-if="savedPlayer.player?.age" class="w-1 h-1 rounded-full bg-neutral-300"></span>
                <span v-if="savedPlayer.player?.age">{{ savedPlayer.player.age }} yrs</span>
              </div>
              <!-- Notes -->
              <p v-if="savedPlayer.notes" class="mt-2 text-sm text-neutral-600 line-clamp-2">
                {{ savedPlayer.notes }}
              </p>
            </div>
          </NuxtLink>
        </div>
      </div>

      <!-- Load More -->
      <div v-if="hasMore" class="mt-10 text-center">
        <button 
          @click="loadMore" 
          :disabled="loadingMore"
          class="px-8 py-3 border-2 border-neutral-200 text-neutral-700 font-semibold rounded-xl hover:bg-neutral-50 hover:border-neutral-300 transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed flex items-center gap-2 mx-auto"
        >
          <svg v-if="loadingMore" class="w-5 h-5 animate-spin" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
          </svg>
          {{ loadingMore ? 'Loading...' : 'Load More' }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Player, ApiResponse } from '~/types'

interface SavedPlayer {
  id: string
  player_id: string
  notes?: string
  saved_at: string
  player: {
    first_name: string
    last_name_init: string
    age: number
    position: string
    country: string
    thumbnail_url?: string
    profile_photo_url?: string
  }
}

definePageMeta({
  layout: 'dashboard',
  middleware: 'auth',
})

const subscriptionStore = useSubscriptionStore()
const savedPlayers = ref<SavedPlayer[]>([])
const loading = ref(true)
const loadingMore = ref(false)
const hasMore = ref(false)
const page = ref(1)

const api = useApi()

// Check subscription on mount
onMounted(async () => {
  await subscriptionStore.fetchSubscription()
  if (subscriptionStore.canSavePlayers) {
    fetchSavedPlayers()
  } else {
    loading.value = false
  }
})

async function fetchSavedPlayers() {
  try {
    const response = await api.get<ApiResponse<{ saved_players: SavedPlayer[]; pagination: { total: number; total_pages: number } }>>(
      `/saved-players?page=${page.value}&limit=12`,
      {},
      true
    )
    
    if (response.success && response.data) {
      if (page.value === 1) {
        savedPlayers.value = response.data.saved_players || []
      } else {
        savedPlayers.value.push(...(response.data.saved_players || []))
      }
      const totalPages = response.data.pagination?.total_pages || 1
      hasMore.value = page.value < totalPages
    }
  } catch (error) {
    console.error('Failed to fetch saved players:', error)
  } finally {
    loading.value = false
    loadingMore.value = false
  }
}

async function unsavePlayer(playerId: string) {
  try {
    const response = await api.delete<ApiResponse<null>>(`/players/${playerId}/save`, true)
    
    if (response.success) {
      savedPlayers.value = savedPlayers.value.filter(p => p.player_id !== playerId)
    }
  } catch (error) {
    console.error('Failed to unsave player:', error)
  }
}

async function loadMore() {
  loadingMore.value = true
  page.value++
  await fetchSavedPlayers()
}

function formatDate(dateString: string): string {
  const date = new Date(dateString)
  const now = new Date()
  const diffMs = now.getTime() - date.getTime()
  const diffDays = Math.floor(diffMs / (1000 * 60 * 60 * 24))
  
  if (diffDays === 0) return 'today'
  if (diffDays === 1) return 'yesterday'
  if (diffDays < 7) return `${diffDays} days ago`
  if (diffDays < 30) return `${Math.floor(diffDays / 7)} weeks ago`
  return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric' })
}
</script>
