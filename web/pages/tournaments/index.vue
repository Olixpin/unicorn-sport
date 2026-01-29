<template>
  <div class="min-h-screen bg-neutral-50">
    <!-- Hero Header -->
    <div class="bg-gradient-to-br from-neutral-900 via-neutral-800 to-neutral-900">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12 lg:py-16">
        <h1 class="font-display text-3xl lg:text-4xl font-bold text-white mb-3">
          Tournaments & Events
        </h1>
        <p class="text-neutral-400 text-lg max-w-2xl">
          Browse public tournaments and discover talented players from competitions across Africa.
        </p>
      </div>
    </div>

    <!-- Filters -->
    <div class="bg-white border-b border-neutral-200 sticky top-0 z-10">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
        <div class="flex flex-wrap gap-3">
          <button 
            @click="toggleFeatured"
            :class="[
              'px-4 py-2 rounded-xl text-sm font-medium transition-colors',
              featuredOnly ? 'bg-primary-500 text-white' : 'bg-neutral-100 text-neutral-700 hover:bg-neutral-200'
            ]"
          >
            ⭐ Featured Only
          </button>
          <select v-model="selectedYear" @change="fetchTournaments" class="px-4 py-2 bg-neutral-100 border-0 rounded-xl text-sm focus:ring-2 focus:ring-primary-500">
            <option value="">All Years</option>
            <option v-for="year in availableYears" :key="year" :value="year">{{ year }}</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="flex justify-center py-20">
      <div class="w-10 h-10 border-4 border-primary-500 border-t-transparent rounded-full animate-spin"></div>
    </div>

    <!-- Tournaments Grid -->
    <div v-else class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
      <!-- Featured Tournaments -->
      <div v-if="featuredTournaments.length > 0 && !featuredOnly" class="mb-12">
        <h2 class="font-display text-2xl font-bold text-neutral-900 mb-6">Featured Events</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <NuxtLink 
            v-for="t in featuredTournaments" 
            :key="t.id" 
            :to="`/tournaments/${t.id}`"
            class="group relative bg-white rounded-2xl border border-neutral-200 overflow-hidden shadow-sm hover:shadow-xl hover:border-primary-300 transition-all duration-300"
          >
            <div class="aspect-[16/9] bg-gradient-to-br from-primary-500 to-primary-700 overflow-hidden">
              <img 
                v-if="t.cover_image_url" 
                :src="t.cover_image_url" 
                :alt="t.name"
                class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
              />
              <div v-else class="w-full h-full flex items-center justify-center">
                <svg class="w-16 h-16 text-white/30" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 4v12l-4-2-4 2V4M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
              </div>
            </div>
            <div class="absolute top-3 left-3">
              <span class="px-3 py-1 bg-amber-500 text-white text-xs font-bold rounded-full shadow-lg">⭐ FEATURED</span>
            </div>
            <div class="p-5">
              <h3 class="font-semibold text-lg text-neutral-900 group-hover:text-primary-600 transition-colors">{{ t.name }}</h3>
              <div class="flex items-center gap-2 text-sm text-neutral-500 mt-2">
                <span>{{ t.year }}</span>
                <span class="w-1 h-1 rounded-full bg-neutral-300"></span>
                <span>{{ t.location }}</span>
              </div>
              <div class="flex items-center gap-2 mt-3">
                <span class="px-2 py-1 bg-primary-50 text-primary-700 text-xs font-medium rounded-lg">{{ t.player_count }} players</span>
                <span v-if="t.start_date" class="text-xs text-neutral-400">{{ formatDate(t.start_date) }}</span>
              </div>
            </div>
          </NuxtLink>
        </div>
      </div>

      <!-- All Tournaments -->
      <div>
        <h2 v-if="!featuredOnly && featuredTournaments.length > 0" class="font-display text-2xl font-bold text-neutral-900 mb-6">All Tournaments</h2>
        
        <div v-if="tournaments.length === 0" class="text-center py-16">
          <div class="w-20 h-20 mx-auto bg-neutral-100 rounded-full flex items-center justify-center mb-4">
            <svg class="w-10 h-10 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-neutral-900">No tournaments found</h3>
          <p class="text-neutral-500 mt-1">Try adjusting your filters</p>
        </div>

        <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
          <NuxtLink 
            v-for="t in tournaments" 
            :key="t.id" 
            :to="`/tournaments/${t.id}`"
            class="group bg-white rounded-xl border border-neutral-200 overflow-hidden hover:shadow-lg hover:border-neutral-300 transition-all duration-200"
          >
            <div class="aspect-[16/10] bg-gradient-to-br from-neutral-200 to-neutral-300 overflow-hidden relative">
              <img 
                v-if="t.cover_image_url" 
                :src="t.cover_image_url" 
                :alt="t.name"
                class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
              />
              <div v-else class="w-full h-full flex items-center justify-center">
                <svg class="w-12 h-12 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 4v12l-4-2-4 2V4M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
              </div>
              <div v-if="t.featured" class="absolute top-2 left-2">
                <span class="px-2 py-0.5 bg-amber-500 text-white text-xs font-bold rounded">⭐</span>
              </div>
            </div>
            <div class="p-4">
              <h3 class="font-semibold text-neutral-900 group-hover:text-primary-600 transition-colors truncate">{{ t.name }}</h3>
              <div class="flex items-center gap-2 text-sm text-neutral-500 mt-1">
                <span>{{ t.year }}</span>
                <span v-if="t.location" class="w-1 h-1 rounded-full bg-neutral-300"></span>
                <span v-if="t.location" class="truncate">{{ t.location }}</span>
              </div>
              <div class="mt-2">
                <span class="text-xs text-primary-600 font-medium">{{ t.player_count }} players →</span>
              </div>
            </div>
          </NuxtLink>
        </div>
      </div>

      <!-- Load More -->
      <div v-if="hasMore" class="mt-10 text-center">
        <button 
          @click="loadMore" 
          :disabled="loadingMore"
          class="px-8 py-3 border-2 border-neutral-200 text-neutral-700 font-semibold rounded-xl hover:bg-neutral-50 hover:border-neutral-300 transition-all disabled:opacity-50"
        >
          {{ loadingMore ? 'Loading...' : 'Load More' }}
        </button>
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
  player_count: number
}

