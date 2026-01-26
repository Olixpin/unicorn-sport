<template>
  <div class="max-w-7xl mx-auto">
    <!-- Page Header -->
    <div class="mb-8">
      <div class="flex flex-col lg:flex-row lg:items-center lg:justify-between gap-4">
        <div>
          <h1 class="font-display text-2xl lg:text-3xl font-bold text-neutral-900">
            Players
          </h1>
          <p class="mt-1 text-neutral-500">Manage all players in the platform</p>
        </div>
        <NuxtLink to="/admin/players/new">
          <button class="flex items-center gap-2 px-5 py-2.5 bg-gradient-to-r from-primary-500 to-primary-600 text-white font-semibold rounded-xl hover:from-primary-600 hover:to-primary-700 shadow-lg shadow-primary-500/25 transition-all duration-200">
            <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
            </svg>
            Add Player
          </button>
        </NuxtLink>
      </div>
    </div>

    <!-- Stats Summary -->
    <div class="grid grid-cols-1 sm:grid-cols-3 gap-4 mb-6">
      <div class="bg-white rounded-xl border border-neutral-200 p-4 flex items-center gap-4">
        <div class="w-10 h-10 bg-primary-100 rounded-lg flex items-center justify-center">
          <svg class="w-5 h-5 text-primary-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
        </div>
        <div>
          <p class="text-2xl font-bold text-neutral-900">{{ totalPlayers }}</p>
          <p class="text-sm text-neutral-500">Total Players</p>
        </div>
      </div>

      <div class="bg-white rounded-xl border border-neutral-200 p-4 flex items-center gap-4">
        <div class="w-10 h-10 bg-emerald-100 rounded-lg flex items-center justify-center">
          <svg class="w-5 h-5 text-emerald-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
        </div>
        <div>
          <p class="text-2xl font-bold text-neutral-900">{{ verifiedCount }}</p>
          <p class="text-sm text-neutral-500">Verified</p>
        </div>
      </div>

      <div class="bg-white rounded-xl border border-neutral-200 p-4 flex items-center gap-4">
        <div class="w-10 h-10 bg-amber-100 rounded-lg flex items-center justify-center">
          <svg class="w-5 h-5 text-amber-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
        </div>
        <div>
          <p class="text-2xl font-bold text-neutral-900">{{ pendingCount }}</p>
          <p class="text-sm text-neutral-500">Pending Review</p>
        </div>
      </div>
    </div>

    <!-- Filters & Search -->
    <div class="bg-white rounded-2xl border border-neutral-200 p-4 mb-6">
      <div class="flex flex-col lg:flex-row gap-4">
        <!-- Search -->
        <div class="flex-1 relative">
          <svg class="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
          <input
            v-model="filters.search"
            type="text"
            placeholder="Search by name, position, country..."
            class="w-full pl-10 pr-4 py-2.5 bg-neutral-50 border-0 rounded-xl text-sm placeholder-neutral-500 focus:outline-none focus:ring-2 focus:ring-primary-500/20 focus:bg-white transition-all"
            @input="debouncedSearch"
          />
        </div>

        <!-- Status Filter -->
        <div class="relative">
          <select
            v-model="filters.status"
            class="appearance-none w-full lg:w-40 px-4 py-2.5 bg-neutral-50 border-0 rounded-xl text-sm text-neutral-700 focus:outline-none focus:ring-2 focus:ring-primary-500/20 cursor-pointer"
            @change="fetchPlayers"
          >
            <option value="">All Status</option>
            <option value="verified">Verified</option>
            <option value="pending">Pending</option>
          </select>
          <svg class="absolute right-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400 pointer-events-none" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
          </svg>
        </div>

        <!-- Academy Filter -->
        <div class="relative">
          <select
            v-model="filters.academy"
            class="appearance-none w-full lg:w-48 px-4 py-2.5 bg-neutral-50 border-0 rounded-xl text-sm text-neutral-700 focus:outline-none focus:ring-2 focus:ring-primary-500/20 cursor-pointer"
            @change="fetchPlayers"
          >
            <option value="">All Academies</option>
            <option v-for="academy in academies" :key="academy.id" :value="academy.id">
              {{ academy.name }}
            </option>
          </select>
          <svg class="absolute right-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400 pointer-events-none" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
          </svg>
        </div>

        <!-- Clear Filters -->
        <button
          v-if="hasActiveFilters"
          @click="clearFilters"
          class="flex items-center gap-2 px-4 py-2.5 text-sm text-neutral-600 hover:text-neutral-900 hover:bg-neutral-100 rounded-xl transition-colors"
        >
          <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
          Clear
        </button>
      </div>
    </div>

    <!-- Players Table -->
    <div class="bg-white rounded-2xl border border-neutral-200 overflow-hidden">
      <!-- Loading State -->
      <div v-if="loading" class="flex flex-col items-center justify-center py-16">
        <div class="w-12 h-12 border-4 border-primary-500/30 border-t-primary-500 rounded-full animate-spin mb-4"></div>
        <p class="text-neutral-500">Loading players...</p>
      </div>

      <!-- Empty State -->
      <div v-else-if="players.length === 0" class="flex flex-col items-center justify-center py-16">
        <div class="w-20 h-20 bg-neutral-100 rounded-full flex items-center justify-center mb-6">
          <svg class="w-10 h-10 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
        </div>
        <h3 class="text-lg font-semibold text-neutral-900 mb-2">No players found</h3>
        <p class="text-neutral-500 mb-6 text-center max-w-sm">
          {{ hasActiveFilters ? 'Try adjusting your filters to find what you\'re looking for.' : 'Get started by adding your first player to the platform.' }}
        </p>
        <NuxtLink v-if="!hasActiveFilters" to="/admin/players/new">
          <button class="flex items-center gap-2 px-5 py-2.5 bg-primary-500 text-white font-semibold rounded-xl hover:bg-primary-600 transition-colors">
            <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
            </svg>
            Add First Player
          </button>
        </NuxtLink>
      </div>

      <!-- Table -->
      <div v-else class="overflow-x-auto">
        <table class="w-full">
          <thead>
            <tr class="border-b border-neutral-200 bg-neutral-50">
              <th class="px-6 py-4 text-left text-xs font-semibold text-neutral-500 uppercase tracking-wider">Player</th>
              <th class="px-6 py-4 text-left text-xs font-semibold text-neutral-500 uppercase tracking-wider">Position</th>
              <th class="px-6 py-4 text-left text-xs font-semibold text-neutral-500 uppercase tracking-wider hidden lg:table-cell">Academy</th>
              <th class="px-6 py-4 text-left text-xs font-semibold text-neutral-500 uppercase tracking-wider">Status</th>
              <th class="px-6 py-4 text-left text-xs font-semibold text-neutral-500 uppercase tracking-wider hidden md:table-cell">Videos</th>
              <th class="px-6 py-4 text-right text-xs font-semibold text-neutral-500 uppercase tracking-wider">Actions</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-neutral-100">
            <tr 
              v-for="player in players" 
              :key="player.id" 
              class="group hover:bg-neutral-50 transition-colors"
            >
              <td class="px-6 py-4">
                <div class="flex items-center gap-4">
                  <div class="w-12 h-12 rounded-xl bg-gradient-to-br from-neutral-100 to-neutral-200 overflow-hidden flex-shrink-0">
                    <img 
                      v-if="player.profile_photo_url"
                      :src="player.profile_photo_url"
                      :alt="player.first_name"
                      class="w-full h-full object-cover"
                    />
                    <div v-else class="w-full h-full flex items-center justify-center text-neutral-400 font-semibold">
                      {{ player.first_name?.charAt(0) }}{{ player.last_name?.charAt(0) }}
                    </div>
                  </div>
                  <div>
                    <p class="font-medium text-neutral-900">{{ player.first_name }} {{ player.last_name }}</p>
                    <p class="text-sm text-neutral-500">{{ player.country }}</p>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4">
                <span class="inline-flex items-center px-2.5 py-1 rounded-lg bg-neutral-100 text-sm font-medium text-neutral-700">
                  {{ player.position || 'N/A' }}
                </span>
              </td>
              <td class="px-6 py-4 hidden lg:table-cell">
                <span class="text-sm text-neutral-600">{{ player.academy_name || 'Unassigned' }}</span>
              </td>
              <td class="px-6 py-4">
                <span 
                  :class="[
                    'inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full text-xs font-medium',
                    player.is_verified 
                      ? 'bg-emerald-100 text-emerald-700' 
                      : 'bg-amber-100 text-amber-700'
                  ]"
                >
                  <span :class="['w-1.5 h-1.5 rounded-full', player.is_verified ? 'bg-emerald-500' : 'bg-amber-500']"></span>
                  {{ player.is_verified ? 'Verified' : 'Pending' }}
                </span>
              </td>
              <td class="px-6 py-4 hidden md:table-cell">
                <span class="text-sm text-neutral-600">{{ player.video_count || 0 }}</span>
              </td>
              <td class="px-6 py-4">
                <div class="flex items-center justify-end gap-2 opacity-0 group-hover:opacity-100 transition-opacity">
                  <NuxtLink :to="`/players/${player.id}`">
                    <button class="p-2 text-neutral-400 hover:text-primary-600 hover:bg-primary-50 rounded-lg transition-colors" title="View Profile">
                      <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                      </svg>
                    </button>
                  </NuxtLink>
                  <NuxtLink :to="`/admin/players/${player.id}`">
                    <button class="p-2 text-neutral-400 hover:text-blue-600 hover:bg-blue-50 rounded-lg transition-colors" title="Edit">
                      <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                      </svg>
                    </button>
                  </NuxtLink>
                  <button 
                    @click="confirmDelete(player)"
                    class="p-2 text-neutral-400 hover:text-red-600 hover:bg-red-50 rounded-lg transition-colors"
                    title="Delete"
                  >
                    <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div v-if="players.length > 0" class="flex flex-col sm:flex-row items-center justify-between gap-4 p-6 border-t border-neutral-100">
        <p class="text-sm text-neutral-500">
          Showing <span class="font-medium text-neutral-900">{{ (page - 1) * perPage + 1 }}</span> to 
          <span class="font-medium text-neutral-900">{{ Math.min(page * perPage, totalPlayers) }}</span> of 
          <span class="font-medium text-neutral-900">{{ totalPlayers }}</span> players
        </p>
        <div class="flex items-center gap-2">
          <button
            :disabled="page === 1"
            @click="page--; fetchPlayers()"
            :class="[
              'flex items-center gap-1 px-4 py-2 text-sm font-medium rounded-lg transition-colors',
              page === 1 
                ? 'text-neutral-300 cursor-not-allowed' 
                : 'text-neutral-600 hover:bg-neutral-100'
            ]"
          >
            <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
            </svg>
            Previous
          </button>
          
          <!-- Page Numbers -->
          <div class="hidden sm:flex items-center gap-1">
            <button
              v-for="pageNum in visiblePages"
              :key="pageNum"
              @click="page = pageNum; fetchPlayers()"
              :class="[
                'w-9 h-9 text-sm font-medium rounded-lg transition-colors',
                page === pageNum 
                  ? 'bg-primary-500 text-white' 
                  : 'text-neutral-600 hover:bg-neutral-100'
              ]"
            >
              {{ pageNum }}
            </button>
          </div>

          <button
            :disabled="page >= totalPages"
            @click="page++; fetchPlayers()"
            :class="[
              'flex items-center gap-1 px-4 py-2 text-sm font-medium rounded-lg transition-colors',
              page >= totalPages 
                ? 'text-neutral-300 cursor-not-allowed' 
                : 'text-neutral-600 hover:bg-neutral-100'
            ]"
          >
            Next
            <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- Delete Confirmation Modal -->
    <Teleport to="body">
      <Transition
        enter-active-class="transition duration-200 ease-out"
        enter-from-class="opacity-0"
        enter-to-class="opacity-100"
        leave-active-class="transition duration-150 ease-in"
        leave-from-class="opacity-100"
        leave-to-class="opacity-0"
      >
        <div v-if="showDeleteModal" class="fixed inset-0 z-50 overflow-y-auto">
          <div class="flex min-h-full items-center justify-center p-4">
            <div class="fixed inset-0 bg-black/50 backdrop-blur-sm" @click="showDeleteModal = false"></div>
            
            <Transition
              enter-active-class="transition duration-200 ease-out"
              enter-from-class="transform scale-95 opacity-0"
              enter-to-class="transform scale-100 opacity-100"
              leave-active-class="transition duration-150 ease-in"
              leave-from-class="transform scale-100 opacity-100"
              leave-to-class="transform scale-95 opacity-0"
            >
              <div v-if="showDeleteModal" class="relative bg-white rounded-2xl shadow-xl max-w-md w-full p-6">
                <div class="flex items-center gap-4 mb-6">
                  <div class="w-12 h-12 bg-red-100 rounded-xl flex items-center justify-center flex-shrink-0">
                    <svg class="w-6 h-6 text-red-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                    </svg>
                  </div>
                  <div>
                    <h3 class="text-lg font-semibold text-neutral-900">Delete Player</h3>
                    <p class="text-sm text-neutral-500">This action cannot be undone</p>
                  </div>
                </div>
                
                <p class="text-neutral-600 mb-6">
                  Are you sure you want to delete <span class="font-semibold">{{ playerToDelete?.first_name }} {{ playerToDelete?.last_name }}</span>? 
                  All associated data including videos and contact requests will be permanently removed.
                </p>

                <div class="flex gap-3">
                  <button 
                    @click="showDeleteModal = false"
                    class="flex-1 px-4 py-2.5 text-sm font-medium text-neutral-700 bg-neutral-100 rounded-xl hover:bg-neutral-200 transition-colors"
                  >
                    Cancel
                  </button>
                  <button 
                    @click="deletePlayer"
                    :disabled="deleting"
                    class="flex-1 flex items-center justify-center gap-2 px-4 py-2.5 text-sm font-medium text-white bg-red-600 rounded-xl hover:bg-red-700 disabled:opacity-50 transition-colors"
                  >
                    <svg v-if="deleting" class="w-4 h-4 animate-spin" fill="none" viewBox="0 0 24 24">
                      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                    {{ deleting ? 'Deleting...' : 'Delete Player' }}
                  </button>
                </div>
              </div>
            </Transition>
          </div>
        </div>
      </Transition>
    </Teleport>
  </div>
