<template>
  <div class="max-w-7xl mx-auto">
    <!-- Welcome Header -->
    <div class="mb-8">
      <div class="flex flex-col lg:flex-row lg:items-center lg:justify-between gap-4">
        <div>
          <h1 class="font-display text-2xl lg:text-3xl font-bold text-neutral-900">
            Welcome back, {{ authStore.user?.first_name || 'Admin' }} 👋
          </h1>
          <p class="mt-1 text-neutral-500">Here's what's happening with your platform today.</p>
        </div>
        <div class="flex items-center gap-3">
          <span class="text-sm text-neutral-500">{{ currentDate }}</span>
          <button 
            @click="loadDashboardData"
            class="flex items-center gap-2 px-4 py-2 bg-white border border-neutral-200 rounded-xl text-sm font-medium text-neutral-700 hover:bg-neutral-50 transition-colors"
          >
            <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
            </svg>
            Refresh
          </button>
        </div>
      </div>
    </div>

    <!-- Stats Grid -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 lg:gap-6 mb-8">
      <!-- Total Players -->
      <div class="group relative bg-white rounded-2xl border border-neutral-200 p-6 hover:shadow-lg hover:shadow-primary-500/5 transition-all duration-300 overflow-hidden">
        <div class="absolute top-0 right-0 w-32 h-32 bg-gradient-to-br from-primary-500/5 to-transparent rounded-bl-full"></div>
        <div class="relative">
          <div class="flex items-center justify-between mb-4">
            <div class="w-12 h-12 bg-gradient-to-br from-primary-500 to-primary-600 rounded-xl flex items-center justify-center shadow-lg shadow-primary-500/30">
              <svg class="w-6 h-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0z" />
              </svg>
            </div>
            <div v-if="stats.growth.players !== 0" :class="['flex items-center gap-1 text-sm font-medium', stats.growth.players >= 0 ? 'text-primary-600' : 'text-red-500']">
              <svg v-if="stats.growth.players >= 0" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 10l7-7m0 0l7 7m-7-7v18" />
              </svg>
              <svg v-else class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14l-7 7m0 0l-7-7m7 7V3" />
              </svg>
              {{ Math.abs(stats.growth.players).toFixed(0) }}%
            </div>
            <span v-else class="text-xs text-neutral-400">—</span>
          </div>
          <div class="space-y-1">
            <p class="text-3xl font-bold text-neutral-900">{{ stats.totalPlayers.toLocaleString() }}</p>
            <p class="text-sm text-neutral-500">Total Players</p>
          </div>
        </div>
      </div>

      <!-- Academies -->
      <div class="group relative bg-white rounded-2xl border border-neutral-200 p-6 hover:shadow-lg hover:shadow-amber-500/5 transition-all duration-300 overflow-hidden">
        <div class="absolute top-0 right-0 w-32 h-32 bg-gradient-to-br from-amber-500/5 to-transparent rounded-bl-full"></div>
        <div class="relative">
          <div class="flex items-center justify-between mb-4">
            <div class="w-12 h-12 bg-gradient-to-br from-amber-500 to-amber-600 rounded-xl flex items-center justify-center shadow-lg shadow-amber-500/30">
              <svg class="w-6 h-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
              </svg>
            </div>
            <div v-if="stats.growth.academies !== 0" :class="['flex items-center gap-1 text-sm font-medium', stats.growth.academies >= 0 ? 'text-amber-600' : 'text-red-500']">
              <svg v-if="stats.growth.academies >= 0" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 10l7-7m0 0l7 7m-7-7v18" />
              </svg>
              <svg v-else class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14l-7 7m0 0l-7-7m7 7V3" />
              </svg>
              {{ Math.abs(stats.growth.academies).toFixed(0) }}%
            </div>
            <span v-else class="text-xs text-neutral-400">—</span>
          </div>
          <div class="space-y-1">
            <p class="text-3xl font-bold text-neutral-900">{{ stats.totalAcademies.toLocaleString() }}</p>
            <p class="text-sm text-neutral-500">Academies</p>
          </div>
        </div>
      </div>

      <!-- Active Scouts -->
      <div class="group relative bg-white rounded-2xl border border-neutral-200 p-6 hover:shadow-lg hover:shadow-blue-500/5 transition-all duration-300 overflow-hidden">
        <div class="absolute top-0 right-0 w-32 h-32 bg-gradient-to-br from-blue-500/5 to-transparent rounded-bl-full"></div>
        <div class="relative">
          <div class="flex items-center justify-between mb-4">
            <div class="w-12 h-12 bg-gradient-to-br from-blue-500 to-blue-600 rounded-xl flex items-center justify-center shadow-lg shadow-blue-500/30">
              <svg class="w-6 h-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197" />
              </svg>
            </div>
            <div v-if="stats.growth.scouts !== 0" :class="['flex items-center gap-1 text-sm font-medium', stats.growth.scouts >= 0 ? 'text-blue-600' : 'text-red-500']">
              <svg v-if="stats.growth.scouts >= 0" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 10l7-7m0 0l7 7m-7-7v18" />
              </svg>
              <svg v-else class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14l-7 7m0 0l-7-7m7 7V3" />
              </svg>
              {{ Math.abs(stats.growth.scouts).toFixed(0) }}%
            </div>
            <span v-else class="text-xs text-neutral-400">—</span>
          </div>
          <div class="space-y-1">
            <p class="text-3xl font-bold text-neutral-900">{{ stats.totalScouts.toLocaleString() }}</p>
            <p class="text-sm text-neutral-500">Active Scouts</p>
          </div>
        </div>
      </div>

      <!-- Total Videos -->
      <div class="group relative bg-white rounded-2xl border border-neutral-200 p-6 hover:shadow-lg hover:shadow-emerald-500/5 transition-all duration-300 overflow-hidden">
        <div class="absolute top-0 right-0 w-32 h-32 bg-gradient-to-br from-emerald-500/5 to-transparent rounded-bl-full"></div>
        <div class="relative">
          <div class="flex items-center justify-between mb-4">
            <div class="w-12 h-12 bg-gradient-to-br from-emerald-500 to-emerald-600 rounded-xl flex items-center justify-center shadow-lg shadow-emerald-500/30">
              <svg class="w-6 h-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
              </svg>
            </div>
            <span class="text-xs text-neutral-400">Total</span>
          </div>
          <div class="space-y-1">
            <p class="text-3xl font-bold text-neutral-900">{{ stats.totalVideos.toLocaleString() }}</p>
            <p class="text-sm text-neutral-500">Player Videos</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Quick Actions -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-4 mb-8">
      <NuxtLink 
        to="/admin/players/new"
        class="group flex items-center gap-4 p-5 bg-white border border-neutral-200 rounded-2xl hover:border-primary-300 hover:shadow-lg hover:shadow-primary-500/5 transition-all duration-300"
      >
        <div class="w-12 h-12 bg-primary-100 rounded-xl flex items-center justify-center group-hover:bg-primary-500 transition-colors">
          <svg class="w-6 h-6 text-primary-600 group-hover:text-white transition-colors" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-5a4 4 0 11-8 0 4 4 0 018 0zM3 20a6 6 0 0112 0v1H3v-1z" />
          </svg>
        </div>
        <div>
          <p class="font-semibold text-neutral-900 group-hover:text-primary-700 transition-colors">Add New Player</p>
          <p class="text-sm text-neutral-500">Register a new talent</p>
        </div>
        <svg class="w-5 h-5 text-neutral-400 ml-auto group-hover:text-primary-600 group-hover:translate-x-1 transition-all" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
        </svg>
      </NuxtLink>

      <NuxtLink 
        to="/admin/academies/new"
        class="group flex items-center gap-4 p-5 bg-white border border-neutral-200 rounded-2xl hover:border-amber-300 hover:shadow-lg hover:shadow-amber-500/5 transition-all duration-300"
      >
        <div class="w-12 h-12 bg-amber-100 rounded-xl flex items-center justify-center group-hover:bg-amber-500 transition-colors">
          <svg class="w-6 h-6 text-amber-600 group-hover:text-white transition-colors" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
          </svg>
        </div>
        <div>
          <p class="font-semibold text-neutral-900 group-hover:text-amber-700 transition-colors">Add Academy</p>
          <p class="text-sm text-neutral-500">Partner with a new academy</p>
        </div>
        <svg class="w-5 h-5 text-neutral-400 ml-auto group-hover:text-amber-600 group-hover:translate-x-1 transition-all" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
        </svg>
      </NuxtLink>

      <NuxtLink 
        to="/admin/videos"
        class="group flex items-center gap-4 p-5 bg-white border border-neutral-200 rounded-2xl hover:border-blue-300 hover:shadow-lg hover:shadow-blue-500/5 transition-all duration-300"
      >
        <div class="w-12 h-12 bg-blue-100 rounded-xl flex items-center justify-center group-hover:bg-blue-500 transition-colors">
          <svg class="w-6 h-6 text-blue-600 group-hover:text-white transition-colors" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
          </svg>
        </div>
        <div>
          <p class="font-semibold text-neutral-900 group-hover:text-blue-700 transition-colors">Upload Video</p>
          <p class="text-sm text-neutral-500">Add highlight reels</p>
        </div>
        <svg class="w-5 h-5 text-neutral-400 ml-auto group-hover:text-blue-600 group-hover:translate-x-1 transition-all" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
        </svg>
      </NuxtLink>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- Recent Players -->
      <div class="bg-white rounded-2xl border border-neutral-200 overflow-hidden">
        <div class="flex items-center justify-between p-6 border-b border-neutral-100">
          <div>
            <h2 class="font-semibold text-neutral-900">Recently Added Players</h2>
            <p class="text-sm text-neutral-500 mt-0.5">Latest talent registrations</p>
          </div>
          <NuxtLink 
            to="/admin/players"
            class="text-sm font-medium text-primary-600 hover:text-primary-700 flex items-center gap-1"
          >
            View All
            <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
            </svg>
          </NuxtLink>
        </div>

        <div v-if="recentPlayers.length === 0" class="p-12 text-center">
          <div class="w-16 h-16 bg-neutral-100 rounded-full flex items-center justify-center mx-auto mb-4">
            <svg class="w-8 h-8 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0z" />
            </svg>
          </div>
          <p class="text-neutral-600 font-medium">No players yet</p>
          <p class="text-sm text-neutral-400 mt-1">Add your first player to get started</p>
        </div>

        <div v-else class="divide-y divide-neutral-100">
          <div 
            v-for="player in recentPlayers" 
            :key="player.id" 
            class="flex items-center justify-between p-4 hover:bg-neutral-50 transition-colors"
          >
            <div class="flex items-center gap-4">
              <div class="w-12 h-12 rounded-xl bg-gradient-to-br from-neutral-100 to-neutral-200 overflow-hidden flex-shrink-0">
                <img 
                  v-if="player.profile_photo_url"
                  :src="player.profile_photo_url"
                  :alt="player.first_name"
                  class="w-full h-full object-cover"
                />
                <div v-else class="w-full h-full flex items-center justify-center text-neutral-400 font-semibold">
                  {{ player.first_name?.charAt(0) }}{{ player.last_name?.charAt(0) }}
                </div>
              </div>
              <div>
                <p class="font-medium text-neutral-900">{{ player.first_name }} {{ player.last_name }}</p>
                <div class="flex items-center gap-2 text-sm text-neutral-500">
                  <span>{{ player.position }}</span>
                  <span class="text-neutral-300">•</span>
                  <span>{{ player.country }}</span>
                  <template v-if="player.academy">
                    <span class="text-neutral-300">•</span>
                    <span class="text-amber-600">{{ player.academy.name }}</span>
                  </template>
                </div>
              </div>
            </div>
            <span 
              :class="[
                'px-3 py-1 rounded-full text-xs font-medium',
                player.verification_status === 'verified'
                  ? 'bg-primary-100 text-primary-700' 
                  : 'bg-amber-100 text-amber-700'
              ]"
            >
              {{ player.verification_status === 'verified' ? 'Verified' : 'Pending' }}
            </span>
          </div>
        </div>
      </div>

      <!-- Pending Contact Requests -->
      <div class="bg-white rounded-2xl border border-neutral-200 overflow-hidden">
        <div class="flex items-center justify-between p-6 border-b border-neutral-100">
          <div>
            <h2 class="font-semibold text-neutral-900">Pending Requests</h2>
            <p class="text-sm text-neutral-500 mt-0.5">Scout contact requests to review</p>
          </div>
          <NuxtLink 
            to="/admin/contacts"
            class="text-sm font-medium text-primary-600 hover:text-primary-700 flex items-center gap-1"
          >
            View All
            <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
            </svg>
          </NuxtLink>
        </div>

        <div v-if="pendingContacts.length === 0" class="p-12 text-center">
          <div class="w-16 h-16 bg-primary-50 rounded-full flex items-center justify-center mx-auto mb-4">
            <svg class="w-8 h-8 text-primary-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
            </svg>
          </div>
          <p class="text-neutral-600 font-medium">All caught up!</p>
          <p class="text-sm text-neutral-400 mt-1">No pending requests to review</p>
        </div>

        <div v-else class="divide-y divide-neutral-100">
          <div 
            v-for="contact in pendingContacts" 
            :key="contact.id" 
            class="p-4 hover:bg-neutral-50 transition-colors"
          >
            <div class="flex items-center justify-between mb-3">
              <div>
                <p class="font-medium text-neutral-900">{{ contact.user_email }}</p>
                <p class="text-sm text-neutral-500">Requesting: {{ contact.player_name }}</p>
              </div>
            </div>
            <div class="flex gap-2">
              <button 
                @click="approveContact(contact.id)"
                class="flex-1 flex items-center justify-center gap-2 px-4 py-2 bg-primary-50 text-primary-700 font-medium rounded-lg hover:bg-primary-100 transition-colors"
              >
                <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                </svg>
                Approve
              </button>
              <button 
                @click="rejectContact(contact.id)"
                class="flex-1 flex items-center justify-center gap-2 px-4 py-2 bg-neutral-100 text-neutral-600 font-medium rounded-lg hover:bg-red-50 hover:text-red-600 transition-colors"
              >
                <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                </svg>
                Reject
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Platform Activity -->
    <div class="mt-8 bg-white rounded-2xl border border-neutral-200 overflow-hidden">
      <div class="p-6 border-b border-neutral-100">
        <h2 class="font-semibold text-neutral-900">Platform Activity</h2>
        <p class="text-sm text-neutral-500 mt-0.5">Recent events across the platform</p>
      </div>
      
      <div class="p-6">
        <div class="flex flex-col items-center justify-center py-8 text-center">
          <div class="w-16 h-16 bg-neutral-100 rounded-full flex items-center justify-center mb-4">
            <svg class="w-8 h-8 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
          </div>
          <p class="text-neutral-600 font-medium">Analytics Coming Soon</p>
          <p class="text-sm text-neutral-400 mt-1 max-w-sm">We're working on detailed analytics and activity tracking for your platform.</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Player, ApiResponse } from '~/types'

