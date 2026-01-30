<template>
  <div class="max-w-7xl mx-auto">
    <!-- Page Header - Compact on mobile -->
    <div class="flex items-center justify-between gap-4 mb-6">
      <div>
        <h1 class="font-display text-xl sm:text-2xl lg:text-3xl font-bold text-neutral-900">
          Tournaments
        </h1>
        <p class="hidden sm:block mt-1 text-neutral-600">Manage tournaments, matches, and player highlights</p>
      </div>
      <button
        @click="openCreateModal"
        class="inline-flex items-center gap-2 px-4 py-2.5 bg-primary-600 text-white rounded-xl text-sm font-semibold hover:bg-primary-700 transition-all shadow-md"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
        </svg>
        <span class="hidden sm:inline">Create Tournament</span>
        <span class="sm:hidden">New</span>
      </button>
    </div>

    <!-- Stats Cards - Horizontal scroll on mobile -->
    <div class="flex gap-3 mb-5 overflow-x-auto pb-2 -mx-4 px-4 sm:mx-0 sm:px-0 sm:grid sm:grid-cols-4 sm:overflow-visible scrollbar-hide">
      <div class="flex-shrink-0 w-[130px] sm:w-auto bg-gradient-to-br from-primary-500 to-emerald-600 rounded-xl p-4 text-white">
        <p class="text-primary-100 text-xs font-medium">Tournaments</p>
        <p class="text-2xl font-bold">{{ tournaments.length }}</p>
      </div>
      <div class="flex-shrink-0 w-[130px] sm:w-auto bg-gradient-to-br from-blue-500 to-indigo-600 rounded-xl p-4 text-white">
        <p class="text-blue-100 text-xs font-medium">Matches</p>
        <p class="text-2xl font-bold">{{ stats.total_matches }}</p>
      </div>
      <div class="flex-shrink-0 w-[130px] sm:w-auto bg-gradient-to-br from-amber-500 to-orange-500 rounded-xl p-4 text-white">
        <p class="text-amber-100 text-xs font-medium">Videos 💰</p>
        <p class="text-2xl font-bold">{{ stats.total_match_videos }}</p>
      </div>
      <div class="flex-shrink-0 w-[130px] sm:w-auto bg-gradient-to-br from-emerald-500 to-green-600 rounded-xl p-4 text-white">
        <p class="text-emerald-100 text-xs font-medium">Highlights 🆓</p>
        <p class="text-2xl font-bold">{{ stats.total_highlights }}</p>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="bg-white rounded-2xl border border-neutral-200 p-16">
      <div class="flex flex-col items-center justify-center">
        <div class="w-12 h-12 border-4 border-primary-500/30 border-t-primary-500 rounded-full animate-spin mb-4"></div>
        <p class="text-neutral-500">Loading tournaments...</p>
      </div>
    </div>

    <!-- Empty State -->
    <div v-else-if="tournaments.length === 0" class="bg-white rounded-2xl border border-neutral-200 p-12 text-center">
      <div class="w-20 h-20 bg-gradient-to-br from-primary-100 to-emerald-100 rounded-2xl flex items-center justify-center mx-auto mb-4">
        <svg class="w-10 h-10 text-primary-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
        </svg>
      </div>
      <h3 class="text-xl font-semibold text-neutral-900 mb-2">Start Building Your Video Library</h3>
      <p class="text-neutral-500 mb-6 max-w-md mx-auto">
        Create a tournament to begin uploading match videos and player highlights. 
        Scouts will discover players through free highlights and pay for full match access.
      </p>
      <button
        @click="openCreateModal"
        class="inline-flex items-center gap-2 px-6 py-3 bg-gradient-to-r from-primary-600 to-emerald-600 text-white rounded-xl font-semibold hover:from-primary-700 hover:to-emerald-700 transition-all shadow-lg shadow-primary-600/25"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
        </svg>
        Create Your First Tournament
      </button>
      
      <!-- How it works -->
      <div class="mt-10 pt-8 border-t border-neutral-100">
        <p class="text-sm font-medium text-neutral-700 mb-4">How it works:</p>
        <div class="flex flex-col sm:flex-row justify-center gap-4 text-sm text-neutral-600">
          <div class="flex items-center gap-2">
            <span class="w-6 h-6 bg-primary-100 text-primary-600 rounded-full flex items-center justify-center text-xs font-bold">1</span>
            Create Tournament
          </div>
          <svg class="w-4 h-4 text-neutral-300 hidden sm:block self-center" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
          </svg>
          <div class="flex items-center gap-2">
            <span class="w-6 h-6 bg-primary-100 text-primary-600 rounded-full flex items-center justify-center text-xs font-bold">2</span>
            Add Matches
          </div>
          <svg class="w-4 h-4 text-neutral-300 hidden sm:block self-center" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
          </svg>
          <div class="flex items-center gap-2">
            <span class="w-6 h-6 bg-primary-100 text-primary-600 rounded-full flex items-center justify-center text-xs font-bold">3</span>
            Upload Videos & Highlights
          </div>
        </div>
      </div>
    </div>

    <!-- Tournaments Grid -->
    <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-3 sm:gap-6">
      <NuxtLink
        v-for="tournament in tournaments"
        :key="tournament.id"
        :to="`/admin/tournaments/${tournament.id}`"
        class="group bg-white rounded-xl sm:rounded-2xl border border-neutral-200 overflow-hidden hover:shadow-xl hover:border-primary-200 hover:-translate-y-0.5 transition-all duration-200 shadow-sm"
      >
        <!-- Mobile: Compact list style -->
        <div class="flex sm:hidden p-3 gap-3">
          <!-- Small thumbnail -->
          <div class="w-16 h-16 bg-gradient-to-br from-primary-500 to-emerald-600 rounded-lg relative flex-shrink-0 overflow-hidden">
            <img
              v-if="tournament.cover_image_url || tournament.thumbnail_url"
              :src="tournament.cover_image_url || tournament.thumbnail_url"
              :alt="tournament.name"
              class="w-full h-full object-cover"
            />
            <div v-else class="w-full h-full flex items-center justify-center">
              <svg class="w-6 h-6 text-white/70" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z" />
              </svg>
            </div>
          </div>
          
          <!-- Content -->
          <div class="flex-1 min-w-0">
            <div class="flex items-start justify-between gap-2">
              <h3 class="font-semibold text-sm text-neutral-900 line-clamp-1">
                {{ tournament.name }}
              </h3>
              <span
                :class="getMobileStatusClass(tournament.status)"
                class="px-2 py-0.5 rounded-full text-[10px] font-medium capitalize flex-shrink-0"
              >
                {{ tournament.status }}
              </span>
            </div>
            <p class="text-xs text-neutral-500 mt-0.5 truncate">
              {{ tournament.location || 'No location' }} · {{ formatDateRange(tournament.start_date, tournament.end_date) }}
            </p>
            <!-- Labeled stats -->
            <div class="flex items-center gap-3 mt-2 text-[11px] text-neutral-500">
              <span><span class="font-medium text-neutral-700">{{ tournament.match_count || 0 }}</span> matches</span>
              <span><span class="font-medium text-neutral-700">{{ tournament.video_count || 0 }}</span> videos</span>
              <span><span class="font-medium text-neutral-700">{{ tournament.highlight_count || 0 }}</span> clips</span>
            </div>
          </div>
          
          <!-- Arrow -->
          <svg class="w-4 h-4 text-neutral-400 flex-shrink-0 self-center" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
          </svg>
        </div>
        
        <!-- Desktop: Card layout -->
        <div class="hidden sm:block">
          <!-- Thumbnail -->
          <div class="w-full aspect-video bg-gradient-to-br from-primary-500 to-emerald-600 relative overflow-hidden">
            <img
              v-if="tournament.cover_image_url || tournament.thumbnail_url"
              :src="tournament.cover_image_url || tournament.thumbnail_url"
              :alt="tournament.name"
              class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
            />
            <div v-else class="w-full h-full flex items-center justify-center">
              <div class="w-20 h-20 rounded-full bg-white/10 flex items-center justify-center backdrop-blur-sm">
                <svg class="w-10 h-10 text-white/70" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z" />
                </svg>
              </div>
            </div>
            <!-- Status Badge -->
            <span
              :class="getMobileStatusClass(tournament.status)"
              class="absolute top-3 right-3 px-2.5 py-1 rounded-full text-xs font-semibold capitalize shadow-lg"
            >
              {{ tournament.status }}
            </span>
          </div>

          <!-- Content -->
          <div class="p-5">
            <h3 class="font-bold text-base text-neutral-900 group-hover:text-primary-600 transition-colors mb-2">
              {{ tournament.name }}
            </h3>
            <div class="flex flex-col gap-1.5 text-sm text-neutral-500">
              <span v-if="tournament.location" class="flex items-center gap-1.5">
                <svg class="w-4 h-4 flex-shrink-0 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                </svg>
                <span class="truncate">{{ tournament.location }}</span>
              </span>
              <span class="flex items-center gap-1.5">
                <svg class="w-4 h-4 flex-shrink-0 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
                <span>{{ formatDateRange(tournament.start_date, tournament.end_date) }}</span>
              </span>
            </div>
            
            <!-- Video Stats -->
            <div class="flex items-center gap-4 mt-3 pt-3 border-t border-neutral-100">
              <span class="flex items-center gap-1.5 text-sm">
                <span class="text-neutral-700 font-medium">{{ tournament.match_count || 0 }}</span>
                <span class="text-neutral-500">matches</span>
              </span>
              <span class="flex items-center gap-1.5 text-sm">
                <span class="text-neutral-700 font-medium">{{ tournament.video_count || 0 }}</span>
                <span class="text-neutral-500">videos</span>
              </span>
              <span class="flex items-center gap-1.5 text-sm">
                <span class="text-neutral-700 font-medium">{{ tournament.highlight_count || 0 }}</span>
                <span class="text-neutral-500">clips</span>
              </span>
            </div>
          </div>
        </div>
      </NuxtLink>
    </div>

    <!-- Create Tournament Modal -->
    <Teleport to="body">
      <Transition
        enter-active-class="transition-opacity duration-200"
        enter-from-class="opacity-0"
        enter-to-class="opacity-100"
        leave-active-class="transition-opacity duration-150"
        leave-from-class="opacity-100"
        leave-to-class="opacity-0"
      >
        <div v-if="showCreateModal" class="fixed inset-0 z-50 overflow-y-auto">
          <div class="fixed inset-0 bg-black/50 backdrop-blur-sm" @click="closeCreateModal"></div>
          <div class="relative min-h-screen flex items-center justify-center p-4">
            <Transition
              enter-active-class="transition-all duration-200"
              enter-from-class="opacity-0 scale-95"
              enter-to-class="opacity-100 scale-100"
              leave-active-class="transition-all duration-150"
              leave-from-class="opacity-100 scale-100"
              leave-to-class="opacity-0 scale-95"
            >
              <div v-if="showCreateModal" class="relative bg-white rounded-2xl shadow-xl w-full max-w-lg overflow-hidden">
                <!-- Modal Header -->
                <div class="px-6 py-4 border-b border-neutral-200 bg-gradient-to-r from-primary-50 to-emerald-50">
                  <div class="flex items-center justify-between">
                    <div class="flex items-center gap-3">
                      <div class="w-10 h-10 bg-gradient-to-br from-primary-500 to-emerald-600 rounded-xl flex items-center justify-center text-white shadow-lg shadow-primary-500/25">
                        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z" />
                        </svg>
                      </div>
                      <div>
                        <h3 class="font-semibold text-neutral-900">Create Tournament</h3>
                        <p class="text-sm text-neutral-500">Add a new tournament to manage videos</p>
                      </div>
                    </div>
                    <button
                      @click="closeCreateModal"
                      class="w-8 h-8 flex items-center justify-center rounded-lg hover:bg-neutral-100 text-neutral-400 hover:text-neutral-600 transition-colors"
                    >
                      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                      </svg>
                    </button>
                  </div>
                </div>

                <!-- Modal Body -->
                <form @submit.prevent="createTournament" class="p-6 space-y-5">
                  <!-- Tournament Name -->
                  <div>
                    <label class="block text-sm font-medium text-neutral-700 mb-2">Tournament Name *</label>
                    <input
                      v-model="form.name"
                      type="text"
                      placeholder="e.g., Summer Cup 2026"
                      class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
                      required
                    />
                  </div>

                  <!-- Date Range -->
                  <div class="grid grid-cols-2 gap-4">
                    <div>
                      <label class="block text-sm font-medium text-neutral-700 mb-2">Start Date *</label>
                      <input
                        v-model="form.start_date"
                        type="date"
                        class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
                        required
                      />
                    </div>
                    <div>
                      <label class="block text-sm font-medium text-neutral-700 mb-2">End Date *</label>
                      <input
                        v-model="form.end_date"
                        type="date"
                        class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
                        required
                      />
                    </div>
                  </div>

                  <!-- Location Section -->
                  <div class="space-y-4">
                    <div class="flex items-center gap-2 text-sm font-medium text-neutral-700">
                      <svg class="w-4 h-4 text-neutral-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
                      </svg>
                      Location
                    </div>

                    <!-- Country Dropdown -->
                    <div>
                      <label class="block text-sm font-medium text-neutral-700 mb-2">Country</label>
                      <div class="relative">
                        <select
                          v-model="form.country"
                          @change="onCountryChange"
                          class="appearance-none w-full px-4 py-3 pr-10 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
                        >
                          <option value="">Select country</option>
                          <option v-for="country in AFRICAN_COUNTRIES" :key="country" :value="country">
                            {{ country }}
                          </option>
                        </select>
                        <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                          <svg class="w-5 h-5 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                          </svg>
                        </div>
                      </div>
                    </div>

                    <!-- State/City Row -->
                    <div class="grid grid-cols-2 gap-4">
                      <!-- State Dropdown -->
                      <div>
                        <label class="block text-sm font-medium text-neutral-700 mb-2">State/Region</label>
                        <div v-if="countryHasLocationData && availableStates.length > 0" class="relative">
                          <select
                            v-model="form.state"
                            @change="onStateChange"
                            :disabled="!form.country"
                            class="appearance-none w-full px-4 py-3 pr-10 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all disabled:bg-neutral-100 disabled:cursor-not-allowed"
                          >
                            <option value="">Select state</option>
                            <option v-for="state in availableStates" :key="state" :value="state">
                              {{ state }}
                            </option>
                          </select>
                          <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                            <svg class="w-5 h-5 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                            </svg>
                          </div>
                        </div>
                        <input
                          v-else
                          v-model="form.state"
                          type="text"
                          :disabled="!form.country"
                          placeholder="Enter state"
                          class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all disabled:bg-neutral-100 disabled:cursor-not-allowed"
                        />
                      </div>

                      <!-- City Dropdown -->
                      <div>
                        <label class="block text-sm font-medium text-neutral-700 mb-2">City</label>
                        <div v-if="countryHasLocationData && availableCities.length > 0" class="relative">
                          <select
                            v-model="form.city"
                            :disabled="!form.state"
                            class="appearance-none w-full px-4 py-3 pr-10 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all disabled:bg-neutral-100 disabled:cursor-not-allowed"
                          >
                            <option value="">Select city</option>
                            <option v-for="city in availableCities" :key="city" :value="city">
                              {{ city }}
                            </option>
                          </select>
                          <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                            <svg class="w-5 h-5 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                            </svg>
                          </div>
                        </div>
                        <input
                          v-else
                          v-model="form.city"
                          type="text"
                          :disabled="!form.country"
                          placeholder="Enter city"
                          class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all disabled:bg-neutral-100 disabled:cursor-not-allowed"
                        />
                      </div>
                    </div>
                  </div>

                  <!-- Description -->
                  <div>
                    <label class="block text-sm font-medium text-neutral-700 mb-2">Description</label>
                    <textarea
                      v-model="form.description"
                      rows="3"
                      placeholder="Brief description of the tournament..."
                      class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all resize-none"
                    ></textarea>
                  </div>

                  <!-- Modal Footer -->
                  <div class="flex gap-3 pt-4 border-t border-neutral-100">
                    <button
                      type="button"
                      @click="closeCreateModal"
                      class="w-auto px-6 py-3 border border-neutral-300 text-neutral-700 rounded-xl text-sm font-medium hover:bg-neutral-50 transition-colors"
                    >
                      Cancel
                    </button>
                    <button
                      type="submit"
                      :disabled="creating || !isFormValid"
                      :class="[
                        'flex-1 inline-flex items-center justify-center gap-2 px-4 py-3 rounded-xl text-sm font-semibold transition-all',
                        isFormValid
                          ? 'bg-gradient-to-r from-primary-600 to-emerald-600 text-white hover:from-primary-700 hover:to-emerald-700 shadow-lg shadow-primary-600/25'
                          : 'bg-neutral-100 text-neutral-400 cursor-not-allowed'
                      ]"
                    >
                      <svg v-if="creating" class="w-4 h-4 animate-spin flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" fill="none" />
                        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z" />
                      </svg>
                      <span class="whitespace-nowrap">{{ creating ? 'Creating...' : 'Create' }}</span>
                    </button>
                  </div>
                </form>
              </div>
            </Transition>
          </div>
        </div>
      </Transition>
    </Teleport>
  </div>
