<template>
  <div class="min-h-screen bg-neutral-50">
    <!-- Hero Header - Light and Inviting -->
    <div class="relative bg-white border-b border-neutral-100 overflow-hidden">
      <!-- Subtle Background Pattern -->
      <div class="absolute inset-0">
        <div class="absolute top-0 right-0 w-[500px] h-[300px] bg-gradient-to-bl from-primary-50/80 to-transparent rounded-full blur-3xl"></div>
        <div class="absolute bottom-0 left-0 w-[400px] h-[250px] bg-gradient-to-tr from-emerald-50/60 to-transparent rounded-full blur-3xl"></div>
      </div>
      
      <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pt-24 lg:pt-28 pb-8 lg:pb-10">
        <div class="flex flex-col lg:flex-row lg:items-end lg:justify-between gap-6">
          <!-- Title & Description -->
          <div class="max-w-xl">
            <div class="inline-flex items-center gap-2 px-3 py-1.5 bg-primary-50 border border-primary-100 rounded-full text-primary-700 text-sm font-medium mb-4">
              <div class="w-1.5 h-1.5 rounded-full bg-primary-500 animate-pulse"></div>
              Live Database
            </div>
            <h1 class="font-display text-3xl lg:text-4xl font-bold text-neutral-900">
              Discover Talent
            </h1>
            <p class="mt-3 text-neutral-600 text-lg">
              Explore verified African football talent. Filter by position, age, and country to find your next signing.
            </p>
          </div>
          
          <!-- Quick Stats Pills -->
          <div class="grid grid-cols-3 gap-2 sm:flex sm:flex-wrap sm:items-center sm:gap-3 w-full lg:w-auto">
            <div class="flex items-center gap-2 px-3 sm:px-4 py-2 sm:py-2.5 bg-white rounded-xl shadow-sm border border-neutral-200">
              <div class="w-7 h-7 sm:w-8 sm:h-8 rounded-lg bg-primary-100 flex items-center justify-center flex-shrink-0">
                <svg class="w-3.5 h-3.5 sm:w-4 sm:h-4 text-primary-600" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
              </div>
              <div class="min-w-0">
                <p class="text-[10px] sm:text-xs text-neutral-500">Players</p>
                <p class="font-bold text-sm sm:text-base text-neutral-900">{{ animatedTotal }}</p>
              </div>
            </div>
            <div class="flex items-center gap-2 px-3 sm:px-4 py-2 sm:py-2.5 bg-white rounded-xl shadow-sm border border-neutral-200">
              <div class="w-7 h-7 sm:w-8 sm:h-8 rounded-lg bg-emerald-100 flex items-center justify-center flex-shrink-0">
                <svg class="w-3.5 h-3.5 sm:w-4 sm:h-4 text-emerald-600" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
                </svg>
              </div>
              <div class="min-w-0">
                <p class="text-[10px] sm:text-xs text-neutral-500">Verified</p>
                <p class="font-bold text-sm sm:text-base text-neutral-900">100%</p>
              </div>
            </div>
            <div class="flex items-center gap-2 px-3 sm:px-4 py-2 sm:py-2.5 bg-white rounded-xl shadow-sm border border-neutral-200">
              <div class="w-7 h-7 sm:w-8 sm:h-8 rounded-lg bg-secondary-100 flex items-center justify-center flex-shrink-0">
                <span class="text-xs sm:text-sm">🌍</span>
              </div>
              <div class="min-w-0">
                <p class="text-[10px] sm:text-xs text-neutral-500">Countries</p>
                <p class="font-bold text-sm sm:text-base text-neutral-900">{{ uniqueCountriesCount }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
      <!-- Search & Filters Card - Elevated -->
      <div class="bg-white rounded-2xl shadow-lg shadow-neutral-200/50 border border-neutral-100 p-5 lg:p-6 -mt-6 relative z-10 mb-8">
        <!-- Search Row -->
        <div class="flex flex-col lg:flex-row gap-3">
          <!-- Search Input -->
          <div class="flex-1 relative">
            <div class="absolute left-4 top-1/2 -translate-y-1/2 w-10 h-10 rounded-lg bg-neutral-100 flex items-center justify-center">
              <svg class="h-5 w-5 text-neutral-500" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </div>
            <input
              v-model="searchQuery"
              type="text"
              placeholder="Search players by name, position, or country..."
              class="w-full pl-16 pr-4 py-4 rounded-xl border border-neutral-200 bg-neutral-50/50 text-neutral-900 placeholder-neutral-400 transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500 focus:bg-white text-base"
              @input="debouncedSearch"
            />
            <!-- Clear button -->
            <button
              v-if="searchQuery"
              @click="searchQuery = ''; debouncedSearch()"
              class="absolute right-4 top-1/2 -translate-y-1/2 p-1.5 rounded-lg text-neutral-400 hover:text-neutral-600 hover:bg-neutral-100 transition-colors"
            >
              <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>

          <!-- Filter Toggle (Mobile) -->
          <button
            class="lg:hidden flex items-center justify-center gap-2 px-5 py-4 rounded-xl border-2 border-neutral-200 text-neutral-700 font-semibold hover:bg-neutral-50 hover:border-neutral-300 transition-all"
            @click="showFilters = !showFilters"
          >
            <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.293A1 1 0 013 6.586V4z" />
            </svg>
            <span>Filters</span>
            <span v-if="activeFiltersCount" class="w-6 h-6 rounded-full bg-primary-500 text-white text-xs font-bold flex items-center justify-center">
              {{ activeFiltersCount }}
            </span>
          </button>
        </div>

        <!-- Filter Options - Better organized grid -->
        <div :class="['mt-5 pt-5 border-t border-neutral-100', showFilters ? 'block' : 'hidden lg:block']">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-sm font-semibold text-neutral-700 flex items-center gap-2">
              <svg class="w-4 h-4 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4" />
              </svg>
              Filter by
            </h3>
            <button 
              v-if="activeFiltersCount > 0"
              class="flex items-center gap-1.5 px-3 py-1.5 text-sm text-primary-600 hover:text-primary-700 hover:bg-primary-50 rounded-lg font-medium transition-colors"
              @click="resetFilters"
            >
              <svg class="w-3.5 h-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
              </svg>
              Reset all
            </button>
          </div>
          
          <div class="grid grid-cols-2 lg:grid-cols-5 gap-3 lg:gap-4">
            <!-- Position -->
            <div class="space-y-1.5">
              <label class="block text-xs font-medium text-neutral-500 uppercase tracking-wider">Position</label>
              <select 
                v-model="filters.position" 
                class="w-full px-3 py-2.5 rounded-lg border border-neutral-200 bg-white text-neutral-700 text-sm focus:outline-none focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500 appearance-none cursor-pointer"
                @change="handleSearch"
              >
                <option value="">All Positions</option>
                <option value="Goalkeeper">🧤 Goalkeeper</option>
                <option value="Defender">🛡️ Defender</option>
                <option value="Midfielder">⚡ Midfielder</option>
                <option value="Forward">⚽ Forward</option>
              </select>
            </div>

            <!-- Country -->
            <div class="space-y-1.5">
              <label class="block text-xs font-medium text-neutral-500 uppercase tracking-wider">Country</label>
              <select 
                v-model="filters.country" 
                class="w-full px-3 py-2.5 rounded-lg border border-neutral-200 bg-white text-neutral-700 text-sm focus:outline-none focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500 appearance-none cursor-pointer"
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

            <!-- Academy -->
            <div class="space-y-1.5">
              <label class="block text-xs font-medium text-neutral-500 uppercase tracking-wider">Academy</label>
              <select 
                v-model="filters.academy_id" 
                class="w-full px-3 py-2.5 rounded-lg border border-neutral-200 bg-white text-neutral-700 text-sm focus:outline-none focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500 appearance-none cursor-pointer"
                @change="handleSearch"
                :disabled="loadingAcademies"
              >
                <option value="">All Academies</option>
                <option v-for="academy in academies" :key="academy.id" :value="academy.id">
                  {{ academy.name }}
                  <template v-if="academy.is_verified"> ✓</template>
                </option>
              </select>
            </div>

            <!-- Age Range -->
            <div class="space-y-1.5">
              <label class="block text-xs font-medium text-neutral-500 uppercase tracking-wider">Age Range</label>
              <select 
                v-model="ageRange" 
                class="w-full px-3 py-2.5 rounded-lg border border-neutral-200 bg-white text-neutral-700 text-sm focus:outline-none focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500 appearance-none cursor-pointer"
                @change="handleAgeChange"
              >
                <option value="">All Ages</option>
                <option value="14-16">14-16 years</option>
                <option value="17-18">17-18 years</option>
                <option value="19-21">19-21 years</option>
                <option value="22-25">22-25 years</option>
              </select>
            </div>

            <!-- Verified Only Toggle -->
            <div class="space-y-1.5">
              <label class="block text-xs font-medium text-neutral-500 uppercase tracking-wider">Status</label>
              <label class="flex items-center gap-3 px-3 py-2.5 rounded-lg border border-neutral-200 bg-white cursor-pointer hover:border-primary-300 transition-colors">
                <div class="relative">
                  <input 
                    type="checkbox" 
                    v-model="filters.verified_only"
                    class="peer sr-only" 
                    @change="handleSearch"
                  />
                  <div class="w-9 h-5 rounded-full bg-neutral-200 peer-checked:bg-primary-500 transition-colors"></div>
                  <div class="absolute left-0.5 top-0.5 w-4 h-4 rounded-full bg-white shadow-sm transition-transform peer-checked:translate-x-4"></div>
                </div>
                <span class="text-sm text-neutral-700 font-medium">Verified only</span>
              </label>
            </div>
          </div>
        </div>
      </div>

      <!-- Results Header - Clean -->
      <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4 mb-6">
        <div>
          <p v-if="searchStore.loading && searchStore.results.length === 0" class="text-neutral-500 flex items-center gap-2">
            <svg class="w-4 h-4 animate-spin" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
            </svg>
            Searching talent pool...
          </p>
          <p v-else class="text-neutral-600">
            Showing <span class="font-bold text-neutral-900">{{ searchStore.results.length }}</span> of 
            <span class="font-bold text-neutral-900">{{ searchStore.pagination.total }}</span> players
            <span v-if="activeFiltersCount > 0" class="text-primary-600">(filtered)</span>
          </p>
        </div>
        
        <div class="flex items-center gap-2">
          <span class="text-sm text-neutral-500">Sort by:</span>
          <select 
            v-model="sortOrder"
            class="px-3 py-2 rounded-lg border border-neutral-200 bg-white text-neutral-700 text-sm font-medium focus:outline-none focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500"
            @change="handleSearch"
          >
            <option value="newest">Most Recent</option>
            <option value="name_asc">Name A-Z</option>
            <option value="name_desc">Name Z-A</option>
          </select>
        </div>
      </div>

      <!-- Loading State - Skeleton Grid -->
      <div v-if="searchStore.loading && searchStore.results.length === 0" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
        <div v-for="n in 8" :key="n" class="bg-white rounded-2xl overflow-hidden border border-neutral-100 animate-pulse">
          <div class="aspect-[3/4] bg-neutral-200"></div>
          <div class="p-4 space-y-3">
            <div class="h-5 bg-neutral-200 rounded w-3/4"></div>
            <div class="h-4 bg-neutral-100 rounded w-1/2"></div>
            <div class="h-4 bg-neutral-100 rounded w-2/3"></div>
          </div>
        </div>
      </div>

      <!-- Results Grid -->
      <div v-else-if="searchStore.results.length > 0" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
        <PlayerCard v-for="player in searchStore.results" :key="player.id" :player="player" />
      </div>

      <!-- Empty State - Friendly & Helpful -->
      <div v-else class="text-center py-16 lg:py-20">
        <div class="w-24 h-24 mx-auto rounded-3xl bg-gradient-to-br from-neutral-100 to-neutral-50 flex items-center justify-center mb-6 shadow-inner">
          <svg class="w-12 h-12 text-neutral-300" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.5">
            <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </div>
        <h3 class="text-xl font-display font-bold text-neutral-900 mb-2">No players match your criteria</h3>
        <p class="text-neutral-500 mb-8 max-w-md mx-auto">
          Try adjusting your filters or search terms to discover more talent.
        </p>
        <div class="flex flex-col sm:flex-row items-center justify-center gap-3">
          <button 
            class="inline-flex items-center gap-2 px-6 py-3 bg-primary-600 text-white rounded-xl font-semibold hover:bg-primary-700 transition-colors shadow-lg shadow-primary-500/25"
            @click="resetFilters"
          >
            <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
            </svg>
            Reset all filters
          </button>
          <NuxtLink
            to="/"
            class="inline-flex items-center gap-2 px-6 py-3 bg-white border-2 border-neutral-200 text-neutral-700 rounded-xl font-semibold hover:border-neutral-300 hover:bg-neutral-50 transition-colors"
          >
            <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
            </svg>
            Back to home
          </NuxtLink>
        </div>
      </div>

      <!-- Load More - Elegant -->
      <div v-if="searchStore.results.length > 0 && searchStore.pagination.page < searchStore.pagination.total_pages" class="mt-12 text-center">
        <div class="inline-flex flex-col items-center gap-4">
          <p class="text-sm text-neutral-500">
            Showing {{ searchStore.results.length }} of {{ searchStore.pagination.total }} players
          </p>
          <button 
            :disabled="searchStore.loading"
            :class="[
              'inline-flex items-center gap-2 px-8 py-4 rounded-2xl font-semibold transition-all duration-200',
              searchStore.loading 
                ? 'bg-neutral-100 text-neutral-400 cursor-not-allowed' 
                : 'bg-white border-2 border-neutral-200 text-neutral-700 hover:border-primary-500 hover:text-primary-600 hover:shadow-lg hover:shadow-primary-500/10'
            ]"
            @click="loadMore"
          >
            <svg v-if="searchStore.loading" class="w-5 h-5 animate-spin" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
            </svg>
            <svg v-else class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7" />
            </svg>
            <span>{{ searchStore.loading ? 'Loading...' : 'Load more players' }}</span>
          </button>
        </div>
      </div>
      
      <!-- End of results indicator -->
      <div v-else-if="searchStore.results.length > 0 && searchStore.pagination.page >= searchStore.pagination.total_pages" class="mt-12 text-center">
        <p class="text-sm text-neutral-400 flex items-center justify-center gap-2">
          <span class="w-8 h-px bg-neutral-200"></span>
          You've seen all {{ searchStore.pagination.total }} players
          <span class="w-8 h-px bg-neutral-200"></span>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  layout: 'default',
  keepalive: false,
})

