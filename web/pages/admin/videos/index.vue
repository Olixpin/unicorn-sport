<template>
  <div>
    <div class="flex items-center justify-between mb-8">
      <div>
        <h1 class="font-display text-2xl lg:text-3xl font-bold text-neutral-900">
          Videos
        </h1>
        <p class="mt-1 text-neutral-600">Manage player highlight videos and reels.</p>
      </div>
    </div>

    <!-- Filters -->
    <div class="bg-white rounded-lg border border-neutral-200 p-4 mb-6">
      <div class="flex flex-col sm:flex-row gap-4">
        <div class="flex-1">
          <UInput
            v-model="searchQuery"
            placeholder="Search by player name..."
            @keyup.enter="fetchVideos"
          />
        </div>
        <select
          v-model="statusFilter"
          class="px-3 py-2 border border-neutral-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500"
          @change="fetchVideos"
        >
          <option value="">All Status</option>
          <option value="pending">Pending Review</option>
          <option value="approved">Approved</option>
          <option value="rejected">Rejected</option>
        </select>
        <UButton variant="outline" @click="fetchVideos">
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

    <!-- Videos Grid -->
    <div v-else>
      <div v-if="videos.length === 0" class="text-center py-12 bg-white rounded-lg border border-neutral-200">
        <svg class="mx-auto h-12 w-12 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
        </svg>
        <h3 class="mt-4 text-sm font-medium text-neutral-900">No videos found</h3>
        <p class="mt-1 text-sm text-neutral-500">Videos uploaded by players will appear here.</p>
      </div>

      <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div
          v-for="video in videos"
          :key="video.id"
          class="bg-white rounded-lg border border-neutral-200 overflow-hidden"
        >
          <!-- Video Thumbnail -->
          <div class="relative aspect-video bg-neutral-100">
            <img
              v-if="video.thumbnail_url"
              :src="video.thumbnail_url"
              :alt="video.title"
              class="w-full h-full object-cover"
            />
            <div v-else class="w-full h-full flex items-center justify-center">
              <svg class="w-12 h-12 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.752 11.168l-3.197-2.132A1 1 0 0010 9.87v4.263a1 1 0 001.555.832l3.197-2.132a1 1 0 000-1.664z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
            <!-- Duration Badge -->
            <div v-if="video.duration" class="absolute bottom-2 right-2 bg-black/70 text-white text-xs px-2 py-1 rounded">
              {{ formatDuration(video.duration) }}
            </div>
            <!-- Status Badge -->
            <div class="absolute top-2 left-2">
              <UBadge :variant="statusVariant(video.status)">
                {{ video.status }}
              </UBadge>
            </div>
          </div>

          <!-- Video Info -->
          <div class="p-4">
            <h3 class="font-medium text-neutral-900 truncate">{{ video.title }}</h3>
            <p class="text-sm text-neutral-500 mt-1">{{ video.player_name }}</p>
            <p class="text-xs text-neutral-400 mt-2">{{ formatDate(video.created_at) }}</p>

            <!-- Actions -->
            <div class="flex items-center justify-between mt-4 pt-4 border-t border-neutral-100">
              <a
                :href="video.video_url"
                target="_blank"
                class="text-sm text-primary-600 hover:text-primary-700"
              >
                Watch Video
              </a>
              <div v-if="video.status === 'pending'" class="flex space-x-2">
                <UButton size="sm" @click="approveVideo(video.id)">
                  Approve
                </UButton>
                <UButton size="sm" variant="outline" class="text-red-600" @click="rejectVideo(video.id)">
                  Reject
                </UButton>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Pagination -->
      <div v-if="totalPages > 1" class="mt-6 flex items-center justify-center space-x-2">
        <UButton
          size="sm"
          variant="outline"
          :disabled="page === 1"
          @click="page--; fetchVideos()"
        >
          Previous
        </UButton>
        <span class="text-sm text-neutral-600">Page {{ page }} of {{ totalPages }}</span>
        <UButton
          size="sm"
          variant="outline"
          :disabled="page === totalPages"
          @click="page++; fetchVideos()"
        >
          Next
        </UButton>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { ApiResponse } from '~/types/index'

interface Video {
  id: string
  title: string
  player_id: string
  player_name?: string
  video_url: string
  thumbnail_url?: string
  duration?: number
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

const videos = ref<Video[]>([])
const loading = ref(true)
const searchQuery = ref('')
const statusFilter = ref('')
const page = ref(1)
const perPage = ref(12)
const total = ref(0)
const totalPages = computed(() => Math.ceil(total.value / perPage.value))

async function fetchVideos() {
  loading.value = true

  try {
    const params = new URLSearchParams()
    params.append('page', page.value.toString())
    params.append('per_page', perPage.value.toString())
    if (searchQuery.value) params.append('q', searchQuery.value)
    if (statusFilter.value) params.append('status', statusFilter.value)

    const response = await $fetch<ApiResponse<{ videos: Video[]; total: number }>>(
      `/admin/videos?${params.toString()}`,
      {
        baseURL: config.public.apiBase,
        headers: {
          Authorization: `Bearer ${authStore.accessToken}`,
        },
      }
    )

    if (response.success && response.data) {
      videos.value = response.data.videos || []
      total.value = response.data.total || 0
    }
  } catch (error) {
    console.error('Failed to fetch videos:', error)
  } finally {
    loading.value = false
  }
}

async function approveVideo(videoId: string) {
  try {
    const response = await $fetch<ApiResponse<null>>(`/admin/videos/${videoId}/approve`, {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: {
        Authorization: `Bearer ${authStore.accessToken}`,
      },
    })

    if (response.success) {
      toast.success('Video Approved', 'The video has been approved and is now visible')
      fetchVideos()
    }
  } catch (error: unknown) {
    const err = error as { data?: { message?: string } }
    toast.error('Error', err.data?.message || 'Failed to approve video')
  }
}

async function rejectVideo(videoId: string) {
  try {
    const response = await $fetch<ApiResponse<null>>(`/admin/videos/${videoId}/reject`, {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: {
        Authorization: `Bearer ${authStore.accessToken}`,
      },
    })

    if (response.success) {
      toast.success('Video Rejected', 'The video has been rejected')
      fetchVideos()
    }
  } catch (error: unknown) {
    const err = error as { data?: { message?: string } }
    toast.error('Error', err.data?.message || 'Failed to reject video')
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

function formatDuration(seconds: number): string {
  const mins = Math.floor(seconds / 60)
  const secs = seconds % 60
  return `${mins}:${secs.toString().padStart(2, '0')}`
}

function formatDate(dateString: string): string {
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
  })
}

onMounted(() => {
  fetchVideos()
})
</script>
