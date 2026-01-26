<template>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
  
  <!-- Global Toast Notifications -->
  <ClientOnly>
    <UiToast />
  </ClientOnly>
</template>

<script setup lang="ts">
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
