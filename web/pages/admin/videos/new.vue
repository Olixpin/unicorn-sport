<template>
  <div class="max-w-4xl mx-auto">
    <!-- Page Header -->
    <div class="mb-8">
      <div class="flex items-center gap-3 mb-4">
        <NuxtLink
          to="/admin/videos"
          class="p-2 hover:bg-neutral-100 rounded-lg transition-colors"
        >
          <svg class="w-5 h-5 text-neutral-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
          </svg>
        </NuxtLink>
        <div>
          <h1 class="font-display text-2xl lg:text-3xl font-bold text-neutral-900">
            Upload Video
          </h1>
          <p class="mt-1 text-neutral-600">Upload player highlights or full match recordings</p>
        </div>
      </div>
    </div>

    <form @submit.prevent="handleSubmit" class="space-y-8">
      <!-- Video Type Selection -->
      <div class="bg-white rounded-2xl border border-neutral-200 p-6 shadow-sm">
        <h2 class="text-lg font-semibold text-neutral-900 mb-4">Video Type</h2>
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <label
            :class="[
              'relative flex flex-col p-5 rounded-xl border-2 cursor-pointer transition-all',
              form.video_type === 'highlight'
                ? 'border-primary-500 bg-primary-50'
                : 'border-neutral-200 hover:border-neutral-300'
            ]"
          >
            <input
              type="radio"
              v-model="form.video_type"
              value="highlight"
              class="sr-only"
            />
            <div class="flex items-center gap-3 mb-2">
              <div :class="['w-10 h-10 rounded-lg flex items-center justify-center', form.video_type === 'highlight' ? 'bg-primary-500 text-white' : 'bg-neutral-100 text-neutral-600']">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z" />
                </svg>
              </div>
              <div>
                <span class="font-semibold text-neutral-900">Player Highlight</span>
                <span class="ml-2 px-2 py-0.5 text-xs font-medium bg-emerald-100 text-emerald-700 rounded-full">FREE</span>
              </div>
            </div>
            <p class="text-sm text-neutral-600">Short clips showcasing player skills (1-5 min, max 1GB)</p>
          </label>

          <label
            :class="[
              'relative flex flex-col p-5 rounded-xl border-2 cursor-pointer transition-all',
              form.video_type === 'full_match'
                ? 'border-primary-500 bg-primary-50'
                : 'border-neutral-200 hover:border-neutral-300'
            ]"
          >
            <input
              type="radio"
              v-model="form.video_type"
              value="full_match"
              class="sr-only"
            />
            <div class="flex items-center gap-3 mb-2">
              <div :class="['w-10 h-10 rounded-lg flex items-center justify-center', form.video_type === 'full_match' ? 'bg-primary-500 text-white' : 'bg-neutral-100 text-neutral-600']">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
                </svg>
              </div>
              <div>
                <span class="font-semibold text-neutral-900">Full Match</span>
                <span class="ml-2 px-2 py-0.5 text-xs font-medium bg-amber-100 text-amber-700 rounded-full">PAID</span>
              </div>
            </div>
            <p class="text-sm text-neutral-600">Complete game recordings (60-90 min, max 25GB)</p>
          </label>
        </div>
      </div>

      <!-- Video Upload -->
      <div class="bg-white rounded-2xl border border-neutral-200 p-6 shadow-sm">
        <h2 class="text-lg font-semibold text-neutral-900 mb-4">Video File</h2>
        
        <!-- Upload Zone -->
        <div
          v-if="!videoFile && !uploadProgress"
          @drop.prevent="handleDrop"
          @dragover.prevent="isDragging = true"
          @dragleave="isDragging = false"
          :class="[
            'relative border-2 border-dashed rounded-xl p-10 text-center transition-all cursor-pointer',
            isDragging ? 'border-primary-500 bg-primary-50' : 'border-neutral-300 hover:border-neutral-400'
          ]"
          @click="videoInput?.click()"
        >
          <input
            ref="videoInput"
            type="file"
            accept="video/mp4,video/quicktime,video/webm,video/x-matroska"
            class="hidden"
            @change="handleFileSelect"
          />
          <div class="flex flex-col items-center">
            <div class="w-16 h-16 bg-neutral-100 rounded-2xl flex items-center justify-center mb-4">
              <svg class="w-8 h-8 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
              </svg>
            </div>
            <p class="text-neutral-900 font-medium mb-1">Drop your video here or click to browse</p>
            <p class="text-sm text-neutral-500">MP4, MOV, WebM, or MKV • Max {{ form.video_type === 'full_match' ? '25GB' : '1GB' }}</p>
          </div>
        </div>

        <!-- Upload Progress -->
        <div v-else-if="uploadProgress !== null && uploadProgress < 100" class="space-y-4">
          <div class="flex items-center gap-4">
            <div class="w-14 h-14 bg-primary-100 rounded-xl flex items-center justify-center">
              <svg class="w-7 h-7 text-primary-600 animate-pulse" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
              </svg>
            </div>
            <div class="flex-1 min-w-0">
              <p class="font-medium text-neutral-900 truncate">{{ videoFile?.name }}</p>
              <p class="text-sm text-neutral-500">{{ formatFileSize(videoFile?.size || 0) }} • Uploading...</p>
            </div>
            <button
              type="button"
              @click="cancelUpload"
              class="p-2 hover:bg-neutral-100 rounded-lg transition-colors"
            >
              <svg class="w-5 h-5 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
          <div class="relative">
            <div class="h-2 bg-neutral-200 rounded-full overflow-hidden">
              <div
                class="h-full bg-gradient-to-r from-primary-500 to-emerald-500 transition-all duration-300"
                :style="{ width: `${uploadProgress}%` }"
              ></div>
            </div>
            <p class="text-sm text-neutral-600 mt-2">{{ uploadProgress }}% uploaded</p>
          </div>
        </div>

        <!-- File Selected / Upload Complete -->
        <div v-else-if="videoFile" class="flex items-center gap-4 p-4 bg-emerald-50 rounded-xl border border-emerald-200">
          <div class="w-12 h-12 bg-emerald-500 rounded-xl flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
            </svg>
          </div>
          <div class="flex-1 min-w-0">
            <p class="font-medium text-neutral-900 truncate">{{ videoFile.name }}</p>
            <p class="text-sm text-emerald-600">{{ formatFileSize(videoFile.size) }} • Ready to save</p>
          </div>
          <button
            type="button"
            @click="removeVideo"
            class="p-2 hover:bg-emerald-100 rounded-lg transition-colors"
          >
            <svg class="w-5 h-5 text-neutral-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
      </div>

      <!-- Thumbnail Upload -->
      <div class="bg-white rounded-xl border border-neutral-200 p-4 sm:p-6">
        <h2 class="text-sm sm:text-base font-semibold text-neutral-900 mb-3">Thumbnail (Optional)</h2>
        
        <div class="flex flex-col sm:flex-row gap-4">
          <!-- Thumbnail Preview -->
          <div class="w-full sm:w-40 h-24 bg-neutral-100 rounded-lg overflow-hidden flex-shrink-0">
            <img
              v-if="thumbnailPreview"
              :src="thumbnailPreview"
              alt="Thumbnail preview"
              class="w-full h-full object-cover"
            />
            <div v-else class="w-full h-full flex items-center justify-center">
              <svg class="w-8 h-8 text-neutral-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
              </svg>
            </div>
          </div>
          
          <!-- Upload Button & Info -->
          <div class="flex flex-col justify-center">
            <input
              ref="thumbnailInput"
              type="file"
              accept="image/jpeg,image/png,image/webp"
              class="hidden"
              @change="handleThumbnailSelect"
            />
            <button
              type="button"
              @click="thumbnailInput?.click()"
              class="w-full sm:w-auto px-4 py-2 border border-neutral-200 rounded-lg text-sm font-medium text-neutral-700 hover:bg-neutral-50 transition-colors"
            >
              {{ thumbnailPreview ? 'Change' : 'Upload' }}
            </button>
            <p class="mt-2 text-[11px] text-neutral-400">JPG, PNG, WebP • Max 10MB • 16:9</p>
          </div>
        </div>
      </div>

      <!-- Video Details -->
      <div class="bg-white rounded-2xl border border-neutral-200 p-6 shadow-sm">
        <h2 class="text-lg font-semibold text-neutral-900 mb-4">Video Details</h2>
        
        <div class="space-y-5">
          <!-- Title -->
          <div>
            <label class="block text-sm font-medium text-neutral-700 mb-1.5">Title *</label>
            <input
              v-model="form.title"
              type="text"
              required
              placeholder="e.g., Amazing goal vs Academy FC"
              class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
            />
          </div>

          <!-- Description -->
          <div>
            <label class="block text-sm font-medium text-neutral-700 mb-1.5">Description</label>
            <textarea
              v-model="form.description"
              rows="3"
              placeholder="Brief description of the video content..."
              class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all resize-none"
            ></textarea>
          </div>

          <!-- Duration -->
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-neutral-700 mb-1.5">Duration (seconds)</label>
              <input
                v-model.number="form.duration_seconds"
                type="number"
                min="0"
                placeholder="180"
                class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-neutral-700 mb-1.5">Tournament</label>
              <select
                v-model="form.tournament_id"
                class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
              >
                <option value="">Select tournament (optional)</option>
                <option v-for="tournament in tournaments" :key="tournament.id" :value="tournament.id">
                  {{ tournament.name }} ({{ tournament.year }})
                </option>
              </select>
            </div>
          </div>
        </div>
      </div>

      <!-- Player Selection -->
      <div class="bg-white rounded-2xl border border-neutral-200 p-6 shadow-sm">
        <h2 class="text-lg font-semibold text-neutral-900 mb-4">Featured Players</h2>
        <p class="text-sm text-neutral-600 mb-4">Select the players featured in this video. The first player selected will be marked as primary.</p>
        
        <!-- Player Search -->
        <div class="relative mb-4">
          <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
            <svg class="w-5 h-5 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </div>
          <input
            v-model="playerSearch"
            type="text"
            placeholder="Search players by name..."
            class="w-full pl-12 pr-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
            @input="searchPlayers"
          />
        </div>

        <!-- Selected Players -->
        <div v-if="selectedPlayers.length > 0" class="mb-4">
          <p class="text-sm font-medium text-neutral-700 mb-2">Selected ({{ selectedPlayers.length }})</p>
          <div class="flex flex-wrap gap-2">
            <div
              v-for="(player, index) in selectedPlayers"
              :key="player.id"
              class="inline-flex items-center gap-2 px-3 py-1.5 bg-primary-50 border border-primary-200 rounded-full"
            >
              <span v-if="index === 0" class="w-4 h-4 bg-primary-500 rounded-full flex items-center justify-center">
                <svg class="w-2.5 h-2.5 text-white" fill="currentColor" viewBox="0 0 20 20">
                  <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                </svg>
              </span>
              <span class="text-sm font-medium text-primary-700">{{ player.first_name }} {{ player.last_name }}</span>
              <button
                type="button"
                @click="removePlayer(player.id)"
                class="text-primary-500 hover:text-primary-700"
              >
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            </div>
          </div>
        </div>

        <!-- Player Results -->
        <div v-if="playerResults.length > 0" class="border border-neutral-200 rounded-xl overflow-hidden">
          <div
            v-for="player in playerResults"
            :key="player.id"
            @click="togglePlayer(player)"
            :class="[
              'flex items-center gap-3 px-4 py-3 cursor-pointer transition-colors',
              isPlayerSelected(player.id) ? 'bg-primary-50' : 'hover:bg-neutral-50'
            ]"
          >
            <div class="w-10 h-10 bg-gradient-to-br from-primary-500 to-emerald-600 rounded-full flex items-center justify-center text-white font-semibold text-sm">
              {{ player.first_name[0] }}{{ player.last_name[0] }}
            </div>
            <div class="flex-1 min-w-0">
              <p class="font-medium text-neutral-900">{{ player.first_name }} {{ player.last_name }}</p>
              <p class="text-sm text-neutral-500">{{ player.position }} • {{ player.country }}</p>
            </div>
            <div v-if="isPlayerSelected(player.id)" class="w-6 h-6 bg-primary-500 rounded-full flex items-center justify-center">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
              </svg>
            </div>
          </div>
        </div>
        
        <p v-else-if="playerSearch && !loadingPlayers" class="text-center text-neutral-500 py-4">
          No players found matching "{{ playerSearch }}"
        </p>
      </div>

      <!-- Submit Button - Desktop only -->
      <div class="hidden lg:flex items-center justify-end gap-4 pt-4">
        <NuxtLink
          to="/admin/videos"
          class="px-6 py-3 text-neutral-700 font-medium hover:bg-neutral-100 rounded-xl transition-colors"
        >
          Cancel
        </NuxtLink>
        <button
          type="submit"
          :disabled="!isFormValid || submitting"
          :class="[
            'px-8 py-3 font-semibold rounded-xl transition-all',
            isFormValid 
              ? 'bg-gradient-to-r from-primary-600 to-emerald-600 text-white hover:from-primary-700 hover:to-emerald-700 shadow-lg shadow-primary-600/25' 
              : 'bg-neutral-100 text-neutral-400 cursor-not-allowed'
          ]"
        >
          <span v-if="submitting" class="flex items-center gap-2">
            <svg class="w-5 h-5 animate-spin" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            Uploading...
          </span>
          <span v-else>{{ isFormValid ? 'Upload Video' : 'Fill Required Fields' }}</span>
        </button>
      </div>
    </form>

    <!-- Mobile Fixed Action Bar -->
    <div class="fixed bottom-0 left-0 right-0 bg-white/90 backdrop-blur-lg border-t border-neutral-200 p-4 z-50 lg:hidden">
      <div class="flex gap-3 max-w-md mx-auto">
        <NuxtLink to="/admin/videos" class="w-auto">
          <button
            type="button"
            class="px-4 py-3 border border-neutral-300 text-neutral-700 rounded-xl text-sm font-medium hover:bg-neutral-50 transition-colors"
          >
            Cancel
          </button>
        </NuxtLink>
        <button
          type="button"
          @click="handleSubmit"
          :disabled="!isFormValid || submitting"
          :class="[
            'flex-1 inline-flex items-center justify-center gap-2 px-4 py-3 rounded-xl text-sm font-semibold transition-all',
            isFormValid 
              ? 'bg-gradient-to-r from-primary-600 to-emerald-600 text-white shadow-lg shadow-primary-600/25' 
              : 'bg-neutral-100 text-neutral-400 cursor-not-allowed'
          ]"
        >
          <svg v-if="submitting" class="w-4 h-4 animate-spin" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          {{ submitting ? 'Uploading...' : (isFormValid ? 'Upload Video' : 'Fill Required Fields') }}
        </button>
      </div>
    </div>

    <!-- Spacer for mobile fixed bar -->
    <div class="lg:hidden h-24"></div>
  </div>
