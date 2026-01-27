<template>
  <div class="min-h-screen bg-neutral-50">
    <!-- Header -->
    <div class="bg-white border-b border-neutral-200 sticky top-0 z-10">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-2xl font-bold text-neutral-900">Tournaments & Videos</h1>
            <p class="text-neutral-500 mt-1">Manage tournaments, matches, and player highlights</p>
          </div>
          <button
            @click="showCreateModal = true"
            class="flex items-center gap-2 px-4 py-2.5 bg-gradient-to-r from-primary-500 to-emerald-600 text-white rounded-xl font-medium hover:from-primary-600 hover:to-emerald-700 transition-all shadow-lg shadow-primary-500/25"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
            </svg>
            Create Tournament
          </button>
        </div>
      </div>
    </div>

    <!-- Content -->
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <!-- Stats Cards -->
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
        <!-- Total Tournaments -->
        <div class="bg-gradient-to-br from-primary-500 to-emerald-600 rounded-2xl p-5 text-white shadow-lg shadow-primary-500/25">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-primary-100 text-sm font-medium">Tournaments</p>
              <p class="text-3xl font-bold mt-1">{{ tournaments.length }}</p>
            </div>
            <div class="w-12 h-12 bg-white/20 rounded-xl flex items-center justify-center">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z" />
              </svg>
            </div>
          </div>
        </div>

        <!-- Total Matches -->
        <div class="bg-gradient-to-br from-blue-500 to-indigo-600 rounded-2xl p-5 text-white shadow-lg shadow-blue-500/25">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-blue-100 text-sm font-medium">Total Matches</p>
              <p class="text-3xl font-bold mt-1">{{ stats.total_matches }}</p>
            </div>
            <div class="w-12 h-12 bg-white/20 rounded-xl flex items-center justify-center">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
              </svg>
            </div>
          </div>
        </div>

        <!-- Full Match Videos (PAID) -->
        <div class="bg-gradient-to-br from-amber-500 to-orange-500 rounded-2xl p-5 text-white shadow-lg shadow-amber-500/25">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-amber-100 text-sm font-medium">Match Videos</p>
              <p class="text-3xl font-bold mt-1">{{ stats.total_match_videos }}</p>
              <p class="text-xs text-amber-200 mt-0.5">💰 Paid content</p>
            </div>
            <div class="w-12 h-12 bg-white/20 rounded-xl flex items-center justify-center">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
              </svg>
            </div>
          </div>
        </div>

        <!-- Player Highlights (FREE) -->
        <div class="bg-gradient-to-br from-emerald-500 to-green-600 rounded-2xl p-5 text-white shadow-lg shadow-emerald-500/25">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-emerald-100 text-sm font-medium">Highlights</p>
              <p class="text-3xl font-bold mt-1">{{ stats.total_highlights }}</p>
              <p class="text-xs text-emerald-200 mt-0.5">🆓 Free content</p>
            </div>
            <div class="w-12 h-12 bg-white/20 rounded-xl flex items-center justify-center">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
              </svg>
            </div>
          </div>
        </div>
      </div>

      <!-- Loading -->
      <div v-if="loading" class="flex justify-center items-center h-64">
        <div class="animate-spin rounded-full h-12 w-12 border-4 border-primary-500 border-t-transparent"></div>
      </div>

      <!-- Empty State -->
      <div v-else-if="tournaments.length === 0" class="bg-white rounded-2xl border border-neutral-200 p-12 text-center">
        <div class="w-20 h-20 bg-gradient-to-br from-primary-100 to-emerald-100 rounded-2xl flex items-center justify-center mx-auto mb-4">
          <svg class="w-10 h-10 text-primary-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
          </svg>
        </div>
        <h3 class="text-xl font-semibold text-neutral-900 mb-2">Start Building Your Video Library</h3>
        <p class="text-neutral-500 mb-6 max-w-md mx-auto">
          Create a tournament to begin uploading match videos and player highlights. 
          Scouts will discover players through free highlights and pay for full match access.
        </p>
        <button
          @click="showCreateModal = true"
          class="px-6 py-3 bg-gradient-to-r from-primary-500 to-emerald-600 text-white rounded-xl font-medium hover:from-primary-600 hover:to-emerald-700 transition-all shadow-lg shadow-primary-500/25"
        >
          Create Your First Tournament
        </button>
        
        <!-- How it works -->
        <div class="mt-10 pt-8 border-t border-neutral-100">
          <p class="text-sm font-medium text-neutral-700 mb-4">How it works:</p>
          <div class="flex flex-col sm:flex-row justify-center gap-4 text-sm text-neutral-600">
            <div class="flex items-center gap-2">
              <span class="w-6 h-6 bg-primary-100 text-primary-600 rounded-full flex items-center justify-center text-xs font-bold">1</span>
              Create Tournament
            </div>
            <svg class="w-4 h-4 text-neutral-300 hidden sm:block self-center" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
            </svg>
            <div class="flex items-center gap-2">
              <span class="w-6 h-6 bg-primary-100 text-primary-600 rounded-full flex items-center justify-center text-xs font-bold">2</span>
              Add Matches
            </div>
            <svg class="w-4 h-4 text-neutral-300 hidden sm:block self-center" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
            </svg>
            <div class="flex items-center gap-2">
              <span class="w-6 h-6 bg-primary-100 text-primary-600 rounded-full flex items-center justify-center text-xs font-bold">3</span>
              Upload Videos & Highlights
            </div>
          </div>
        </div>
      </div>

      <!-- Tournaments Grid -->
      <div v-else>
        <div class="flex items-center justify-between mb-4">
          <h2 class="text-lg font-semibold text-neutral-900">All Tournaments</h2>
        </div>
        
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <NuxtLink
            v-for="tournament in tournaments"
            :key="tournament.id"
            :to="`/admin/tournaments/${tournament.id}`"
            class="group bg-white rounded-2xl border border-neutral-200 overflow-hidden hover:shadow-lg hover:border-primary-200 transition-all"
          >
            <!-- Thumbnail -->
            <div class="aspect-video bg-gradient-to-br from-primary-500 to-emerald-600 relative">
              <img
                v-if="tournament.thumbnail_url"
                :src="tournament.thumbnail_url"
                :alt="tournament.name"
                class="w-full h-full object-cover"
              />
              <div v-else class="w-full h-full flex items-center justify-center">
                <svg class="w-16 h-16 text-white/50" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z" />
                </svg>
              </div>
              <!-- Status Badge -->
              <span
                :class="getStatusClass(tournament.status)"
                class="absolute top-3 right-3 px-2.5 py-1 rounded-full text-xs font-medium capitalize"
              >
                {{ tournament.status }}
              </span>
            </div>

            <!-- Content -->
            <div class="p-5">
              <h3 class="font-semibold text-neutral-900 group-hover:text-primary-600 transition-colors mb-2">
                {{ tournament.name }}
              </h3>
              <div class="flex items-center gap-4 text-sm text-neutral-500">
                <span class="flex items-center gap-1">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                  </svg>
                  {{ formatDateRange(tournament.start_date, tournament.end_date) }}
                </span>
              </div>
              
              <!-- Video Stats -->
              <div class="flex items-center gap-4 mt-3 pt-3 border-t border-neutral-100">
                <span class="flex items-center gap-1 text-sm">
                  <svg class="w-4 h-4 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                  </svg>
                  <span class="text-neutral-700 font-medium">{{ tournament.match_count || 0 }}</span>
                  <span class="text-neutral-500">matches</span>
                </span>
                <span class="flex items-center gap-1 text-sm">
                  <svg class="w-4 h-4 text-amber-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
                  </svg>
                  <span class="text-neutral-700 font-medium">{{ tournament.video_count || 0 }}</span>
                  <span class="text-neutral-500">videos</span>
                </span>
                <span class="flex items-center gap-1 text-sm">
                  <svg class="w-4 h-4 text-emerald-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
                  </svg>
                  <span class="text-neutral-700 font-medium">{{ tournament.highlight_count || 0 }}</span>
                  <span class="text-neutral-500">highlights</span>
                </span>
              </div>
            </div>
          </NuxtLink>
        </div>
      </div>
    </div>

    <!-- Create Tournament Modal -->
    <Teleport to="body">
      <div v-if="showCreateModal" class="fixed inset-0 z-50 overflow-y-auto">
        <div class="fixed inset-0 bg-black/50" @click="showCreateModal = false"></div>
        <div class="relative min-h-screen flex items-center justify-center p-4">
          <div class="relative bg-white rounded-2xl shadow-xl w-full max-w-lg p-6">
            <h3 class="text-xl font-semibold mb-6">Create Tournament</h3>

            <form @submit.prevent="createTournament" class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-1">Name *</label>
                <input
                  v-model="form.name"
                  type="text"
                  class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                  placeholder="e.g., Summer Cup 2026"
                  required
                />
              </div>

              <div class="grid grid-cols-2 gap-4">
                <div>
                  <label class="block text-sm font-medium text-neutral-700 mb-1">Start Date *</label>
                  <input
                    v-model="form.start_date"
                    type="date"
                    class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                    required
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-neutral-700 mb-1">End Date *</label>
                  <input
                    v-model="form.end_date"
                    type="date"
                    class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                    required
                  />
                </div>
              </div>

              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-1">Location</label>
                <input
                  v-model="form.location"
                  type="text"
                  class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                  placeholder="e.g., Lagos, Nigeria"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-1">Description</label>
                <textarea
                  v-model="form.description"
                  rows="3"
                  class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent resize-none"
                  placeholder="Brief description of the tournament..."
                ></textarea>
              </div>

              <div class="flex justify-end gap-3 pt-4">
                <button
                  type="button"
                  @click="showCreateModal = false"
                  class="px-4 py-2.5 border border-neutral-200 rounded-xl hover:bg-neutral-50 transition-colors"
                >
                  Cancel
                </button>
                <button
                  type="submit"
                  :disabled="creating"
                  class="px-6 py-2.5 bg-gradient-to-r from-primary-500 to-emerald-600 text-white rounded-xl font-medium hover:from-primary-600 hover:to-emerald-700 transition-all disabled:opacity-50"
                >
                  {{ creating ? 'Creating...' : 'Create Tournament' }}
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </Teleport>
  </div>
