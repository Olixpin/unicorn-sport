<template>
  <UiModal
    v-model="isOpen"
    size="video"
    :show-close-button="true"
    :no-padding="true"
    close-button-class="!text-white hover:!bg-white/20 !right-2 !top-2 z-20"
    backdrop-class="!bg-black/80"
    panel-class="!bg-neutral-900 !rounded-xl overflow-hidden"
    @closed="onClosed"
  >
    <div class="relative">
      <!-- Video Container -->
      <div class="relative aspect-video bg-black">
        <!-- Loading State with Timeout -->
        <div 
          v-if="loading && !hasTimedOut" 
          class="absolute inset-0 flex items-center justify-center bg-neutral-900"
        >
          <div class="flex flex-col items-center gap-4">
            <div class="w-12 h-12 border-4 border-primary-500/30 border-t-primary-500 rounded-full animate-spin" />
            <span class="text-white/60 text-sm">Loading video...</span>
          </div>
        </div>

        <!-- Error State -->
        <div 
          v-else-if="error || hasTimedOut" 
          class="absolute inset-0 flex items-center justify-center bg-neutral-900"
        >
          <div class="flex flex-col items-center gap-4 text-center px-8">
            <div class="w-16 h-16 rounded-full bg-red-500/20 flex items-center justify-center">
              <svg class="w-8 h-8 text-red-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
            <div>
              <p class="text-white font-medium">Unable to load video</p>
              <p class="text-white/60 text-sm mt-1">{{ error || 'Connection timed out' }}</p>
            </div>
            <button 
              @click="retry"
              class="px-4 py-2 bg-white/10 hover:bg-white/20 text-white rounded-lg text-sm transition-colors"
            >
              Try Again
            </button>
          </div>
        </div>

        <!-- Demo Mode Banner (shows when using fallback content) -->
        <div 
          v-if="isDemoMode && !loading" 
          class="absolute top-4 right-4 z-10 flex items-center gap-2 px-3 py-1.5 bg-secondary-500/90 backdrop-blur-sm rounded-full"
        >
          <svg class="w-3.5 h-3.5 text-neutral-900" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd" />
          </svg>
          <span class="text-neutral-900 text-xs font-semibold">DEMO PREVIEW</span>
        </div>

        <!-- YouTube Embed -->
        <iframe
          v-if="effectiveYoutubeId && !error && !hasTimedOut"
          :src="`https://www.youtube.com/embed/${effectiveYoutubeId}?autoplay=1&rel=0&modestbranding=1&showinfo=0`"
          class="absolute inset-0 w-full h-full"
          frameborder="0"
          allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
          allowfullscreen
          @load="onIframeLoad"
        />

        <!-- Native Video Player -->
        <video
          v-else-if="videoUrl && !error && !hasTimedOut"
          ref="videoRef"
          :src="videoUrl"
          :poster="thumbnailUrl"
          class="absolute inset-0 w-full h-full object-contain"
          controls
          autoplay
          playsinline
          @loadstart="loading = true"
          @canplay="onVideoCanPlay"
          @error="onVideoError"
        >
          <source :src="videoUrl" :type="videoType" />
          Your browser does not support the video tag.
        </video>

        <!-- No Content State (graceful fallback) -->
        <div 
          v-if="!effectiveYoutubeId && !videoUrl && !loading && !error" 
          class="absolute inset-0 flex items-center justify-center overflow-hidden"
        >
          <!-- Dynamic Background - African Football Vibe -->
          <div class="absolute inset-0">
            <!-- Base gradient with brand colors -->
            <div class="absolute inset-0 bg-gradient-to-br from-primary-950 via-neutral-950 to-emerald-950"></div>
            
            <!-- Subtle pitch lines pattern -->
            <div class="absolute inset-0 opacity-[0.08]">
              <svg class="w-full h-full" viewBox="0 0 800 450" preserveAspectRatio="xMidYMid slice">
                <!-- Pitch outline -->
                <rect x="40" y="30" width="720" height="390" fill="none" stroke="white" stroke-width="2"/>
                <!-- Center line -->
                <line x1="400" y1="30" x2="400" y2="420" stroke="white" stroke-width="2"/>
                <!-- Center circle -->
                <circle cx="400" cy="225" r="60" fill="none" stroke="white" stroke-width="2"/>
                <circle cx="400" cy="225" r="3" fill="white"/>
                <!-- Penalty boxes -->
                <rect x="40" y="120" width="100" height="210" fill="none" stroke="white" stroke-width="2"/>
                <rect x="660" y="120" width="100" height="210" fill="none" stroke="white" stroke-width="2"/>
                <!-- Goal boxes -->
                <rect x="40" y="165" width="40" height="120" fill="none" stroke="white" stroke-width="2"/>
                <rect x="720" y="165" width="40" height="120" fill="none" stroke="white" stroke-width="2"/>
              </svg>
            </div>
            
            <!-- Animated gradient orbs (brand colors) -->
            <div class="absolute top-0 left-1/4 w-80 h-80 bg-primary-500/20 rounded-full blur-3xl animate-pulse"></div>
            <div class="absolute bottom-0 right-1/4 w-64 h-64 bg-secondary-500/15 rounded-full blur-3xl animate-pulse" style="animation-delay: 1s;"></div>
            <div class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-96 h-96 bg-emerald-500/10 rounded-full blur-3xl animate-pulse" style="animation-delay: 2s;"></div>
            
            <!-- African-inspired geometric pattern (subtle) -->
            <div class="absolute inset-0 opacity-[0.04]">
              <svg class="w-full h-full" viewBox="0 0 100 100" preserveAspectRatio="none">
                <defs>
                  <pattern id="africanPattern" x="0" y="0" width="20" height="20" patternUnits="userSpaceOnUse">
                    <!-- Kente-inspired diamond pattern -->
                    <path d="M10 0 L20 10 L10 20 L0 10 Z" fill="none" stroke="white" stroke-width="0.5"/>
                    <circle cx="10" cy="10" r="2" fill="white" opacity="0.5"/>
                  </pattern>
                </defs>
                <rect width="100" height="100" fill="url(#africanPattern)"/>
              </svg>
            </div>
            
            <!-- Floating particles (like stadium lights) -->
            <div class="absolute inset-0 overflow-hidden">
              <div class="absolute top-[10%] left-[15%] w-1 h-1 bg-secondary-400 rounded-full animate-float opacity-60"></div>
              <div class="absolute top-[20%] right-[20%] w-1.5 h-1.5 bg-primary-400 rounded-full animate-float opacity-50" style="animation-delay: 0.5s;"></div>
              <div class="absolute bottom-[30%] left-[25%] w-1 h-1 bg-white rounded-full animate-float opacity-40" style="animation-delay: 1s;"></div>
              <div class="absolute top-[40%] right-[30%] w-2 h-2 bg-secondary-500 rounded-full animate-float opacity-30" style="animation-delay: 1.5s;"></div>
              <div class="absolute bottom-[20%] right-[15%] w-1 h-1 bg-primary-300 rounded-full animate-float opacity-50" style="animation-delay: 2s;"></div>
            </div>
          </div>
          
          <!-- Content -->
          <div class="relative z-10 flex flex-col items-center gap-6 text-center max-w-md px-8">
            <!-- Animated Football Icon with African flag colors ring -->
            <div class="relative">
              <!-- Outer ring with gradient (African flag colors) -->
              <div class="absolute -inset-3 rounded-full bg-gradient-to-r from-primary-500 via-secondary-500 to-red-500 opacity-20 blur-sm animate-spin-slow"></div>
              
              <div class="relative w-28 h-28 rounded-full bg-gradient-to-br from-primary-500/30 via-emerald-500/20 to-secondary-500/30 flex items-center justify-center backdrop-blur-sm border border-white/10">
                <!-- Football icon -->
                <svg class="w-14 h-14 text-white" viewBox="0 0 24 24" fill="none">
                  <!-- Outer circle -->
                  <circle cx="12" cy="12" r="10" stroke="currentColor" stroke-width="1.5" opacity="0.8"/>
                  <!-- Pentagon pattern -->
                  <path d="M12 2 L12 6 M12 18 L12 22" stroke="currentColor" stroke-width="1" opacity="0.5"/>
                  <path d="M2 12 L6 12 M18 12 L22 12" stroke="currentColor" stroke-width="1" opacity="0.5"/>
                  <path d="M4.93 4.93 L7.76 7.76 M16.24 16.24 L19.07 19.07" stroke="currentColor" stroke-width="1" opacity="0.5"/>
                  <path d="M4.93 19.07 L7.76 16.24 M16.24 7.76 L19.07 4.93" stroke="currentColor" stroke-width="1" opacity="0.5"/>
                  <!-- Center pentagon -->
                  <path d="M12 7 L15.5 10 L14 14.5 L10 14.5 L8.5 10 Z" fill="currentColor" opacity="0.3" stroke="currentColor" stroke-width="1"/>
                </svg>
                
                <!-- Pulse effect -->
                <div class="absolute inset-0 rounded-full bg-primary-500/30 animate-ping opacity-30"></div>
              </div>
              
              <!-- "Coming Soon" badge -->
              <div class="absolute -bottom-2 -right-2 px-3 py-1 bg-gradient-to-r from-secondary-500 to-secondary-600 rounded-full shadow-lg">
                <span class="text-neutral-900 text-xs font-bold tracking-wide">SOON</span>
              </div>
            </div>
            
            <!-- Text Content -->
            <div>
              <h3 class="text-white text-2xl font-bold mb-3">
                Highlights Coming Soon
              </h3>
              <p class="text-neutral-300 text-sm leading-relaxed">
                We're curating the best African football highlights. 
                <span class="text-primary-400">Be among the first</span> to discover emerging talent from across the continent.
              </p>
            </div>
            
            <!-- Stats teaser (social proof) -->
            <div class="flex items-center gap-6 text-center">
              <div>
                <div class="text-2xl font-bold text-white">500+</div>
                <div class="text-xs text-neutral-400">Players</div>
              </div>
              <div class="w-px h-8 bg-white/10"></div>
              <div>
                <div class="text-2xl font-bold text-white">15+</div>
                <div class="text-xs text-neutral-400">Countries</div>
              </div>
              <div class="w-px h-8 bg-white/10"></div>
              <div>
                <div class="text-2xl font-bold text-white">50+</div>
                <div class="text-xs text-neutral-400">Scouts</div>
              </div>
            </div>
            
            <!-- CTA Buttons -->
            <div class="flex flex-col sm:flex-row gap-3 w-full sm:w-auto">
              <NuxtLink 
                to="/discover"
                class="px-6 py-3 bg-gradient-to-r from-primary-500 to-emerald-500 hover:from-primary-600 hover:to-emerald-600 text-white rounded-xl text-sm font-semibold transition-all duration-300 flex items-center justify-center gap-2 shadow-lg shadow-primary-500/25"
                @click="isOpen = false"
              >
                <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                </svg>
                Discover Players
              </NuxtLink>
              <button 
                @click="isOpen = false"
                class="px-6 py-3 bg-white/5 hover:bg-white/10 text-white rounded-xl text-sm font-medium transition-colors border border-white/10"
              >
                Maybe Later
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Video Info Bar (Optional) -->
      <div v-if="showInfo && (title || player) && !error && !hasTimedOut" class="p-4 bg-neutral-900 border-t border-white/10">
        <div class="flex items-center gap-4">
          <!-- Player Avatar -->
          <div v-if="player" class="flex-shrink-0">
            <div 
              v-if="player.avatarUrl"
              class="w-12 h-12 rounded-xl overflow-hidden"
            >
              <img :src="player.avatarUrl" :alt="player.name" class="w-full h-full object-cover" />
            </div>
            <div 
              v-else 
              class="w-12 h-12 rounded-xl bg-gradient-to-br from-primary-500 to-emerald-500 flex items-center justify-center"
            >
              <span class="text-white font-bold">{{ getInitials(player.name) }}</span>
            </div>
          </div>

          <!-- Video Info -->
          <div class="flex-1 min-w-0">
            <div class="flex items-center gap-2">
              <h4 class="text-white font-semibold truncate">{{ effectiveTitle }}</h4>
              <span v-if="isDemoMode" class="px-2 py-0.5 bg-secondary-500/20 text-secondary-400 text-xs rounded-full">Demo</span>
            </div>
            <p v-if="player" class="text-neutral-400 text-sm truncate">
              {{ player.name }}
              <span v-if="player.position"> • {{ player.position }}</span>
              <span v-if="player.country"> • {{ player.country }}</span>
            </p>
            <p v-else-if="isDemoMode" class="text-neutral-400 text-sm">
              Sample highlight reel showcasing platform capabilities
            </p>
          </div>

          <!-- Action Buttons -->
          <div class="flex items-center gap-2">
            <NuxtLink 
              to="/discover"
              class="px-4 py-2 bg-primary-500 hover:bg-primary-600 text-white rounded-lg text-sm font-medium transition-colors flex items-center gap-2"
              @click="isOpen = false"
            >
              <span>Explore Players</span>
              <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3" />
              </svg>
            </NuxtLink>
          </div>
        </div>
      </div>
    </div>
  </UiModal>
