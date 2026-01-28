<template>
  <div class="max-w-7xl mx-auto">
    <!-- Page Header -->
    <div class="mb-8">
      <div class="flex flex-col lg:flex-row lg:items-center lg:justify-between gap-4">
        <div>
          <h1 class="font-display text-2xl lg:text-3xl font-bold text-neutral-900">
            Users
          </h1>
          <p class="mt-1 text-neutral-500">Manage user accounts, roles and permissions</p>
        </div>
        <div class="flex items-center gap-3">
          <!-- Export Button -->
          <button
            @click="exportUsers"
            class="flex items-center gap-2 px-4 py-2 text-sm font-medium text-neutral-700 bg-white border border-neutral-200 rounded-xl hover:bg-neutral-50 transition-colors"
          >
            <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
            </svg>
            Export CSV
          </button>
          <span class="text-sm text-neutral-500">{{ total }} total users</span>
        </div>
      </div>
    </div>

    <!-- Stats Summary -->
    <div class="grid grid-cols-2 sm:grid-cols-4 gap-3 sm:gap-4 mb-6">
      <div class="bg-white rounded-xl border border-neutral-200 p-3 sm:p-4 flex items-center gap-3 sm:gap-4">
        <div class="w-9 h-9 sm:w-10 sm:h-10 bg-primary-100 rounded-lg flex items-center justify-center flex-shrink-0">
          <svg class="w-4 h-4 sm:w-5 sm:h-5 text-primary-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197" />
          </svg>
        </div>
        <div class="min-w-0">
          <p class="text-xl sm:text-2xl font-bold text-neutral-900">{{ total }}</p>
          <p class="text-xs sm:text-sm text-neutral-500">Total Users</p>
        </div>
      </div>

      <div class="bg-white rounded-xl border border-neutral-200 p-3 sm:p-4 flex items-center gap-3 sm:gap-4">
        <div class="w-9 h-9 sm:w-10 sm:h-10 bg-blue-100 rounded-lg flex items-center justify-center flex-shrink-0">
          <svg class="w-4 h-4 sm:w-5 sm:h-5 text-blue-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </div>
        <div class="min-w-0">
          <p class="text-xl sm:text-2xl font-bold text-neutral-900">{{ scoutCount }}</p>
          <p class="text-xs sm:text-sm text-neutral-500">Scouts</p>
        </div>
      </div>

      <div class="bg-white rounded-xl border border-neutral-200 p-3 sm:p-4 flex items-center gap-3 sm:gap-4">
        <div class="w-9 h-9 sm:w-10 sm:h-10 bg-red-100 rounded-lg flex items-center justify-center flex-shrink-0">
          <svg class="w-4 h-4 sm:w-5 sm:h-5 text-red-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
          </svg>
        </div>
        <div class="min-w-0">
          <p class="text-xl sm:text-2xl font-bold text-neutral-900">{{ adminCount }}</p>
          <p class="text-xs sm:text-sm text-neutral-500">Admins</p>
        </div>
      </div>

      <div class="bg-white rounded-xl border border-neutral-200 p-3 sm:p-4 flex items-center gap-3 sm:gap-4">
        <div class="w-9 h-9 sm:w-10 sm:h-10 bg-emerald-100 rounded-lg flex items-center justify-center flex-shrink-0">
          <svg class="w-4 h-4 sm:w-5 sm:h-5 text-emerald-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
        </div>
        <div class="min-w-0">
          <p class="text-xl sm:text-2xl font-bold text-neutral-900">{{ activeCount }}</p>
          <p class="text-xs sm:text-sm text-neutral-500">Active</p>
        </div>
      </div>
    </div>

    <!-- Search & Filters -->
    <div class="bg-white rounded-2xl border border-neutral-200 p-4 mb-6">
      <div class="flex flex-col lg:flex-row gap-4">
        <!-- Search -->
        <div class="flex-1 relative">
          <svg class="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
          <input
            v-model="searchQuery"
            type="text"
            placeholder="Search by email or name..."
            class="w-full pl-10 pr-4 py-2.5 bg-neutral-50 border-0 rounded-xl text-sm placeholder-neutral-500 focus:outline-none focus:ring-2 focus:ring-primary-500/20 focus:bg-white transition-all"
            @keyup.enter="fetchUsers"
          />
        </div>

        <!-- Role Filter -->
        <div class="relative">
          <select
            v-model="roleFilter"
            class="appearance-none w-full lg:w-40 px-4 py-2.5 bg-neutral-50 border-0 rounded-xl text-sm text-neutral-700 focus:outline-none focus:ring-2 focus:ring-primary-500/20 cursor-pointer"
            @change="fetchUsers"
          >
            <option value="">All Roles</option>
            <option value="user">User</option>
            <option value="scout">Scout</option>
            <option value="admin">Admin</option>
          </select>
          <svg class="absolute right-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400 pointer-events-none" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
          </svg>
        </div>

        <!-- Search Button -->
        <button 
          @click="fetchUsers"
          class="flex items-center gap-2 px-5 py-2.5 bg-neutral-900 text-white font-medium rounded-xl hover:bg-neutral-800 transition-colors"
        >
          <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
          Search
        </button>
      </div>
    </div>

    <!-- Users Table -->
    <div class="bg-white rounded-2xl border border-neutral-200 overflow-hidden">
      <!-- Loading State -->
      <div v-if="loading" class="flex flex-col items-center justify-center py-16">
        <div class="w-12 h-12 border-4 border-primary-500/30 border-t-primary-500 rounded-full animate-spin mb-4"></div>
        <p class="text-neutral-500">Loading users...</p>
      </div>

      <!-- Empty State -->
      <div v-else-if="users.length === 0" class="flex flex-col items-center justify-center py-16">
        <div class="w-20 h-20 bg-neutral-100 rounded-full flex items-center justify-center mb-6">
          <svg class="w-10 h-10 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197" />
          </svg>
        </div>
        <h3 class="text-lg font-semibold text-neutral-900 mb-2">No users found</h3>
        <p class="text-neutral-500 text-center max-w-sm">
          Try adjusting your search filters to find what you're looking for.
        </p>
      </div>

      <!-- Table -->
      <div v-else class="overflow-x-auto">
        <table class="w-full">
          <thead>
            <tr class="border-b border-neutral-200 bg-neutral-50">
              <th class="px-6 py-4 text-left text-xs font-semibold text-neutral-500 uppercase tracking-wider">User</th>
              <th class="px-6 py-4 text-left text-xs font-semibold text-neutral-500 uppercase tracking-wider">Role</th>
              <th class="px-6 py-4 text-left text-xs font-semibold text-neutral-500 uppercase tracking-wider hidden md:table-cell">Subscription</th>
              <th class="px-6 py-4 text-left text-xs font-semibold text-neutral-500 uppercase tracking-wider">Status</th>
              <th class="px-6 py-4 text-left text-xs font-semibold text-neutral-500 uppercase tracking-wider hidden lg:table-cell">Joined</th>
              <th class="px-6 py-4 text-right text-xs font-semibold text-neutral-500 uppercase tracking-wider">Actions</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-neutral-100">
            <tr 
              v-for="user in users" 
              :key="user.id" 
              class="group hover:bg-neutral-50 transition-colors"
            >
              <td class="px-4 sm:px-6 py-4">
                <div class="flex items-center gap-3 sm:gap-4">
                  <div 
                    :class="[
                      'w-10 h-10 sm:w-11 sm:h-11 rounded-xl flex items-center justify-center font-semibold text-xs sm:text-sm flex-shrink-0',
                      getRoleColorClass(user.role)
                    ]"
                  >
                    {{ getInitials(user) }}
                  </div>
                  <div class="min-w-0">
                    <p class="font-medium text-neutral-900 truncate">{{ user.first_name || '' }} {{ user.last_name || '' }}</p>
                    <p class="text-xs sm:text-sm text-neutral-500 truncate max-w-[150px] sm:max-w-none">{{ user.email }}</p>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4">
                <div class="relative inline-block">
                  <select
                    :value="user.role"
                    @change="updateUserRole(user.id, ($event.target as HTMLSelectElement).value)"
                    :class="[
                      'appearance-none px-3 py-1.5 pr-8 text-xs font-medium rounded-lg border cursor-pointer focus:outline-none focus:ring-2 focus:ring-offset-1',
                      getRoleBadgeClass(user.role)
                    ]"
                  >
                    <option value="user">User</option>
                    <option value="scout">Scout</option>
                    <option value="admin">Admin</option>
                  </select>
                  <svg class="absolute right-2 top-1/2 -translate-y-1/2 w-3 h-3 text-current pointer-events-none" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                  </svg>
                </div>
              </td>
              <td class="px-6 py-4 hidden md:table-cell">
                <span 
                  :class="[
                    'inline-flex items-center px-2.5 py-1 rounded-full text-xs font-medium',
                    getSubscriptionClass(user.subscription_tier)
                  ]"
                >
                  {{ user.subscription_tier || 'Free' }}
                </span>
              </td>
              <td class="px-6 py-4">
                <span 
                  :class="[
                    'inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full text-xs font-medium',
                    user.is_active 
                      ? 'bg-emerald-100 text-emerald-700' 
                      : 'bg-neutral-100 text-neutral-600'
                  ]"
                >
                  <span :class="['w-1.5 h-1.5 rounded-full', user.is_active ? 'bg-emerald-500' : 'bg-neutral-400']"></span>
                  {{ user.is_active ? 'Active' : 'Inactive' }}
                </span>
              </td>
              <td class="px-6 py-4 hidden lg:table-cell">
                <span class="text-sm text-neutral-600">{{ formatDate(user.created_at) }}</span>
              </td>
              <td class="px-6 py-4">
                <div class="flex items-center justify-end gap-2">
                  <button 
                    @click="toggleUserStatus(user)"
                    :class="[
                      'px-3 py-1.5 text-xs font-medium rounded-lg transition-colors',
                      user.is_active 
                        ? 'text-red-600 hover:bg-red-50' 
                        : 'text-emerald-600 hover:bg-emerald-50'
                    ]"
                  >
                    {{ user.is_active ? 'Deactivate' : 'Activate' }}
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div v-if="users.length > 0" class="flex flex-col sm:flex-row items-center justify-between gap-4 p-6 border-t border-neutral-100">
        <p class="text-sm text-neutral-500">
          Showing <span class="font-medium text-neutral-900">{{ (page - 1) * perPage + 1 }}</span> to 
          <span class="font-medium text-neutral-900">{{ Math.min(page * perPage, total) }}</span> of 
          <span class="font-medium text-neutral-900">{{ total }}</span> users
        </p>
        <div class="flex items-center gap-2">
          <button
            :disabled="page === 1"
            @click="page--; fetchUsers()"
            :class="[
              'flex items-center gap-1 px-4 py-2 text-sm font-medium rounded-lg transition-colors',
              page === 1 
                ? 'text-neutral-300 cursor-not-allowed' 
                : 'text-neutral-600 hover:bg-neutral-100'
            ]"
          >
            <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
            </svg>
            Previous
          </button>
          <button
            :disabled="page >= totalPages"
            @click="page++; fetchUsers()"
            :class="[
              'flex items-center gap-1 px-4 py-2 text-sm font-medium rounded-lg transition-colors',
              page >= totalPages 
                ? 'text-neutral-300 cursor-not-allowed' 
                : 'text-neutral-600 hover:bg-neutral-100'
            ]"
          >
            Next
            <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
            </svg>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { ApiResponse } from '~/types/index'

