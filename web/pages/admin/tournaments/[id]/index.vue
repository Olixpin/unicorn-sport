<template>
  <div class="max-w-7xl mx-auto">
    <!-- Page Header -->
    <div class="mb-8">
      <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
        <div class="flex items-center gap-4">
          <NuxtLink 
            to="/admin/tournaments" 
            class="w-10 h-10 flex items-center justify-center rounded-xl bg-neutral-100 hover:bg-neutral-200 text-neutral-600 transition-colors"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
            </svg>
          </NuxtLink>
          <div>
            <h1 class="font-display text-2xl lg:text-3xl font-bold text-neutral-900">
              Tournament Details
            </h1>
            <p class="mt-1 text-neutral-600">Manage matches and upload videos</p>
          </div>
        </div>
        <div class="flex items-center gap-3">
          <button
            @click="showEditModal = true"
            class="inline-flex items-center gap-2 px-5 py-2.5 bg-neutral-100 text-neutral-700 rounded-xl text-sm font-semibold hover:bg-neutral-200 transition-colors"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
            </svg>
            Edit Tournament
          </button>
        </div>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex flex-col items-center justify-center py-20">
      <div class="w-12 h-12 border-4 border-primary-200 border-t-primary-600 rounded-full animate-spin"></div>
      <p class="mt-4 text-neutral-600">Loading tournament...</p>
    </div>

    <!-- Content -->
    <div v-else-if="tournament" class="space-y-6">
      <!-- Tournament Header -->
      <!-- Tournament Header Card -->
      <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
        <div class="px-6 py-4 border-b border-neutral-200 bg-gradient-to-r from-primary-50 to-emerald-50">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 bg-gradient-to-br from-primary-500 to-emerald-600 rounded-xl flex items-center justify-center text-white shadow-lg shadow-primary-500/25">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z" />
              </svg>
            </div>
            <div>
              <h2 class="font-semibold text-neutral-900">Tournament Overview</h2>
              <p class="text-sm text-neutral-500">Details and stats</p>
            </div>
          </div>
        </div>
        
        <div class="p-6">
          <div class="flex flex-col md:flex-row md:items-start gap-6">
            <!-- Thumbnail -->
            <div class="w-full md:w-64 h-40 bg-gradient-to-br from-primary-500 to-emerald-600 rounded-xl flex items-center justify-center overflow-hidden">
              <img
                v-if="tournament.thumbnail_url"
                :src="tournament.thumbnail_url"
                :alt="tournament.name"
                class="w-full h-full object-cover"
              />
              <svg v-else class="w-16 h-16 text-white/50" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z" />
              </svg>
            </div>

            <!-- Info -->
            <div class="flex-1">
              <div class="flex items-center gap-3 mb-2">
                <h3 class="text-xl font-bold text-neutral-900">{{ tournament.name }}</h3>
                <span :class="getStatusClass(tournament.status)" class="px-3 py-1 rounded-full text-sm font-medium capitalize">
                  {{ tournament.status }}
                </span>
              </div>
              <p v-if="tournament.description" class="text-neutral-600 mb-4">{{ tournament.description }}</p>
              <div class="flex flex-wrap items-center gap-6 text-sm text-neutral-600">
                <span class="flex items-center gap-2">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                  </svg>
                  {{ formatDateRange(tournament.start_date, tournament.end_date) }}
                </span>
                <span v-if="tournament.location" class="flex items-center gap-2">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                  </svg>
                  {{ tournament.location }}
                </span>
                <span class="flex items-center gap-2 font-medium text-primary-600">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
                  </svg>
                  {{ matches.length }} matches
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Matches Section -->
      <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
        <div class="px-6 py-4 border-b border-neutral-200 bg-gradient-to-r from-blue-50 to-indigo-50">
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-3">
              <div class="w-10 h-10 bg-gradient-to-br from-blue-500 to-indigo-600 rounded-xl flex items-center justify-center text-white shadow-lg shadow-blue-500/25">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
                </svg>
              </div>
              <div>
                <h2 class="font-semibold text-neutral-900">Matches</h2>
                <p class="text-sm text-neutral-500">Manage matches and upload videos</p>
              </div>
            </div>
            <button
              @click="showCreateMatchModal = true"
              class="inline-flex items-center gap-2 px-4 py-2 bg-gradient-to-r from-primary-600 to-emerald-600 text-white rounded-xl text-sm font-semibold hover:from-primary-700 hover:to-emerald-700 transition-all shadow-lg shadow-primary-600/25"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
              </svg>
              Add Match
            </button>
          </div>
        </div>

        <div class="p-6">

        <!-- Empty State -->
        <div v-if="matches.length === 0" class="text-center py-12">
          <div class="w-16 h-16 bg-neutral-100 rounded-2xl flex items-center justify-center mx-auto mb-4">
            <svg class="w-8 h-8 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
            </svg>
          </div>
          <p class="text-neutral-500 mb-4">No matches added to this tournament yet</p>
          <button
            @click="showCreateMatchModal = true"
            class="text-primary-600 hover:text-primary-700 font-medium"
          >
            Add the first match
          </button>
        </div>

        <!-- Matches List -->
        <div v-else class="space-y-3">
          <NuxtLink
            v-for="match in matches"
            :key="match.id"
            :to="`/admin/tournaments/${tournamentId}/matches/${match.id}`"
            class="flex items-center justify-between p-4 bg-neutral-50 rounded-xl hover:bg-neutral-100 transition-colors group"
          >
            <div class="flex items-center gap-4">
              <!-- Match Number/Stage -->
              <div class="w-12 h-12 bg-gradient-to-br from-primary-500 to-emerald-600 rounded-xl flex items-center justify-center text-white font-bold">
                {{ match.match_number || '#' }}
              </div>

              <div>
                <p class="font-medium text-neutral-900 group-hover:text-primary-600 transition-colors">
                  {{ match.title }}
                </p>
                <div class="flex items-center gap-3 text-sm text-neutral-500 mt-0.5">
                  <span v-if="match.home_team && match.away_team">
                    {{ match.home_team }}
                    <span v-if="match.home_score !== null" class="font-medium text-neutral-700">{{ match.home_score }}</span>
                    vs
                    <span v-if="match.away_score !== null" class="font-medium text-neutral-700">{{ match.away_score }}</span>
                    {{ match.away_team }}
                  </span>
                  <span class="flex items-center gap-1">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                    </svg>
                    {{ formatDate(match.match_date) }}
                  </span>
                  <span v-if="match.stage" class="px-2 py-0.5 bg-neutral-200 text-neutral-600 rounded text-xs">
                    {{ formatStage(match.stage) }}
                  </span>
                </div>
              </div>
            </div>

            <div class="flex items-center gap-6">
              <!-- Video Status -->
              <div class="flex items-center gap-3 text-sm">
                <span
                  :class="match.has_video ? 'text-emerald-600' : 'text-neutral-400'"
                  class="flex items-center gap-1"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
                  </svg>
                  {{ match.has_video ? 'Video' : 'No video' }}
                </span>
                <span class="text-neutral-600">
                  {{ match.player_count || 0 }} players
                </span>
                <span class="text-primary-600 font-medium">
                  {{ match.highlight_count || 0 }} highlights
                </span>
              </div>

              <!-- Status Badge -->
              <span :class="getMatchStatusClass(match.status)" class="px-2.5 py-1 rounded-full text-xs font-medium capitalize">
                {{ match.status }}
              </span>

              <!-- Arrow -->
              <svg class="w-5 h-5 text-neutral-400 group-hover:text-primary-600 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
              </svg>
            </div>
          </NuxtLink>
        </div>
        </div>
      </div>
    </div>

    <!-- Create Match Modal -->
    <Teleport to="body">
      <div v-if="showCreateMatchModal" class="fixed inset-0 z-50 overflow-y-auto">
        <div class="fixed inset-0 bg-black/50" @click="showCreateMatchModal = false"></div>
        <div class="relative min-h-screen flex items-center justify-center p-4">
          <div class="relative bg-white rounded-2xl shadow-xl w-full max-w-lg p-6">
            <h3 class="text-xl font-semibold mb-6">Add Match</h3>

            <form @submit.prevent="createMatch" class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-1">Match Title *</label>
                <input
                  v-model="matchForm.title"
                  type="text"
                  class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                  placeholder="e.g., Quarter Final - Game 1"
                  required
                />
              </div>

              <div class="grid grid-cols-2 gap-4">
                <div>
                  <label class="block text-sm font-medium text-neutral-700 mb-1">Home Team</label>
                  <input
                    v-model="matchForm.home_team"
                    type="text"
                    class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                    placeholder="Team A"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-neutral-700 mb-1">Away Team</label>
                  <input
                    v-model="matchForm.away_team"
                    type="text"
                    class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                    placeholder="Team B"
                  />
                </div>
              </div>

              <div class="grid grid-cols-2 gap-4">
                <div>
                  <label class="block text-sm font-medium text-neutral-700 mb-1">Match Date *</label>
                  <input
                    v-model="matchForm.match_date"
                    type="datetime-local"
                    :min="matchDateMin ? matchDateMin + 'T00:00' : undefined"
                    :max="matchDateMax ? matchDateMax + 'T23:59' : undefined"
                    :class="['w-full px-4 py-3 border rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent', !isMatchDateValid ? 'border-rose-300 bg-rose-50' : 'border-neutral-200']"
                    required
                  />
                  <p v-if="matchDateError" class="mt-1 text-xs text-rose-600">{{ matchDateError }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-neutral-700 mb-1">Stage</label>
                  <select
                    v-model="matchForm.stage"
                    class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent bg-white"
                  >
                    <option value="">Select stage</option>
                    <option value="group">Group Stage</option>
                    <option value="round_of_16">Round of 16</option>
                    <option value="quarterfinal">Quarter Final</option>
                    <option value="semifinal">Semi Final</option>
                    <option value="final">Final</option>
                  </select>
                </div>
              </div>

              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-1">Location</label>
                <input
                  v-model="matchForm.location"
                  type="text"
                  class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                  placeholder="Stadium name"
                />
              </div>

              <div class="flex justify-end gap-3 pt-4">
                <button
                  type="button"
                  @click="showCreateMatchModal = false"
                  class="px-4 py-2.5 border border-neutral-200 rounded-xl hover:bg-neutral-50 transition-colors"
                >
                  Cancel
                </button>
                <button
                  type="submit"
                  :disabled="creatingMatch || !isMatchDateValid"
                  class="px-6 py-2.5 bg-gradient-to-r from-primary-500 to-emerald-600 text-white rounded-xl font-medium hover:from-primary-600 hover:to-emerald-700 transition-all disabled:opacity-50 disabled:cursor-not-allowed"
                >
                  {{ creatingMatch ? 'Creating...' : 'Create Match' }}
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
}

