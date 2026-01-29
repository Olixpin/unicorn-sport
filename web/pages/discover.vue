<template>
  <div class="min-h-screen bg-neutral-50">
    <!-- Compact Header -->
    <div class="bg-white border-b border-neutral-200">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pt-16 lg:pt-20 pb-3">
        <div class="flex items-center justify-between">
          <h1 class="font-display text-xl lg:text-2xl font-bold text-neutral-900">
            Discover Players
          </h1>
          <NuxtLink 
            to="/dashboard" 
            class="flex items-center gap-1.5 px-3 py-1.5 bg-neutral-900 text-white rounded-lg text-sm font-medium hover:bg-neutral-800 transition-colors"
          >
            <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM9.555 7.168A1 1 0 008 8v4a1 1 0 001.555.832l3-2a1 1 0 000-1.664l-3-2z" clip-rule="evenodd"/>
            </svg>
            Watch Feed
          </NuxtLink>
        </div>
      </div>
    </div>

    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-3">
      <!-- Search & Filters - Compact -->
      <div class="bg-white rounded-xl shadow-sm border border-neutral-200 p-3 mb-4">
        <!-- Search Row -->
        <div class="flex gap-2">
          <!-- Search Input -->
          <div class="flex-1 relative">
            <svg class="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
            <input
              v-model="searchQuery"
              type="text"
              placeholder="Search..."
              class="w-full pl-9 pr-3 py-2 rounded-lg border border-neutral-200 text-neutral-900 placeholder-neutral-400 text-sm focus:outline-none focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500"
              @input="debouncedSearch"
            />
          </div>

          <!-- Quick Filters - Always Visible -->
          <select 
            v-model="filters.position" 
            class="px-2 py-2 pr-7 rounded-lg border border-neutral-200 bg-white text-neutral-700 text-sm focus:outline-none focus:border-primary-500 appearance-none bg-[url('data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20fill%3D%22none%22%20viewBox%3D%220%200%2020%2020%22%3E%3Cpath%20stroke%3D%22%236b7280%22%20stroke-linecap%3D%22round%22%20stroke-linejoin%3D%22round%22%20stroke-width%3D%221.5%22%20d%3D%22m6%208%204%204%204-4%22%2F%3E%3C%2Fsvg%3E')] bg-[length:1.25rem] bg-[right_0.25rem_center] bg-no-repeat"
            @change="handleSearch"
          >
            <option value="">Position</option>
            <option value="Goalkeeper">GK</option>
            <option value="Defender">DEF</option>
            <option value="Midfielder">MID</option>
            <option value="Forward">FWD</option>
          </select>

          <select 
            v-model="filters.country" 
            class="hidden sm:block px-2 py-2 pr-7 rounded-lg border border-neutral-200 bg-white text-neutral-700 text-sm focus:outline-none focus:border-primary-500 appearance-none bg-[url('data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20fill%3D%22none%22%20viewBox%3D%220%200%2020%2020%22%3E%3Cpath%20stroke%3D%22%236b7280%22%20stroke-linecap%3D%22round%22%20stroke-linejoin%3D%22round%22%20stroke-width%3D%221.5%22%20d%3D%22m6%208%204%204%204-4%22%2F%3E%3C%2Fsvg%3E')] bg-[length:1.25rem] bg-[right_0.25rem_center] bg-no-repeat"
            @change="handleSearch"
          >
            <option value="">Country</option>
            <option value="Nigeria">Nigeria</option>
            <option value="Ghana">Ghana</option>
            <option value="Kenya">Kenya</option>
            <option value="South Africa">S. Africa</option>
            <option value="Cameroon">Cameroon</option>
            <option value="Senegal">Senegal</option>
          </select>

          <select 
            v-model="ageRange" 
            class="hidden md:block px-2 py-2 pr-7 rounded-lg border border-neutral-200 bg-white text-neutral-700 text-sm focus:outline-none focus:border-primary-500 appearance-none bg-[url('data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20fill%3D%22none%22%20viewBox%3D%220%200%2020%2020%22%3E%3Cpath%20stroke%3D%22%236b7280%22%20stroke-linecap%3D%22round%22%20stroke-linejoin%3D%22round%22%20stroke-width%3D%221.5%22%20d%3D%22m6%208%204%204%204-4%22%2F%3E%3C%2Fsvg%3E')] bg-[length:1.25rem] bg-[right_0.25rem_center] bg-no-repeat"
            @change="handleAgeChange"
          >
            <option value="">Age</option>
            <option value="14-16">14-16</option>
            <option value="17-18">17-18</option>
            <option value="19-21">19-21</option>
            <option value="22-25">22+</option>
          </select>

          <button 
            v-if="activeFiltersCount > 0"
            class="hidden sm:block px-2 py-2 text-sm text-primary-600 hover:text-primary-700 font-medium whitespace-nowrap"
            @click="resetFilters"
          >
            Clear
          </button>
        </div>

        <!-- Mobile Extra Filters Row -->
        <div class="flex sm:hidden gap-2 mt-2">
          <select 
            v-model="filters.country" 
            class="flex-1 px-2 py-2 pr-7 rounded-lg border border-neutral-200 bg-white text-neutral-700 text-sm appearance-none bg-[url('data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20fill%3D%22none%22%20viewBox%3D%220%200%2020%2020%22%3E%3Cpath%20stroke%3D%22%236b7280%22%20stroke-linecap%3D%22round%22%20stroke-linejoin%3D%22round%22%20stroke-width%3D%221.5%22%20d%3D%22m6%208%204%204%204-4%22%2F%3E%3C%2Fsvg%3E')] bg-[length:1.25rem] bg-[right_0.25rem_center] bg-no-repeat"
            @change="handleSearch"
          >
            <option value="">Country</option>
            <option value="Nigeria">Nigeria</option>
            <option value="Ghana">Ghana</option>
            <option value="Kenya">Kenya</option>
          </select>
          <select 
            v-model="ageRange" 
            class="flex-1 px-2 py-2 pr-7 rounded-lg border border-neutral-200 bg-white text-neutral-700 text-sm appearance-none bg-[url('data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20fill%3D%22none%22%20viewBox%3D%220%200%2020%2020%22%3E%3Cpath%20stroke%3D%22%236b7280%22%20stroke-linecap%3D%22round%22%20stroke-linejoin%3D%22round%22%20stroke-width%3D%221.5%22%20d%3D%22m6%208%204%204%204-4%22%2F%3E%3C%2Fsvg%3E')] bg-[length:1.25rem] bg-[right_0.25rem_center] bg-no-repeat"
            @change="handleAgeChange"
          >
            <option value="">Age</option>
            <option value="14-16">14-16</option>
            <option value="17-18">17-18</option>
            <option value="19-21">19-21</option>
          </select>
          <button 
            v-if="activeFiltersCount > 0"
            class="px-3 py-2 text-sm text-primary-600 font-medium"
            @click="resetFilters"
          >
            Clear
          </button>
        </div>
      </div>

      <!-- Results Count -->
      <div class="flex items-center justify-between mb-3">
        <p class="text-sm text-neutral-500">
          <span v-if="searchStore.loading" class="flex items-center gap-2">
            <svg class="w-3 h-3 animate-spin" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
            </svg>
            Loading...
          </span>
          <span v-else>
            <span class="font-semibold text-neutral-900">{{ searchStore.pagination.total }}</span> players
          </span>
        </p>
        <select 
          v-model="sortOrder"
          class="px-2 py-1 pr-6 rounded-md border border-neutral-200 text-neutral-600 text-xs focus:outline-none focus:border-primary-500 appearance-none bg-[url('data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20fill%3D%22none%22%20viewBox%3D%220%200%2020%2020%22%3E%3Cpath%20stroke%3D%22%236b7280%22%20stroke-linecap%3D%22round%22%20stroke-linejoin%3D%22round%22%20stroke-width%3D%221.5%22%20d%3D%22m6%208%204%204%204-4%22%2F%3E%3C%2Fsvg%3E')] bg-[length:1rem] bg-[right_0.25rem_center] bg-no-repeat"
          @change="handleSearch"
        >
          <option value="newest">Newest</option>
          <option value="name_asc">A-Z</option>
        </select>
      </div>

      <!-- Loading State -->
      <div v-if="searchStore.loading && searchStore.results.length === 0" class="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4">
        <div v-for="n in 10" :key="n" class="bg-white rounded-lg overflow-hidden border border-neutral-200 animate-pulse">
          <div class="aspect-[4/5] bg-neutral-200"></div>
          <div class="p-3 space-y-2">
            <div class="h-4 bg-neutral-200 rounded w-3/4"></div>
            <div class="h-3 bg-neutral-100 rounded w-1/2"></div>
          </div>
        </div>
      </div>

      <!-- Results Grid -->
      <div v-else-if="searchStore.results.length > 0" class="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4">
        <PlayerCard v-for="player in searchStore.results" :key="player.id" :player="player" />
      </div>

      <!-- Empty State -->
      <div v-else class="text-center py-12">
        <div class="w-16 h-16 mx-auto rounded-full bg-neutral-100 flex items-center justify-center mb-4">
          <svg class="w-8 h-8 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.5">
            <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </div>
        <h3 class="text-lg font-semibold text-neutral-900 mb-1">No players found</h3>
        <p class="text-neutral-500 text-sm mb-4">Try adjusting your filters</p>
        <button 
          class="text-sm text-primary-600 font-medium hover:text-primary-700"
          @click="resetFilters"
        >
          Clear all filters
        </button>
      </div>

      <!-- Load More -->
      <div v-if="searchStore.results.length > 0 && searchStore.pagination.page < searchStore.pagination.total_pages" class="mt-8 text-center">
        <button 
          :disabled="searchStore.loading"
          class="px-6 py-2.5 rounded-lg border border-neutral-200 text-neutral-700 text-sm font-medium hover:bg-neutral-50 disabled:opacity-50"
          @click="loadMore"
        >
          {{ searchStore.loading ? 'Loading...' : 'Load more' }}
        </button>
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
const ageRange = ref('')
const sortOrder = ref('newest')
const filters = reactive({
  position: '',
  country: '',
})

// Count active filters
const activeFiltersCount = computed(() => {
  let count = 0
  if (filters.position) count++
  if (filters.country) count++
  if (ageRange.value) count++
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
  searchStore.resetFilters()
  searchStore.search()
}

const loadMore = () => {
  searchStore.loadMore()
}

// Initial search on mount
onMounted(async () => {
  searchQuery.value = ''
  ageRange.value = ''
  filters.position = ''
  filters.country = ''
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
