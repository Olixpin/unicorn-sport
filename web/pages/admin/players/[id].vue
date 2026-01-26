<template>
  <div>
    <div class="mb-8">
      <div class="flex items-center justify-between">
        <div class="flex items-center">
          <NuxtLink to="/admin/players" class="text-neutral-500 hover:text-neutral-700">
            <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
            </svg>
          </NuxtLink>
          <h1 class="ml-4 font-display text-2xl lg:text-3xl font-bold text-neutral-900">
            Edit Player
          </h1>
        </div>
        <UBadge v-if="player" :variant="player.is_verified ? 'success' : 'warning'">
          {{ player.is_verified ? 'Verified' : 'Pending Verification' }}
        </UBadge>
      </div>
      <p class="mt-2 text-neutral-600 ml-9">Update player information.</p>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex justify-center py-12">
      <USpinner size="lg" />
    </div>

    <!-- Not Found -->
    <div v-else-if="!player" class="text-center py-12">
      <p class="text-neutral-500">Player not found</p>
      <NuxtLink to="/admin/players">
        <UButton variant="outline" class="mt-4">Back to Players</UButton>
      </NuxtLink>
    </div>

    <!-- Edit Form -->
    <form v-else @submit.prevent="handleSubmit" class="max-w-3xl">
      <!-- Basic Information -->
      <UCard class="mb-6">
        <template #header>
          <h2 class="font-semibold text-neutral-900">Basic Information</h2>
        </template>

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <UInput
            v-model="form.first_name"
            label="First Name *"
            placeholder="John"
            :error="errors.first_name"
          />
          <UInput
            v-model="form.last_name"
            label="Last Name *"
            placeholder="Doe"
            :error="errors.last_name"
          />
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mt-4">
          <UInput
            v-model="form.date_of_birth"
            label="Date of Birth *"
            type="date"
            :error="errors.date_of_birth"
          />
          <div>
            <label class="block text-sm font-medium text-neutral-700 mb-1">Position *</label>
            <select
              v-model="form.position"
              class="w-full px-3 py-2 border border-neutral-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500"
            >
              <option value="">Select position</option>
              <option v-for="pos in PLAYER_POSITIONS" :key="pos" :value="pos">
                {{ pos }}
              </option>
            </select>
          </div>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-3 gap-4 mt-4">
          <div>
            <label class="block text-sm font-medium text-neutral-700 mb-1">Preferred Foot</label>
            <select
              v-model="form.preferred_foot"
              class="w-full px-3 py-2 border border-neutral-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500"
            >
              <option value="">Select</option>
              <option value="right">Right</option>
              <option value="left">Left</option>
              <option value="both">Both</option>
            </select>
          </div>
          <UInput
            v-model.number="form.height_cm"
            label="Height (cm)"
            type="number"
            placeholder="175"
          />
          <UInput
            v-model.number="form.weight_kg"
            label="Weight (kg)"
            type="number"
            placeholder="70"
          />
        </div>
      </UCard>

      <!-- Location -->
      <UCard class="mb-6">
        <template #header>
          <h2 class="font-semibold text-neutral-900">Location</h2>
        </template>

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm font-medium text-neutral-700 mb-1">Country *</label>
            <select
              v-model="form.country"
              class="w-full px-3 py-2 border border-neutral-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500"
            >
              <option value="">Select country</option>
              <option v-for="country in AFRICAN_COUNTRIES" :key="country" :value="country">
                {{ country }}
              </option>
            </select>
          </div>
          <UInput
            v-model="form.state"
            label="State/Region"
            placeholder="Lagos State"
          />
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mt-4">
          <UInput
            v-model="form.city"
            label="City"
            placeholder="Lagos"
          />
          <UInput
            v-model="form.school_name"
            label="School Name"
            placeholder="Secondary School"
          />
        </div>
      </UCard>

      <!-- Verification Status (Admin Only) -->
      <UCard class="mb-6">
        <template #header>
          <h2 class="font-semibold text-neutral-900">Verification Status</h2>
        </template>

        <div class="flex items-center justify-between">
          <div>
            <p class="font-medium text-neutral-900">Player Verification</p>
            <p class="text-sm text-neutral-500">Toggle to verify or unverify this player</p>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input
              type="checkbox"
              v-model="form.is_verified"
              class="sr-only peer"
            />
            <div class="w-11 h-6 bg-neutral-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-primary-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-neutral-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-primary-600"></div>
          </label>
        </div>
      </UCard>

      <!-- Actions -->
      <div class="flex justify-between">
        <UButton
          type="button"
          variant="outline"
          class="border-red-300 text-red-600 hover:bg-red-50"
          @click="confirmDelete"
        >
          Delete Player
        </UButton>
        <div class="flex space-x-4">
          <NuxtLink to="/admin/players">
            <UButton variant="ghost">Cancel</UButton>
          </NuxtLink>
          <UButton type="submit" :loading="submitting">
            Save Changes
          </UButton>
        </div>
      </div>
    </form>

    <!-- Delete Confirmation Modal -->
    <Teleport to="body">
      <div v-if="showDeleteModal" class="fixed inset-0 z-50 overflow-y-auto">
        <div class="flex min-h-full items-center justify-center p-4">
          <div class="fixed inset-0 bg-black/50" @click="showDeleteModal = false"></div>

          <div class="relative bg-white rounded-xl shadow-xl max-w-md w-full p-6">
            <h3 class="text-lg font-semibold text-neutral-900 mb-4">Delete Player</h3>
            <p class="text-neutral-600 mb-6">
              Are you sure you want to delete {{ player?.first_name }} {{ player?.last_name }}?
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
import { PLAYER_POSITIONS, AFRICAN_COUNTRIES } from '~/schemas/player'
import type { Player, ApiResponse } from '~/types/index'

