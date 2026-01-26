<template>
  <footer class="bg-neutral-950 text-white relative overflow-hidden">
    <!-- Football Pitch Pattern Background -->
    <div class="absolute inset-0 opacity-[0.03]">
      <svg class="w-full h-full" viewBox="0 0 800 400" preserveAspectRatio="xMidYMid slice">
        <rect x="40" y="20" width="720" height="360" fill="none" stroke="currentColor" stroke-width="2"/>
        <line x1="400" y1="20" x2="400" y2="380" stroke="currentColor" stroke-width="2"/>
        <circle cx="400" cy="200" r="60" fill="none" stroke="currentColor" stroke-width="2"/>
        <rect x="40" y="100" width="100" height="200" fill="none" stroke="currentColor" stroke-width="2"/>
        <rect x="660" y="100" width="100" height="200" fill="none" stroke="currentColor" stroke-width="2"/>
      </svg>
    </div>

    <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pt-20 pb-12">
      <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-8 lg:gap-12">
        <!-- Brand Column -->
        <div class="col-span-2">
          <NuxtLink to="/" class="flex items-center gap-3 mb-6">
            <div class="w-12 h-12 rounded-2xl bg-gradient-to-br from-primary-500 to-emerald-500 flex items-center justify-center">
              <svg class="w-7 h-7 text-white" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
                <circle cx="12" cy="12" r="10"/>
                <path d="M12 2v4M12 18v4M2 12h4M18 12h4M5.6 5.6l2.8 2.8M15.6 15.6l2.8 2.8M5.6 18.4l2.8-2.8M15.6 8.4l2.8-2.8"/>
              </svg>
            </div>
            <div>
              <span class="font-display text-xl font-bold text-white">Unicorn Sport</span>
              <div class="text-xs text-neutral-500">Africa's Football Talent Hub</div>
            </div>
          </NuxtLink>
          <p class="text-neutral-400 text-sm leading-relaxed mb-6 max-w-xs">
            Connecting Africa's brightest football talent with scouts and clubs worldwide. Verified profiles, HD highlights, direct connections.
          </p>
          <div class="flex gap-3">
            <a v-for="social in socials" :key="social.name" :href="social.url" target="_blank" rel="noopener noreferrer"
              class="w-10 h-10 rounded-xl bg-white/5 border border-white/10 flex items-center justify-center text-neutral-400 hover:bg-primary-500 hover:border-primary-500 hover:text-white transition-all duration-300">
              <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24" v-html="social.icon"></svg>
            </a>
          </div>
        </div>

        <!-- Platform Links -->
        <div>
          <h4 class="font-semibold text-white mb-4 flex items-center gap-2">
            <span class="w-1.5 h-1.5 rounded-full bg-primary-500"></span>
            Platform
          </h4>
          <ul class="space-y-3">
            <li v-for="link in platformLinks" :key="link.label">
              <NuxtLink :to="link.to" class="text-neutral-400 hover:text-white transition-colors text-sm flex items-center gap-2 group">
                <span class="w-0 h-px bg-primary-500 group-hover:w-3 transition-all duration-300"></span>
                {{ link.label }}
              </NuxtLink>
            </li>
          </ul>
        </div>

        <!-- For Scouts -->
        <div>
          <h4 class="font-semibold text-white mb-4 flex items-center gap-2">
            <span class="w-1.5 h-1.5 rounded-full bg-emerald-500"></span>
            For Scouts
          </h4>
          <ul class="space-y-3">
            <li v-for="link in scoutLinks" :key="link.label">
              <NuxtLink :to="link.to" class="text-neutral-400 hover:text-white transition-colors text-sm flex items-center gap-2 group">
                <span class="w-0 h-px bg-emerald-500 group-hover:w-3 transition-all duration-300"></span>
                {{ link.label }}
              </NuxtLink>
            </li>
          </ul>
        </div>

        <!-- For Academies -->
        <div>
          <h4 class="font-semibold text-white mb-4 flex items-center gap-2">
            <span class="w-1.5 h-1.5 rounded-full bg-secondary-500"></span>
            For Academies
          </h4>
          <ul class="space-y-3">
            <li v-for="link in academyLinks" :key="link.label">
              <NuxtLink :to="link.to" class="text-neutral-400 hover:text-white transition-colors text-sm flex items-center gap-2 group">
                <span class="w-0 h-px bg-secondary-500 group-hover:w-3 transition-all duration-300"></span>
                {{ link.label }}
              </NuxtLink>
            </li>
          </ul>
        </div>

        <!-- Newsletter -->
        <div class="col-span-2 md:col-span-4 lg:col-span-1">
          <h4 class="font-semibold text-white mb-4 flex items-center gap-2">
            <span class="w-1.5 h-1.5 rounded-full bg-primary-500"></span>
            Stay Updated
          </h4>
          <p class="text-neutral-400 text-sm mb-4">Get talent alerts and platform updates.</p>
          <form @submit.prevent="subscribeNewsletter" class="space-y-3">
            <input v-model="email" type="email" placeholder="your@email.com"
              class="w-full px-4 py-3 bg-white/5 border border-white/10 rounded-xl text-white placeholder-neutral-500 text-sm focus:outline-none focus:border-primary-500 transition-all"/>
            <button type="submit" class="w-full px-4 py-3 bg-gradient-to-r from-primary-500 to-emerald-500 text-white font-semibold rounded-xl text-sm hover:shadow-lg hover:shadow-primary-500/25 transition-all">
              Subscribe
            </button>
          </form>
        </div>
      </div>

      <!-- Trust Badges -->
      <div class="mt-16 pt-8 border-t border-white/10">
        <div class="flex flex-wrap items-center justify-center gap-8 mb-8">
          <div class="flex items-center gap-2 text-neutral-500 text-sm">
            <svg class="w-5 h-5 text-primary-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
            </svg>
            <span>Verified Profiles</span>
          </div>
          <div class="flex items-center gap-2 text-neutral-500 text-sm">
            <svg class="w-5 h-5 text-primary-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
            </svg>
            <span>Secure Platform</span>
          </div>
          <div class="flex items-center gap-2 text-neutral-500 text-sm">
            <svg class="w-5 h-5 text-primary-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.055 11H5a2 2 0 012 2v1a2 2 0 002 2 2 2 0 012 2v2.945M8 3.935V5.5A2.5 2.5 0 0010.5 8h.5a2 2 0 012 2 2 2 0 104 0 2 2 0 012-2h1.064M15 20.488V18a2 2 0 012-2h3.064M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <span>15+ African Countries</span>
          </div>
          <div class="flex items-center gap-2 text-neutral-500 text-sm">
            <svg class="w-5 h-5 text-primary-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
            </svg>
            <span>200+ Active Scouts</span>
          </div>
        </div>
      </div>

      <!-- Bottom Bar -->
      <div class="mt-8 pt-8 border-t border-white/10 flex flex-col md:flex-row items-center justify-between gap-4">
        <div class="text-neutral-500 text-sm">© {{ new Date().getFullYear() }} Unicorn Sport. All rights reserved.</div>
        <div class="flex flex-wrap items-center gap-6 text-sm">
          <NuxtLink to="/privacy" class="text-neutral-500 hover:text-white transition-colors">Privacy Policy</NuxtLink>
          <NuxtLink to="/terms" class="text-neutral-500 hover:text-white transition-colors">Terms of Service</NuxtLink>
          <NuxtLink to="/contact" class="text-neutral-500 hover:text-white transition-colors">Contact</NuxtLink>
        </div>
      </div>
    </div>

    <!-- Decorative Football -->
    <div class="absolute bottom-0 right-0 w-64 h-64 opacity-5 translate-x-1/3 translate-y-1/3">
      <svg viewBox="0 0 100 100" fill="currentColor">
        <circle cx="50" cy="50" r="48" fill="none" stroke="currentColor" stroke-width="2"/>
        <polygon points="50,15 65,35 60,55 40,55 35,35" fill="currentColor"/>
        <polygon points="50,85 35,65 40,45 60,45 65,65" fill="currentColor"/>
      </svg>
    </div>
  </footer>
