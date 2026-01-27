<template>
  <div class="bg-neutral-50 rounded-xl p-4">
    <div class="flex items-center gap-3">
      <!-- Logo -->
      <div class="w-10 h-10 rounded-lg bg-white border-2 border-neutral-200 overflow-hidden flex-shrink-0">
        <img
          v-if="academy.logo_url"
          :src="academy.logo_url"
          :alt="academy.name"
          class="w-full h-full object-cover"
        />
        <div 
          v-else 
          class="w-full h-full bg-gradient-to-br from-secondary-400 to-secondary-600 flex items-center justify-center"
        >
          <span class="text-white font-bold text-sm">{{ academy.name?.charAt(0) || '?' }}</span>
        </div>
      </div>
      
      <!-- Info -->
      <div class="min-w-0 flex-1">
        <div class="flex items-center gap-2">
          <p class="font-semibold text-neutral-900 truncate">{{ academy.name }}</p>
          <span 
            v-if="showVerifiedBadge && academy.is_verified"
            class="inline-flex items-center gap-1 px-1.5 py-0.5 bg-emerald-100 text-emerald-700 rounded text-xs font-medium"
          >
            <svg class="w-3 h-3" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M6.267 3.455a3.066 3.066 0 001.745-.723 3.066 3.066 0 013.976 0 3.066 3.066 0 001.745.723 3.066 3.066 0 012.812 2.812c.051.643.304 1.254.723 1.745a3.066 3.066 0 010 3.976 3.066 3.066 0 00-.723 1.745 3.066 3.066 0 01-2.812 2.812 3.066 3.066 0 00-1.745.723 3.066 3.066 0 01-3.976 0 3.066 3.066 0 00-1.745-.723 3.066 3.066 0 01-2.812-2.812 3.066 3.066 0 00-.723-1.745 3.066 3.066 0 010-3.976 3.066 3.066 0 00.723-1.745 3.066 3.066 0 012.812-2.812zm7.44 5.252a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
            </svg>
            Verified
          </span>
        </div>
        <p class="text-sm text-neutral-500 truncate">
          <span v-if="academy.city">{{ academy.city }}, </span>{{ academy.country }}
        </p>
      </div>
      
      <!-- Optional slot for actions -->
      <slot name="actions" />
    </div>
    
    <!-- Optional additional info -->
    <div v-if="showDetails" class="mt-3 pt-3 border-t border-neutral-200 grid grid-cols-2 gap-2 text-sm">
      <div v-if="academy.player_count !== undefined" class="flex items-center gap-1.5 text-neutral-600">
        <svg class="w-4 h-4 text-primary-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
        </svg>
        <span class="font-medium">{{ academy.player_count }}</span> players
      </div>
      <div v-if="academy.founded_year" class="flex items-center gap-1.5 text-neutral-600">
        <svg class="w-4 h-4 text-secondary-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
        </svg>
        Est. {{ academy.founded_year }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Academy } from '~/schemas/academy'

interface Props {
  academy: Academy
  showVerifiedBadge?: boolean
  showDetails?: boolean
}

withDefaults(defineProps<Props>(), {
  showVerifiedBadge: false,
  showDetails: false,
})
</script>