interface PendingContact {
  id: string
  user_email: string
  player_name: string
}

definePageMeta({
  layout: 'admin',
  middleware: 'admin',
})

const api = useApi()
const authStore = useAuthStore()

const currentDate = computed(() => {
  return new Date().toLocaleDateString('en-US', {
    weekday: 'long',
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
})

const stats = reactive({
  totalPlayers: 0,
  totalAcademies: 0,
  totalScouts: 0,
  totalVideos: 0,
  verifiedPlayers: 0,
  pendingPlayers: 0,
  growth: {
    players: 0,
    academies: 0,
    scouts: 0,
  },
})

const recentPlayers = ref<Player[]>([])
const pendingContacts = ref<PendingContact[]>([])

interface StatsResponse {
  total_players: number
  total_academies: number
  total_scouts: number
  total_videos: number
  verified_players: number
  pending_players: number
  growth: {
    players: number
    academies: number
    scouts: number
  }
}

async function loadDashboardData() {
  try {
    // Fetch admin stats
    const statsResponse = await api.get<ApiResponse<StatsResponse>>('/admin/stats', {}, true)
    if (statsResponse.success && statsResponse.data) {
      const data = statsResponse.data
      stats.totalPlayers = data.total_players || 0
      stats.totalAcademies = data.total_academies || 0
      stats.totalScouts = data.total_scouts || 0
      stats.totalVideos = data.total_videos || 0
      stats.verifiedPlayers = data.verified_players || 0
      stats.pendingPlayers = data.pending_players || 0
      stats.growth = data.growth || { players: 0, academies: 0, scouts: 0 }
    }

    // Fetch recent players
    const playersResponse = await api.get<ApiResponse<{ players: Player[] }>>('/admin/players?limit=5&sort=created_at', {}, true)
    if (playersResponse.success && playersResponse.data) {
      recentPlayers.value = playersResponse.data.players || []
    }

    // Fetch pending contacts
    const contactsResponse = await api.get<ApiResponse<{ requests: PendingContact[] }>>('/admin/contact-requests?status=pending', {}, true)
    if (contactsResponse.success && contactsResponse.data) {
      pendingContacts.value = contactsResponse.data.requests || []
    }
  } catch (error) {
    console.error('Failed to load admin dashboard:', error)
  }
}

async function approveContact(id: string) {
  try {
    await api.put(`/admin/contact-requests/${id}/approve`, {}, true)
    pendingContacts.value = pendingContacts.value.filter(c => c.id !== id)
  } catch (error) {
    console.error('Failed to approve contact:', error)
  }
}

async function rejectContact(id: string) {
  try {
    await api.put(`/admin/contact-requests/${id}/reject`, {}, true)
    pendingContacts.value = pendingContacts.value.filter(c => c.id !== id)
  } catch (error) {
    console.error('Failed to reject contact:', error)
  }
}

onMounted(() => {
  loadDashboardData()
})
</script>