</template>

<script setup lang="ts">
const email = ref('')

const socials = [
  { name: 'Twitter/X', url: '#', icon: '<path d="M18.244 2.25h3.308l-7.227 8.26 8.502 11.24H16.17l-5.214-6.817L4.99 21.75H1.68l7.73-8.835L1.254 2.25H8.08l4.713 6.231zm-1.161 17.52h1.833L7.084 4.126H5.117z"/>' },
  { name: 'Instagram', url: '#', icon: '<path d="M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zM12 0C8.741 0 8.333.014 7.053.072 2.695.272.273 2.69.073 7.052.014 8.333 0 8.741 0 12c0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98C8.333 23.986 8.741 24 12 24c3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98C15.668.014 15.259 0 12 0zm0 5.838a6.162 6.162 0 100 12.324 6.162 6.162 0 000-12.324zM12 16a4 4 0 110-8 4 4 0 010 8zm6.406-11.845a1.44 1.44 0 100 2.881 1.44 1.44 0 000-2.881z"/>' },
  { name: 'YouTube', url: '#', icon: '<path d="M23.498 6.186a3.016 3.016 0 0 0-2.122-2.136C19.505 3.545 12 3.545 12 3.545s-7.505 0-9.377.505A3.017 3.017 0 0 0 .502 6.186C0 8.07 0 12 0 12s0 3.93.502 5.814a3.016 3.016 0 0 0 2.122 2.136c1.871.505 9.376.505 9.376.505s7.505 0 9.377-.505a3.015 3.015 0 0 0 2.122-2.136C24 15.93 24 12 24 12s0-3.93-.502-5.814zM9.545 15.568V8.432L15.818 12l-6.273 3.568z"/>' },
  { name: 'LinkedIn', url: '#', icon: '<path d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433c-1.144 0-2.063-.926-2.063-2.065 0-1.138.92-2.063 2.063-2.063 1.14 0 2.064.925 2.064 2.063 0 1.139-.925 2.065-2.064 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z"/>' }
]

const platformLinks = [
  { label: 'Discover Players', to: '/discover' },
  { label: 'Browse Academies', to: '/academies' },
  { label: 'How It Works', to: '/how-it-works' },
  { label: 'Pricing', to: '/pricing' }
]

const scoutLinks = [
  { label: 'Create Account', to: '/auth/register' },
  { label: 'Scout Dashboard', to: '/dashboard' },
  { label: 'Search Filters', to: '/discover' },
  { label: 'Contact Players', to: '/pricing' }
]

const academyLinks = [
  { label: 'Register Academy', to: '/academy/register' },
  { label: 'Add Players', to: '/academy/players' },
  { label: 'Upload Videos', to: '/academy/videos' },
  { label: 'Partner Program', to: '/partners' }
]

function subscribeNewsletter() {
  if (email.value) {
    console.log('Subscribe:', email.value)
    email.value = ''
  }
}
</script>
