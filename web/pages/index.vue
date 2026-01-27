<template>
  <div class="overflow-hidden">
    <!-- Hero Section - Clean African Football Atmosphere -->
    <section class="relative min-h-screen flex items-center bg-neutral-950 overflow-hidden">
      <!-- Simplified Background - Subtle but Evocative -->
      <div class="absolute inset-0">
        <!-- Base gradient with warm African undertones -->
        <div class="absolute inset-0 bg-gradient-to-b from-neutral-950 via-neutral-950 to-primary-950/30"></div>
        
        <!-- Single soft glow (stadium light feel) -->
        <div class="absolute top-0 left-1/2 -translate-x-1/2 w-[800px] h-[500px] bg-gradient-radial from-primary-500/8 via-transparent to-transparent"></div>
        
        <!-- Subtle warm horizon -->
        <div class="absolute bottom-0 inset-x-0 h-1/3 bg-gradient-to-t from-secondary-900/10 to-transparent"></div>
      </div>

      <!-- Single Pattern Layer - Football Pitch (very subtle) -->
      <div class="absolute inset-0 opacity-[0.025]">
        <svg class="w-full h-full" viewBox="0 0 1200 800" preserveAspectRatio="xMidYMid slice">
          <circle cx="600" cy="400" r="120" fill="none" stroke="white" stroke-width="2"/>
          <circle cx="600" cy="400" r="4" fill="white"/>
          <line x1="600" y1="0" x2="600" y2="800" stroke="white" stroke-width="2"/>
        </svg>
      </div>

      <!-- Content - Two Column Layout -->
      <div class="relative z-10 max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-32 w-full">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-12 lg:gap-20 items-center">
          <!-- Left: Text Content -->
          <div class="text-center lg:text-left">
            <!-- Badge -->
            <div class="inline-flex items-center gap-2 px-4 py-2 rounded-full bg-white/5 border border-white/10 backdrop-blur-sm mb-8">
              <span class="relative flex h-2 w-2">
                <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-primary-400 opacity-75"></span>
                <span class="relative inline-flex rounded-full h-2 w-2 bg-primary-500"></span>
              </span>
              <span class="text-sm text-neutral-300">Now scouting in 15+ African countries</span>
            </div>

            <!-- Main Headline -->
            <h1 class="font-display text-5xl sm:text-6xl lg:text-7xl font-bold tracking-tight text-white leading-[0.9]">
              <span class="block">The Future of</span>
              <span class="block mt-2 bg-gradient-to-r from-primary-400 via-secondary-400 to-primary-400 bg-clip-text text-transparent bg-[length:200%_auto] animate-gradient">
                African Football
              </span>
            </h1>

            <p class="mt-8 text-xl text-neutral-400 max-w-xl mx-auto lg:mx-0 leading-relaxed">
              Where elite scouts discover tomorrow's superstars. 
              <span class="text-white">Verified players.</span> 
              <span class="text-white">HD highlights.</span> 
              <span class="text-white">Direct connections.</span>
            </p>

            <!-- CTA Buttons -->
            <div class="mt-10 flex flex-col sm:flex-row gap-4 justify-center lg:justify-start">
              <NuxtLink to="/discover">
                <button class="group relative px-8 py-4 bg-white text-neutral-900 font-semibold rounded-xl overflow-hidden transition-all duration-300 hover:scale-105 hover:shadow-[0_0_40px_rgba(34,197,94,0.4)]">
                  <span class="relative z-10 flex items-center gap-2">
                    <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                    </svg>
                    Explore Players
                    <svg class="w-4 h-4 transition-transform group-hover:translate-x-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3" />
                    </svg>
                  </span>
                </button>
              </NuxtLink>
              <NuxtLink to="/auth/register">
                <button class="px-8 py-4 bg-white/5 text-white font-semibold rounded-xl border border-white/10 backdrop-blur-sm transition-all duration-300 hover:bg-white/10 hover:border-white/20 flex items-center gap-2">
                  <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                  </svg>
                  Start Free Trial
                </button>
              </NuxtLink>
            </div>

            <!-- Trust Badges -->
            <div class="mt-12 flex flex-wrap items-center justify-center lg:justify-start gap-6 text-neutral-500">
              <div class="flex items-center gap-2">
                <svg class="w-5 h-5 text-primary-500" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                </svg>
                <span class="text-sm">Verified profiles</span>
              </div>
              <div class="flex items-center gap-2">
                <svg class="w-5 h-5 text-primary-500" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                </svg>
                <span class="text-sm">50+ Academies</span>
              </div>
              <div class="flex items-center gap-2">
                <svg class="w-5 h-5 text-primary-500" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                </svg>
                <span class="text-sm">HD Highlights</span>
              </div>
            </div>
          </div>

          <!-- Right: Video Player Mockup -->
          <div class="relative hidden lg:block">
            <!-- Main Video Card -->
            <div class="relative group">
              <!-- Glowing Border Effect -->
              <div class="absolute -inset-1 bg-gradient-to-r from-primary-500 via-emerald-500 to-primary-500 rounded-3xl blur opacity-30 group-hover:opacity-50 transition-opacity"></div>
              
              <div class="relative bg-neutral-900 rounded-2xl overflow-hidden border border-white/10">
                <!-- Video Preview Area with Animated Action -->
                <div class="aspect-video bg-gradient-to-br from-neutral-800 to-neutral-900 relative overflow-hidden">
                  <!-- Animated Stadium Background -->
                  <div class="absolute inset-0 bg-[url('data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNDAiIGhlaWdodD0iNDAiIHZpZXdCb3g9IjAgMCA0MCA0MCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48cGF0aCBkPSJNMjAgMEw0MCAyMEwyMCA0MEwwIDIwTDIwIDBaIiBmaWxsPSJub25lIiBzdHJva2U9InJnYmEoMjU1LDI1NSwyNTUsMC4wMykiIHN0cm9rZS13aWR0aD0iMSIvPjwvc3ZnPg==')] opacity-50"></div>
                  
                  <!-- Animated Player Silhouette -->
                  <div class="absolute inset-0 flex items-center justify-center">
                    <div class="relative w-48 h-48">
                      <!-- Running player animation container -->
                      <div class="absolute inset-0 flex items-center justify-center animate-player-run">
                        <svg class="w-32 h-32 text-white/20" viewBox="0 0 100 100" fill="currentColor">
                          <ellipse cx="50" cy="85" rx="20" ry="5" class="animate-pulse opacity-30" />
                          <circle cx="50" cy="25" r="12" />
                          <path d="M50 37 L50 60 M35 45 L50 50 L65 45 M40 80 L50 60 L60 80" stroke="currentColor" stroke-width="6" stroke-linecap="round" fill="none" />
                        </svg>
                      </div>
                      
                      <!-- Animated football -->
                      <div class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 animate-ball-bounce">
                        <div class="w-10 h-10 rounded-full bg-gradient-to-br from-white to-neutral-300 shadow-lg flex items-center justify-center">
                          <svg class="w-6 h-6 text-neutral-800" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
                            <circle cx="12" cy="12" r="10"/>
                            <path d="M12 2v4M12 18v4M2 12h4M18 12h4M5.64 5.64l2.83 2.83M15.54 15.54l2.83 2.83M5.64 18.36l2.83-2.83M15.54 8.46l2.83-2.83"/>
                          </svg>
                        </div>
                      </div>
                      
                      <!-- Skill move trail effect -->
                      <div class="absolute inset-0">
                        <div class="absolute top-1/4 left-1/4 w-2 h-2 rounded-full bg-primary-400/50 animate-trail-1"></div>
                        <div class="absolute top-1/3 left-1/3 w-1.5 h-1.5 rounded-full bg-emerald-400/40 animate-trail-2"></div>
                        <div class="absolute top-1/2 left-1/2 w-1 h-1 rounded-full bg-white/30 animate-trail-3"></div>
                      </div>
                    </div>
                  </div>
                  
                  <!-- Live indicator -->
                  <div class="absolute top-4 left-4 flex items-center gap-2 px-3 py-1.5 bg-red-500/90 backdrop-blur-sm rounded-full">
                    <span class="relative flex h-2 w-2">
                      <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-white opacity-75"></span>
                      <span class="relative inline-flex rounded-full h-2 w-2 bg-white"></span>
                    </span>
                    <span class="text-white text-xs font-semibold">HIGHLIGHT REEL</span>
                  </div>
                  
                  <!-- Play Button - Opens Video Modal -->
                  <button 
                    @click="videoModal.open()"
                    class="absolute inset-0 flex items-center justify-center group/play"
                    aria-label="Play highlight reel"
                  >
                    <div class="w-20 h-20 rounded-full bg-white/10 backdrop-blur-md border border-white/20 flex items-center justify-center transition-all group-hover/play:bg-primary-500 group-hover/play:scale-110 group-hover/play:border-primary-400">
                      <svg class="w-8 h-8 text-white ml-1" fill="currentColor" viewBox="0 0 24 24">
                        <path d="M8 5v14l11-7z"/>
                      </svg>
                    </div>
                  </button>
                  
                  <!-- Video Controls Bar -->
                  <div class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/80 to-transparent p-4">
                    <div class="flex items-center gap-3">
                      <div class="w-10 h-10 rounded-full bg-white/20 flex items-center justify-center">
                        <svg class="w-5 h-5 text-white" fill="currentColor" viewBox="0 0 24 24">
                          <path d="M8 5v14l11-7z"/>
                        </svg>
                      </div>
                      <div class="flex-1">
                        <div class="h-1 bg-white/20 rounded-full overflow-hidden">
                          <div class="h-full w-1/3 bg-primary-500 rounded-full"></div>
                        </div>
                      </div>
                      <span class="text-white/60 text-sm font-mono">2:34</span>
                    </div>
                  </div>
                </div>
                
                <!-- Player Info Bar -->
                <div class="p-4 flex items-center gap-4 bg-neutral-900/80 border-t border-white/5">
                  <div class="w-12 h-12 rounded-xl bg-gradient-to-br from-primary-500 to-emerald-500 flex items-center justify-center text-white font-bold text-lg">
                    AK
                  </div>
                  <div class="flex-1">
                    <div class="text-white font-semibold">Amara Koné</div>
                    <div class="text-neutral-400 text-sm">Striker • 18 years • Ghana</div>
                  </div>
                  <div class="flex items-center gap-1 px-3 py-1.5 rounded-full bg-primary-500/20 text-primary-400 text-sm font-medium">
                    <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                      <path fill-rule="evenodd" d="M6.267 3.455a3.066 3.066 0 001.745-.723 3.066 3.066 0 013.976 0 3.066 3.066 0 001.745.723 3.066 3.066 0 012.812 2.812c.051.643.304 1.254.723 1.745a3.066 3.066 0 010 3.976 3.066 3.066 0 00-.723 1.745 3.066 3.066 0 01-2.812 2.812 3.066 3.066 0 00-1.745.723 3.066 3.066 0 01-3.976 0 3.066 3.066 0 00-1.745-.723 3.066 3.066 0 01-2.812-2.812 3.066 3.066 0 00-.723-1.745 3.066 3.066 0 010-3.976 3.066 3.066 0 00.723-1.745 3.066 3.066 0 012.812-2.812zm7.44 5.252a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                    </svg>
                    Verified
                  </div>
                </div>
              </div>
            </div>

            <!-- Floating Stats Cards -->
            <div class="absolute -left-8 top-1/4 bg-white rounded-xl p-4 shadow-2xl shadow-black/20 animate-float">
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 rounded-lg bg-primary-100 flex items-center justify-center">
                  <svg class="w-5 h-5 text-primary-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.752 11.168l-3.197-2.132A1 1 0 0010 9.87v4.263a1 1 0 001.555.832l3.197-2.132a1 1 0 000-1.664z" />
                  </svg>
                </div>
                <div>
                  <div class="text-sm font-semibold text-neutral-900">1,200+</div>
                  <div class="text-xs text-neutral-500">Video Highlights</div>
                </div>
              </div>
            </div>

            <div class="absolute -right-4 bottom-1/4 bg-white rounded-xl p-4 shadow-2xl shadow-black/20 animate-float" style="animation-delay: 1s;">
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 rounded-lg bg-secondary-100 flex items-center justify-center">
                  <svg class="w-5 h-5 text-secondary-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z" />
                  </svg>
                </div>
                <div>
                  <div class="text-sm font-semibold text-neutral-900">Top Rated</div>
                  <div class="text-xs text-neutral-500">By Scouts</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Scroll Indicator -->
      <div class="absolute bottom-8 left-1/2 -translate-x-1/2 animate-bounce">
        <svg class="w-6 h-6 text-neutral-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14l-7 7m0 0l-7-7m7 7V3" />
        </svg>
      </div>
    </section>

    <!-- Floating Stats Bar -->
    <section class="relative -mt-16 z-20">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="bg-white rounded-2xl shadow-2xl shadow-neutral-900/10 border border-neutral-200/50 p-8 grid grid-cols-2 md:grid-cols-4 gap-8">
          <div class="text-center" v-for="stat in statsData" :key="stat.label">
            <div class="font-display text-4xl lg:text-5xl font-bold text-neutral-900">
              {{ stat.value }}<span class="text-primary-500">+</span>
            </div>
            <div class="mt-2 text-neutral-500 font-medium">{{ stat.label }}</div>
          </div>
        </div>
      </div>
    </section>

    <!-- Featured Players - Horizontal Scroll Showcase -->
    <section class="py-24 lg:py-32 bg-white relative overflow-hidden">
      <!-- Background decoration -->
      <div class="absolute top-0 right-0 w-1/3 h-full bg-gradient-to-l from-primary-50 to-transparent"></div>
      
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 relative">
        <!-- Section Header with Line -->
        <div class="flex items-center gap-8 mb-16">
          <div class="flex-1">
            <div class="flex items-center gap-4 mb-4">
              <div class="w-12 h-12 rounded-2xl bg-gradient-to-br from-primary-500 to-emerald-500 flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z" />
                </svg>
              </div>
              <h2 class="font-display text-4xl lg:text-5xl font-bold text-neutral-900">
                Rising Stars
              </h2>
            </div>
            <p class="text-lg text-neutral-500 max-w-xl">
              Hand-picked by our scouts. Verified by academies. Ready for their breakthrough.
            </p>
          </div>
          <NuxtLink 
            to="/discover" 
            class="hidden lg:flex items-center gap-3 px-6 py-3 bg-neutral-900 text-white rounded-full font-semibold hover:bg-neutral-800 transition-all group"
          >
            <span>Explore All</span>
            <svg class="w-5 h-5 transition-transform group-hover:translate-x-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3" />
            </svg>
          </NuxtLink>
        </div>

        <!-- Loading State -->
        <div v-if="loading" class="flex justify-center py-20">
          <div class="relative">
            <div class="w-20 h-20 border-4 border-neutral-200 rounded-full"></div>
            <div class="w-20 h-20 border-4 border-primary-500 rounded-full animate-spin absolute inset-0 border-t-transparent"></div>
          </div>
        </div>

        <!-- Players Showcase - Magazine Style -->
        <div v-else-if="featuredPlayers.length > 0" class="grid grid-cols-12 gap-6">
          <!-- Large Featured Card -->
          <div class="col-span-12 lg:col-span-7 group">
            <NuxtLink :to="`/players/${featuredPlayers[0]?.id}`" class="block relative h-[500px] lg:h-[600px] rounded-3xl overflow-hidden">
              <img 
                v-if="featuredPlayers[0]?.profile_photo_url"
                :src="featuredPlayers[0].profile_photo_url"
                :alt="featuredPlayers[0].first_name"
                class="w-full h-full object-cover transition-transform duration-700 group-hover:scale-105"
              />
              <div v-else class="w-full h-full bg-gradient-to-br from-primary-400 to-emerald-500"></div>
              
              <!-- Overlay -->
              <div class="absolute inset-0 bg-gradient-to-t from-black/80 via-black/20 to-transparent"></div>
              
              <!-- Content -->
              <div class="absolute inset-0 p-8 flex flex-col justify-between">
                <!-- Top badges -->
                <div class="flex items-start justify-between">
                  <div v-if="featuredPlayers[0]?.is_verified" class="flex items-center gap-2 px-4 py-2 bg-white/20 backdrop-blur-md rounded-full border border-white/30">
                    <svg class="w-4 h-4 text-primary-400" fill="currentColor" viewBox="0 0 20 20">
                      <path fill-rule="evenodd" d="M6.267 3.455a3.066 3.066 0 001.745-.723 3.066 3.066 0 013.976 0 3.066 3.066 0 001.745.723 3.066 3.066 0 012.812 2.812c.051.643.304 1.254.723 1.745a3.066 3.066 0 010 3.976 3.066 3.066 0 00-.723 1.745 3.066 3.066 0 01-2.812 2.812 3.066 3.066 0 00-1.745.723 3.066 3.066 0 01-3.976 0 3.066 3.066 0 00-1.745-.723 3.066 3.066 0 01-2.812-2.812 3.066 3.066 0 00-.723-1.745 3.066 3.066 0 010-3.976 3.066 3.066 0 00.723-1.745 3.066 3.066 0 012.812-2.812zm7.44 5.252a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                    </svg>
                    <span class="text-white text-sm font-medium">Verified Talent</span>
                  </div>
                  <span class="px-4 py-2 bg-secondary-500 text-neutral-900 rounded-full text-sm font-bold">
                    Featured
                  </span>
                </div>
                
                <!-- Bottom info -->
                <div>
                  <div class="flex flex-wrap gap-2 mb-4">
                    <span class="px-3 py-1 bg-white/20 backdrop-blur-md rounded-full text-white text-sm border border-white/20">
                      {{ featuredPlayers[0]?.position }}
                    </span>
                    <span class="px-3 py-1 bg-white/20 backdrop-blur-md rounded-full text-white text-sm border border-white/20">
                      {{ calculateAge(featuredPlayers[0]?.date_of_birth) }} years
                    </span>
                    <span class="px-3 py-1 bg-white/20 backdrop-blur-md rounded-full text-white text-sm border border-white/20">
                      {{ featuredPlayers[0]?.country }}
                    </span>
                  </div>
                  <h3 class="text-4xl lg:text-5xl font-bold text-white mb-2">
                    {{ featuredPlayers[0]?.first_name }} {{ featuredPlayers[0]?.last_name }}
                  </h3>
                  <p v-if="featuredPlayers[0]?.academy_name" class="text-neutral-300 text-lg">
                    {{ featuredPlayers[0]?.academy_name }}
                  </p>
                </div>
              </div>
            </NuxtLink>
          </div>

          <!-- Right side smaller cards -->
          <div class="col-span-12 lg:col-span-5 grid grid-cols-2 gap-4">
            <div 
              v-for="(player, index) in featuredPlayers.slice(1, 5)" 
              :key="player.id"
              class="group relative rounded-2xl overflow-hidden"
              :class="index < 2 ? 'h-[200px] lg:h-[290px]' : 'h-[200px] lg:h-[290px]'"
            >
              <NuxtLink :to="`/players/${player.id}`" class="block h-full">
                <img 
                  v-if="player.profile_photo_url"
                  :src="player.profile_photo_url"
                  :alt="player.first_name"
                  class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-110"
                />
                <div v-else class="w-full h-full bg-gradient-to-br from-neutral-200 to-neutral-300 flex items-center justify-center">
                  <span class="text-4xl font-bold text-neutral-400">{{ player.first_name?.charAt(0) }}</span>
                </div>
                
                <!-- Overlay -->
                <div class="absolute inset-0 bg-gradient-to-t from-black/70 via-transparent to-transparent opacity-80 group-hover:opacity-100 transition-opacity"></div>
                
                <!-- Info -->
                <div class="absolute bottom-0 left-0 right-0 p-4">
                  <div class="text-xs text-primary-400 font-semibold mb-1">{{ player.position }}</div>
                  <h4 class="text-white font-bold text-lg leading-tight">{{ player.first_name }} {{ player.last_name }}</h4>
                  <p class="text-neutral-400 text-sm">{{ player.country }}</p>
                </div>
                
                <!-- Hover Arrow -->
                <div class="absolute top-4 right-4 w-10 h-10 rounded-full bg-white/20 backdrop-blur-sm flex items-center justify-center opacity-0 group-hover:opacity-100 transition-all transform translate-x-2 group-hover:translate-x-0">
                  <svg class="w-5 h-5 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3" />
                  </svg>
                </div>
              </NuxtLink>
            </div>
          </div>
        </div>

        <!-- Empty State -->
        <div v-else class="text-center py-20">
          <div class="w-24 h-24 mx-auto mb-6 rounded-3xl bg-gradient-to-br from-neutral-100 to-neutral-200 flex items-center justify-center">
            <svg class="w-12 h-12 text-neutral-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
            </svg>
          </div>
          <h3 class="text-2xl font-bold text-neutral-900 mb-2">Talent Coming Soon</h3>
          <p class="text-neutral-500">Our scouts are discovering amazing players. Check back soon!</p>
        </div>

        <!-- Mobile CTA -->
        <div class="lg:hidden mt-8 text-center">
          <NuxtLink 
            to="/discover" 
            class="inline-flex items-center gap-3 px-8 py-4 bg-neutral-900 text-white rounded-full font-semibold"
          >
            <span>Explore All Players</span>
            <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3" />
            </svg>
          </NuxtLink>
        </div>
      </div>
    </section>

    <!-- Featured Highlights Section -->
    <section v-if="featuredHighlights.length > 0" class="py-20 lg:py-28 bg-neutral-950 relative overflow-hidden">
      <!-- Background Elements -->
      <div class="absolute inset-0">
        <div class="absolute top-1/4 right-0 w-96 h-96 bg-emerald-500/10 rounded-full blur-3xl"></div>
        <div class="absolute bottom-1/4 left-0 w-72 h-72 bg-primary-500/10 rounded-full blur-3xl"></div>
      </div>

      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 relative">
        <!-- Header -->
        <div class="flex items-end justify-between mb-12">
          <div>
            <div class="inline-flex items-center gap-2 px-4 py-2 rounded-full bg-emerald-500/10 border border-emerald-500/20 mb-4">
              <span class="w-2 h-2 rounded-full bg-emerald-500 animate-pulse"></span>
              <span class="text-sm text-emerald-400 font-medium">Free to Watch</span>
            </div>
            <h2 class="font-display text-3xl lg:text-5xl font-bold text-white">
              Latest <span class="text-transparent bg-clip-text bg-gradient-to-r from-emerald-400 to-primary-400">Highlights</span>
            </h2>
            <p class="mt-3 text-neutral-400 max-w-xl">
              Watch the best moments from our verified players - goals, skills, saves and more
            </p>
          </div>
          <NuxtLink 
            to="/discover"
            class="hidden lg:inline-flex items-center gap-2 text-emerald-400 hover:text-emerald-300 font-medium transition-colors"
          >
            View All Players
            <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3" />
            </svg>
          </NuxtLink>
        </div>

        <!-- Highlights Grid -->
        <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 lg:gap-6">
          <div
            v-for="highlight in featuredHighlights.slice(0, 8)"
            :key="highlight.id"
            class="group relative rounded-2xl overflow-hidden cursor-pointer bg-neutral-900"
            @click="openHighlight(highlight)"
          >
            <!-- Thumbnail -->
            <div class="aspect-video relative">
              <img
                v-if="highlight.thumbnail_url"
                :src="highlight.thumbnail_url"
                :alt="highlight.title || 'Highlight'"
                class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-110"
              />
              <div v-else class="w-full h-full bg-gradient-to-br from-neutral-800 to-neutral-900 flex items-center justify-center">
                <span class="text-4xl">{{ getHighlightIcon(highlight.highlight_type) }}</span>
              </div>

              <!-- Play Button Overlay -->
              <div class="absolute inset-0 bg-black/40 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity">
                <div class="w-14 h-14 rounded-full bg-white/20 backdrop-blur-sm flex items-center justify-center border border-white/30">
                  <svg class="w-6 h-6 text-white ml-1" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M8 5v14l11-7z" />
                  </svg>
                </div>
              </div>

              <!-- Duration Badge -->
              <div v-if="highlight.duration_seconds" class="absolute bottom-2 right-2 px-2 py-1 bg-black/70 rounded text-white text-xs font-mono">
                {{ formatDuration(highlight.duration_seconds) }}
              </div>

              <!-- Type Badge -->
              <div class="absolute top-2 left-2 px-2 py-1 bg-emerald-500/90 rounded-lg text-white text-xs font-medium flex items-center gap-1">
                <span>{{ getHighlightIcon(highlight.highlight_type) }}</span>
                <span class="capitalize">{{ highlight.highlight_type }}</span>
              </div>
            </div>

            <!-- Info -->
            <div class="p-3">
              <div class="flex items-center gap-2 mb-1">
                <div class="w-6 h-6 rounded-full bg-gradient-to-br from-primary-500 to-emerald-500 flex items-center justify-center text-white text-xs font-bold">
                  {{ highlight.player?.first_name?.charAt(0) || '?' }}
                </div>
                <div class="flex-1 min-w-0">
                  <p class="text-white text-sm font-medium truncate">
                    {{ highlight.player?.first_name }} {{ highlight.player?.last_name }}
                  </p>
                </div>
              </div>
              <p class="text-neutral-500 text-xs truncate">
                {{ highlight.player?.position }}
              </p>
            </div>
          </div>
        </div>

        <!-- Mobile CTA -->
        <div class="lg:hidden mt-8 text-center">
          <NuxtLink 
            to="/discover"
            class="inline-flex items-center gap-2 text-emerald-400 font-medium"
          >
            View All Players
            <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3" />
            </svg>
          </NuxtLink>
        </div>
      </div>
    </section>

    <!-- Highlight Video Modal -->
    <Teleport to="body">
      <div v-if="showHighlightModal && selectedHighlight" class="fixed inset-0 z-50 overflow-y-auto">
        <div class="fixed inset-0 bg-black/90" @click="showHighlightModal = false"></div>
        <div class="relative min-h-screen flex items-center justify-center p-4">
          <div class="relative bg-black rounded-2xl shadow-xl w-full max-w-4xl overflow-hidden">
            <!-- Close button -->
            <button
              @click="showHighlightModal = false"
              class="absolute top-4 right-4 z-10 p-2 bg-black/50 hover:bg-black/70 text-white rounded-full transition-colors"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>

            <!-- Video Player -->
            <div class="aspect-video bg-black">
              <video
                v-if="selectedHighlight.stream_url"
                :src="selectedHighlight.stream_url"
                controls
                autoplay
                class="w-full h-full"
              >
                Your browser does not support the video tag.
              </video>
            </div>

            <!-- Highlight Info -->
            <div class="p-4 bg-neutral-900">
              <div class="flex items-center gap-4">
                <div class="w-12 h-12 rounded-xl bg-gradient-to-br from-primary-500 to-emerald-500 flex items-center justify-center text-white font-bold text-lg">
                  {{ selectedHighlight.player?.first_name?.charAt(0) }}{{ selectedHighlight.player?.last_name?.charAt(0) }}
                </div>
                <div class="flex-1">
                  <div class="flex items-center gap-2 mb-1">
                    <span class="text-lg font-semibold text-white">
                      {{ selectedHighlight.player?.first_name }} {{ selectedHighlight.player?.last_name }}
                    </span>
                    <span class="px-2 py-0.5 bg-emerald-500/20 text-emerald-400 rounded text-sm">
                      {{ getHighlightIcon(selectedHighlight.highlight_type) }} {{ selectedHighlight.highlight_type }}
                    </span>
                  </div>
                  <p class="text-neutral-400 text-sm">
                    {{ selectedHighlight.player?.position }}
                    <span v-if="selectedHighlight.match"> • {{ selectedHighlight.match.title }}</span>
                  </p>
                </div>
                <NuxtLink
                  v-if="selectedHighlight.player?.id"
                  :to="`/players/${selectedHighlight.player.id}`"
                  class="px-4 py-2 bg-primary-500 hover:bg-primary-600 text-white rounded-lg text-sm font-medium transition-colors"
                  @click="showHighlightModal = false"
                >
                  View Profile
                </NuxtLink>
              </div>
            </div>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- How It Works - Timeline Style -->
    <section class="py-24 lg:py-32 bg-neutral-950 text-white relative overflow-hidden">
      <!-- Background Elements -->
      <div class="absolute inset-0">
        <div class="absolute top-1/4 left-0 w-72 h-72 bg-primary-500/10 rounded-full blur-3xl"></div>
        <div class="absolute bottom-1/4 right-0 w-96 h-96 bg-emerald-500/10 rounded-full blur-3xl"></div>
      </div>
      
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 relative">
        <!-- Header -->
        <div class="text-center mb-20">
          <div class="inline-flex items-center gap-2 px-4 py-2 rounded-full bg-white/5 border border-white/10 mb-6">
            <span class="w-2 h-2 rounded-full bg-primary-500"></span>
            <span class="text-sm text-neutral-400">Simple 3-Step Process</span>
          </div>
          <h2 class="font-display text-4xl lg:text-6xl font-bold">
            How It <span class="text-transparent bg-clip-text bg-gradient-to-r from-primary-400 to-emerald-400">Works</span>
          </h2>
        </div>

        <!-- Steps - Horizontal on Desktop -->
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-8 lg:gap-4">
          <div 
            v-for="(step, index) in steps" 
            :key="index"
            class="relative group"
          >
            <!-- Connector Line (Desktop) -->
            <div v-if="index < 2" class="hidden lg:block absolute top-16 left-1/2 w-full h-px bg-gradient-to-r from-primary-500/50 to-transparent"></div>
            
            <!-- Card -->
            <div class="relative bg-white/5 backdrop-blur-sm border border-white/10 rounded-3xl p-8 hover:bg-white/10 hover:border-white/20 transition-all duration-500 group-hover:-translate-y-2">
              <!-- Number Badge -->
              <div class="absolute -top-5 left-8 w-10 h-10 rounded-2xl bg-gradient-to-br from-primary-500 to-emerald-500 flex items-center justify-center font-bold text-lg shadow-lg shadow-primary-500/25">
                {{ index + 1 }}
              </div>
              
              <!-- Icon Container -->
              <div class="w-16 h-16 rounded-2xl bg-gradient-to-br from-white/10 to-white/5 border border-white/10 flex items-center justify-center mb-6 mt-4 group-hover:scale-110 transition-transform duration-300">
                <svg class="w-8 h-8 text-primary-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" :d="step.iconPath" />
                </svg>
              </div>
              
              <h3 class="text-2xl font-bold text-white mb-3">{{ step.title }}</h3>
              <p class="text-neutral-400 leading-relaxed">{{ step.description }}</p>
              
              <!-- Decorative -->
              <div class="absolute bottom-0 right-0 w-32 h-32 bg-gradient-to-tl from-primary-500/10 to-transparent rounded-tl-full opacity-0 group-hover:opacity-100 transition-opacity"></div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Testimonial / Social Proof - Modern Card Style -->
    <section class="py-24 lg:py-32 bg-neutral-50 relative overflow-hidden">
      <!-- Background -->
      <div class="absolute inset-0">
        <div class="absolute top-0 left-1/2 -translate-x-1/2 w-full max-w-4xl h-full bg-gradient-to-b from-primary-100/30 via-transparent to-transparent"></div>
      </div>
      
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 relative">
        <!-- Section Header -->
        <div class="text-center mb-16">
          <div class="inline-flex items-center gap-2 px-4 py-2 rounded-full bg-white border border-neutral-200 shadow-sm mb-6">
            <svg class="w-4 h-4 text-secondary-500" fill="currentColor" viewBox="0 0 20 20">
              <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
            </svg>
            <span class="text-sm font-semibold text-neutral-700">What Scouts Say</span>
          </div>
          <h2 class="font-display text-4xl lg:text-5xl font-bold text-neutral-900">
            Trusted by <span class="text-transparent bg-clip-text bg-gradient-to-r from-primary-600 to-emerald-500">Professionals</span>
          </h2>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-12 gap-8">
          <!-- Rotating Testimonial Card -->
          <div class="lg:col-span-7 relative">
            <div class="bg-white rounded-3xl p-8 lg:p-12 shadow-xl shadow-neutral-900/5 border border-neutral-100 relative overflow-hidden min-h-[400px]">
              <!-- Quote mark decoration -->
              <div class="absolute -top-4 -left-4 w-24 h-24 bg-gradient-to-br from-primary-500/10 to-transparent rounded-full"></div>
              <svg class="w-16 h-16 text-primary-200 mb-6" fill="currentColor" viewBox="0 0 24 24">
                <path d="M14.017 21v-7.391c0-5.704 3.731-9.57 8.983-10.609l.995 2.151c-2.432.917-3.995 3.638-3.995 5.849h4v10h-9.983zm-14.017 0v-7.391c0-5.704 3.748-9.57 9-10.609l.996 2.151c-2.433.917-3.996 3.638-3.996 5.849h3.983v10h-9.983z"/>
              </svg>
              
              <!-- Testimonial Content with Transition -->
              <Transition name="testimonial-fade" mode="out-in">
                <div v-if="currentTestimonial" :key="currentTestimonialIndex">
                  <blockquote class="text-2xl lg:text-3xl font-medium leading-relaxed text-neutral-800 mb-8">
                    "{{ currentTestimonial.quote }}"
                  </blockquote>
                  
                  <div class="flex items-center gap-4">
                    <div class="w-14 h-14 rounded-2xl overflow-hidden" :class="currentTestimonial.avatarBg">
                      <img 
                        v-if="currentTestimonial.avatar" 
                        :src="currentTestimonial.avatar" 
                        :alt="currentTestimonial.name"
                        class="w-full h-full object-cover"
                      />
                      <div v-else class="w-full h-full flex items-center justify-center bg-gradient-to-br from-primary-500 to-emerald-500">
                        <span class="text-xl font-bold text-white">{{ currentTestimonial.initials }}</span>
                      </div>
                    </div>
                    <div>
                      <div class="font-bold text-neutral-900 text-lg">{{ currentTestimonial.name }}</div>
                      <div class="text-neutral-500">{{ currentTestimonial.role }}</div>
                    </div>
                  </div>
                </div>
              </Transition>
              
              <!-- 5 stars -->
              <div class="absolute top-8 right-8 flex gap-1">
                <svg v-for="i in 5" :key="i" class="w-5 h-5 text-secondary-500" fill="currentColor" viewBox="0 0 20 20">
                  <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                </svg>
              </div>
              
              <!-- Testimonial Navigation Dots -->
              <div class="absolute bottom-8 left-8 flex items-center gap-3">
                <button 
                  v-for="(_, index) in testimonials" 
                  :key="index"
                  @click="currentTestimonialIndex = index"
                  class="w-2.5 h-2.5 rounded-full transition-all duration-300"
                  :class="index === currentTestimonialIndex ? 'bg-primary-500 w-8' : 'bg-neutral-300 hover:bg-neutral-400'"
                  :aria-label="`View testimonial ${index + 1}`"
                />
              </div>
            </div>
          </div>
          
          <!-- Stats Cards -->
          <div class="lg:col-span-5 grid grid-cols-2 gap-4">
            <div 
              v-for="(feature, index) in features" 
              :key="feature.label"
              class="group p-6 rounded-2xl bg-white border border-neutral-200 shadow-sm hover:shadow-lg hover:border-primary-200 transition-all duration-300 hover:-translate-y-1"
              :class="index === 0 ? 'bg-gradient-to-br from-primary-500 to-emerald-500 border-0' : ''"
            >
              <div 
                :class="[
                  'text-4xl font-bold mb-2',
                  index === 0 ? 'text-white' : 'text-neutral-900'
                ]"
              >
                {{ feature.value }}
              </div>
              <div 
                :class="[
                  'text-sm font-medium',
                  index === 0 ? 'text-white/80' : 'text-neutral-500'
                ]"
              >
                {{ feature.label }}
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- CTA Section - Immersive Full-Width -->
    <section class="relative py-32 lg:py-40 bg-neutral-900 overflow-hidden">
      <!-- Animated Background -->
      <div class="absolute inset-0">
        <div class="absolute inset-0 bg-[radial-gradient(ellipse_at_top_right,_var(--tw-gradient-stops))] from-primary-900/40 via-neutral-900 to-neutral-900"></div>
        <div class="absolute top-0 left-1/4 w-[600px] h-[600px] bg-primary-500/20 rounded-full blur-[120px] animate-pulse"></div>
        <div class="absolute bottom-0 right-1/4 w-[500px] h-[500px] bg-emerald-500/20 rounded-full blur-[100px] animate-pulse" style="animation-delay: 1s;"></div>
        <!-- Grid Pattern -->
        <div class="absolute inset-0 bg-[linear-gradient(rgba(255,255,255,.02)_1px,transparent_1px),linear-gradient(90deg,rgba(255,255,255,.02)_1px,transparent_1px)] bg-[size:64px_64px]"></div>
      </div>
      
      <div class="relative max-w-5xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <!-- Badge with Urgency -->
        <div class="inline-flex items-center gap-3 px-5 py-2.5 rounded-full bg-white/5 border border-white/10 backdrop-blur-sm mb-8">
          <span class="relative flex h-2 w-2">
            <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-primary-400 opacity-75"></span>
            <span class="relative inline-flex rounded-full h-2 w-2 bg-primary-500"></span>
          </span>
          <span class="text-sm text-neutral-300">Join 200+ scouts</span>
          <span class="text-xs px-2 py-0.5 bg-secondary-500/20 text-secondary-400 rounded-full font-medium">{{ weeklySignups }} signed up this week</span>
        </div>
        
        <h2 class="font-display text-5xl lg:text-7xl font-bold text-white leading-tight mb-6">
          Ready to Discover<br/>
          <span class="text-transparent bg-clip-text bg-gradient-to-r from-primary-400 via-emerald-400 to-primary-400 bg-[length:200%_auto] animate-gradient">
            the Next Superstar?
          </span>
        </h2>
        
        <p class="text-xl lg:text-2xl text-neutral-400 max-w-2xl mx-auto mb-12 leading-relaxed">
          The platform trusted by clubs and scouts worldwide to find Africa's brightest football talent.
        </p>
        
        <div class="flex flex-col sm:flex-row gap-4 justify-center items-center">
          <NuxtLink 
            to="/auth/register"
            class="group relative px-10 py-5 bg-white text-neutral-900 font-bold text-lg rounded-2xl transition-all duration-300 hover:shadow-2xl hover:shadow-white/20 hover:scale-105 overflow-hidden"
          >
            <span class="relative z-10 flex items-center gap-3">
              Get Started Free
              <svg class="w-5 h-5 transition-transform group-hover:translate-x-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3" />
              </svg>
            </span>
          </NuxtLink>
          
          <NuxtLink 
            to="/discover"
            class="group flex items-center gap-3 px-10 py-5 text-white font-semibold text-lg rounded-2xl border border-white/20 bg-white/5 backdrop-blur-sm transition-all duration-300 hover:bg-white/10 hover:border-white/30"
          >
            <svg class="w-6 h-6 text-primary-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.752 11.168l-3.197-2.132A1 1 0 0010 9.87v4.263a1 1 0 001.555.832l3.197-2.132a1 1 0 000-1.664zM21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            Browse Players
          </NuxtLink>
        </div>
        
        <!-- Trust Indicators -->
        <div class="mt-16 pt-12 border-t border-white/10">
          <p class="text-neutral-500 text-sm mb-6">Trusted by professionals from</p>
          <div class="flex flex-wrap justify-center items-center gap-8 lg:gap-16 opacity-40">
            <div class="text-white font-bold text-xl tracking-wider">PREMIER LEAGUE</div>
            <div class="text-white font-bold text-xl tracking-wider">LA LIGA</div>
            <div class="text-white font-bold text-xl tracking-wider">BUNDESLIGA</div>
            <div class="text-white font-bold text-xl tracking-wider">SERIE A</div>
          </div>
        </div>
      </div>
    </section>

    <!-- Across Africa - Interactive Map Style -->
    <section class="py-24 lg:py-32 bg-white relative overflow-hidden">
      <!-- Subtle background decoration -->
      <div class="absolute top-0 left-0 w-full h-px bg-gradient-to-r from-transparent via-primary-200 to-transparent"></div>
      
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-16 items-center">
          <!-- Left: Content -->
          <div>
            <div class="inline-flex items-center gap-2 px-4 py-2 rounded-full bg-primary-50 border border-primary-200 mb-6">
              <svg class="w-4 h-4 text-primary-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.055 11H5a2 2 0 012 2v1a2 2 0 002 2 2 2 0 012 2v2.945M8 3.935V5.5A2.5 2.5 0 0010.5 8h.5a2 2 0 012 2 2 2 0 104 0 2 2 0 012-2h1.064M15 20.488V18a2 2 0 012-2h3.064M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
              <span class="text-sm font-semibold text-primary-700">Continental Coverage</span>
            </div>
            
            <h2 class="font-display text-4xl lg:text-5xl font-bold text-neutral-900 mb-6">
              Talent From<br/>
              <span class="text-transparent bg-clip-text bg-gradient-to-r from-primary-600 to-emerald-500">
                Across Africa
              </span>
            </h2>
            
            <p class="text-lg text-neutral-600 mb-8 leading-relaxed">
              Our network spans the entire continent, partnering with 50+ academies in 15 countries to bring you verified, professional-grade talent.
            </p>
            
            <!-- Stats Row -->
            <div class="grid grid-cols-3 gap-6 mb-8">
              <div class="text-center p-4 rounded-2xl bg-neutral-50">
                <div class="text-3xl font-bold text-primary-600 mb-1">15+</div>
                <div class="text-sm text-neutral-500">Countries</div>
              </div>
              <div class="text-center p-4 rounded-2xl bg-neutral-50">
                <div class="text-3xl font-bold text-primary-600 mb-1">50+</div>
                <div class="text-sm text-neutral-500">Academies</div>
              </div>
              <div class="text-center p-4 rounded-2xl bg-neutral-50">
                <div class="text-3xl font-bold text-primary-600 mb-1">500+</div>
                <div class="text-sm text-neutral-500">Players</div>
              </div>
            </div>
            
            <NuxtLink 
              to="/discover" 
              class="inline-flex items-center gap-2 text-primary-600 font-semibold hover:text-primary-700 transition-colors group"
            >
              View Partner Academies
              <svg class="w-5 h-5 transition-transform group-hover:translate-x-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3" />
              </svg>
            </NuxtLink>
          </div>
          
          <!-- Right: Country Grid -->
          <div class="relative">
            <!-- Background glow -->
            <div class="absolute inset-0 bg-gradient-to-br from-primary-100/50 to-emerald-100/50 rounded-3xl blur-3xl scale-110"></div>
            
            <div class="relative bg-white rounded-3xl border border-neutral-200 p-8 shadow-xl shadow-neutral-900/5">
              <div class="grid grid-cols-3 gap-3">
                <NuxtLink 
                  v-for="(country, index) in countries" 
                  :key="country.name"
                  :to="`/discover?country=${encodeURIComponent(country.name)}`"
                  class="group relative overflow-hidden rounded-2xl transition-all duration-300 cursor-pointer"
                  :class="[
                    index === 0 ? 'col-span-2 row-span-2 bg-gradient-to-br from-primary-500 to-emerald-500 p-6 hover:from-primary-600 hover:to-emerald-600' : 'bg-neutral-50 hover:bg-primary-50 p-4 border border-neutral-100 hover:border-primary-300 hover:shadow-md'
                  ]"
                >
                  <!-- Country with flag -->
                  <div class="flex items-center gap-2">
                    <span class="text-xl">{{ country.flag }}</span>
                    <span 
                      :class="[
                        'font-bold transition-colors',
                        index === 0 ? 'text-white text-2xl' : 'text-neutral-700 group-hover:text-primary-700 text-sm'
                      ]"
                    >
                      {{ country.name }}
                    </span>
                  </div>
                  
                  <!-- Player count badge -->
                  <div 
                    v-if="country.playerCount"
                    :class="[
                      'mt-2 text-xs font-medium',
                      index === 0 ? 'text-white/80' : 'text-neutral-500 group-hover:text-primary-600'
                    ]"
                  >
                    {{ country.playerCount }} players
                  </div>
                  
                  <!-- Arrow on hover -->
                  <div 
                    class="absolute top-2 right-2 w-6 h-6 rounded-full flex items-center justify-center opacity-0 group-hover:opacity-100 transition-all transform translate-x-2 group-hover:translate-x-0"
                    :class="index === 0 ? 'bg-white/20' : 'bg-primary-100'"
                  >
                    <svg :class="index === 0 ? 'text-white' : 'text-primary-600'" class="w-3 h-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3" />
                    </svg>
                  </div>
                  
                  <div 
                    v-if="index === 0"
                    class="absolute bottom-4 right-4 flex items-center gap-2 text-white/80 text-sm"
                  >
                    <span>Most players</span>
                    <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 10l7-7m0 0l7 7m-7-7v18" />
                    </svg>
                  </div>
                </NuxtLink>
              </div>
              
              <!-- More countries indicator -->
              <div class="mt-4 pt-4 border-t border-neutral-100 flex items-center justify-between">
                <span class="text-neutral-500 text-sm">And more countries coming soon...</span>
                <div class="flex -space-x-2">
                  <div class="w-8 h-8 rounded-full bg-neutral-200 border-2 border-white flex items-center justify-center text-xs font-bold text-neutral-500">+5</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Video Player Modal -->
    <UiVideoPlayerModal
      v-model="videoModal.isOpen.value"
      :video-url="featuredVideo?.streamUrl"
      :youtube-id="featuredVideo?.youtubeId"
      :thumbnail-url="featuredVideo?.thumbnailUrl"
      :title="featuredVideo?.title"
      :player="featuredVideoPlayer"
      :show-info="true"
      :allow-demo="true"
      :load-timeout="8000"
    />
  </div>
