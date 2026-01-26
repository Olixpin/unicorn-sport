<template>
  <div class="max-w-7xl mx-auto">
    <!-- Page Header -->
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4 mb-8">
      <div>
        <h1 class="font-display text-2xl lg:text-3xl font-bold text-neutral-900">
          Videos
        </h1>
        <p class="mt-1 text-neutral-600">Manage player highlight videos and reels for review.</p>
      </div>
      <div class="flex items-center gap-3">
        <button
          @click="fetchVideos"
          class="inline-flex items-center gap-2 px-4 py-2.5 border border-neutral-300 rounded-xl text-sm font-medium text-neutral-700 bg-white hover:bg-neutral-50 transition-colors"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
          Refresh
        </button>
      </div>
    </div>

    <!-- Stats Cards -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 mb-6">
      <!-- Total Videos -->
      <div class="bg-gradient-to-br from-rose-500 to-pink-600 rounded-2xl p-5 text-white shadow-lg shadow-rose-500/25">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-rose-100 text-sm font-medium">Total Videos</p>
            <p class="text-3xl font-bold mt-1">{{ total }}</p>
          </div>
          <div class="w-12 h-12 bg-white/20 rounded-xl flex items-center justify-center">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
            </svg>
          </div>
        </div>
      </div>

      <!-- Pending Review -->
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

      <!-- Approved -->
      <div class="bg-gradient-to-br from-emerald-500 to-green-600 rounded-2xl p-5 text-white shadow-lg shadow-emerald-500/25">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-emerald-100 text-sm font-medium">Approved</p>
            <p class="text-3xl font-bold mt-1">{{ approvedCount }}</p>
          </div>
          <div class="w-12 h-12 bg-white/20 rounded-xl flex items-center justify-center">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>

      <!-- Rejected -->
      <div class="bg-gradient-to-br from-slate-500 to-gray-600 rounded-2xl p-5 text-white shadow-lg shadow-slate-500/25">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-slate-200 text-sm font-medium">Rejected</p>
            <p class="text-3xl font-bold mt-1">{{ rejectedCount }}</p>
          </div>
          <div class="w-12 h-12 bg-white/20 rounded-xl flex items-center justify-center">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728A9 9 0 015.636 5.636m12.728 12.728L5.636 5.636" />
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
            placeholder="Search by player name or video title..."
            class="w-full pl-12 pr-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
            @keyup.enter="fetchVideos"
          />
        </div>

        <!-- Status Filter -->
        <div class="flex flex-wrap items-center gap-2">
          <button
            v-for="status in statusOptions"
            :key="status.value"
            @click="statusFilter = status.value; fetchVideos()"
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
          @click="fetchVideos"
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
      <p class="mt-4 text-neutral-600">Loading videos...</p>
    </div>

    <!-- Videos Grid -->
    <div v-else>
      <!-- Empty State -->
      <div v-if="videos.length === 0" class="bg-white rounded-2xl border border-neutral-200 px-6 py-16 text-center">
        <div class="w-16 h-16 bg-neutral-100 rounded-2xl flex items-center justify-center mx-auto mb-4">
          <svg class="w-8 h-8 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
          </svg>
        </div>
        <h3 class="font-semibold text-neutral-900 mb-1">No videos found</h3>
        <p class="text-neutral-500 text-sm mb-6">
          {{ hasActiveFilters ? 'Try adjusting your filters to see more results.' : 'Videos uploaded by players will appear here for review.' }}
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

      <!-- Videos -->
      <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div
          v-for="video in videos"
          :key="video.id"
          class="bg-white rounded-2xl border border-neutral-200 overflow-hidden hover:shadow-lg hover:border-primary-200 transition-all group"
        >
          <!-- Video Thumbnail -->
          <div class="relative aspect-video bg-neutral-900">
            <img
              v-if="video.thumbnail_url"
              :src="video.thumbnail_url"
              :alt="video.title"
              class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
            />
            <div v-else class="w-full h-full flex items-center justify-center">
              <svg class="w-16 h-16 text-neutral-700" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.752 11.168l-3.197-2.132A1 1 0 0010 9.87v4.263a1 1 0 001.555.832l3.197-2.132a1 1 0 000-1.664z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
            
            <!-- Play Button Overlay -->
            <a
              :href="video.video_url"
              target="_blank"
              class="absolute inset-0 flex items-center justify-center bg-black/30 opacity-0 group-hover:opacity-100 transition-opacity"
            >
              <div class="w-16 h-16 bg-white/90 rounded-full flex items-center justify-center shadow-lg">
                <svg class="w-8 h-8 text-primary-600 ml-1" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M8 5v14l11-7z" />
                </svg>
              </div>
            </a>
            
            <!-- Duration Badge -->
            <div v-if="video.duration" class="absolute bottom-3 right-3 bg-black/80 text-white text-xs px-2 py-1 rounded-lg font-medium">
              {{ formatDuration(video.duration) }}
            </div>
            
            <!-- Status Badge -->
            <div class="absolute top-3 left-3">
              <span :class="getStatusBadgeClass(video.status)">
                <span class="w-1.5 h-1.5 rounded-full" :class="getStatusDotClass(video.status)"></span>
                {{ capitalizeFirst(video.status) }}
              </span>
            </div>
          </div>

          <!-- Video Info -->
          <div class="p-5">
            <h3 class="font-semibold text-neutral-900 truncate group-hover:text-primary-600 transition-colors">
              {{ video.title }}
            </h3>
            
            <div class="flex items-center gap-2 mt-2">
              <div class="w-6 h-6 bg-gradient-to-br from-primary-500 to-emerald-600 rounded-full flex items-center justify-center text-white text-xs font-semibold">
                {{ getInitials(video.player_name) }}
              </div>
              <span class="text-sm text-neutral-600">{{ video.player_name }}</span>
            </div>
            
            <p class="text-xs text-neutral-400 mt-3">Uploaded {{ formatDate(video.created_at) }}</p>

            <!-- Actions -->
            <div class="flex items-center justify-between mt-4 pt-4 border-t border-neutral-100">
              <a
                :href="video.video_url"
                target="_blank"
                class="inline-flex items-center gap-1.5 text-sm text-primary-600 hover:text-primary-700 font-medium transition-colors"
              >
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
                </svg>
                Watch
              </a>
              
              <div v-if="video.status === 'pending'" class="flex items-center gap-2">
                <button
                  @click="approveVideo(video.id)"
                  :disabled="processing === video.id"
                  class="inline-flex items-center gap-1.5 px-3 py-1.5 bg-emerald-50 text-emerald-700 rounded-lg text-sm font-medium hover:bg-emerald-100 transition-colors disabled:opacity-50"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                  </svg>
                  Approve
                </button>
                <button
                  @click="rejectVideo(video.id)"
                  :disabled="processing === video.id"
                  class="inline-flex items-center gap-1.5 px-3 py-1.5 bg-rose-50 text-rose-700 rounded-lg text-sm font-medium hover:bg-rose-100 transition-colors disabled:opacity-50"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                  </svg>
                  Reject
                </button>
              </div>
              
              <span v-else-if="video.status === 'approved'" class="inline-flex items-center gap-1.5 text-emerald-600 text-sm">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                Approved
              </span>
              
              <span v-else class="inline-flex items-center gap-1.5 text-neutral-500 text-sm">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728A9 9 0 015.636 5.636m12.728 12.728L5.636 5.636" />
                </svg>
                Rejected
              </span>
            </div>
          </div>
        </div>
      </div>

      <!-- Pagination -->
      <div v-if="totalPages > 1" class="mt-8 bg-white rounded-2xl border border-neutral-200 px-6 py-4 shadow-sm">
        <div class="flex flex-col sm:flex-row items-center justify-between gap-4">
          <p class="text-sm text-neutral-600">
            Showing <span class="font-medium">{{ (page - 1) * perPage + 1 }}</span> to 
            <span class="font-medium">{{ Math.min(page * perPage, total) }}</span> of 
            <span class="font-medium">{{ total }}</span> videos
          </p>
          <div class="flex items-center gap-2">
            <button
              :disabled="page === 1"
              @click="page--; fetchVideos()"
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
                @click="page = pageNum; fetchVideos()"
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
              @click="page++; fetchVideos()"
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
const processing = ref<string | null>(null)
const searchQuery = ref('')
const statusFilter = ref('')
const page = ref(1)
const perPage = ref(12)
const total = ref(0)
const totalPages = computed(() => Math.ceil(total.value / perPage.value))

