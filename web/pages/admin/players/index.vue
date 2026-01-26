<template>
  <div>
    <div class="mb-8 flex items-center justify-between">
      <div>
        <h1 class="font-display text-2xl lg:text-3xl font-bold text-neutral-900">
          Players
        </h1>
        <p class="mt-2 text-neutral-600">Manage all players in the system.</p>
      </div>
      <NuxtLink to="/admin/players/new">
        <UButton>
          <svg class="w-5 h-5 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
          </svg>
          Add Player
        </UButton>
      </NuxtLink>
    </div>

    <!-- Filters -->
    <UCard class="mb-6">
      <div class="flex flex-wrap gap-4">
        <div class="flex-1 min-w-[200px]">
          <UInput
            v-model="filters.search"
            placeholder="Search players..."
            @input="debouncedSearch"
          >
            <template #prefix>
              <svg class="w-5 h-5 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </template>
          </UInput>
        </div>

        <select
          v-model="filters.status"
          class="px-3 py-2 border border-neutral-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-primary-500"
          @change="fetchPlayers"
        >
          <option value="">All Status</option>
          <option value="verified">Verified</option>
          <option value="pending">Pending</option>
        </select>

        <select
          v-model="filters.academy"
          class="px-3 py-2 border border-neutral-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-primary-500"
          @change="fetchPlayers"
        >
          <option value="">All Academies</option>
          <option v-for="academy in academies" :key="academy.id" :value="academy.id">
            {{ academy.name }}
          </option>
        </select>
      </div>
    </UCard>

    <!-- Players Table -->
    <UCard>
      <div v-if="loading" class="flex justify-center py-12">
        <USpinner size="lg" />
      </div>

      <div v-else-if="players.length === 0" class="text-center py-12 text-neutral-500">
        No players found
      </div>

      <div v-else class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-neutral-50">
            <tr>
              <th class="px-4 py-3 text-left text-xs font-medium text-neutral-500 uppercase tracking-wider">Player</th>
              <th class="px-4 py-3 text-left text-xs font-medium text-neutral-500 uppercase tracking-wider">Position</th>
              <th class="px-4 py-3 text-left text-xs font-medium text-neutral-500 uppercase tracking-wider">Academy</th>
              <th class="px-4 py-3 text-left text-xs font-medium text-neutral-500 uppercase tracking-wider">Status</th>
              <th class="px-4 py-3 text-left text-xs font-medium text-neutral-500 uppercase tracking-wider">Videos</th>
              <th class="px-4 py-3 text-right text-xs font-medium text-neutral-500 uppercase tracking-wider">Actions</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-neutral-200">
            <tr v-for="player in players" :key="player.id" class="hover:bg-neutral-50">
              <td class="px-4 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="w-10 h-10 rounded-full bg-neutral-200 overflow-hidden">
                    <img 
                      v-if="player.profile_photo_url"
                      :src="player.profile_photo_url"
                      :alt="player.first_name"
                      class="w-full h-full object-cover"
                    />
                  </div>
                  <div class="ml-3">
                    <p class="font-medium text-neutral-900">{{ player.first_name }} {{ player.last_name }}</p>
                    <p class="text-sm text-neutral-500">{{ player.country }}</p>
                  </div>
                </div>
              </td>
              <td class="px-4 py-4 whitespace-nowrap text-sm text-neutral-600">
                {{ player.position }}
              </td>
              <td class="px-4 py-4 whitespace-nowrap text-sm text-neutral-600">
                {{ player.academy_name || 'N/A' }}
              </td>
              <td class="px-4 py-4 whitespace-nowrap">
                <UBadge :variant="player.is_verified ? 'success' : 'warning'" size="sm">
                  {{ player.is_verified ? 'Verified' : 'Pending' }}
                </UBadge>
              </td>
              <td class="px-4 py-4 whitespace-nowrap text-sm text-neutral-600">
                {{ player.video_count || 0 }}
              </td>
              <td class="px-4 py-4 whitespace-nowrap text-right text-sm">
                <div class="flex justify-end space-x-2">
                  <NuxtLink :to="`/admin/players/${player.id}`">
                    <UButton size="sm" variant="ghost">Edit</UButton>
                  </NuxtLink>
                  <UButton 
                    size="sm" 
                    variant="ghost" 
                    class="text-red-600"
                    @click="confirmDelete(player)"
                  >
                    Delete
                  </UButton>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div v-if="totalPages > 1" class="mt-6 flex items-center justify-between">
        <p class="text-sm text-neutral-500">
          Showing {{ (page - 1) * perPage + 1 }} to {{ Math.min(page * perPage, totalPlayers) }} of {{ totalPlayers }} players
        </p>
        <div class="flex space-x-2">
          <UButton
            variant="outline"
            size="sm"
            :disabled="page === 1"
            @click="page--; fetchPlayers()"
          >
            Previous
          </UButton>
          <UButton
            variant="outline"
            size="sm"
            :disabled="page >= totalPages"
            @click="page++; fetchPlayers()"
          >
            Next
          </UButton>
        </div>
      </div>
    </UCard>

    <!-- Delete Confirmation Modal -->
    <Teleport to="body">
      <div v-if="showDeleteModal" class="fixed inset-0 z-50 overflow-y-auto">
        <div class="flex min-h-full items-center justify-center p-4">
          <div class="fixed inset-0 bg-black/50" @click="showDeleteModal = false"></div>
          
          <div class="relative bg-white rounded-xl shadow-xl max-w-md w-full p-6">
            <h3 class="text-lg font-semibold text-neutral-900 mb-4">Delete Player</h3>
            <p class="text-neutral-600 mb-6">
              Are you sure you want to delete {{ playerToDelete?.first_name }} {{ playerToDelete?.last_name }}? 
              This action cannot be undone.
            </p>

            <div class="flex justify-end space-x-3">
              <UButton variant="ghost" @click="showDeleteModal = false">Cancel</UButton>
              <UButton 
                variant="outline" 
                class="border-red-300 text-red-600 hover:bg-red-50"
                @click="deletePlayer"
                :loading="deleting"
              >
                Delete
              </UButton>
            </div>
          </div>
        </div>
      </div>
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

const filters = reactive({
  search: '',
  status: '',
  academy: '',
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