</template>

<script setup lang="ts">
import type { ApiResponse } from '~/types/index'

interface Tournament {
  id: string
  name: string
  description?: string
  start_date: string
  end_date: string
  location?: string
  status: string
  thumbnail_url?: string
  match_count?: number
  video_count?: number
  highlight_count?: number
}

interface VideoStats {
  total_matches: number
  total_match_videos: number
  total_highlights: number
}

definePageMeta({
  layout: 'admin',
  middleware: 'admin',
})

const config = useRuntimeConfig()
const authStore = useAuthStore()
const toast = useToast()
const router = useRouter()

const loading = ref(true)
const tournaments = ref<Tournament[]>([])
const stats = ref<VideoStats>({
  total_matches: 0,
  total_match_videos: 0,
  total_highlights: 0,
})
const showCreateModal = ref(false)
const creating = ref(false)

const form = reactive({
  name: '',
  start_date: '',
  end_date: '',
  location: '',
  description: '',
})

async function fetchTournaments() {
  try {
    const response = await $fetch<ApiResponse<{ tournaments: Tournament[] }>>('/admin/events', {
      baseURL: config.public.apiBase,
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
    })
    if (response.success && response.data) {
      tournaments.value = response.data.tournaments || []
      
      // Calculate aggregate stats from tournaments
      stats.value = {
        total_matches: tournaments.value.reduce((sum, t) => sum + (t.match_count || 0), 0),
        total_match_videos: tournaments.value.reduce((sum, t) => sum + (t.video_count || 0), 0),
        total_highlights: tournaments.value.reduce((sum, t) => sum + (t.highlight_count || 0), 0),
      }
    }
  } catch (error) {
    console.error('Failed to fetch tournaments:', error)
    toast.error('Error', 'Failed to load tournaments')
  } finally {
    loading.value = false
  }
}

