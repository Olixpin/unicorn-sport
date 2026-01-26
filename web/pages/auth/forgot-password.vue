<template>
  <div>
    <!-- Success State -->
    <div v-if="sent" class="text-center">
      <!-- Success icon -->
      <div class="mx-auto w-16 h-16 rounded-full bg-gradient-to-br from-primary-500 to-emerald-500 flex items-center justify-center mb-6">
        <svg class="w-8 h-8 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M5 13l4 4L19 7" />
        </svg>
      </div>
      
      <h1 class="text-2xl sm:text-3xl font-display font-bold text-neutral-900 mb-3">
        Check your email
      </h1>
      
      <p class="text-neutral-600 mb-2">
        We sent a password reset link to
      </p>
      <p class="text-neutral-900 font-semibold mb-8">
        {{ form.email }}
      </p>

      <div class="space-y-4">
        <NuxtLink to="/auth/login">
          <button class="w-full px-6 py-3.5 rounded-xl bg-primary-600 text-white font-semibold hover:bg-primary-700 transition-all duration-200 hover:shadow-lg hover:shadow-primary-500/25">
            Back to sign in
          </button>
        </NuxtLink>
        
        <p class="text-sm text-neutral-500">
          Didn't receive the email? 
          <button 
            @click="resendEmail" 
            :disabled="resendCooldown > 0"
            class="text-primary-600 hover:text-primary-700 font-medium disabled:text-neutral-400 disabled:cursor-not-allowed"
          >
            {{ resendCooldown > 0 ? `Resend in ${resendCooldown}s` : 'Click to resend' }}
          </button>
        </p>
      </div>
    </div>

    <!-- Form State -->
    <div v-else>
      <!-- Header -->
      <div class="mb-8">
        <NuxtLink 
          to="/auth/login" 
          class="inline-flex items-center gap-2 text-sm text-neutral-500 hover:text-neutral-700 transition-colors mb-6"
        >
          <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
          </svg>
          Back to sign in
        </NuxtLink>
        
        <h1 class="text-2xl sm:text-3xl font-display font-bold text-neutral-900">
          Forgot password?
        </h1>
        <p class="mt-2 text-neutral-600">
          No worries, we'll send you reset instructions.
        </p>
      </div>

      <!-- Form -->
      <form @submit.prevent="handleSubmit" class="space-y-6">
        <div>
          <label for="email" class="block text-sm font-medium text-neutral-700 mb-1.5">
            Email address
          </label>
          <input
            id="email"
            v-model="form.email"
            type="email"
            autocomplete="email"
            required
            placeholder="you@example.com"
            :class="[
              'w-full px-4 py-3 rounded-xl border bg-white text-neutral-900 placeholder-neutral-400 transition-all duration-200',
              'focus:outline-none focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500',
              error ? 'border-red-300 bg-red-50' : 'border-neutral-200 hover:border-neutral-300'
            ]"
          />
          <p v-if="error" class="mt-1.5 text-sm text-red-600">{{ error }}</p>
        </div>

        <!-- Submit button -->
        <button
          type="submit"
          :disabled="loading"
          :class="[
            'w-full flex items-center justify-center gap-2 px-6 py-3.5 rounded-xl font-semibold text-white transition-all duration-200',
            loading 
              ? 'bg-primary-400 cursor-not-allowed' 
              : 'bg-primary-600 hover:bg-primary-700 hover:shadow-lg hover:shadow-primary-500/25 active:scale-[0.98]'
          ]"
        >
          <svg v-if="loading" class="w-5 h-5 animate-spin" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          <span>{{ loading ? 'Sending...' : 'Reset password' }}</span>
        </button>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { ApiResponse } from '~/types'

definePageMeta({
  layout: 'auth',
  middleware: 'guest',
})

const config = useRuntimeConfig()

const loading = ref(false)
const sent = ref(false)
const error = ref('')
const resendCooldown = ref(0)

const form = reactive({
  email: '',
})

let cooldownInterval: ReturnType<typeof setInterval> | null = null

async function handleSubmit() {
  loading.value = true
  error.value = ''

  try {
    const response = await $fetch<ApiResponse<null>>('/auth/forgot-password', {
      method: 'POST',
      baseURL: config.public.apiBase,
      body: { email: form.email },
    })

    if (response.success) {
      sent.value = true
      startResendCooldown()
    } else {
      error.value = response.message || 'Failed to send reset link'
    }
  } catch (e: unknown) {
    error.value = (e as { data?: { message?: string } })?.data?.message || 'Something went wrong. Please try again.'
  } finally {
    loading.value = false
  }
}

async function resendEmail() {
  if (resendCooldown.value > 0) return
  
  loading.value = true
  try {
    const response = await $fetch<ApiResponse<null>>('/auth/forgot-password', {
      method: 'POST',
      baseURL: config.public.apiBase,
      body: { email: form.email },
    })
    
    if (response.success) {
      startResendCooldown()
    }
  } catch (e) {
    // Silently fail for resend
  } finally {
    loading.value = false
  }
}

function startResendCooldown() {
  resendCooldown.value = 60
  
  if (cooldownInterval) clearInterval(cooldownInterval)
  
  cooldownInterval = setInterval(() => {
    resendCooldown.value--
    if (resendCooldown.value <= 0 && cooldownInterval) {
      clearInterval(cooldownInterval)
    }
  }, 1000)
}

onUnmounted(() => {
  if (cooldownInterval) clearInterval(cooldownInterval)
})
</script>