const searchStore = useSearchStore()

const searchQuery = ref('')
const showFilters = ref(false)
const ageRange = ref('')
const sortOrder = ref('newest')
const filters = reactive({
  position: '',
  country: '',
  academy_id: '',
  verified_only: false,
})

// Academy list for filter dropdown
interface Academy {
  id: string
  name: string
  country?: string
  is_verified: boolean
}
const academies = ref<Academy[]>([])
const loadingAcademies = ref(false)

// Fetch academies on mount
const fetchAcademies = async () => {
  loadingAcademies.value = true
  try {
    const api = useApi()
    const response = await api.get<{ success: boolean; data: { academies: Academy[] } }>('/academies')
    if (response.success && response.data) {
      academies.value = response.data.academies || []
    }
  } catch (error) {
    console.error('Failed to fetch academies:', error)
  } finally {
    loadingAcademies.value = false
  }
}

// Animated total count for visual engagement
const animatedTotal = ref(0)
watch(() => searchStore.pagination.total, (newTotal) => {
  if (newTotal === 0) {
    animatedTotal.value = 0
    return
  }
  // Quick animate to the new total
  const duration = 500
  const steps = 20
  const increment = newTotal / steps
  let current = 0
  const timer = setInterval(() => {
    current += increment
    if (current >= newTotal) {
      animatedTotal.value = newTotal
      clearInterval(timer)
    } else {
      animatedTotal.value = Math.floor(current)
    }
  }, duration / steps)
}, { immediate: true })

