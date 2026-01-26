<template>
  <div class="min-h-screen flex items-center justify-center bg-neutral-50">
    <div class="text-center px-4">
      <h1 class="font-display text-9xl font-bold text-primary-600">
        {{ error?.statusCode || '404' }}
      </h1>
      <h2 class="mt-4 text-2xl font-semibold text-neutral-900">
        {{ errorTitle }}
      </h2>
      <p class="mt-2 text-neutral-600 max-w-md mx-auto">
        {{ errorMessage }}
      </p>
      <div class="mt-8 flex justify-center gap-4">
        <UButton @click="handleError">
          Go Home
        </UButton>
        <UButton variant="outline" @click="$router.back()">
          Go Back
        </UButton>
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