</template>

<script setup lang="ts">
import type { ApiResponse } from '~/types/index'
import { AFRICAN_COUNTRIES, getStatesForCountry, getCitiesForState, hasLocationData } from '~/data/locations'

interface Tournament {
  id: string
  name: string
  description?: string
  start_date: string
  end_date: string
  location?: string
  status: string
  thumbnail_url?: string
  cover_image_url?: string
  match_count?: number
  video_count?: number
  highlight_count?: number
}

interface VideoStats {
  total_matches: number
  total_match_videos: number
  total_highlights: number
}

definePageMeta({
  layout: 'admin',
  middleware: 'admin',
})

const config = useRuntimeConfig()
const authStore = useAuthStore()
const toast = useToast()
const router = useRouter()

const loading = ref(true)
const tournaments = ref<Tournament[]>([])
const stats = ref<VideoStats>({
  total_matches: 0,
  total_match_videos: 0,
  total_highlights: 0,
})
const showCreateModal = ref(false)
const creating = ref(false)

const form = reactive({
  name: '',
  start_date: '',
  end_date: '',
  country: '',
  state: '',
  city: '',
  description: '',
})

// Location helpers
const countryHasLocationData = computed(() => hasLocationData(form.country))
const availableStates = computed(() => getStatesForCountry(form.country))
const availableCities = computed(() => getCitiesForState(form.country, form.state))

