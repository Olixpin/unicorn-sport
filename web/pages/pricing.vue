<template>
  <div class="min-h-screen bg-neutral-50">
    <!-- Hero Header -->
    <div class="relative bg-neutral-950 overflow-hidden">
      <!-- Background -->
      <div class="absolute inset-0">
        <div class="absolute inset-0 bg-gradient-to-br from-neutral-950 via-primary-950/20 to-neutral-950"></div>
        <div class="absolute top-0 left-1/2 -translate-x-1/2 w-[1000px] h-[500px] bg-gradient-to-b from-primary-500/10 via-transparent to-transparent"></div>
      </div>
      
      <div class="relative max-w-5xl mx-auto px-4 sm:px-6 lg:px-8 pt-28 lg:pt-32 pb-20 text-center">
        <div class="inline-flex items-center gap-2 px-4 py-2 rounded-full bg-white/5 border border-white/10 mb-6">
          <span class="text-sm text-neutral-300">Simple, transparent pricing</span>
        </div>
        
        <h1 class="font-display text-4xl lg:text-5xl font-bold text-white mb-4">
          Choose the plan that's right for you
        </h1>
        <p class="text-lg text-neutral-400 max-w-2xl mx-auto">
          Start free and upgrade as you grow. All plans include access to verified African football talent.
        </p>

        <!-- Billing toggle -->
        <div class="mt-10 inline-flex items-center gap-4 p-1.5 bg-white/5 rounded-xl border border-white/10">
          <button 
            :class="[
              'px-5 py-2 rounded-lg text-sm font-medium transition-all duration-200',
              !isAnnual ? 'bg-white text-neutral-900' : 'text-neutral-400 hover:text-white'
            ]"
            @click="isAnnual = false"
          >
            Monthly
          </button>
          <button 
            :class="[
              'px-5 py-2 rounded-lg text-sm font-medium transition-all duration-200 flex items-center gap-2',
              isAnnual ? 'bg-white text-neutral-900' : 'text-neutral-400 hover:text-white'
            ]"
            @click="isAnnual = true"
          >
            Annual
            <span class="px-2 py-0.5 bg-primary-500 text-white text-xs font-bold rounded-full">Save 20%</span>
          </button>
        </div>
      </div>
    </div>

    <!-- Pricing Cards -->
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 -mt-8 pb-20">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <!-- Free Tier -->
        <div class="bg-white rounded-2xl border border-neutral-200/80 shadow-sm overflow-hidden">
          <div class="p-6">
            <div class="flex items-center gap-3 mb-4">
              <div class="w-10 h-10 rounded-xl bg-neutral-100 flex items-center justify-center">
                <svg class="w-5 h-5 text-neutral-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                </svg>
              </div>
              <h3 class="text-lg font-semibold text-neutral-900">Free</h3>
            </div>
            
            <div class="mb-4">
              <span class="text-4xl font-bold text-neutral-900">$0</span>
              <span class="text-neutral-500">/month</span>
            </div>
            <p class="text-sm text-neutral-500 mb-6">Perfect for exploring the platform</p>

            <NuxtLink to="/discover" class="block">
              <button class="w-full py-3 px-4 rounded-xl border-2 border-neutral-200 text-neutral-700 font-semibold hover:bg-neutral-50 hover:border-neutral-300 transition-all duration-200">
                Get Started
              </button>
            </NuxtLink>
          </div>
          
          <div class="border-t border-neutral-100 p-6">
            <p class="text-xs font-semibold text-neutral-400 uppercase tracking-wider mb-4">Includes</p>
            <ul class="space-y-3">
              <li class="flex items-center gap-3 text-sm text-neutral-600">
                <svg class="w-5 h-5 text-primary-500 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                </svg>
                Browse player profiles
              </li>
              <li class="flex items-center gap-3 text-sm text-neutral-600">
                <svg class="w-5 h-5 text-primary-500 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                </svg>
                Watch highlight videos
              </li>
              <li class="flex items-center gap-3 text-sm text-neutral-600">
                <svg class="w-5 h-5 text-primary-500 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                </svg>
                Basic search filters
              </li>
            </ul>
          </div>
        </div>

        <!-- Scout Tier -->
        <div class="bg-white rounded-2xl border border-neutral-200/80 shadow-sm overflow-hidden">
          <div class="p-6">
            <div class="flex items-center gap-3 mb-4">
              <div class="w-10 h-10 rounded-xl bg-primary-100 flex items-center justify-center">
                <svg class="w-5 h-5 text-primary-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                </svg>
              </div>
              <h3 class="text-lg font-semibold text-neutral-900">Scout</h3>
            </div>
            
            <div class="mb-4">
              <span class="text-4xl font-bold text-neutral-900">{{ isAnnual ? '$79' : '$99' }}</span>
              <span class="text-neutral-500">/month</span>
              <span v-if="isAnnual" class="ml-2 text-sm text-primary-600 font-medium">Save $240/yr</span>
            </div>
            <p class="text-sm text-neutral-500 mb-6">For active talent scouts</p>

            <button 
              :disabled="loading === 'scout'"
              :class="[
                'w-full py-3 px-4 rounded-xl font-semibold transition-all duration-200 flex items-center justify-center gap-2',
                loading === 'scout' 
                  ? 'bg-neutral-100 text-neutral-400 cursor-not-allowed' 
                  : 'border-2 border-neutral-200 text-neutral-700 hover:bg-neutral-50 hover:border-neutral-300'
              ]"
              @click="subscribe('scout')"
            >
              <svg v-if="loading === 'scout'" class="w-5 h-5 animate-spin" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
              </svg>
              Subscribe
            </button>
          </div>
          
          <div class="border-t border-neutral-100 p-6">
            <p class="text-xs font-semibold text-neutral-400 uppercase tracking-wider mb-4">Everything in Free, plus</p>
            <ul class="space-y-3">
              <li class="flex items-center gap-3 text-sm text-neutral-600">
                <svg class="w-5 h-5 text-primary-500 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                </svg>
                Full match recordings
              </li>
              <li class="flex items-center gap-3 text-sm text-neutral-600">
                <svg class="w-5 h-5 text-primary-500 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                </svg>
                Save to favorites
              </li>
              <li class="flex items-center gap-3 text-sm text-neutral-600">
                <svg class="w-5 h-5 text-primary-500 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                </svg>
                Advanced filters
              </li>
            </ul>
          </div>
        </div>

        <!-- Pro Tier (Popular) -->
        <div class="relative bg-neutral-900 rounded-2xl shadow-xl overflow-hidden ring-2 ring-primary-500">
          <div class="absolute top-0 inset-x-0 h-1 bg-gradient-to-r from-primary-500 to-secondary-500"></div>
          <div class="absolute top-4 right-4">
            <span class="px-3 py-1 bg-primary-500 text-white text-xs font-bold rounded-full">
              MOST POPULAR
            </span>
          </div>
          
          <div class="p-6 pt-8">
            <div class="flex items-center gap-3 mb-4">
              <div class="w-10 h-10 rounded-xl bg-primary-500/20 flex items-center justify-center">
                <svg class="w-5 h-5 text-primary-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
                </svg>
              </div>
              <h3 class="text-lg font-semibold text-white">Pro</h3>
            </div>
            
            <div class="mb-4">
              <span class="text-4xl font-bold text-white">{{ isAnnual ? '$239' : '$299' }}</span>
              <span class="text-neutral-400">/month</span>
              <span v-if="isAnnual" class="ml-2 text-sm text-primary-400 font-medium">Save $720/yr</span>
            </div>
            <p class="text-sm text-neutral-400 mb-6">For serious recruitment</p>

            <button 
              :disabled="loading === 'pro'"
              :class="[
                'w-full py-3 px-4 rounded-xl font-semibold transition-all duration-200 flex items-center justify-center gap-2',
                loading === 'pro' 
                  ? 'bg-primary-400 cursor-not-allowed' 
                  : 'bg-primary-500 text-white hover:bg-primary-600 hover:shadow-lg hover:shadow-primary-500/25'
              ]"
              @click="subscribe('pro')"
            >
              <svg v-if="loading === 'pro'" class="w-5 h-5 animate-spin" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
              </svg>
              Subscribe
            </button>
          </div>
          
          <div class="border-t border-white/10 p-6">
            <p class="text-xs font-semibold text-neutral-500 uppercase tracking-wider mb-4">Everything in Scout, plus</p>
            <ul class="space-y-3">
              <li class="flex items-center gap-3 text-sm text-neutral-300">
                <svg class="w-5 h-5 text-primary-400 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                </svg>
                Contact players directly
              </li>
              <li class="flex items-center gap-3 text-sm text-neutral-300">
                <svg class="w-5 h-5 text-primary-400 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                </svg>
                Priority support
              </li>
              <li class="flex items-center gap-3 text-sm text-neutral-300">
                <svg class="w-5 h-5 text-primary-400 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                </svg>
                Export data (PDF)
              </li>
              <li class="flex items-center gap-3 text-sm text-neutral-300">
                <svg class="w-5 h-5 text-primary-400 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                </svg>
                10 contacts/month
              </li>
            </ul>
          </div>
        </div>

        <!-- Club Tier -->
        <div class="bg-white rounded-2xl border border-neutral-200/80 shadow-sm overflow-hidden">
          <div class="p-6">
            <div class="flex items-center gap-3 mb-4">
              <div class="w-10 h-10 rounded-xl bg-secondary-100 flex items-center justify-center">
                <svg class="w-5 h-5 text-secondary-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
                </svg>
              </div>
              <h3 class="text-lg font-semibold text-neutral-900">Club</h3>
            </div>
            
            <div class="mb-4">
              <span class="text-4xl font-bold text-neutral-900">{{ isAnnual ? '$799' : '$999' }}</span>
              <span class="text-neutral-500">/month</span>
              <span v-if="isAnnual" class="ml-2 text-sm text-primary-600 font-medium">Save $2,400/yr</span>
            </div>
            <p class="text-sm text-neutral-500 mb-6">For clubs & agencies</p>

            <button 
              class="w-full py-3 px-4 rounded-xl border-2 border-neutral-200 text-neutral-700 font-semibold hover:bg-neutral-50 hover:border-neutral-300 transition-all duration-200"
              @click="contactSales"
            >
              Contact Sales
            </button>
          </div>
          
          <div class="border-t border-neutral-100 p-6">
            <p class="text-xs font-semibold text-neutral-400 uppercase tracking-wider mb-4">Everything in Pro, plus</p>
            <ul class="space-y-3">
              <li class="flex items-center gap-3 text-sm text-neutral-600">
                <svg class="w-5 h-5 text-primary-500 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                </svg>
                API access
              </li>
              <li class="flex items-center gap-3 text-sm text-neutral-600">
                <svg class="w-5 h-5 text-primary-500 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                </svg>
                Dedicated manager
              </li>
              <li class="flex items-center gap-3 text-sm text-neutral-600">
                <svg class="w-5 h-5 text-primary-500 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                </svg>
                Unlimited contacts
              </li>
              <li class="flex items-center gap-3 text-sm text-neutral-600">
                <svg class="w-5 h-5 text-primary-500 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                </svg>
                Custom reporting
              </li>
            </ul>
          </div>
        </div>
      </div>

      <!-- Trust badges -->
      <div class="mt-16 flex flex-wrap items-center justify-center gap-8 text-neutral-400">
        <div class="flex items-center gap-2">
          <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
          </svg>
          <span class="text-sm">Secure payments</span>
        </div>
        <div class="flex items-center gap-2">
          <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
          <span class="text-sm">Cancel anytime</span>
        </div>
        <div class="flex items-center gap-2">
          <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z" />
          </svg>
          <span class="text-sm">All major cards accepted</span>
        </div>
      </div>
    </div>

    <!-- FAQ Section -->
    <div class="bg-white border-t border-neutral-200">
      <div class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 py-20">
        <div class="text-center mb-12">
          <h2 class="font-display text-3xl font-bold text-neutral-900 mb-4">
            Frequently Asked Questions
          </h2>
          <p class="text-neutral-600">Have questions? We've got answers.</p>
        </div>

        <div class="space-y-4">
          <div 
            v-for="(faq, index) in faqs" 
            :key="index"
            class="border border-neutral-200 rounded-xl overflow-hidden"
          >
            <button 
              class="w-full flex items-center justify-between p-5 text-left hover:bg-neutral-50 transition-colors"
              @click="toggleFaq(index)"
            >
              <span class="font-semibold text-neutral-900">{{ faq.question }}</span>
              <svg 
                :class="[
                  'w-5 h-5 text-neutral-400 transition-transform duration-200',
                  openFaq === index ? 'rotate-180' : ''
                ]" 
                fill="none" 
                viewBox="0 0 24 24" 
                stroke="currentColor"
              >
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
              </svg>
            </button>
            <div 
              v-show="openFaq === index"
              class="px-5 pb-5 text-neutral-600"
            >
              {{ faq.answer }}
            </div>
          </div>
        </div>

        <!-- CTA -->
        <div class="mt-12 text-center">
          <p class="text-neutral-600 mb-4">Still have questions?</p>
          <a 
            href="mailto:support@unicornsport.com" 
            class="inline-flex items-center gap-2 text-primary-600 font-semibold hover:text-primary-700"
          >
            Contact our team
            <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3" />
            </svg>
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { SubscriptionTier } from '~/types'