const pendingCount = computed(() => videos.value.filter(v => v.status === 'pending').length)
const approvedCount = computed(() => videos.value.filter(v => v.status === 'approved').length)
const rejectedCount = computed(() => videos.value.filter(v => v.status === 'rejected').length)

const hasActiveFilters = computed(() => searchQuery.value || statusFilter.value)

const statusOptions = [
  { value: '', label: 'All', activeClass: 'bg-neutral-900 text-white' },
  { value: 'pending', label: 'Pending', activeClass: 'bg-amber-500 text-white' },
  { value: 'approved', label: 'Approved', activeClass: 'bg-emerald-500 text-white' },
  { value: 'rejected', label: 'Rejected', activeClass: 'bg-neutral-500 text-white' },
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
  fetchVideos()
}

function getInitials(name?: string): string {
  if (!name) return '?'
  return name.split(' ').map(n => n[0]).join('').toUpperCase().slice(0, 2)
}

function getStatusBadgeClass(status: string): string {
  const baseClass = 'inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full text-xs font-medium backdrop-blur-sm'
  switch (status) {
    case 'approved':
      return `${baseClass} bg-emerald-500/90 text-white`
    case 'rejected':
      return `${baseClass} bg-neutral-500/90 text-white`
    default:
      return `${baseClass} bg-amber-500/90 text-white`
  }
}

function getStatusDotClass(status: string): string {
  switch (status) {
    case 'approved':
      return 'bg-white'
    case 'rejected':
      return 'bg-white'
    default:
      return 'bg-white'
  }
}

function capitalizeFirst(str: string): string {
  return str.charAt(0).toUpperCase() + str.slice(1)
}

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
  processing.value = videoId
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
  } finally {
    processing.value = null
  }
}

async function rejectVideo(videoId: string) {
  processing.value = videoId
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
  } finally {
    processing.value = null
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
