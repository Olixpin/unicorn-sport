<template>
  <div class="space-y-4">
    <!-- Header -->
    <div class="flex items-center justify-between">
      <div>
        <h1 class="text-xl sm:text-2xl font-bold text-neutral-900">Talent Feed</h1>
        <p class="text-neutral-500 text-sm">Watch player highlights</p>
      </div>
      <div class="flex items-center gap-2">
        <NuxtLink to="/dashboard/saved" class="relative">
          <button class="w-10 h-10 rounded-xl bg-white border border-neutral-200 flex items-center justify-center hover:bg-neutral-50 transition-colors">
            <svg class="w-5 h-5 text-neutral-600" fill="currentColor" viewBox="0 0 24 24">
              <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/>
            </svg>
            <span v-if="savedCount > 0" class="absolute -top-1 -right-1 w-5 h-5 bg-rose-500 rounded-full text-[10px] text-white font-bold flex items-center justify-center">
              {{ savedCount > 9 ? '9+' : savedCount }}
            </span>
          </button>
        </NuxtLink>
        <NuxtLink to="/discover">
          <button class="w-10 h-10 rounded-xl bg-primary-500 flex items-center justify-center hover:bg-primary-600 transition-colors">
            <svg class="w-5 h-5 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </button>
        </NuxtLink>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex items-center justify-center py-20">
      <div class="text-center">
        <div class="w-12 h-12 border-3 border-primary-500 border-t-transparent rounded-full animate-spin mx-auto mb-4" />
        <p class="text-neutral-500 text-sm">Loading highlights...</p>
      </div>
    </div>

    <!-- Empty State -->
    <div v-else-if="allHighlights.length === 0" class="flex items-center justify-center py-20">
      <div class="text-center">
        <div class="w-20 h-20 mx-auto bg-neutral-100 rounded-full flex items-center justify-center mb-6">
          <svg class="w-10 h-10 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
          </svg>
        </div>
        <h2 class="text-neutral-900 text-xl font-bold mb-2">No highlights yet</h2>
        <p class="text-neutral-500 text-sm mb-6">Be the first to discover rising talent</p>
        <NuxtLink to="/discover">
          <button class="px-6 py-3 bg-primary-500 text-white font-semibold rounded-xl hover:bg-primary-600 transition-colors">
            Browse Players
          </button>
        </NuxtLink>
      </div>
    </div>

    <!-- Video Feed - Vertical Cards -->
    <div v-else class="space-y-4">
      <div 
        v-for="(highlight, index) in allHighlights" 
        :key="highlight.id"
        class="rounded-2xl overflow-hidden shadow-xl"
      >
        <!-- Video Container - Full bleed like Instagram Reels -->
        <div class="relative aspect-[9/16] max-h-[75vh]">
          <video
            :id="`video-${index}`"
            class="w-full h-full object-cover"
            :src="highlight.stream_url"
            :poster="highlight.thumbnail_url || undefined"
            loop
            muted
            playsinline
            preload="auto"
            @click="togglePlayPause(index)"
            @timeupdate="updateProgress($event, index)"
            @loadeddata="onVideoLoaded(index)"
            @error="onVideoError(index)"
          />
          
          <!-- Play/Pause Overlay (shows when paused) -->
          <div 
            v-if="!isPlaying[index]"
            class="absolute inset-0 flex items-center justify-center cursor-pointer"
            @click="togglePlayPause(index)"
          >
            <svg class="w-16 h-16 text-white/90 drop-shadow-lg" fill="currentColor" viewBox="0 0 24 24">
              <path d="M8 5v14l11-7z"/>
            </svg>
          </div>
          
          <!-- Gradient overlays -->
          <div class="absolute inset-0 bg-gradient-to-b from-black/40 via-transparent to-black/60 pointer-events-none" />

          <!-- Right Side Actions -->
          <div class="absolute right-3 bottom-24 flex flex-col items-center gap-4">
            <!-- Save/Shortlist (Heart) - Only for paid users -->
            <button 
              v-if="highlight.player && subscriptionStore.canSavePlayers"
              class="flex flex-col items-center gap-0.5 group"
              @click.stop="toggleSaveHighlight(highlight)"
            >
              <svg 
                class="w-8 h-8 transition-all duration-200 drop-shadow-lg" 
                :class="isHighlightSaved(highlight) ? 'text-rose-500 scale-110' : 'text-white group-hover:scale-110 group-active:scale-95'"
                :fill="isHighlightSaved(highlight) ? 'currentColor' : 'none'" 
                viewBox="0 0 24 24" 
                stroke="currentColor" 
                stroke-width="2"
              >
                <path stroke-linecap="round" stroke-linejoin="round" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
              </svg>
              <span class="text-white text-[10px] font-medium drop-shadow">
                {{ isHighlightSaved(highlight) ? 'Saved' : 'Save' }}
              </span>
            </button>
            
            <!-- Upgrade prompt for free users - Shows heart with lock -->
            <NuxtLink v-if="!subscriptionStore.canSavePlayers && highlight.player" to="/pricing">
              <button class="flex flex-col items-center gap-0.5 group relative">
                <svg class="w-8 h-8 text-white/70 drop-shadow-lg group-hover:scale-110 transition-transform" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
                </svg>
                <!-- Lock badge -->
                <div class="absolute top-4 right-0 w-4 h-4 bg-amber-500 rounded-full flex items-center justify-center">
                  <svg class="w-2 h-2 text-white" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z" clip-rule="evenodd" />
                  </svg>
                </div>
                <span class="text-amber-400 text-[10px] font-medium drop-shadow">Pro</span>
              </button>
            </NuxtLink>

            <!-- View Profile -->
            <NuxtLink v-if="highlight.player" :to="`/players/${highlight.player.id}`">
              <button class="flex flex-col items-center gap-0.5 group">
                <svg class="w-8 h-8 text-white drop-shadow-lg group-hover:scale-110 transition-transform" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                </svg>
                <span class="text-white text-[10px] font-medium drop-shadow">Profile</span>
              </button>
            </NuxtLink>

            <!-- Mute/Unmute -->
            <button 
              class="flex flex-col items-center gap-0.5 group"
              @click.stop="toggleMute(index)"
            >
              <svg v-if="isMuted" class="w-8 h-8 text-white drop-shadow-lg group-hover:scale-110 transition-transform" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M5.586 15H4a1 1 0 01-1-1v-4a1 1 0 011-1h1.586l4.707-4.707C10.923 3.663 12 4.109 12 5v14c0 .891-1.077 1.337-1.707.707L5.586 15z" />
                <path stroke-linecap="round" stroke-linejoin="round" d="M17 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2" />
              </svg>
              <svg v-else class="w-8 h-8 text-white drop-shadow-lg group-hover:scale-110 transition-transform" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M15.536 8.464a5 5 0 010 7.072m2.828-9.9a9 9 0 010 12.728M5.586 15H4a1 1 0 01-1-1v-4a1 1 0 011-1h1.586l4.707-4.707C10.923 3.663 12 4.109 12 5v14c0 .891-1.077 1.337-1.707.707L5.586 15z" />
              </svg>
              <span class="text-white text-[10px] font-medium drop-shadow">{{ isMuted ? 'Sound' : 'Mute' }}</span>
            </button>
          </div>

          <!-- Bottom Player Info -->
          <div class="absolute bottom-4 left-4 right-16">
            <!-- Player Info Row -->
            <div v-if="highlight.player" class="flex items-center gap-3 mb-2">
              <NuxtLink :to="`/players/${highlight.player.id}`">
                <div class="relative w-11 h-11 rounded-full bg-gradient-to-br from-primary-500 to-emerald-500 flex items-center justify-center ring-2 ring-white/30 shadow-lg">
                  <span class="text-white font-bold text-lg">{{ highlight.player.first_name?.charAt(0) }}</span>
                  <!-- Verified badge -->
                  <div v-if="highlight.player.verification_status === 'verified'" class="absolute -bottom-0.5 -right-0.5 w-4 h-4 bg-primary-500 rounded-full flex items-center justify-center ring-2 ring-neutral-900">
                    <svg class="w-2.5 h-2.5 text-white" fill="currentColor" viewBox="0 0 20 20">
                      <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                    </svg>
                  </div>
                </div>
              </NuxtLink>
              <div class="flex-1 min-w-0">
                <NuxtLink :to="`/players/${highlight.player.id}`">
                  <h3 class="text-white font-bold text-sm truncate hover:underline inline-flex items-center gap-1">
                    {{ highlight.player.first_name }} {{ highlight.player.last_name || '' }}
                  </h3>
                </NuxtLink>
                <!-- Scout-critical stats line -->
                <p class="text-white/70 text-xs flex items-center gap-1.5 flex-wrap">
                  <span>{{ getCountryFlag(highlight.player.country) }}</span>
                  <span v-if="highlight.player.date_of_birth">{{ getPlayerAge(highlight.player.date_of_birth) }}y</span>
                  <span class="text-white/40">•</span>
                  <span>{{ highlight.player.position }}</span>
                  <template v-if="highlight.player.height_cm">
                    <span class="text-white/40">•</span>
                    <span>{{ highlight.player.height_cm }}cm</span>
                  </template>
                  <template v-if="highlight.player.preferred_foot">
                    <span class="text-white/40">•</span>
                    <span>{{ formatFoot(highlight.player.preferred_foot) }}</span>
                  </template>
                </p>
                <!-- Academy info -->
                <p v-if="highlight.player.academy" class="text-white/50 text-xs flex items-center gap-1 mt-0.5">
                  <span>{{ highlight.player.academy.name }}</span>
                  <svg v-if="highlight.player.academy.is_verified" class="w-3 h-3 text-primary-400" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M6.267 3.455a3.066 3.066 0 001.745-.723 3.066 3.066 0 013.976 0 3.066 3.066 0 001.745.723 3.066 3.066 0 012.812 2.812c.051.643.304 1.254.723 1.745a3.066 3.066 0 010 3.976 3.066 3.066 0 00-.723 1.745 3.066 3.066 0 01-2.812 2.812 3.066 3.066 0 00-1.745.723 3.066 3.066 0 01-3.976 0 3.066 3.066 0 00-1.745-.723 3.066 3.066 0 01-2.812-2.812 3.066 3.066 0 00-.723-1.745 3.066 3.066 0 010-3.976 3.066 3.066 0 00.723-1.745 3.066 3.066 0 012.812-2.812zm7.44 5.252a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                  </svg>
                </p>
              </div>
            </div>

            <!-- Description / Caption (Instagram style) -->
            <p class="text-white text-sm leading-relaxed mb-2">
              {{ getHighlightCaption(highlight) }}
            </p>

            <!-- Hashtags + Match context -->
            <div class="flex flex-wrap items-center gap-1.5">
              <span class="text-primary-400 text-xs font-medium">#{{ formatHighlightType(highlight.highlight_type) }}</span>
              <span v-if="highlight.player" class="text-primary-400 text-xs font-medium">#{{ highlight.player.position?.replace(' ', '') }}</span>
              <span v-if="highlight.match" class="text-white/50 text-xs">• {{ highlight.match.title }}</span>
            </div>
          </div>

          <!-- Progress bar -->
          <div class="absolute bottom-0 left-0 right-0 h-1 bg-white/20">
            <div 
              class="h-full bg-primary-500 transition-all duration-100"
              :style="{ width: `${videoProgress[index] || 0}%` }"
            />
          </div>
          
          <!-- Video error fallback -->
          <div 
            v-if="videoErrors[index]" 
            class="absolute inset-0 flex items-center justify-center bg-neutral-900"
          >
            <div class="text-center p-4">
              <svg class="w-12 h-12 text-neutral-600 mx-auto mb-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
              </svg>
              <p class="text-neutral-500 text-sm">Video unavailable</p>
            </div>
          </div>
        </div>
      </div>
      
      <!-- Load more / End message -->
      <div class="text-center py-8">
        <p class="text-neutral-500 text-sm">{{ allHighlights.length }} highlight{{ allHighlights.length !== 1 ? 's' : '' }}</p>
        <NuxtLink to="/discover">
          <button class="mt-3 px-5 py-2 bg-neutral-100 text-neutral-700 font-medium text-sm rounded-xl hover:bg-neutral-200 transition-colors">
            Discover More Players
          </button>
        </NuxtLink>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { ApiResponse } from '~/types'

