<template>
  <div class="max-w-7xl mx-auto">
    <!-- Page Header -->
    <div class="flex items-center justify-between gap-4 mb-6">
      <div>
        <h1 class="font-display text-xl sm:text-2xl lg:text-3xl font-bold text-neutral-900">
          Contact Requests
        </h1>
        <p class="hidden sm:block mt-1 text-neutral-600">Manage and review scout contact requests to players.</p>
      </div>
      <button
        @click="fetchContacts"
        class="inline-flex items-center justify-center gap-2 w-10 h-10 sm:w-auto sm:h-auto sm:px-4 sm:py-2.5 border border-neutral-200 rounded-xl text-sm font-medium text-neutral-700 bg-white hover:bg-neutral-50 transition-colors"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
        </svg>
        <span class="hidden sm:inline">Refresh</span>
      </button>
    </div>

    <!-- Stats Cards - Horizontal scroll on mobile -->
    <div class="flex gap-3 overflow-x-auto pb-2 -mx-4 px-4 sm:mx-0 sm:px-0 sm:grid sm:grid-cols-4 sm:overflow-visible mb-6">
      <!-- Total Requests -->
      <div class="flex-shrink-0 w-[130px] sm:w-auto bg-white rounded-xl border border-neutral-200 p-3 sm:p-4">
        <p class="text-[10px] sm:text-xs font-medium text-neutral-500 uppercase tracking-wide">Total</p>
        <p class="text-xl sm:text-2xl font-bold text-neutral-900 mt-0.5">{{ total }}</p>
      </div>

      <!-- Pending -->
      <button 
        @click="statusFilter = 'pending'; fetchContacts()"
        class="flex-shrink-0 w-[130px] sm:w-auto bg-white rounded-xl border border-neutral-200 p-3 sm:p-4 text-left hover:border-amber-300 hover:bg-amber-50/50 transition-all"
        :class="{ 'border-amber-400 bg-amber-50': statusFilter === 'pending' }"
      >
        <p class="text-[10px] sm:text-xs font-medium text-amber-600 uppercase tracking-wide">Pending</p>
        <p class="text-xl sm:text-2xl font-bold text-neutral-900 mt-0.5">{{ pendingCount }}</p>
      </button>

      <!-- Approved -->
      <button 
        @click="statusFilter = 'approved'; fetchContacts()"
        class="flex-shrink-0 w-[130px] sm:w-auto bg-white rounded-xl border border-neutral-200 p-3 sm:p-4 text-left hover:border-emerald-300 hover:bg-emerald-50/50 transition-all"
        :class="{ 'border-emerald-400 bg-emerald-50': statusFilter === 'approved' }"
      >
        <p class="text-[10px] sm:text-xs font-medium text-emerald-600 uppercase tracking-wide">Approved</p>
        <p class="text-xl sm:text-2xl font-bold text-neutral-900 mt-0.5">{{ approvedCount }}</p>
      </button>

      <!-- Rejected -->
      <button 
        @click="statusFilter = 'rejected'; fetchContacts()"
        class="flex-shrink-0 w-[130px] sm:w-auto bg-white rounded-xl border border-neutral-200 p-3 sm:p-4 text-left hover:border-rose-300 hover:bg-rose-50/50 transition-all"
        :class="{ 'border-rose-400 bg-rose-50': statusFilter === 'rejected' }"
      >
        <p class="text-[10px] sm:text-xs font-medium text-rose-600 uppercase tracking-wide">Rejected</p>
        <p class="text-xl sm:text-2xl font-bold text-neutral-900 mt-0.5">{{ rejectedCount }}</p>
      </button>
    </div>

    <!-- Filters & Search -->
    <div class="bg-white rounded-xl border border-neutral-200 p-3 sm:p-4 mb-6">
      <div class="flex flex-col sm:flex-row gap-3">
        <!-- Search Input -->
        <div class="flex-1 relative">
          <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
            <svg class="w-4 h-4 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </div>
          <input
            v-model="searchQuery"
            type="text"
            placeholder="Search by name or email..."
            class="w-full pl-10 pr-4 py-2.5 bg-neutral-50 border border-neutral-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
            @keyup.enter="fetchContacts"
          />
        </div>

        <!-- Status Filter Pills - Hidden on mobile since stats are clickable -->
        <div class="hidden sm:flex items-center gap-2">
          <button
            v-for="status in statusOptions"
            :key="status.value"
            @click="statusFilter = status.value; fetchContacts()"
            :class="[
              'px-3 py-2 rounded-lg text-xs font-medium transition-all',
              statusFilter === status.value
                ? 'bg-primary-600 text-white'
                : 'bg-neutral-100 text-neutral-600 hover:bg-neutral-200'
            ]"
          >
            {{ status.label }}
          </button>
        </div>

        <!-- Search Button -->
        <button
          @click="fetchContacts"
          class="w-full sm:w-auto px-5 py-2.5 bg-gradient-to-r from-primary-600 to-emerald-600 text-white rounded-lg text-sm font-semibold hover:from-primary-700 hover:to-emerald-700 transition-all shadow-lg shadow-primary-600/25"
        >
          Search
        </button>

        <!-- Clear Filters -->
        <button
          v-if="hasActiveFilters"
          @click="clearFilters"
          class="hidden sm:block px-3 py-2 text-neutral-500 hover:text-neutral-700 text-sm font-medium transition-colors"
        >
          Clear
        </button>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex flex-col items-center justify-center py-20">
      <div class="w-12 h-12 border-4 border-primary-200 border-t-primary-600 rounded-full animate-spin"></div>
      <p class="mt-4 text-neutral-600">Loading contact requests...</p>
    </div>

    <!-- Contacts Table -->
    <div v-else class="bg-white rounded-2xl border border-neutral-200 overflow-hidden shadow-sm">
      <!-- Table Header -->
      <div class="px-6 py-4 border-b border-neutral-200 bg-neutral-50/50">
        <div class="flex items-center justify-between">
          <h3 class="font-semibold text-neutral-900">All Requests</h3>
          <span class="text-sm text-neutral-500">{{ total }} total requests</span>
        </div>
      </div>

      <!-- Table -->
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-neutral-200">
          <thead>
            <tr class="bg-neutral-50">
              <th class="px-4 sm:px-6 py-3 sm:py-4 text-left text-xs font-semibold text-neutral-500 uppercase tracking-wider">
                Scout
              </th>
              <th class="px-4 sm:px-6 py-3 sm:py-4 text-left text-xs font-semibold text-neutral-500 uppercase tracking-wider hidden sm:table-cell">
                Player
              </th>
              <th class="px-4 sm:px-6 py-3 sm:py-4 text-left text-xs font-semibold text-neutral-500 uppercase tracking-wider hidden lg:table-cell">
                Message
              </th>
              <th class="px-4 sm:px-6 py-3 sm:py-4 text-left text-xs font-semibold text-neutral-500 uppercase tracking-wider">
                Status
              </th>
              <th class="px-4 sm:px-6 py-3 sm:py-4 text-left text-xs font-semibold text-neutral-500 uppercase tracking-wider hidden md:table-cell">
                Date
              </th>
              <th class="px-4 sm:px-6 py-3 sm:py-4 text-right text-xs font-semibold text-neutral-500 uppercase tracking-wider">
                Actions
              </th>
            </tr>
          </thead>
          <tbody class="divide-y divide-neutral-100">
            <tr 
              v-for="contact in contacts" 
              :key="contact.id" 
              class="hover:bg-neutral-50/50 transition-colors group"
            >
              <!-- Scout Info -->
              <td class="px-4 sm:px-6 py-3 sm:py-4">
                <div class="flex items-center gap-2 sm:gap-3">
                  <div class="w-8 h-8 sm:w-10 sm:h-10 bg-gradient-to-br from-purple-500 to-indigo-600 rounded-lg sm:rounded-xl flex items-center justify-center text-white font-semibold text-xs sm:text-sm shadow-lg shadow-purple-500/25 flex-shrink-0">
                    {{ getInitials(contact.scout_name) }}
                  </div>
                  <div class="min-w-0">
                    <div class="font-medium text-neutral-900 text-sm truncate">{{ contact.scout_name || 'Unknown' }}</div>
                    <div class="text-xs text-neutral-500 truncate max-w-[120px] sm:max-w-none hidden sm:block">{{ contact.scout_email }}</div>
                  </div>
                </div>
              </td>

              <!-- Player Info -->
              <td class="px-4 sm:px-6 py-3 sm:py-4 hidden sm:table-cell">
                <div class="flex items-center gap-2 sm:gap-3">
                  <div class="w-8 h-8 sm:w-10 sm:h-10 bg-gradient-to-br from-primary-500 to-emerald-600 rounded-lg sm:rounded-xl flex items-center justify-center text-white font-semibold text-xs sm:text-sm shadow-lg shadow-primary-500/25 flex-shrink-0">
                    {{ getInitials(contact.player_name) }}
                  </div>
                  <div class="font-medium text-neutral-900 text-sm truncate">{{ contact.player_name || 'Unknown' }}</div>
                </div>
              </td>

              <!-- Message -->
              <td class="px-4 sm:px-6 py-3 sm:py-4 hidden lg:table-cell">
                <div class="max-w-xs">
                  <p 
                    v-if="contact.message" 
                    class="text-sm text-neutral-600 truncate cursor-pointer hover:text-neutral-900 transition-colors"
                    :title="contact.message"
                    @click="showMessageModal(contact)"
                  >
                    {{ contact.message }}
                  </p>
                  <span v-else class="text-sm text-neutral-400 italic">No message</span>
                </div>
              </td>

              <!-- Status -->
              <td class="px-4 sm:px-6 py-3 sm:py-4">
                <span :class="getStatusBadgeClass(contact.status)" class="text-xs">
                  <span class="w-1.5 h-1.5 rounded-full" :class="getStatusDotClass(contact.status)"></span>
                  {{ capitalizeFirst(contact.status) }}
                </span>
              </td>

              <!-- Date -->
              <td class="px-4 sm:px-6 py-3 sm:py-4 hidden md:table-cell">
                <div>
                  <div class="text-sm font-medium text-neutral-900">{{ formatDate(contact.created_at) }}</div>
                  <div class="text-xs text-neutral-500">{{ formatTime(contact.created_at) }}</div>
                </div>
              </td>

              <!-- Actions -->
              <td class="px-6 py-4 text-right">
                <div v-if="contact.status === 'pending'" class="flex items-center justify-end gap-2 opacity-0 group-hover:opacity-100 transition-opacity">
                  <button
                    @click="updateStatus(contact.id, 'approved')"
                    :disabled="updating === contact.id"
                    class="inline-flex items-center gap-1.5 px-3 py-1.5 bg-emerald-50 text-emerald-700 rounded-lg text-sm font-medium hover:bg-emerald-100 transition-colors disabled:opacity-50"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                    </svg>
                    Approve
                  </button>
                  <button
                    @click="updateStatus(contact.id, 'rejected')"
                    :disabled="updating === contact.id"
                    class="inline-flex items-center gap-1.5 px-3 py-1.5 bg-rose-50 text-rose-700 rounded-lg text-sm font-medium hover:bg-rose-100 transition-colors disabled:opacity-50"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                    </svg>
                    Reject
                  </button>
                </div>
                <div v-else class="flex items-center justify-end">
                  <span v-if="contact.status === 'approved'" class="inline-flex items-center gap-1.5 text-emerald-600 text-sm">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                    Approved
                  </span>
                  <span v-else class="inline-flex items-center gap-1.5 text-rose-600 text-sm">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                    Rejected
                  </span>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Empty State -->
      <div v-if="contacts.length === 0 && !loading" class="px-6 py-16 text-center">
        <div class="w-16 h-16 bg-neutral-100 rounded-2xl flex items-center justify-center mx-auto mb-4">
          <svg class="w-8 h-8 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 4H6a2 2 0 00-2 2v12a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-2m-4-1v8m0 0l3-3m-3 3L9 8m-5 5h2.586a1 1 0 01.707.293l2.414 2.414a1 1 0 00.707.293h3.172a1 1 0 00.707-.293l2.414-2.414a1 1 0 01.707-.293H20" />
          </svg>
        </div>
        <h3 class="font-semibold text-neutral-900 mb-1">No contact requests found</h3>
        <p class="text-neutral-500 text-sm mb-6">
          {{ hasActiveFilters ? 'Try adjusting your filters to see more results.' : 'Contact requests from scouts will appear here.' }}
        </p>
        <button
          v-if="hasActiveFilters"
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
      <div v-if="totalPages > 1" class="px-6 py-4 border-t border-neutral-200 bg-neutral-50/50">
        <div class="flex flex-col sm:flex-row items-center justify-between gap-4">
          <p class="text-sm text-neutral-600">
            Showing <span class="font-medium">{{ (page - 1) * perPage + 1 }}</span> to 
            <span class="font-medium">{{ Math.min(page * perPage, total) }}</span> of 
            <span class="font-medium">{{ total }}</span> requests
          </p>
          <div class="flex items-center gap-2">
            <button
              :disabled="page === 1"
              @click="page--; fetchContacts()"
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
                @click="page = pageNum; fetchContacts()"
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
              @click="page++; fetchContacts()"
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

    <!-- Message Modal -->
    <Teleport to="body">
      <div 
        v-if="selectedContact" 
        class="fixed inset-0 z-50 flex items-center justify-center p-4"
        @click.self="selectedContact = null"
      >
        <div class="absolute inset-0 bg-black/50 backdrop-blur-sm"></div>
        <div class="relative bg-white rounded-2xl shadow-2xl max-w-lg w-full p-6">
          <button
            @click="selectedContact = null"
            class="absolute top-4 right-4 w-8 h-8 flex items-center justify-center rounded-lg hover:bg-neutral-100 transition-colors"
          >
            <svg class="w-5 h-5 text-neutral-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
          
          <div class="mb-4">
            <h3 class="text-lg font-semibold text-neutral-900">Contact Message</h3>
            <p class="text-sm text-neutral-500">From {{ selectedContact.scout_name }} to {{ selectedContact.player_name }}</p>
          </div>
          
          <div class="bg-neutral-50 rounded-xl p-4 mb-6">
            <p class="text-neutral-700 whitespace-pre-wrap">{{ selectedContact.message || 'No message provided' }}</p>
          </div>
          
          <div class="flex items-center justify-between text-sm text-neutral-500">
            <span>{{ formatDate(selectedContact.created_at) }} at {{ formatTime(selectedContact.created_at) }}</span>
            <span :class="getStatusBadgeClass(selectedContact.status)">
              {{ capitalizeFirst(selectedContact.status) }}
            </span>
          </div>
        </div>
      </div>
    </Teleport>
  </div>
