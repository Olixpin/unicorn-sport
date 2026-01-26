<template>
  <div class="max-w-7xl mx-auto">
    <!-- Page Header -->
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4 mb-8">
      <div>
        <h1 class="font-display text-2xl lg:text-3xl font-bold text-neutral-900">
          Academies
        </h1>
        <p class="mt-1 text-neutral-600">Manage football academies and training centers worldwide.</p>
      </div>
      <NuxtLink to="/admin/academies/new">
        <button class="inline-flex items-center gap-2 px-5 py-2.5 bg-gradient-to-r from-primary-600 to-emerald-600 text-white rounded-xl text-sm font-semibold hover:from-primary-700 hover:to-emerald-700 transition-all shadow-lg shadow-primary-600/25">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
          </svg>
          Add Academy
        </button>
      </NuxtLink>
    </div>

    <!-- Stats Cards -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 mb-6">
      <!-- Total Academies -->
      <div class="bg-gradient-to-br from-violet-500 to-purple-600 rounded-2xl p-5 text-white shadow-lg shadow-violet-500/25">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-violet-100 text-sm font-medium">Total Academies</p>
            <p class="text-3xl font-bold mt-1">{{ total }}</p>
          </div>
          <div class="w-12 h-12 bg-white/20 rounded-xl flex items-center justify-center">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
            </svg>
          </div>
        </div>
      </div>

      <!-- Verified -->
      <div class="bg-gradient-to-br from-emerald-500 to-green-600 rounded-2xl p-5 text-white shadow-lg shadow-emerald-500/25">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-emerald-100 text-sm font-medium">Verified</p>
            <p class="text-3xl font-bold mt-1">{{ verifiedCount }}</p>
          </div>
          <div class="w-12 h-12 bg-white/20 rounded-xl flex items-center justify-center">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z" />
            </svg>
          </div>
        </div>
      </div>

      <!-- Pending -->
      <div class="bg-gradient-to-br from-amber-500 to-orange-500 rounded-2xl p-5 text-white shadow-lg shadow-amber-500/25">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-amber-100 text-sm font-medium">Pending Review</p>
            <p class="text-3xl font-bold mt-1">{{ pendingCount }}</p>
          </div>
          <div class="w-12 h-12 bg-white/20 rounded-xl flex items-center justify-center">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>

      <!-- Total Players -->
      <div class="bg-gradient-to-br from-blue-500 to-indigo-600 rounded-2xl p-5 text-white shadow-lg shadow-blue-500/25">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-blue-100 text-sm font-medium">Total Players</p>
            <p class="text-3xl font-bold mt-1">{{ totalPlayers }}</p>
          </div>
          <div class="w-12 h-12 bg-white/20 rounded-xl flex items-center justify-center">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- Filters & Search -->
    <div class="bg-white rounded-2xl border border-neutral-200 p-4 mb-6 shadow-sm">
      <div class="flex flex-col lg:flex-row gap-4">
        <!-- Search Input -->
        <div class="flex-1 relative">
          <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
            <svg class="w-5 h-5 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </div>
          <input
            v-model="searchQuery"
            type="text"
            placeholder="Search by academy name, city..."
            class="w-full pl-12 pr-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
            @keyup.enter="fetchAcademies"
          />
        </div>

        <!-- Country Filter -->
        <div class="relative">
          <select
            v-model="countryFilter"
            class="appearance-none px-4 py-3 pr-10 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all min-w-[180px]"
            @change="fetchAcademies"
          >
            <option value="">All Countries</option>
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

        <!-- Status Filter -->
        <div class="flex items-center gap-2">
          <button
            v-for="status in statusOptions"
            :key="status.value"
            @click="statusFilter = status.value; fetchAcademies()"
            :class="[
              'px-4 py-2.5 rounded-xl text-sm font-medium transition-all',
              statusFilter === status.value
                ? status.activeClass
                : 'bg-neutral-100 text-neutral-600 hover:bg-neutral-200'
            ]"
          >
            {{ status.label }}
          </button>
        </div>

        <!-- Search Button -->
        <button
          @click="fetchAcademies"
          class="px-6 py-2.5 bg-primary-600 text-white rounded-xl text-sm font-semibold hover:bg-primary-700 transition-colors shadow-lg shadow-primary-600/25"
        >
          Search
        </button>

        <!-- Clear Filters -->
        <button
          v-if="hasActiveFilters"
          @click="clearFilters"
          class="px-4 py-2.5 text-neutral-600 hover:text-neutral-900 text-sm font-medium transition-colors"
        >
          Clear
        </button>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex flex-col items-center justify-center py-20">
      <div class="w-12 h-12 border-4 border-primary-200 border-t-primary-600 rounded-full animate-spin"></div>
      <p class="mt-4 text-neutral-600">Loading academies...</p>
    </div>

    <!-- Academies Grid -->
    <div v-else>
      <div v-if="academies.length > 0" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-6">
        <div 
          v-for="academy in academies" 
          :key="academy.id"
          class="bg-white rounded-2xl border border-neutral-200 overflow-hidden hover:shadow-lg hover:border-primary-200 transition-all group"
        >
          <!-- Academy Header -->
          <div class="relative h-32 bg-gradient-to-br from-violet-500 to-purple-600">
            <div class="absolute inset-0 bg-black/10"></div>
            <div class="absolute top-4 right-4">
              <span 
                v-if="academy.is_verified"
                class="inline-flex items-center gap-1 px-2.5 py-1 bg-white/20 backdrop-blur-sm rounded-full text-xs font-medium text-white"
              >
                <svg class="w-3.5 h-3.5" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M6.267 3.455a3.066 3.066 0 001.745-.723 3.066 3.066 0 013.976 0 3.066 3.066 0 001.745.723 3.066 3.066 0 012.812 2.812c.051.643.304 1.254.723 1.745a3.066 3.066 0 010 3.976 3.066 3.066 0 00-.723 1.745 3.066 3.066 0 01-2.812 2.812 3.066 3.066 0 00-1.745.723 3.066 3.066 0 01-3.976 0 3.066 3.066 0 00-1.745-.723 3.066 3.066 0 01-2.812-2.812 3.066 3.066 0 00-.723-1.745 3.066 3.066 0 010-3.976 3.066 3.066 0 00.723-1.745 3.066 3.066 0 012.812-2.812zm7.44 5.252a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                </svg>
                Verified
              </span>
              <span 
                v-else
                class="inline-flex items-center gap-1 px-2.5 py-1 bg-amber-500/90 backdrop-blur-sm rounded-full text-xs font-medium text-white"
              >
                <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                Pending
              </span>
            </div>
            <!-- Logo -->
            <div class="absolute -bottom-8 left-6">
              <div class="w-16 h-16 rounded-xl bg-white shadow-lg border-4 border-white overflow-hidden">
                <img
                  v-if="academy.logo_url"
                  :src="academy.logo_url"
                  :alt="academy.name"
                  class="w-full h-full object-cover"
                />
                <div v-else class="w-full h-full bg-gradient-to-br from-secondary-400 to-secondary-600 flex items-center justify-center">
                  <span class="text-white font-bold text-lg">{{ academy.name.charAt(0) }}</span>
                </div>
              </div>
            </div>
          </div>

          <!-- Academy Content -->
          <div class="pt-12 px-6 pb-6">
            <h3 class="font-semibold text-lg text-neutral-900 mb-1 group-hover:text-primary-600 transition-colors">
              {{ academy.name }}
            </h3>
            <p class="text-sm text-neutral-500 mb-4 flex items-center gap-1">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
              </svg>
              {{ academy.city || 'Unknown City' }}, {{ academy.country }}
            </p>

            <div class="flex items-center justify-between mb-4">
              <div class="flex items-center gap-4 text-sm">
                <div class="flex items-center gap-1.5 text-neutral-600">
                  <svg class="w-4 h-4 text-primary-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
                  </svg>
                  <span class="font-medium">{{ academy.player_count || 0 }}</span> players
                </div>
                <div v-if="academy.founded_year" class="flex items-center gap-1.5 text-neutral-600">
                  <svg class="w-4 h-4 text-secondary-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                  </svg>
                  Est. {{ academy.founded_year }}
                </div>
              </div>
            </div>

            <div class="flex items-center gap-2">
              <NuxtLink :to="`/admin/academies/${academy.id}`" class="flex-1">
                <button class="w-full px-4 py-2.5 bg-neutral-100 text-neutral-700 rounded-xl text-sm font-medium hover:bg-neutral-200 transition-colors">
                  Edit Details
                </button>
              </NuxtLink>
              <button
                v-if="!academy.is_verified"
                @click="verifyAcademy(academy.id)"
                class="px-4 py-2.5 bg-emerald-50 text-emerald-700 rounded-xl text-sm font-medium hover:bg-emerald-100 transition-colors"
              >
                Verify
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div v-else class="bg-white rounded-2xl border border-neutral-200 px-6 py-16 text-center">
        <div class="w-16 h-16 bg-neutral-100 rounded-2xl flex items-center justify-center mx-auto mb-4">
          <svg class="w-8 h-8 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
          </svg>
        </div>
        <h3 class="font-semibold text-neutral-900 mb-1">No academies found</h3>
        <p class="text-neutral-500 text-sm mb-6">
          {{ hasActiveFilters ? 'Try adjusting your filters to see more results.' : 'Get started by adding your first academy.' }}
        </p>
        <NuxtLink v-if="!hasActiveFilters" to="/admin/academies/new">
          <button class="inline-flex items-center gap-2 px-5 py-2.5 bg-gradient-to-r from-primary-600 to-emerald-600 text-white rounded-xl text-sm font-semibold hover:from-primary-700 hover:to-emerald-700 transition-all shadow-lg shadow-primary-600/25">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
            </svg>
            Add First Academy
          </button>
        </NuxtLink>
        <button
          v-else
          @click="clearFilters"
          class="inline-flex items-center gap-2 px-4 py-2 text-primary-600 hover:text-primary-700 font-medium text-sm transition-colors"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
          Clear filters
        </button>
      </div>

      <!-- Pagination -->
      <div v-if="totalPages > 1" class="mt-6 bg-white rounded-2xl border border-neutral-200 px-6 py-4 shadow-sm">
        <div class="flex flex-col sm:flex-row items-center justify-between gap-4">
          <p class="text-sm text-neutral-600">
            Showing <span class="font-medium">{{ (page - 1) * perPage + 1 }}</span> to 
            <span class="font-medium">{{ Math.min(page * perPage, total) }}</span> of 
            <span class="font-medium">{{ total }}</span> academies
          </p>
          <div class="flex items-center gap-2">
            <button
              :disabled="page === 1"
              @click="page--; fetchAcademies()"
              class="inline-flex items-center gap-1 px-3 py-2 border border-neutral-300 rounded-lg text-sm font-medium text-neutral-700 bg-white hover:bg-neutral-50 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
              </svg>
              Previous
            </button>
            
            <div class="hidden sm:flex items-center gap-1">
              <button
                v-for="pageNum in visiblePages"
                :key="pageNum"
                @click="page = pageNum; fetchAcademies()"
                :class="[
                  'w-10 h-10 rounded-lg text-sm font-medium transition-colors',
                  page === pageNum
                    ? 'bg-primary-600 text-white shadow-lg shadow-primary-600/25'
                    : 'text-neutral-700 hover:bg-neutral-100'
                ]"
              >
                {{ pageNum }}
              </button>
            </div>

            <button
              :disabled="page === totalPages"
              @click="page++; fetchAcademies()"
              class="inline-flex items-center gap-1 px-3 py-2 border border-neutral-300 rounded-lg text-sm font-medium text-neutral-700 bg-white hover:bg-neutral-50 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
            >
              Next
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
              </svg>
            </button>
          </div>
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
const toast = useToast()

