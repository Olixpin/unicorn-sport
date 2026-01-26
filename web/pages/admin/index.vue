<template>
  <div>
    <div class="mb-8">
      <h1 class="font-display text-2xl lg:text-3xl font-bold text-neutral-900">
        Admin Dashboard
      </h1>
      <p class="mt-2 text-neutral-600">Platform overview and quick actions.</p>
    </div>

    <!-- Stats Grid -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
      <UCard>
        <div class="flex items-center">
          <div class="p-3 bg-primary-100 rounded-lg">
            <svg class="w-6 h-6 text-primary-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0z" />
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm text-neutral-500">Total Players</p>
            <p class="text-2xl font-bold text-neutral-900">{{ stats.totalPlayers }}</p>
          </div>
        </div>
      </UCard>

      <UCard>
        <div class="flex items-center">
          <div class="p-3 bg-secondary-100 rounded-lg">
            <svg class="w-6 h-6 text-secondary-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5" />
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm text-neutral-500">Academies</p>
            <p class="text-2xl font-bold text-neutral-900">{{ stats.totalAcademies }}</p>
          </div>
        </div>
      </UCard>

      <UCard>
        <div class="flex items-center">
          <div class="p-3 bg-blue-100 rounded-lg">
            <svg class="w-6 h-6 text-blue-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197" />
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm text-neutral-500">Subscribers</p>
            <p class="text-2xl font-bold text-neutral-900">{{ stats.totalSubscribers }}</p>
          </div>
        </div>
      </UCard>

      <UCard>
        <div class="flex items-center">
          <div class="p-3 bg-green-100 rounded-lg">
            <svg class="w-6 h-6 text-green-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm text-neutral-500">Monthly Revenue</p>
            <p class="text-2xl font-bold text-neutral-900">${{ stats.monthlyRevenue.toLocaleString() }}</p>
          </div>
        </div>
      </UCard>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- Recent Players -->
      <UCard>
        <template #header>
          <div class="flex items-center justify-between">
            <h2 class="font-semibold text-neutral-900">Recently Added Players</h2>
            <NuxtLink to="/admin/players">
              <UButton size="sm" variant="ghost">View All</UButton>
            </NuxtLink>
          </div>
        </template>

        <div class="space-y-3">
          <div v-for="player in recentPlayers" :key="player.id" class="flex items-center justify-between py-2">
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
                <p class="font-medium text-neutral-900 text-sm">{{ player.first_name }} {{ player.last_name }}</p>
                <p class="text-xs text-neutral-500">{{ player.position }}</p>
              </div>
            </div>
            <UBadge :variant="player.is_verified ? 'success' : 'default'" size="sm">
              {{ player.is_verified ? 'Verified' : 'Pending' }}
            </UBadge>
          </div>
        </div>
      </UCard>

      <!-- Pending Contact Requests -->
      <UCard>
        <template #header>
          <div class="flex items-center justify-between">
            <h2 class="font-semibold text-neutral-900">Pending Contact Requests</h2>
            <NuxtLink to="/admin/contacts">
              <UButton size="sm" variant="ghost">View All</UButton>
            </NuxtLink>
          </div>
        </template>

        <div v-if="pendingContacts.length === 0" class="text-center py-6 text-neutral-500">
          No pending requests
        </div>

        <div v-else class="space-y-3">
          <div v-for="contact in pendingContacts" :key="contact.id" class="flex items-center justify-between py-2">
            <div>
              <p class="font-medium text-neutral-900 text-sm">{{ contact.user_email }}</p>
              <p class="text-xs text-neutral-500">Requesting: {{ contact.player_name }}</p>
            </div>
            <div class="flex space-x-2">
              <UButton size="sm" variant="ghost" class="text-green-600" @click="approveContact(contact.id)">
                Approve
              </UButton>
              <UButton size="sm" variant="ghost" class="text-red-600" @click="rejectContact(contact.id)">
                Reject
              </UButton>
            </div>
          </div>
        </div>
      </UCard>
    </div>

    <!-- Quick Actions -->
    <div class="mt-8">
      <h2 class="text-lg font-semibold text-neutral-900 mb-4">Quick Actions</h2>
      <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
        <NuxtLink to="/admin/players/new">
          <UCard class="hover:border-primary-300 cursor-pointer transition-colors">
            <div class="flex items-center">
              <div class="p-2 bg-primary-100 rounded-lg">
                <svg class="w-5 h-5 text-primary-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
                </svg>
              </div>
              <span class="ml-3 font-medium text-neutral-900">Add New Player</span>
            </div>
          </UCard>
        </NuxtLink>

        <NuxtLink to="/admin/academies/new">
          <UCard class="hover:border-primary-300 cursor-pointer transition-colors">
            <div class="flex items-center">
              <div class="p-2 bg-secondary-100 rounded-lg">
                <svg class="w-5 h-5 text-secondary-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
                </svg>
              </div>
              <span class="ml-3 font-medium text-neutral-900">Add New Academy</span>
            </div>
          </UCard>
        </NuxtLink>

        <NuxtLink to="/admin/videos/upload">
          <UCard class="hover:border-primary-300 cursor-pointer transition-colors">
            <div class="flex items-center">
              <div class="p-2 bg-blue-100 rounded-lg">
                <svg class="w-5 h-5 text-blue-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
                </svg>
              </div>
              <span class="ml-3 font-medium text-neutral-900">Upload Video</span>
            </div>
          </UCard>
        </NuxtLink>
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

const stats = reactive({
  totalPlayers: 0,
  totalAcademies: 0,
  totalSubscribers: 0,
  monthlyRevenue: 0,
})

const recentPlayers = ref<Player[]>([])
const pendingContacts = ref<PendingContact[]>([])

async function loadDashboardData() {
  try {
    // Fetch admin stats
    const statsResponse = await api.get<ApiResponse<typeof stats>>('/admin/stats', {}, true)
    if (statsResponse.success && statsResponse.data) {
      Object.assign(stats, statsResponse.data)
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
