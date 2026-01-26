<template>
  <div>
    <div class="flex items-center justify-between mb-8">
      <div>
        <h1 class="font-display text-2xl lg:text-3xl font-bold text-neutral-900">
          Users
        </h1>
        <p class="mt-1 text-neutral-600">Manage user accounts and permissions.</p>
      </div>
    </div>

    <!-- Search & Filters -->
    <div class="bg-white rounded-lg border border-neutral-200 p-4 mb-6">
      <div class="flex flex-col sm:flex-row gap-4">
        <div class="flex-1">
          <UInput
            v-model="searchQuery"
            placeholder="Search by email or name..."
            @keyup.enter="fetchUsers"
          />
        </div>
        <select
          v-model="roleFilter"
          class="px-3 py-2 border border-neutral-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500"
          @change="fetchUsers"
        >
          <option value="">All Roles</option>
          <option value="user">User</option>
          <option value="scout">Scout</option>
          <option value="admin">Admin</option>
        </select>
        <UButton variant="outline" @click="fetchUsers">Search</UButton>
      </div>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="flex justify-center py-12">
      <USpinner size="lg" />
    </div>

    <!-- Users Table -->
    <div v-else class="bg-white rounded-lg border border-neutral-200 overflow-hidden">
      <table class="min-w-full divide-y divide-neutral-200">
        <thead class="bg-neutral-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-neutral-500 uppercase tracking-wider">
              User
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-neutral-500 uppercase tracking-wider">
              Role
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-neutral-500 uppercase tracking-wider">
              Subscription
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-neutral-500 uppercase tracking-wider">
              Status
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-neutral-500 uppercase tracking-wider">
              Joined
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-neutral-500 uppercase tracking-wider">
              Actions
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-neutral-200">
          <tr v-for="user in users" :key="user.id" class="hover:bg-neutral-50">
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center">
                <div class="h-10 w-10 flex-shrink-0">
                  <div class="h-10 w-10 rounded-full bg-primary-100 flex items-center justify-center">
                    <span class="text-primary-600 font-medium text-sm">
                      {{ getInitials(user) }}
                    </span>
                  </div>
                </div>
                <div class="ml-4">
                  <div class="text-sm font-medium text-neutral-900">
                    {{ user.first_name }} {{ user.last_name }}
                  </div>
                  <div class="text-sm text-neutral-500">{{ user.email }}</div>
                </div>
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <select
                :value="user.role"
                @change="updateUserRole(user.id, ($event.target as HTMLSelectElement).value)"
                class="text-sm border border-neutral-300 rounded-lg px-2 py-1 focus:outline-none focus:ring-2 focus:ring-primary-500"
              >
                <option value="user">User</option>
                <option value="scout">Scout</option>
                <option value="admin">Admin</option>
              </select>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <UBadge :variant="subscriptionVariant(user.subscription_tier)">
                {{ user.subscription_tier || 'Free' }}
              </UBadge>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <UBadge :variant="user.is_active ? 'success' : 'error'">
                {{ user.is_active ? 'Active' : 'Inactive' }}
              </UBadge>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-neutral-500">
              {{ formatDate(user.created_at) }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
              <div class="flex items-center justify-end space-x-2">
                <UButton
                  size="sm"
                  :variant="user.is_active ? 'ghost' : 'outline'"
                  @click="toggleUserStatus(user)"
                >
                  {{ user.is_active ? 'Deactivate' : 'Activate' }}
                </UButton>
              </div>
            </td>
          </tr>
          <tr v-if="users.length === 0">
            <td colspan="6" class="px-6 py-12 text-center text-neutral-500">
              No users found.
            </td>
          </tr>
        </tbody>
      </table>

      <!-- Pagination -->
      <div v-if="totalPages > 1" class="bg-white px-4 py-3 flex items-center justify-between border-t border-neutral-200">
        <div class="text-sm text-neutral-700">
          Showing {{ (page - 1) * perPage + 1 }} to {{ Math.min(page * perPage, total) }} of {{ total }} users
        </div>
        <div class="flex space-x-2">
          <UButton
            size="sm"
            variant="outline"
            :disabled="page === 1"
            @click="page--; fetchUsers()"
          >
            Previous
          </UButton>
          <UButton
            size="sm"
            variant="outline"
            :disabled="page === totalPages"
            @click="page++; fetchUsers()"
          >
            Next
          </UButton>
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

function subscriptionVariant(tier?: string) {
  switch (tier?.toLowerCase()) {
    case 'pro':
      return 'default'
    case 'scout':
      return 'success'
    case 'club':
      return 'warning'
    default:
      return 'secondary'
  }
}

function formatDate(dateString: string): string {
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
  })
}

onMounted(() => {
  fetchUsers()
})
</script>
