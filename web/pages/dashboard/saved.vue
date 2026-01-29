<template>
  <div>
    <!-- Header -->
    <div class="mb-8">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="font-display text-2xl lg:text-3xl font-bold text-neutral-900">
            Saved Players
          </h1>
          <p class="mt-1 text-neutral-500">Players you've saved for later review</p>
        </div>
        <div v-if="savedPlayers.length > 0" class="hidden sm:flex items-center gap-3">
          <span class="px-3 py-1.5 bg-neutral-100 text-neutral-600 text-sm font-medium rounded-lg">
            {{ savedPlayers.length }} {{ savedPlayers.length === 1 ? 'player' : 'players' }}
          </span>
        </div>
      </div>
      
      <!-- Filters -->
      <div v-if="subscriptionStore.canSavePlayers && !loading" class="mt-4 flex flex-wrap gap-3">
        <select v-model="filterPriority" @change="applyFilters" class="px-4 py-2 bg-neutral-100 border-0 rounded-xl text-sm focus:ring-2 focus:ring-primary-500">
          <option value="">All Priorities</option>
          <option value="high">🔴 High</option>
          <option value="medium">🟡 Medium</option>
          <option value="low">🔵 Low</option>
        </select>
        <select v-model="filterTag" @change="applyFilters" class="px-4 py-2 bg-neutral-100 border-0 rounded-xl text-sm focus:ring-2 focus:ring-primary-500">
          <option value="">All Tags</option>
          <option v-for="tag in allTags" :key="tag.id" :value="tag.id">{{ tag.name }}</option>
        </select>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex justify-center py-16">
      <div class="w-10 h-10 border-4 border-primary-500 border-t-transparent rounded-full animate-spin"></div>
    </div>

    <!-- Subscription Gate -->
    <div v-if="!subscriptionStore.canSavePlayers && !loading" class="bg-white rounded-2xl border border-neutral-200 p-12 text-center">
      <div class="w-20 h-20 mx-auto bg-gradient-to-br from-amber-100 to-orange-100 rounded-2xl flex items-center justify-center mb-6">
        <svg class="w-10 h-10 text-amber-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
        </svg>
      </div>
      <h3 class="text-xl font-semibold text-neutral-900 mb-2">Upgrade to Save Players</h3>
      <p class="text-neutral-500 mb-6 max-w-sm mx-auto">Scout tier or higher is required to save players to your shortlist.</p>
      <NuxtLink to="/pricing">
        <button class="px-6 py-3 bg-primary-500 text-white font-semibold rounded-xl hover:bg-primary-600 transition-colors shadow-lg shadow-primary-500/25">
          View Plans
        </button>
      </NuxtLink>
    </div>

    <!-- Empty State -->
    <div v-else-if="savedPlayers.length === 0 && !loading" class="bg-white rounded-2xl border border-neutral-200 p-12 text-center">
      <div class="w-20 h-20 mx-auto bg-gradient-to-br from-pink-100 to-rose-100 rounded-2xl flex items-center justify-center mb-6">
        <svg class="w-10 h-10 text-rose-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
        </svg>
      </div>
      <h3 class="text-xl font-semibold text-neutral-900 mb-2">No saved players yet</h3>
      <p class="text-neutral-500 mb-6 max-w-sm mx-auto">Start browsing to find talented players and save them to your shortlist.</p>
      <NuxtLink to="/discover">
        <button class="px-6 py-3 bg-primary-500 text-white font-semibold rounded-xl hover:bg-primary-600 transition-colors shadow-lg shadow-primary-500/25">
          Browse Players
        </button>
      </NuxtLink>
    </div>

    <!-- Saved Players Grid -->
    <div v-else-if="savedPlayers.length > 0">
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
        <div 
          v-for="savedPlayer in savedPlayers" 
          :key="savedPlayer.id" 
          class="group relative bg-white rounded-2xl border border-neutral-200 overflow-hidden shadow-sm hover:shadow-xl hover:border-neutral-300 transition-all duration-300"
        >
          <!-- Remove Button -->
          <button
            @click.stop="unsavePlayer(savedPlayer.player_id)"
            class="absolute top-3 right-3 z-10 p-2.5 bg-white/90 backdrop-blur-sm rounded-xl shadow-lg opacity-0 group-hover:opacity-100 hover:bg-red-50 transition-all duration-200"
            title="Remove from saved"
          >
            <svg class="w-5 h-5 text-red-500" fill="currentColor" viewBox="0 0 24 24">
              <path d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
            </svg>
          </button>
          
          <!-- Edit Button -->
          <button
            @click.stop="openEditModal(savedPlayer)"
            class="absolute top-3 right-14 z-10 p-2.5 bg-white/90 backdrop-blur-sm rounded-xl shadow-lg opacity-0 group-hover:opacity-100 hover:bg-primary-50 transition-all duration-200"
            title="Edit notes & tags"
          >
            <svg class="w-5 h-5 text-primary-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
            </svg>
          </button>
          
          <!-- Priority Badge -->
          <span 
            v-if="savedPlayer.priority" 
            :class="['absolute top-3 left-3 z-10 px-2 py-1 text-xs font-semibold rounded-lg capitalize', getPriorityColor(savedPlayer.priority)]"
          >
            {{ savedPlayer.priority }}
          </span>

          <NuxtLink :to="`/players/${savedPlayer.player_id}`">
            <!-- Photo -->
            <div class="aspect-[4/5] bg-neutral-200 overflow-hidden">
              <img 
                v-if="savedPlayer.player?.profile_photo_url || savedPlayer.player?.thumbnail_url"
                :src="savedPlayer.player?.profile_photo_url || savedPlayer.player?.thumbnail_url"
                :alt="`${savedPlayer.player?.first_name}`"
                class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
              />
              <div v-else class="w-full h-full flex items-center justify-center bg-gradient-to-br from-neutral-200 to-neutral-300">
                <svg class="h-16 w-16 text-neutral-400" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z"/>
                </svg>
              </div>
            </div>

            <!-- Info -->
            <div class="p-5">
              <div class="flex items-start justify-between mb-2">
                <h3 class="font-semibold text-lg text-neutral-900 group-hover:text-primary-600 transition-colors">
                  {{ savedPlayer.player?.first_name }} {{ savedPlayer.player?.last_name_init }}.
                </h3>
                <span class="text-xs text-neutral-400">
                  Saved {{ formatDate(savedPlayer.saved_at) }}
                </span>
              </div>
              <div class="flex items-center gap-2 text-sm text-neutral-500">
                <span>{{ savedPlayer.player?.position || 'Unknown' }}</span>
                <span class="w-1 h-1 rounded-full bg-neutral-300"></span>
                <span>{{ savedPlayer.player?.country || 'Unknown' }}</span>
                <span v-if="savedPlayer.player?.age" class="w-1 h-1 rounded-full bg-neutral-300"></span>
                <span v-if="savedPlayer.player?.age">{{ savedPlayer.player.age }} yrs</span>
              </div>
              <!-- Notes -->
              <div v-if="savedPlayer.notes" class="mt-3 p-3 bg-neutral-50 rounded-lg border border-neutral-100">
                <div class="flex items-start gap-2">
                  <svg class="w-4 h-4 text-neutral-400 mt-0.5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                  </svg>
                  <p class="text-sm text-neutral-600 line-clamp-2">{{ savedPlayer.notes }}</p>
                </div>
              </div>
              
              <!-- Tags -->
              <div v-if="savedPlayer.tags && savedPlayer.tags.length > 0" class="mt-3 flex flex-wrap gap-1.5">
                <span 
                  v-for="tag in savedPlayer.tags" 
                  :key="tag.id" 
                  class="px-2 py-0.5 text-xs font-medium rounded-full"
                  :style="{ backgroundColor: tag.color + '20', color: tag.color }"
                >
                  {{ tag.name }}
                </span>
              </div>
              
              <div v-if="!savedPlayer.notes && (!savedPlayer.tags || savedPlayer.tags.length === 0)" class="mt-3 text-xs text-neutral-400 italic">No notes or tags</div>
            </div>
          </NuxtLink>
        </div>
      </div>

      <!-- Load More -->
      <div v-if="hasMore" class="mt-10 text-center">
        <button 
          @click="loadMore" 
          :disabled="loadingMore"
          class="px-8 py-3 border-2 border-neutral-200 text-neutral-700 font-semibold rounded-xl hover:bg-neutral-50 hover:border-neutral-300 transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed flex items-center gap-2 mx-auto"
        >
          <svg v-if="loadingMore" class="w-5 h-5 animate-spin" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
          </svg>
          {{ loadingMore ? 'Loading...' : 'Load More' }}
        </button>
      </div>
    </div>
    
    <!-- Edit Modal -->
    <Teleport to="body">
      <div v-if="showEditModal" class="fixed inset-0 z-50 flex items-center justify-center p-4">
        <div class="absolute inset-0 bg-black/50" @click="closeEditModal"></div>
        <div class="relative bg-white rounded-2xl shadow-2xl w-full max-w-md p-6 max-h-[90vh] overflow-y-auto">
          <div class="flex items-center justify-between mb-6">
            <h3 class="text-xl font-bold text-neutral-900">Edit Saved Player</h3>
            <button @click="closeEditModal" class="p-2 hover:bg-neutral-100 rounded-lg transition-colors">
              <svg class="w-5 h-5 text-neutral-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
          
          <div v-if="editingPlayer" class="space-y-5">
            <!-- Player Preview -->
            <div class="flex items-center gap-3 p-3 bg-neutral-50 rounded-xl">
              <img 
                v-if="editingPlayer.player?.profile_photo_url" 
                :src="editingPlayer.player.profile_photo_url" 
                class="w-12 h-12 rounded-lg object-cover"
              />
              <div v-else class="w-12 h-12 rounded-lg bg-neutral-200 flex items-center justify-center">
                <svg class="w-6 h-6 text-neutral-400" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z"/>
                </svg>
              </div>
              <div>
                <p class="font-semibold text-neutral-900">{{ editingPlayer.player?.first_name }} {{ editingPlayer.player?.last_name_init }}.</p>
                <p class="text-sm text-neutral-500">{{ editingPlayer.player?.position }}</p>
              </div>
            </div>
            
            <!-- Priority -->
            <div>
              <label class="block text-sm font-medium text-neutral-700 mb-2">Priority</label>
              <div class="flex gap-2">
                <button 
                  @click="editPriority = ''" 
                  :class="['px-4 py-2 rounded-xl text-sm font-medium transition-colors', editPriority === '' ? 'bg-neutral-900 text-white' : 'bg-neutral-100 text-neutral-600 hover:bg-neutral-200']"
                >
                  None
                </button>
                <button 
                  @click="editPriority = 'high'" 
                  :class="['px-4 py-2 rounded-xl text-sm font-medium transition-colors', editPriority === 'high' ? 'bg-red-500 text-white' : 'bg-red-50 text-red-600 hover:bg-red-100']"
                >
                  🔴 High
                </button>
                <button 
                  @click="editPriority = 'medium'" 
                  :class="['px-4 py-2 rounded-xl text-sm font-medium transition-colors', editPriority === 'medium' ? 'bg-amber-500 text-white' : 'bg-amber-50 text-amber-600 hover:bg-amber-100']"
                >
                  🟡 Medium
                </button>
                <button 
                  @click="editPriority = 'low'" 
                  :class="['px-4 py-2 rounded-xl text-sm font-medium transition-colors', editPriority === 'low' ? 'bg-blue-500 text-white' : 'bg-blue-50 text-blue-600 hover:bg-blue-100']"
                >
                  🔵 Low
                </button>
              </div>
            </div>
            
            <!-- Tags -->
            <div>
              <label class="block text-sm font-medium text-neutral-700 mb-2">Tags</label>
              <div v-if="allTags.length > 0" class="flex flex-wrap gap-2">
                <button 
                  v-for="tag in allTags" 
                  :key="tag.id"
                  @click="editTags.includes(tag.id) ? editTags = editTags.filter(t => t !== tag.id) : editTags.push(tag.id)"
                  :class="['px-3 py-1.5 rounded-full text-sm font-medium transition-colors border-2', editTags.includes(tag.id) ? 'border-current' : 'border-transparent opacity-60 hover:opacity-100']"
                  :style="{ backgroundColor: tag.color + '20', color: tag.color }"
                >
                  {{ tag.name }}
                </button>
              </div>
              <p v-else class="text-sm text-neutral-400">No tags created yet. Create tags to organize your players.</p>
            </div>
            
            <!-- Notes -->
            <div>
              <label class="block text-sm font-medium text-neutral-700 mb-2">Notes</label>
              <textarea 
                v-model="editNotes"
                rows="4"
                class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-primary-500 resize-none"
                placeholder="Add notes about this player..."
              ></textarea>
            </div>
            
            <!-- Actions -->
            <div class="flex gap-3 pt-2">
              <button 
                @click="closeEditModal" 
                class="flex-1 px-4 py-3 border border-neutral-200 text-neutral-700 font-semibold rounded-xl hover:bg-neutral-50 transition-colors"
              >
                Cancel
              </button>
              <button 
                @click="saveChanges" 
                :disabled="saving"
                class="flex-1 px-4 py-3 bg-primary-500 text-white font-semibold rounded-xl hover:bg-primary-600 transition-colors disabled:opacity-50"
              >
                {{ saving ? 'Saving...' : 'Save Changes' }}
              </button>
            </div>
          </div>
        </div>
      </div>
    </Teleport>
  </div>
