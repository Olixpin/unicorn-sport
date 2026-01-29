<template>
  <div class="space-y-4 sm:space-y-6">
    <!-- Featured Highlight Hero - Video First -->
    <div v-if="featuredHighlight" class="relative overflow-hidden bg-neutral-950 rounded-2xl sm:rounded-3xl shadow-2xl">
      <div class="relative aspect-[16/9] sm:aspect-[2.4/1]">
        <!-- Auto-playing Video Background -->
        <div class="absolute inset-0">
          <video
            v-if="featuredHighlight.stream_url"
            ref="heroVideoRef"
            class="w-full h-full object-cover"
            :src="featuredHighlight.stream_url"
            :poster="featuredHighlight.thumbnail_url"
            loop
            muted
            playsinline
            autoplay
            @click="toggleHeroMute"
          />
          <img 
            v-else-if="featuredHighlight.thumbnail_url"
            :src="featuredHighlight.thumbnail_url"
            :alt="featuredHighlight.title"
            class="w-full h-full object-cover"
          />
          <div v-else class="w-full h-full bg-gradient-to-br from-primary-600 to-emerald-600" />
        </div>
        
        <!-- Gradient Overlays -->
        <div class="absolute inset-0 bg-gradient-to-r from-black/80 via-black/40 to-transparent" />
        <div class="absolute inset-0 bg-gradient-to-t from-black/70 via-transparent to-black/20" />
        
        <!-- Content -->
        <div class="absolute inset-0 flex flex-col justify-between p-4 sm:p-6 lg:p-8">
          <!-- Top Row -->
          <div class="flex items-start justify-between">
            <div class="flex items-center gap-2">
              <div class="w-2 h-2 bg-red-500 rounded-full animate-pulse" />
              <span class="text-white/90 text-xs sm:text-sm font-medium">Featured Highlight</span>
            </div>
            
            <!-- Mute/Unmute Button -->
            <button 
              v-if="featuredHighlight.stream_url"
              class="w-9 h-9 sm:w-10 sm:h-10 rounded-full bg-black/40 backdrop-blur-sm flex items-center justify-center hover:bg-black/60 transition-colors"
              @click="toggleHeroMute"
            >
              <svg v-if="heroMuted" class="w-4 h-4 sm:w-5 sm:h-5 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M5.586 15H4a1 1 0 01-1-1v-4a1 1 0 011-1h1.586l4.707-4.707C10.923 3.663 12 4.109 12 5v14c0 .891-1.077 1.337-1.707.707L5.586 15z" />
                <path stroke-linecap="round" stroke-linejoin="round" d="M17 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2" />
              </svg>
              <svg v-else class="w-4 h-4 sm:w-5 sm:h-5 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M15.536 8.464a5 5 0 010 7.072m2.828-9.9a9 9 0 010 12.728M5.586 15H4a1 1 0 01-1-1v-4a1 1 0 011-1h1.586l4.707-4.707C10.923 3.663 12 4.109 12 5v14c0 .891-1.077 1.337-1.707.707L5.586 15z" />
              </svg>
            </button>
          </div>
          
          <!-- Bottom Content -->
          <div class="flex flex-col sm:flex-row sm:items-end sm:justify-between gap-4">
            <!-- Player Info -->
            <div v-if="featuredHighlight.player" class="flex items-end gap-3 sm:gap-4">
              <NuxtLink :to="`/players/${featuredHighlight.player.id}`">
                <div class="w-14 h-14 sm:w-16 sm:h-16 lg:w-20 lg:h-20 rounded-xl sm:rounded-2xl overflow-hidden ring-2 ring-white/20 shadow-xl flex-shrink-0 bg-gradient-to-br from-primary-500 to-emerald-500 flex items-center justify-center">
                  <span class="text-white font-bold text-xl sm:text-2xl">{{ featuredHighlight.player.first_name?.charAt(0) }}</span>
                </div>
              </NuxtLink>
              
              <div class="min-w-0">
                <div class="flex items-center gap-2 mb-1">
                  <h2 class="text-lg sm:text-xl lg:text-2xl font-bold text-white truncate">
                    {{ featuredHighlight.player.first_name }} {{ featuredHighlight.player.last_name?.charAt(0) || '' }}.
                  </h2>
                </div>
                
                <div class="flex flex-wrap items-center gap-1.5 sm:gap-2">
                  <span class="px-2 py-0.5 sm:px-2.5 sm:py-1 bg-white/15 backdrop-blur-sm rounded-lg text-xs sm:text-sm text-white font-medium">
                    {{ featuredHighlight.player.position }}
                  </span>
                  <span v-if="featuredHighlight.highlight_type" class="px-2 py-0.5 sm:px-2.5 sm:py-1 bg-primary-500/80 backdrop-blur-sm rounded-lg text-xs sm:text-sm text-white font-medium">
                    {{ formatHighlightType(featuredHighlight.highlight_type) }}
                  </span>
                  <span v-if="featuredHighlight.duration_seconds" class="px-2 py-0.5 bg-black/40 backdrop-blur-sm rounded-lg text-xs text-white/80">
                    {{ formatDuration(featuredHighlight.duration_seconds) }}
                  </span>
                </div>
              </div>
            </div>
            
            <!-- Action Buttons -->
            <div class="flex items-center gap-2 sm:gap-3">
              <NuxtLink v-if="featuredHighlight.player" :to="`/players/${featuredHighlight.player.id}`" class="flex-1 sm:flex-initial">
                <button class="w-full sm:w-auto flex items-center justify-center gap-2 px-4 sm:px-5 py-2.5 sm:py-3 bg-white text-neutral-900 font-semibold text-sm rounded-xl hover:bg-white/90 transition-all shadow-lg">
                  <svg class="w-4 h-4 sm:w-5 sm:h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                  </svg>
                  View Profile
                </button>
              </NuxtLink>
              
              <button 
                v-if="subscriptionStore.canSavePlayers && featuredHighlight.player"
                class="flex items-center justify-center gap-2 px-4 py-2.5 sm:py-3 rounded-xl font-semibold text-sm transition-all"
                :class="isFeaturedSaved 
                  ? 'bg-primary-500 text-white' 
                  : 'bg-white/15 text-white backdrop-blur-sm border border-white/20 hover:bg-white/25'"
                @click="toggleSaveFeatured"
              >
                <svg class="w-5 h-5" :fill="isFeaturedSaved ? 'currentColor' : 'none'" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
                </svg>
                <span class="hidden sm:inline">{{ isFeaturedSaved ? 'Saved' : 'Save' }}</span>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- Loading state for featured -->
    <div v-else-if="loading" class="relative overflow-hidden bg-neutral-200 rounded-2xl sm:rounded-3xl animate-pulse">
      <div class="aspect-[16/9] sm:aspect-[2.4/1]" />
    </div>

    <!-- Quick Actions Row -->
    <div class="flex flex-col sm:flex-row gap-3 sm:gap-4">
      <!-- Discover CTA -->
      <NuxtLink to="/discover" class="flex-1">
        <div class="h-full group flex items-center justify-between px-4 sm:px-5 py-4 bg-gradient-to-r from-primary-500 to-emerald-500 rounded-2xl hover:shadow-lg hover:shadow-primary-500/25 hover:scale-[1.01] transition-all">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 sm:w-11 sm:h-11 rounded-xl bg-white/20 flex items-center justify-center">
              <svg class="w-5 h-5 sm:w-6 sm:h-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </div>
            <div class="text-left">
              <p class="text-white font-bold text-sm sm:text-base">Discover Players</p>
              <p class="text-white/70 text-xs sm:text-sm">Browse verified talent</p>
            </div>
          </div>
          <svg class="w-5 h-5 text-white/70 group-hover:translate-x-1 transition-transform" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M9 5l7 7-7 7" />
          </svg>
        </div>
      </NuxtLink>
      
      <!-- My Shortlist -->
      <NuxtLink to="/dashboard/saved" class="flex-1">
        <div class="h-full flex items-center justify-between px-4 sm:px-5 py-4 bg-white rounded-2xl border border-neutral-200 hover:border-primary-300 hover:shadow-lg transition-all group">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 sm:w-11 sm:h-11 rounded-xl bg-gradient-to-br from-rose-500 to-pink-500 flex items-center justify-center shadow-lg shadow-rose-500/25">
              <svg class="w-5 h-5 sm:w-6 sm:h-6 text-white" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/>
              </svg>
            </div>
            <div class="text-left">
              <p class="text-neutral-900 font-bold text-sm sm:text-base">My Shortlist</p>
              <p class="text-neutral-500 text-xs sm:text-sm">{{ savedCount }} players saved</p>
            </div>
          </div>
          <svg class="w-5 h-5 text-neutral-300 group-hover:text-primary-500 group-hover:translate-x-1 transition-all" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M9 5l7 7-7 7" />
          </svg>
        </div>
      </NuxtLink>
    </div>

    <!-- More Highlights Grid -->
    <div class="bg-white rounded-2xl border border-neutral-200 overflow-hidden">
      <div class="flex items-center justify-between p-4 sm:p-5 border-b border-neutral-100">
        <div class="flex items-center gap-2 sm:gap-3">
          <svg class="w-5 h-5 text-primary-500" fill="currentColor" viewBox="0 0 24 24">
            <path d="M8 5v14l11-7z"/>
          </svg>
          <h2 class="font-semibold text-sm sm:text-base text-neutral-900">More Highlights</h2>
          <span v-if="highlights.length" class="px-2 py-0.5 bg-primary-100 text-primary-700 text-[10px] sm:text-xs font-bold rounded-full">
            {{ highlights.length }} clips
          </span>
        </div>
        <NuxtLink to="/discover" class="text-xs sm:text-sm text-primary-600 hover:text-primary-700 font-medium">
          See all
        </NuxtLink>
      </div>

      <div v-if="loading" class="p-12 flex justify-center">
        <div class="w-10 h-10 border-3 border-primary-500 border-t-transparent rounded-full animate-spin" />
      </div>

      <div v-else-if="highlights.length === 0" class="p-12 text-center">
        <div class="w-16 h-16 mx-auto bg-neutral-100 rounded-2xl flex items-center justify-center mb-4">
          <svg class="w-8 h-8 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
          </svg>
        </div>
        <p class="text-neutral-500 mb-4">No highlights yet</p>
        <NuxtLink to="/discover">
          <button class="px-4 py-2 bg-primary-500 text-white text-sm font-medium rounded-xl hover:bg-primary-600 transition-colors">
            Browse Players
          </button>
        </NuxtLink>
      </div>

      <!-- Video Highlights Grid -->
      <div v-else class="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4">
        <div 
          v-for="highlight in highlights" 
          :key="highlight.id"
          class="group relative aspect-video overflow-hidden border-b border-r border-neutral-100 cursor-pointer"
          @click="playHighlight(highlight)"
        >
          <!-- Video Thumbnail -->
          <div class="absolute inset-0">
            <img 
              v-if="highlight.thumbnail_url"
              :src="highlight.thumbnail_url"
              :alt="highlight.title"
              class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
            />
            <div v-else class="w-full h-full flex items-center justify-center bg-gradient-to-br from-neutral-800 to-neutral-900">
              <svg class="w-10 h-10 text-white/30" fill="currentColor" viewBox="0 0 24 24">
                <path d="M8 5v14l11-7z"/>
              </svg>
            </div>
          </div>
          
          <!-- Play Button Overlay -->
          <div class="absolute inset-0 flex items-center justify-center bg-black/20 group-hover:bg-black/40 transition-colors">
            <div class="w-10 h-10 sm:w-12 sm:h-12 rounded-full bg-white/90 flex items-center justify-center shadow-lg group-hover:scale-110 transition-transform">
              <svg class="w-5 h-5 sm:w-6 sm:h-6 text-neutral-900 ml-0.5" fill="currentColor" viewBox="0 0 24 24">
                <path d="M8 5v14l11-7z"/>
              </svg>
            </div>
          </div>
          
          <!-- Duration Badge -->
          <div v-if="highlight.duration_seconds" class="absolute top-2 right-2 px-1.5 py-0.5 bg-black/70 rounded text-[10px] text-white font-medium">
            {{ formatDuration(highlight.duration_seconds) }}
          </div>
          
          <!-- Gradient Overlay -->
          <div class="absolute inset-x-0 bottom-0 bg-gradient-to-t from-black/80 via-black/40 to-transparent h-16" />
          
          <!-- Player Info -->
          <div v-if="highlight.player" class="absolute bottom-0 left-0 right-0 p-2 sm:p-3">
            <p class="font-semibold text-white text-xs sm:text-sm truncate">
              {{ highlight.player.first_name }} {{ highlight.player.last_name?.charAt(0) || '' }}.
            </p>
            <div class="flex items-center gap-1.5 mt-0.5">
              <span class="text-white/80 text-[10px] sm:text-xs">{{ highlight.player.position }}</span>
              <span v-if="highlight.highlight_type" class="text-white/60 text-[10px] sm:text-xs">· {{ formatHighlightType(highlight.highlight_type) }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- Video Player Modal -->
    <Teleport to="body">
      <div 
        v-if="playingHighlight" 
        class="fixed inset-0 z-50 bg-black/95 flex items-center justify-center p-4"
        @click.self="closeVideoPlayer"
      >
        <button 
          class="absolute top-4 right-4 w-10 h-10 rounded-full bg-white/10 hover:bg-white/20 flex items-center justify-center transition-colors"
          @click="closeVideoPlayer"
        >
          <svg class="w-6 h-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
        
        <div class="w-full max-w-4xl">
          <video
            ref="modalVideoRef"
            class="w-full rounded-xl"
            :src="playingHighlight.stream_url"
            :poster="playingHighlight.thumbnail_url"
            controls
            autoplay
          />
          
          <div v-if="playingHighlight.player" class="mt-4 flex items-center justify-between">
            <div class="flex items-center gap-3">
              <div class="w-10 h-10 rounded-lg bg-gradient-to-br from-primary-500 to-emerald-500 flex items-center justify-center">
                <span class="text-white font-bold">{{ playingHighlight.player.first_name?.charAt(0) }}</span>
              </div>
              <div>
                <p class="text-white font-semibold">{{ playingHighlight.player.first_name }} {{ playingHighlight.player.last_name?.charAt(0) || '' }}.</p>
                <p class="text-white/60 text-sm">{{ playingHighlight.player.position }}</p>
              </div>
            </div>
            
            <NuxtLink :to="`/players/${playingHighlight.player.id}`" @click="closeVideoPlayer">
              <button class="px-4 py-2 bg-white text-neutral-900 font-semibold text-sm rounded-xl hover:bg-white/90 transition-colors">
                View Profile
              </button>
            </NuxtLink>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- Compact Stats Row -->
    <div class="grid grid-cols-3 gap-3 sm:gap-4">
      <NuxtLink to="/dashboard/saved" class="bg-white rounded-xl sm:rounded-2xl p-3 sm:p-4 border border-neutral-200 hover:border-primary-300 hover:shadow-lg transition-all group">
        <div class="text-center">
          <p class="text-xl sm:text-2xl font-bold text-neutral-900">{{ savedCount }}</p>
          <p class="text-xs text-neutral-500 mt-0.5">Saved</p>
        </div>
      </NuxtLink>
      
      <NuxtLink to="/dashboard/contacts" class="bg-white rounded-xl sm:rounded-2xl p-3 sm:p-4 border border-neutral-200 hover:border-amber-300 hover:shadow-lg transition-all group">
        <div class="text-center">
          <p class="text-xl sm:text-2xl font-bold text-neutral-900">{{ contactsCount }}</p>
          <p class="text-xs text-neutral-500 mt-0.5">Contacts</p>
          <span v-if="pendingContacts > 0" class="inline-flex items-center gap-1 mt-1 text-[10px] font-medium text-amber-600">
            <span class="w-1.5 h-1.5 bg-amber-500 rounded-full animate-pulse" />
            {{ pendingContacts }} pending
          </span>
        </div>
      </NuxtLink>
      
      <div class="bg-white rounded-xl sm:rounded-2xl p-3 sm:p-4 border border-neutral-200">
        <div class="text-center">
          <p class="text-xl sm:text-2xl font-bold text-neutral-900">{{ viewsCount }}</p>
          <p class="text-xs text-neutral-500 mt-0.5">Viewed</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { ApiResponse } from '~/types'

// Highlight type from API
interface Highlight {
  id: string
  highlight_type: string
  title?: string
  thumbnail_url?: string
  stream_url: string
  duration_seconds?: number
  view_count: number
  created_at: string
  player?: {
    id: string
    first_name: string
    last_name?: string
    position: string
  }
  match?: {
    id: string
    title: string
  }
}

definePageMeta({
  layout: 'dashboard',
  middleware: 'auth',
})

const subscriptionStore = useSubscriptionStore()
const config = useRuntimeConfig()
const api = useApi()

const loading = ref(true)
const savedCount = ref(0)
const contactsCount = ref(0)
const pendingContacts = ref(0)
const viewsCount = ref(0)
const savedPlayerIds = ref<string[]>([])

// Video highlights
const highlights = ref<Highlight[]>([])
const featuredHighlight = ref<Highlight | null>(null)
const playingHighlight = ref<Highlight | null>(null)
const heroVideoRef = ref<HTMLVideoElement | null>(null)
const modalVideoRef = ref<HTMLVideoElement | null>(null)
const heroMuted = ref(true)

const isFeaturedSaved = computed(() => {
  if (!featuredHighlight.value?.player) return false
  return savedPlayerIds.value.includes(featuredHighlight.value.player.id)
})

// Format duration (seconds to mm:ss)
function formatDuration(seconds?: number): string {
  if (!seconds) return '0:00'
  const mins = Math.floor(seconds / 60)
  const secs = seconds % 60
  return `${mins}:${secs.toString().padStart(2, '0')}`
}

// Format highlight type for display
function formatHighlightType(type: string): string {
  const types: Record<string, string> = {
    'goal': 'Goal',
    'assist': 'Assist',
    'skill': 'Skill',
    'save': 'Save',
    'tackle': 'Tackle',
    'dribble': 'Dribble',
    'other': 'Highlight'
  }
  return types[type] || 'Highlight'
}

// Toggle hero video mute
function toggleHeroMute() {
  heroMuted.value = !heroMuted.value
  if (heroVideoRef.value) {
    heroVideoRef.value.muted = heroMuted.value
  }
}

// Play highlight in modal
function playHighlight(highlight: Highlight) {
  playingHighlight.value = highlight
  // Pause hero video when modal opens
  if (heroVideoRef.value) {
    heroVideoRef.value.pause()
  }
}

// Close video player modal
function closeVideoPlayer() {
  if (modalVideoRef.value) {
    modalVideoRef.value.pause()
  }
  playingHighlight.value = null
  // Resume hero video
  if (heroVideoRef.value) {
    heroVideoRef.value.play()
  }
}

// Toggle save for featured player
async function toggleSaveFeatured() {
  if (!featuredHighlight.value?.player) return
  const playerId = featuredHighlight.value.player.id
  
  try {
    if (isFeaturedSaved.value) {
      await api.delete(`/players/${playerId}/save`, true)
      savedPlayerIds.value = savedPlayerIds.value.filter(id => id !== playerId)
      savedCount.value = Math.max(0, savedCount.value - 1)
    } else {
      await api.post(`/players/${playerId}/save`, {}, true)
      savedPlayerIds.value.push(playerId)
      savedCount.value++
    }
  } catch (error) {
    console.error('Failed to toggle save:', error)
  }
}

onMounted(async () => {
  await subscriptionStore.fetchSubscription()
  
  try {
    // Get saved players with IDs for tracking
    if (subscriptionStore.canSavePlayers) {
      const savedResponse = await api.get<ApiResponse<{ saved_players: { player_id: string }[]; pagination: { total: number } }>>(
        '/saved-players?limit=100',
        {},
        true
      )
      if (savedResponse.success && savedResponse.data) {
        savedCount.value = savedResponse.data.pagination?.total || savedResponse.data.saved_players?.length || 0
        savedPlayerIds.value = savedResponse.data.saved_players?.map(sp => sp.player_id) || []
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

    // Fetch featured highlights (video-first approach)
    try {
      const highlightsResponse = await $fetch<ApiResponse<Highlight[]>>('/highlights/featured', {
        baseURL: config.public.apiBase,
      })
      if (highlightsResponse.success && highlightsResponse.data?.length) {
        const allHighlights = highlightsResponse.data
        // First one is featured (auto-play)
        featuredHighlight.value = allHighlights[0] ?? null
        // Rest go in grid
        highlights.value = allHighlights.slice(1)
      }
    } catch {
      console.log('Highlights endpoint not available')
    }
  } catch (error) {
    console.error('Failed to load dashboard data:', error)
  } finally {
    loading.value = false
  }
})
</script>
