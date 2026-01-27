<template>
  <div>
    <!-- Header -->
    <div class="mb-8">
      <h1 class="text-2xl sm:text-3xl font-display font-bold text-neutral-900">
        Create your account
      </h1>
      <p class="mt-2 text-neutral-600">
        Start discovering African football talent today
      </p>
    </div>

    <!-- Form -->
    <form @submit.prevent="handleRegister" class="space-y-5">
      <!-- Name fields -->
      <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <div>
          <label for="first_name" class="block text-sm font-medium text-neutral-700 mb-1.5">
            First name
          </label>
          <input
            id="first_name"
            v-model="form.first_name"
            type="text"
            autocomplete="given-name"
            required
            placeholder="John"
            :class="[
              'w-full px-4 py-3 rounded-xl border bg-white text-neutral-900 placeholder-neutral-400 transition-all duration-200',
              'focus:outline-none focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500',
              errors.first_name ? 'border-red-300 bg-red-50' : 'border-neutral-200 hover:border-neutral-300'
            ]"
          />
          <p v-if="errors.first_name" class="mt-1.5 text-sm text-red-600">{{ errors.first_name }}</p>
        </div>
        <div>
          <label for="last_name" class="block text-sm font-medium text-neutral-700 mb-1.5">
            Last name
          </label>
          <input
            id="last_name"
            v-model="form.last_name"
            type="text"
            autocomplete="family-name"
            required
            placeholder="Doe"
            :class="[
              'w-full px-4 py-3 rounded-xl border bg-white text-neutral-900 placeholder-neutral-400 transition-all duration-200',
              'focus:outline-none focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500',
              errors.last_name ? 'border-red-300 bg-red-50' : 'border-neutral-200 hover:border-neutral-300'
            ]"
          />
          <p v-if="errors.last_name" class="mt-1.5 text-sm text-red-600">{{ errors.last_name }}</p>
        </div>
      </div>

      <!-- Email -->
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

      <!-- Password -->
      <div>
        <label for="password" class="block text-sm font-medium text-neutral-700 mb-1.5">
          Password
        </label>
        <div class="relative">
          <input
            id="password"
            v-model="form.password"
            :type="showPassword ? 'text' : 'password'"
            autocomplete="new-password"
            required
            placeholder="Minimum 8 characters"
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
        <!-- Password strength indicator -->
        <div v-if="form.password" class="mt-2 flex items-center gap-2">
          <div class="flex-1 h-1 rounded-full bg-neutral-200 overflow-hidden">
            <div 
              :class="[
                'h-full rounded-full transition-all duration-300',
                passwordStrength.color
              ]"
              :style="{ width: passwordStrength.width }"
            ></div>
          </div>
          <span :class="['text-xs font-medium', passwordStrength.textColor]">
            {{ passwordStrength.label }}
          </span>
        </div>
      </div>

      <!-- User type selection -->
      <div>
        <label class="block text-sm font-medium text-neutral-700 mb-3">
          I am a
        </label>
        <div class="grid grid-cols-2 gap-3">
          <button
            type="button"
            :class="[
              'relative p-4 rounded-xl border-2 text-center transition-all duration-200',
              form.user_type === 'scout' 
                ? 'border-primary-500 bg-primary-50 ring-2 ring-primary-500/20' 
                : 'border-neutral-200 hover:border-neutral-300 hover:bg-neutral-50'
            ]"
            @click="form.user_type = 'scout'"
          >
            <div 
              v-if="form.user_type === 'scout'"
              class="absolute top-2 right-2 w-5 h-5 rounded-full bg-primary-500 flex items-center justify-center"
            >
              <svg class="w-3 h-3 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="3">
                <path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7" />
              </svg>
            </div>
            <div class="w-12 h-12 mx-auto mb-3 rounded-xl bg-gradient-to-br from-primary-500 to-emerald-500 flex items-center justify-center">
              <svg class="w-6 h-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </div>
            <span :class="['font-semibold', form.user_type === 'scout' ? 'text-primary-700' : 'text-neutral-700']">Scout</span>
            <p class="mt-1 text-xs text-neutral-500">Find players</p>
          </button>
          
          <button
            type="button"
            :class="[
              'relative p-4 rounded-xl border-2 text-center transition-all duration-200',
              form.user_type === 'agency' 
                ? 'border-primary-500 bg-primary-50 ring-2 ring-primary-500/20' 
                : 'border-neutral-200 hover:border-neutral-300 hover:bg-neutral-50'
            ]"
            @click="form.user_type = 'agency'"
          >
            <div 
              v-if="form.user_type === 'agency'"
              class="absolute top-2 right-2 w-5 h-5 rounded-full bg-primary-500 flex items-center justify-center"
            >
              <svg class="w-3 h-3 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="3">
                <path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7" />
              </svg>
            </div>
            <div class="w-12 h-12 mx-auto mb-3 rounded-xl bg-gradient-to-br from-secondary-500 to-orange-500 flex items-center justify-center">
              <svg class="w-6 h-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
              </svg>
            </div>
            <span :class="['font-semibold', form.user_type === 'agency' ? 'text-primary-700' : 'text-neutral-700']">Agency / Club</span>
            <p class="mt-1 text-xs text-neutral-500">Recruit talent</p>
          </button>
        </div>
        <p v-if="errors.user_type" class="mt-2 text-sm text-red-600">{{ errors.user_type }}</p>
      </div>

      <!-- Terms -->
      <label class="flex items-start gap-3 cursor-pointer group">
        <div class="relative mt-0.5">
          <input 
            type="checkbox" 
            v-model="acceptTerms"
            class="peer sr-only" 
          />
          <div class="w-5 h-5 rounded border-2 border-neutral-300 bg-white peer-checked:bg-primary-500 peer-checked:border-primary-500 transition-all duration-200 flex items-center justify-center group-hover:border-neutral-400 peer-checked:group-hover:border-primary-600">
            <svg class="w-3 h-3 text-white opacity-0 peer-checked:opacity-100 transition-opacity" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="3">
              <path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7" />
            </svg>
          </div>
        </div>
        <span class="text-sm text-neutral-600 group-hover:text-neutral-900 transition-colors">
          I agree to the 
          <NuxtLink to="/terms" class="text-primary-600 hover:underline">Terms of Service</NuxtLink> 
          and 
          <NuxtLink to="/privacy" class="text-primary-600 hover:underline">Privacy Policy</NuxtLink>
        </span>
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
        :disabled="loading || !acceptTerms"
        :class="[
          'w-full flex items-center justify-center gap-2 px-6 py-3.5 rounded-xl font-semibold text-white transition-all duration-200',
          loading || !acceptTerms
            ? 'bg-primary-300 cursor-not-allowed' 
            : 'bg-primary-600 hover:bg-primary-700 hover:shadow-lg hover:shadow-primary-500/25 active:scale-[0.98]'
        ]"
      >
        <svg v-if="loading" class="w-5 h-5 animate-spin" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
        <span>{{ loading ? 'Creating account...' : 'Create account' }}</span>
      </button>
    </form>

    <!-- Sign in link -->
    <p class="mt-8 text-center text-sm text-neutral-600">
      Already have an account?
      <NuxtLink 
        to="/auth/login" 
        class="text-primary-600 hover:text-primary-700 font-semibold ml-1"
      >
        Sign in →
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

