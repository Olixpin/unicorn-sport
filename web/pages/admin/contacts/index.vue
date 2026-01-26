<template>
  <div>
    <div class="flex items-center justify-between mb-8">
      <div>
        <h1 class="font-display text-2xl lg:text-3xl font-bold text-neutral-900">
          Contact Requests
        </h1>
        <p class="mt-1 text-neutral-600">Manage scout contact requests to players.</p>
      </div>
    </div>

    <!-- Filters -->
    <div class="bg-white rounded-lg border border-neutral-200 p-4 mb-6">
      <div class="flex flex-col sm:flex-row gap-4">
        <div class="flex-1">
          <UInput
            v-model="searchQuery"
            placeholder="Search by player or scout..."
            @keyup.enter="fetchContacts"
          />
        </div>
        <select
          v-model="statusFilter"
          class="px-3 py-2 border border-neutral-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500"
          @change="fetchContacts"
        >
          <option value="">All Status</option>
          <option value="pending">Pending</option>
          <option value="approved">Approved</option>
          <option value="rejected">Rejected</option>
        </select>
        <UButton variant="outline" @click="fetchContacts">Search</UButton>
      </div>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="flex justify-center py-12">
      <USpinner size="lg" />
    </div>

    <!-- Contacts Table -->
    <div v-else class="bg-white rounded-lg border border-neutral-200 overflow-hidden">
      <table class="min-w-full divide-y divide-neutral-200">
        <thead class="bg-neutral-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-neutral-500 uppercase tracking-wider">
              Scout
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-neutral-500 uppercase tracking-wider">
              Player
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-neutral-500 uppercase tracking-wider">
              Message
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-neutral-500 uppercase tracking-wider">
              Status
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-neutral-500 uppercase tracking-wider">
              Date
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-neutral-500 uppercase tracking-wider">
              Actions
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-neutral-200">
          <tr v-for="contact in contacts" :key="contact.id" class="hover:bg-neutral-50">
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm font-medium text-neutral-900">{{ contact.scout_name }}</div>
              <div class="text-sm text-neutral-500">{{ contact.scout_email }}</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm font-medium text-neutral-900">{{ contact.player_name }}</div>
            </td>
            <td class="px-6 py-4">
              <p class="text-sm text-neutral-600 max-w-xs truncate">
                {{ contact.message || 'No message' }}
              </p>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <UBadge :variant="statusVariant(contact.status)">
                {{ contact.status }}
              </UBadge>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-neutral-500">
              {{ formatDate(contact.created_at) }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
              <div v-if="contact.status === 'pending'" class="flex items-center justify-end space-x-2">
                <UButton size="sm" @click="updateStatus(contact.id, 'approved')">
                  Approve
                </UButton>
                <UButton
                  size="sm"
                  variant="outline"
                  class="text-red-600"
                  @click="updateStatus(contact.id, 'rejected')"
                >
                  Reject
                </UButton>
              </div>
              <span v-else class="text-neutral-400">—</span>
            </td>
          </tr>
          <tr v-if="contacts.length === 0">
            <td colspan="6" class="px-6 py-12 text-center text-neutral-500">
              No contact requests found.
            </td>
          </tr>
        </tbody>
      </table>

      <!-- Pagination -->
      <div v-if="totalPages > 1" class="bg-white px-4 py-3 flex items-center justify-between border-t border-neutral-200">
        <div class="text-sm text-neutral-700">
          Showing {{ (page - 1) * perPage + 1 }} to {{ Math.min(page * perPage, total) }} of {{ total }} requests
        </div>
        <div class="flex space-x-2">
          <UButton
            size="sm"
            variant="outline"
            :disabled="page === 1"
            @click="page--; fetchContacts()"
          >
            Previous
          </UButton>
          <UButton
            size="sm"
            variant="outline"
            :disabled="page === totalPages"
            @click="page++; fetchContacts()"
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
const searchQuery = ref('')
const statusFilter = ref('')
const page = ref(1)
const perPage = ref(20)
const total = ref(0)
const totalPages = computed(() => Math.ceil(total.value / perPage.value))

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
  }
}

function statusVariant(status: string) {
  switch (status) {
    case 'approved':
      return 'success'
    case 'rejected':
      return 'error'
    default:
      return 'warning'
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
  fetchContacts()
})
</script>
