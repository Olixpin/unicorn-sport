<template>
  <div class="min-h-screen bg-neutral-950">
    <!-- Hero Header - Premium Dark with Pattern -->
    <div class="relative overflow-hidden">
      <!-- Background Layers -->
      <div class="absolute inset-0">
        <div class="absolute inset-0 bg-gradient-to-b from-neutral-950 via-neutral-950 to-neutral-900"></div>
        
        <!-- Premium Noise Texture -->
        <div class="absolute inset-0 opacity-[0.03]" :style="{ backgroundImage: noiseTexture }"></div>
        
        <!-- Glowing Orbs -->
        <div class="absolute top-0 left-1/4 w-[600px] h-[400px] bg-primary-500/10 rounded-full blur-[120px]"></div>
        <div class="absolute top-0 right-1/4 w-[500px] h-[300px] bg-emerald-500/10 rounded-full blur-[100px]"></div>
        
        <!-- Grid Pattern -->
        <div class="absolute inset-0 bg-[linear-gradient(rgba(255,255,255,.015)_1px,transparent_1px),linear-gradient(90deg,rgba(255,255,255,.015)_1px,transparent_1px)] bg-[size:64px_64px]"></div>
        
        <!-- Vignette -->
        <div class="absolute inset-0 bg-[radial-gradient(ellipse_at_center,transparent_0%,rgba(0,0,0,0.3)_100%)]"></div>
      </div>
      
      <div class="relative max-w-5xl mx-auto px-4 sm:px-6 lg:px-8 pt-16 lg:pt-20 pb-32 lg:pb-40 text-center">
        <!-- Badge -->
        <div class="inline-flex items-center gap-2 px-3 py-1.5 rounded-full bg-white/5 border border-white/10 backdrop-blur-sm mb-6">
          <span class="relative flex h-2 w-2">
            <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-emerald-400 opacity-75"></span>
            <span class="relative inline-flex rounded-full h-2 w-2 bg-emerald-500"></span>
          </span>
          <span class="text-xs text-neutral-300">Simple, transparent pricing</span>
        </div>
        
        <h1 class="font-display text-3xl sm:text-4xl lg:text-5xl font-bold text-white mb-4 leading-tight">
          Invest in <span class="text-transparent bg-clip-text bg-gradient-to-r from-primary-400 via-emerald-400 to-primary-400 bg-[length:200%_auto] animate-gradient">Tomorrow's Stars</span>
        </h1>
        <p class="text-base lg:text-lg text-neutral-400 max-w-xl mx-auto">
          Start free and upgrade as you grow. Access Africa's most promising verified football talent.
        </p>

        <!-- Billing Toggle - Premium Style -->
        <div class="mt-8 inline-flex items-center gap-1 p-1 bg-white/5 rounded-xl border border-white/10 backdrop-blur-sm">
          <button 
            :class="[
              'px-5 py-2.5 rounded-lg text-sm font-semibold transition-all duration-300',
              !isAnnual ? 'bg-white text-neutral-900 shadow-lg' : 'text-neutral-400 hover:text-white'
            ]"
            @click="isAnnual = false"
          >
            Monthly
          </button>
          <button 
            :class="[
              'px-5 py-2.5 rounded-lg text-sm font-semibold transition-all duration-300 flex items-center gap-2',
              isAnnual ? 'bg-white text-neutral-900 shadow-lg' : 'text-neutral-400 hover:text-white'
            ]"
            @click="isAnnual = true"
          >
            Annual
            <span class="px-1.5 py-0.5 bg-gradient-to-r from-primary-500 to-emerald-500 text-white text-[10px] font-bold rounded-full">-20%</span>
          </button>
        </div>
      </div>
    </div>

    <!-- Pricing Cards -->
    <div class="relative max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 pb-24 -mt-20 lg:-mt-28">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-5">
        
        <!-- Free Tier -->
        <div class="group relative flex">
          <div class="absolute -inset-px bg-gradient-to-b from-neutral-700/50 to-transparent rounded-2xl opacity-0 group-hover:opacity-100 transition-all duration-300"></div>
          <div class="relative flex flex-col w-full bg-neutral-900/80 backdrop-blur-sm rounded-2xl border border-neutral-800/80 overflow-hidden transition-all duration-300 group-hover:border-neutral-700 group-hover:-translate-y-1">
            <!-- Card Header -->
            <div class="p-6 pb-4">
              <div class="flex items-center gap-3 mb-6">
                <div class="w-11 h-11 rounded-xl bg-neutral-800 flex items-center justify-center">
                  <svg class="w-5 h-5 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                  </svg>
                </div>
                <div>
                  <h3 class="font-semibold text-white">Explorer</h3>
                  <p class="text-xs text-neutral-500">Start discovering</p>
                </div>
              </div>
              
              <div class="mb-6">
                <div class="flex items-baseline gap-1">
                  <span class="text-4xl font-bold text-white">$0</span>
                  <span class="text-neutral-500 text-sm">/mo</span>
                </div>
                <p class="text-xs text-neutral-500 mt-1 h-4">Free forever</p>
              </div>

              <NuxtLink to="/discover" class="block">
                <button class="w-full py-2.5 px-4 rounded-lg bg-neutral-800 text-white text-sm font-medium hover:bg-neutral-700 transition-all duration-200">
                  Get Started
                </button>
              </NuxtLink>
            </div>
            
            <!-- Features -->
            <div class="flex-1 border-t border-neutral-800/80 p-6 pt-5">
              <p class="text-[11px] font-medium text-neutral-500 uppercase tracking-wide mb-4">What's included</p>
              <ul class="space-y-3">
                <li class="flex items-start gap-2.5 text-sm text-neutral-300">
                  <svg class="w-4 h-4 text-neutral-600 mt-0.5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" /></svg>
                  Browse player profiles
                </li>
                <li class="flex items-start gap-2.5 text-sm text-neutral-300">
                  <svg class="w-4 h-4 text-neutral-600 mt-0.5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" /></svg>
                  Watch highlights
                </li>
                <li class="flex items-start gap-2.5 text-sm text-neutral-300">
                  <svg class="w-4 h-4 text-neutral-600 mt-0.5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" /></svg>
                  Basic filters
                </li>
              </ul>
            </div>
          </div>
        </div>

        <!-- Scout Tier -->
        <div class="group relative flex">
          <div class="absolute -inset-px bg-gradient-to-b from-primary-500/30 to-transparent rounded-2xl opacity-0 group-hover:opacity-100 transition-all duration-300"></div>
          <div class="relative flex flex-col w-full bg-neutral-900/80 backdrop-blur-sm rounded-2xl border border-neutral-800/80 overflow-hidden transition-all duration-300 group-hover:border-primary-500/40 group-hover:-translate-y-1">
            <!-- Card Header -->
            <div class="p-6 pb-4">
              <div class="flex items-center gap-3 mb-6">
                <div class="w-11 h-11 rounded-xl bg-primary-500/10 flex items-center justify-center">
                  <svg class="w-5 h-5 text-primary-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                  </svg>
                </div>
                <div>
                  <h3 class="font-semibold text-white">Scout</h3>
                  <p class="text-xs text-neutral-500">For active scouts</p>
                </div>
              </div>
              
              <div class="mb-6">
                <div class="flex items-baseline gap-1">
                  <span class="text-4xl font-bold text-white">{{ isAnnual ? '$79' : '$99' }}</span>
                  <span class="text-neutral-500 text-sm">/mo</span>
                </div>
                <p v-if="isAnnual" class="text-xs text-primary-400 mt-1 h-4 font-medium">Save $240/year</p>
                <p v-else class="text-xs text-neutral-500 mt-1 h-4">Billed monthly</p>
              </div>

              <button 
                :disabled="loading === 'scout'"
                :class="[
                  'w-full py-2.5 px-4 rounded-lg text-sm font-medium transition-all duration-200',
                  loading === 'scout' 
                    ? 'bg-neutral-800 text-neutral-500 cursor-not-allowed' 
                    : 'bg-neutral-800 text-white hover:bg-neutral-700'
                ]"
                @click="subscribe('scout')"
              >
                <span v-if="loading === 'scout'" class="flex items-center justify-center gap-2">
                  <svg class="w-4 h-4 animate-spin" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
                  </svg>
                  Processing...
                </span>
                <span v-else>Get Started</span>
              </button>
            </div>
            
            <!-- Features -->
            <div class="flex-1 border-t border-neutral-800/80 p-6 pt-5">
              <p class="text-[11px] font-medium text-neutral-500 uppercase tracking-wide mb-4">Everything in Explorer, plus</p>
              <ul class="space-y-3">
                <li class="flex items-start gap-2.5 text-sm text-neutral-300">
                  <svg class="w-4 h-4 text-primary-500 mt-0.5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" /></svg>
                  Full match recordings
                </li>
                <li class="flex items-start gap-2.5 text-sm text-neutral-300">
                  <svg class="w-4 h-4 text-primary-500 mt-0.5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" /></svg>
                  Save to shortlist
                </li>
                <li class="flex items-start gap-2.5 text-sm text-neutral-300">
                  <svg class="w-4 h-4 text-primary-500 mt-0.5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" /></svg>
                  Advanced filters
                </li>
                <li class="flex items-start gap-2.5 text-sm text-neutral-300">
                  <svg class="w-4 h-4 text-primary-500 mt-0.5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" /></svg>
                  Player comparison
                </li>
              </ul>
            </div>
          </div>
        </div>

        <!-- Pro Tier (Popular) -->
        <div class="group relative flex">
          <!-- Glow effect -->
          <div class="absolute -inset-[2px] bg-gradient-to-b from-emerald-400 via-primary-500 to-emerald-600 rounded-2xl opacity-100"></div>
          <div class="absolute -inset-[2px] bg-gradient-to-b from-emerald-400 via-primary-500 to-emerald-600 rounded-2xl blur-xl opacity-30 group-hover:opacity-50 transition-opacity"></div>
          
          <div class="relative flex flex-col w-full bg-neutral-900 rounded-2xl overflow-hidden transition-all duration-300 group-hover:-translate-y-1">
            <!-- Popular Badge -->
            <div class="absolute top-0 inset-x-0 flex justify-center">
              <span class="px-3 py-1 bg-gradient-to-r from-emerald-500 to-primary-500 text-white text-[10px] font-bold uppercase tracking-wide rounded-b-lg">
                Most Popular
              </span>
            </div>
            
            <!-- Card Header -->
            <div class="p-6 pb-4 pt-10">
              <div class="flex items-center gap-3 mb-6">
                <div class="w-11 h-11 rounded-xl bg-gradient-to-br from-primary-500 to-emerald-500 flex items-center justify-center shadow-lg shadow-primary-500/20">
                  <svg class="w-5 h-5 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
                  </svg>
                </div>
                <div>
                  <h3 class="font-semibold text-white">Pro</h3>
                  <p class="text-xs text-neutral-400">For serious scouts</p>
                </div>
              </div>
              
              <div class="mb-6">
                <div class="flex items-baseline gap-1">
                  <span class="text-4xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-emerald-400 to-primary-400">{{ isAnnual ? '$239' : '$299' }}</span>
                  <span class="text-neutral-400 text-sm">/mo</span>
                </div>
                <p v-if="isAnnual" class="text-xs text-emerald-400 mt-1 h-4 font-medium">Save $720/year</p>
                <p v-else class="text-xs text-neutral-400 mt-1 h-4">Billed monthly</p>
              </div>

              <button 
                :disabled="loading === 'pro'"
                :class="[
                  'w-full py-2.5 px-4 rounded-lg text-sm font-semibold transition-all duration-200',
                  loading === 'pro' 
                    ? 'bg-primary-600/50 cursor-not-allowed text-white/50' 
                    : 'bg-gradient-to-r from-emerald-500 to-primary-500 text-white hover:shadow-lg hover:shadow-primary-500/25 hover:scale-[1.02]'
                ]"
                @click="subscribe('pro')"
              >
                <span v-if="loading === 'pro'" class="flex items-center justify-center gap-2">
                  <svg class="w-4 h-4 animate-spin" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
                  </svg>
                  Processing...
                </span>
                <span v-else>Start Free Trial</span>
              </button>
            </div>
            
            <!-- Features -->
            <div class="flex-1 border-t border-neutral-800 p-6 pt-5">
              <p class="text-[11px] font-medium text-neutral-400 uppercase tracking-wide mb-4">Everything in Scout, plus</p>
              <ul class="space-y-3">
                <li class="flex items-start gap-2.5 text-sm text-white font-medium">
                  <svg class="w-4 h-4 text-emerald-400 mt-0.5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" /></svg>
                  Contact players directly
                </li>
                <li class="flex items-start gap-2.5 text-sm text-neutral-300">
                  <svg class="w-4 h-4 text-emerald-400 mt-0.5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" /></svg>
                  10 contact requests/mo
                </li>
                <li class="flex items-start gap-2.5 text-sm text-neutral-300">
                  <svg class="w-4 h-4 text-emerald-400 mt-0.5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" /></svg>
                  Export PDF reports
                </li>
                <li class="flex items-start gap-2.5 text-sm text-neutral-300">
                  <svg class="w-4 h-4 text-emerald-400 mt-0.5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" /></svg>
                  Priority support
                </li>
              </ul>
            </div>
          </div>
        </div>

        <!-- Club Tier -->
        <div class="group relative flex">
          <div class="absolute -inset-px bg-gradient-to-b from-secondary-500/30 to-transparent rounded-2xl opacity-0 group-hover:opacity-100 transition-all duration-300"></div>
          <div class="relative flex flex-col w-full bg-neutral-900/80 backdrop-blur-sm rounded-2xl border border-neutral-800/80 overflow-hidden transition-all duration-300 group-hover:border-secondary-500/40 group-hover:-translate-y-1">
            <!-- Card Header -->
            <div class="p-6 pb-4">
              <div class="flex items-center gap-3 mb-6">
                <div class="w-11 h-11 rounded-xl bg-secondary-500/10 flex items-center justify-center">
                  <svg class="w-5 h-5 text-secondary-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
                  </svg>
                </div>
                <div>
                  <h3 class="font-semibold text-white">Club</h3>
                  <p class="text-xs text-neutral-500">For clubs & agencies</p>
                </div>
              </div>
              
              <div class="mb-6">
                <div class="flex items-baseline gap-1">
                  <span class="text-4xl font-bold text-white">{{ isAnnual ? '$799' : '$999' }}</span>
                  <span class="text-neutral-500 text-sm">/mo</span>
                </div>
                <p v-if="isAnnual" class="text-xs text-secondary-400 mt-1 h-4 font-medium">Save $2,400/year</p>
                <p v-else class="text-xs text-neutral-500 mt-1 h-4">Billed monthly</p>
              </div>

              <button 
                class="w-full py-2.5 px-4 rounded-lg border border-secondary-500/40 text-secondary-400 text-sm font-medium hover:bg-secondary-500/10 hover:border-secondary-500/60 transition-all duration-200"
                @click="contactSales"
              >
                Contact Sales
              </button>
            </div>
            
            <!-- Features -->
            <div class="flex-1 border-t border-neutral-800/80 p-6 pt-5">
              <p class="text-[11px] font-medium text-neutral-500 uppercase tracking-wide mb-4">Everything in Pro, plus</p>
              <ul class="space-y-3">
                <li class="flex items-start gap-2.5 text-sm text-white font-medium">
                  <svg class="w-4 h-4 text-secondary-400 mt-0.5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" /></svg>
                  Unlimited contacts
                </li>
                <li class="flex items-start gap-2.5 text-sm text-neutral-300">
                  <svg class="w-4 h-4 text-secondary-400 mt-0.5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" /></svg>
                  Account manager
                </li>
                <li class="flex items-start gap-2.5 text-sm text-neutral-300">
                  <svg class="w-4 h-4 text-secondary-400 mt-0.5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" /></svg>
                  API access
                </li>
                <li class="flex items-start gap-2.5 text-sm text-neutral-300">
                  <svg class="w-4 h-4 text-secondary-400 mt-0.5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" /></svg>
                  Custom analytics
                </li>
              </ul>
            </div>
          </div>
        </div>
      </div>

      <!-- Trust Badges -->
      <div class="mt-20 flex flex-wrap items-center justify-center gap-10 lg:gap-16">
        <div class="flex items-center gap-3 text-neutral-500">
          <div class="w-10 h-10 rounded-xl bg-neutral-800/50 flex items-center justify-center">
            <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
            </svg>
          </div>
          <span class="text-sm font-medium">SSL Secured</span>
        </div>
        <div class="flex items-center gap-3 text-neutral-500">
          <div class="w-10 h-10 rounded-xl bg-neutral-800/50 flex items-center justify-center">
            <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
            </svg>
          </div>
          <span class="text-sm font-medium">Cancel Anytime</span>
        </div>
        <div class="flex items-center gap-3 text-neutral-500">
          <div class="w-10 h-10 rounded-xl bg-neutral-800/50 flex items-center justify-center">
            <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z" />
            </svg>
          </div>
          <span class="text-sm font-medium">All Cards Accepted</span>
        </div>
        <div class="flex items-center gap-3 text-neutral-500">
          <div class="w-10 h-10 rounded-xl bg-neutral-800/50 flex items-center justify-center">
            <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <span class="text-sm font-medium">24/7 Support</span>
        </div>
      </div>
    </div>

    <!-- Comparison Table -->
    <section class="py-28 lg:py-32 bg-neutral-900 border-t border-neutral-800">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <h2 class="font-display text-3xl lg:text-4xl font-bold text-white mb-4">
            Compare <span class="text-transparent bg-clip-text bg-gradient-to-r from-primary-400 to-emerald-400">All Features</span>
          </h2>
          <p class="text-neutral-400">Everything you need to find your next star player</p>
        </div>

        <div class="overflow-x-auto">
          <table class="w-full">
            <thead>
              <tr class="border-b border-neutral-800">
                <th class="text-left py-4 px-4 text-sm font-semibold text-neutral-400">Features</th>
                <th class="text-center py-4 px-4 text-sm font-semibold text-neutral-400">Explorer</th>
                <th class="text-center py-4 px-4 text-sm font-semibold text-neutral-400">Scout</th>
                <th class="text-center py-4 px-4 text-sm font-semibold text-transparent bg-clip-text bg-gradient-to-r from-primary-400 to-emerald-400">Pro</th>
                <th class="text-center py-4 px-4 text-sm font-semibold text-secondary-400">Club</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-neutral-800">
              <tr v-for="feature in comparisonFeatures" :key="feature.name" class="hover:bg-neutral-800/50 transition-colors">
                <td class="py-4 px-4 text-sm text-neutral-300">{{ feature.name }}</td>
                <td class="text-center py-4 px-4">
                  <span v-if="feature.explorer === true" class="text-primary-400">✓</span>
                  <span v-else-if="feature.explorer === false" class="text-neutral-600">—</span>
                  <span v-else class="text-sm text-neutral-400">{{ feature.explorer }}</span>
                </td>
                <td class="text-center py-4 px-4">
                  <span v-if="feature.scout === true" class="text-primary-400">✓</span>
                  <span v-else-if="feature.scout === false" class="text-neutral-600">—</span>
                  <span v-else class="text-sm text-neutral-400">{{ feature.scout }}</span>
                </td>
                <td class="text-center py-4 px-4 bg-primary-500/5">
                  <span v-if="feature.pro === true" class="text-emerald-400">✓</span>
                  <span v-else-if="feature.pro === false" class="text-neutral-600">—</span>
                  <span v-else class="text-sm text-emerald-400 font-medium">{{ feature.pro }}</span>
                </td>
                <td class="text-center py-4 px-4">
                  <span v-if="feature.club === true" class="text-secondary-400">✓</span>
                  <span v-else-if="feature.club === false" class="text-neutral-600">—</span>
                  <span v-else class="text-sm text-secondary-400 font-medium">{{ feature.club }}</span>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </section>

    <!-- FAQ Section -->
    <section class="py-28 lg:py-32 bg-neutral-950 border-t border-neutral-800">
      <div class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <div class="inline-flex items-center gap-2 px-4 py-2 rounded-full bg-white/5 border border-white/10 mb-6">
            <svg class="w-4 h-4 text-primary-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <span class="text-sm text-neutral-400">Got Questions?</span>
          </div>
          <h2 class="font-display text-3xl lg:text-4xl font-bold text-white mb-4">
            Frequently Asked <span class="text-transparent bg-clip-text bg-gradient-to-r from-primary-400 to-emerald-400">Questions</span>
          </h2>
        </div>

        <div class="space-y-4">
          <div 
            v-for="(faq, index) in faqs" 
            :key="index"
            class="group"
          >
            <div class="bg-neutral-900 border border-neutral-800 rounded-2xl overflow-hidden transition-all duration-300 hover:border-neutral-700">
              <button 
                class="w-full flex items-center justify-between p-6 text-left transition-colors"
                @click="toggleFaq(index)"
              >
                <span class="font-semibold text-white pr-4">{{ faq.question }}</span>
                <div :class="[
                  'w-8 h-8 rounded-xl flex items-center justify-center flex-shrink-0 transition-all duration-300',
                  openFaq === index ? 'bg-primary-500 rotate-180' : 'bg-neutral-800'
                ]">
                  <svg class="w-4 h-4 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                  </svg>
                </div>
              </button>
              <div 
                v-show="openFaq === index"
                class="px-6 pb-6 text-neutral-400 leading-relaxed"
              >
                {{ faq.answer }}
              </div>
            </div>
          </div>
        </div>

        <!-- Contact CTA -->
        <div class="mt-16 text-center p-8 rounded-2xl bg-gradient-to-r from-primary-500/10 to-emerald-500/10 border border-primary-500/20">
          <p class="text-white font-semibold mb-2">Still have questions?</p>
          <p class="text-neutral-400 text-sm mb-6">Our team is ready to help you find the perfect plan</p>
          <a 
            href="mailto:support@unicornsport.com" 
            class="inline-flex items-center gap-2 px-6 py-3 bg-white text-neutral-900 rounded-xl font-semibold hover:scale-105 transition-transform"
          >
            <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
            </svg>
            Contact Support
          </a>
        </div>
      </div>
    </section>

    <!-- Final CTA -->
    <section class="py-28 lg:py-32 bg-neutral-900 border-t border-neutral-800 relative overflow-hidden">
      <div class="absolute inset-0">
        <div class="absolute top-0 left-1/4 w-[500px] h-[500px] bg-primary-500/10 rounded-full blur-[120px]"></div>
        <div class="absolute bottom-0 right-1/4 w-[400px] h-[400px] bg-emerald-500/10 rounded-full blur-[100px]"></div>
      </div>
      
      <div class="relative max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h2 class="font-display text-4xl lg:text-5xl font-bold text-white mb-6">
          Ready to discover the next <span class="text-transparent bg-clip-text bg-gradient-to-r from-primary-400 to-emerald-400">superstar</span>?
        </h2>
        <p class="text-xl text-neutral-400 mb-10">
          Join hundreds of scouts and clubs already finding talent on Unicorn Sport
        </p>
        <div class="flex flex-col sm:flex-row gap-4 justify-center">
          <NuxtLink 
            to="/auth/register"
            class="group inline-flex items-center gap-3 px-8 py-4 bg-white text-neutral-900 rounded-xl font-bold hover:scale-105 transition-transform"
          >
            Start Free Trial
            <svg class="w-5 h-5 transition-transform group-hover:translate-x-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3" />
            </svg>
          </NuxtLink>
          <NuxtLink 
            to="/discover"
            class="inline-flex items-center gap-3 px-8 py-4 border border-neutral-700 text-white rounded-xl font-semibold hover:bg-neutral-800 transition-colors"
          >
            <svg class="w-5 h-5 text-primary-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.752 11.168l-3.197-2.132A1 1 0 0010 9.87v4.263a1 1 0 001.555.832l3.197-2.132a1 1 0 000-1.664z" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            Explore Players First
          </NuxtLink>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import type { SubscriptionTier } from '~/types'