</template>

<script setup lang="ts">
import type { Player, ApiResponse } from '~/types'

interface ScoutTag {
  id: string
  name: string
  color: string
}

interface SavedPlayer {
  id: string
  player_id: string
  notes?: string
  priority: 'high' | 'medium' | 'low' | ''
  tags?: ScoutTag[]
  saved_at: string
  player: {
    first_name: string
    last_name_init: string
    age: number
    position: string
    country: string
    thumbnail_url?: string
    profile_photo_url?: string
  }
}

definePageMeta({
  layout: 'dashboard',
  middleware: 'auth',
})

const subscriptionStore = useSubscriptionStore()
const savedPlayers = ref<SavedPlayer[]>([])
const allTags = ref<ScoutTag[]>([])
const loading = ref(true)
const loadingMore = ref(false)
const hasMore = ref(false)
const page = ref(1)

// Filter state
const filterPriority = ref('')
const filterTag = ref('')

// Edit modal state
const showEditModal = ref(false)
const editingPlayer = ref<SavedPlayer | null>(null)
const editNotes = ref('')
const editPriority = ref<'high' | 'medium' | 'low' | ''>('')
const editTags = ref<string[]>([])
const saving = ref(false)

const api = useApi()

// Check subscription on mount
onMounted(async () => {
  await subscriptionStore.fetchSubscription()
  if (subscriptionStore.canSavePlayers) {
    await Promise.all([fetchSavedPlayers(), fetchTags()])
  } else {
    loading.value = false
  }
})