definePageMeta({
  layout: 'admin',
  middleware: 'admin',
})

const route = useRoute()
const router = useRouter()
const toast = useToast()
const config = useRuntimeConfig()
const authStore = useAuthStore()

const playerId = route.params.id as string

const player = ref<Player | null>(null)
const loading = ref(true)
const submitting = ref(false)
const deleting = ref(false)
const showDeleteModal = ref(false)
const errors = reactive<Record<string, string>>({})

const form = reactive({
  first_name: '',
  last_name: '',
  date_of_birth: '',
  position: '',
  preferred_foot: '' as '' | 'left' | 'right' | 'both',
  height_cm: undefined as number | undefined,
  weight_kg: undefined as number | undefined,
  country: '',
  state: '',
  city: '',
  school_name: '',
  is_verified: false,
})

// Fetch player data
onMounted(async () => {
  try {
    const response = await $fetch<ApiResponse<{ player: Player }>>(`/admin/players/${playerId}`, {
      baseURL: config.public.apiBase,
      headers: {
        Authorization: `Bearer ${authStore.accessToken}`,
      },
    })

    if (response.success && response.data?.player) {
      player.value = response.data.player
      // Populate form
      Object.assign(form, {
        first_name: player.value.first_name,
        last_name: player.value.last_name,
        date_of_birth: player.value.date_of_birth.split('T')[0],
        position: player.value.position,
        preferred_foot: player.value.preferred_foot || '',
        height_cm: player.value.height_cm,
        weight_kg: player.value.weight_kg,
        country: player.value.country,
        state: player.value.state || '',
        city: player.value.city || '',
        school_name: player.value.school_name || '',
        is_verified: player.value.is_verified || false,
      })
    }
  } catch (error) {
    console.error('Failed to fetch player:', error)
    toast.error('Error', 'Failed to load player data')
  } finally {
    loading.value = false
  }
})

async function handleSubmit() {
  submitting.value = true

  try {
    const payload = {
      ...form,
      preferred_foot: form.preferred_foot || undefined,
    }

    const response = await $fetch<ApiResponse<{ player: Player }>>(`/admin/players/${playerId}`, {
      baseURL: config.public.apiBase,
      method: 'PUT',
      headers: {
        Authorization: `Bearer ${authStore.accessToken}`,
      },
      body: payload,
    })

    if (response.success) {
      toast.success('Player Updated', 'The player has been successfully updated')
      router.push('/admin/players')
    } else {
      toast.error('Error', response.message || 'Failed to update player')
    }
  } catch (error: unknown) {
    const err = error as { data?: { message?: string } }
    toast.error('Error', err.data?.message || 'An unexpected error occurred')
  } finally {
    submitting.value = false
  }
}

function confirmDelete() {
  showDeleteModal.value = true
}

async function deletePlayer() {
  deleting.value = true

  try {
    const response = await $fetch<ApiResponse<null>>(`/admin/players/${playerId}`, {
      baseURL: config.public.apiBase,
      method: 'DELETE',
      headers: {
        Authorization: `Bearer ${authStore.accessToken}`,
      },
    })

    if (response.success) {
      toast.success('Player Deleted', 'The player has been successfully deleted')
      router.push('/admin/players')
    } else {
      toast.error('Error', response.message || 'Failed to delete player')
    }
  } catch (error: unknown) {
    const err = error as { data?: { message?: string } }
    toast.error('Error', err.data?.message || 'An unexpected error occurred')
  } finally {
    deleting.value = false
    showDeleteModal.value = false
  }
}
</script>