</template>

<script setup lang="ts">
import type { Player, ApiResponse } from '~/types'

interface Academy {
  id: string
  name: string
}

definePageMeta({
  layout: 'admin',
  middleware: 'admin',
})

const api = useApi()

const players = ref<Player[]>([])
const academies = ref<Academy[]>([])
const loading = ref(true)
const page = ref(1)
const perPage = 20
const totalPlayers = ref(0)
const totalPages = computed(() => Math.ceil(totalPlayers.value / perPage))

// Computed stats
const verifiedCount = computed(() => players.value.filter(p => p.is_verified).length)
const pendingCount = computed(() => players.value.filter(p => !p.is_verified).length)

const filters = reactive({
  search: '',
  status: '',
  academy: '',
})

const hasActiveFilters = computed(() => {
  return filters.search || filters.status || filters.academy
})

const visiblePages = computed(() => {
  const pages = []
  const total = totalPages.value
  const current = page.value
  
  let start = Math.max(1, current - 2)
  let end = Math.min(total, start + 4)
  
  if (end - start < 4) {
    start = Math.max(1, end - 4)
  }
  
  for (let i = start; i <= end; i++) {
    pages.push(i)
  }
  
  return pages
})

const showDeleteModal = ref(false)
const playerToDelete = ref<Player | null>(null)
const deleting = ref(false)

