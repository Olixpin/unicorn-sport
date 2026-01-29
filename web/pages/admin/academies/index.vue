<template>
  <div class="max-w-7xl mx-auto">
    <!-- Page Header -->
    <div class="flex items-center justify-between gap-4 mb-6">
      <div>
        <h1 class="font-display text-xl sm:text-2xl lg:text-3xl font-bold text-neutral-900">
          Academies
        </h1>
        <p class="hidden sm:block mt-1 text-neutral-600">Manage football academies and training centers worldwide.</p>
      </div>
      <NuxtLink to="/admin/academies/new">
        <button class="inline-flex items-center gap-2 px-4 py-2.5 bg-primary-600 text-white rounded-xl text-sm font-semibold hover:bg-primary-700 transition-all shadow-md">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
          </svg>
          <span class="hidden sm:inline">Add Academy</span>
          <span class="sm:hidden">Add</span>
        </button>
      </NuxtLink>
    </div>

    <!-- Stats Cards - Compact horizontal scroll on mobile -->
    <div class="flex gap-3 mb-5 overflow-x-auto pb-2 -mx-4 px-4 sm:mx-0 sm:px-0 sm:grid sm:grid-cols-4 sm:overflow-visible scrollbar-hide">
      <!-- Total Academies -->
      <div class="flex-shrink-0 w-[140px] sm:w-auto bg-gradient-to-br from-violet-500 to-purple-600 rounded-xl p-4 text-white">
        <p class="text-violet-100 text-xs font-medium">Total</p>
        <p class="text-2xl font-bold">{{ stats.totalAcademies }}</p>
      </div>

      <!-- Verified -->
      <button 
        @click="setStatusFilter('verified')"
        class="flex-shrink-0 w-[140px] sm:w-auto bg-gradient-to-br from-emerald-500 to-green-600 rounded-xl p-4 text-white text-left transition-all active:scale-95"
        :class="{ 'ring-2 ring-white/60 ring-offset-2 ring-offset-emerald-600': statusFilter === 'verified' }"
      >
        <p class="text-emerald-100 text-xs font-medium">Verified</p>
        <p class="text-2xl font-bold">{{ stats.verifiedCount }}</p>
      </button>

      <!-- Pending -->
      <button 
        @click="setStatusFilter('pending')"
        class="flex-shrink-0 w-[140px] sm:w-auto bg-gradient-to-br from-amber-500 to-orange-500 rounded-xl p-4 text-white text-left transition-all active:scale-95"
        :class="{ 'ring-2 ring-white/60 ring-offset-2 ring-offset-amber-600': statusFilter === 'pending' }"
      >
        <p class="text-amber-100 text-xs font-medium">Pending</p>
        <p class="text-2xl font-bold">{{ stats.pendingCount }}</p>
      </button>

      <!-- Total Players -->
      <div class="flex-shrink-0 w-[140px] sm:w-auto bg-gradient-to-br from-blue-500 to-indigo-600 rounded-xl p-4 text-white">
        <p class="text-blue-100 text-xs font-medium">Players</p>
        <p class="text-2xl font-bold">{{ stats.totalPlayers }}</p>
      </div>
    </div>

    <!-- Mobile: Compact Search & Filter -->
    <div class="sm:hidden mb-4 space-y-3">
      <!-- Search -->
      <div class="relative">
        <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
          <svg class="w-5 h-5 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </div>
        <input
          v-model="searchQuery"
          type="text"
          placeholder="Search academies..."
          class="w-full pl-10 pr-4 py-2.5 bg-white border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500"
          @keyup.enter="fetchAcademies"
        />
      </div>
      
      <!-- Filters Row -->
      <div class="flex gap-2">
        <div class="flex-1 relative">
          <select
            v-model="countryFilter"
            class="w-full appearance-none pl-3 pr-10 py-2.5 bg-white border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500"
            @change="page = 1; fetchAcademies()"
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
        <button
          @click="fetchAcademies"
          class="px-5 py-2.5 bg-primary-600 text-white rounded-xl text-sm font-semibold"
        >
          Go
        </button>
      </div>
    </div>

    <!-- Desktop: Full Filter Bar -->
    <div class="hidden sm:block bg-white rounded-2xl border border-neutral-200 p-4 mb-6 shadow-sm">
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
            class="w-full pl-12 pr-10 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
          />
          <!-- Loading indicator for search -->
          <div v-if="isSearching" class="absolute inset-y-0 right-0 pr-4 flex items-center">
            <svg class="w-4 h-4 text-neutral-400 animate-spin" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
          </div>
          <!-- Clear search button -->
          <button 
            v-else-if="searchQuery" 
            @click="searchQuery = ''"
            class="absolute inset-y-0 right-0 pr-4 flex items-center text-neutral-400 hover:text-neutral-600"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <!-- Country Filter -->
        <div class="relative">
          <select
            v-model="countryFilter"
            class="appearance-none px-4 py-3 pr-10 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all min-w-[180px]"
            @change="page = 1; fetchAcademies()"
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

    <!-- Active Filters Chips (visible when filtering) -->
    <div v-if="hasActiveFilters" class="mb-4 flex items-center gap-2 flex-wrap">
      <span class="text-xs text-neutral-500">Filters:</span>
      <span v-if="searchQuery" class="inline-flex items-center gap-1 px-2 py-0.5 bg-primary-100 text-primary-700 rounded-full text-xs font-medium">
        "{{ searchQuery }}"
        <button @click="searchQuery = ''; fetchAcademies()" class="hover:text-primary-900">
          <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </span>
      <span v-if="countryFilter" class="inline-flex items-center gap-1 px-2 py-0.5 bg-violet-100 text-violet-700 rounded-full text-xs font-medium">
        {{ countryFilter }}
        <button @click="countryFilter = ''; fetchAcademies()" class="hover:text-violet-900">
          <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </span>
      <span v-if="statusFilter" class="inline-flex items-center gap-1 px-2 py-0.5 rounded-full text-xs font-medium"
        :class="statusFilter === 'verified' ? 'bg-emerald-100 text-emerald-700' : 'bg-amber-100 text-amber-700'"
      >
        {{ statusFilter === 'verified' ? 'Verified' : 'Pending' }}
        <button @click="statusFilter = ''; fetchAcademies()" :class="statusFilter === 'verified' ? 'hover:text-emerald-900' : 'hover:text-amber-900'">
          <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </span>
      <button @click="clearFilters" class="text-xs text-neutral-500 hover:text-neutral-700 underline">
        Clear
      </button>
      <span class="text-xs text-neutral-400 ml-auto">{{ total }} result{{ total !== 1 ? 's' : '' }}</span>
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
                @click="openVerifyModal(academy)"
                class="px-4 py-2.5 bg-emerald-50 text-emerald-700 rounded-xl text-sm font-medium hover:bg-emerald-100 transition-colors"
              >
                Verify
              </button>
              <button
                v-else
                @click="openRevokeModal(academy)"
                class="px-4 py-2.5 bg-amber-50 text-amber-700 rounded-xl text-sm font-medium hover:bg-amber-100 transition-colors"
              >
                Revoke
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

    <!-- Verification Modal (handles both verify and revoke) -->
    <AdminAcademyVerificationModal
      v-model="showModal"
      :academy="selectedAcademy"
      :mode="modalMode"
      :loading="verifying"
      @confirm="handleConfirmVerification"
      @close="closeModal"
    />
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
const isSearching = ref(false)
const searchQuery = ref('')
const countryFilter = ref('')
const statusFilter = ref('')
const page = ref(1)
const perPage = ref(12)
const total = ref(0)
const totalPages = computed(() => Math.ceil(total.value / perPage.value))