</template>

<script setup lang="ts">
import type { ApiResponse } from '~/types/index'

interface Player {
  id: string
  first_name: string
  last_name: string
  position: string
  country: string
}

interface Tournament {
  id: string
  name: string
  year: number
}

interface UploadSession {
  session_id: string
  upload_url: string
  upload_method: string
  s3_key: string
  expires_in: number
}

definePageMeta({
  layout: 'admin',
  middleware: 'admin',
})

const config = useRuntimeConfig()
const authStore = useAuthStore()
const router = useRouter()
const toast = useToast()

// Form state
const form = reactive({
  video_type: 'highlight',
  title: '',
  description: '',
  duration_seconds: null as number | null,
  tournament_id: '',
})

// File state
const videoFile = ref<File | null>(null)
const thumbnailFile = ref<File | null>(null)
const thumbnailPreview = ref<string | null>(null)
const uploadProgress = ref<number | null>(null)
const uploadSession = ref<UploadSession | null>(null)
const isDragging = ref(false)

// Template refs
const videoInput = ref<HTMLInputElement | null>(null)
const thumbnailInput = ref<HTMLInputElement | null>(null)

// Player selection
const playerSearch = ref('')
const playerResults = ref<Player[]>([])
const selectedPlayers = ref<Player[]>([])
const loadingPlayers = ref(false)

