<template>
  <div class="max-w-7xl mx-auto">
    <!-- Page Header -->
    <div class="mb-8">
      <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
        <div class="flex items-center gap-4">
          <NuxtLink 
            to="/admin/tournaments" 
            class="w-10 h-10 flex items-center justify-center rounded-xl bg-neutral-100 hover:bg-neutral-200 text-neutral-600 transition-colors"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
            </svg>
          </NuxtLink>
          <div>
            <h1 class="font-display text-2xl lg:text-3xl font-bold text-neutral-900">
              Tournament Details
            </h1>
            <p class="mt-1 text-neutral-600">Manage matches and upload videos</p>
          </div>
        </div>
        <div class="flex items-center gap-3">
          <button
            @click="showEditModal = true"
            class="inline-flex items-center gap-2 px-5 py-2.5 bg-neutral-100 text-neutral-700 rounded-xl text-sm font-semibold hover:bg-neutral-200 transition-colors"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
            </svg>
            Edit Tournament
          </button>
        </div>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex flex-col items-center justify-center py-20">
      <div class="w-12 h-12 border-4 border-primary-200 border-t-primary-600 rounded-full animate-spin"></div>
      <p class="mt-4 text-neutral-600">Loading tournament...</p>
    </div>

    <!-- Content -->
    <div v-else-if="tournament" class="space-y-6">
      <!-- Tournament Header -->
      <!-- Tournament Header Card - Simplified on mobile -->
      <div class="bg-white sm:rounded-2xl sm:border sm:border-neutral-200 sm:shadow-sm overflow-hidden -mx-4 sm:mx-0">
        <!-- Desktop header only -->
        <div class="hidden sm:block px-6 py-4 border-b border-neutral-200 bg-gradient-to-r from-primary-50 to-emerald-50">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 bg-gradient-to-br from-primary-500 to-emerald-600 rounded-xl flex items-center justify-center text-white shadow-lg shadow-primary-500/25">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z" />
              </svg>
            </div>
            <div>
              <h2 class="font-semibold text-neutral-900">Tournament Overview</h2>
              <p class="text-sm text-neutral-500">Details and stats</p>
            </div>
          </div>
        </div>
        
        <div class="p-4 sm:p-6">
          <div class="flex flex-col sm:flex-row sm:items-start gap-4 sm:gap-6">
            <!-- Thumbnail - horizontal on mobile, larger on desktop -->
            <div class="flex sm:block gap-4">
              <div class="w-20 h-20 sm:w-64 sm:h-40 bg-gradient-to-br from-primary-500 to-emerald-600 rounded-xl flex items-center justify-center overflow-hidden flex-shrink-0">
                <img
                  v-if="tournament.thumbnail_url"
                  :src="tournament.thumbnail_url"
                  :alt="tournament.name"
                  class="w-full h-full object-cover"
                />
                <div v-else class="w-10 h-10 sm:w-16 sm:h-16 rounded-full bg-white/10 flex items-center justify-center backdrop-blur-sm">
                  <svg class="w-5 h-5 sm:w-8 sm:h-8 text-white/70" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z" />
                  </svg>
                </div>
              </div>
              
              <!-- Mobile-only: Title and status next to image -->
              <div class="sm:hidden flex-1 min-w-0">
                <h3 class="text-base font-bold text-neutral-900 line-clamp-2 mb-1">{{ tournament.name }}</h3>
                <span :class="getStatusClass(tournament.status)" class="inline-block px-2 py-0.5 rounded-full text-[11px] font-medium capitalize">
                  {{ tournament.status }}
                </span>
                <p v-if="tournament.description" class="text-xs text-neutral-500 line-clamp-2 mt-1.5">{{ tournament.description }}</p>
              </div>
            </div>

            <!-- Info - Desktop -->
            <div class="flex-1 hidden sm:block">
              <div class="flex items-center gap-3 mb-2">
                <h3 class="text-xl font-bold text-neutral-900">{{ tournament.name }}</h3>
                <span :class="getStatusClass(tournament.status)" class="px-3 py-1 rounded-full text-sm font-medium capitalize">
                  {{ tournament.status }}
                </span>
              </div>
              <p v-if="tournament.description" class="text-neutral-600 mb-4">{{ tournament.description }}</p>
              <div class="flex flex-wrap items-center gap-6 text-sm text-neutral-600">
                <span class="flex items-center gap-2">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                  </svg>
                  {{ formatDateRange(tournament.start_date, tournament.end_date) }}
                </span>
                <span v-if="tournament.location" class="flex items-center gap-2">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                  </svg>
                  {{ tournament.location }}
                </span>
                <span class="flex items-center gap-2 font-medium text-primary-600">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
                  </svg>
                  {{ matches.length }} matches
                </span>
              </div>
            </div>
          </div>
          
          <!-- Mobile Meta Info Bar -->
          <div class="sm:hidden flex flex-wrap items-center gap-x-3 gap-y-1.5 mt-3 text-[11px] text-neutral-500">
            <span class="flex items-center gap-1">
              <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
              </svg>
              {{ formatDateRange(tournament.start_date, tournament.end_date) }}
            </span>
            <span v-if="tournament.location" class="flex items-center gap-1">
              <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
              </svg>
              {{ tournament.location }}
            </span>
            <span class="flex items-center gap-1 font-medium text-primary-600">
              <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
              </svg>
              {{ matches.length }} matches
            </span>
          </div>
        </div>
      </div>

      <!-- Matches Section - Simplified on mobile -->
      <div class="sm:bg-white sm:rounded-2xl sm:border sm:border-neutral-200 sm:shadow-sm overflow-hidden -mx-4 sm:mx-0">
        <!-- Mobile: Simple header -->
        <div class="sm:hidden flex items-center justify-between px-4 py-3 border-b border-neutral-100">
          <h2 class="font-semibold text-neutral-900">Matches</h2>
          <button
            @click="showCreateMatchModal = true"
            class="inline-flex items-center gap-1.5 px-3 py-1.5 bg-primary-600 text-white rounded-lg text-xs font-semibold"
          >
            <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
            </svg>
            Add
          </button>
        </div>
        
        <!-- Desktop: Full header -->
        <div class="hidden sm:block px-6 py-4 border-b border-neutral-200 bg-gradient-to-r from-blue-50 to-indigo-50">
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-3">
              <div class="w-10 h-10 bg-gradient-to-br from-blue-500 to-indigo-600 rounded-xl flex items-center justify-center text-white shadow-lg shadow-blue-500/25">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
                </svg>
              </div>
              <div>
                <h2 class="font-semibold text-neutral-900">Matches</h2>
                <p class="text-sm text-neutral-500">Manage matches and upload videos</p>
              </div>
            </div>
            <button
              @click="showCreateMatchModal = true"
              class="inline-flex items-center gap-2 px-4 py-2 bg-gradient-to-r from-primary-600 to-emerald-600 text-white rounded-xl text-sm font-semibold hover:from-primary-700 hover:to-emerald-700 transition-all shadow-lg shadow-primary-600/25"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
              </svg>
              Add Match
            </button>
          </div>
        </div>

        <div class="p-4 sm:p-6">

        <!-- Empty State -->
        <div v-if="matches.length === 0" class="text-center py-8 sm:py-12">
          <div class="w-12 h-12 sm:w-16 sm:h-16 bg-neutral-100 rounded-xl sm:rounded-2xl flex items-center justify-center mx-auto mb-3 sm:mb-4">
            <svg class="w-6 h-6 sm:w-8 sm:h-8 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
            </svg>
          </div>
          <p class="text-sm text-neutral-500 mb-3">No matches added yet</p>
          <button
            @click="showCreateMatchModal = true"
            class="text-primary-600 hover:text-primary-700 font-medium text-sm"
          >
            Add the first match
          </button>
        </div>

        <!-- Matches List -->
        <div v-else class="space-y-3">
          <NuxtLink
            v-for="match in matches"
            :key="match.id"
            :to="`/admin/tournaments/${tournamentId}/matches/${match.id}`"
            class="block p-3 sm:p-4 bg-neutral-50 rounded-xl hover:bg-neutral-100 transition-colors group"
          >
            <!-- Mobile Layout -->
            <div class="sm:hidden">
              <div class="flex items-start gap-3">
                <!-- Match Number/Stage -->
                <div class="w-10 h-10 bg-gradient-to-br from-primary-500 to-emerald-600 rounded-lg flex items-center justify-center text-white font-bold text-sm flex-shrink-0">
                  {{ match.match_number || '#' }}
                </div>

                <div class="flex-1 min-w-0">
                  <p class="font-semibold text-neutral-900 group-hover:text-primary-600 transition-colors text-sm">
                    {{ match.title }}
                  </p>
                  <div class="flex flex-wrap items-center gap-x-2 gap-y-1 text-xs text-neutral-500 mt-1">
                    <span v-if="match.home_team && match.away_team" class="font-medium">
                      {{ match.home_team }}
                      <span v-if="match.home_score !== null" class="text-neutral-900">{{ match.home_score }}</span>
                      <span class="text-neutral-400">vs</span>
                      <span v-if="match.away_score !== null" class="text-neutral-900">{{ match.away_score }}</span>
                      {{ match.away_team }}
                    </span>
                  </div>
                </div>
              </div>
              
              <!-- Mobile Meta Row -->
              <div class="flex items-center justify-between mt-3 pt-3 border-t border-neutral-200/50">
                <div class="flex items-center gap-2">
                  <span class="flex items-center gap-1 text-[11px] text-neutral-500">
                    <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                    </svg>
                    {{ formatDate(match.match_date) }}
                  </span>
                  <span v-if="match.stage" class="px-1.5 py-0.5 bg-neutral-200 text-neutral-600 rounded text-[10px] font-medium">
                    {{ formatStage(match.stage) }}
                  </span>
                </div>
                <span
                  :class="match.has_video ? 'text-emerald-600' : 'text-neutral-400'"
                  class="flex items-center gap-1 text-xs font-medium"
                >
                  <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
                  </svg>
                  {{ match.has_video ? 'Video' : 'No video' }}
                </span>
              </div>
            </div>

            <!-- Desktop Layout -->
            <div class="hidden sm:flex items-center justify-between">
              <div class="flex items-center gap-4">
                <!-- Match Number/Stage -->
                <div class="w-12 h-12 bg-gradient-to-br from-primary-500 to-emerald-600 rounded-xl flex items-center justify-center text-white font-bold">
                  {{ match.match_number || '#' }}
                </div>

                <div>
                  <p class="font-medium text-neutral-900 group-hover:text-primary-600 transition-colors">
                    {{ match.title }}
                  </p>
                  <div class="flex items-center gap-3 text-sm text-neutral-500 mt-0.5">
                    <span v-if="match.home_team && match.away_team">
                      {{ match.home_team }}
                      <span v-if="match.home_score !== null" class="font-medium text-neutral-700">{{ match.home_score }}</span>
                      vs
                      <span v-if="match.away_score !== null" class="font-medium text-neutral-700">{{ match.away_score }}</span>
                      {{ match.away_team }}
                    </span>
                    <span class="flex items-center gap-1">
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                      </svg>
                      {{ formatDate(match.match_date) }}
                    </span>
                    <span v-if="match.stage" class="px-2 py-0.5 bg-neutral-200 text-neutral-600 rounded text-xs">
                      {{ formatStage(match.stage) }}
                    </span>
                  </div>
                </div>
              </div>

              <div class="flex items-center gap-6">
                <!-- Video Status -->
                <div class="flex items-center gap-3 text-sm">
                  <span
                    :class="match.has_video ? 'text-emerald-600' : 'text-neutral-400'"
                    class="flex items-center gap-1"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
                    </svg>
                    {{ match.has_video ? 'Video' : 'No video' }}
                  </span>
                  <span class="text-neutral-600">
                    {{ match.player_count || 0 }} players
                  </span>
                  <span class="text-primary-600 font-medium">
                    {{ match.highlight_count || 0 }} highlights
                  </span>
                </div>

                <!-- Status Badge -->
                <span :class="getMatchStatusClass(match.status)" class="px-2.5 py-1 rounded-full text-xs font-medium capitalize">
                  {{ match.status }}
                </span>

                <!-- Arrow -->
                <svg class="w-5 h-5 text-neutral-400 group-hover:text-primary-600 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
                </svg>
              </div>
            </div>
          </NuxtLink>
        </div>
        </div>
      </div>
    </div>

    <!-- Create Match Modal -->
    <Teleport to="body">
      <div v-if="showCreateMatchModal" class="fixed inset-0 z-50 overflow-y-auto">
        <div class="fixed inset-0 bg-black/50" @click="showCreateMatchModal = false"></div>
        <div class="relative min-h-screen flex items-center justify-center p-4">
          <div class="relative bg-white rounded-2xl shadow-xl w-full max-w-lg p-6">
            <h3 class="text-xl font-semibold mb-6">Add Match</h3>

            <form @submit.prevent="createMatch" class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-1">Match Title *</label>
                <input
                  v-model="matchForm.title"
                  type="text"
                  class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                  placeholder="e.g., Quarter Final - Game 1"
                  required
                />
              </div>

              <div class="grid grid-cols-2 gap-4">
                <div>
                  <label class="block text-sm font-medium text-neutral-700 mb-1">Home Team</label>
                  <input
                    v-model="matchForm.home_team"
                    type="text"
                    class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                    placeholder="Team A"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-neutral-700 mb-1">Away Team</label>
                  <input
                    v-model="matchForm.away_team"
                    type="text"
                    class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                    placeholder="Team B"
                  />
                </div>
              </div>

              <div class="grid grid-cols-2 gap-4">
                <div>
                  <label class="block text-sm font-medium text-neutral-700 mb-1">Match Date *</label>
                  <input
                    v-model="matchForm.match_date"
                    type="datetime-local"
                    :min="matchDateMin ? matchDateMin + 'T00:00' : undefined"
                    :max="matchDateMax ? matchDateMax + 'T23:59' : undefined"
                    :class="['w-full px-4 py-3 border rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent', !isMatchDateValid ? 'border-rose-300 bg-rose-50' : 'border-neutral-200']"
                    required
                  />
                  <p v-if="matchDateError" class="mt-1 text-xs text-rose-600">{{ matchDateError }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-neutral-700 mb-1">Stage</label>
                  <select
                    v-model="matchForm.stage"
                    class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent bg-white"
                  >
                    <option value="">Select stage</option>
                    <option value="group">Group Stage</option>
                    <option value="round_of_16">Round of 16</option>
                    <option value="quarterfinal">Quarter Final</option>
                    <option value="semifinal">Semi Final</option>
                    <option value="final">Final</option>
                  </select>
                </div>
              </div>

              <!-- Location Section - Google-style Cascading Selection -->
              <div class="space-y-4">
                <!-- Country indicator (if detected) -->
                <div v-if="detectedCountry" class="flex items-center gap-2 px-3 py-2 bg-primary-50 border border-primary-100 rounded-xl">
                  <svg class="w-4 h-4 text-primary-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
                  </svg>
                  <span class="text-sm text-primary-700">Showing venues in <strong>{{ detectedCountry }}</strong></span>
                </div>

                <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                  <!-- City Dropdown -->
                  <div>
                    <label class="block text-sm font-medium text-neutral-700 mb-1">
                      City
                      <span class="text-neutral-400 font-normal">(optional)</span>
                    </label>
                    <div class="relative">
                      <select
                        v-model="matchForm.city"
                        class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent bg-white appearance-none cursor-pointer"
                      >
                        <option value="">Select city</option>
                        <option v-for="city in availableCities" :key="city" :value="city">{{ city }}</option>
                      </select>
                      <div class="absolute right-4 top-1/2 -translate-y-1/2 pointer-events-none">
                        <svg class="w-4 h-4 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                        </svg>
                      </div>
                    </div>
                    <p class="mt-1 text-xs text-neutral-500">Where the match will be played</p>
                  </div>

                  <!-- Venue - Dropdown if suggestions available, else text input -->
                  <div>
                    <label class="block text-sm font-medium text-neutral-700 mb-1">
                      Venue
                      <span class="text-neutral-400 font-normal">(optional)</span>
                    </label>
                    
                    <!-- Show dropdown if we have venue suggestions -->
                    <div v-if="availableVenues.length > 0" class="relative">
                      <select
                        v-model="matchForm.venue_name"
                        class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent bg-white appearance-none cursor-pointer"
                      >
                        <option value="">Select venue</option>
                        <option v-for="venue in availableVenues" :key="venue" :value="venue">{{ venue }}</option>
                        <option value="__custom__">✏️ Enter custom venue...</option>
                      </select>
                      <div class="absolute right-4 top-1/2 -translate-y-1/2 pointer-events-none">
                        <svg class="w-4 h-4 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                        </svg>
                      </div>
                    </div>
                    
                    <!-- Show text input if no suggestions or custom selected -->
                    <input
                      v-else
                      v-model="matchForm.venue_name"
                      type="text"
                      class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                      placeholder="e.g., City Stadium"
                    />
                    
                    <!-- Custom venue input (shows when __custom__ is selected) -->
                    <input
                      v-if="matchForm.venue_name === '__custom__'"
                      v-model="customVenueName"
                      type="text"
                      class="w-full mt-2 px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                      placeholder="Enter stadium/field name"
                    />
                    <p class="mt-1 text-xs text-neutral-500">
                      {{ matchForm.city ? `Stadiums in ${matchForm.city}` : 'Select city for venue suggestions' }}
                    </p>
                  </div>
                </div>
              </div>

              <!-- Modal Footer -->
              <div class="flex gap-3 pt-4 border-t border-neutral-100">
                <button
                  type="button"
                  @click="showCreateMatchModal = false"
                  class="w-auto px-6 py-3 border border-neutral-300 text-neutral-700 rounded-xl text-sm font-medium hover:bg-neutral-50 transition-colors"
                >
                  Cancel
                </button>
                <button
                  type="submit"
                  :disabled="creatingMatch || !isMatchFormValid"
                  :class="[
                    'flex-1 inline-flex items-center justify-center gap-2 px-4 py-3 rounded-xl text-sm font-semibold transition-all',
                    isMatchFormValid
                      ? 'bg-gradient-to-r from-primary-600 to-emerald-600 text-white hover:from-primary-700 hover:to-emerald-700 shadow-lg shadow-primary-600/25'
                      : 'bg-neutral-100 text-neutral-400 cursor-not-allowed'
                  ]"
                >
                  <svg v-if="creatingMatch" class="w-4 h-4 animate-spin flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" fill="none" />
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z" />
                  </svg>
                  <span class="whitespace-nowrap">{{ creatingMatch ? 'Creating...' : 'Create' }}</span>
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- Edit Tournament Modal -->
    <Teleport to="body">
      <div v-if="showEditModal" class="fixed inset-0 z-50 flex items-center justify-center p-4">
        <div class="fixed inset-0 bg-black/50" @click="showEditModal = false"></div>
        <div class="relative bg-white rounded-2xl shadow-xl w-full max-w-lg max-h-[90vh] overflow-y-auto">
          <div class="sticky top-0 bg-white px-6 py-4 border-b border-neutral-200 flex items-center justify-between">
            <h3 class="text-xl font-semibold text-neutral-900">Edit Tournament</h3>
            <button
              @click="showEditModal = false"
              class="w-8 h-8 flex items-center justify-center rounded-lg hover:bg-neutral-100 text-neutral-500"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
          
          <form @submit.prevent="updateTournament" class="p-6 space-y-5">
            <div>
              <label class="block text-sm font-medium text-neutral-700 mb-1">Tournament Name *</label>
              <input
                v-model="editForm.name"
                type="text"
                class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                placeholder="e.g., Africa Youth Cup 2026"
                required
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-neutral-700 mb-1">Description</label>
              <textarea
                v-model="editForm.description"
                rows="3"
                class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent resize-none"
                placeholder="Brief description of the tournament..."
              ></textarea>
            </div>

            <div class="grid grid-cols-2 gap-4">
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-1">Start Date *</label>
                <input
                  v-model="editForm.start_date"
                  type="date"
                  class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                  required
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-1">End Date *</label>
                <input
                  v-model="editForm.end_date"
                  type="date"
                  class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                  required
                />
              </div>
            </div>

            <div>
              <label class="block text-sm font-medium text-neutral-700 mb-1">Location</label>
              <input
                v-model="editForm.location"
                type="text"
                class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                placeholder="e.g., Lagos, Nigeria"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-neutral-700 mb-1">Status</label>
              <select
                v-model="editForm.status"
                class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent bg-white"
              >
                <option value="draft">Draft</option>
                <option value="upcoming">Upcoming</option>
                <option value="active">Active</option>
                <option value="completed">Completed</option>
              </select>
            </div>

            <!-- Cover Image Upload -->
            <div>
              <label class="block text-sm font-medium text-neutral-700 mb-1">Cover Image</label>
              <div class="space-y-3">
                <!-- Preview -->
                <div v-if="coverImagePreview || editForm.cover_image_url" class="relative">
                  <img
                    :src="coverImagePreview || editForm.cover_image_url"
                    alt="Cover preview"
                    class="w-full h-40 object-cover rounded-xl border border-neutral-200"
                  />
                  <button
                    v-if="editForm.cover_image_url || coverImagePreview"
                    type="button"
                    @click="removeCoverImage"
                    class="absolute top-2 right-2 w-8 h-8 bg-white/90 hover:bg-white rounded-lg flex items-center justify-center text-neutral-600 hover:text-red-600 shadow transition-colors"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                    </svg>
                  </button>
                </div>
                
                <!-- Upload button -->
                <div v-if="!coverImagePreview && !editForm.cover_image_url" class="relative">
                  <input
                    ref="coverImageInput"
                    type="file"
                    accept="image/jpeg,image/png,image/webp"
                    class="hidden"
                    @change="handleCoverImageUpload"
                  />
                  <button
                    type="button"
                    @click="$refs.coverImageInput.click()"
                    :disabled="uploadingCoverImage"
                    class="w-full h-32 border-2 border-dashed border-neutral-300 rounded-xl flex flex-col items-center justify-center gap-2 hover:border-primary-400 hover:bg-primary-50/50 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
                  >
                    <svg v-if="!uploadingCoverImage" class="w-8 h-8 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                    </svg>
                    <div v-else class="w-6 h-6 border-2 border-primary-500/30 border-t-primary-500 rounded-full animate-spin"></div>
                    <span class="text-sm text-neutral-500">
                      {{ uploadingCoverImage ? 'Uploading...' : 'Click to upload cover image' }}
                    </span>
                    <span class="text-xs text-neutral-400">JPEG, PNG or WebP • Max 10MB</span>
                  </button>
                </div>
                
                <!-- Replace button when image exists -->
                <div v-else class="flex gap-2">
                  <input
                    ref="coverImageInput"
                    type="file"
                    accept="image/jpeg,image/png,image/webp"
                    class="hidden"
                    @change="handleCoverImageUpload"
                  />
                  <button
                    type="button"
                    @click="$refs.coverImageInput.click()"
                    :disabled="uploadingCoverImage"
                    class="flex-1 px-3 py-2 text-sm border border-neutral-200 rounded-lg hover:bg-neutral-50 transition-colors disabled:opacity-50"
                  >
                    {{ uploadingCoverImage ? 'Uploading...' : 'Replace image' }}
                  </button>
                </div>
              </div>
            </div>

            <!-- Visibility Toggles -->
            <div class="space-y-4 pt-2">
              <div class="flex items-center justify-between">
                <div>
                  <label class="block text-sm font-medium text-neutral-700">Make Public</label>
                  <p class="text-xs text-neutral-500">Show this tournament on the public Tournaments page</p>
                </div>
                <button
                  type="button"
                  @click="editForm.is_public = !editForm.is_public"
                  :class="[
                    'relative inline-flex h-6 w-11 flex-shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2',
                    editForm.is_public ? 'bg-primary-600' : 'bg-neutral-200'
                  ]"
                >
                  <span
                    :class="[
                      'pointer-events-none inline-block h-5 w-5 transform rounded-full bg-white shadow ring-0 transition duration-200 ease-in-out',
                      editForm.is_public ? 'translate-x-5' : 'translate-x-0'
                    ]"
                  />
                </button>
              </div>

              <div class="flex items-center justify-between">
                <div>
                  <label class="block text-sm font-medium text-neutral-700">Featured Tournament</label>
                  <p class="text-xs text-neutral-500">Highlight this tournament at the top of listings</p>
                </div>
                <button
                  type="button"
                  @click="editForm.featured = !editForm.featured"
                  :class="[
                    'relative inline-flex h-6 w-11 flex-shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2',
                    editForm.featured ? 'bg-primary-600' : 'bg-neutral-200'
                  ]"
                >
                  <span
                    :class="[
                      'pointer-events-none inline-block h-5 w-5 transform rounded-full bg-white shadow ring-0 transition duration-200 ease-in-out',
                      editForm.featured ? 'translate-x-5' : 'translate-x-0'
                    ]"
                  />
                </button>
              </div>
            </div>

            <div class="flex justify-end gap-3 pt-4">
              <button
                type="button"
                @click="showEditModal = false"
                class="px-4 py-2.5 border border-neutral-200 rounded-xl hover:bg-neutral-50 transition-colors"
              >
                Cancel
              </button>
              <button
                type="submit"
                :disabled="updatingTournament"
                class="px-6 py-2.5 bg-gradient-to-r from-primary-500 to-emerald-600 text-white rounded-xl font-medium hover:from-primary-600 hover:to-emerald-700 transition-all disabled:opacity-50 disabled:cursor-not-allowed"
              >
                {{ updatingTournament ? 'Saving...' : 'Save Changes' }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </Teleport>
  </div>
</template>

<script setup lang="ts">
import type { ApiResponse } from '~/types/index'
import { AFRICAN_LOCATIONS } from '~/data/locations'

interface Tournament {
  id: string
  name: string
  description?: string
  start_date: string
  end_date: string
  location?: string
  status: string
  thumbnail_url?: string
  is_public?: boolean
  featured?: boolean
  cover_image_url?: string
}

interface Match {
  id: string
  title: string
  match_date: string
  location?: string
  stage?: string
  home_team?: string
  away_team?: string
  home_score?: number
  away_score?: number
  status: string
  match_number?: number
  has_video: boolean
  player_count?: number
  highlight_count?: number
}

definePageMeta({
  layout: 'admin',
  middleware: 'admin',
})

const route = useRoute()
const config = useRuntimeConfig()
const authStore = useAuthStore()
const toast = useToast()
const router = useRouter()

const tournamentId = route.params.id as string

const loading = ref(true)
const tournament = ref<Tournament | null>(null)
const matches = ref<Match[]>([])

// Venue data organized by country and city for smart filtering
const VENUE_DATA: Record<string, Record<string, string[]>> = {
  'Nigeria': {
    'Lagos': ['Teslim Balogun Stadium', 'Onikan Stadium', 'Agege Stadium', 'National Stadium Surulere'],
    'Abuja': ['MKO Abiola Stadium', 'Abuja National Stadium'],
    'Port Harcourt': ['Adokiye Amiesimaka Stadium'],
    'Uyo': ['Godswill Akpabio Stadium', 'Nest of Champions'],
    'Benin City': ['Samuel Ogbemudia Stadium'],
    'Enugu': ['Nnamdi Azikiwe Stadium'],
    'Kaduna': ['Ahmadu Bello Stadium'],
    'Kano': ['Sani Abacha Stadium'],
    'Calabar': ['U.J. Esuene Stadium'],
    'Aba': ['Enyimba Stadium'],
    'Ibadan': ['Lekan Salami Stadium', 'Adamasingba Stadium'],
    'Ilorin': ['Kwara State Stadium'],
    'Jos': ['Rwang Pam Stadium'],
  },
  'Ghana': {
    'Accra': ['Accra Sports Stadium', 'El-Wak Stadium'],
    'Kumasi': ['Baba Yara Stadium'],
    'Cape Coast': ['Cape Coast Stadium'],
    'Tamale': ['Tamale Stadium', 'Aliu Mahama Stadium'],
    'Sekondi-Takoradi': ['Essipong Stadium'],
  },
  'South Africa': {
    'Johannesburg': ['FNB Stadium', 'Ellis Park Stadium', 'Orlando Stadium', 'Bidvest Stadium'],
    'Durban': ['Moses Mabhida Stadium', 'King Zwelithini Stadium'],
    'Cape Town': ['Cape Town Stadium', 'Newlands Stadium', 'Athlone Stadium'],
    'Pretoria': ['Loftus Versfeld Stadium'],
    'Port Elizabeth': ['Nelson Mandela Bay Stadium'],
    'Bloemfontein': ['Free State Stadium'],
    'Polokwane': ['Peter Mokaba Stadium'],
    'Rustenburg': ['Royal Bafokeng Stadium'],
    'Nelspruit': ['Mbombela Stadium'],
  },
  'Egypt': {
    'Cairo': ['Cairo International Stadium', 'Al Ahly Stadium', 'Zamalek Stadium', '30 June Stadium'],
    'Alexandria': ['Borg El Arab Stadium', 'Alexandria Stadium'],
    'Suez': ['Suez Stadium'],
  },
  'Kenya': {
    'Nairobi': ['Kasarani Stadium', 'Nyayo National Stadium'],
    'Mombasa': ['Mombasa Municipal Stadium'],
    'Kisumu': ['Moi Stadium Kisumu'],
  },
  'Cameroon': {
    'Yaoundé': ['Ahmadou Ahidjo Stadium', 'Stade Omnisport de Yaoundé'],
    'Douala': ['Stade de la Réunification'],
    'Garoua': ['Roumde Adjia Stadium'],
    'Limbe': ['Limbe Stadium'],
    'Bafoussam': ['Kouekong Stadium'],
  },
  'Senegal': {
    'Dakar': ['Stade Léopold Sédar Senghor', 'Stade Demba Diop'],
    'Thiès': ['Stade Lat Dior'],
  },
  'Morocco': {
    'Casablanca': ['Stade Mohammed V', 'Complexe Sportif Mohammed V'],
    'Rabat': ['Moulay Abdellah Stadium'],
    'Marrakech': ['Grand Stade de Marrakech'],
    'Tangier': ['Ibn Batouta Stadium'],
    'Agadir': ['Adrar Stadium'],
    'Fez': ['Fez Stadium'],
  },
  'Ivory Coast': {
    'Abidjan': ['Stade Félix Houphouët-Boigny', 'Stade Olympique d\'Ebimpé'],
    'Bouaké': ['Stade de la Paix'],
    'Yamoussoukro': ['Yamoussoukro Stadium'],
  },
}

// Detect country from tournament location
const detectedCountry = computed(() => {
  const loc = tournament.value?.location?.toLowerCase() || ''
  for (const country of Object.keys(VENUE_DATA)) {
    if (loc.includes(country.toLowerCase())) {
      return country
    }
  }
  // Check cities too
  for (const [country, cities] of Object.entries(VENUE_DATA)) {
    for (const city of Object.keys(cities)) {
      if (loc.includes(city.toLowerCase())) {
        return country
      }
    }
  }
  return null
})

// Get cities for the detected country (or all if no country detected)
const availableCities = computed(() => {
  const country = detectedCountry.value
  if (country && VENUE_DATA[country]) {
    return Object.keys(VENUE_DATA[country]).sort()
  }
  // Fallback: all cities from AFRICAN_LOCATIONS
  const cities: string[] = []
  for (const data of Object.values(AFRICAN_LOCATIONS)) {
    for (const stateCities of Object.values(data.states)) {
      cities.push(...stateCities)
    }
  }
  return [...new Set(cities)].sort().slice(0, 50) // Limit to 50 for performance
})

// Get venue suggestions based on selected city
const availableVenues = computed(() => {
  const city = matchForm.city
  const country = detectedCountry.value
  
  if (city && country && VENUE_DATA[country]?.[city]) {
    return VENUE_DATA[country][city]
  }
  
  // If city selected but not in our data, check all countries
  if (city) {
    for (const cities of Object.values(VENUE_DATA)) {
      if (cities[city]) {
        return cities[city]
      }
    }
  }
  
  // If no city selected but country detected, show all venues in that country
  if (country && VENUE_DATA[country]) {
    const venues: string[] = []
    for (const cityVenues of Object.values(VENUE_DATA[country])) {
      venues.push(...cityVenues)
    }
    return venues
  }
  
  return []
})
const showEditModal = ref(false)
const showCreateMatchModal = ref(false)
const creatingMatch = ref(false)
const updatingTournament = ref(false)

// Cover image upload state
const uploadingCoverImage = ref(false)
const coverImagePreview = ref<string | null>(null)
const coverImageInput = ref<HTMLInputElement | null>(null)

const matchForm = reactive({
  title: '',
  home_team: '',
  away_team: '',
  match_date: '',
  stage: '',
  venue_name: '',
  city: '',
})

// Edit tournament form
const editForm = reactive({
  name: '',
  description: '',
  start_date: '',
  end_date: '',
  location: '',
  status: 'draft',
  is_public: false,
  featured: false,
  cover_image_url: '',
})

// Initialize edit form when modal opens
watch(showEditModal, (isOpen) => {
  if (isOpen && tournament.value) {
    editForm.name = tournament.value.name
    editForm.description = tournament.value.description || ''
    editForm.start_date = tournament.value.start_date?.split('T')[0] || ''
    editForm.end_date = tournament.value.end_date?.split('T')[0] || ''
    editForm.location = tournament.value.location || ''
    editForm.status = tournament.value.status || 'draft'
    editForm.is_public = tournament.value.is_public || false
    editForm.featured = tournament.value.featured || false
    editForm.cover_image_url = tournament.value.cover_image_url || ''
    // Reset preview when opening modal
    coverImagePreview.value = null
  }
})

// For custom venue entry when user selects "Enter custom venue..."
const customVenueName = ref('')

// Match date validation - must be within tournament dates
const matchDateMin = computed(() => {
  if (!tournament.value?.start_date) return ''
  return tournament.value.start_date.split('T')[0]
})

const matchDateMax = computed(() => {
  if (!tournament.value?.end_date) return ''
  return tournament.value.end_date.split('T')[0]
})

const isMatchDateValid = computed(() => {
  if (!matchForm.match_date || !tournament.value) return true
  const matchDate = new Date(matchForm.match_date)
  const startDate = new Date(tournament.value.start_date)
  const endDate = new Date(tournament.value.end_date)
  // Set to start/end of day for comparison
  startDate.setHours(0, 0, 0, 0)
  endDate.setHours(23, 59, 59, 999)
  return matchDate >= startDate && matchDate <= endDate
})

const isMatchFormValid = computed(() => {
  return matchForm.title.trim() && matchForm.match_date && isMatchDateValid.value
})

const matchDateError = computed(() => {
  if (!matchForm.match_date || isMatchDateValid.value) return ''
  return `Match date must be between ${formatDate(tournament.value!.start_date)} and ${formatDate(tournament.value!.end_date)}`
})

async function fetchTournament() {
  try {
    // Backend returns "events" not "tournaments"
    const response = await $fetch<ApiResponse<{ events: Tournament[] }>>('/admin/events', {
      baseURL: config.public.apiBase,
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
    })
    if (response.success && response.data) {
      tournament.value = response.data.events?.find(t => t.id === tournamentId) || null
    }
  } catch (error) {
    console.error('Failed to fetch tournament:', error)
  }
}

async function fetchMatches() {
  try {
    // Backend returns data as array directly, not { matches: [...] }
    const response = await $fetch<ApiResponse<Match[]>>(`/admin/tournaments/${tournamentId}/matches`, {
      baseURL: config.public.apiBase,
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
    })
    if (response.success && response.data) {
      matches.value = response.data || []
    }
  } catch (error) {
    console.error('Failed to fetch matches:', error)
  } finally {
    loading.value = false
  }
}

async function createMatch() {
  creatingMatch.value = true
  try {
    // Determine final venue name (handle custom entry)
    const finalVenue = matchForm.venue_name === '__custom__' 
      ? customVenueName.value 
      : matchForm.venue_name
    
    // Format date as YYYY-MM-DD (backend expects this format)
    const formattedDate = matchForm.match_date 
      ? matchForm.match_date.split('T')[0] 
      : undefined
    
    const response = await $fetch<ApiResponse<Match>>(`/admin/tournaments/${tournamentId}/matches`, {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: {
        tournament_id: tournamentId, // Backend requires this
        title: matchForm.title,
        home_team: matchForm.home_team || undefined,
        away_team: matchForm.away_team || undefined,
        match_date: formattedDate,
        stage: matchForm.stage || undefined,
        // Combine venue and city into location
        location: [finalVenue, matchForm.city].filter(Boolean).join(', ') || undefined,
      },
    })
    if (response.success && response.data) {
      toast.success('Match Created', 'Match has been added to the tournament')
      showCreateMatchModal.value = false
      Object.assign(matchForm, { title: '', home_team: '', away_team: '', match_date: '', stage: '', venue_name: '', city: '' })
      customVenueName.value = ''
      router.push(`/admin/tournaments/${tournamentId}/matches/${response.data.id}`)
    }
  } catch (error: any) {
    toast.error('Error', error.data?.message || 'Failed to create match')
  } finally {
    creatingMatch.value = false
  }
}

async function updateTournament() {
  updatingTournament.value = true
  try {
    const response = await $fetch<ApiResponse<Tournament>>(`/admin/events/${tournamentId}`, {
      baseURL: config.public.apiBase,
      method: 'PUT',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: {
        name: editForm.name,
        description: editForm.description || undefined,
        start_date: editForm.start_date,
        end_date: editForm.end_date,
        location: editForm.location || undefined,
        status: editForm.status,
        is_public: editForm.is_public,
        featured: editForm.featured,
        cover_image_url: editForm.cover_image_url || undefined,
      },
    })
    if (response.success) {
      toast.success('Tournament Updated', 'Tournament details have been saved')
      showEditModal.value = false
      await fetchTournament()
    }
  } catch (error: any) {
    toast.error('Error', error.data?.message || 'Failed to update tournament')
  } finally {
    updatingTournament.value = false
  }
}

// Cover image upload handlers
async function handleCoverImageUpload(event: Event) {
  const input = event.target as HTMLInputElement
  const file = input.files?.[0]
  if (!file) return

  // Validate file type
  const allowedTypes = ['image/jpeg', 'image/png', 'image/webp']
  if (!allowedTypes.includes(file.type)) {
    toast.error('Invalid File', 'Please upload a JPEG, PNG, or WebP image')
    return
  }

  // Validate file size (10MB max)
  if (file.size > 10 * 1024 * 1024) {
    toast.error('File Too Large', 'Maximum file size is 10MB')
    return
  }

  uploadingCoverImage.value = true

  try {
    // Step 1: Get presigned URL from backend
    const initResponse = await $fetch<ApiResponse<{
      upload_url: string
      s3_key: string
      session_id: string
    }>>('/admin/upload/init', {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: {
        upload_type: 'cover_image',
        content_type: file.type,
        file_name: file.name,
        file_size: file.size,
      },
    })

    if (!initResponse.success || !initResponse.data) {
      throw new Error('Failed to initialize upload')
    }

    const { upload_url, s3_key } = initResponse.data

    // Step 2: Upload directly to S3
    const uploadResponse = await fetch(upload_url, {
      method: 'PUT',
      headers: {
        'Content-Type': file.type,
      },
      body: file,
    })

    if (!uploadResponse.ok) {
      throw new Error('Failed to upload to S3')
    }

    // Step 3: Update the form with the S3 URL
    editForm.cover_image_url = `s3://unicorn-sport-media/${s3_key}`

    // Create preview URL for display
    coverImagePreview.value = URL.createObjectURL(file)

    toast.success('Image Uploaded', 'Cover image uploaded successfully')
  } catch (error: any) {
    console.error('Upload failed:', error)
    toast.error('Upload Failed', error.message || 'Failed to upload cover image')
  } finally {
    uploadingCoverImage.value = false
    // Reset input
    if (input) input.value = ''
  }
}

function removeCoverImage() {
  editForm.cover_image_url = ''
  coverImagePreview.value = null
}

function formatDateRange(start: string, end: string): string {
  const startDate = new Date(start)
  const endDate = new Date(end)
  const opts: Intl.DateTimeFormatOptions = { month: 'short', day: 'numeric' }
  return `${startDate.toLocaleDateString('en-US', opts)} - ${endDate.toLocaleDateString('en-US', opts)}, ${endDate.getFullYear()}`
}

function formatDate(date: string): string {
  return new Date(date).toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' })
}

function formatStage(stage: string): string {
  const stages: Record<string, string> = {
    group: 'Group',
    round_of_16: 'R16',
    quarterfinal: 'QF',
    semifinal: 'SF',
    final: 'Final',
  }
  return stages[stage] || stage
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

function getMatchStatusClass(status: string): string {
  const classes: Record<string, string> = {
    scheduled: 'bg-blue-100 text-blue-700',
    in_progress: 'bg-amber-100 text-amber-700',
    completed: 'bg-emerald-100 text-emerald-700',
    cancelled: 'bg-neutral-100 text-neutral-700',
  }
  return classes[status] || classes.scheduled!
}

onMounted(async () => {
  await Promise.all([fetchTournament(), fetchMatches()])
})
</script>