let searchTimeout: ReturnType<typeof setTimeout>

function debouncedSearch() {
  clearTimeout(searchTimeout)
  searchTimeout = setTimeout(() => {
    page.value = 1
    fetchPlayers()
  }, 300)
}

function clearFilters() {
  filters.search = ''
  filters.status = ''
  filters.academy = ''
  page.value = 1
  fetchPlayers()
}

async function fetchPlayers() {
  loading.value = true
  try {
    const params = new URLSearchParams({
      page: page.value.toString(),
      limit: perPage.toString(),
    })
    
    if (filters.search) params.append('search', filters.search)
    if (filters.status) params.append('verified', filters.status === 'verified' ? 'true' : 'false')
    if (filters.academy) params.append('academy_id', filters.academy)

    const response = await api.get<ApiResponse<{ players: Player[]; total: number }>>(
      `/admin/players?${params}`,
      {},
      true
    )
    
    if (response.success && response.data) {
      players.value = response.data.players || []
      totalPlayers.value = response.data.total || 0
    }
  } catch (error) {
    console.error('Failed to fetch players:', error)
  } finally {
    loading.value = false
  }
}

async function fetchAcademies() {
  try {
    const response = await api.get<ApiResponse<{ academies: Academy[] }>>('/admin/academies', {}, true)
    if (response.success && response.data) {
      academies.value = response.data.academies || []
    }
  } catch (error) {
    console.error('Failed to fetch academies:', error)
  }
}

function confirmDelete(player: Player) {
  playerToDelete.value = player
  showDeleteModal.value = true
}

async function deletePlayer() {
  if (!playerToDelete.value) return
  
  deleting.value = true
  try {
    await api.delete(`/admin/players/${playerToDelete.value.id}`, true)
    players.value = players.value.filter(p => p.id !== playerToDelete.value?.id)
    totalPlayers.value--
    showDeleteModal.value = false
    playerToDelete.value = null
  } catch (error) {
    console.error('Failed to delete player:', error)
  } finally {
    deleting.value = false
  }
}

onMounted(() => {
  fetchPlayers()
  fetchAcademies()
})
</script>