</template>

<script setup lang="ts">
import type { Player, ApiResponse } from '~/types'

interface Testimonial {
  quote: string
  name: string
  role: string
  initials: string
  avatar: string | undefined
  avatarBg: string
}

interface Country {
  name: string
  flag: string
  playerCount: number
}

interface FeaturedVideo {
  id: string
  title: string
  streamUrl?: string
  youtubeId?: string
  thumbnailUrl?: string
  durationSeconds?: number
}

interface FeaturedVideoPlayer {
  id?: string
  name: string
  position?: string
  country?: string
  avatarUrl?: string
}

definePageMeta({
  layout: 'default',
})

const config = useRuntimeConfig()
const loading = ref(true)
const featuredPlayers = ref<Player[]>([])
const featuredHighlights = ref<FeaturedHighlight[]>([])
const selectedHighlight = ref<FeaturedHighlight | null>(null)
const showHighlightModal = ref(false)

// Highlight type interface
interface FeaturedHighlight {
  id: string
  highlight_type: string
  title?: string
  thumbnail_url?: string
  stream_url: string
  duration_seconds?: number
  view_count: number
  player?: {
    id: string
    first_name: string
    last_name: string
    position: string
  }
  match?: {
    id: string
    title: string
  }
}

// Video modal state using the composable
const videoModal = useModal()

