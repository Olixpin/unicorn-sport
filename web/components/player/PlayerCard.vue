<template>
  <div class="card-hover block group relative">
    <!-- Save Button (overlaid) -->
    <button
      v-if="authStore.isAuthenticated && subscriptionStore.canSavePlayers"
      @click.prevent="toggleSave"
      :disabled="saving"
      class="absolute top-3 left-3 z-10 p-2 rounded-lg transition-all duration-200"
      :class="[
        isSaved 
          ? 'bg-primary-500 text-white shadow-lg' 
          : 'bg-white/90 text-neutral-600 backdrop-blur-sm hover:bg-white hover:text-primary-600 shadow-md'
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
      <!-- Player Image -->
      <div class="aspect-[3/4] relative overflow-hidden bg-neutral-200">
        <img
          v-if="playerImageUrl"
          :src="playerImageUrl"
          :alt="`${player.first_name} ${playerLastInitial}.`"
          class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
          @error="handleImageError"
        />
        <div v-else class="w-full h-full flex items-center justify-center bg-gradient-to-br from-neutral-100 to-neutral-200">
          <div class="flex flex-col items-center gap-2">
            <div class="w-16 h-16 rounded-full bg-neutral-300 flex items-center justify-center">
              <span class="text-2xl font-bold text-neutral-500">{{ playerInitials }}</span>
            </div>
          </div>
        </div>

        <!-- Verified Badge -->
        <div 
          v-if="player.verification_status === 'verified'"
          class="absolute top-3 right-3 bg-primary-500 text-white px-2 py-1 rounded-full text-xs font-medium flex items-center space-x-1"
        >
          <svg class="h-3 w-3" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"/>
          </svg>
          <span>Verified</span>
        </div>
      </div>

      <!-- Player Info -->
      <div class="p-4">
        <h3 class="font-semibold text-neutral-900 truncate">
          {{ player.first_name }} {{ playerLastInitial }}.
        </h3>
        <p class="text-sm text-neutral-600 mt-1">
          {{ player.position }}<span v-if="playerAge != null"> • {{ playerAge }} years</span>
        </p>
        <p class="text-sm text-neutral-500 mt-1 flex items-center">
          <span>{{ player.city || player.state || '' }}{{ player.city || player.state ? ', ' : '' }}{{ player.country }}</span>
          <span class="ml-1">{{ countryFlag }}</span>
        </p>
        
        <!-- CTA -->
        <div class="mt-3 pt-3 border-t border-neutral-100">
          <span class="text-primary-600 text-sm font-medium group-hover:text-primary-700 flex items-center">
            <svg class="h-4 w-4 mr-1" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM9.555 7.168A1 1 0 008 8v4a1 1 0 001.555.832l3-2a1 1 0 000-1.664l-3-2z" clip-rule="evenodd"/>
            </svg>
            Watch Highlights
          </span>
        </div>
      </div>
    </NuxtLink>
  </div>
</template>

<script setup lang="ts">
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

// Safe last name initial - handles both last_name and last_name_init
const playerLastInitial = computed(() => {
  if (props.player?.last_name_init) return props.player.last_name_init
  if (props.player?.last_name) return props.player.last_name.charAt(0)
  return ''
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

// Handle image load errors
const imageError = ref(false)
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
  if (props.player?.age != null && !isNaN(props.player.age)) {
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
  return age
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
