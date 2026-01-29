<template>
  <div class="group relative bg-white rounded-2xl overflow-hidden shadow-sm hover:shadow-xl hover:shadow-primary-500/10 border border-neutral-100 hover:border-primary-200 transition-all duration-300 hover:-translate-y-1">
    <!-- Quick Actions (overlaid) - Save Button -->
    <button
      v-if="authStore.isAuthenticated && subscriptionStore.canSavePlayers"
      @click.prevent="toggleSave"
      :disabled="saving"
      class="absolute top-3 left-3 z-10 p-2.5 rounded-xl transition-all duration-200"
      :class="[
        isSaved 
          ? 'bg-primary-500 text-white shadow-lg shadow-primary-500/30' 
          : 'bg-white/90 text-neutral-500 backdrop-blur-sm hover:bg-white hover:text-primary-600 shadow-lg'
      ]"
      :title="isSaved ? 'Remove from saved' : 'Save player'"
    >
      <svg v-if="saving" class="w-4 h-4 animate-spin" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
      </svg>
      <svg v-else class="w-4 h-4" :fill="isSaved ? 'currentColor' : 'none'" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
        <path stroke-linecap="round" stroke-linejoin="round" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
      </svg>
    </button>

    <NuxtLink :to="`/players/${player.id}`">
      <!-- Player Visual - Video Thumbnail Priority -->
      <div class="aspect-[4/5] relative overflow-hidden">
        <img
          v-if="playerImageUrl && !imageError"
          :src="playerImageUrl"
          :alt="`${player.first_name} ${playerLastInitial}.`"
          class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
          @error="handleImageError"
        />
        <!-- Premium Fallback Avatar -->
        <div v-else class="w-full h-full flex items-center justify-center" :class="avatarGradient">
          <div class="flex flex-col items-center gap-3">
            <div class="w-20 h-20 rounded-2xl bg-white/20 backdrop-blur-sm flex items-center justify-center border border-white/30 shadow-lg">
              <span class="text-3xl font-bold text-white drop-shadow-sm">{{ playerInitials }}</span>
            </div>
            <div class="px-3 py-1 rounded-full bg-white/20 backdrop-blur-sm">
              <span class="text-xs font-medium text-white/90">No Photo</span>
            </div>
          </div>
        </div>

        <!-- Position Badge -->
        <div class="absolute top-3 right-3 z-10">
          <div 
            class="px-2.5 py-1 rounded-lg text-xs font-bold shadow-lg"
            :class="positionBadgeClass"
          >
            {{ positionAbbrev }}
          </div>
        </div>

        <!-- Video Count Badge (if has videos) -->
        <div 
          v-if="player.video_count && player.video_count > 0"
          class="absolute bottom-3 left-3 z-10 flex items-center gap-1.5 px-2.5 py-1.5 rounded-lg bg-black/70 backdrop-blur-sm text-white text-xs font-semibold shadow-lg"
        >
          <svg class="w-3.5 h-3.5" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM9.555 7.168A1 1 0 008 8v4a1 1 0 001.555.832l3-2a1 1 0 000-1.664l-3-2z" clip-rule="evenodd"/>
          </svg>
          {{ player.video_count }} {{ player.video_count === 1 ? 'video' : 'videos' }}
        </div>

        <!-- Verified Badge -->
        <div 
          v-if="player.is_verified || player.verification_status === 'verified'"
          class="absolute bottom-3 right-3 bg-emerald-500 text-white p-1.5 rounded-lg shadow-lg"
          title="Verified Player"
        >
          <svg class="h-3.5 w-3.5" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"/>
          </svg>
        </div>

        <!-- Gradient overlay for text readability -->
        <div class="absolute inset-x-0 bottom-0 h-24 bg-gradient-to-t from-black/60 via-black/20 to-transparent pointer-events-none"></div>
      </div>

      <!-- Player Info -->
      <div class="p-4">
        <!-- Name & Country -->
        <div class="flex items-start justify-between gap-2">
          <div class="min-w-0 flex-1">
            <h3 class="font-display font-bold text-neutral-900 truncate text-lg group-hover:text-primary-600 transition-colors">
              {{ player.first_name }} {{ playerLastInitial }}.
            </h3>
            <p class="text-sm text-neutral-500 flex items-center gap-1 mt-0.5">
              <span class="text-base">{{ countryFlag }}</span>
              <span>{{ player.country }}</span>
            </p>
          </div>
        </div>

        <!-- Academy Badge (if available) -->
        <div v-if="player.academy_name" class="mt-2">
          <span class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full bg-primary-50 text-primary-700 text-xs font-medium">
            <svg class="w-3 h-3" fill="currentColor" viewBox="0 0 20 20">
              <path d="M10.394 2.08a1 1 0 00-.788 0l-7 3a1 1 0 000 1.84L5.25 8.051a.999.999 0 01.356-.257l4-1.714a1 1 0 11.788 1.838L7.667 9.088l1.94.831a1 1 0 00.787 0l7-3a1 1 0 000-1.838l-7-3zM3.31 9.397L5 10.12v4.102a8.969 8.969 0 00-1.05-.174 1 1 0 01-.89-.89 11.115 11.115 0 01.25-3.762zM9.3 16.573A9.026 9.026 0 007 14.935v-3.957l1.818.78a3 3 0 002.364 0l5.508-2.361a11.026 11.026 0 01.25 3.762 1 1 0 01-.89.89 8.968 8.968 0 00-5.35 2.524 1 1 0 01-1.4 0z"/>
            </svg>
            {{ player.academy_name }}
          </span>
        </div>

        <!-- Stats Row -->
        <div class="flex items-center gap-3 mt-3 pt-3 border-t border-neutral-100">
          <!-- Age -->
          <div class="flex-1 text-center">
            <p class="text-xs text-neutral-400 uppercase tracking-wide">Age</p>
            <p class="font-semibold text-neutral-800 text-sm mt-0.5">{{ displayAge }}</p>
          </div>
          <!-- Height -->
          <div class="flex-1 text-center border-l border-neutral-100">
            <p class="text-xs text-neutral-400 uppercase tracking-wide">Height</p>
            <p class="font-semibold text-neutral-800 text-sm mt-0.5">{{ displayHeight }}</p>
          </div>
          <!-- Foot -->
          <div class="flex-1 text-center border-l border-neutral-100">
            <p class="text-xs text-neutral-400 uppercase tracking-wide">Foot</p>
            <p class="font-semibold text-neutral-800 text-sm mt-0.5">{{ displayFoot }}</p>
          </div>
        </div>
        
        <!-- CTA Button -->
        <div class="mt-4">
          <span class="flex items-center justify-center gap-2 w-full py-2.5 px-4 bg-neutral-50 group-hover:bg-gradient-to-r group-hover:from-primary-500 group-hover:to-emerald-500 text-neutral-600 group-hover:text-white rounded-xl text-sm font-semibold transition-all duration-300">
            <svg class="h-4 w-4" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM9.555 7.168A1 1 0 008 8v4a1 1 0 001.555.832l3-2a1 1 0 000-1.664l-3-2z" clip-rule="evenodd"/>
            </svg>
            View Profile
          </span>
        </div>
      </div>
    </NuxtLink>
  </div>
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

// Image URL - try thumbnail first (from list API), then profile photo
const playerImageUrl = computed(() => {
  return props.player?.thumbnail_url || props.player?.profile_photo_url || null
})

// Player initials for fallback avatar
const playerInitials = computed(() => {
  const first = props.player?.first_name?.charAt(0) || ''
  const last = playerLastInitial.value
  return (first + last).toUpperCase()
})

// Dynamic gradient based on position for avatar background
const avatarGradient = computed(() => {
  const position = props.player?.position?.toLowerCase() || ''
  if (position.includes('goalkeeper') || position.includes('gk')) {
    return 'bg-gradient-to-br from-amber-400 to-orange-500'
  } else if (position.includes('defender') || position.includes('back')) {
    return 'bg-gradient-to-br from-blue-400 to-indigo-500'
  } else if (position.includes('midfielder') || position.includes('mid')) {
    return 'bg-gradient-to-br from-primary-400 to-emerald-500'
  } else if (position.includes('forward') || position.includes('striker') || position.includes('winger')) {
    return 'bg-gradient-to-br from-rose-400 to-red-500'
  }
  return 'bg-gradient-to-br from-primary-400 to-emerald-500'
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