</template>

<script setup lang="ts">
import type { ApiResponse } from '~/types/index'

interface ContactRequest {
  id: string
  scout_id: string
  scout_name?: string
  scout_email?: string
  player_id: string
  player_name?: string
  message?: string
  status: 'pending' | 'approved' | 'rejected'
  created_at: string
}

definePageMeta({
  layout: 'admin',
  middleware: 'admin',
})

const config = useRuntimeConfig()
const authStore = useAuthStore()
const toast = useToast()

const contacts = ref<ContactRequest[]>([])
const loading = ref(true)
const updating = ref<string | null>(null)
const searchQuery = ref('')
const statusFilter = ref('')
const page = ref(1)
const perPage = ref(20)
const total = ref(0)
const selectedContact = ref<ContactRequest | null>(null)

const totalPages = computed(() => Math.ceil(total.value / perPage.value))

const pendingCount = computed(() => contacts.value.filter(c => c.status === 'pending').length)
const approvedCount = computed(() => contacts.value.filter(c => c.status === 'approved').length)
const rejectedCount = computed(() => contacts.value.filter(c => c.status === 'rejected').length)

const hasActiveFilters = computed(() => searchQuery.value || statusFilter.value)

const statusOptions = [
  { value: '', label: 'All' },
  { value: 'pending', label: 'Pending' },
  { value: 'approved', label: 'Approved' },
  { value: 'rejected', label: 'Rejected' },
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
  statusFilter.value = ''
  page.value = 1
  fetchContacts()
}

