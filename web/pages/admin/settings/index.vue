<template>
  <div class="max-w-4xl mx-auto">
    <!-- Page Header -->
    <div class="mb-6">
      <h1 class="font-display text-xl sm:text-2xl lg:text-3xl font-bold text-neutral-900">
        Settings
      </h1>
      <p class="hidden sm:block mt-1 text-sm text-neutral-500">Configure platform settings and preferences</p>
    </div>

    <div class="space-y-4 sm:space-y-6">
      <!-- General Settings -->
      <div class="bg-white rounded-xl border border-neutral-200 overflow-hidden">
        <div class="px-4 sm:px-6 py-3 sm:py-4 border-b border-neutral-100 bg-neutral-50/50">
          <div class="flex items-center gap-3">
            <div class="w-8 h-8 bg-primary-100 rounded-lg flex items-center justify-center">
              <svg class="w-4 h-4 text-primary-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
              </svg>
            </div>
            <div>
              <h2 class="font-semibold text-neutral-900 text-sm sm:text-base">General Settings</h2>
              <p class="text-xs sm:text-sm text-neutral-500">Basic platform configuration</p>
            </div>
          </div>
        </div>

        <div class="p-4 sm:p-6 space-y-4">
          <div>
            <label class="block text-xs sm:text-sm font-medium text-neutral-700 mb-1.5">Site Name</label>
            <input
              v-model="settings.site_name"
              type="text"
              placeholder="Unicorn Sport"
              class="w-full px-3 sm:px-4 py-2 sm:py-2.5 rounded-lg border border-neutral-200 bg-white text-sm focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500 transition-all"
            />
          </div>

          <div>
            <label class="block text-xs sm:text-sm font-medium text-neutral-700 mb-1.5">Site Description</label>
            <textarea
              v-model="settings.site_description"
              rows="3"
              placeholder="Discover African football talent..."
              class="w-full px-3 sm:px-4 py-2 sm:py-2.5 rounded-lg border border-neutral-200 bg-white text-sm focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500 transition-all resize-none"
            ></textarea>
          </div>

          <div>
            <label class="block text-xs sm:text-sm font-medium text-neutral-700 mb-1.5">Contact Email</label>
            <input
              v-model="settings.contact_email"
              type="email"
              placeholder="support@unicornsport.com"
              class="w-full px-3 sm:px-4 py-2 sm:py-2.5 rounded-lg border border-neutral-200 bg-white text-sm focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500 transition-all"
            />
          </div>
        </div>
      </div>

      <!-- API & Integrations -->
      <div class="bg-white rounded-xl border border-neutral-200 overflow-hidden">
        <div class="px-4 sm:px-6 py-3 sm:py-4 border-b border-neutral-100 bg-neutral-50/50">
          <div class="flex items-center gap-3">
            <div class="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center">
              <svg class="w-4 h-4 text-blue-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1" />
              </svg>
            </div>
            <div>
              <h2 class="font-semibold text-neutral-900 text-sm sm:text-base">API & Integrations</h2>
              <p class="text-xs sm:text-sm text-neutral-500">External services</p>
            </div>
          </div>
        </div>

        <div class="p-3 sm:p-6 space-y-2 sm:space-y-3">
          <!-- AWS S3 -->
          <div class="flex items-center justify-between gap-3 p-3 bg-neutral-50 rounded-lg">
            <div class="flex items-center gap-3">
              <div class="w-8 h-8 bg-orange-100 rounded-lg flex items-center justify-center flex-shrink-0">
                <svg class="w-4 h-4 text-orange-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
                </svg>
              </div>
              <div class="min-w-0">
                <p class="font-medium text-neutral-900 text-sm">AWS S3</p>
                <p class="text-[11px] text-neutral-500">Storage</p>
              </div>
            </div>
            <span 
              :class="[
                'inline-flex items-center gap-1 px-2 py-0.5 rounded-full text-[10px] font-medium',
                settings.s3_configured 
                  ? 'bg-emerald-100 text-emerald-700' 
                  : 'bg-neutral-100 text-neutral-500'
              ]"
            >
              <span :class="['w-1 h-1 rounded-full', settings.s3_configured ? 'bg-emerald-500' : 'bg-neutral-400']"></span>
              {{ settings.s3_configured ? 'Connected' : 'Not Set' }}
            </span>
          </div>

          <!-- Stripe -->
          <div class="flex items-center justify-between gap-3 p-3 bg-neutral-50 rounded-lg">
            <div class="flex items-center gap-3">
              <div class="w-8 h-8 bg-purple-100 rounded-lg flex items-center justify-center flex-shrink-0">
                <svg class="w-4 h-4 text-purple-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z" />
                </svg>
              </div>
              <div class="min-w-0">
                <p class="font-medium text-neutral-900 text-sm">Stripe</p>
                <p class="text-[11px] text-neutral-500">Payments</p>
              </div>
            </div>
            <span 
              :class="[
                'inline-flex items-center gap-1 px-2 py-0.5 rounded-full text-[10px] font-medium',
                settings.stripe_configured 
                  ? 'bg-emerald-100 text-emerald-700' 
                  : 'bg-neutral-100 text-neutral-500'
              ]"
            >
              <span :class="['w-1 h-1 rounded-full', settings.stripe_configured ? 'bg-emerald-500' : 'bg-neutral-400']"></span>
              {{ settings.stripe_configured ? 'Connected' : 'Not Set' }}
            </span>
          </div>

          <!-- Email Service -->
          <div class="flex items-center justify-between gap-3 p-3 bg-neutral-50 rounded-lg">
            <div class="flex items-center gap-3">
              <div class="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center flex-shrink-0">
                <svg class="w-4 h-4 text-blue-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                </svg>
              </div>
              <div class="min-w-0">
                <p class="font-medium text-neutral-900 text-sm">Email</p>
                <p class="text-[11px] text-neutral-500">Transactional</p>
              </div>
            </div>
            <span 
              :class="[
                'inline-flex items-center gap-1 px-2 py-0.5 rounded-full text-[10px] font-medium',
                settings.email_configured 
                  ? 'bg-emerald-100 text-emerald-700' 
                  : 'bg-neutral-100 text-neutral-500'
              ]"
            >
              <span :class="['w-1 h-1 rounded-full', settings.email_configured ? 'bg-emerald-500' : 'bg-neutral-400']"></span>
              {{ settings.email_configured ? 'Connected' : 'Not Set' }}
            </span>
          </div>
        </div>
      </div>

      <!-- Feature Flags -->
      <div class="bg-white rounded-xl border border-neutral-200 overflow-hidden">
        <div class="px-4 py-3 border-b border-neutral-100 bg-neutral-50/50">
          <div class="flex items-center gap-2.5">
            <div class="w-7 h-7 bg-amber-100 rounded-lg flex items-center justify-center">
              <svg class="w-3.5 h-3.5 text-amber-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 21v-4m0 0V5a2 2 0 012-2h6.5l1 1H21l-3 6 3 6h-8.5l-1-1H5a2 2 0 00-2 2zm9-13.5V9" />
              </svg>
            </div>
            <div>
              <h2 class="font-semibold text-neutral-900 text-sm">Feature Flags</h2>
              <p class="text-[11px] text-neutral-500 hidden sm:block">Enable or disable platform features</p>
            </div>
          </div>
        </div>

        <div class="p-4 space-y-2">
          <!-- Registration -->
          <div class="flex items-center justify-between p-3 bg-neutral-50 rounded-lg">
            <div class="min-w-0">
              <p class="font-medium text-neutral-900 text-sm">User Registration</p>
              <p class="text-[11px] text-neutral-500">Allow new signups</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer flex-shrink-0">
              <input
                type="checkbox"
                v-model="settings.registration_enabled"
                class="sr-only peer"
              />
              <div class="w-11 h-6 bg-neutral-300 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-primary-500/20 rounded-full peer peer-checked:after:translate-x-5 after:content-[''] after:absolute after:top-0.5 after:left-0.5 after:bg-white after:rounded-full after:h-5 after:w-5 after:shadow-sm after:transition-all peer-checked:bg-primary-500"></div>
            </label>
          </div>

          <!-- Video Uploads -->
          <div class="flex items-center justify-between p-3 bg-neutral-50 rounded-lg">
            <div class="min-w-0">
              <p class="font-medium text-neutral-900 text-sm">Video Uploads</p>
              <p class="text-[11px] text-neutral-500">Allow player videos</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer flex-shrink-0">
              <input
                type="checkbox"
                v-model="settings.video_uploads_enabled"
                class="sr-only peer"
              />
              <div class="w-11 h-6 bg-neutral-300 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-primary-500/20 rounded-full peer peer-checked:after:translate-x-5 after:content-[''] after:absolute after:top-0.5 after:left-0.5 after:bg-white after:rounded-full after:h-5 after:w-5 after:shadow-sm after:transition-all peer-checked:bg-primary-500"></div>
            </label>
          </div>

          <!-- Contact Requests -->
          <div class="flex items-center justify-between p-3 bg-neutral-50 rounded-lg">
            <div class="min-w-0">
              <p class="font-medium text-neutral-900 text-sm">Contact Requests</p>
              <p class="text-[11px] text-neutral-500">Allow scout requests</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer flex-shrink-0">
              <input
                type="checkbox"
                v-model="settings.contact_requests_enabled"
                class="sr-only peer"
              />
              <div class="w-11 h-6 bg-neutral-300 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-primary-500/20 rounded-full peer peer-checked:after:translate-x-5 after:content-[''] after:absolute after:top-0.5 after:left-0.5 after:bg-white after:rounded-full after:h-5 after:w-5 after:shadow-sm after:transition-all peer-checked:bg-primary-500"></div>
            </label>
          </div>
        </div>
      </div>

      <!-- Danger Zone -->
      <div class="bg-white rounded-xl border border-neutral-200 overflow-hidden">
        <div class="px-4 py-3 border-b border-neutral-100">
          <div class="flex items-center gap-2.5">
            <div class="w-7 h-7 bg-red-50 rounded-lg flex items-center justify-center">
              <svg class="w-3.5 h-3.5 text-red-500" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
              </svg>
            </div>
            <h2 class="font-semibold text-neutral-900 text-sm">Danger Zone</h2>
          </div>
        </div>

        <div class="p-4">
          <div class="flex items-center justify-between gap-3 p-3 bg-neutral-50 rounded-lg">
            <div class="min-w-0">
              <p class="font-medium text-neutral-900 text-sm">Clear Cache</p>
              <p class="text-[11px] text-neutral-500">Remove cached data</p>
            </div>
            <button class="flex-shrink-0 px-3 py-1.5 text-xs font-medium text-red-600 bg-white border border-red-200 rounded-lg hover:bg-red-50 active:scale-[0.98] transition-all">
              Clear
            </button>
          </div>
        </div>
      </div>
    </div>
    
    <!-- Fixed Bottom Save Bar -->
    <div class="fixed bottom-0 left-0 right-0 lg:left-64 bg-white/90 backdrop-blur-lg border-t border-neutral-200 px-4 py-2.5 z-40">
      <div class="max-w-4xl mx-auto flex items-center justify-between gap-3">
        <div class="flex items-center gap-2 text-xs text-neutral-500">
          <svg class="w-3.5 h-3.5 text-emerald-500" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
          </svg>
          <span class="hidden sm:inline">Auto-saved</span>
        </div>
        <button 
          @click="saveSettings" 
          :disabled="saving || !hasChanges"
          :class="[
            'flex items-center gap-1.5 px-4 py-2 text-sm font-medium rounded-lg transition-all',
            hasChanges 
              ? 'text-white bg-gradient-to-r from-primary-600 to-emerald-600 shadow-sm hover:shadow-md active:scale-[0.98]' 
              : 'text-neutral-400 bg-neutral-100 cursor-not-allowed'
          ]"
        >
          <svg v-if="saving" class="w-3.5 h-3.5 animate-spin" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          <svg v-else class="w-3.5 h-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2.5">
            <path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7" />
          </svg>
          <span>{{ saving ? 'Saving...' : 'Save' }}</span>
        </button>
      </div>
    </div>
    
    <!-- Bottom padding for fixed bar -->
    <div class="h-16"></div>
  </div>