// Form validation - all required fields must be filled
const isFormValid = computed(() => {
  return form.name.trim() && form.start_date && form.end_date
})

function onCountryChange() {
  form.state = ''
  form.city = ''
}

function onStateChange() {
  form.city = ''
}

function openCreateModal() {
  // Reset form
  form.name = ''
  form.start_date = ''
  form.end_date = ''
  form.country = ''
  form.state = ''
  form.city = ''
  form.description = ''
  showCreateModal.value = true
}

function closeCreateModal() {
  showCreateModal.value = false
}

// Build location string from parts
function buildLocationString(): string {
  const parts = [form.city, form.state, form.country].filter(Boolean)
  return parts.join(', ')
}

async function fetchTournaments() {
  try {
    const response = await $fetch<ApiResponse<{ events: Tournament[] }>>('/admin/events', {
      baseURL: config.public.apiBase,
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
    })
    if (response.success && response.data) {
      // Backend returns "events" not "tournaments"
      tournaments.value = response.data.events || []
      
      // Calculate aggregate stats from tournaments
      stats.value = {
        total_matches: tournaments.value.reduce((sum, t) => sum + (t.match_count || 0), 0),
        total_match_videos: tournaments.value.reduce((sum, t) => sum + (t.video_count || 0), 0),
        total_highlights: tournaments.value.reduce((sum, t) => sum + (t.highlight_count || 0), 0),
      }
    }
  } catch (error) {
    console.error('Failed to fetch tournaments:', error)
    toast.error('Error', 'Failed to load tournaments')
  } finally {
    loading.value = false
  }
}

