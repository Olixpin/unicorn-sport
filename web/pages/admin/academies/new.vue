<template>
  <div class="max-w-6xl mx-auto">
    <!-- Page Header -->
    <div class="mb-8">
      <div class="flex items-center gap-4 mb-2">
        <NuxtLink 
          to="/admin/academies" 
          class="w-10 h-10 flex items-center justify-center rounded-xl bg-neutral-100 hover:bg-neutral-200 text-neutral-600 transition-colors"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
          </svg>
        </NuxtLink>
        <div>
          <h1 class="font-display text-2xl lg:text-3xl font-bold text-neutral-900">
            Add New Academy
          </h1>
          <p class="mt-1 text-neutral-600">Register a new football academy to the platform.</p>
        </div>
      </div>
    </div>

    <form @submit.prevent="handleSubmit" class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <!-- Main Content (2 columns) -->
      <div class="lg:col-span-2 space-y-6">
        <!-- Basic Information -->
        <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
          <div class="px-6 py-4 border-b border-neutral-200 bg-gradient-to-r from-violet-50 to-purple-50">
            <div class="flex items-center gap-3">
              <div class="w-10 h-10 bg-gradient-to-br from-violet-500 to-purple-600 rounded-xl flex items-center justify-center text-white shadow-lg shadow-violet-500/25">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
                </svg>
              </div>
              <div>
                <h2 class="font-semibold text-neutral-900">Basic Information</h2>
                <p class="text-sm text-neutral-500">Academy name and details</p>
              </div>
            </div>
          </div>

          <div class="p-6 space-y-5">
            <div>
              <label class="block text-sm font-medium text-neutral-700 mb-2">Academy Name *</label>
              <input
                v-model="form.name"
                type="text"
                placeholder="Lagos Football Academy"
                :class="[
                  'w-full px-4 py-3 border rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all',
                  errors.name ? 'border-red-300 bg-red-50' : 'border-neutral-200 bg-neutral-50'
                ]"
              />
              <p v-if="errors.name" class="mt-1.5 text-sm text-red-600">{{ errors.name }}</p>
            </div>

            <div>
              <label class="block text-sm font-medium text-neutral-700 mb-2">Description</label>
              <textarea
                v-model="form.description"
                rows="4"
                placeholder="Brief description of the academy, training philosophy, achievements..."
                class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all resize-none"
              ></textarea>
            </div>

            <div class="grid grid-cols-1 sm:grid-cols-2 gap-5">
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-2">Founded Year</label>
                <input
                  v-model.number="form.founded_year"
                  type="number"
                  placeholder="2005"
                  min="1900"
                  :max="new Date().getFullYear()"
                  class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-2">Logo URL</label>
                <input
                  v-model="form.logo_url"
                  type="url"
                  placeholder="https://example.com/logo.png"
                  class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
                />
              </div>
            </div>
          </div>
        </div>

        <!-- Contact Information -->
        <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
          <div class="px-6 py-4 border-b border-neutral-200 bg-gradient-to-r from-blue-50 to-indigo-50">
            <div class="flex items-center gap-3">
              <div class="w-10 h-10 bg-gradient-to-br from-blue-500 to-indigo-600 rounded-xl flex items-center justify-center text-white shadow-lg shadow-blue-500/25">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                </svg>
              </div>
              <div>
                <h2 class="font-semibold text-neutral-900">Contact Information</h2>
                <p class="text-sm text-neutral-500">How scouts can reach the academy</p>
              </div>
            </div>
          </div>

          <div class="p-6 space-y-5">
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-5">
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-2">Email</label>
                <input
                  v-model="form.email"
                  type="email"
                  placeholder="info@academy.com"
                  :class="[
                    'w-full px-4 py-3 border rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all',
                    errors.email ? 'border-red-300 bg-red-50' : 'border-neutral-200 bg-neutral-50'
                  ]"
                />
                <p v-if="errors.email" class="mt-1.5 text-sm text-red-600">{{ errors.email }}</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-2">Phone</label>
                <input
                  v-model="form.phone"
                  type="tel"
                  placeholder="+234 800 000 0000"
                  class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
                />
              </div>
            </div>

            <div>
              <label class="block text-sm font-medium text-neutral-700 mb-2">Website</label>
              <input
                v-model="form.website"
                type="url"
                placeholder="https://www.academy.com"
                :class="[
                  'w-full px-4 py-3 border rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all',
                  errors.website ? 'border-red-300 bg-red-50' : 'border-neutral-200 bg-neutral-50'
                ]"
              />
              <p v-if="errors.website" class="mt-1.5 text-sm text-red-600">{{ errors.website }}</p>
            </div>
          </div>
        </div>

        <!-- Location -->
        <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
          <div class="px-6 py-4 border-b border-neutral-200 bg-gradient-to-r from-emerald-50 to-green-50">
            <div class="flex items-center gap-3">
              <div class="w-10 h-10 bg-gradient-to-br from-emerald-500 to-green-600 rounded-xl flex items-center justify-center text-white shadow-lg shadow-emerald-500/25">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
              </div>
              <div>
                <h2 class="font-semibold text-neutral-900">Location</h2>
                <p class="text-sm text-neutral-500">Where the academy is based</p>
              </div>
            </div>
          </div>

          <div class="p-6 space-y-5">
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-5">
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-2">Country *</label>
                <div class="relative">
                  <select
                    v-model="form.country"
                    :class="[
                      'appearance-none w-full px-4 py-3 pr-10 border rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all',
                      errors.country ? 'border-red-300 bg-red-50' : 'border-neutral-200 bg-neutral-50'
                    ]"
                  >
                    <option value="">Select country</option>
                    <option v-for="country in AFRICAN_COUNTRIES" :key="country" :value="country">
                      {{ country }}
                    </option>
                  </select>
                  <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                    <svg class="w-5 h-5 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                    </svg>
                  </div>
                </div>
                <p v-if="errors.country" class="mt-1.5 text-sm text-red-600">{{ errors.country }}</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-2">State/Region</label>
                <input
                  v-model="form.state"
                  type="text"
                  placeholder="Lagos State"
                  class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
                />
              </div>
            </div>

            <div class="grid grid-cols-1 sm:grid-cols-2 gap-5">
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-2">City</label>
                <input
                  v-model="form.city"
                  type="text"
                  placeholder="Lagos"
                  class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-2">Address</label>
                <input
                  v-model="form.address"
                  type="text"
                  placeholder="123 Football Lane"
                  class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
                />
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Sidebar -->
      <div class="space-y-6">
        <!-- Logo Preview -->
        <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
          <div class="px-6 py-4 border-b border-neutral-200">
            <h3 class="font-semibold text-neutral-900">Logo Preview</h3>
          </div>
          <div class="p-6 flex flex-col items-center">
            <div class="w-24 h-24 rounded-xl bg-neutral-100 border-2 border-dashed border-neutral-300 flex items-center justify-center overflow-hidden">
              <img
                v-if="form.logo_url"
                :src="form.logo_url"
                :alt="form.name || 'Academy Logo'"
                class="w-full h-full object-cover"
                @error="form.logo_url = ''"
              />
              <svg v-else class="w-10 h-10 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
              </svg>
            </div>
            <p class="mt-3 text-sm text-neutral-500 text-center">
              Enter a valid image URL to preview the academy logo
            </p>
          </div>
        </div>

        <!-- Verification Status -->
        <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
          <div class="px-6 py-4 border-b border-neutral-200">
            <h3 class="font-semibold text-neutral-900">Initial Status</h3>
          </div>
          <div class="p-6">
            <label class="flex items-center gap-3 cursor-pointer">
              <div class="relative">
                <input
                  v-model="form.is_verified"
                  type="checkbox"
                  class="sr-only peer"
                />
                <div class="w-11 h-6 bg-neutral-200 rounded-full peer-checked:bg-emerald-500 transition-colors"></div>
                <div class="absolute left-0.5 top-0.5 w-5 h-5 bg-white rounded-full shadow peer-checked:translate-x-5 transition-transform"></div>
              </div>
              <div>
                <span class="text-sm font-medium text-neutral-900">Mark as Verified</span>
                <p class="text-xs text-neutral-500">Pre-approve this academy</p>
              </div>
            </label>
          </div>
        </div>

        <!-- Actions -->
        <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
          <div class="p-6 space-y-3">
            <button
              type="submit"
              :disabled="submitting"
              class="w-full inline-flex items-center justify-center gap-2 px-6 py-3 bg-gradient-to-r from-primary-600 to-emerald-600 text-white rounded-xl text-sm font-semibold hover:from-primary-700 hover:to-emerald-700 disabled:opacity-50 disabled:cursor-not-allowed transition-all shadow-lg shadow-primary-600/25"
            >
              <svg v-if="submitting" class="w-5 h-5 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" fill="none" />
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z" />
              </svg>
              <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
              </svg>
              {{ submitting ? 'Creating...' : 'Create Academy' }}
            </button>
            
            <NuxtLink to="/admin/academies" class="block">
              <button
                type="button"
                class="w-full px-6 py-3 border border-neutral-300 text-neutral-700 rounded-xl text-sm font-medium hover:bg-neutral-50 transition-colors"
              >
                Cancel
              </button>
            </NuxtLink>
          </div>
        </div>

        <!-- Tips -->
        <div class="bg-gradient-to-br from-amber-50 to-orange-50 rounded-2xl border border-amber-200 p-5">
          <div class="flex items-start gap-3">
            <div class="w-8 h-8 bg-amber-100 rounded-lg flex items-center justify-center flex-shrink-0">
              <svg class="w-5 h-5 text-amber-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
            <div>
              <h4 class="text-sm font-semibold text-amber-900 mb-1">Tips for adding academies</h4>
              <ul class="text-xs text-amber-700 space-y-1">
                <li>• Use the official academy name</li>
                <li>• Add accurate contact details</li>
                <li>• Include a high-quality logo</li>
                <li>• Verify location information</li>
              </ul>
            </div>
          </div>
        </div>
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

const form = reactive<Partial<AcademyFormData> & { is_verified?: boolean }>({
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
  is_verified: false,
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
