<template>
  <NuxtLink 
    :to="`/players/${player.id}`"
    class="group block bg-white rounded-xl overflow-hidden border border-neutral-200 hover:border-primary-400 hover:shadow-md transition-all duration-200"
  >
    <!-- Player Image -->
    <div class="aspect-[4/3] relative overflow-hidden bg-neutral-100">
      <img
        v-if="playerImageUrl && !imageError"
        :src="playerImageUrl"
        :alt="`${player.first_name} ${playerLastInitial}.`"
        class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
        @error="handleImageError"
      />
      <!-- Fallback Avatar -->
      <div v-else class="w-full h-full flex items-center justify-center">
        <span class="text-3xl font-bold text-neutral-300">{{ playerInitials }}</span>
      </div>

      <!-- Position Badge - Top Right -->
      <div 
        class="absolute top-2 right-2 px-2 py-0.5 rounded-md text-xs font-bold"
        :class="positionBadgeClass"
      >
        {{ positionAbbrev }}
      </div>

      <!-- Video Count - Bottom Left -->
      <div 
        v-if="hasVideos"
        class="absolute bottom-2 left-2 flex items-center gap-1 px-2 py-1 rounded-md bg-black/80 text-white text-xs font-semibold"
      >
        <svg class="w-3 h-3" fill="currentColor" viewBox="0 0 20 20">
          <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM9.555 7.168A1 1 0 008 8v4a1 1 0 001.555.832l3-2a1 1 0 000-1.664l-3-2z" clip-rule="evenodd"/>
        </svg>
        {{ player.video_count }}
      </div>

      <!-- Save Button -->
      <button
        v-if="authStore.isAuthenticated && subscriptionStore.canSavePlayers"
        @click.prevent="toggleSave"
        :disabled="saving"
        class="absolute top-2 left-2 p-1.5 rounded-lg bg-white/90 hover:bg-white transition-all"
        :class="isSaved ? 'text-rose-500' : 'text-neutral-400 hover:text-rose-500'"
      >
        <svg class="w-4 h-4" :fill="isSaved ? 'currentColor' : 'none'" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
        </svg>
      </button>
    </div>

    <!-- Player Info -->
    <div class="p-3">
      <!-- Name Row -->
      <h3 class="font-semibold text-neutral-900 truncate group-hover:text-primary-600 transition-colors">
        {{ player.first_name }} {{ playerLastInitial }}.
      </h3>
      
      <!-- Stats Row -->
      <div class="flex items-center gap-1.5 mt-1 text-xs text-neutral-500">
        <span>{{ countryFlag }}</span>
        <span class="font-medium text-neutral-700">{{ displayAge }}</span>
        <span class="text-neutral-300">·</span>
        <span class="font-medium text-neutral-700">{{ displayHeight }}</span>
        <span v-if="player.academy_name" class="text-neutral-300">·</span>
        <span v-if="player.academy_name" class="truncate text-primary-600">{{ player.academy_name }}</span>
      </div>
    </div>
  </NuxtLink>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import type { Player, ApiResponse } from '~/types'

interface Props {
  player: Player
  initialSaved?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  initialSaved: false
})

const emit = defineEmits<{
  (e: 'saved', playerId: string): void
  (e: 'unsaved', playerId: string): void
}>()

const authStore = useAuthStore()
const subscriptionStore = useSubscriptionStore()
const api = useApi()

const isSaved = ref(props.initialSaved)
const saving = ref(false)
const imageError = ref(false)

// Safe last name initial - handles both last_name and last_name_init
// Remove trailing period if present since we add it in the template
const playerLastInitial = computed(() => {
  let initial = ''
  if (props.player?.last_name_init) {
    initial = props.player.last_name_init
  } else if (props.player?.last_name) {
    initial = props.player.last_name.charAt(0)
  }
  // Remove trailing period if the API already added one
  return initial.replace(/\.$/, '')
})

// Image URL priority: video thumbnail (action shot) > thumbnail > profile photo
// Video thumbnails show players in action which is more valuable for scouts
const playerImageUrl = computed(() => {
  return props.player?.video_thumbnail_url || props.player?.thumbnail_url || props.player?.profile_photo_url || null
})

// Check if player has videos
const hasVideos = computed(() => {
  return props.player?.video_count && props.player.video_count > 0
})

