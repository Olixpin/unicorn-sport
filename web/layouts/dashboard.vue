<template>
  <div class="min-h-screen bg-black lg:bg-neutral-100">
    <div class="flex">
      <!-- Sidebar -->
      <aside class="hidden lg:flex lg:flex-col w-72 bg-white border-r border-neutral-200 min-h-screen fixed left-0 top-0 bottom-0">
        <!-- Logo Header -->
        <div class="p-4 border-b border-neutral-100">
          <NuxtLink to="/" class="flex items-center gap-3 px-3 py-2 rounded-xl hover:bg-neutral-50 transition-colors group">
            <div class="w-10 h-10 rounded-xl bg-gradient-to-br from-primary-500 to-emerald-500 flex items-center justify-center shadow-lg shadow-primary-500/20 group-hover:scale-105 transition-transform">
              <span class="text-xl">🦄</span>
            </div>
            <div>
              <span class="font-display font-bold text-lg text-neutral-900">Unicorn Sport</span>
              <p class="text-xs text-neutral-400">Scout Dashboard</p>
            </div>
          </NuxtLink>
        </div>

        <nav class="flex-1 p-4 space-y-1 overflow-y-auto">
          <NuxtLink 
            to="/dashboard" 
            class="group flex items-center px-4 py-3 text-neutral-600 rounded-xl hover:bg-neutral-50 transition-colors"
            active-class="!bg-primary-50 !text-primary-700"
          >
            <div class="w-10 h-10 rounded-xl bg-neutral-100 group-hover:bg-neutral-200 flex items-center justify-center mr-3 transition-colors group-[.router-link-active]:bg-primary-100">
              <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
              </svg>
            </div>
            <div>
              <span class="font-medium">Dashboard</span>
              <p class="text-xs text-neutral-400">Overview & activity</p>
            </div>
          </NuxtLink>

          <NuxtLink 
            to="/discover" 
            class="group flex items-center px-4 py-3 text-neutral-600 rounded-xl hover:bg-neutral-50 transition-colors"
            active-class="!bg-primary-50 !text-primary-700"
          >
            <div class="w-10 h-10 rounded-xl bg-neutral-100 group-hover:bg-neutral-200 flex items-center justify-center mr-3 transition-colors group-[.router-link-active]:bg-primary-100">
              <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </div>
            <div>
              <span class="font-medium">Discover</span>
              <p class="text-xs text-neutral-400">Find new talent</p>
            </div>
          </NuxtLink>

          <NuxtLink 
            to="/dashboard/saved" 
            class="group flex items-center px-4 py-3 text-neutral-600 rounded-xl hover:bg-neutral-50 transition-colors"
            active-class="!bg-primary-50 !text-primary-700"
          >
            <div class="w-10 h-10 rounded-xl bg-neutral-100 group-hover:bg-neutral-200 flex items-center justify-center mr-3 transition-colors group-[.router-link-active]:bg-primary-100">
              <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
              </svg>
            </div>
            <div>
              <span class="font-medium">Saved Players</span>
              <p class="text-xs text-neutral-400">Your shortlist</p>
            </div>
          </NuxtLink>

          <NuxtLink 
            to="/dashboard/contacts" 
            class="group flex items-center px-4 py-3 text-neutral-600 rounded-xl hover:bg-neutral-50 transition-colors"
            active-class="!bg-primary-50 !text-primary-700"
          >
            <div class="w-10 h-10 rounded-xl bg-neutral-100 group-hover:bg-neutral-200 flex items-center justify-center mr-3 transition-colors group-[.router-link-active]:bg-primary-100">
              <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
              </svg>
            </div>
            <div>
              <span class="font-medium">Contact Requests</span>
              <p class="text-xs text-neutral-400">Player inquiries</p>
            </div>
          </NuxtLink>

          <NuxtLink 
            to="/dashboard/settings" 
            class="group flex items-center px-4 py-3 text-neutral-600 rounded-xl hover:bg-neutral-50 transition-colors"
            active-class="!bg-primary-50 !text-primary-700"
          >
            <div class="w-10 h-10 rounded-xl bg-neutral-100 group-hover:bg-neutral-200 flex items-center justify-center mr-3 transition-colors group-[.router-link-active]:bg-primary-100">
              <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
              </svg>
            </div>
            <div>
              <span class="font-medium">Settings</span>
              <p class="text-xs text-neutral-400">Account & billing</p>
            </div>
          </NuxtLink>
        </nav>

        <!-- Subscription Status -->
        <div class="p-4 border-t border-neutral-200">
          <div class="relative rounded-2xl p-5 overflow-hidden" :class="subStore.tier === 'free' ? 'bg-neutral-900' : 'bg-gradient-to-br from-primary-600 to-primary-700'">
            <div v-if="subStore.tier !== 'free'" class="absolute top-0 right-0 w-32 h-32 bg-white/10 rounded-full -translate-y-1/2 translate-x-1/2"></div>
            <div class="relative">
              <div class="flex items-center justify-between mb-3">
                <span class="text-sm text-white/70">Current Plan</span>
                <span class="px-2.5 py-1 bg-white/20 text-white text-xs font-bold rounded-full">
                  {{ subStore.tier.toUpperCase() }}
                </span>
              </div>
              <p class="text-white font-semibold mb-3">
                {{ subStore.tier === 'free' ? 'Upgrade for more' : 'Premium access' }}
              </p>
              <NuxtLink v-if="subStore.tier === 'free'" to="/pricing">
                <button class="w-full py-2.5 bg-white text-neutral-900 font-semibold rounded-xl hover:bg-neutral-100 transition-colors text-sm">
                  Upgrade Now
                </button>
              </NuxtLink>
              <NuxtLink v-else to="/dashboard/settings">
                <button class="w-full py-2.5 bg-white/20 text-white font-medium rounded-xl hover:bg-white/30 transition-colors text-sm">
                  Manage Plan
                </button>
              </NuxtLink>
            </div>
          </div>
        </div>
      </aside>

      <!-- Main Content -->
      <main class="flex-1 lg:ml-72 min-h-screen pb-16 lg:pb-0">
        <!-- Top Bar (Mobile Logo + Menu) - Hidden on dashboard for TikTok feel -->
        <div v-if="!isFeedPage" class="lg:hidden sticky top-0 z-40 bg-white border-b border-neutral-200 px-4 py-3">
          <div class="flex items-center justify-between">
            <NuxtLink to="/" class="flex items-center gap-2">
              <div class="w-8 h-8 rounded-lg bg-gradient-to-br from-primary-500 to-emerald-500 flex items-center justify-center">
                <span class="text-sm">🦄</span>
              </div>
              <span class="font-display font-bold text-neutral-900">Unicorn Sport</span>
            </NuxtLink>
            <button 
              @click="mobileMenuOpen = !mobileMenuOpen"
              class="p-2 rounded-lg text-neutral-600 hover:bg-neutral-100 transition-colors"
            >
              <svg class="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
              </svg>
            </button>
          </div>
        </div>

        <!-- Mobile Menu Overlay -->
        <Transition
          enter-active-class="transition-opacity duration-200"
          enter-from-class="opacity-0"
          enter-to-class="opacity-100"
          leave-active-class="transition-opacity duration-200"
          leave-from-class="opacity-100"
          leave-to-class="opacity-0"
        >
          <div v-if="mobileMenuOpen" class="lg:hidden fixed inset-0 z-50 bg-black/50" @click="mobileMenuOpen = false"></div>
        </Transition>

        <!-- Mobile Sidebar -->
        <Transition
          enter-active-class="transition-transform duration-300"
          enter-from-class="-translate-x-full"
          enter-to-class="translate-x-0"
          leave-active-class="transition-transform duration-300"
          leave-from-class="translate-x-0"
          leave-to-class="-translate-x-full"
        >
          <div v-if="mobileMenuOpen" class="lg:hidden fixed inset-y-0 left-0 z-50 w-72 bg-white shadow-xl">
            <div class="p-4 border-b border-neutral-100 flex items-center justify-between">
              <NuxtLink to="/" class="flex items-center gap-2" @click="mobileMenuOpen = false">
                <div class="w-8 h-8 rounded-lg bg-gradient-to-br from-primary-500 to-emerald-500 flex items-center justify-center">
                  <span class="text-sm">🦄</span>
                </div>
                <span class="font-display font-bold text-neutral-900">Unicorn Sport</span>
              </NuxtLink>
              <button @click="mobileMenuOpen = false" class="p-2 rounded-lg text-neutral-400 hover:text-neutral-600 hover:bg-neutral-100">
                <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            </div>
            <nav class="p-4 space-y-1">
              <NuxtLink 
                to="/dashboard" 
                class="flex items-center gap-3 px-4 py-3 text-neutral-600 rounded-xl hover:bg-neutral-50 transition-colors"
                active-class="!bg-primary-50 !text-primary-700"
                @click="mobileMenuOpen = false"
              >
                <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
                </svg>
                <span class="font-medium">Dashboard</span>
              </NuxtLink>
              <NuxtLink 
                to="/discover" 
                class="flex items-center gap-3 px-4 py-3 text-neutral-600 rounded-xl hover:bg-neutral-50 transition-colors"
                active-class="!bg-primary-50 !text-primary-700"
                @click="mobileMenuOpen = false"
              >
                <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                </svg>
                <span class="font-medium">Discover</span>
              </NuxtLink>
              <NuxtLink 
                to="/dashboard/saved" 
                class="flex items-center gap-3 px-4 py-3 text-neutral-600 rounded-xl hover:bg-neutral-50 transition-colors"
                active-class="!bg-primary-50 !text-primary-700"
                @click="mobileMenuOpen = false"
              >
                <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
                </svg>
                <span class="font-medium">Saved Players</span>
              </NuxtLink>
              <NuxtLink 
                to="/dashboard/contacts" 
                class="flex items-center gap-3 px-4 py-3 text-neutral-600 rounded-xl hover:bg-neutral-50 transition-colors"
                active-class="!bg-primary-50 !text-primary-700"
                @click="mobileMenuOpen = false"
              >
                <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                </svg>
                <span class="font-medium">Contacts</span>
              </NuxtLink>
              <NuxtLink 
                to="/dashboard/settings" 
                class="flex items-center gap-3 px-4 py-3 text-neutral-600 rounded-xl hover:bg-neutral-50 transition-colors"
                active-class="!bg-primary-50 !text-primary-700"
                @click="mobileMenuOpen = false"
              >
                <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
                <span class="font-medium">Settings</span>
              </NuxtLink>
            </nav>
          </div>
        </Transition>

        <!-- Page Content -->
        <div class="lg:p-8" :class="isFeedPage ? 'p-0' : 'p-4 sm:p-4'">
          <slot />
        </div>
      </main>
    </div>

    <!-- Bottom Navigation (Mobile Only) - TikTok Style -->
    <nav class="lg:hidden fixed bottom-0 left-0 right-0 z-50 bg-black/95 backdrop-blur-lg border-t border-white/10 safe-area-pb">
      <div class="flex items-center justify-around h-14">
        <!-- Home -->
        <NuxtLink 
          to="/dashboard" 
          class="flex flex-col items-center justify-center flex-1 h-full"
          :class="route.path === '/dashboard' ? 'text-white' : 'text-white/50'"
        >
          <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
            <path d="M10 20v-6h4v6h5v-8h3L12 3 2 12h3v8z"/>
          </svg>
          <span class="text-[10px] mt-0.5 font-medium">Home</span>
        </NuxtLink>

        <!-- Discover -->
        <NuxtLink 
          to="/discover" 
          class="flex flex-col items-center justify-center flex-1 h-full"
          :class="route.path === '/discover' ? 'text-white' : 'text-white/50'"
        >
          <svg class="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
          <span class="text-[10px] mt-0.5 font-medium">Discover</span>
        </NuxtLink>

        <!-- Saved -->
        <NuxtLink 
          to="/dashboard/saved" 
          class="flex flex-col items-center justify-center flex-1 h-full"
          :class="route.path === '/dashboard/saved' ? 'text-white' : 'text-white/50'"
        >
          <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
            <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/>
          </svg>
          <span class="text-[10px] mt-0.5 font-medium">Saved</span>
        </NuxtLink>

        <!-- Menu / More -->
        <button 
          @click="mobileMenuOpen = !mobileMenuOpen"
          class="flex flex-col items-center justify-center flex-1 h-full text-white/50"
        >
          <svg class="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M4 6h16M4 12h16M4 18h16" />
          </svg>
          <span class="text-[10px] mt-0.5 font-medium">Menu</span>
        </button>
      </div>
    </nav>
  </div>
</template>

<script setup lang="ts">
const route = useRoute()
const subStore = useSubscriptionStore()
const mobileMenuOpen = ref(false)

// Check if we're on the main feed page for TikTok-style layout
const isFeedPage = computed(() => route.path === '/dashboard')
</script>

<style scoped>
/* Safe area padding for devices with home indicator */
.safe-area-pb {
  padding-bottom: env(safe-area-inset-bottom, 0px);
}
</style>
