<template>
  <div class="min-h-screen bg-neutral-50">
    <!-- Hero Header -->
    <div class="relative bg-neutral-950 overflow-hidden">
      <!-- Background -->
      <div class="absolute inset-0">
        <div class="absolute inset-0 bg-gradient-to-br from-neutral-950 via-primary-950/30 to-neutral-950"></div>
        <div class="absolute top-0 left-1/2 -translate-x-1/2 w-[800px] h-[400px] bg-primary-500/10 rounded-full blur-[120px]"></div>
        <!-- Subtle pitch pattern -->
        <div class="absolute inset-0 opacity-[0.02]">
          <svg class="w-full h-full" viewBox="0 0 800 400" preserveAspectRatio="xMidYMid slice">
            <circle cx="400" cy="200" r="80" fill="none" stroke="white" stroke-width="2"/>
            <circle cx="400" cy="200" r="3" fill="white"/>
          </svg>
        </div>
      </div>
      
      <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pt-28 lg:pt-32 pb-16 lg:pb-20">
        <div class="max-w-2xl">
          <h1 class="font-display text-4xl lg:text-5xl font-bold text-white">
            Discover Players
          </h1>
          <p class="mt-4 text-lg text-neutral-300">
            Search through verified African football talent. Find your next signing.
          </p>
        </div>
        
        <!-- Quick stats -->
        <div class="mt-8 flex flex-wrap items-center gap-6">
          <div class="flex items-center gap-2">
            <div class="w-2 h-2 rounded-full bg-primary-500 animate-pulse"></div>
            <span class="text-sm text-neutral-400">
              <span class="text-white font-semibold">{{ searchStore.pagination.total || '500+' }}</span> players available
            </span>
          </div>
          <div class="flex items-center gap-2">
            <div class="w-2 h-2 rounded-full bg-secondary-500"></div>
            <span class="text-sm text-neutral-400">
              <span class="text-white font-semibold">15+</span> countries
            </span>
          </div>
          <div class="flex items-center gap-2">
            <div class="w-2 h-2 rounded-full bg-emerald-500"></div>
            <span class="text-sm text-neutral-400">
              <span class="text-white font-semibold">HD</span> video verified
            </span>
          </div>
        </div>
      </div>
    </div>

    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <!-- Search & Filters Card -->
      <div class="bg-white rounded-2xl shadow-sm border border-neutral-200/80 p-6 -mt-8 relative z-10 mb-8">
        <div class="flex flex-col lg:flex-row gap-4">
          <!-- Search Input -->
          <div class="flex-1">
            <div class="relative">
              <svg class="absolute left-4 top-1/2 -translate-y-1/2 h-5 w-5 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
              <input
                v-model="searchQuery"
                type="text"
                placeholder="Search by name, position, country..."
                class="w-full pl-12 pr-4 py-3.5 rounded-xl border border-neutral-200 bg-neutral-50 text-neutral-900 placeholder-neutral-400 transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500 focus:bg-white"
                @input="debouncedSearch"
              />
            </div>
          </div>

          <!-- Filters Toggle (Mobile) -->
          <button
            class="lg:hidden flex items-center justify-center gap-2 px-4 py-3 rounded-xl border border-neutral-200 text-neutral-700 font-medium hover:bg-neutral-50 transition-colors"
            @click="showFilters = !showFilters"
          >
            <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.293A1 1 0 013 6.586V4z" />
            </svg>
            Filters
            <span v-if="activeFiltersCount" class="w-5 h-5 rounded-full bg-primary-500 text-white text-xs flex items-center justify-center">
              {{ activeFiltersCount }}
            </span>
          </button>
        </div>

        <!-- Filter Options -->
        <div :class="['mt-6 grid grid-cols-2 lg:grid-cols-5 gap-4', showFilters ? 'block' : 'hidden lg:grid']">
          <!-- Position -->
          <div>
            <label class="block text-xs font-medium text-neutral-500 mb-1.5">Position</label>
            <select 
              v-model="filters.position" 
              class="w-full px-3 py-2.5 rounded-lg border border-neutral-200 bg-white text-neutral-700 text-sm focus:outline-none focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500"
              @change="handleSearch"
            >
              <option value="">All Positions</option>
              <option value="Goalkeeper">Goalkeeper</option>
              <option value="Defender">Defender</option>
              <option value="Midfielder">Midfielder</option>
              <option value="Forward">Forward</option>
            </select>
          </div>

          <!-- Country -->
          <div>
            <label class="block text-xs font-medium text-neutral-500 mb-1.5">Country</label>
            <select 
              v-model="filters.country" 
              class="w-full px-3 py-2.5 rounded-lg border border-neutral-200 bg-white text-neutral-700 text-sm focus:outline-none focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500"
              @change="handleSearch"
            >
              <option value="">All Countries</option>
              <option value="Nigeria">🇳🇬 Nigeria</option>
              <option value="Ghana">🇬🇭 Ghana</option>
              <option value="Kenya">🇰🇪 Kenya</option>
              <option value="South Africa">🇿🇦 South Africa</option>
              <option value="Cameroon">🇨🇲 Cameroon</option>
              <option value="Senegal">🇸🇳 Senegal</option>
              <option value="Egypt">🇪🇬 Egypt</option>
              <option value="Morocco">🇲🇦 Morocco</option>
            </select>
          </div>

          <!-- Age Range -->
          <div>
            <label class="block text-xs font-medium text-neutral-500 mb-1.5">Age Range</label>
            <select 
              v-model="ageRange" 
              class="w-full px-3 py-2.5 rounded-lg border border-neutral-200 bg-white text-neutral-700 text-sm focus:outline-none focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500"
              @change="handleAgeChange"
            >
              <option value="">All Ages</option>
              <option value="14-16">14-16 years</option>
              <option value="17-18">17-18 years</option>
              <option value="19-21">19-21 years</option>
              <option value="22-25">22-25 years</option>
            </select>
          </div>

          <!-- Verified Only -->
          <div class="flex items-end">
            <label class="flex items-center gap-3 cursor-pointer group py-2.5">
              <div class="relative">
                <input 
                  type="checkbox" 
                  v-model="filters.verified_only"
                  class="peer sr-only" 
                  @change="handleSearch"
                />
                <div class="w-5 h-5 rounded border-2 border-neutral-300 bg-white peer-checked:bg-primary-500 peer-checked:border-primary-500 transition-all duration-200 flex items-center justify-center group-hover:border-neutral-400">
                  <svg class="w-3 h-3 text-white opacity-0 peer-checked:opacity-100 transition-opacity" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="3">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7" />
                  </svg>
                </div>
              </div>
              <span class="text-sm text-neutral-700 font-medium">Verified only</span>
            </label>
          </div>

          <!-- Reset -->
          <div class="flex items-end">
            <button 
              v-if="activeFiltersCount > 0"
              class="flex items-center gap-2 px-4 py-2.5 text-sm text-primary-600 hover:text-primary-700 hover:bg-primary-50 rounded-lg font-medium transition-colors"
              @click="resetFilters"
            >
              <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
              Clear filters
            </button>
          </div>
        </div>
      </div>

      <!-- Results Header -->
      <div class="flex items-center justify-between mb-6">
        <div>
          <p v-if="searchStore.loading" class="text-neutral-500 flex items-center gap-2">
            <svg class="w-4 h-4 animate-spin" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
            </svg>
            Searching...
          </p>
          <p v-else class="text-neutral-600">
            <span class="font-semibold text-neutral-900">{{ searchStore.pagination.total }}</span> players found
          </p>
        </div>
        <select 
          class="px-3 py-2 rounded-lg border border-neutral-200 bg-white text-neutral-700 text-sm focus:outline-none focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500"
          @change="handleSearch"
        >
          <option value="newest">Most Recent</option>
          <option value="name_asc">Name A-Z</option>
          <option value="name_desc">Name Z-A</option>
        </select>
      </div>

      <!-- Loading State -->
      <div v-if="searchStore.loading && searchStore.results.length === 0" class="flex flex-col items-center justify-center py-20">
        <div class="w-12 h-12 border-4 border-primary-500/30 border-t-primary-500 rounded-full animate-spin mb-4"></div>
        <p class="text-neutral-500">Loading players...</p>
      </div>

      <!-- Results Grid -->
      <div v-else-if="searchStore.results.length > 0" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
        <PlayerCard v-for="player in searchStore.results" :key="player.id" :player="player" />
      </div>

      <!-- Empty State -->
      <div v-else class="text-center py-20">
        <div class="w-20 h-20 mx-auto rounded-full bg-neutral-100 flex items-center justify-center mb-6">
          <svg class="w-10 h-10 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </div>
        <h3 class="text-xl font-display font-bold text-neutral-900 mb-2">No players found</h3>
        <p class="text-neutral-500 mb-6 max-w-md mx-auto">
          We couldn't find any players matching your criteria. Try adjusting your filters.
        </p>
        <button 
          class="inline-flex items-center gap-2 px-5 py-2.5 bg-primary-600 text-white rounded-xl font-medium hover:bg-primary-700 transition-colors"
          @click="resetFilters"
        >
          <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
          Reset all filters
        </button>
      </div>

      <!-- Load More -->
      <div v-if="searchStore.pagination.page < searchStore.pagination.total_pages" class="mt-12 text-center">
        <button 
          :disabled="searchStore.loading"
          :class="[
            'inline-flex items-center gap-2 px-8 py-3.5 rounded-xl font-semibold transition-all duration-200',
            searchStore.loading 
              ? 'bg-neutral-100 text-neutral-400 cursor-not-allowed' 
              : 'bg-white border-2 border-neutral-200 text-neutral-700 hover:border-primary-500 hover:text-primary-600'
          ]"
          @click="loadMore"
        >
          <svg v-if="searchStore.loading" class="w-5 h-5 animate-spin" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
          </svg>
          <span>{{ searchStore.loading ? 'Loading...' : 'Load More Players' }}</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  layout: 'default',
})