</template>

<script setup lang="ts">
import { ref, computed, watch, onUnmounted } from 'vue'

// =============================================================================
// TYPES
// =============================================================================

interface PlayerInfo {
  id?: string
  name: string
  position?: string
  country?: string
  avatarUrl?: string
}

interface VideoPlayerModalProps {
  modelValue: boolean
  videoUrl?: string
  youtubeId?: string
  thumbnailUrl?: string
  title?: string
  player?: PlayerInfo
  showInfo?: boolean
  videoType?: string
  /** Use demo content if no video provided */
  allowDemo?: boolean
  /** Timeout in ms before showing error (default: 10000) */
  loadTimeout?: number
}

interface VideoPlayerModalEmits {
  (e: 'update:modelValue', value: boolean): void
  (e: 'close'): void
  (e: 'closed'): void
  (e: 'view-profile', playerId: string): void
}

// =============================================================================
// DEMO VIDEOS - Curated African football content
// These are publicly embeddable YouTube videos showcasing African football
// =============================================================================

const DEMO_VIDEOS = [
  {
    id: 'xYMXhmH39Xg', // African football skills compilation
    title: 'African Football Rising Stars',
    description: 'Showcasing the best emerging talent from across Africa',
  },
  {
    id: 'QhWGNy6Az0Q', // African Cup of Nations highlights
    title: 'AFCON Best Moments',
    description: 'Highlights from Africa Cup of Nations',
  },
  {
    id: '2Z4m4lnjxkY', // Youth football Africa
    title: 'Future Champions',
    description: 'Young African players making their mark',
  },
]