// Debounced search - auto-search as user types
let searchTimeout: ReturnType<typeof setTimeout> | null = null
watch(searchQuery, (newVal, oldVal) => {
  if (newVal !== oldVal) {
    if (searchTimeout) clearTimeout(searchTimeout)
    isSearching.value = true
    searchTimeout = setTimeout(() => {
      page.value = 1
      fetchAcademies()
    }, 400)
  }
})

// Stats from API
const stats = reactive({
  totalAcademies: 0,
  verifiedCount: 0,
  pendingCount: 0,
  totalPlayers: 0,
})

// Modal state - unified for both verify and revoke
const showModal = ref(false)
const selectedAcademy = ref<Academy | null>(null)
const modalMode = ref<'verify' | 'revoke'>('verify')
const verifying = ref(false)

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

function setStatusFilter(status: string) {
  // Toggle off if same filter clicked again
  if (statusFilter.value === status) {
    statusFilter.value = ''
  } else {
    statusFilter.value = status
  }
  page.value = 1
  fetchAcademies()
}

function openVerifyModal(academy: Academy) {
  selectedAcademy.value = academy
  modalMode.value = 'verify'
  showModal.value = true
}

function openRevokeModal(academy: Academy) {
  selectedAcademy.value = academy
  modalMode.value = 'revoke'
  showModal.value = true
}

function closeModal() {
  showModal.value = false
  selectedAcademy.value = null
}

async function handleConfirmVerification(data: { academy: Academy; reason?: string }) {
  verifying.value = true
  const isVerify = modalMode.value === 'verify'

  try {
    const response = await $fetch<ApiResponse<Academy>>(`/admin/academies/${data.academy.id}`, {
      baseURL: config.public.apiBase,
      method: 'PUT',
      headers: {
        Authorization: `Bearer ${authStore.accessToken}`,
      },
      body: {
        ...data.academy,
        is_verified: isVerify,
      },
    })

    if (response.success) {
      toast.success(
        isVerify ? 'Academy Verified' : 'Verification Revoked',
        isVerify 
          ? 'The academy has been successfully verified.' 
          : 'The academy verification has been revoked.'
      )
      fetchAcademies()
    } else {
      toast.error('Error', response.message || `Failed to ${isVerify ? 'verify' : 'revoke'} academy`)
    }
  } catch (error: unknown) {
    const err = error as { data?: { message?: string } }
    toast.error('Error', err.data?.message || `Failed to ${isVerify ? 'verify' : 'revoke'} academy`)
  } finally {
    verifying.value = false
    closeModal()
  }
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

    const response = await $fetch<ApiResponse<{ 
      academies: Academy[]
      total: number
      stats?: {
        total_academies: number
        verified_count: number
        pending_count: number
        total_players: number
      }
    }>>(
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
      
      // Update stats from API
      if (response.data.stats) {
        stats.totalAcademies = response.data.stats.total_academies
        stats.verifiedCount = response.data.stats.verified_count
        stats.pendingCount = response.data.stats.pending_count
        stats.totalPlayers = response.data.stats.total_players
      }
    }
  } catch (error) {
    console.error('Failed to fetch academies:', error)
  } finally {
    loading.value = false
    isSearching.value = false
  }
}

onMounted(() => {
  fetchAcademies()
})
</script>