async function fetchTags() {
  try {
    const response = await api.get<ApiResponse<{ tags: ScoutTag[] }>>('/tags', {}, true)
    if (response.success && response.data) {
      allTags.value = response.data.tags || []
    }
  } catch (error) {
    console.error('Failed to fetch tags:', error)
  }
}

async function fetchSavedPlayers() {
  try {
    let url = `/saved-players?page=${page.value}&limit=12`
    if (filterPriority.value) url += `&priority=${filterPriority.value}`
    if (filterTag.value) url += `&tag_id=${filterTag.value}`
    
    const response = await api.get<ApiResponse<{ saved_players: SavedPlayer[]; pagination: { total: number; total_pages: number } }>>(
      url,
      {},
      true
    )
    
    if (response.success && response.data) {
      if (page.value === 1) {
        savedPlayers.value = response.data.saved_players || []
      } else {
        savedPlayers.value.push(...(response.data.saved_players || []))
      }
      const totalPages = response.data.pagination?.total_pages || 1
      hasMore.value = page.value < totalPages
    }
  } catch (error) {
    console.error('Failed to fetch saved players:', error)
  } finally {
    loading.value = false
    loadingMore.value = false
  }
}

async function unsavePlayer(playerId: string) {
  try {
    const response = await api.delete<ApiResponse<null>>(`/players/${playerId}/save`, true)
    
    if (response.success) {
      savedPlayers.value = savedPlayers.value.filter(p => p.player_id !== playerId)
    }
  } catch (error) {
    console.error('Failed to unsave player:', error)
  }
}

