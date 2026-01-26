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

    <!-- Empty State -->
    <div v-else-if="savedPlayers.length === 0" class="bg-white rounded-2xl border border-neutral-200 p-12 text-center">
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
    <div v-else>
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
        <div 
          v-for="player in savedPlayers" 
          :key="player.id" 
          class="group relative bg-white rounded-2xl border border-neutral-200 overflow-hidden shadow-sm hover:shadow-xl hover:border-neutral-300 transition-all duration-300"
        >
          <!-- Remove Button -->
          <button
            @click.stop="unsavePlayer(player.id)"
            class="absolute top-3 right-3 z-10 p-2.5 bg-white/90 backdrop-blur-sm rounded-xl shadow-lg opacity-0 group-hover:opacity-100 hover:bg-red-50 transition-all duration-200"
            title="Remove from saved"
          >
            <svg class="w-5 h-5 text-red-500" fill="currentColor" viewBox="0 0 24 24">
              <path d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
            </svg>
          </button>

          <NuxtLink :to="`/players/${player.id}`">
            <!-- Photo -->
            <div class="aspect-[4/5] bg-neutral-200 overflow-hidden">
              <img 
                v-if="player.profile_photo_url"
                :src="player.profile_photo_url"
                :alt="`${player.first_name} ${player.last_name}`"
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
                  {{ player.first_name }} {{ player.last_name }}
                </h3>
                <span 
                  v-if="player.verification_status === 'verified'" 
                  class="flex-shrink-0 px-2 py-0.5 bg-primary-100 text-primary-700 text-xs font-medium rounded-full"
                >
                  Verified
                </span>
              </div>
              <div class="flex items-center gap-2 text-sm text-neutral-500">
                <span>{{ player.position }}</span>
                <span class="w-1 h-1 rounded-full bg-neutral-300"></span>
                <span>{{ player.country }}</span>
              </div>
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

definePageMeta({
  layout: 'dashboard',
  middleware: 'auth',
})

const savedPlayers = ref<Player[]>([])
const loading = ref(true)
const loadingMore = ref(false)
const hasMore = ref(false)
const page = ref(1)

const api = useApi()

async function fetchSavedPlayers() {
  try {
    const response = await api.get<ApiResponse<{ players: Player[]; has_more: boolean }>>(
      `/saved-players?page=${page.value}&limit=12`,
      {},
      true
    )
    
    if (response.success && response.data) {
      if (page.value === 1) {
        savedPlayers.value = response.data.players || []
      } else {
        savedPlayers.value.push(...(response.data.players || []))
      }
      hasMore.value = response.data.has_more || false
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
    const response = await api.delete<ApiResponse<null>>(`/saved-players/${playerId}`, true)
    
    if (response.success) {
      savedPlayers.value = savedPlayers.value.filter(p => p.id !== playerId)
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

onMounted(() => {
  fetchSavedPlayers()
})
</script>