interface User {
  id: string
  email: string
  first_name?: string
  last_name?: string
  role: string
  subscription_tier?: string
  is_active: boolean
  created_at: string
}

definePageMeta({
  layout: 'admin',
  middleware: 'admin',
})

const config = useRuntimeConfig()
const authStore = useAuthStore()
const toast = useToast()

const users = ref<User[]>([])
const loading = ref(true)
const searchQuery = ref('')
const roleFilter = ref('')
const page = ref(1)
const perPage = ref(20)
const total = ref(0)
const totalPages = computed(() => Math.ceil(total.value / perPage.value))

// Computed stats
const scoutCount = computed(() => users.value.filter(u => u.role === 'scout').length)
const adminCount = computed(() => users.value.filter(u => u.role === 'admin').length)
const activeCount = computed(() => users.value.filter(u => u.is_active).length)

function getRoleColorClass(role: string): string {
  switch (role) {
    case 'admin':
      return 'bg-red-100 text-red-700'
    case 'scout':
      return 'bg-blue-100 text-blue-700'
    default:
      return 'bg-neutral-100 text-neutral-700'
  }
}

function getRoleBadgeClass(role: string): string {
  switch (role) {
    case 'admin':
      return 'bg-red-50 border-red-200 text-red-700 focus:ring-red-500/20'
    case 'scout':
      return 'bg-blue-50 border-blue-200 text-blue-700 focus:ring-blue-500/20'
    default:
      return 'bg-neutral-50 border-neutral-200 text-neutral-700 focus:ring-neutral-500/20'
  }
}