interface Highlight {
  id: string
  highlight_type: string
  title?: string
  description?: string
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
    country: string
    date_of_birth: string
    height_cm?: number
    preferred_foot?: string
    verification_status: string
    academy?: {
      id: string
      name: string
      is_verified: boolean
    }
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
const savedPlayerIds = ref<string[]>([])

// Video feed state
const allHighlights = ref<Highlight[]>([])
const isMuted = ref(true)
const isPlaying = ref<Record<number, boolean>>({})
const videoProgress = ref<Record<number, number>>({})
const videoErrors = ref<Record<number, boolean>>({})

// Check if player is saved
function isHighlightSaved(highlight: Highlight): boolean {
  if (!highlight.player) return false
  return savedPlayerIds.value.includes(highlight.player.id)
}

// Toggle save for a highlight's player
async function toggleSaveHighlight(highlight: Highlight) {
  if (!highlight.player || !subscriptionStore.canSavePlayers) return
  
  const playerId = highlight.player.id
  
  try {
    if (isHighlightSaved(highlight)) {
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

// Get video element by index
function getVideoElement(index: number): HTMLVideoElement | null {
  return document.getElementById(`video-${index}`) as HTMLVideoElement | null
}

// Toggle play/pause
function togglePlayPause(index: number) {
  const video = getVideoElement(index)
  if (!video) return
  
  if (video.paused) {
    // Pause all other videos first
    allHighlights.value.forEach((_, i) => {
      if (i !== index) {
        const otherVideo = getVideoElement(i)
        if (otherVideo) {
          otherVideo.pause()
          isPlaying.value[i] = false
        }
      }
    })
    video.play().catch(() => {})
    isPlaying.value[index] = true
  } else {
    video.pause()
    isPlaying.value[index] = false
  }
}

// Toggle mute for all videos
function toggleMute(index: number) {
  isMuted.value = !isMuted.value
  allHighlights.value.forEach((_, i) => {
    const video = getVideoElement(i)
    if (video) {
      video.muted = isMuted.value
    }
  })
}

// Update progress bar
function updateProgress(event: Event, index: number) {
  const video = event.target as HTMLVideoElement
  if (video.duration) {
    videoProgress.value[index] = (video.currentTime / video.duration) * 100
  }
}

// Handle video loaded
function onVideoLoaded(index: number) {
  videoErrors.value[index] = false
  // Auto-play first video
  if (index === 0) {
    const video = getVideoElement(0)
    if (video) {
      video.muted = true
      video.play().then(() => {
        isPlaying.value[0] = true
      }).catch(() => {
        isPlaying.value[0] = false
      })
    }
  }
}

// Handle video error
function onVideoError(index: number) {
  videoErrors.value[index] = true
  isPlaying.value[index] = false
}

// Format highlight type
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

// Get emoji for highlight type
function getTypeEmoji(type: string): string {
  const emojis: Record<string, string> = {
    'goal': '⚽🔥',
    'assist': '🎯✨',
    'skill': '🪄💫',
    'save': '🧤🛡️',
    'tackle': '💪⚡',
    'dribble': '🌀✨',
    'other': '🌟'
  }
  return emojis[type] || '🌟'
}

// Calculate age from date of birth
function getPlayerAge(dateOfBirth: string): number {
  const today = new Date()
  const birth = new Date(dateOfBirth)
  let age = today.getFullYear() - birth.getFullYear()
  const monthDiff = today.getMonth() - birth.getMonth()
  if (monthDiff < 0 || (monthDiff === 0 && today.getDate() < birth.getDate())) {
    age--
  }
  return age
}

// Get country flag emoji from country name
function getCountryFlag(country: string): string {
  const countryFlags: Record<string, string> = {
    'Nigeria': '🇳🇬',
    'Ghana': '🇬🇭',
    'Cameroon': '🇨🇲',
    'Senegal': '🇸🇳',
    'Egypt': '🇪🇬',
    'Morocco': '🇲🇦',
    'South Africa': '🇿🇦',
    'Ivory Coast': '🇨🇮',
    'Mali': '🇲🇱',
    'Algeria': '🇩🇿',
    'Tunisia': '🇹🇳',
    'Kenya': '🇰🇪',
    'Uganda': '🇺🇬',
    'Tanzania': '🇹🇿',
    'Congo': '🇨🇩',
    'Angola': '🇦🇴',
    'Zimbabwe': '🇿🇼',
    'Zambia': '🇿🇲',
  }
  return countryFlags[country] || '🌍'
}

// Format preferred foot
function formatFoot(foot?: string): string {
  if (!foot) return ''
  const footMap: Record<string, string> = {
    'left': 'L',
    'right': 'R',
    'both': 'L/R'
  }
  return footMap[foot.toLowerCase()] || foot.charAt(0).toUpperCase()
}

// Generate caption for highlight (Instagram-style)
function getHighlightCaption(highlight: Highlight): string {
  // If has custom description, use it
  if (highlight.description) {
    return highlight.description
  }
  
  // Auto-generate engaging caption
  const playerName = highlight.player?.first_name || 'This player'
  const type = highlight.highlight_type
  const emoji = getTypeEmoji(type)
  
  const captions: Record<string, string[]> = {
    'goal': [
      `${playerName} finds the back of the net! ${emoji}`,
      `Clinical finish by ${playerName} ${emoji}`,
      `What a strike! ${emoji}`
    ],
    'assist': [
      `Vision and precision from ${playerName} ${emoji}`,
      `Perfect delivery by ${playerName} ${emoji}`,
      `${playerName} with the key pass ${emoji}`
    ],
    'skill': [
      `${playerName} showing off the technique ${emoji}`,
      `Silky smooth from ${playerName} ${emoji}`,
      `Watch this magic ${emoji}`
    ],
    'save': [
      `${playerName} keeping it out! ${emoji}`,
      `Incredible reflexes ${emoji}`,
      `What a stop by ${playerName} ${emoji}`
    ],
    'tackle': [
      `${playerName} wins it clean ${emoji}`,
      `Perfectly timed challenge ${emoji}`,
      `Solid defending from ${playerName} ${emoji}`
    ],
    'dribble': [
      `${playerName} dancing past defenders ${emoji}`,
      `Can't stop this ${emoji}`,
      `Smooth moves from ${playerName} ${emoji}`
    ]
  }
  
  const typeOptions = captions[type] || [`${playerName} showing quality ${emoji}`]
  // Use highlight id to deterministically pick a caption (so it doesn't change on re-render)
  const idx = highlight.id.charCodeAt(0) % typeOptions.length
  
  return typeOptions[idx] || `${playerName} showing quality ${emoji}`
}

onMounted(async () => {
  await subscriptionStore.fetchSubscription()
  
  try {
    // Get saved players with IDs
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

    // Fetch all highlights
    try {
      const highlightsResponse = await $fetch<ApiResponse<Highlight[]>>('/highlights/featured', {
        baseURL: config.public.apiBase,
      })
      if (highlightsResponse.success && highlightsResponse.data?.length) {
        allHighlights.value = highlightsResponse.data
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

