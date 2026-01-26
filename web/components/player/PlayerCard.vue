<template>
  <NuxtLink :to="`/players/${player.id}`" class="card-hover block group">
    <!-- Player Image -->
    <div class="aspect-[3/4] relative overflow-hidden bg-neutral-200">
      <img
        v-if="player.profile_photo_url"
        :src="player.profile_photo_url"
        :alt="`${player.first_name} ${player.last_name}`"
        class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
      />
      <div v-else class="w-full h-full flex items-center justify-center bg-neutral-200">
        <svg class="h-20 w-20 text-neutral-400" fill="currentColor" viewBox="0 0 24 24">
          <path d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z"/>
        </svg>
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
        {{ player.first_name }} {{ player.last_name.charAt(0) }}.
      </h3>
      <p class="text-sm text-neutral-600 mt-1">
        {{ player.position }} • {{ playerAge }} years
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
</template>

<script setup lang="ts">
import type { Player } from '~/types'

interface Props {
  player: Player
}

const props = defineProps<Props>()

const playerAge = computed(() => {
  const birthDate = new Date(props.player.date_of_birth)
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

const countryFlag = computed(() => countryFlags[props.player.country] || '🌍')
</script>