function getSubscriptionClass(tier?: string): string {
  switch (tier?.toLowerCase()) {
    case 'pro':
      return 'bg-primary-100 text-primary-700'
    case 'scout':
      return 'bg-amber-100 text-amber-700'
    case 'club':
      return 'bg-purple-100 text-purple-700'
    default:
      return 'bg-neutral-100 text-neutral-600'
  }
}

async function fetchUsers() {
  loading.value = true

  try {
    const params = new URLSearchParams()
    params.append('page', page.value.toString())
    params.append('per_page', perPage.value.toString())
    if (searchQuery.value) params.append('q', searchQuery.value)
    if (roleFilter.value) params.append('role', roleFilter.value)

    const response = await $fetch<ApiResponse<{ users: User[]; total: number }>>(
      `/admin/users?${params.toString()}`,
      {
        baseURL: config.public.apiBase,
        headers: {
          Authorization: `Bearer ${authStore.accessToken}`,
        },
      }
    )

    if (response.success && response.data) {
      users.value = response.data.users || []
      total.value = response.data.total || 0
    }
  } catch (error) {
    console.error('Failed to fetch users:', error)
  } finally {
    loading.value = false
  }
}

async function updateUserRole(userId: string, newRole: string) {
  try {
    const response = await $fetch<ApiResponse<null>>(`/admin/users/${userId}/role`, {
      baseURL: config.public.apiBase,
      method: 'PATCH',
      headers: {
        Authorization: `Bearer ${authStore.accessToken}`,
      },
      body: { role: newRole },
    })

    if (response.success) {
      toast.success('Role Updated', 'User role has been updated')
    }
  } catch (error: unknown) {
    const err = error as { data?: { message?: string } }
    toast.error('Error', err.data?.message || 'Failed to update role')
    fetchUsers() // Refresh to revert
  }
}