// Tournaments
const tournaments = ref<Tournament[]>([])

// UI state
const submitting = ref(false)

const isFormValid = computed(() => {
  return form.title && videoFile.value && uploadProgress.value === 100
})

// Fetch tournaments on mount
onMounted(async () => {
  try {
    const response = await $fetch<ApiResponse<{ events: Tournament[] }>>('/admin/events', {
      baseURL: config.public.apiBase,
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
    })
    if (response.success && response.data) {
      tournaments.value = response.data.events || []
    }
  } catch (error) {
    console.error('Failed to fetch tournaments:', error)
  }
})

// File handling
function handleDrop(e: DragEvent) {
  isDragging.value = false
  const files = e.dataTransfer?.files
  if (files && files.length > 0 && files[0]) {
    handleVideoFile(files[0])
  }
}

function handleFileSelect(e: Event) {
  const input = e.target as HTMLInputElement
  if (input.files && input.files.length > 0 && input.files[0]) {
    handleVideoFile(input.files[0])
  }
}

async function handleVideoFile(file: File) {
  // Validate file type
  const allowedTypes = ['video/mp4', 'video/quicktime', 'video/webm', 'video/x-matroska']
  if (!allowedTypes.includes(file.type)) {
    toast.error('Invalid File', 'Please upload a valid video file (MP4, MOV, WebM, or MKV)')
    return
  }

  // Validate file size
  const maxSize = form.video_type === 'full_match' ? 25 * 1024 * 1024 * 1024 : 1 * 1024 * 1024 * 1024
  if (file.size > maxSize) {
    toast.error('File Too Large', `Maximum file size is ${form.video_type === 'full_match' ? '25GB' : '1GB'}`)
    return
  }

  videoFile.value = file
  
  // Auto-fill title from filename if empty
  if (!form.title) {
    form.title = file.name.replace(/\.[^/.]+$/, '').replace(/[-_]/g, ' ')
  }

  // Start upload
  await uploadVideoFile(file)
}

