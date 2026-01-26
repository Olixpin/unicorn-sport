<template>
  <div class="min-h-screen bg-neutral-50">
    <!-- Loading -->
    <div v-if="loading" class="min-h-screen flex items-center justify-center">
      <div class="text-center">
        <div class="w-12 h-12 border-4 border-primary-500 border-t-transparent rounded-full animate-spin mx-auto"></div>
        <p class="mt-4 text-neutral-500">Loading player profile...</p>
      </div>
    </div>

    <!-- Error -->
    <div v-else-if="error" class="min-h-screen flex items-center justify-center">
      <div class="max-w-md mx-auto px-4 text-center">
        <div class="w-20 h-20 mx-auto bg-neutral-100 rounded-full flex items-center justify-center mb-6">
          <svg class="w-10 h-10 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
        </div>
        <h1 class="text-2xl font-bold text-neutral-900">Player Not Found</h1>
        <p class="mt-2 text-neutral-600">The player you're looking for doesn't exist or has been removed.</p>
        <NuxtLink to="/discover">
          <button class="mt-6 px-6 py-3 bg-primary-500 text-white font-semibold rounded-xl hover:bg-primary-600 transition-colors">
            Browse Players
          </button>
        </NuxtLink>
      </div>
    </div>

    <!-- Player Profile -->
    <div v-else-if="player">
      <!-- Hero Header -->
      <div class="relative bg-neutral-950 overflow-hidden">
        <div class="absolute inset-0">
          <div class="absolute inset-0 bg-gradient-to-br from-neutral-950 via-primary-950/20 to-neutral-950"></div>
          <div class="absolute top-0 left-1/4 w-[600px] h-[400px] bg-gradient-to-b from-primary-500/10 via-transparent to-transparent blur-3xl"></div>
        </div>
        
        <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pt-24 lg:pt-28 pb-8 lg:pb-12">
          <div class="flex flex-col lg:flex-row gap-8 lg:gap-12">
            <!-- Photo -->
            <div class="flex-shrink-0">
              <div class="w-48 lg:w-64 mx-auto lg:mx-0">
                <div class="aspect-[3/4] rounded-2xl overflow-hidden bg-neutral-800 ring-4 ring-white/10 shadow-2xl">
                  <img
                    v-if="player.profile_photo_url"
                    :src="player.profile_photo_url"
                    :alt="`${player.first_name} ${player.last_name}`"
                    class="w-full h-full object-cover"
                  />
                  <div v-else class="w-full h-full flex items-center justify-center bg-gradient-to-br from-neutral-700 to-neutral-800">
                    <svg class="h-20 w-20 text-neutral-500" fill="currentColor" viewBox="0 0 24 24">
                      <path d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z"/>
                    </svg>
                  </div>
                </div>
              </div>
            </div>

            <!-- Info -->
            <div class="flex-1 text-center lg:text-left">
              <div class="flex flex-col lg:flex-row items-center lg:items-start gap-3 mb-4">
                <h1 class="font-display text-3xl lg:text-4xl font-bold text-white">
                  {{ player.first_name }} {{ player.last_name }}
                </h1>
                <span 
                  v-if="player.verification_status === 'verified'" 
                  class="inline-flex items-center gap-1.5 px-3 py-1 bg-primary-500/20 text-primary-400 text-sm font-medium rounded-full border border-primary-500/30"
                >
                  <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
                  </svg>
                  Age Verified
                </span>
              </div>
              
              <div class="flex flex-wrap items-center justify-center lg:justify-start gap-3 text-neutral-300">
                <span class="inline-flex items-center gap-2 px-3 py-1.5 bg-white/5 rounded-lg">
                  <svg class="w-4 h-4 text-primary-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                  </svg>
                  {{ player.position }}
                </span>
                <span class="inline-flex items-center gap-2 px-3 py-1.5 bg-white/5 rounded-lg">
                  <span class="text-lg">{{ countryFlag }}</span>
                  {{ player.country }}
                </span>
                <span v-if="player.preferred_foot" class="inline-flex items-center gap-2 px-3 py-1.5 bg-white/5 rounded-lg">
                  {{ player.preferred_foot }} foot
                </span>
              </div>

              <!-- Stats Grid -->
              <div class="mt-8 grid grid-cols-2 sm:grid-cols-4 gap-4">
                <div class="bg-white/5 backdrop-blur-sm rounded-xl p-4 border border-white/10">
                  <div class="text-xs uppercase tracking-wider text-neutral-500 mb-1">Age</div>
                  <div class="text-2xl font-bold text-white">{{ playerAge }}</div>
                </div>
                <div class="bg-white/5 backdrop-blur-sm rounded-xl p-4 border border-white/10">
                  <div class="text-xs uppercase tracking-wider text-neutral-500 mb-1">Height</div>
                  <div class="text-2xl font-bold text-white">{{ player.height_cm || '-' }}<span class="text-sm text-neutral-400 ml-1">cm</span></div>
                </div>
                <div class="bg-white/5 backdrop-blur-sm rounded-xl p-4 border border-white/10">
                  <div class="text-xs uppercase tracking-wider text-neutral-500 mb-1">Weight</div>
                  <div class="text-2xl font-bold text-white">{{ player.weight_kg || '-' }}<span class="text-sm text-neutral-400 ml-1">kg</span></div>
                </div>
                <div class="bg-white/5 backdrop-blur-sm rounded-xl p-4 border border-white/10">
                  <div class="text-xs uppercase tracking-wider text-neutral-500 mb-1">Year</div>
                  <div class="text-2xl font-bold text-white">{{ player.tournament_year || '-' }}</div>
                </div>
              </div>

              <!-- Additional Info -->
              <div class="mt-6 flex flex-wrap items-center justify-center lg:justify-start gap-4 text-sm text-neutral-400">
                <span v-if="player.city || player.state" class="flex items-center gap-1.5">
                  <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                  </svg>
                  {{ [player.city, player.state].filter(Boolean).join(', ') }}
                </span>
                <span v-if="player.school_name" class="flex items-center gap-1.5">
                  <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 14l9-5-9-5-9 5 9 5z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 14l6.16-3.422a12.083 12.083 0 01.665 6.479A11.952 11.952 0 0012 20.055a11.952 11.952 0 00-6.824-2.998 12.078 12.078 0 01.665-6.479L12 14z" />
                  </svg>
                  {{ player.school_name }}
                </span>
                <span v-if="player.tournament" class="flex items-center gap-1.5">
                  <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 3v4M3 5h4M6 17v4m-2-2h4m5-16l2.286 6.857L21 12l-5.714 2.143L13 21l-2.286-6.857L5 12l5.714-2.143L13 3z" />
                  </svg>
                  {{ player.tournament.name }}
                </span>
              </div>

              <!-- Action Buttons -->
              <div class="mt-8 flex flex-wrap items-center justify-center lg:justify-start gap-4">
                <button 
                  v-if="subStore.canSavePlayers"
                  :disabled="savingPlayer"
                  :class="[
                    'px-6 py-3 rounded-xl font-semibold transition-all duration-200 flex items-center gap-2',
                    isSaved 
                      ? 'bg-primary-500 text-white' 
                      : 'bg-white/10 text-white hover:bg-white/20 border border-white/20'
                  ]"
                  @click="toggleSavePlayer"
                >
                  <svg v-if="savingPlayer" class="w-5 h-5 animate-spin" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
                  </svg>
                  <svg v-else class="w-5 h-5" :fill="isSaved ? 'currentColor' : 'none'" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
                  </svg>
                  {{ isSaved ? 'Saved' : 'Save Player' }}
                </button>

                <button 
                  v-if="subStore.canRequestContact"
                  class="px-6 py-3 bg-primary-500 text-white rounded-xl font-semibold hover:bg-primary-600 transition-colors flex items-center gap-2 shadow-lg shadow-primary-500/25"
                  @click="showContactModal = true"
                >
                  <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                  </svg>
                  Request Contact
                </button>

                <NuxtLink v-if="!authStore.isAuthenticated" to="/pricing">
                  <button class="px-6 py-3 bg-secondary-500 text-neutral-900 rounded-xl font-semibold hover:bg-secondary-400 transition-colors flex items-center gap-2">
                    <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                    </svg>
                    Upgrade to Contact
                  </button>
                </NuxtLink>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Videos Section -->
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <!-- Highlights (Free) -->
        <div class="mb-16">
          <div class="flex items-center justify-between mb-8">
            <div>
              <h2 class="font-display text-2xl font-bold text-neutral-900">Highlights</h2>
              <p class="text-neutral-500 mt-1">Free to watch</p>
            </div>
            <span class="px-3 py-1 bg-primary-100 text-primary-700 text-sm font-medium rounded-full">
              {{ highlights.length }} {{ highlights.length === 1 ? 'video' : 'videos' }}
            </span>
          </div>

          <div v-if="highlights.length > 0" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
            <div 
              v-for="video in highlights" 
              :key="video.id"
              class="group bg-white rounded-2xl border border-neutral-200 overflow-hidden shadow-sm hover:shadow-xl hover:border-neutral-300 transition-all duration-300 cursor-pointer"
              @click="playVideo(video)"
            >
              <div class="aspect-video relative bg-neutral-200 overflow-hidden">
                <img 
                  v-if="video.thumbnail_url"
                  :src="video.thumbnail_url"
                  :alt="video.title"
                  class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
                />
                <div class="absolute inset-0 bg-gradient-to-t from-black/60 via-transparent to-transparent"></div>
                <div class="absolute inset-0 flex items-center justify-center">
                  <div class="w-16 h-16 bg-white/95 rounded-full flex items-center justify-center shadow-lg group-hover:scale-110 transition-transform duration-300">
                    <svg class="w-7 h-7 text-primary-600 ml-1" fill="currentColor" viewBox="0 0 24 24">
                      <path d="M8 5v14l11-7z"/>
                    </svg>
                  </div>
                </div>
                <div class="absolute bottom-3 right-3 bg-black/80 text-white text-xs font-medium px-2.5 py-1 rounded-lg">
                  {{ formatDuration(video.duration_seconds) }}
                </div>
              </div>
              <div class="p-5">
                <h3 class="font-semibold text-neutral-900 truncate group-hover:text-primary-600 transition-colors">{{ video.title }}</h3>
                <p v-if="video.description" class="text-sm text-neutral-500 mt-2 line-clamp-2">{{ video.description }}</p>
              </div>
            </div>
          </div>

          <div v-else class="bg-neutral-100 rounded-2xl p-12 text-center">
            <svg class="w-12 h-12 mx-auto text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
            </svg>
            <p class="mt-4 text-neutral-600">No highlight videos available yet</p>
          </div>
        </div>

        <!-- Full Matches (Paid) -->
        <div>
          <div class="flex items-center justify-between mb-8">
            <div>
              <h2 class="font-display text-2xl font-bold text-neutral-900">Full Matches</h2>
              <p class="text-neutral-500 mt-1">Scout+ tier required</p>
            </div>
            <span v-if="fullMatches.length > 0" class="px-3 py-1 bg-secondary-100 text-secondary-700 text-sm font-medium rounded-full">
              {{ fullMatches.length }} {{ fullMatches.length === 1 ? 'match' : 'matches' }}
            </span>
          </div>

          <!-- Upgrade prompt -->
          <div v-if="!subStore.canAccessFullMatches" class="relative bg-gradient-to-br from-neutral-900 to-neutral-800 rounded-2xl p-8 lg:p-12 text-center overflow-hidden">
            <div class="absolute inset-0 bg-[url('/grid.svg')] opacity-5"></div>
            <div class="relative">
              <div class="w-20 h-20 mx-auto bg-white/10 rounded-2xl flex items-center justify-center mb-6">
                <svg class="w-10 h-10 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                </svg>
              </div>
              <h3 class="text-2xl font-bold text-white mb-2">Unlock Full Match Recordings</h3>
              <p class="text-neutral-400 max-w-md mx-auto mb-8">Get complete context with full 90-minute match recordings. See how players perform throughout an entire game.</p>
              <NuxtLink to="/pricing">
                <button class="px-8 py-4 bg-primary-500 text-white font-semibold rounded-xl hover:bg-primary-600 transition-colors shadow-lg shadow-primary-500/25">
                  View Pricing Plans
                </button>
              </NuxtLink>
            </div>
          </div>

          <!-- Full matches grid -->
          <div v-else-if="fullMatches.length > 0" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
            <div 
              v-for="video in fullMatches" 
              :key="video.id"
              class="group bg-white rounded-2xl border border-neutral-200 overflow-hidden shadow-sm hover:shadow-xl hover:border-neutral-300 transition-all duration-300 cursor-pointer"
              @click="playVideo(video)"
            >
              <div class="aspect-video relative bg-neutral-200 overflow-hidden">
                <img 
                  v-if="video.thumbnail_url"
                  :src="video.thumbnail_url"
                  :alt="video.title"
                  class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
                />
                <div class="absolute inset-0 bg-gradient-to-t from-black/60 via-transparent to-transparent"></div>
                <div class="absolute inset-0 flex items-center justify-center">
                  <div class="w-16 h-16 bg-white/95 rounded-full flex items-center justify-center shadow-lg group-hover:scale-110 transition-transform duration-300">
                    <svg class="w-7 h-7 text-primary-600 ml-1" fill="currentColor" viewBox="0 0 24 24">
                      <path d="M8 5v14l11-7z"/>
                    </svg>
                  </div>
                </div>
                <div class="absolute top-3 left-3 px-2.5 py-1 bg-secondary-500 text-neutral-900 text-xs font-bold rounded-lg">
                  FULL MATCH
                </div>
                <div class="absolute bottom-3 right-3 bg-black/80 text-white text-xs font-medium px-2.5 py-1 rounded-lg">
                  {{ formatDuration(video.duration_seconds) }}
                </div>
              </div>
              <div class="p-5">
                <h3 class="font-semibold text-neutral-900 truncate group-hover:text-primary-600 transition-colors">{{ video.title }}</h3>
              </div>
            </div>
          </div>

          <div v-else class="bg-neutral-100 rounded-2xl p-12 text-center">
            <svg class="w-12 h-12 mx-auto text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
            </svg>
            <p class="mt-4 text-neutral-600">No full match recordings available yet</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Player, Video, ApiResponse } from '~/types'

