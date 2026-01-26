<template>
  <div>
    <div class="mb-8">
      <h1 class="font-display text-2xl lg:text-3xl font-bold text-neutral-900">
        Settings
      </h1>
      <p class="mt-1 text-neutral-600">Configure platform settings and preferences.</p>
    </div>

    <!-- General Settings -->
    <UCard class="mb-6">
      <template #header>
        <h2 class="font-semibold text-neutral-900">General Settings</h2>
      </template>

      <div class="space-y-4">
        <UInput
          v-model="settings.site_name"
          label="Site Name"
          placeholder="Unicorn Sport"
        />

        <div>
          <label class="block text-sm font-medium text-neutral-700 mb-1">Site Description</label>
          <textarea
            v-model="settings.site_description"
            rows="3"
            class="w-full px-3 py-2 border border-neutral-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500"
            placeholder="Discover African football talent..."
          ></textarea>
        </div>

        <UInput
          v-model="settings.contact_email"
          label="Contact Email"
          type="email"
          placeholder="support@unicornsport.com"
        />
      </div>
    </UCard>

    <!-- API & Integrations -->
    <UCard class="mb-6">
      <template #header>
        <h2 class="font-semibold text-neutral-900">API & Integrations</h2>
      </template>

      <div class="space-y-4">
        <div class="flex items-center justify-between p-4 bg-neutral-50 rounded-lg">
          <div>
            <p class="font-medium text-neutral-900">AWS S3 Storage</p>
            <p class="text-sm text-neutral-500">Video and image storage</p>
          </div>
          <UBadge :variant="settings.s3_configured ? 'success' : 'warning'">
            {{ settings.s3_configured ? 'Connected' : 'Not Configured' }}
          </UBadge>
        </div>

        <div class="flex items-center justify-between p-4 bg-neutral-50 rounded-lg">
          <div>
            <p class="font-medium text-neutral-900">Stripe Payments</p>
            <p class="text-sm text-neutral-500">Subscription billing</p>
          </div>
          <UBadge :variant="settings.stripe_configured ? 'success' : 'warning'">
            {{ settings.stripe_configured ? 'Connected' : 'Not Configured' }}
          </UBadge>
        </div>

        <div class="flex items-center justify-between p-4 bg-neutral-50 rounded-lg">
          <div>
            <p class="font-medium text-neutral-900">Email Service</p>
            <p class="text-sm text-neutral-500">Transactional emails</p>
          </div>
          <UBadge :variant="settings.email_configured ? 'success' : 'warning'">
            {{ settings.email_configured ? 'Connected' : 'Not Configured' }}
          </UBadge>
        </div>
      </div>
    </UCard>

    <!-- Feature Flags -->
    <UCard class="mb-6">
      <template #header>
        <h2 class="font-semibold text-neutral-900">Feature Flags</h2>
      </template>

      <div class="space-y-4">
        <div class="flex items-center justify-between">
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
            <div class="w-11 h-6 bg-neutral-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-primary-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-neutral-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-primary-600"></div>
          </label>
        </div>

        <div class="flex items-center justify-between">
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
            <div class="w-11 h-6 bg-neutral-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-primary-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-neutral-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-primary-600"></div>
          </label>
        </div>

        <div class="flex items-center justify-between">
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
            <div class="w-11 h-6 bg-neutral-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-primary-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-neutral-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-primary-600"></div>
          </label>
        </div>
      </div>
    </UCard>

    <!-- Save Button -->
    <div class="flex justify-end">
      <UButton @click="saveSettings" :loading="saving">
        Save Settings
      </UButton>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { ApiResponse } from '~/types/index'

definePageMeta({
  layout: 'admin',
  middleware: 'admin',
})

const config = useRuntimeConfig()
const authStore = useAuthStore()
const toast = useToast()

const saving = ref(false)

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
    const response = await $fetch<ApiResponse<typeof settings>>('/admin/settings', {
      baseURL: config.public.apiBase,
      headers: {
        Authorization: `Bearer ${authStore.accessToken}`,
      },
    })

    if (response.success && response.data) {
      Object.assign(settings, response.data)
    }
  } catch (error) {
    console.error('Failed to fetch settings:', error)
  }
})

async function saveSettings() {
  saving.value = true

  try {
    const response = await $fetch<ApiResponse<null>>('/admin/settings', {
      baseURL: config.public.apiBase,
      method: 'PUT',
      headers: {
        Authorization: `Bearer ${authStore.accessToken}`,
      },
      body: settings,
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