async function loadMore() {
  loadingMore.value = true
  page.value++
  await fetchSavedPlayers()
}

function applyFilters() {
  page.value = 1
  loading.value = true
  fetchSavedPlayers()
}

function openEditModal(savedPlayer: SavedPlayer) {
  editingPlayer.value = savedPlayer
  editNotes.value = savedPlayer.notes || ''
  editPriority.value = savedPlayer.priority || ''
  editTags.value = savedPlayer.tags?.map(t => t.id) || []
  showEditModal.value = true
}

function closeEditModal() {
  showEditModal.value = false
  editingPlayer.value = null
}

async function saveChanges() {
  if (!editingPlayer.value) return
  saving.value = true
  
  try {
    const response = await api.patch<ApiResponse<SavedPlayer>>(
      `/saved-players/${editingPlayer.value.id}`,
      {
        notes: editNotes.value,
        priority: editPriority.value,
        tag_ids: editTags.value,
      },
      true
    )
    
    if (response.success && response.data) {
      // Update the player in the list
      const index = savedPlayers.value.findIndex(p => p.id === editingPlayer.value!.id)
      if (index !== -1) {
        savedPlayers.value[index] = response.data
      }
      closeEditModal()
    }
  } catch (error) {
    console.error('Failed to save changes:', error)
  } finally {
    saving.value = false
  }
}

function getPriorityColor(priority: string): string {
  switch (priority) {
    case 'high': return 'bg-red-100 text-red-700'
    case 'medium': return 'bg-amber-100 text-amber-700'
    case 'low': return 'bg-blue-100 text-blue-700'
    default: return 'bg-neutral-100 text-neutral-600'
  }
}

function formatDate(dateString: string): string {
  const date = new Date(dateString)
  const now = new Date()
  const diffMs = now.getTime() - date.getTime()
  const diffDays = Math.floor(diffMs / (1000 * 60 * 60 * 24))
  
  if (diffDays === 0) return 'today'
  if (diffDays === 1) return 'yesterday'
  if (diffDays < 7) return `${diffDays} days ago`
  if (diffDays < 30) return `${Math.floor(diffDays / 7)} weeks ago`
  return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric' })
}
</script>