// =============================================================================
// PROPS & EMITS
// =============================================================================

const props = withDefaults(defineProps<VideoPlayerModalProps>(), {
  showInfo: true,
  videoType: 'video/mp4',
  allowDemo: true,
  loadTimeout: 10000,
})

const emit = defineEmits<VideoPlayerModalEmits>()

// =============================================================================
// STATE
// =============================================================================

const isOpen = computed({
  get: () => props.modelValue,
  set: (value: boolean) => emit('update:modelValue', value),
})

const videoRef = ref<HTMLVideoElement | null>(null)
const loading = ref(true)
const error = ref<string | null>(null)
const hasTimedOut = ref(false)
const loadTimeoutId = ref<ReturnType<typeof setTimeout> | null>(null)

// =============================================================================
// COMPUTED
// =============================================================================

// Determine if we're in demo mode (using fallback content)
const isDemoMode = computed(() => {
  return props.allowDemo && !props.videoUrl && !props.youtubeId
})

// Get the effective YouTube ID (props or demo fallback)
const effectiveYoutubeId = computed(() => {
  if (props.youtubeId) return props.youtubeId
  if (props.videoUrl) return null // Use native video instead
  if (props.allowDemo && DEMO_VIDEOS.length > 0) {
    // Pick a random demo video
    const randomIndex = Math.floor(Math.random() * DEMO_VIDEOS.length)
    const demoVideo = DEMO_VIDEOS[randomIndex]
    return demoVideo ? demoVideo.id : null
  }
  return null
})

