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
        <nav class="hidden md:flex items-center space-x-8">
          <NuxtLink 
            to="/discover" 
            :class="[
              'font-medium transition-colors relative group',
              isScrolled || !isHomePage 
                ? 'text-neutral-600 hover:text-primary-600' 
                : 'text-neutral-300 hover:text-white'
            ]"
          >
            Discover
            <span class="absolute -bottom-1 left-0 w-0 h-0.5 bg-primary-500 transition-all group-hover:w-full"></span>
          </NuxtLink>
          <NuxtLink 
            to="/pricing" 
            :class="[
              'font-medium transition-colors relative group',
              isScrolled || !isHomePage 
                ? 'text-neutral-600 hover:text-primary-600' 
                : 'text-neutral-300 hover:text-white'
            ]"
          >
            Pricing
            <span class="absolute -bottom-1 left-0 w-0 h-0.5 bg-primary-500 transition-all group-hover:w-full"></span>
          </NuxtLink>
        </nav>

        <!-- Auth Buttons -->
        <div class="hidden md:flex items-center space-x-4">
          <template v-if="authStore.isAuthenticated">
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
            <NuxtLink 
              to="/dashboard" 
              :class="[
                'font-medium transition-colors',
                isScrolled || !isHomePage 
                  ? 'text-neutral-600 hover:text-primary-600' 
                  : 'text-neutral-300 hover:text-white'
              ]"
            >
              Dashboard
            </NuxtLink>
            <button 
              @click="handleLogout" 
              :class="[
                'font-medium transition-colors',
                isScrolled || !isHomePage 
                  ? 'text-neutral-600 hover:text-red-600' 
                  : 'text-neutral-300 hover:text-red-400'
              ]"
            >
              Logout
            </button>
          </template>
          <template v-else>
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
                  'px-5 py-2.5 rounded-lg font-semibold text-sm transition-all duration-300',
                  isScrolled || !isHomePage 
                    ? 'bg-primary-600 text-white hover:bg-primary-700 shadow-lg shadow-primary-600/20' 
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
        enter-active-class="transition-all duration-300 ease-out"
        enter-from-class="opacity-0 -translate-y-2"
        enter-to-class="opacity-100 translate-y-0"
        leave-active-class="transition-all duration-200 ease-in"
        leave-from-class="opacity-100 translate-y-0"
        leave-to-class="opacity-0 -translate-y-2"
      >
        <div 
          v-if="mobileMenuOpen" 
          class="md:hidden fixed inset-0 top-16 z-50 bg-neutral-900"
        >
          <nav class="flex flex-col space-y-2 p-6">
            <NuxtLink 
              to="/discover" 
              class="font-medium text-lg text-white hover:text-primary-400 transition-colors py-3"
              @click="mobileMenuOpen = false"
            >
              Discover
            </NuxtLink>
            
            <!-- Live status badge -->
            <div class="py-2">
              <div class="inline-flex items-center gap-2 px-4 py-2 rounded-full bg-neutral-800 border border-neutral-700">
                <span class="relative flex h-2.5 w-2.5">
                  <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-primary-400 opacity-75"></span>
                  <span class="relative inline-flex rounded-full h-2.5 w-2.5 bg-primary-500"></span>
                </span>
                <span class="text-sm text-neutral-300">Now scouting in 15+ African countries</span>
              </div>
            </div>
            
            <NuxtLink 
              to="/pricing" 
              class="font-medium text-lg text-white hover:text-primary-400 transition-colors py-3"
              @click="mobileMenuOpen = false"
            >
              Pricing
            </NuxtLink>
            <template v-if="authStore.isAuthenticated">
              <NuxtLink 
                v-if="authStore.isAdmin"
                to="/admin" 
                class="font-medium text-lg text-white hover:text-primary-400 transition-colors py-3"
                @click="mobileMenuOpen = false"
              >
                Admin
              </NuxtLink>
              <NuxtLink 
                to="/dashboard" 
                class="font-medium text-lg text-white hover:text-primary-400 transition-colors py-3"
                @click="mobileMenuOpen = false"
              >
                Dashboard
              </NuxtLink>
              <button 
                @click="handleLogout; mobileMenuOpen = false" 
                class="font-medium text-lg text-white hover:text-red-400 transition-colors text-left py-3"
              >
                Logout
              </button>
            </template>
            <template v-else>
              <NuxtLink 
                to="/auth/login" 
                class="font-medium text-lg text-white hover:text-primary-400 transition-colors py-3"
                @click="mobileMenuOpen = false"
              >
                Login
              </NuxtLink>
              <NuxtLink to="/auth/register" @click="mobileMenuOpen = false">
                <button class="w-full px-5 py-3 bg-primary-600 text-white rounded-xl font-semibold text-sm hover:bg-primary-700 transition-colors">
                  Get Started
                </button>
              </NuxtLink>
            </template>
          </nav>
        </div>
      </Transition>
    </div>
  </header>
</template>

<script setup lang="ts">
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