definePageMeta({
  layout: 'default',
  keepalive: false,
})

useSeoMeta({
  title: 'Pricing - Unicorn Sport',
  description: 'Simple, transparent pricing for scouts, clubs, and agencies. Start free and upgrade as you grow.',
})

const authStore = useAuthStore()
const subStore = useSubscriptionStore()
const router = useRouter()

const loading = ref<SubscriptionTier | null>(null)
const isAnnual = ref(true)
const openFaq = ref<number | null>(null)

// Noise texture for premium background
const noiseTexture = "url(\"data:image/svg+xml,%3Csvg viewBox='0 0 256 256' xmlns='http://www.w3.org/2000/svg'%3E%3Cfilter id='noise'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.8' numOctaves='4' stitchTiles='stitch'/%3E%3C/filter%3E%3Crect width='100%25' height='100%25' filter='url(%23noise)'/%3E%3C/svg%3E\")"

const comparisonFeatures = [
  { name: 'Browse player profiles', explorer: true, scout: true, pro: true, club: true },
  { name: 'Watch highlight videos', explorer: true, scout: true, pro: true, club: true },
  { name: 'Basic search filters', explorer: true, scout: true, pro: true, club: true },
  { name: 'Full match recordings', explorer: false, scout: true, pro: true, club: true },
  { name: 'Save to shortlist', explorer: false, scout: true, pro: true, club: true },
  { name: 'Advanced filters', explorer: false, scout: true, pro: true, club: true },
  { name: 'Player comparison', explorer: false, scout: true, pro: true, club: true },
  { name: 'Contact players', explorer: false, scout: false, pro: '10/mo', club: 'Unlimited' },
  { name: 'Export PDF reports', explorer: false, scout: false, pro: true, club: true },
  { name: 'Priority support', explorer: false, scout: false, pro: true, club: true },
  { name: 'API access', explorer: false, scout: false, pro: false, club: true },
  { name: 'Dedicated manager', explorer: false, scout: false, pro: false, club: true },
  { name: 'Custom reporting', explorer: false, scout: false, pro: false, club: true },
]

