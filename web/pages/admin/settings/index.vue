<template>
  <div class="max-w-4xl mx-auto">
    <!-- Page Header -->
    <div class="mb-8">
      <h1 class="font-display text-2xl lg:text-3xl font-bold text-neutral-900">
        Settings
      </h1>
      <p class="mt-1 text-neutral-500">Configure platform settings and preferences</p>
    </div>

    <div class="space-y-6">
      <!-- General Settings -->
      <div class="bg-white rounded-2xl border border-neutral-200 overflow-hidden">
        <div class="px-6 py-4 border-b border-neutral-100 bg-neutral-50">
          <div class="flex items-center gap-3">
            <div class="w-8 h-8 bg-primary-100 rounded-lg flex items-center justify-center">
              <svg class="w-4 h-4 text-primary-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
              </svg>
            </div>
            <div>
              <h2 class="font-semibold text-neutral-900">General Settings</h2>
              <p class="text-sm text-neutral-500">Basic platform configuration</p>
            </div>
          </div>
        </div>

        <div class="p-6 space-y-5">
          <div>
            <label class="block text-sm font-medium text-neutral-700 mb-2">Site Name</label>
            <input
              v-model="settings.site_name"
              type="text"
              placeholder="Unicorn Sport"
              class="w-full px-4 py-2.5 rounded-xl border border-neutral-200 bg-white text-sm focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500 transition-all duration-200"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-neutral-700 mb-2">Site Description</label>
            <textarea
              v-model="settings.site_description"
              rows="3"
              placeholder="Discover African football talent..."
              class="w-full px-4 py-2.5 rounded-xl border border-neutral-200 bg-white text-sm focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500 transition-all duration-200 resize-none"
            ></textarea>
          </div>

          <div>
            <label class="block text-sm font-medium text-neutral-700 mb-2">Contact Email</label>
            <input
              v-model="settings.contact_email"
              type="email"
              placeholder="support@unicornsport.com"
              class="w-full px-4 py-2.5 rounded-xl border border-neutral-200 bg-white text-sm focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500 transition-all duration-200"
            />
          </div>
        </div>
      </div>

      <!-- API & Integrations -->
      <div class="bg-white rounded-2xl border border-neutral-200 overflow-hidden">
        <div class="px-6 py-4 border-b border-neutral-100 bg-neutral-50">
          <div class="flex items-center gap-3">
            <div class="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center">
              <svg class="w-4 h-4 text-blue-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1" />
              </svg>
            </div>
            <div>
              <h2 class="font-semibold text-neutral-900">API & Integrations</h2>
              <p class="text-sm text-neutral-500">External services and connections</p>
            </div>
          </div>
        </div>

        <div class="p-6 space-y-4">
          <!-- AWS S3 -->
          <div class="flex items-center justify-between p-4 bg-neutral-50 rounded-xl">
            <div class="flex items-center gap-4">
              <div class="w-10 h-10 bg-orange-100 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5 text-orange-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
                </svg>
              </div>
              <div>
                <p class="font-medium text-neutral-900">AWS S3 Storage</p>
                <p class="text-sm text-neutral-500">Video and image storage</p>
              </div>
            </div>
            <span 
              :class="[
                'inline-flex items-center gap-1.5 px-3 py-1 rounded-full text-xs font-medium',
                settings.s3_configured 
                  ? 'bg-emerald-100 text-emerald-700' 
                  : 'bg-amber-100 text-amber-700'
              ]"
            >
              <span :class="['w-1.5 h-1.5 rounded-full', settings.s3_configured ? 'bg-emerald-500' : 'bg-amber-500']"></span>
              {{ settings.s3_configured ? 'Connected' : 'Not Configured' }}
            </span>
          </div>

          <!-- Stripe -->
          <div class="flex items-center justify-between p-4 bg-neutral-50 rounded-xl">
            <div class="flex items-center gap-4">
              <div class="w-10 h-10 bg-purple-100 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5 text-purple-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z" />
                </svg>
              </div>
              <div>
                <p class="font-medium text-neutral-900">Stripe Payments</p>
                <p class="text-sm text-neutral-500">Subscription billing</p>
              </div>
            </div>
            <span 
              :class="[
                'inline-flex items-center gap-1.5 px-3 py-1 rounded-full text-xs font-medium',
                settings.stripe_configured 
                  ? 'bg-emerald-100 text-emerald-700' 
                  : 'bg-amber-100 text-amber-700'
              ]"
            >
              <span :class="['w-1.5 h-1.5 rounded-full', settings.stripe_configured ? 'bg-emerald-500' : 'bg-amber-500']"></span>
              {{ settings.stripe_configured ? 'Connected' : 'Not Configured' }}
            </span>
          </div>

          <!-- Email Service -->
          <div class="flex items-center justify-between p-4 bg-neutral-50 rounded-xl">
            <div class="flex items-center gap-4">
              <div class="w-10 h-10 bg-blue-100 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5 text-blue-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                </svg>
              </div>
              <div>
                <p class="font-medium text-neutral-900">Email Service</p>
                <p class="text-sm text-neutral-500">Transactional emails</p>
              </div>
            </div>
            <span 
              :class="[
                'inline-flex items-center gap-1.5 px-3 py-1 rounded-full text-xs font-medium',
                settings.email_configured 
                  ? 'bg-emerald-100 text-emerald-700' 
                  : 'bg-amber-100 text-amber-700'
              ]"
            >
              <span :class="['w-1.5 h-1.5 rounded-full', settings.email_configured ? 'bg-emerald-500' : 'bg-amber-500']"></span>
              {{ settings.email_configured ? 'Connected' : 'Not Configured' }}
            </span>
          </div>
        </div>
      </div>

      <!-- Feature Flags -->
      <div class="bg-white rounded-2xl border border-neutral-200 overflow-hidden">
        <div class="px-6 py-4 border-b border-neutral-100 bg-neutral-50">
          <div class="flex items-center gap-3">
            <div class="w-8 h-8 bg-amber-100 rounded-lg flex items-center justify-center">
              <svg class="w-4 h-4 text-amber-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 21v-4m0 0V5a2 2 0 012-2h6.5l1 1H21l-3 6 3 6h-8.5l-1-1H5a2 2 0 00-2 2zm9-13.5V9" />
              </svg>
            </div>
            <div>
              <h2 class="font-semibold text-neutral-900">Feature Flags</h2>
              <p class="text-sm text-neutral-500">Enable or disable platform features</p>
            </div>
          </div>
        </div>

        <div class="p-6 space-y-4">
          <!-- Registration -->
          <div class="flex items-center justify-between p-4 bg-neutral-50 rounded-xl">
            <div>
              <p class="font-medium text-neutral-900">User Registration</p>
              <p class="text-sm text-neutral-500">Allow new users to register</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input
                type="checkbox"
                v-model="settings.registration_enabled"
                class="sr-only peer"
              />
              <div class="w-12 h-6 bg-neutral-300 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-primary-500/20 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[3px] after:bg-white after:rounded-full after:h-5 after:w-5 after:shadow after:transition-all peer-checked:bg-primary-500"></div>
            </label>
          </div>

          <!-- Video Uploads -->
          <div class="flex items-center justify-between p-4 bg-neutral-50 rounded-xl">
            <div>
              <p class="font-medium text-neutral-900">Video Uploads</p>
              <p class="text-sm text-neutral-500">Allow players to upload videos</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input
                type="checkbox"
                v-model="settings.video_uploads_enabled"
                class="sr-only peer"
              />
              <div class="w-12 h-6 bg-neutral-300 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-primary-500/20 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[3px] after:bg-white after:rounded-full after:h-5 after:w-5 after:shadow after:transition-all peer-checked:bg-primary-500"></div>
            </label>
          </div>

          <!-- Contact Requests -->
          <div class="flex items-center justify-between p-4 bg-neutral-50 rounded-xl">
            <div>
              <p class="font-medium text-neutral-900">Contact Requests</p>
              <p class="text-sm text-neutral-500">Allow scouts to request player contact info</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input
                type="checkbox"
                v-model="settings.contact_requests_enabled"
                class="sr-only peer"
              />
              <div class="w-12 h-6 bg-neutral-300 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-primary-500/20 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[3px] after:bg-white after:rounded-full after:h-5 after:w-5 after:shadow after:transition-all peer-checked:bg-primary-500"></div>
            </label>
          </div>
        </div>
      </div>

      <!-- Danger Zone -->
      <div class="bg-white rounded-2xl border border-red-200 overflow-hidden">
        <div class="px-6 py-4 border-b border-red-100 bg-red-50">
          <div class="flex items-center gap-3">
            <div class="w-8 h-8 bg-red-100 rounded-lg flex items-center justify-center">
              <svg class="w-4 h-4 text-red-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
              </svg>
            </div>
            <div>
              <h2 class="font-semibold text-red-900">Danger Zone</h2>
              <p class="text-sm text-red-600">Irreversible and destructive actions</p>
            </div>
          </div>
        </div>

        <div class="p-6">
          <div class="flex items-center justify-between">
            <div>
              <p class="font-medium text-neutral-900">Clear All Cache</p>
              <p class="text-sm text-neutral-500">Remove all cached data and regenerate</p>
            </div>
            <button class="px-4 py-2 text-sm font-medium text-red-600 border border-red-200 rounded-xl hover:bg-red-50 transition-colors">
              Clear Cache
            </button>
          </div>
        </div>
      </div>

      <!-- Save Button -->
      <div class="flex items-center justify-between pt-4">
        <p class="text-sm text-neutral-500">
          Changes are saved automatically when you modify settings
        </p>
        <button 
          @click="saveSettings" 
          :disabled="saving"
          class="flex items-center gap-2 px-6 py-2.5 bg-gradient-to-r from-primary-500 to-primary-600 text-white font-semibold rounded-xl hover:from-primary-600 hover:to-primary-700 shadow-lg shadow-primary-500/25 disabled:opacity-50 transition-all duration-200"
        >
          <svg v-if="saving" class="w-5 h-5 animate-spin" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          <svg v-else class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
          </svg>
          {{ saving ? 'Saving...' : 'Save Settings' }}
        </button>
      </div>
    </div>
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
    }
  } catch (error) {
    console.error('Failed to fetch settings:', error)
  } finally {
    loading.value = false
  }
})

async function saveSettings() {
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

    const response = await api<ApiResponse<null>>('/admin/settings', {
      method: 'PUT',
      body: { settings: settingsMap },
    })

    if (response.success) {
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
