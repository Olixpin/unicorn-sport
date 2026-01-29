<template>
  <div class="min-h-screen bg-neutral-50">
    <!-- Loading -->
    <div v-if="loading" class="min-h-screen flex items-center justify-center">
      <div class="w-10 h-10 border-4 border-primary-500 border-t-transparent rounded-full animate-spin"></div>
    </div>

    <!-- Error -->
    <div v-else-if="error" class="min-h-screen flex items-center justify-center">
      <div class="text-center">
        <h1 class="text-2xl font-bold text-neutral-900">Tournament Not Found</h1>
        <p class="text-neutral-500 mt-2">This tournament doesn't exist or isn't public.</p>
        <NuxtLink to="/tournaments">
          <button class="mt-4 px-6 py-3 bg-primary-500 text-white rounded-xl font-semibold hover:bg-primary-600 transition-colors">
            Browse Tournaments
          </button>
        </NuxtLink>
      </div>
    </div>

    <!-- Tournament Detail -->
    <div v-else-if="tournament">
      <!-- Hero -->
      <div class="relative bg-gradient-to-br from-neutral-900 via-neutral-800 to-neutral-900 overflow-hidden">
        <img 
          v-if="tournament.cover_image_url" 
          :src="tournament.cover_image_url" 
          :alt="tournament.name"
          class="absolute inset-0 w-full h-full object-cover opacity-30"
        />
        <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12 lg:py-16">
          <NuxtLink to="/tournaments" class="inline-flex items-center gap-2 text-neutral-400 hover:text-white transition-colors mb-4">
            <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
            </svg>
            Back to Tournaments
          </NuxtLink>
          <div class="flex items-start gap-3">
            <h1 class="font-display text-3xl lg:text-4xl font-bold text-white">{{ tournament.name }}</h1>
            <span v-if="tournament.featured" class="px-3 py-1 bg-amber-500 text-white text-xs font-bold rounded-full">⭐ FEATURED</span>
          </div>
          <div class="flex flex-wrap items-center gap-4 mt-4 text-neutral-300">
            <span class="flex items-center gap-2">
              <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
              </svg>
              {{ tournament.year }}
            </span>
            <span v-if="tournament.location" class="flex items-center gap-2">
              <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
              </svg>
              {{ tournament.location }}
            </span>
            <span class="flex items-center gap-2">
              <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
              </svg>
              {{ totalPlayers }} players
            </span>
          </div>
        </div>
      </div>

      <!-- Filters -->
      <div class="bg-white border-b border-neutral-200">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
          <select v-model="selectedPosition" @change="applyFilters" class="px-4 py-2 bg-neutral-100 border-0 rounded-xl text-sm focus:ring-2 focus:ring-primary-500">
            <option value="">All Positions</option>
            <option value="Goalkeeper">Goalkeeper</option>
            <option value="Defender">Defender</option>
            <option value="Midfielder">Midfielder</option>
            <option value="Forward">Forward</option>
          </select>
        </div>
      </div>

      <!-- Players Grid -->
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
        <div v-if="players.length === 0" class="text-center py-16">
          <p class="text-neutral-500">No players found for this filter.</p>
        </div>

        <div v-else class="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4 lg:gap-6">
          <NuxtLink 
            v-for="player in players" 
            :key="player.id" 
            :to="`/players/${player.id}`"
            class="group bg-white rounded-xl border border-neutral-200 overflow-hidden hover:shadow-lg hover:border-neutral-300 transition-all"
          >
            <div class="aspect-[3/4] bg-neutral-100 overflow-hidden">
              <img 
                v-if="player.profile_photo_url" 
                :src="player.profile_photo_url" 
                :alt="player.first_name"
                class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
              />
              <div v-else class="w-full h-full flex items-center justify-center">
                <svg class="w-12 h-12 text-neutral-300" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z"/>
                </svg>
              </div>
            </div>
            <div class="p-3 lg:p-4">
              <h3 class="font-semibold text-neutral-900 truncate group-hover:text-primary-600 transition-colors">
                {{ player.first_name }} {{ player.last_name_init }}.
              </h3>
              <div class="flex items-center gap-1.5 text-sm text-neutral-500 mt-1">
                <span>{{ player.position }}</span>
                <span class="w-1 h-1 rounded-full bg-neutral-300"></span>
                <span>{{ player.age }}y</span>
              </div>
              <p v-if="player.academy_name" class="text-xs text-primary-600 mt-1 truncate">{{ player.academy_name }}</p>
            </div>
          </NuxtLink>
        </div>

        <!-- Load More -->
        <div v-if="hasMore" class="mt-10 text-center">
          <button 
            @click="loadMore" 
            :disabled="loadingMore"
            class="px-8 py-3 border-2 border-neutral-200 text-neutral-700 font-semibold rounded-xl hover:bg-neutral-50 hover:border-neutral-300 transition-all disabled:opacity-50"
          >
            {{ loadingMore ? 'Loading...' : 'Load More Players' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { ApiResponse } from '~/types'

interface Tournament {
  id: number
  name: string
  year: number
  location?: string
  start_date?: string
  end_date?: string
  featured: boolean
  cover_image_url?: string
}

interface TournamentPlayer {
  id: string
  first_name: string
  last_name_init: string
  age: number
  position: string
  country: string
  profile_photo_url?: string
  academy_name?: string
  video_count: number
}

definePageMeta({ layout: 'default' })

const route = useRoute()
const config = useRuntimeConfig()

const loading = ref(true)
const error = ref(false)
const loadingMore = ref(false)
const tournament = ref<Tournament | null>(null)
const players = ref<TournamentPlayer[]>([])
const hasMore = ref(false)
const page = ref(1)
const totalPlayers = ref(0)
const selectedPosition = ref('')

useHead(() => ({
  title: tournament.value ? `${tournament.value.name} - Unicorn Sport` : 'Tournament - Unicorn Sport',
}))

onMounted(() => {
  fetchTournament()
})

async function fetchTournament() {
  const tournamentId = route.params.id as string
  
  try {
    let url = `/tournaments/${tournamentId}?page=${page.value}&limit=20`
    if (selectedPosition.value) url += `&position=${selectedPosition.value}`

    const response = await $fetch<ApiResponse<{ tournament: Tournament; players: TournamentPlayer[]; pagination: { total: number; total_pages: number } }>>(url, {
      baseURL: config.public.apiBase,
    })

    if (response.success && response.data) {
      tournament.value = response.data.tournament
      if (page.value === 1) {
        players.value = response.data.players || []
      } else {
        players.value.push(...(response.data.players || []))
      }
      totalPlayers.value = response.data.pagination?.total || 0
      hasMore.value = page.value < (response.data.pagination?.total_pages || 1)
    } else {
      error.value = true
    }
  } catch (e) {
    console.error('Failed to fetch tournament:', e)
    error.value = true
  } finally {
    loading.value = false
    loadingMore.value = false
  }
}

function applyFilters() {
  page.value = 1
  loading.value = true
  fetchTournament()
}

async function loadMore() {
  loadingMore.value = true
  page.value++
  await fetchTournament()
}
</script>