definePageMeta({
  layout: 'default',
})

const route = useRoute()
const config = useRuntimeConfig()
const authStore = useAuthStore()
const subStore = useSubscriptionStore()

const loading = ref(true)
const error = ref(false)
const player = ref<Player | null>(null)
const highlights = ref<Video[]>([])
const fullMatches = ref<Video[]>([])
const isSaved = ref(false)
const savingPlayer = ref(false)
const showContactModal = ref(false)

// Country flags
const countryFlags: Record<string, string> = {
  'Nigeria': '🇳🇬',
  'Ghana': '🇬🇭',
  'Kenya': '🇰🇪',
  'South Africa': '🇿🇦',
  'Cameroon': '🇨🇲',
  'Senegal': '🇸🇳',
}

const countryFlag = computed(() => player.value ? (countryFlags[player.value.country] || '🌍') : '')

const playerAge = computed(() => {
  if (!player.value) return 0
  const birthDate = new Date(player.value.date_of_birth)
  const today = new Date()
  let age = today.getFullYear() - birthDate.getFullYear()
  const monthDiff = today.getMonth() - birthDate.getMonth()
  if (monthDiff < 0 || (monthDiff === 0 && today.getDate() < birthDate.getDate())) {
    age--
  }
  return age
})

const formatDuration = (seconds?: number) => {
  if (!seconds) return '0:00'
  const mins = Math.floor(seconds / 60)
  const secs = seconds % 60
  return `${mins}:${secs.toString().padStart(2, '0')}`
}