// Featured video for homepage hero
const featuredVideo = ref<FeaturedVideo | null>(null)
const featuredVideoPlayer = computed<FeaturedVideoPlayer | undefined>(() => {
  // Use the first featured player as the video player info
  const player = featuredPlayers.value[0]
  if (!player) return undefined
  return {
    id: player.id,
    name: `${player.first_name} ${player.last_name}`,
    position: player.position,
    country: player.country,
    avatarUrl: player.profile_photo_url,
  }
})

const statsData = [
  { value: '500', label: 'Verified Players' },
  { value: '15', label: 'Countries' },
  { value: '200', label: 'Active Scouts' },
  { value: '50', label: 'Academy Partners' },
]

const features = [
  { value: '100%', label: 'Verified Profiles' },
  { value: '4K', label: 'Video Quality' },
  { value: '24hr', label: 'Response Time' },
  { value: '∞', label: 'Player Access' },
]

const countries: Country[] = [
  { name: 'Nigeria', flag: '🇳🇬', playerCount: 120 },
  { name: 'Ghana', flag: '🇬🇭', playerCount: 85 },
  { name: 'Senegal', flag: '🇸🇳', playerCount: 72 },
  { name: 'Cameroon', flag: '🇨🇲', playerCount: 58 },
  { name: 'Egypt', flag: '🇪🇬', playerCount: 45 },
  { name: 'Morocco', flag: '🇲🇦', playerCount: 42 },
  { name: 'South Africa', flag: '🇿🇦', playerCount: 38 },
  { name: 'Kenya', flag: '🇰🇪', playerCount: 25 },
  { name: 'Ivory Coast', flag: '🇨🇮', playerCount: 35 },
  { name: 'Mali', flag: '🇲🇱', playerCount: 28 },
  { name: 'DR Congo', flag: '🇨🇩', playerCount: 22 },
  { name: 'Algeria', flag: '🇩🇿', playerCount: 18 },
  { name: 'Tunisia', flag: '🇹🇳', playerCount: 15 },
  { name: 'Uganda', flag: '🇺🇬', playerCount: 12 },
  { name: 'Tanzania', flag: '🇹🇿', playerCount: 10 }
]