</template>

<script setup lang="ts">
import type { ApiResponse } from '~/types/index'

definePageMeta({
  layout: 'admin',
  middleware: 'admin',
})

const api = useApi()
const toast = useToast()

const saving = ref(false)
const loading = ref(true)

// Track original settings to detect changes
const originalSettings = ref<Record<string, unknown>>({})

const settings = reactive({
  site_name: 'Unicorn Sport',
  site_description: 'Discover and connect with emerging African football talent',
  contact_email: 'support@unicornsport.com',
  s3_configured: false,
  stripe_configured: false,
  email_configured: false,
  registration_enabled: true,
  video_uploads_enabled: true,
  contact_requests_enabled: true,
})

// Check if any settings have changed
const hasChanges = computed(() => {
  return Object.keys(originalSettings.value).length > 0 && 
    Object.keys(settings).some(key => settings[key as keyof typeof settings] !== originalSettings.value[key])
})

// Fetch settings on mount
onMounted(async () => {
  try {
    const response = await api.get<ApiResponse<Record<string, string>>>('/admin/settings', {}, true)

    if (response.success && response.data) {
      // Map key-value pairs to settings object
      const data = response.data
      settings.site_name = data.site_name || 'Unicorn Sport'
      settings.site_description = data.site_description || 'Discover and connect with emerging African football talent'
      settings.contact_email = data.contact_email || 'support@unicornsport.com'
      settings.s3_configured = data.s3_configured === 'true'
      settings.stripe_configured = data.stripe_configured === 'true'
      settings.email_configured = data.email_configured === 'true'
      settings.registration_enabled = data.registration_enabled !== 'false'
      settings.video_uploads_enabled = data.video_uploads_enabled !== 'false'
      settings.contact_requests_enabled = data.contact_requests_enabled !== 'false'
      
      // Store original settings for change detection
      originalSettings.value = { ...settings }
    }
  } catch (error) {
    console.error('Failed to fetch settings:', error)
  } finally {
    loading.value = false
  }
})

async function saveSettings() {
  if (!hasChanges.value) return
  
  saving.value = true

  try {
    // Convert settings to key-value format
    const settingsMap: Record<string, string> = {
      site_name: settings.site_name,
      site_description: settings.site_description,
      contact_email: settings.contact_email,
      s3_configured: settings.s3_configured.toString(),
      stripe_configured: settings.stripe_configured.toString(),
      email_configured: settings.email_configured.toString(),
      registration_enabled: settings.registration_enabled.toString(),
      video_uploads_enabled: settings.video_uploads_enabled.toString(),
      contact_requests_enabled: settings.contact_requests_enabled.toString(),
    }

    const response = await api.put<ApiResponse<null>>('/admin/settings', { settings: settingsMap }, true)

    if (response.success) {
      originalSettings.value = { ...settings }
      toast.success('Settings Saved', 'Platform settings have been updated')
    } else {
      toast.error('Error', response.message || 'Failed to save settings')
    }
  } catch (error: unknown) {
    const err = error as { data?: { message?: string } }
    toast.error('Error', err.data?.message || 'An unexpected error occurred')
  } finally {
    saving.value = false
  }
}
</script>
