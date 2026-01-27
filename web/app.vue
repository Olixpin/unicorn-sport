<template>
  <NuxtLayout>
    <NuxtPage :key="route.path" />
  </NuxtLayout>
  
  <!-- Global Toast Notifications -->
  <ClientOnly>
    <UiToast />
    <UiConfirmDialog />
  </ClientOnly>
</template>

<script setup lang="ts">
const route = useRoute()

// Initialize auth state on app mount
const authStore = useAuthStore()
const subscriptionStore = useSubscriptionStore()

// Initialize auth from cookies (works on both server and client)
await authStore.initFromStorage()

// If authenticated, fetch subscription status
if (authStore.isAuthenticated) {
  subscriptionStore.fetchSubscription()
}
</script>