definePageMeta({ layout: 'default' })

useHead({
  title: 'Tournaments & Events - Unicorn Sport',
  meta: [{ name: 'description', content: 'Browse public football tournaments and discover talented players from competitions across Africa.' }]
})

const config = useRuntimeConfig()

const loading = ref(true)
const loadingMore = ref(false)
const tournaments = ref<Tournament[]>([])
const featuredTournaments = ref<Tournament[]>([])
const hasMore = ref(false)
const page = ref(1)
const featuredOnly = ref(false)
const selectedYear = ref('')

// Generate available years (current year back to 2020)
const currentYear = new Date().getFullYear()
const availableYears = Array.from({ length: currentYear - 2019 }, (_, i) => currentYear - i)

onMounted(() => {
  fetchTournaments()
})

async function fetchTournaments() {
  if (page.value === 1) loading.value = true
  
  try {
    let url = `/tournaments?page=${page.value}&limit=20`
    if (featuredOnly.value) url += '&featured=true'
    if (selectedYear.value) url += `&year=${selectedYear.value}`

    const response = await $fetch<ApiResponse<{ tournaments: Tournament[]; pagination: { total: number; total_pages: number } }>>(url, {
      baseURL: config.public.apiBase,
    })

    if (response.success && response.data) {
      if (page.value === 1) {
        tournaments.value = response.data.tournaments || []
        // Separate featured for first page
        if (!featuredOnly.value) {
          featuredTournaments.value = tournaments.value.filter(t => t.featured).slice(0, 3)
          tournaments.value = tournaments.value.filter(t => !t.featured || featuredTournaments.value.length === 0)
        }
      } else {
        tournaments.value.push(...(response.data.tournaments || []))
      }
      hasMore.value = page.value < (response.data.pagination?.total_pages || 1)
    }
  } catch (error) {
    console.error('Failed to fetch tournaments:', error)
  } finally {
    loading.value = false
    loadingMore.value = false
  }
}

function toggleFeatured() {
  featuredOnly.value = !featuredOnly.value
  page.value = 1
  featuredTournaments.value = []
  fetchTournaments()
}

async function loadMore() {
  loadingMore.value = true
  page.value++
  await fetchTournaments()
}

function formatDate(dateString: string): string {
  return new Date(dateString).toLocaleDateString('en-US', { month: 'short', year: 'numeric' })
}
</script>