definePageMeta({
  layout: 'default',
  keepalive: false,
})

const authStore = useAuthStore()
const subStore = useSubscriptionStore()
const router = useRouter()

const loading = ref<SubscriptionTier | null>(null)
const isAnnual = ref(false)
const openFaq = ref<number | null>(null)

const faqs = [
  {
    question: 'How does billing work?',
    answer: 'All plans are billed monthly or annually (with a 20% discount). You can upgrade, downgrade, or cancel at any time. Changes take effect at the start of your next billing cycle.',
  },
  {
    question: 'Can I cancel anytime?',
    answer: "Yes, you can cancel your subscription at any time. You'll continue to have access until the end of your current billing period. No questions asked.",
  },
  {
    question: 'How do contact requests work?',
    answer: 'Pro and Club tier subscribers can submit contact requests for players. Our academy team reviews each request and facilitates introductions between scouts and players within 48 hours.',
  },
  {
    question: 'Are players verified?',
    answer: 'Yes! All players on our platform have their age verified through official documents (NIN or Passport) before participating in our events. This is our core value proposition.',
  },
  {
    question: 'What payment methods do you accept?',
    answer: 'We accept all major credit cards (Visa, Mastercard, American Express), as well as bank transfers for Club tier subscriptions.',
  },
]

const toggleFaq = (index: number) => {
  openFaq.value = openFaq.value === index ? null : index
}

const subscribe = async (tier: SubscriptionTier) => {
  if (!authStore.isAuthenticated) {
    router.push(`/auth/register?plan=${tier}`)
    return
  }

  loading.value = tier
  try {
    const checkoutUrl = await subStore.createCheckoutSession(tier)
    if (checkoutUrl) {
      window.location.href = checkoutUrl
    }
  } catch (error) {
    console.error('Failed to create checkout session:', error)
  } finally {
    loading.value = null
  }
}

const contactSales = () => {
  window.location.href = 'mailto:sales@unicornsport.com?subject=Club Tier Inquiry'
}
</script>
