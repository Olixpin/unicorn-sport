<template>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
  
  <!-- Global Toast Notifications -->
  <UToast />
</template>

<script setup lang="ts">
// Initialize auth state on app mount
const authStore = useAuthStore()
const subscriptionStore = useSubscriptionStore()

onMounted(() => {
  // Check for stored token and initialize auth
  authStore.initFromStorage()
  
  // If authenticated, fetch subscription status
  if (authStore.isAuthenticated) {
    subscriptionStore.fetchSubscription()
  }
})
</script>
