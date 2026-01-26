<template>
  <div>
    <div class="flex items-center justify-between mb-8">
      <div>
        <h1 class="font-display text-2xl lg:text-3xl font-bold text-neutral-900">
          Academies
        </h1>
        <p class="mt-1 text-neutral-600">Manage football academies and training centers.</p>
      </div>
      <NuxtLink to="/admin/academies/new">
        <UButton>
          <svg class="w-4 h-4 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
          </svg>
          Add Academy
        </UButton>
      </NuxtLink>
    </div>

    <!-- Search & Filters -->
    <div class="bg-white rounded-lg border border-neutral-200 p-4 mb-6">
      <div class="flex flex-col sm:flex-row gap-4">
        <div class="flex-1">
          <UInput
            v-model="searchQuery"
            placeholder="Search academies..."
            @keyup.enter="fetchAcademies"
          />
        </div>
        <select
          v-model="countryFilter"
          class="px-3 py-2 border border-neutral-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500"
          @change="fetchAcademies"
        >
          <option value="">All Countries</option>
          <option v-for="country in AFRICAN_COUNTRIES" :key="country" :value="country">
            {{ country }}
          </option>
        </select>
        <UButton variant="outline" @click="fetchAcademies">
          <svg class="w-4 h-4 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
          Search
        </UButton>
      </div>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="flex justify-center py-12">
      <USpinner size="lg" />
    </div>

    <!-- Academies Table -->
    <div v-else class="bg-white rounded-lg border border-neutral-200 overflow-hidden">
      <table class="min-w-full divide-y divide-neutral-200">
        <thead class="bg-neutral-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-neutral-500 uppercase tracking-wider">
              Academy
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-neutral-500 uppercase tracking-wider">
              Location
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-neutral-500 uppercase tracking-wider">
              Players
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-neutral-500 uppercase tracking-wider">
              Status
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-neutral-500 uppercase tracking-wider">
              Actions
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-neutral-200">
          <tr v-for="academy in academies" :key="academy.id" class="hover:bg-neutral-50">
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center">
                <div class="h-10 w-10 flex-shrink-0">
                  <img
                    v-if="academy.logo_url"
                    :src="academy.logo_url"
                    :alt="academy.name"
                    class="h-10 w-10 rounded-full object-cover"
                  />
                  <div v-else class="h-10 w-10 rounded-full bg-secondary-100 flex items-center justify-center">
                    <span class="text-secondary-600 font-medium">{{ academy.name.charAt(0) }}</span>
                  </div>
                </div>
                <div class="ml-4">
                  <div class="text-sm font-medium text-neutral-900">{{ academy.name }}</div>
                  <div v-if="academy.founded_year" class="text-sm text-neutral-500">
                    Est. {{ academy.founded_year }}
                  </div>
                </div>
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm text-neutral-900">{{ academy.city || '-' }}</div>
              <div class="text-sm text-neutral-500">{{ academy.country }}</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-neutral-500">
              {{ academy.player_count || 0 }} players
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <UBadge :variant="academy.is_verified ? 'success' : 'warning'">
                {{ academy.is_verified ? 'Verified' : 'Pending' }}
              </UBadge>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
              <div class="flex items-center justify-end space-x-2">
                <NuxtLink :to="`/admin/academies/${academy.id}`">
                  <UButton size="sm" variant="ghost">Edit</UButton>
                </NuxtLink>
              </div>
            </td>
          </tr>
          <tr v-if="academies.length === 0">
            <td colspan="5" class="px-6 py-12 text-center text-neutral-500">
              No academies found.
            </td>
          </tr>
        </tbody>
      </table>

      <!-- Pagination -->
      <div v-if="totalPages > 1" class="bg-white px-4 py-3 flex items-center justify-between border-t border-neutral-200">
        <div class="text-sm text-neutral-700">
          Showing {{ (page - 1) * perPage + 1 }} to {{ Math.min(page * perPage, total) }} of {{ total }} academies
        </div>
        <div class="flex space-x-2">
          <UButton
            size="sm"
            variant="outline"
            :disabled="page === 1"
            @click="page--; fetchAcademies()"
          >
            Previous
          </UButton>
          <UButton
            size="sm"
            variant="outline"
            :disabled="page === totalPages"
            @click="page++; fetchAcademies()"
          >
            Next
          </UButton>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { AFRICAN_COUNTRIES } from '~/schemas/player'
import type { Academy } from '~/schemas/academy'
import type { ApiResponse } from '~/types/index'

definePageMeta({
  layout: 'admin',
  middleware: 'admin',
})

const config = useRuntimeConfig()
const authStore = useAuthStore()

const academies = ref<Academy[]>([])
const loading = ref(true)
const searchQuery = ref('')
const countryFilter = ref('')
const page = ref(1)
const perPage = ref(20)
const total = ref(0)
const totalPages = computed(() => Math.ceil(total.value / perPage.value))

async function fetchAcademies() {
  loading.value = true

  try {
    const params = new URLSearchParams()
    params.append('page', page.value.toString())
    params.append('per_page', perPage.value.toString())
    if (searchQuery.value) params.append('q', searchQuery.value)
    if (countryFilter.value) params.append('country', countryFilter.value)

    const response = await $fetch<ApiResponse<{ academies: Academy[]; total: number }>>(
      `/admin/academies?${params.toString()}`,
      {
        baseURL: config.public.apiBase,
        headers: {
          Authorization: `Bearer ${authStore.accessToken}`,
        },
      }
    )

    if (response.success && response.data) {
      academies.value = response.data.academies || []
      total.value = response.data.total || 0
    }
  } catch (error) {
    console.error('Failed to fetch academies:', error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchAcademies()
})
</script>