async function createTournament() {
  creating.value = true
  try {
    const response = await $fetch<ApiResponse<Tournament>>('/admin/events', {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: form,
    })
    if (response.success && response.data) {
      toast.success('Tournament Created', 'Tournament has been created successfully')
      showCreateModal.value = false
      router.push(`/admin/tournaments/${response.data.id}`)
    }
  } catch (error: any) {
    toast.error('Error', error.data?.message || 'Failed to create tournament')
  } finally {
    creating.value = false
  }
}

function formatDateRange(start: string, end: string): string {
  const startDate = new Date(start)
  const endDate = new Date(end)
  const opts: Intl.DateTimeFormatOptions = { month: 'short', day: 'numeric' }
  
  if (startDate.getFullYear() === endDate.getFullYear() && startDate.getMonth() === endDate.getMonth()) {
    return `${startDate.toLocaleDateString('en-US', opts)} - ${endDate.getDate()}, ${endDate.getFullYear()}`
  }
  return `${startDate.toLocaleDateString('en-US', opts)} - ${endDate.toLocaleDateString('en-US', opts)}, ${endDate.getFullYear()}`
}

function getStatusClass(status: string): string {
  const classes: Record<string, string> = {
    draft: 'bg-neutral-100 text-neutral-700',
    upcoming: 'bg-blue-100 text-blue-700',
    active: 'bg-emerald-100 text-emerald-700',
    completed: 'bg-amber-100 text-amber-700',
  }
  return classes[status] || classes.draft!
}

onMounted(() => {
  fetchTournaments()
})
</script>