const toggleSavePlayer = async () => {
  if (!player.value) return
  savingPlayer.value = true
  
  try {
    const api = useApi()
    if (isSaved.value) {
      await api.delete(`/players/${player.value.id}/save`, true)
      isSaved.value = false
    } else {
      await api.post(`/players/${player.value.id}/save`, {}, true)
      isSaved.value = true
    }
  } catch (error) {
    console.error('Failed to toggle save:', error)
  } finally {
    savingPlayer.value = false
  }
}

const playVideo = (video: Video) => {
  // TODO: Open video player modal
  console.log('Play video:', video)
}

// Fetch player data
onMounted(async () => {
  const playerId = route.params.id as string
  
  try {
    const response = await $fetch<ApiResponse<Player>>(`/players/${playerId}`, {
      baseURL: config.public.apiBase,
    })
    
    if (response.success && response.data) {
      player.value = response.data
      
      // Fetch videos
      if (response.data.videos) {
        highlights.value = response.data.videos.filter(v => v.video_type === 'highlight')
        fullMatches.value = response.data.videos.filter(v => v.video_type === 'full_match')
      }
    } else {
      error.value = true
    }
  } catch (e) {
    console.error('Failed to fetch player:', e)
    error.value = true
  } finally {
    loading.value = false
  }
})
</script>