// Testimonials data with rotation
const testimonials: Testimonial[] = [
  {
    quote: "Unicorn Sport has revolutionized how we scout African talent. The verified profiles and HD highlight reels save us countless hours of research.",
    name: "James Martinez",
    role: "Head Scout, European Club",
    initials: "JM",
    avatar: undefined,
    avatarBg: "bg-gradient-to-br from-primary-500 to-emerald-500"
  },
  {
    quote: "The quality of players on this platform is exceptional. We've signed three talents in the past year, all discovered through Unicorn Sport.",
    name: "Sarah Williams",
    role: "Recruitment Director, Premier League Club",
    initials: "SW",
    avatar: undefined,
    avatarBg: "bg-gradient-to-br from-purple-500 to-pink-500"
  },
  {
    quote: "Finally, a platform that understands African football. The verification process gives us confidence in every profile we review.",
    name: "Mohamed Al-Hassan",
    role: "Technical Director, Middle East Academy",
    initials: "MA",
    avatar: undefined,
    avatarBg: "bg-gradient-to-br from-amber-500 to-orange-500"
  }
]

const currentTestimonialIndex = ref(0)
const weeklySignups = ref(47)

// Computed property for current testimonial to avoid undefined access
const currentTestimonial = computed(() => testimonials[currentTestimonialIndex.value] as Testimonial)

