<template>
  <div class="min-h-screen bg-neutral-100">
    <!-- Admin Header -->
    <header class="bg-neutral-900 text-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex items-center justify-between h-16">
          <div class="flex items-center">
            <NuxtLink to="/admin" class="flex items-center">
              <span class="text-xl font-display font-bold text-primary-400">Unicorn Sport</span>
              <span class="ml-2 px-2 py-0.5 bg-red-600 text-xs rounded">ADMIN</span>
            </NuxtLink>
          </div>

          <div class="flex items-center space-x-4">
            <NuxtLink to="/" class="text-sm text-neutral-300 hover:text-white">
              View Site
            </NuxtLink>
            <div class="h-8 w-8 rounded-full bg-neutral-700 flex items-center justify-center">
              <span class="text-sm font-medium">{{ authStore.user?.email?.charAt(0).toUpperCase() }}</span>
            </div>
          </div>
        </div>
      </div>
    </header>

    <div class="flex">
      <!-- Sidebar -->
      <aside class="w-64 min-h-[calc(100vh-4rem)] bg-white border-r border-neutral-200">
        <nav class="p-4 space-y-1">
          <NuxtLink
            to="/admin"
            class="flex items-center px-3 py-2 text-sm font-medium rounded-lg hover:bg-neutral-100"
            :class="{ 'bg-primary-50 text-primary-700': $route.path === '/admin' }"
          >
            <svg class="w-5 h-5 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
            </svg>
            Dashboard
          </NuxtLink>

          <NuxtLink
            to="/admin/players"
            class="flex items-center px-3 py-2 text-sm font-medium rounded-lg hover:bg-neutral-100"
            :class="{ 'bg-primary-50 text-primary-700': $route.path.startsWith('/admin/players') }"
          >
            <svg class="w-5 h-5 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
            </svg>
            Players
          </NuxtLink>

          <NuxtLink
            to="/admin/academies"
            class="flex items-center px-3 py-2 text-sm font-medium rounded-lg hover:bg-neutral-100"
            :class="{ 'bg-primary-50 text-primary-700': $route.path.startsWith('/admin/academies') }"
          >
            <svg class="w-5 h-5 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
            </svg>
            Academies
          </NuxtLink>

          <NuxtLink
            to="/admin/videos"
            class="flex items-center px-3 py-2 text-sm font-medium rounded-lg hover:bg-neutral-100"
            :class="{ 'bg-primary-50 text-primary-700': $route.path.startsWith('/admin/videos') }"
          >
            <svg class="w-5 h-5 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
            </svg>
            Videos
          </NuxtLink>

          <NuxtLink
            to="/admin/users"
            class="flex items-center px-3 py-2 text-sm font-medium rounded-lg hover:bg-neutral-100"
            :class="{ 'bg-primary-50 text-primary-700': $route.path.startsWith('/admin/users') }"
          >
            <svg class="w-5 h-5 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13 5.197v1H21v-1a4 4 0 00-4-4h-1m-4-2a4 4 0 110-8 4 4 0 010 8z" />
            </svg>
            Users
          </NuxtLink>

          <NuxtLink
            to="/admin/contacts"
            class="flex items-center px-3 py-2 text-sm font-medium rounded-lg hover:bg-neutral-100"
            :class="{ 'bg-primary-50 text-primary-700': $route.path.startsWith('/admin/contacts') }"
          >
            <svg class="w-5 h-5 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
            </svg>
            Contact Requests
          </NuxtLink>

          <div class="pt-4 mt-4 border-t border-neutral-200">
            <NuxtLink
              to="/admin/settings"
              class="flex items-center px-3 py-2 text-sm font-medium rounded-lg hover:bg-neutral-100"
              :class="{ 'bg-primary-50 text-primary-700': $route.path.startsWith('/admin/settings') }"
            >
              <svg class="w-5 h-5 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
              </svg>
              Settings
            </NuxtLink>
          </div>
        </nav>
      </aside>

      <!-- Main Content -->
      <main class="flex-1 p-8">
        <slot />
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
const authStore = useAuthStore()
</script>
