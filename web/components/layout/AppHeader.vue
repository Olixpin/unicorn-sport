<template>
  <header 
    :class="[
      'fixed top-0 left-0 right-0 z-50 transition-all duration-300',
      mobileMenuOpen 
        ? 'bg-neutral-900 border-b border-neutral-800'
        : isScrolled || !isHomePage 
          ? 'bg-white/95 backdrop-blur-md border-b border-neutral-200 shadow-sm' 
          : 'bg-transparent'
    ]"
  >
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex items-center justify-between h-16 lg:h-20">
        <!-- Logo -->
        <NuxtLink to="/" class="flex items-center space-x-2 group">
          <span class="text-2xl group-hover:scale-110 transition-transform">🦄</span>
          <span 
            :class="[
              'font-display font-bold text-xl transition-colors',
              mobileMenuOpen ? 'text-white' : isScrolled || !isHomePage ? 'text-primary-600' : 'text-white'
            ]"
          >
            Unicorn Sport
          </span>
        </NuxtLink>

        <!-- Desktop Navigation -->
        <nav class="hidden md:flex items-center space-x-6">
          <NuxtLink 
            to="/discover" 
            :class="[
              'font-medium transition-colors',
              isScrolled || !isHomePage 
                ? 'text-neutral-600 hover:text-primary-600' 
                : 'text-neutral-300 hover:text-white'
            ]"
          >
            Discover
          </NuxtLink>
          <NuxtLink 
            to="/dashboard" 
            :class="[
              'flex items-center gap-1.5 font-medium transition-colors',
              isScrolled || !isHomePage 
                ? 'text-neutral-600 hover:text-primary-600' 
                : 'text-neutral-300 hover:text-white'
            ]"
          >
            <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM9.555 7.168A1 1 0 008 8v4a1 1 0 001.555.832l3-2a1 1 0 000-1.664l-3-2z" clip-rule="evenodd"/>
            </svg>
            Feed
          </NuxtLink>
        </nav>

        <!-- Auth Buttons -->
        <div class="hidden md:flex items-center space-x-4">
          <template v-if="authStore.isAuthenticated">
            <NuxtLink 
              to="/dashboard/saved" 
              :class="[
                'font-medium transition-colors',
                isScrolled || !isHomePage 
                  ? 'text-neutral-600 hover:text-primary-600' 
                  : 'text-neutral-300 hover:text-white'
              ]"
            >
              Saved
            </NuxtLink>
            <NuxtLink 
              v-if="authStore.isAdmin" 
              to="/admin" 
              :class="[
                'font-medium transition-colors',
                isScrolled || !isHomePage 
                  ? 'text-neutral-600 hover:text-primary-600' 
                  : 'text-neutral-300 hover:text-white'
              ]"
            >
              Admin
            </NuxtLink>
            <button 
              @click="handleLogout" 
              :class="[
                'font-medium transition-colors',
                isScrolled || !isHomePage 
                  ? 'text-neutral-500 hover:text-neutral-700' 
                  : 'text-neutral-400 hover:text-white'
              ]"
            >
              Logout
            </button>
          </template>
          <template v-else>
            <NuxtLink 
              to="/pricing" 
              :class="[
                'font-medium transition-colors',
                isScrolled || !isHomePage 
                  ? 'text-neutral-600 hover:text-primary-600' 
                  : 'text-neutral-300 hover:text-white'
              ]"
            >
              Pricing
            </NuxtLink>
            <NuxtLink 
              to="/auth/login" 
              :class="[
                'font-medium transition-colors',
                isScrolled || !isHomePage 
                  ? 'text-neutral-600 hover:text-primary-600' 
                  : 'text-neutral-300 hover:text-white'
              ]"
            >
              Login
            </NuxtLink>
            <NuxtLink to="/auth/register">
              <button 
                :class="[
                  'px-4 py-2 rounded-lg font-semibold text-sm transition-all',
                  isScrolled || !isHomePage 
                    ? 'bg-primary-600 text-white hover:bg-primary-700' 
                    : 'bg-white text-neutral-900 hover:bg-neutral-100'
                ]"
              >
                Get Started
              </button>
            </NuxtLink>
          </template>
        </div>

        <!-- Mobile menu button -->
        <button 
          class="md:hidden p-2 rounded-lg transition-colors"
          :class="mobileMenuOpen ? 'text-white hover:bg-white/10' : isScrolled || !isHomePage ? 'text-neutral-600 hover:bg-neutral-100' : 'text-white hover:bg-white/10'"
          @click="mobileMenuOpen = !mobileMenuOpen"
        >
          <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path 
              v-if="!mobileMenuOpen"
              stroke-linecap="round" 
              stroke-linejoin="round" 
              stroke-width="2" 
              d="M4 6h16M4 12h16M4 18h16" 
            />
            <path 
              v-else
              stroke-linecap="round" 
              stroke-linejoin="round" 
              stroke-width="2" 
              d="M6 18L18 6M6 6l12 12" 
            />
          </svg>
        </button>
      </div>

      <!-- Mobile Navigation -->
      <Transition
        enter-active-class="transition-all duration-300 ease-[cubic-bezier(0.32,0.72,0,1)]"
        enter-from-class="opacity-0 max-h-0"
        enter-to-class="opacity-100 max-h-screen"
        leave-active-class="transition-all duration-200 ease-[cubic-bezier(0.32,0.72,0,1)]"
        leave-from-class="opacity-100 max-h-screen"
        leave-to-class="opacity-0 max-h-0"
      >
        <div 
          v-if="mobileMenuOpen" 
          class="md:hidden fixed inset-0 top-16 z-50 bg-neutral-900 overflow-hidden"
        >
          <nav class="flex flex-col p-6 space-y-1">
            <!-- Primary Actions -->
            <NuxtLink 
              to="/discover" 
              class="flex items-center gap-3 text-white hover:text-primary-400 transition-all py-3 px-3 rounded-lg hover:bg-white/5 transform transition-transform duration-300 delay-[50ms]"
              :class="mobileMenuOpen ? 'translate-x-0 opacity-100' : '-translate-x-4 opacity-0'"
              @click="mobileMenuOpen = false"
            >
              <svg class="w-5 h-5 text-neutral-500" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
              Discover Players
            </NuxtLink>
            
            <NuxtLink 
              to="/dashboard" 
              class="flex items-center gap-3 text-white hover:text-primary-400 transition-all py-3 px-3 rounded-lg hover:bg-white/5 transform transition-transform duration-300 delay-[100ms]"
              :class="mobileMenuOpen ? 'translate-x-0 opacity-100' : '-translate-x-4 opacity-0'"
              @click="mobileMenuOpen = false"
            >
              <svg class="w-5 h-5 text-neutral-500" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM9.555 7.168A1 1 0 008 8v4a1 1 0 001.555.832l3-2a1 1 0 000-1.664l-3-2z" clip-rule="evenodd"/>
              </svg>
              Watch Feed
            </NuxtLink>

            <template v-if="authStore.isAuthenticated">
              <NuxtLink 
                to="/dashboard/saved" 
                class="flex items-center gap-3 text-white hover:text-primary-400 transition-all py-3 px-3 rounded-lg hover:bg-white/5 transform transition-transform duration-300 delay-[150ms]"
                :class="mobileMenuOpen ? 'translate-x-0 opacity-100' : '-translate-x-4 opacity-0'"
                @click="mobileMenuOpen = false"
              >
                <svg class="w-5 h-5 text-neutral-500" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z" />
                </svg>
                Saved Players
              </NuxtLink>

              <NuxtLink 
                v-if="authStore.isAdmin"
                to="/admin" 
                class="flex items-center gap-3 text-white hover:text-primary-400 transition-all py-3 px-3 rounded-lg hover:bg-white/5 transform transition-transform duration-300 delay-[200ms]"
                :class="mobileMenuOpen ? 'translate-x-0 opacity-100' : '-translate-x-4 opacity-0'"
                @click="mobileMenuOpen = false"
              >
                <svg class="w-5 h-5 text-neutral-500" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                  <path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
                Admin
              </NuxtLink>

              <!-- Divider -->
              <div class="border-t border-neutral-800 my-3"></div>

              <!-- Logout -->
              <button 
                @click="() => { handleLogout(); mobileMenuOpen = false }" 
                class="flex items-center gap-3 text-neutral-400 hover:text-white transition-all py-3 px-3 rounded-lg hover:bg-white/5 text-left transform transition-transform duration-300 delay-[250ms]"
                :class="mobileMenuOpen ? 'translate-x-0 opacity-100' : '-translate-x-4 opacity-0'"
              >
                <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
                </svg>
                Logout
              </button>
            </template>

            <template v-else>
              <NuxtLink 
                to="/pricing" 
                class="flex items-center gap-3 text-white hover:text-primary-400 transition-all py-3 px-3 rounded-lg hover:bg-white/5 transform transition-transform duration-300 delay-[150ms]"
                :class="mobileMenuOpen ? 'translate-x-0 opacity-100' : '-translate-x-4 opacity-0'"
                @click="mobileMenuOpen = false"
              >
                <svg class="w-5 h-5 text-neutral-500" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                Pricing
              </NuxtLink>

              <div 
                class="mt-6 space-y-3 transform transition-all duration-300 delay-[200ms]"
                :class="mobileMenuOpen ? 'translate-y-0 opacity-100' : 'translate-y-4 opacity-0'"
              >
                <NuxtLink 
                  to="/auth/login" 
                  class="block w-full text-center py-3 text-white border border-neutral-700 rounded-lg font-medium hover:bg-neutral-800 transition-colors"
                  @click="mobileMenuOpen = false"
                >
                  Login
                </NuxtLink>
                <NuxtLink 
                  to="/auth/register" 
                  class="block w-full text-center py-3 bg-primary-600 text-white rounded-lg font-medium hover:bg-primary-700 transition-colors"
                  @click="mobileMenuOpen = false"
                >
                  Get Started
                </NuxtLink>
              </div>
            </template>
          </nav>
        </div>
      </Transition>
    </div>
  </header>
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted, onUnmounted } from 'vue'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()
const mobileMenuOpen = ref(false)
const isScrolled = ref(false)

const isHomePage = computed(() => route.path === '/')

// Close mobile menu on route change
watch(() => route.path, () => {
  mobileMenuOpen.value = false
})

const handleLogout = () => {
  authStore.logout()
  navigateTo('/')
}

onMounted(() => {
  const handleScroll = () => {
    isScrolled.value = window.scrollY > 50
  }
  
  window.addEventListener('scroll', handleScroll)
  handleScroll()
  
  onUnmounted(() => {
    window.removeEventListener('scroll', handleScroll)
  })
})
</script>