// Get the effective title
const effectiveTitle = computed(() => {
  if (props.title) return props.title
  if (isDemoMode.value) {
    const video = DEMO_VIDEOS.find(v => v.id === effectiveYoutubeId.value)
    return video?.title || 'African Football Highlights'
  }
  return 'Highlight Reel'
})

// =============================================================================
// METHODS
// =============================================================================

function getInitials(name: string): string {
  return name
    .split(' ')
    .map((n) => n[0])
    .join('')
    .toUpperCase()
    .slice(0, 2)
}

function startLoadTimeout() {
  clearLoadTimeout()
  loadTimeoutId.value = setTimeout(() => {
    if (loading.value) {
      hasTimedOut.value = true
      loading.value = false
    }
  }, props.loadTimeout)
}

function clearLoadTimeout() {
  if (loadTimeoutId.value) {
    clearTimeout(loadTimeoutId.value)
    loadTimeoutId.value = null
  }
}

function onClosed() {
  // Pause video when modal closes
  if (videoRef.value) {
    videoRef.value.pause()
    videoRef.value.currentTime = 0
  }
  error.value = null
  loading.value = true
  hasTimedOut.value = false
  clearLoadTimeout()
  emit('closed')
}

function onVideoError(e: Event) {
  clearLoadTimeout()
  loading.value = false
  error.value = 'Failed to load video. Please try again.'
  console.error('Video error:', e)
}

function onVideoCanPlay() {
  clearLoadTimeout()
  loading.value = false
}

function onIframeLoad() {
  clearLoadTimeout()
  loading.value = false
}

function retry() {
  error.value = null
  hasTimedOut.value = false
  loading.value = true
  startLoadTimeout()
  // Force video reload
  if (videoRef.value) {
    videoRef.value.load()
  }
}

// =============================================================================
// WATCHERS
// =============================================================================

watch(() => props.modelValue, (open: boolean) => {
  if (open) {
    loading.value = true
    error.value = null
    hasTimedOut.value = false
    startLoadTimeout()
  } else {
    clearLoadTimeout()
  }
})

// =============================================================================
// LIFECYCLE
// =============================================================================

onUnmounted(() => {
  clearLoadTimeout()
})
</script>