// Auto-rotate testimonials
let testimonialInterval: ReturnType<typeof setInterval> | null = null
onMounted(() => {
  testimonialInterval = setInterval(() => {
    currentTestimonialIndex.value = (currentTestimonialIndex.value + 1) % testimonials.length
  }, 5000)
})

onUnmounted(() => {
  if (testimonialInterval) clearInterval(testimonialInterval)
})

const steps = [
  {
    iconPath: 'M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z',
    title: 'Browse & Discover',
    description: 'Search through hundreds of verified player profiles. Filter by position, age, country, and more to find your perfect prospect.'
  },
  {
    iconPath: 'M14.752 11.168l-3.197-2.132A1 1 0 0010 9.87v4.263a1 1 0 001.555.832l3.197-2.132a1 1 0 000-1.664zM21 12a9 9 0 11-18 0 9 9 0 0118 0z',
    title: 'Watch Highlights',
    description: 'Review professional-quality highlight reels and full match recordings. See players in action before making contact.'
  },
  {
    iconPath: 'M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z',
    title: 'Connect Directly',
    description: 'Submit contact requests through our platform. We facilitate secure introductions between scouts and player representatives.'
  }
]

function calculateAge(dateOfBirth?: string): number | string {
  if (!dateOfBirth) return '—'
  const today = new Date()
  const birthDate = new Date(dateOfBirth)
  let age = today.getFullYear() - birthDate.getFullYear()
  const m = today.getMonth() - birthDate.getMonth()
  if (m < 0 || (m === 0 && today.getDate() < birthDate.getDate())) {
    age--
  }
  return age
}

