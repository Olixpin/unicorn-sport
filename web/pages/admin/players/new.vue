<template>
  <div>
    <div class="mb-8">
      <div class="flex items-center">
        <NuxtLink to="/admin/players" class="text-neutral-500 hover:text-neutral-700">
          <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
          </svg>
        </NuxtLink>
        <h1 class="ml-4 font-display text-2xl lg:text-3xl font-bold text-neutral-900">
          Add New Player
        </h1>
      </div>
      <p class="mt-2 text-neutral-600 ml-9">Create a new player profile in the system.</p>
    </div>

    <form @submit.prevent="handleSubmit" class="max-w-3xl">
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
              :class="{ 'border-red-500': errors.position }"
            >
              <option value="">Select position</option>
              <option v-for="pos in PLAYER_POSITIONS" :key="pos" :value="pos">
                {{ pos }}
              </option>
            </select>
            <p v-if="errors.position" class="mt-1 text-sm text-red-500">{{ errors.position }}</p>
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
              :class="{ 'border-red-500': errors.country }"
            >
              <option value="">Select country</option>
              <option v-for="country in AFRICAN_COUNTRIES" :key="country" :value="country">
                {{ country }}
              </option>
            </select>
            <p v-if="errors.country" class="mt-1 text-sm text-red-500">{{ errors.country }}</p>
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

      <!-- Academy Assignment -->
      <UCard class="mb-6">
        <template #header>
          <h2 class="font-semibold text-neutral-900">Academy Assignment</h2>
        </template>

        <div>
          <label class="block text-sm font-medium text-neutral-700 mb-1">Academy</label>
          <select
            v-model="form.academy_id"
            class="w-full px-3 py-2 border border-neutral-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500"
          >
            <option value="">No academy assigned</option>
            <option v-for="academy in academies" :key="academy.id" :value="academy.id">
              {{ academy.name }}
            </option>
          </select>
        </div>
      </UCard>

      <!-- Actions -->
      <div class="flex justify-end space-x-4">
        <NuxtLink to="/admin/players">
          <UButton variant="ghost">Cancel</UButton>
        </NuxtLink>
        <UButton type="submit" :loading="submitting">
          Create Player
        </UButton>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { playerSchema, PLAYER_POSITIONS, AFRICAN_COUNTRIES } from '~/schemas/player'
import type { ApiResponse } from '~/types/index'

interface Academy {
  id: string
  name: string
}

definePageMeta({
  layout: 'admin',
  middleware: 'admin',
})

const router = useRouter()
const toast = useToast()
const config = useRuntimeConfig()

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
  academy_id: '',
})

const errors = reactive<Record<string, string>>({})
const submitting = ref(false)
const academies = ref<Academy[]>([])

// Fetch academies for dropdown
onMounted(async () => {
  try {
    const response = await $fetch<ApiResponse<{ academies: Academy[] }>>('/admin/academies', {
      baseURL: config.public.apiBase,
      headers: {
        Authorization: `Bearer ${useAuthStore().accessToken}`,
      },
    })
    if (response.success && response.data) {
      academies.value = response.data.academies || []
    }
  } catch (error) {
    console.error('Failed to fetch academies:', error)
  }
})

function validateForm(): boolean {
  // Clear previous errors
  Object.keys(errors).forEach((key) => delete errors[key])

  const result = playerSchema.safeParse({
    ...form,
    preferred_foot: form.preferred_foot || undefined,
    academy_id: form.academy_id || undefined,
  })

  if (!result.success) {
    result.error.errors.forEach((err) => {
      const field = err.path[0] as string
      errors[field] = err.message
    })
    return false
  }

  return true
}

async function handleSubmit() {
  if (!validateForm()) {
    toast.error('Validation Error', 'Please fix the errors in the form')
    return
  }

  submitting.value = true

  try {
    const authStore = useAuthStore()
    const payload = {
      ...form,
      preferred_foot: form.preferred_foot || undefined,
      academy_id: form.academy_id || undefined,
    }

    const response = await $fetch<ApiResponse<{ player: { id: string } }>>('/admin/players', {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: {
        Authorization: `Bearer ${authStore.accessToken}`,
      },
      body: payload,
    })

    if (response.success) {
      toast.success('Player Created', 'The player has been successfully created')
      router.push('/admin/players')
    } else {
      toast.error('Error', response.message || 'Failed to create player')
    }
  } catch (error: unknown) {
    const err = error as { data?: { message?: string } }
    toast.error('Error', err.data?.message || 'An unexpected error occurred')
  } finally {
    submitting.value = false
  }
}
</script>