// Player initials for fallback avatar
const playerInitials = computed(() => {
  const first = props.player?.first_name?.charAt(0) || ''
  const last = playerLastInitial.value
  return (first + last).toUpperCase()
})

// Position badge styling based on position type
const positionBadgeClass = computed(() => {
  const position = props.player?.position?.toLowerCase() || ''
  if (position.includes('goalkeeper') || position.includes('gk')) {
    return 'bg-amber-500 text-white'
  } else if (position.includes('defender') || position.includes('back')) {
    return 'bg-blue-500 text-white'
  } else if (position.includes('midfielder') || position.includes('mid')) {
    return 'bg-emerald-500 text-white'
  } else if (position.includes('forward') || position.includes('striker') || position.includes('winger')) {
    return 'bg-rose-500 text-white'
  }
  return 'bg-neutral-700 text-white'
})

// Short position abbreviation
const positionAbbrev = computed(() => {
  const position = props.player?.position?.toLowerCase() || ''
  if (position.includes('goalkeeper')) return 'GK'
  if (position.includes('center back') || position.includes('centre back')) return 'CB'
  if (position.includes('left back')) return 'LB'
  if (position.includes('right back')) return 'RB'
  if (position.includes('defender')) return 'DEF'
  if (position.includes('defensive mid')) return 'CDM'
  if (position.includes('central mid')) return 'CM'
  if (position.includes('attacking mid')) return 'CAM'
  if (position.includes('midfielder')) return 'MID'
  if (position.includes('left wing')) return 'LW'
  if (position.includes('right wing')) return 'RW'
  if (position.includes('winger')) return 'WNG'
  if (position.includes('striker')) return 'ST'
  if (position.includes('forward')) return 'FWD'
  return props.player?.position?.substring(0, 3).toUpperCase() || '—'
})

// Handle image load errors
function handleImageError() {
  imageError.value = true
}

async function toggleSave() {
  saving.value = true
  try {
    if (isSaved.value) {
      await api.delete<ApiResponse<null>>(`/players/${props.player.id}/save`, true)
      isSaved.value = false
      emit('unsaved', props.player.id)
    } else {
      await api.post<ApiResponse<null>>(`/players/${props.player.id}/save`, {}, true)
      isSaved.value = true
      emit('saved', props.player.id)
    }
  } catch (error) {
    console.error('Failed to toggle save:', error)
  } finally {
    saving.value = false
  }
}

const playerAge = computed(() => {
  // First use the pre-calculated age from API if available
  if (props.player?.age != null && !isNaN(props.player.age) && props.player.age > 0) {
    return props.player.age
  }
  // Fallback: calculate from date_of_birth
  if (!props.player?.date_of_birth) return null
  const birthDate = new Date(props.player.date_of_birth)
  if (isNaN(birthDate.getTime())) return null
  const today = new Date()
  let age = today.getFullYear() - birthDate.getFullYear()
  const monthDiff = today.getMonth() - birthDate.getMonth()
  if (monthDiff < 0 || (monthDiff === 0 && today.getDate() < birthDate.getDate())) {
    age--
  }
  // Don't return invalid ages
  if (age <= 0 || age > 50) return null
  return age
})

// Display age with proper fallback
const displayAge = computed(() => {
  if (playerAge.value != null && playerAge.value > 0) {
    return `${playerAge.value}y`
  }
  return '—'
})

// Display height with proper fallback
const displayHeight = computed(() => {
  if (props.player?.height_cm && props.player.height_cm > 0) {
    return `${props.player.height_cm}cm`
  }
  return '—'
})

// Display preferred foot
const displayFoot = computed(() => {
  const foot = props.player?.preferred_foot
  if (foot === 'left') return 'L'
  if (foot === 'right') return 'R'
  if (foot === 'both') return 'Both'
  return '—'
})

// Simple country flag mapping (can be expanded)
const countryFlags: Record<string, string> = {
  'Nigeria': '🇳🇬',
  'Ghana': '🇬🇭',
  'Kenya': '🇰🇪',
  'South Africa': '🇿🇦',
  'Cameroon': '🇨🇲',
  'Senegal': '🇸🇳',
  'Egypt': '🇪🇬',
  'Morocco': '🇲🇦',
  'Algeria': '🇩🇿',
  'Tunisia': '🇹🇳',
  'Ivory Coast': '🇨🇮',
}

const countryFlag = computed(() => {
  if (!props.player?.country) return '🌍'
  return countryFlags[props.player.country] || '🌍'
})
</script>