// Count unique countries from the results (derived from data)
const uniqueCountriesCount = computed(() => {
  if (searchStore.results.length === 0) return 0
  const countries = new Set(searchStore.results.map(p => p.country).filter(Boolean))
  return countries.size
})

// Count active filters
const activeFiltersCount = computed(() => {
  let count = 0
  if (filters.position) count++
  if (filters.country) count++
  if (filters.academy_id) count++
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
  searchStore.setFilter('academy_id', filters.academy_id || undefined)
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
  filters.academy_id = ''
  filters.verified_only = false
  searchStore.resetFilters()
  searchStore.search()
}

const loadMore = () => {
  searchStore.loadMore()
}

// Initial search on mount - non-blocking
onMounted(async () => {
  // Fetch academies for filter dropdown
  fetchAcademies()
  
  // Reset local state
  searchQuery.value = ''
  ageRange.value = ''
  filters.position = ''
  filters.country = ''
  filters.academy_id = ''
  filters.verified_only = false
  // Reset store and search
  searchStore.resetFilters()
  await searchStore.search()
})

// Refresh on route changes to this page
const route = useRoute()
watch(() => route.fullPath, async (newPath, oldPath) => {
  if (newPath.startsWith('/discover') && oldPath && !oldPath.startsWith('/discover')) {
    searchQuery.value = ''
    ageRange.value = ''
    filters.position = ''
    filters.country = ''
    filters.academy_id = ''
    filters.verified_only = false
    searchStore.resetFilters()
    await searchStore.search()
  }
})
</script>
