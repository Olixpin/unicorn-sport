<template>
  <div class="min-h-screen flex items-center justify-center bg-neutral-950">
    <div class="text-center px-4 max-w-lg mx-auto">
      <!-- Glowing background effect -->
      <div class="absolute inset-0 overflow-hidden pointer-events-none">
        <div class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[500px] h-[500px] bg-primary-500/10 rounded-full blur-[100px]"></div>
      </div>
      
      <div class="relative">
        <h1 class="font-display text-8xl sm:text-9xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-primary-400 to-emerald-400">
          {{ error?.statusCode || '404' }}
        </h1>
        <h2 class="mt-4 text-2xl sm:text-3xl font-semibold text-white">
          {{ errorTitle }}
        </h2>
        <p class="mt-4 text-neutral-400 max-w-md mx-auto">
          {{ errorMessage }}
        </p>
        <div class="mt-8 flex flex-col sm:flex-row justify-center gap-4">
          <button 
            @click="handleError"
            class="px-6 py-3 bg-gradient-to-r from-primary-500 to-emerald-500 text-white font-semibold rounded-xl hover:shadow-lg hover:shadow-primary-500/25 transition-all"
          >
            Go Home
          </button>
          <button 
            @click="$router.back()"
            class="px-6 py-3 bg-white/5 text-white font-semibold rounded-xl border border-white/10 hover:bg-white/10 transition-all"
          >
            Go Back
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { NuxtError } from '#app'

const props = defineProps<{
  error: NuxtError
}>()

const errorTitle = computed(() => {
  switch (props.error?.statusCode) {
    case 404:
      return 'Page Not Found'
    case 403:
      return 'Access Denied'
    case 500:
      return 'Server Error'
    default:
      return 'Something Went Wrong'
  }
})

const errorMessage = computed(() => {
  switch (props.error?.statusCode) {
    case 404:
      return "The page you're looking for doesn't exist or has been moved."
    case 403:
      return "You don't have permission to access this page."
    case 500:
      return 'An unexpected error occurred. Please try again later.'
    default:
      return props.error?.message || 'An unexpected error occurred.'
  }
})

function handleError() {
  clearError({ redirect: '/' })
}
</script>