async function createTournament() {
  creating.value = true
  try {
    // Extract year from start_date
    const year = form.start_date ? new Date(form.start_date).getFullYear() : new Date().getFullYear()
    
    const payload = {
      name: form.name,
      year: year,
      start_date: form.start_date,
      end_date: form.end_date,
      location: buildLocationString() || undefined,
      country: form.country || undefined,
      description: form.description || undefined,
    }
    
    const response = await $fetch<ApiResponse<Tournament>>('/admin/events', {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: payload,
    })
    if (response.success && response.data) {
      toast.success('Tournament Created', 'Tournament has been created successfully')
      closeCreateModal()
      router.push(`/admin/tournaments/${response.data.id}`)
    }
  } catch (error: any) {
    toast.error('Error', error.data?.message || 'Failed to create tournament')
  } finally {
    creating.value = false
  }
}

function formatDateRange(start: string, end: string): string {
  const startDate = new Date(start)
  const endDate = new Date(end)
  const opts: Intl.DateTimeFormatOptions = { month: 'short', day: 'numeric' }
  
  if (startDate.getFullYear() === endDate.getFullYear() && startDate.getMonth() === endDate.getMonth()) {
    return `${startDate.toLocaleDateString('en-US', opts)} - ${endDate.getDate()}, ${endDate.getFullYear()}`
  }
  return `${startDate.toLocaleDateString('en-US', opts)} - ${endDate.toLocaleDateString('en-US', opts)}, ${endDate.getFullYear()}`
}

function getStatusClass(status: string): string {
  const classes: Record<string, string> = {
    draft: 'bg-neutral-100 text-neutral-700',
    upcoming: 'bg-blue-100 text-blue-700',
    active: 'bg-emerald-100 text-emerald-700',
    completed: 'bg-amber-100 text-amber-700',
  }
  return classes[status] || classes.draft!
}

// High contrast status badge for mobile (white background with colored text)
function getMobileStatusClass(status: string): string {
  const classes: Record<string, string> = {
    draft: 'bg-white text-neutral-700 border border-neutral-200',
    upcoming: 'bg-white text-blue-600 border border-blue-200',
    active: 'bg-white text-emerald-600 border border-emerald-200',
    completed: 'bg-white text-amber-600 border border-amber-200',
  }
  return classes[status] || classes.draft!
}

onMounted(() => {
  fetchTournaments()
})
</script>
