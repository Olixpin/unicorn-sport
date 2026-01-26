<template>
  <div class="min-h-screen flex">
    <!-- Left Panel - Branding & Social Proof -->
    <div class="hidden lg:flex lg:w-1/2 xl:w-[55%] relative bg-neutral-950 overflow-hidden">
      <!-- Background layers -->
      <div class="absolute inset-0">
        <!-- Base gradient -->
        <div class="absolute inset-0 bg-gradient-to-br from-neutral-950 via-primary-950/30 to-neutral-950"></div>
        
        <!-- Soft glow -->
        <div class="absolute top-1/4 left-1/2 -translate-x-1/2 w-[600px] h-[600px] bg-primary-500/10 rounded-full blur-[150px]"></div>
        <div class="absolute bottom-0 left-0 w-[400px] h-[400px] bg-secondary-500/10 rounded-full blur-[120px]"></div>
        
        <!-- Subtle pitch pattern -->
        <div class="absolute inset-0 opacity-[0.03]">
          <svg class="w-full h-full" viewBox="0 0 800 800" preserveAspectRatio="xMidYMid slice">
            <circle cx="400" cy="400" r="150" fill="none" stroke="white" stroke-width="2"/>
            <circle cx="400" cy="400" r="5" fill="white"/>
            <line x1="400" y1="0" x2="400" y2="800" stroke="white" stroke-width="2"/>
          </svg>
        </div>
      </div>

      <!-- Content -->
      <div class="relative z-10 flex flex-col justify-between p-12 xl:p-16 w-full">
        <!-- Logo -->
        <div>
          <NuxtLink to="/" class="inline-flex items-center gap-3 group">
            <span class="text-4xl transition-transform group-hover:scale-110">🦄</span>
            <span class="font-display font-bold text-2xl text-white">Unicorn Sport</span>
          </NuxtLink>
        </div>

        <!-- Main message -->
        <div class="max-w-lg">
          <h1 class="font-display text-4xl xl:text-5xl font-bold text-white leading-tight">
            Discover Africa's next 
            <span class="bg-gradient-to-r from-primary-400 to-secondary-400 bg-clip-text text-transparent">football superstars</span>
          </h1>
          <p class="mt-6 text-lg text-neutral-400 leading-relaxed">
            Join elite scouts and agencies worldwide who trust Unicorn Sport to find verified talent across 15+ African countries.
          </p>

          <!-- Stats -->
          <div class="mt-10 flex items-center gap-8">
            <div>
              <div class="text-3xl font-bold text-white">500+</div>
              <div class="text-sm text-neutral-500">Verified Players</div>
            </div>
            <div class="w-px h-12 bg-white/10"></div>
            <div>
              <div class="text-3xl font-bold text-white">50+</div>
              <div class="text-sm text-neutral-500">Scouts Active</div>
            </div>
            <div class="w-px h-12 bg-white/10"></div>
            <div>
              <div class="text-3xl font-bold text-white">15+</div>
              <div class="text-sm text-neutral-500">Countries</div>
            </div>
          </div>
        </div>

        <!-- Testimonial -->
        <div class="max-w-lg">
          <div class="flex items-start gap-4 p-6 rounded-2xl bg-white/5 border border-white/10 backdrop-blur-sm">
            <div class="flex-shrink-0 w-12 h-12 rounded-full bg-gradient-to-br from-primary-500 to-emerald-500 flex items-center justify-center">
              <span class="text-white font-bold text-lg">JM</span>
            </div>
            <div>
              <p class="text-neutral-300 text-sm leading-relaxed">
                "Found 3 players for our academy within the first week. The video quality and verification system is outstanding."
              </p>
              <div class="mt-3 flex items-center gap-2">
                <span class="text-white font-medium text-sm">James Mensah</span>
                <span class="text-neutral-600">•</span>
                <span class="text-neutral-500 text-sm">Head Scout, Premier FC</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Right Panel - Form -->
    <div class="flex-1 flex flex-col min-h-screen bg-white">
      <!-- Mobile header -->
      <div class="lg:hidden p-6 border-b border-neutral-100">
        <NuxtLink to="/" class="inline-flex items-center gap-2">
          <span class="text-3xl">🦄</span>
          <span class="font-display font-bold text-xl text-primary-600">Unicorn Sport</span>
        </NuxtLink>
      </div>

      <!-- Form container -->
      <div class="flex-1 flex items-center justify-center p-6 sm:p-12">
        <div class="w-full max-w-md">
          <slot />
        </div>
      </div>

      <!-- Footer -->
      <div class="p-6 text-center text-sm text-neutral-500 border-t border-neutral-100">
        <p>© {{ new Date().getFullYear() }} Unicorn Sport. All rights reserved.</p>
      </div>
    </div>
  </div>
</template>