const loading = ref(false)
const showPassword = ref(false)
const acceptTerms = ref(false)

const form = reactive({
  first_name: '',
  last_name: '',
  email: '',
  password: '',
  user_type: 'scout' as 'scout' | 'agency',
})

const errors = reactive({
  first_name: '',
  last_name: '',
  email: '',
  password: '',
  user_type: '',
  general: '',
})

// Password strength calculation
const passwordStrength = computed(() => {
  const password = form.password
  if (!password) return { width: '0%', color: 'bg-neutral-200', textColor: 'text-neutral-400', label: '' }
  
  let score = 0
  if (password.length >= 8) score++
  if (password.length >= 12) score++
  if (/[a-z]/.test(password) && /[A-Z]/.test(password)) score++
  if (/\d/.test(password)) score++
  if (/[^a-zA-Z0-9]/.test(password)) score++
  
  if (score <= 1) return { width: '20%', color: 'bg-red-500', textColor: 'text-red-600', label: 'Weak' }
  if (score === 2) return { width: '40%', color: 'bg-orange-500', textColor: 'text-orange-600', label: 'Fair' }
  if (score === 3) return { width: '60%', color: 'bg-yellow-500', textColor: 'text-yellow-600', label: 'Good' }
  if (score === 4) return { width: '80%', color: 'bg-primary-500', textColor: 'text-primary-600', label: 'Strong' }
  return { width: '100%', color: 'bg-emerald-500', textColor: 'text-emerald-600', label: 'Excellent' }
})

const handleRegister = async () => {
  // Reset errors
  Object.keys(errors).forEach(key => {
    (errors as Record<string, string>)[key] = ''
  })

  // Validation
  if (!form.first_name) {
    errors.first_name = 'Please enter your first name'
    return
  }
  if (!form.last_name) {
    errors.last_name = 'Please enter your last name'
    return
  }
  if (!form.email) {
    errors.email = 'Please enter your email address'
    return
  }
  if (!form.password || form.password.length < 8) {
    errors.password = 'Password must be at least 8 characters'
    return
  }
  if (!form.user_type) {
    errors.user_type = 'Please select your account type'
    return
  }

  loading.value = true
  try {
    const success = await authStore.register({
      first_name: form.first_name,
      last_name: form.last_name,
      email: form.email,
      password: form.password,
      organization_type: form.user_type === 'agency' ? 'agency' : 'independent',
    })
    if (success) {
      router.push('/dashboard')
    } else {
      errors.general = 'This email may already be in use. Try signing in instead.'
    }
  } catch (error) {
    errors.general = 'Something went wrong. Please try again later.'
  } finally {
    loading.value = false
  }
}
</script>
