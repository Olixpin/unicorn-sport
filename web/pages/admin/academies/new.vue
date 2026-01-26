<template>
  <div>
    <div class="mb-8">
      <div class="flex items-center">
        <NuxtLink to="/admin/academies" class="text-neutral-500 hover:text-neutral-700">
          <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
          </svg>
        </NuxtLink>
        <h1 class="ml-4 font-display text-2xl lg:text-3xl font-bold text-neutral-900">
          Add New Academy
        </h1>
      </div>
      <p class="mt-2 text-neutral-600 ml-9">Register a new football academy.</p>
    </div>

    <form @submit.prevent="handleSubmit" class="max-w-3xl">
      <!-- Basic Information -->
      <UCard class="mb-6">
        <template #header>
          <h2 class="font-semibold text-neutral-900">Basic Information</h2>
        </template>

        <UInput
          v-model="form.name"
          label="Academy Name *"
          placeholder="Lagos Football Academy"
          :error="errors.name"
          class="mb-4"
        />

        <div>
          <label class="block text-sm font-medium text-neutral-700 mb-1">Description</label>
          <textarea
            v-model="form.description"
            rows="3"
            class="w-full px-3 py-2 border border-neutral-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500"
            placeholder="Brief description of the academy..."
          ></textarea>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mt-4">
          <UInput
            v-model.number="form.founded_year"
            label="Founded Year"
            type="number"
            placeholder="2005"
          />
          <UInput
            v-model="form.logo_url"
            label="Logo URL"
            placeholder="https://..."
          />
        </div>
      </UCard>

      <!-- Contact Information -->
      <UCard class="mb-6">
        <template #header>
          <h2 class="font-semibold text-neutral-900">Contact Information</h2>
        </template>

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <UInput
            v-model="form.email"
            label="Email"
            type="email"
            placeholder="info@academy.com"
            :error="errors.email"
          />
          <UInput
            v-model="form.phone"
            label="Phone"
            placeholder="+234 800 000 0000"
          />
        </div>

        <UInput
          v-model="form.website"
          label="Website"
          placeholder="https://www.academy.com"
          :error="errors.website"
          class="mt-4"
        />
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
            v-model="form.address"
            label="Address"
            placeholder="123 Football Lane"
          />
        </div>
      </UCard>

      <!-- Actions -->
      <div class="flex justify-end space-x-4">
        <NuxtLink to="/admin/academies">
          <UButton variant="ghost">Cancel</UButton>
        </NuxtLink>
        <UButton type="submit" :loading="submitting">
          Create Academy
        </UButton>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { academySchema, type AcademyFormData } from '~/schemas/academy'
import { AFRICAN_COUNTRIES } from '~/schemas/player'
import type { ApiResponse } from '~/types/index'

definePageMeta({
  layout: 'admin',
  middleware: 'admin',
})

const router = useRouter()
const toast = useToast()
const config = useRuntimeConfig()
const authStore = useAuthStore()

const form = reactive<Partial<AcademyFormData>>({
  name: '',
  description: '',
  country: '',
  state: '',
  city: '',
  address: '',
  phone: '',
  email: '',
  website: '',
  founded_year: undefined,
  logo_url: '',
})

const errors = reactive<Record<string, string>>({})
const submitting = ref(false)

function validateForm(): boolean {
  Object.keys(errors).forEach((key) => delete errors[key])

  const result = academySchema.safeParse(form)

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
    const response = await $fetch<ApiResponse<{ academy: { id: string } }>>('/admin/academies', {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: {
        Authorization: `Bearer ${authStore.accessToken}`,
      },
      body: form,
    })

    if (response.success) {
      toast.success('Academy Created', 'The academy has been successfully created')
      router.push('/admin/academies')
    } else {
      toast.error('Error', response.message || 'Failed to create academy')
    }
  } catch (error: unknown) {
    const err = error as { data?: { message?: string } }
    toast.error('Error', err.data?.message || 'An unexpected error occurred')
  } finally {
    submitting.value = false
  }
}
</script>