interface Match {
  id: string
  title: string
  match_date: string
  location?: string
  stage?: string
  home_team?: string
  away_team?: string
  home_score?: number
  away_score?: number
  status: string
  match_number?: number
  has_video: boolean
  player_count?: number
  highlight_count?: number
}

definePageMeta({
  layout: 'admin',
  middleware: 'admin',
})

const route = useRoute()
const config = useRuntimeConfig()
const authStore = useAuthStore()
const toast = useToast()
const router = useRouter()

const tournamentId = route.params.id as string

const loading = ref(true)
const tournament = ref<Tournament | null>(null)
const matches = ref<Match[]>([])
const showEditModal = ref(false)
const showCreateMatchModal = ref(false)
const creatingMatch = ref(false)

const matchForm = reactive({
  title: '',
  home_team: '',
  away_team: '',
  match_date: '',
  stage: '',
  location: '',
})

// Match date validation - must be within tournament dates
const matchDateMin = computed(() => {
  if (!tournament.value?.start_date) return ''
  return tournament.value.start_date.split('T')[0]
})

const matchDateMax = computed(() => {
  if (!tournament.value?.end_date) return ''
  return tournament.value.end_date.split('T')[0]
})

const isMatchDateValid = computed(() => {
  if (!matchForm.match_date || !tournament.value) return true
  const matchDate = new Date(matchForm.match_date)
  const startDate = new Date(tournament.value.start_date)
  const endDate = new Date(tournament.value.end_date)
  // Set to start/end of day for comparison
  startDate.setHours(0, 0, 0, 0)
  endDate.setHours(23, 59, 59, 999)
  return matchDate >= startDate && matchDate <= endDate
})

