<template>
  <div>
    <!-- Header -->
    <div class="mb-8">
      <h1 class="text-2xl sm:text-3xl font-display font-bold text-neutral-900">
        Welcome back
      </h1>
      <p class="mt-2 text-neutral-600">
        Sign in to continue discovering talent
      </p>
    </div>

    <!-- Form -->
    <form @submit.prevent="handleLogin" class="space-y-5">
      <div>
        <label for="email" class="block text-sm font-medium text-neutral-700 mb-1.5">
          Email address
        </label>
        <input
          id="email"
          v-model="form.email"
          type="email"
          autocomplete="email"
          required
          placeholder="you@example.com"
          :class="[
            'w-full px-4 py-3 rounded-xl border bg-white text-neutral-900 placeholder-neutral-400 transition-all duration-200',
            'focus:outline-none focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500',
            errors.email ? 'border-red-300 bg-red-50' : 'border-neutral-200 hover:border-neutral-300'
          ]"
        />
        <p v-if="errors.email" class="mt-1.5 text-sm text-red-600">{{ errors.email }}</p>
      </div>

      <div>
        <div class="flex items-center justify-between mb-1.5">
          <label for="password" class="block text-sm font-medium text-neutral-700">
            Password
          </label>
          <NuxtLink 
            to="/auth/forgot-password" 
            class="text-sm text-primary-600 hover:text-primary-700 font-medium"
          >
            Forgot?
          </NuxtLink>
        </div>
        <div class="relative">
          <input
            id="password"
            v-model="form.password"
            :type="showPassword ? 'text' : 'password'"
            autocomplete="current-password"
            required
            placeholder="••••••••"
            :class="[
              'w-full px-4 py-3 pr-12 rounded-xl border bg-white text-neutral-900 placeholder-neutral-400 transition-all duration-200',
              'focus:outline-none focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500',
              errors.password ? 'border-red-300 bg-red-50' : 'border-neutral-200 hover:border-neutral-300'
            ]"
          />
          <button
            type="button"
            @click="showPassword = !showPassword"
            class="absolute right-3 top-1/2 -translate-y-1/2 p-1 text-neutral-400 hover:text-neutral-600 transition-colors"
          >
            <svg v-if="showPassword" class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l3.59 3.59m0 0A9.953 9.953 0 0112 5c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21" />
            </svg>
            <svg v-else class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
            </svg>
          </button>
        </div>
        <p v-if="errors.password" class="mt-1.5 text-sm text-red-600">{{ errors.password }}</p>
      </div>

      <!-- Remember me -->
      <label class="flex items-center gap-3 cursor-pointer group">
        <div class="relative">
          <input 
            type="checkbox" 
            v-model="rememberMe"
            class="peer sr-only" 
          />
          <div class="w-5 h-5 rounded border-2 border-neutral-300 bg-white peer-checked:bg-primary-500 peer-checked:border-primary-500 transition-all duration-200 flex items-center justify-center group-hover:border-neutral-400 peer-checked:group-hover:border-primary-600">
            <svg class="w-3 h-3 text-white opacity-0 peer-checked:opacity-100 transition-opacity" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="3">
              <path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7" />
            </svg>
          </div>
        </div>
        <span class="text-sm text-neutral-600 group-hover:text-neutral-900 transition-colors">Remember me for 30 days</span>
      </label>

      <!-- Error message -->
      <div 
        v-if="errors.general" 
        class="flex items-center gap-3 p-4 rounded-xl bg-red-50 border border-red-100"
      >
        <div class="flex-shrink-0 w-10 h-10 rounded-full bg-red-100 flex items-center justify-center">
          <svg class="w-5 h-5 text-red-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
        </div>
        <p class="text-sm text-red-700">{{ errors.general }}</p>
      </div>

      <!-- Submit button -->
      <button
        type="submit"
        :disabled="loading"
        :class="[
          'w-full flex items-center justify-center gap-2 px-6 py-3.5 rounded-xl font-semibold text-white transition-all duration-200',
          loading 
            ? 'bg-primary-400 cursor-not-allowed' 
            : 'bg-primary-600 hover:bg-primary-700 hover:shadow-lg hover:shadow-primary-500/25 active:scale-[0.98]'
        ]"
      >
        <svg v-if="loading" class="w-5 h-5 animate-spin" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
        <span>{{ loading ? 'Signing in...' : 'Sign in' }}</span>
      </button>
    </form>

    <!-- Sign up link -->
    <p class="mt-8 text-center text-sm text-neutral-600">
      Don't have an account?
      <NuxtLink 
        to="/auth/register" 
        class="text-primary-600 hover:text-primary-700 font-semibold ml-1"
      >
        Create one free →
      </NuxtLink>
    </p>
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  layout: 'auth',
  middleware: 'guest',
})

const authStore = useAuthStore()
const router = useRouter()
const route = useRoute()

const loading = ref(false)
const showPassword = ref(false)
const rememberMe = ref(false)

const form = reactive({
  email: '',
  password: '',
})

const errors = reactive({
  email: '',
  password: '',
  general: '',
})

const handleLogin = async () => {
  // Reset errors
  errors.email = ''
  errors.password = ''
  errors.general = ''

  // Basic validation
  if (!form.email) {
    errors.email = 'Email is required'
    return
  }
  if (!form.password) {
    errors.password = 'Password is required'
    return
  }

  loading.value = true
  try {
    const success = await authStore.login(form.email, form.password)
    if (success) {
      // Fetch subscription status
      const subStore = useSubscriptionStore()
      await subStore.fetchSubscription()
      
      // Redirect to intended page or dashboard
      const redirect = route.query.redirect as string || '/dashboard'
      router.push(redirect)
    } else {
      errors.general = 'Invalid email or password. Please try again.'
    }
  } catch (error) {
    errors.general = 'Something went wrong. Please try again later.'
  } finally {
    loading.value = false
  }
}
</script>
