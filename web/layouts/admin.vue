<template>
  <div class="min-h-screen bg-neutral-50">
    <!-- Admin Header -->
    <header class="fixed top-0 left-0 right-0 z-50 bg-white border-b border-neutral-200">
      <div class="flex items-center justify-between h-16 px-4 lg:px-6">
        <!-- Left: Logo & Toggle -->
        <div class="flex items-center gap-4">
          <button 
            @click="sidebarOpen = !sidebarOpen"
            class="lg:hidden p-2 rounded-lg hover:bg-neutral-100 transition-colors"
          >
            <svg class="w-5 h-5 text-neutral-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
            </svg>
          </button>
          
          <NuxtLink to="/admin" class="flex items-center gap-3">
            <div class="w-9 h-9 bg-gradient-to-br from-primary-500 to-primary-600 rounded-xl flex items-center justify-center shadow-lg shadow-primary-500/25">
              <span class="text-white text-lg">🦄</span>
            </div>
            <div class="hidden sm:block">
              <span class="font-display font-bold text-neutral-900">Unicorn Sport</span>
              <span class="ml-2 px-2 py-0.5 bg-red-100 text-red-700 text-xs font-bold rounded-full">ADMIN</span>
            </div>
          </NuxtLink>
        </div>

        <!-- Right: Actions -->
        <div class="flex items-center gap-3">
          <!-- Search -->
          <div class="hidden md:flex items-center">
            <div class="relative">
              <svg class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
              <input 
                type="text" 
                placeholder="Search..." 
                class="w-64 pl-9 pr-4 py-2 bg-neutral-100 border-0 rounded-xl text-sm placeholder-neutral-500 focus:outline-none focus:ring-2 focus:ring-primary-500/20 focus:bg-white transition-all"
              />
              <kbd class="absolute right-3 top-1/2 -translate-y-1/2 px-1.5 py-0.5 bg-neutral-200 text-neutral-500 text-xs rounded">⌘K</kbd>
            </div>
          </div>

          <!-- View Site -->
          <NuxtLink 
            to="/" 
            class="hidden sm:flex items-center gap-2 px-3 py-2 text-sm text-neutral-600 hover:text-neutral-900 hover:bg-neutral-100 rounded-lg transition-colors"
          >
            <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
            </svg>
            View Site
          </NuxtLink>

          <!-- Notifications -->
          <button class="relative p-2 text-neutral-500 hover:text-neutral-700 hover:bg-neutral-100 rounded-lg transition-colors">
            <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
            </svg>
            <span class="absolute top-1.5 right-1.5 w-2 h-2 bg-red-500 rounded-full"></span>
          </button>

          <!-- Profile Dropdown -->
          <div class="relative">
            <button 
              @click="profileOpen = !profileOpen"
              class="flex items-center gap-2 p-1.5 rounded-xl hover:bg-neutral-100 transition-colors"
            >
              <div class="w-8 h-8 rounded-lg bg-gradient-to-br from-primary-500 to-emerald-500 flex items-center justify-center shadow-md">
                <span class="text-white text-sm font-bold">{{ userInitials }}</span>
              </div>
              <svg class="w-4 h-4 text-neutral-400 hidden sm:block" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
              </svg>
            </button>

            <!-- Dropdown -->
            <Transition
              enter-active-class="transition duration-100 ease-out"
              enter-from-class="transform scale-95 opacity-0"
              enter-to-class="transform scale-100 opacity-100"
              leave-active-class="transition duration-75 ease-in"
              leave-from-class="transform scale-100 opacity-100"
              leave-to-class="transform scale-95 opacity-0"
            >
              <div 
                v-if="profileOpen"
                class="absolute right-0 mt-2 w-56 bg-white rounded-xl shadow-lg border border-neutral-200 py-2 z-50"
              >
                <div class="px-4 py-3 border-b border-neutral-100">
                  <p class="font-medium text-neutral-900 text-sm">{{ authStore.user?.first_name }} {{ authStore.user?.last_name }}</p>
                  <p class="text-xs text-neutral-500 truncate">{{ authStore.user?.email }}</p>
                </div>
                <NuxtLink to="/admin/settings" class="flex items-center gap-3 px-4 py-2 text-sm text-neutral-700 hover:bg-neutral-50">
                  <svg class="w-4 h-4 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                  </svg>
                  Settings
                </NuxtLink>
                <button 
                  @click="logout"
                  class="w-full flex items-center gap-3 px-4 py-2 text-sm text-red-600 hover:bg-red-50"
                >
                  <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
                  </svg>
                  Sign out
                </button>
              </div>
            </Transition>
          </div>
        </div>
      </div>
    </header>

    <div class="flex pt-16">
      <!-- Mobile Overlay -->
      <div 
        v-if="sidebarOpen" 
        class="fixed inset-0 bg-black/50 z-40 lg:hidden"
        @click="sidebarOpen = false"
      ></div>

      <!-- Sidebar -->
      <aside 
        :class="[
          'fixed lg:sticky top-16 left-0 z-40 w-72 h-[calc(100vh-4rem)] bg-white border-r border-neutral-200 transition-transform duration-300',
          sidebarOpen ? 'translate-x-0' : '-translate-x-full lg:translate-x-0'
        ]"
      >
        <div class="flex flex-col h-full">
          <!-- Navigation -->
          <nav class="flex-1 p-4 space-y-1 overflow-y-auto">
            <!-- Main Section -->
            <p class="px-3 py-2 text-xs font-semibold text-neutral-400 uppercase tracking-wider">Main</p>
            
            <NuxtLink
              to="/admin"
              class="group flex items-center gap-3 px-3 py-2.5 rounded-xl transition-all duration-200"
              :class="isActive('/admin') ? 'bg-primary-50 text-primary-700' : 'text-neutral-600 hover:bg-neutral-100'"
            >
              <div 
                :class="[
                  'w-9 h-9 rounded-lg flex items-center justify-center transition-colors',
                  isActive('/admin') ? 'bg-primary-100' : 'bg-neutral-100 group-hover:bg-neutral-200'
                ]"
              >
                <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 5a1 1 0 011-1h14a1 1 0 011 1v2a1 1 0 01-1 1H5a1 1 0 01-1-1V5zM4 13a1 1 0 011-1h6a1 1 0 011 1v6a1 1 0 01-1 1H5a1 1 0 01-1-1v-6zM16 13a1 1 0 011-1h2a1 1 0 011 1v6a1 1 0 01-1 1h-2a1 1 0 01-1-1v-6z" />
                </svg>
              </div>
              <span class="font-medium">Dashboard</span>
            </NuxtLink>

            <!-- Content Section -->
            <p class="px-3 py-2 mt-6 text-xs font-semibold text-neutral-400 uppercase tracking-wider">Content</p>

            <NuxtLink
              to="/admin/players"
              class="group flex items-center gap-3 px-3 py-2.5 rounded-xl transition-all duration-200"
              :class="isActive('/admin/players') ? 'bg-primary-50 text-primary-700' : 'text-neutral-600 hover:bg-neutral-100'"
            >
              <div 
                :class="[
                  'w-9 h-9 rounded-lg flex items-center justify-center transition-colors',
                  isActive('/admin/players') ? 'bg-primary-100' : 'bg-neutral-100 group-hover:bg-neutral-200'
                ]"
              >
                <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
              </div>
              <span class="font-medium">Players</span>
            </NuxtLink>

            <NuxtLink
              to="/admin/academies"
              class="group flex items-center gap-3 px-3 py-2.5 rounded-xl transition-all duration-200"
              :class="isActive('/admin/academies') ? 'bg-primary-50 text-primary-700' : 'text-neutral-600 hover:bg-neutral-100'"
            >
              <div 
                :class="[
                  'w-9 h-9 rounded-lg flex items-center justify-center transition-colors',
                  isActive('/admin/academies') ? 'bg-primary-100' : 'bg-neutral-100 group-hover:bg-neutral-200'
                ]"
              >
                <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
                </svg>
              </div>
              <span class="font-medium">Academies</span>
            </NuxtLink>

            <NuxtLink
              to="/admin/videos"
              class="group flex items-center gap-3 px-3 py-2.5 rounded-xl transition-all duration-200"
              :class="isActive('/admin/videos') ? 'bg-primary-50 text-primary-700' : 'text-neutral-600 hover:bg-neutral-100'"
            >
              <div 
                :class="[
                  'w-9 h-9 rounded-lg flex items-center justify-center transition-colors',
                  isActive('/admin/videos') ? 'bg-primary-100' : 'bg-neutral-100 group-hover:bg-neutral-200'
                ]"
              >
                <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
                </svg>
              </div>
              <span class="font-medium">Videos</span>
            </NuxtLink>

            <NuxtLink
              to="/admin/tournaments"
              class="group flex items-center gap-3 px-3 py-2.5 rounded-xl transition-all duration-200"
              :class="isActive('/admin/tournaments') ? 'bg-primary-50 text-primary-700' : 'text-neutral-600 hover:bg-neutral-100'"
            >
              <div 
                :class="[
                  'w-9 h-9 rounded-lg flex items-center justify-center transition-colors',
                  isActive('/admin/tournaments') ? 'bg-primary-100' : 'bg-neutral-100 group-hover:bg-neutral-200'
                ]"
              >
                <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z" />
                </svg>
              </div>
              <span class="font-medium">Tournaments</span>
            </NuxtLink>

            <!-- Management Section -->
            <p class="px-3 py-2 mt-6 text-xs font-semibold text-neutral-400 uppercase tracking-wider">Management</p>

            <NuxtLink
              to="/admin/users"
              class="group flex items-center gap-3 px-3 py-2.5 rounded-xl transition-all duration-200"
              :class="isActive('/admin/users') ? 'bg-primary-50 text-primary-700' : 'text-neutral-600 hover:bg-neutral-100'"
            >
              <div 
                :class="[
                  'w-9 h-9 rounded-lg flex items-center justify-center transition-colors',
                  isActive('/admin/users') ? 'bg-primary-100' : 'bg-neutral-100 group-hover:bg-neutral-200'
                ]"
              >
                <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13 5.197v1H21v-1a4 4 0 00-4-4h-1m-4-2a4 4 0 110-8 4 4 0 010 8z" />
                </svg>
              </div>
              <span class="font-medium">Users</span>
            </NuxtLink>

            <NuxtLink
              to="/admin/contacts"
              class="group flex items-center gap-3 px-3 py-2.5 rounded-xl transition-all duration-200"
              :class="isActive('/admin/contacts') ? 'bg-primary-50 text-primary-700' : 'text-neutral-600 hover:bg-neutral-100'"
            >
              <div 
                :class="[
                  'w-9 h-9 rounded-lg flex items-center justify-center transition-colors',
                  isActive('/admin/contacts') ? 'bg-primary-100' : 'bg-neutral-100 group-hover:bg-neutral-200'
                ]"
              >
                <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                </svg>
              </div>
              <span class="font-medium">Contact Requests</span>
            </NuxtLink>

            <NuxtLink
              to="/admin/settings"
              class="group flex items-center gap-3 px-3 py-2.5 rounded-xl transition-all duration-200"
              :class="isActive('/admin/settings') ? 'bg-primary-50 text-primary-700' : 'text-neutral-600 hover:bg-neutral-100'"
            >
              <div 
                :class="[
                  'w-9 h-9 rounded-lg flex items-center justify-center transition-colors',
                  isActive('/admin/settings') ? 'bg-primary-100' : 'bg-neutral-100 group-hover:bg-neutral-200'
                ]"
              >
                <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
              </div>
              <span class="font-medium">Settings</span>
            </NuxtLink>
          </nav>

          <!-- Quick Add Button -->
          <div class="p-4 border-t border-neutral-200">
            <NuxtLink to="/admin/players/new">
              <button class="w-full flex items-center justify-center gap-2 px-4 py-3 bg-gradient-to-r from-primary-500 to-primary-600 text-white font-semibold rounded-xl hover:from-primary-600 hover:to-primary-700 shadow-lg shadow-primary-500/25 transition-all duration-200">
                <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
                </svg>
                Add New Player
              </button>
            </NuxtLink>
          </div>
        </div>
      </aside>

      <!-- Main Content -->
      <main class="flex-1 min-h-[calc(100vh-4rem)] p-6 lg:p-8">
        <slot />
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()

const sidebarOpen = ref(false)
const profileOpen = ref(false)

const userInitials = computed(() => {
  const first = authStore.user?.first_name?.charAt(0) || ''
  const last = authStore.user?.last_name?.charAt(0) || ''
  return (first + last).toUpperCase() || 'A'
})

const isActive = (path: string) => {
  if (path === '/admin') {
    return route.path === '/admin'
  }
  return route.path.startsWith(path)
}

const logout = () => {
  authStore.logout()
  router.push('/auth/login')
}

// Close profile dropdown when clicking outside
onMounted(() => {
  document.addEventListener('click', (e) => {
    const target = e.target as HTMLElement
    if (!target.closest('.relative')) {
      profileOpen.value = false
    }
  })
})

// Close sidebar on route change (mobile)
watch(() => route.path, () => {
  sidebarOpen.value = false
})
</script>