const matchDateError = computed(() => {
  if (!matchForm.match_date || isMatchDateValid.value) return ''
  return `Match date must be between ${formatDate(tournament.value!.start_date)} and ${formatDate(tournament.value!.end_date)}`
})

async function fetchTournament() {
  try {
    // Backend returns "events" not "tournaments"
    const response = await $fetch<ApiResponse<{ events: Tournament[] }>>('/admin/events', {
      baseURL: config.public.apiBase,
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
    })
    if (response.success && response.data) {
      tournament.value = response.data.events?.find(t => t.id === tournamentId) || null
    }
  } catch (error) {
    console.error('Failed to fetch tournament:', error)
  }
}

async function fetchMatches() {
  try {
    const response = await $fetch<ApiResponse<{ matches: Match[] }>>(`/admin/tournaments/${tournamentId}/matches`, {
      baseURL: config.public.apiBase,
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
    })
    if (response.success && response.data) {
      matches.value = response.data.matches || []
    }
  } catch (error) {
    console.error('Failed to fetch matches:', error)
  } finally {
    loading.value = false
  }
}

async function createMatch() {
  creatingMatch.value = true
  try {
    const response = await $fetch<ApiResponse<Match>>(`/admin/tournaments/${tournamentId}/matches`, {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: {
        ...matchForm,
        match_date: matchForm.match_date ? new Date(matchForm.match_date).toISOString() : undefined,
      },
    })
    if (response.success && response.data) {
      toast.success('Match Created', 'Match has been added to the tournament')
      showCreateMatchModal.value = false
      Object.assign(matchForm, { title: '', home_team: '', away_team: '', match_date: '', stage: '', location: '' })
      router.push(`/admin/tournaments/${tournamentId}/matches/${response.data.id}`)
    }
  } catch (error: any) {
    toast.error('Error', error.data?.message || 'Failed to create match')
  } finally {
    creatingMatch.value = false
  }
}

function formatDateRange(start: string, end: string): string {
  const startDate = new Date(start)
  const endDate = new Date(end)
  const opts: Intl.DateTimeFormatOptions = { month: 'short', day: 'numeric' }
  return `${startDate.toLocaleDateString('en-US', opts)} - ${endDate.toLocaleDateString('en-US', opts)}, ${endDate.getFullYear()}`
}

function formatDate(date: string): string {
  return new Date(date).toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' })
}

function formatStage(stage: string): string {
  const stages: Record<string, string> = {
    group: 'Group',
    round_of_16: 'R16',
    quarterfinal: 'QF',
    semifinal: 'SF',
    final: 'Final',
  }
  return stages[stage] || stage
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

function getMatchStatusClass(status: string): string {
  const classes: Record<string, string> = {
    scheduled: 'bg-blue-100 text-blue-700',
    in_progress: 'bg-amber-100 text-amber-700',
    completed: 'bg-emerald-100 text-emerald-700',
    cancelled: 'bg-neutral-100 text-neutral-700',
  }
  return classes[status] || classes.scheduled!
}

onMounted(async () => {
  await Promise.all([fetchTournament(), fetchMatches()])
})
</script>