const academies = ref<Academy[]>([])
const loading = ref(true)
const searchQuery = ref('')
const countryFilter = ref('')
const statusFilter = ref('')
const page = ref(1)
const perPage = ref(12)
const total = ref(0)
const totalPages = computed(() => Math.ceil(total.value / perPage.value))

const verifiedCount = computed(() => academies.value.filter(a => a.is_verified).length)
const pendingCount = computed(() => academies.value.filter(a => !a.is_verified).length)
const totalPlayers = computed(() => academies.value.reduce((sum, a) => sum + (a.player_count || 0), 0))

const hasActiveFilters = computed(() => searchQuery.value || countryFilter.value || statusFilter.value)

const statusOptions = [
  { value: '', label: 'All', activeClass: 'bg-neutral-900 text-white' },
  { value: 'verified', label: 'Verified', activeClass: 'bg-emerald-500 text-white' },
  { value: 'pending', label: 'Pending', activeClass: 'bg-amber-500 text-white' },
]

const visiblePages = computed(() => {
  const pages: number[] = []
  const maxVisible = 5
  let start = Math.max(1, page.value - Math.floor(maxVisible / 2))
  const end = Math.min(totalPages.value, start + maxVisible - 1)
  
  if (end - start < maxVisible - 1) {
    start = Math.max(1, end - maxVisible + 1)
  }
  
  for (let i = start; i <= end; i++) {
    pages.push(i)
  }
  return pages
})

function clearFilters() {
  searchQuery.value = ''
  countryFilter.value = ''
  statusFilter.value = ''
  page.value = 1
  fetchAcademies()
}

async function fetchAcademies() {
  loading.value = true

  try {
    const params = new URLSearchParams()
    params.append('page', page.value.toString())
    params.append('per_page', perPage.value.toString())
    if (searchQuery.value) params.append('q', searchQuery.value)
    if (countryFilter.value) params.append('country', countryFilter.value)
    if (statusFilter.value) params.append('status', statusFilter.value)

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

async function verifyAcademy(academyId: string) {
  try {
    const response = await $fetch<ApiResponse<null>>(`/admin/academies/${academyId}/verify`, {
      baseURL: config.public.apiBase,
      method: 'PATCH',
      headers: {
        Authorization: `Bearer ${authStore.accessToken}`,
      },
    })

    if (response.success) {
      toast.success('Academy Verified', 'The academy has been successfully verified.')
      fetchAcademies()
    }
  } catch (error: unknown) {
    const err = error as { data?: { message?: string } }
    toast.error('Error', err.data?.message || 'Failed to verify academy')
  }
}

onMounted(() => {
  fetchAcademies()
})
</script>