const searchStore = useSearchStore()

const searchQuery = ref('')
const showFilters = ref(false)
const ageRange = ref('')
const filters = reactive({
  position: '',
  country: '',
  verified_only: false,
})

// Count active filters
const activeFiltersCount = computed(() => {
  let count = 0
  if (filters.position) count++
  if (filters.country) count++
  if (ageRange.value) count++
  if (filters.verified_only) count++
  if (searchQuery.value) count++
  return count
})

// Debounced search
let debounceTimer: ReturnType<typeof setTimeout>
const debouncedSearch = () => {
  clearTimeout(debounceTimer)
  debounceTimer = setTimeout(() => {
    searchStore.setQuery(searchQuery.value)
    handleSearch()
  }, 300)
}

const handleSearch = async () => {
  searchStore.setFilter('position', filters.position || undefined)
  searchStore.setFilter('country', filters.country || undefined)
  searchStore.setFilter('verified_only', filters.verified_only)
  await searchStore.search()
}

const handleAgeChange = () => {
  if (ageRange.value) {
    const [min, max] = ageRange.value.split('-').map(Number)
    searchStore.setFilter('age_min', min)
    searchStore.setFilter('age_max', max)
  } else {
    searchStore.setFilter('age_min', undefined)
    searchStore.setFilter('age_max', undefined)
  }
  handleSearch()
}

const resetFilters = () => {
  searchQuery.value = ''
  ageRange.value = ''
  filters.position = ''
  filters.country = ''
  filters.verified_only = false
  searchStore.resetFilters()
  searchStore.search()
}

const loadMore = () => {
  searchStore.loadMore()
}

// Initial search on mount
onMounted(() => {
  searchStore.search()
})
</script>