const faqs = [
  {
    question: 'How does billing work?',
    answer: 'All plans are billed monthly or annually (with a 20% discount). You can upgrade, downgrade, or cancel at any time. Changes take effect at the start of your next billing cycle.',
  },
  {
    question: 'Can I cancel my subscription anytime?',
    answer: "Absolutely! You can cancel your subscription at any time with no questions asked. You'll continue to have access until the end of your current billing period.",
  },
  {
    question: 'How do contact requests work?',
    answer: 'Pro and Club tier subscribers can submit contact requests for players. Our academy team reviews each request and facilitates introductions between scouts and players within 48 hours.',
  },
  {
    question: 'Are all players verified?',
    answer: 'Yes! All players on our platform have their age verified through official documents (NIN or Passport) before participating in our events. This age verification is our core value proposition.',
  },
  {
    question: 'What payment methods do you accept?',
    answer: 'We accept all major credit cards (Visa, Mastercard, American Express), as well as bank transfers for Club tier subscriptions. All payments are processed securely via Stripe.',
  },
  {
    question: 'Is there a free trial?',
    answer: "Yes! You can start with our Explorer tier completely free, forever. When you're ready to unlock more features, you can upgrade to Scout, Pro, or Club tier.",
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

<style scoped>
@keyframes gradient {
  0%, 100% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
}

.animate-gradient {
  animation: gradient 6s ease infinite;
}
</style>