async function toggleUserStatus(user: User) {
  try {
    const endpoint = user.is_active ? 'deactivate' : 'activate'
    const response = await $fetch<ApiResponse<null>>(`/admin/users/${user.id}/${endpoint}`, {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: {
        Authorization: `Bearer ${authStore.accessToken}`,
      },
    })

    if (response.success) {
      toast.success('Status Updated', `User has been ${user.is_active ? 'deactivated' : 'activated'}`)
      fetchUsers()
    }
  } catch (error: unknown) {
    const err = error as { data?: { message?: string } }
    toast.error('Error', err.data?.message || 'Failed to update status')
  }
}

function getInitials(user: User): string {
  const first = user.first_name?.charAt(0) || ''
  const last = user.last_name?.charAt(0) || ''
  return (first + last).toUpperCase() || user.email.charAt(0).toUpperCase()
}

function formatDate(dateString: string): string {
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
  })
}

async function exportUsers() {
  try {
    const response = await fetch(`${config.public.apiUrl}/admin/export/users`, {
      headers: {
        'Authorization': `Bearer ${authStore.accessToken}`
      }
    })
    const blob = await response.blob()
    const url = window.URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = `users-export-${new Date().toISOString().split('T')[0]}.csv`
    document.body.appendChild(a)
    a.click()
    document.body.removeChild(a)
    window.URL.revokeObjectURL(url)
  } catch (error) {
    console.error('Export failed:', error)
  }
}

onMounted(() => {
  fetchUsers()
})
</script>