async function uploadVideoFile(file: File) {
  try {
    uploadProgress.value = 0

    // 1. Initialize upload session
    const initResponse = await $fetch<ApiResponse<UploadSession>>('/admin/upload/init', {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: {
        upload_type: form.video_type,
        file_name: file.name,
        content_type: file.type,
        file_size: file.size,
      },
    })

    if (!initResponse.success || !initResponse.data) {
      throw new Error('Failed to initialize upload')
    }

    uploadSession.value = initResponse.data

    // 2. Upload file directly to S3
    if (initResponse.data.upload_method === 'direct' && initResponse.data.upload_url) {
      await uploadDirect(file, initResponse.data.upload_url)
    } else {
      // For large files, use multipart upload
      await uploadMultipart(file)
    }

    uploadProgress.value = 100
    toast.success('Upload Complete', 'Video uploaded successfully')
  } catch (error) {
    console.error('Upload failed:', error)
    toast.error('Upload Failed', 'Failed to upload video. Please try again.')
    uploadProgress.value = null
    uploadSession.value = null
  }
}

async function uploadDirect(file: File, uploadUrl: string) {
  return new Promise<void>((resolve, reject) => {
    const xhr = new XMLHttpRequest()
    
    xhr.upload.addEventListener('progress', (e) => {
      if (e.lengthComputable) {
        uploadProgress.value = Math.round((e.loaded / e.total) * 99) // Reserve 1% for finalization
      }
    })

    xhr.addEventListener('load', () => {
      if (xhr.status >= 200 && xhr.status < 300) {
        resolve()
      } else {
        reject(new Error(`Upload failed with status ${xhr.status}`))
      }
    })

    xhr.addEventListener('error', () => reject(new Error('Upload failed')))
    xhr.addEventListener('abort', () => reject(new Error('Upload cancelled')))

    xhr.open('PUT', uploadUrl)
    xhr.setRequestHeader('Content-Type', file.type)
    xhr.send(file)
  })
}