function getInitials(name?: string): string {
  if (!name) return '?'
  return name.split(' ').map(n => n[0]).join('').toUpperCase().slice(0, 2)
}

function getStatusBadgeClass(status: string): string {
  const baseClass = 'inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full text-xs font-medium'
  switch (status) {
    case 'approved':
      return `${baseClass} bg-emerald-50 text-emerald-700`
    case 'rejected':
      return `${baseClass} bg-rose-50 text-rose-700`
    default:
      return `${baseClass} bg-amber-50 text-amber-700`
  }
}

function getStatusDotClass(status: string): string {
  switch (status) {
    case 'approved':
      return 'bg-emerald-500'
    case 'rejected':
      return 'bg-rose-500'
    default:
      return 'bg-amber-500'
  }
}

function capitalizeFirst(str: string): string {
  return str.charAt(0).toUpperCase() + str.slice(1)
}

function showMessageModal(contact: ContactRequest) {
  selectedContact.value = contact
}

async function fetchContacts() {
  loading.value = true

  try {
    const params = new URLSearchParams()
    params.append('page', page.value.toString())
    params.append('per_page', perPage.value.toString())
    if (searchQuery.value) params.append('q', searchQuery.value)
    if (statusFilter.value) params.append('status', statusFilter.value)

    const response = await $fetch<ApiResponse<{ contacts: ContactRequest[]; total: number }>>(
      `/admin/contacts?${params.toString()}`,
      {
        baseURL: config.public.apiBase,
        headers: {
          Authorization: `Bearer ${authStore.accessToken}`,
        },
      }
    )

    if (response.success && response.data) {
      contacts.value = response.data.contacts || []
      total.value = response.data.total || 0
    }
  } catch (error) {
    console.error('Failed to fetch contacts:', error)
  } finally {
    loading.value = false
  }
}

async function updateStatus(contactId: string, status: 'approved' | 'rejected') {
  updating.value = contactId
  try {
    const response = await $fetch<ApiResponse<null>>(`/admin/contacts/${contactId}/status`, {
      baseURL: config.public.apiBase,
      method: 'PATCH',
      headers: {
        Authorization: `Bearer ${authStore.accessToken}`,
      },
      body: { status },
    })

    if (response.success) {
      toast.success('Status Updated', `Contact request has been ${status}`)
      fetchContacts()
    }
  } catch (error: unknown) {
    const err = error as { data?: { message?: string } }
    toast.error('Error', err.data?.message || 'Failed to update status')
  } finally {
    updating.value = null
  }
}

function formatDate(dateString: string): string {
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
  })
}

function formatTime(dateString: string): string {
  return new Date(dateString).toLocaleTimeString('en-US', {
    hour: '2-digit',
    minute: '2-digit',
  })
}

onMounted(() => {
  fetchContacts()
})
</script>
