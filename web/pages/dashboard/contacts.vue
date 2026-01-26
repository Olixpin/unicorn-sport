<template>
  <div>
    <!-- Header -->
    <div class="mb-8">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="font-display text-2xl lg:text-3xl font-bold text-neutral-900">
            Contact Requests
          </h1>
          <p class="mt-1 text-neutral-500">Players you've requested to contact through their academies</p>
        </div>
        <div v-if="contactRequests.length > 0" class="hidden sm:flex items-center gap-3">
          <span class="px-3 py-1.5 bg-neutral-100 text-neutral-600 text-sm font-medium rounded-lg">
            {{ contactRequests.length }} {{ contactRequests.length === 1 ? 'request' : 'requests' }}
          </span>
        </div>
      </div>
    </div>

    <!-- Subscription Gate -->
    <div v-if="!subscriptionStore.canContactPlayers" class="mb-8">
      <div class="relative bg-gradient-to-br from-amber-50 to-orange-50 rounded-2xl p-6 border border-amber-200 overflow-hidden">
        <div class="absolute top-0 right-0 w-32 h-32 bg-gradient-to-br from-amber-200/50 to-transparent rounded-full -translate-y-1/2 translate-x-1/2"></div>
        <div class="relative flex items-center gap-4">
          <div class="w-12 h-12 rounded-xl bg-amber-100 flex items-center justify-center flex-shrink-0">
            <svg class="w-6 h-6 text-amber-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
            </svg>
          </div>
          <div class="flex-1">
            <p class="font-semibold text-amber-900">Upgrade to contact players</p>
            <p class="text-sm text-amber-700 mt-0.5">Scout plan or higher required to request player contact.</p>
          </div>
          <NuxtLink to="/pricing">
            <button class="px-5 py-2.5 bg-amber-500 text-white font-semibold rounded-xl hover:bg-amber-600 transition-colors shadow-lg shadow-amber-500/25">
              Upgrade
            </button>
          </NuxtLink>
        </div>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex justify-center py-16">
      <div class="w-10 h-10 border-4 border-primary-500 border-t-transparent rounded-full animate-spin"></div>
    </div>

    <!-- Empty State -->
    <div v-else-if="contactRequests.length === 0" class="bg-white rounded-2xl border border-neutral-200 p-12 text-center">
      <div class="w-20 h-20 mx-auto bg-gradient-to-br from-amber-100 to-orange-100 rounded-2xl flex items-center justify-center mb-6">
        <svg class="w-10 h-10 text-amber-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
        </svg>
      </div>
      <h3 class="text-xl font-semibold text-neutral-900 mb-2">No contact requests yet</h3>
      <p class="text-neutral-500 mb-6 max-w-sm mx-auto">Find a player you're interested in and request to contact them through their academy.</p>
      <NuxtLink to="/discover">
        <button class="px-6 py-3 bg-primary-500 text-white font-semibold rounded-xl hover:bg-primary-600 transition-colors shadow-lg shadow-primary-500/25">
          Browse Players
        </button>
      </NuxtLink>
    </div>

    <!-- Contact Requests List -->
    <div v-else class="space-y-4">
      <div 
        v-for="request in contactRequests" 
        :key="request.id"
        class="bg-white rounded-2xl border border-neutral-200 p-5 hover:shadow-md hover:border-neutral-300 transition-all duration-200"
      >
        <div class="flex items-start justify-between">
          <div class="flex items-center gap-4">
            <div class="w-16 h-16 rounded-xl bg-neutral-200 overflow-hidden flex-shrink-0">
              <img 
                v-if="request.player?.profile_photo_url"
                :src="request.player.profile_photo_url"
                :alt="request.player.first_name"
                class="w-full h-full object-cover"
              />
              <div v-else class="w-full h-full flex items-center justify-center bg-gradient-to-br from-neutral-200 to-neutral-300">
                <span class="text-neutral-500 text-lg font-medium">{{ request.player?.first_name?.charAt(0) }}</span>
              </div>
            </div>
            <div>
              <p class="font-semibold text-lg text-neutral-900">
                {{ request.player?.first_name }} {{ request.player?.last_name }}
              </p>
              <p class="text-sm text-neutral-500">
                {{ request.player?.position }} • {{ request.player?.academy_name || 'Academy' }}
              </p>
              <div class="flex items-center gap-2 mt-1.5">
                <svg class="w-4 h-4 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <span class="text-xs text-neutral-400">Requested {{ formatDate(request.created_at) }}</span>
              </div>
            </div>
          </div>
          
          <div class="flex items-center gap-3">
            <span 
              :class="[
                'px-3 py-1 text-xs font-bold rounded-full',
                request.status === 'approved' ? 'bg-green-100 text-green-700' : '',
                request.status === 'pending' ? 'bg-amber-100 text-amber-700' : '',
                request.status === 'rejected' ? 'bg-red-100 text-red-700' : ''
              ]"
            >
              {{ request.status.toUpperCase() }}
            </span>
            <NuxtLink :to="`/players/${request.player_id}`">
              <button class="px-4 py-2 text-sm font-medium text-neutral-600 hover:text-neutral-900 hover:bg-neutral-100 rounded-lg transition-colors">
                View Profile
              </button>
            </NuxtLink>
          </div>
        </div>

        <!-- Contact Info (if approved) -->
        <div v-if="request.status === 'approved' && request.contact_info" class="mt-5 pt-5 border-t border-neutral-100">
          <p class="text-xs font-semibold text-neutral-400 uppercase tracking-wider mb-3">Contact Information</p>
          <div class="bg-green-50 rounded-xl p-4 border border-green-100">
            <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
              <div v-if="request.contact_info.email">
                <p class="text-xs text-neutral-500 mb-1">Email</p>
                <a :href="`mailto:${request.contact_info.email}`" class="text-primary-600 hover:text-primary-700 font-medium text-sm">
                  {{ request.contact_info.email }}
                </a>
              </div>
              <div v-if="request.contact_info.phone">
                <p class="text-xs text-neutral-500 mb-1">Phone</p>
                <p class="font-medium text-sm text-neutral-900">{{ request.contact_info.phone }}</p>
              </div>
              <div v-if="request.contact_info.academy_contact">
                <p class="text-xs text-neutral-500 mb-1">Academy Contact</p>
                <p class="font-medium text-sm text-neutral-900">{{ request.contact_info.academy_contact }}</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Rejection Reason -->
        <div v-if="request.status === 'rejected' && request.rejection_reason" class="mt-5 pt-5 border-t border-neutral-100">
          <div class="bg-red-50 rounded-xl p-4 border border-red-100">
            <p class="text-sm text-red-700">
              <span class="font-semibold">Reason:</span> {{ request.rejection_reason }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Player, ApiResponse } from '~/types'

interface ContactRequest {
  id: string
  player_id: string
  player?: Player
  status: 'pending' | 'approved' | 'rejected'
  created_at: string
  contact_info?: {
    email?: string
    phone?: string
    academy_contact?: string
  }
  rejection_reason?: string
}

definePageMeta({
  layout: 'dashboard',
  middleware: 'auth',
})

const subscriptionStore = useSubscriptionStore()
const api = useApi()

const contactRequests = ref<ContactRequest[]>([])
const loading = ref(true)

function formatDate(dateString: string): string {
  const date = new Date(dateString)
  return date.toLocaleDateString('en-US', {
    month: 'short',
    day: 'numeric',
    year: 'numeric',
  })
}

async function fetchContactRequests() {
  try {
    const response = await api.get<ApiResponse<{ requests: ContactRequest[] }>>(
      '/contact-requests',
      {},
      true
    )
    
    if (response.success && response.data) {
      contactRequests.value = response.data.requests || []
    }
  } catch (error) {
    console.error('Failed to fetch contact requests:', error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchContactRequests()
})
</script>