async function uploadMultipart(file: File) {
  // Initialize multipart upload
  const initResponse = await $fetch<ApiResponse<{ upload_id: string; total_parts: number; part_size: number }>>('/admin/upload/multipart/init', {
    baseURL: config.public.apiBase,
    method: 'POST',
    headers: { Authorization: `Bearer ${authStore.accessToken}` },
    body: { session_id: uploadSession.value?.session_id },
  })

  if (!initResponse.success || !initResponse.data) {
    throw new Error('Failed to initialize multipart upload')
  }

  const { total_parts, part_size } = initResponse.data
  const completedParts: { part_number: number; etag: string }[] = []

  // Upload each part
  for (let partNumber = 1; partNumber <= total_parts; partNumber++) {
    const start = (partNumber - 1) * part_size
    const end = Math.min(start + part_size, file.size)
    const chunk = file.slice(start, end)

    // Get presigned URL for this part
    const partUrlResponse = await $fetch<ApiResponse<{ upload_url: string }>>('/admin/upload/multipart/part-url', {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: { session_id: uploadSession.value?.session_id, part_number: partNumber },
    })

    if (!partUrlResponse.success || !partUrlResponse.data) {
      throw new Error(`Failed to get URL for part ${partNumber}`)
    }

    // Upload the chunk
    const response = await fetch(partUrlResponse.data.upload_url, {
      method: 'PUT',
      body: chunk,
    })

    if (!response.ok) {
      throw new Error(`Failed to upload part ${partNumber}`)
    }

    const etag = response.headers.get('ETag') || ''
    completedParts.push({ part_number: partNumber, etag: etag.replace(/"/g, '') })

    // Update progress
    uploadProgress.value = Math.round((partNumber / total_parts) * 99)
  }

  // Complete multipart upload
  await $fetch('/admin/upload/multipart/complete', {
    baseURL: config.public.apiBase,
    method: 'POST',
    headers: { Authorization: `Bearer ${authStore.accessToken}` },
    body: { session_id: uploadSession.value?.session_id, parts: completedParts },
  })
}

function cancelUpload() {
  // TODO: Implement upload cancellation
  videoFile.value = null
  uploadProgress.value = null
  uploadSession.value = null
}

function removeVideo() {
  videoFile.value = null
  uploadProgress.value = null
  uploadSession.value = null
}

// Thumbnail handling
function handleThumbnailSelect(e: Event) {
  const input = e.target as HTMLInputElement
  if (input.files && input.files.length > 0) {
    const file = input.files[0]
    if (!file) return
    
    // Validate
    if (!['image/jpeg', 'image/png', 'image/webp'].includes(file.type)) {
      toast.error('Invalid File', 'Please upload a JPG, PNG, or WebP image')
      return
    }
    
    if (file.size > 10 * 1024 * 1024) {
      toast.error('File Too Large', 'Maximum thumbnail size is 10MB')
      return
    }

    thumbnailFile.value = file
    thumbnailPreview.value = URL.createObjectURL(file)
  }
}

// Player search
let searchTimeout: ReturnType<typeof setTimeout> | null = null

function searchPlayers() {
  if (searchTimeout) clearTimeout(searchTimeout)
  
  searchTimeout = setTimeout(async () => {
    if (!playerSearch.value || playerSearch.value.length < 2) {
      playerResults.value = []
      return
    }

    loadingPlayers.value = true
    try {
      const response = await $fetch<ApiResponse<{ players: Player[] }>>(`/admin/players?q=${encodeURIComponent(playerSearch.value)}&limit=10`, {
        baseURL: config.public.apiBase,
        headers: { Authorization: `Bearer ${authStore.accessToken}` },
      })
      if (response.success && response.data) {
        playerResults.value = response.data.players || []
      }
    } catch (error) {
      console.error('Failed to search players:', error)
    } finally {
      loadingPlayers.value = false
    }
  }, 300)
}

function togglePlayer(player: Player) {
  const index = selectedPlayers.value.findIndex(p => p.id === player.id)
  if (index >= 0) {
    selectedPlayers.value.splice(index, 1)
  } else {
    selectedPlayers.value.push(player)
  }
}

function removePlayer(playerId: string) {
  selectedPlayers.value = selectedPlayers.value.filter(p => p.id !== playerId)
}

function isPlayerSelected(playerId: string): boolean {
  return selectedPlayers.value.some(p => p.id === playerId)
}

// Form submission
async function handleSubmit() {
  if (!isFormValid.value || !uploadSession.value) return

  submitting.value = true

  try {
    // Upload thumbnail if provided
    let thumbnailUrl: string | undefined
    if (thumbnailFile.value) {
      const thumbResponse = await $fetch<ApiResponse<UploadSession>>('/admin/upload/init', {
        baseURL: config.public.apiBase,
        method: 'POST',
        headers: { Authorization: `Bearer ${authStore.accessToken}` },
        body: {
          upload_type: 'thumbnail',
          file_name: thumbnailFile.value.name,
          content_type: thumbnailFile.value.type,
          file_size: thumbnailFile.value.size,
        },
      })

      if (thumbResponse.success && thumbResponse.data?.upload_url) {
        await fetch(thumbResponse.data.upload_url, {
          method: 'PUT',
          body: thumbnailFile.value,
          headers: { 'Content-Type': thumbnailFile.value.type },
        })
        // Construct the URL based on the s3_key
        thumbnailUrl = thumbResponse.data.s3_key
      }
    }

    // Create video record
    const response = await $fetch<ApiResponse<{ id: string }>>('/admin/videos', {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: {
        video_type: form.video_type,
        title: form.title,
        description: form.description || undefined,
        session_id: uploadSession.value.session_id,
        thumbnail_url: thumbnailUrl,
        duration_seconds: form.duration_seconds || undefined,
        tournament_id: form.tournament_id || undefined,
        player_ids: selectedPlayers.value.map(p => p.id),
      },
    })

    if (response.success) {
      toast.success('Video Created', 'Video has been uploaded and saved successfully')
      router.push('/admin/videos')
    } else {
      throw new Error('Failed to create video')
    }
  } catch (error) {
    console.error('Failed to create video:', error)
    toast.error('Error', 'Failed to save video. Please try again.')
  } finally {
    submitting.value = false
  }
}

// Helpers
function formatFileSize(bytes: number): string {
  if (bytes === 0) return '0 Bytes'
  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}
</script>