function formatDuration(seconds?: number): string {
  if (!seconds) return '0:00'
  const mins = Math.floor(seconds / 60)
  const secs = seconds % 60
  return `${mins}:${secs.toString().padStart(2, '0')}`
}

function getHighlightIcon(type: string): string {
  const icons: Record<string, string> = {
    goal: '⚽',
    assist: '🎯',
    skill: '✨',
    save: '🧤',
    dribbling: '🦵',
    speed: '⚡',
    defending: '🛡️',
    passing: '➡️',
    shooting: '🎯',
    heading: '🗣️',
  }
  return icons[type.toLowerCase()] || '🎬'
}

function openHighlight(highlight: FeaturedHighlight) {
  selectedHighlight.value = highlight
  showHighlightModal.value = true
}

// Fetch featured players and highlights
onMounted(async () => {
  try {
    // Fetch players and highlights in parallel
    const [playersRes, highlightsRes] = await Promise.all([
      $fetch<ApiResponse<{ players: Player[] }>>('/players/featured', {
        baseURL: config.public.apiBase,
      }),
      $fetch<ApiResponse<FeaturedHighlight[]>>('/highlights/featured', {
        baseURL: config.public.apiBase,
      }),
    ])

    if (playersRes.success && playersRes.data?.players) {
      featuredPlayers.value = playersRes.data.players.slice(0, 5)
    }

    if (highlightsRes.success && highlightsRes.data) {
      featuredHighlights.value = highlightsRes.data
    }
  } catch (error) {
    console.error('Failed to fetch homepage data:', error)
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
@keyframes gradient {
  0%, 100% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
}

.animate-gradient {
  animation: gradient 6s ease infinite;
}

@keyframes float {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-10px); }
}

.animate-float {
  animation: float 3s ease-in-out infinite;
}

/* Video mockup animations */
@keyframes player-run {
  0%, 100% { transform: translateX(-5px) scale(1); }
  25% { transform: translateX(0) scale(1.02); }
  50% { transform: translateX(5px) scale(1); }
  75% { transform: translateX(0) scale(0.98); }
}

.animate-player-run {
  animation: player-run 2s ease-in-out infinite;
}

@keyframes ball-bounce {
  0%, 100% { transform: translate(-50%, -50%) translateY(0); }
  25% { transform: translate(-50%, -50%) translateY(-15px) translateX(10px); }
  50% { transform: translate(-50%, -50%) translateY(0) translateX(20px); }
  75% { transform: translate(-50%, -50%) translateY(-10px) translateX(10px); }
}

.animate-ball-bounce {
  animation: ball-bounce 2s ease-in-out infinite;
}

@keyframes trail-1 {
  0%, 100% { opacity: 0; transform: scale(0); }
  50% { opacity: 0.6; transform: scale(1); }
}

@keyframes trail-2 {
  0%, 100% { opacity: 0; transform: scale(0); }
  60% { opacity: 0.5; transform: scale(1); }
}

@keyframes trail-3 {
  0%, 100% { opacity: 0; transform: scale(0); }
  70% { opacity: 0.4; transform: scale(1); }
}

.animate-trail-1 { animation: trail-1 2s ease-in-out infinite; }
.animate-trail-2 { animation: trail-2 2s ease-in-out infinite 0.2s; }
.animate-trail-3 { animation: trail-3 2s ease-in-out infinite 0.4s; }

/* Testimonial transition */
.testimonial-fade-enter-active,
.testimonial-fade-leave-active {
  transition: all 0.4s ease;
}

.testimonial-fade-enter-from {
  opacity: 0;
  transform: translateX(20px);
}

.testimonial-fade-leave-to {
  opacity: 0;
  transform: translateX(-20px);
}
</style>
