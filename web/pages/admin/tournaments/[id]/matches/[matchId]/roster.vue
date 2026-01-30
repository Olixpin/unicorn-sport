<template>
  <div class="min-h-screen bg-neutral-50">
    <!-- Header -->
    <div class="bg-white border-b border-neutral-200 sticky top-0 z-20">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-4">
            <NuxtLink
              :to="`/admin/tournaments/${tournamentId}/matches/${matchId}`"
              class="flex items-center gap-2 text-neutral-600 hover:text-neutral-900 transition-colors"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
              </svg>
              Back to Match
            </NuxtLink>
          </div>
          <div class="flex items-center gap-3">
            <span class="text-sm text-neutral-500">
              {{ rosterPlayers.length }} players in roster
            </span>
            <button
              v-if="hasUnsavedChanges"
              @click="saveAllChanges"
              :disabled="saving"
              class="px-4 py-2 bg-primary-500 text-white rounded-lg hover:bg-primary-600 transition-colors disabled:opacity-50 flex items-center gap-2"
            >
              <svg v-if="saving" class="w-4 h-4 animate-spin" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              {{ saving ? 'Saving...' : 'Save Changes' }}
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="flex justify-center items-center h-64">
      <div class="animate-spin rounded-full h-12 w-12 border-4 border-primary-500 border-t-transparent"></div>
    </div>

    <!-- Error State -->
    <div v-else-if="loadError || !match" class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16 text-center">
      <div class="w-16 h-16 mx-auto mb-4 rounded-full bg-red-100 flex items-center justify-center">
        <svg class="w-8 h-8 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
        </svg>
      </div>
      <h2 class="text-xl font-semibold text-neutral-900 mb-2">{{ loadError ? 'Failed to load roster' : 'Match not found' }}</h2>
      <p class="text-neutral-500 mb-4">{{ loadError || 'Unable to load match data. Please try again.' }}</p>
      <NuxtLink
        :to="`/admin/tournaments/${tournamentId}/matches/${matchId}`"
        class="inline-flex items-center gap-2 px-4 py-2 bg-primary-500 text-white rounded-lg hover:bg-primary-600 transition-colors"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
        </svg>
        Back to Match
      </NuxtLink>
    </div>

    <div v-else class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <!-- Match Info Banner -->
      <div class="bg-gradient-to-r from-primary-600 to-emerald-600 rounded-2xl p-4 sm:p-6 mb-6 sm:mb-8 text-white">
        <h1 class="text-xl sm:text-2xl font-bold">{{ match?.title }}</h1>
        <div class="flex flex-wrap items-center gap-x-3 gap-y-1 mt-2 text-sm text-white/80">
          <span class="flex items-center gap-1">
            <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
            </svg>
            {{ formatDate(match?.match_date) }}
          </span>
          <span v-if="match?.location" class="flex items-center gap-1">
            <span class="text-white/40">•</span>
            {{ match.location }}
          </span>
        </div>
        <div v-if="match?.home_team && match?.away_team" class="mt-2 text-base sm:text-lg font-medium">
          {{ match.home_team }} vs {{ match.away_team }}
        </div>
        
        <!-- Stats Cards -->
        <div class="flex gap-2 sm:gap-3 mt-4">
          <div class="flex-1 text-center px-3 py-2 bg-white/15 rounded-xl">
            <div class="text-2xl sm:text-3xl font-bold">{{ rosterPlayers.length }}</div>
            <div class="text-xs text-white/70">Players</div>
          </div>
          <div class="flex-1 text-center px-3 py-2 bg-white/15 rounded-xl">
            <div class="text-2xl sm:text-3xl font-bold">{{ starters.length }}</div>
            <div class="text-xs text-white/70">Starters</div>
          </div>
          <div class="flex-1 text-center px-3 py-2 bg-white/15 rounded-xl">
            <div class="text-2xl sm:text-3xl font-bold">{{ totalMinutes }}</div>
            <div class="text-xs text-white/70">Minutes</div>
          </div>
        </div>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- Left Column: Add Players -->
        <div class="lg:col-span-1 space-y-6">
          <!-- Search Box -->
          <div class="bg-white rounded-2xl border border-neutral-200 p-6">
            <h2 class="text-lg font-semibold text-neutral-900 mb-4 flex items-center gap-2">
              <svg class="w-5 h-5 text-primary-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
              Add Players
            </h2>
            
            <div class="relative">
              <input
                v-model="searchQuery"
                type="text"
                placeholder="Search players by name..."
                class="w-full px-4 py-3 pl-10 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                @input="debouncedSearch"
              />
              <svg class="w-5 h-5 text-neutral-400 absolute left-3 top-1/2 -translate-y-1/2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </div>

            <!-- Search Results -->
            <div v-if="searchResults.length > 0" class="mt-4 space-y-2 max-h-64 overflow-y-auto">
              <button
                v-for="player in searchResults"
                :key="player.id"
                @click="quickAddPlayer(player)"
                :disabled="isPlayerInRoster(player.id) || addingPlayerId === player.id"
                class="w-full flex items-center gap-3 p-3 rounded-xl transition-all text-left"
                :class="isPlayerInRoster(player.id) 
                  ? 'bg-neutral-100 opacity-50 cursor-not-allowed' 
                  : addingPlayerId === player.id
                    ? 'bg-primary-50 ring-2 ring-primary-200'
                    : 'bg-neutral-50 hover:bg-primary-50 hover:ring-2 hover:ring-primary-200'"
              >
                <div class="w-10 h-10 rounded-full overflow-hidden bg-gradient-to-br from-primary-500 to-emerald-600 flex items-center justify-center text-white text-sm font-medium flex-shrink-0">
                  <img v-if="player.profile_photo_url" :src="player.profile_photo_url" class="w-full h-full object-cover" />
                  <span v-else>{{ getInitials(player.first_name, player.last_name) }}</span>
                </div>
                <div class="flex-1 min-w-0">
                  <p class="font-medium text-neutral-900 truncate">{{ player.first_name }} {{ player.last_name }}</p>
                  <p class="text-xs text-neutral-500">{{ player.position }} • {{ player.academy?.name || 'No academy' }}</p>
                </div>
                <div v-if="addingPlayerId === player.id" class="flex-shrink-0">
                  <div class="w-5 h-5 border-2 border-primary-500 border-t-transparent rounded-full animate-spin"></div>
                </div>
                <div v-else-if="isPlayerInRoster(player.id)" class="text-xs text-emerald-600 font-medium">
                  ✓ Added
                </div>
                <svg v-else class="w-5 h-5 text-primary-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                </svg>
              </button>
            </div>

            <p v-else-if="searchQuery && !searching" class="mt-4 text-sm text-neutral-500 text-center py-4">
              No players found for "{{ searchQuery }}"
            </p>

            <div v-if="searching" class="mt-4 flex justify-center py-4">
              <div class="animate-spin rounded-full h-6 w-6 border-2 border-primary-500 border-t-transparent"></div>
            </div>
          </div>

          <!-- Suggested Players -->
          <div class="bg-white rounded-2xl border border-neutral-200 p-6">
            <h2 class="text-lg font-semibold text-neutral-900 mb-4 flex items-center gap-2">
              <svg class="w-5 h-5 text-amber-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z" />
              </svg>
              Suggested
            </h2>

            <!-- From Tournament -->
            <div v-if="tournamentPlayers.length > 0" class="mb-4">
              <p class="text-xs font-medium text-neutral-500 uppercase tracking-wide mb-2">From This Tournament</p>
              <div class="flex flex-wrap gap-2">
                <button
                  v-for="player in tournamentPlayers.slice(0, 6)"
                  :key="player.id"
                  @click="quickAddPlayer(player)"
                  :disabled="isPlayerInRoster(player.id) || addingPlayerId === player.id"
                  class="flex items-center gap-2 px-3 py-2 rounded-full text-sm transition-all"
                  :class="isPlayerInRoster(player.id) 
                    ? 'bg-emerald-100 text-emerald-700' 
                    : addingPlayerId === player.id
                      ? 'bg-primary-100 text-primary-700'
                      : 'bg-neutral-100 hover:bg-primary-100 hover:text-primary-700'"
                >
                  <span>{{ player.first_name }} {{ player.last_name?.charAt(0) }}.</span>
                  <div v-if="addingPlayerId === player.id" class="w-3 h-3 border-2 border-primary-500 border-t-transparent rounded-full animate-spin"></div>
                  <span v-else-if="isPlayerInRoster(player.id)" class="text-emerald-600">✓</span>
                  <svg v-else class="w-4 h-4 text-primary-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                  </svg>
                </button>
              </div>
            </div>

            <!-- Recent Players -->
            <div v-if="recentPlayers.length > 0">
              <p class="text-xs font-medium text-neutral-500 uppercase tracking-wide mb-2">Recently Used</p>
              <div class="flex flex-wrap gap-2">
                <button
                  v-for="player in recentPlayers.slice(0, 6)"
                  :key="player.id"
                  @click="quickAddPlayer(player)"
                  :disabled="isPlayerInRoster(player.id) || addingPlayerId === player.id"
                  class="flex items-center gap-2 px-3 py-2 rounded-full text-sm transition-all"
                  :class="isPlayerInRoster(player.id) 
                    ? 'bg-emerald-100 text-emerald-700' 
                    : addingPlayerId === player.id
                      ? 'bg-primary-100 text-primary-700'
                      : 'bg-neutral-100 hover:bg-primary-100 hover:text-primary-700'"
                >
                  <span>{{ player.first_name }} {{ player.last_name?.charAt(0) }}.</span>
                  <div v-if="addingPlayerId === player.id" class="w-3 h-3 border-2 border-primary-500 border-t-transparent rounded-full animate-spin"></div>
                  <span v-else-if="isPlayerInRoster(player.id)" class="text-emerald-600">✓</span>
                  <svg v-else class="w-4 h-4 text-primary-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                  </svg>
                </button>
              </div>
            </div>

            <p v-if="tournamentPlayers.length === 0 && recentPlayers.length === 0" class="text-sm text-neutral-500 text-center py-4">
              No suggestions available. Use search to find players.
            </p>
          </div>
        </div>

        <!-- Right Column: Roster List -->
        <div class="lg:col-span-2">
          <!-- Roster Header -->
          <div class="flex items-center justify-between mb-4">
            <h2 class="text-base sm:text-lg font-semibold text-neutral-900 flex items-center gap-2">
              <svg class="w-5 h-5 text-primary-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
              </svg>
              Match Roster
            </h2>
            
            <!-- View Toggle -->
            <div class="flex items-center gap-1 bg-neutral-100 rounded-lg p-1">
              <button
                @click="viewMode = 'list'"
                class="px-3 py-1.5 rounded-md text-sm font-medium transition-all"
                :class="viewMode === 'list' ? 'bg-white shadow text-primary-600' : 'text-neutral-600 hover:text-neutral-900'"
              >
                List
              </button>
              <button
                @click="viewMode = 'formation'"
                class="px-3 py-1.5 rounded-md text-sm font-medium transition-all"
                :class="viewMode === 'formation' ? 'bg-white shadow text-primary-600' : 'text-neutral-600 hover:text-neutral-900'"
              >
                Formation
              </button>
            </div>
          </div>

          <!-- List View -->
          <div v-if="viewMode === 'list'" class="bg-white rounded-2xl border border-neutral-200 divide-y divide-neutral-100">
              <!-- Position Groups -->
              <div v-for="group in positionGroups" :key="group.name" class="px-6 py-4">
                <div class="flex items-center gap-2 mb-3">
                  <span 
                    class="w-3 h-3 rounded-full"
                    :class="{
                      'bg-amber-500': group.name === 'Goalkeepers',
                      'bg-blue-500': group.name === 'Defenders',
                      'bg-emerald-500': group.name === 'Midfielders',
                      'bg-rose-500': group.name === 'Forwards'
                    }"
                  ></span>
                  <h3 class="text-sm font-semibold text-neutral-700 uppercase tracking-wide">
                    {{ group.name }} ({{ group.players.length }})
                  </h3>
                </div>

                <div v-if="group.players.length > 0" class="space-y-2">
                  <div
                    v-for="rp in group.players"
                    :key="rp.id"
                    class="p-3 sm:p-4 bg-neutral-50 rounded-xl hover:bg-neutral-100 transition-colors group"
                  >
                    <!-- Mobile: Two rows, Desktop: Single row -->
                    <div class="flex items-center gap-3">
                      <!-- Player Avatar -->
                      <div class="w-10 h-10 sm:w-12 sm:h-12 rounded-full overflow-hidden bg-gradient-to-br from-primary-500 to-emerald-600 flex items-center justify-center text-white text-sm font-medium flex-shrink-0">
                        <img v-if="rp.player?.profile_photo_url" :src="rp.player.profile_photo_url" class="w-full h-full object-cover" />
                        <span v-else>{{ getInitials(rp.player?.first_name, rp.player?.last_name) }}</span>
                      </div>
                      
                      <!-- Name & Position -->
                      <div class="flex-1 min-w-0">
                        <p class="font-semibold text-neutral-900 text-sm sm:text-base truncate">
                          {{ rp.player?.first_name }} {{ rp.player?.last_name }}
                        </p>
                        <p class="text-xs sm:text-sm text-neutral-500">
                          {{ rp.position_played || rp.player?.position || 'Position TBD' }}
                        </p>
                      </div>

                      <!-- Desktop Stats -->
                      <div class="hidden sm:flex items-center gap-3">
                        <div class="flex items-center gap-1">
                          <span class="text-xs text-neutral-500">Minutes</span>
                          <input
                            v-model.number="rp.minutes_played"
                            type="number"
                            min="0"
                            max="120"
                            class="w-16 px-2 py-1 text-center border border-neutral-200 rounded-lg text-sm focus:ring-2 focus:ring-primary-500"
                            placeholder="90"
                            @change="markAsChanged(rp)"
                          />
                        </div>
                      </div>

                      <!-- Remove Button -->
                      <button
                        @click="removePlayer(rp)"
                        :disabled="removingPlayerId === rp.player_id"
                        class="p-2 text-neutral-400 hover:text-rose-500 sm:opacity-0 sm:group-hover:opacity-100 transition-all disabled:opacity-100"
                        title="Remove from roster"
                      >
                        <div v-if="removingPlayerId === rp.player_id" class="w-5 h-5 border-2 border-rose-400 border-t-transparent rounded-full animate-spin"></div>
                        <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                        </svg>
                      </button>
                    </div>
                    
                    <!-- Mobile Stats Row -->
                    <div class="flex items-center gap-2 mt-3 sm:hidden">
                      <div class="flex items-center gap-1 flex-1">
                        <span class="text-[10px] text-neutral-500 uppercase">Minutes Played</span>
                        <input
                          v-model.number="rp.minutes_played"
                          type="number"
                          min="0"
                          max="120"
                          class="w-full px-2 py-1.5 text-center border border-neutral-200 rounded-lg text-sm focus:ring-2 focus:ring-primary-500"
                          placeholder="90"
                          @change="markAsChanged(rp)"
                        />
                      </div>
                    </div>
                  </div>
                </div>

                <p v-else class="text-sm text-neutral-400 py-2">No {{ group.name.toLowerCase() }} added</p>
              </div>

              <!-- Empty State -->
              <div v-if="rosterPlayers.length === 0" class="px-6 py-16 text-center">
                <div class="w-20 h-20 bg-neutral-100 rounded-full flex items-center justify-center mx-auto mb-4">
                  <svg class="w-10 h-10 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
                  </svg>
                </div>
                <h3 class="text-lg font-semibold text-neutral-900 mb-2">No players yet</h3>
                <p class="text-neutral-500 mb-4">Search or select from suggestions to add players</p>
              </div>
            </div>

            <!-- Formation View -->
            <div v-if="viewMode === 'formation'" class="space-y-4">
              <!-- Formation Selector -->
              <div class="flex items-center justify-center gap-3">
                <label class="text-sm font-medium text-neutral-700">Formation:</label>
                <div class="relative">
                  <select
                    v-model="selectedFormation"
                    class="appearance-none pl-4 pr-10 py-2 bg-white border border-neutral-200 rounded-xl text-sm font-medium focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                  >
                    <option v-for="f in formations" :key="f.value" :value="f.value">{{ f.label }}</option>
                  </select>
                  <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                    <svg class="w-4 h-4 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                    </svg>
                  </div>
                </div>
              </div>

              <!-- Main Grid: Pitch + Bench -->
              <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
                <!-- Football Pitch -->
                <div class="lg:col-span-2">
                  <div 
                    class="relative aspect-[3/4] max-w-lg mx-auto bg-gradient-to-b from-emerald-600 to-emerald-700 rounded-2xl overflow-hidden shadow-xl"
                    @dragover.prevent
                    @drop="handlePitchDrop"
                  >
                    <!-- Pitch markings -->
                    <div class="absolute inset-0 pointer-events-none">
                      <!-- Outer border -->
                      <div class="absolute inset-3 border-2 border-white/40 rounded-lg"></div>
                      <!-- Center line -->
                      <div class="absolute top-1/2 left-3 right-3 h-px bg-white/40"></div>
                      <!-- Center circle -->
                      <div class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-28 h-28 border-2 border-white/40 rounded-full"></div>
                      <div class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-2 h-2 bg-white/40 rounded-full"></div>
                      <!-- Top penalty area -->
                      <div class="absolute top-3 left-1/2 -translate-x-1/2 w-40 h-16 border-2 border-t-0 border-white/40"></div>
                      <div class="absolute top-3 left-1/2 -translate-x-1/2 w-20 h-8 border-2 border-t-0 border-white/40"></div>
                      <!-- Bottom penalty area -->
                      <div class="absolute bottom-3 left-1/2 -translate-x-1/2 w-40 h-16 border-2 border-b-0 border-white/40"></div>
                      <div class="absolute bottom-3 left-1/2 -translate-x-1/2 w-20 h-8 border-2 border-b-0 border-white/40"></div>
                      <!-- Corner arcs -->
                      <div class="absolute top-3 left-3 w-4 h-4 border-r-2 border-white/40 rounded-br-full"></div>
                      <div class="absolute top-3 right-3 w-4 h-4 border-l-2 border-white/40 rounded-bl-full"></div>
                      <div class="absolute bottom-3 left-3 w-4 h-4 border-r-2 border-white/40 rounded-tr-full"></div>
                      <div class="absolute bottom-3 right-3 w-4 h-4 border-l-2 border-white/40 rounded-tl-full"></div>
                    </div>

                    <!-- Formation Positions - Dynamically positioned based on formation -->
                    <div class="absolute inset-6 flex flex-col justify-between py-4">
                      <!-- Forwards Row -->
                      <div class="flex justify-center gap-3">
                        <template v-for="(player, idx) in formationByPosition.forwards.slice(0, getFormationRow(0))" :key="player.id">
                          <div
                            draggable="true"
                            @dragstart="startDrag(player)"
                            @dragend="endDrag"
                            class="relative group cursor-grab active:cursor-grabbing"
                          >
                            <div 
                              class="w-12 h-12 md:w-14 md:h-14 bg-white rounded-full flex items-center justify-center text-xs font-bold shadow-lg transition-all hover:scale-110 hover:shadow-xl ring-2 ring-white/50"
                              :class="{ 'ring-primary-400 ring-4': draggingPlayer?.id === player.id }"
                            >
                              <img v-if="player.player?.profile_photo_url" :src="player.player.profile_photo_url" class="w-full h-full rounded-full object-cover" />
                              <span v-else class="text-emerald-700">{{ getInitials(player.player?.first_name, player.player?.last_name) }}</span>
                            </div>
                            <div v-if="player.jersey_number" class="absolute -top-1 -right-1 w-5 h-5 bg-primary-500 text-white text-[10px] font-bold rounded-full flex items-center justify-center">
                              {{ player.jersey_number }}
                            </div>
                            <!-- Tooltip -->
                            <div class="absolute -bottom-8 left-1/2 -translate-x-1/2 opacity-0 group-hover:opacity-100 transition-opacity whitespace-nowrap bg-black/80 text-white text-xs px-2 py-1 rounded pointer-events-none z-10">
                              {{ player.player?.first_name }} {{ player.player?.last_name }}
                            </div>
                          </div>
                        </template>
                        <template v-for="n in Math.max(0, getFormationRow(0) - formationByPosition.forwards.length)" :key="'fw-empty-' + n">
                          <div class="w-12 h-12 md:w-14 md:h-14 bg-white/20 rounded-full flex items-center justify-center text-white/50 text-xs border-2 border-dashed border-white/30">
                            FW
                          </div>
                        </template>
                      </div>

                      <!-- Attacking Midfielders Row (for formations like 4-2-3-1) -->
                      <div v-if="selectedFormation === '4-2-3-1'" class="flex justify-center gap-3">
                        <template v-for="(player, idx) in formationByPosition.midfielders.slice(0, 3)" :key="player.id">
                          <div
                            draggable="true"
                            @dragstart="startDrag(player)"
                            @dragend="endDrag"
                            class="relative group cursor-grab active:cursor-grabbing"
                          >
                            <div class="w-12 h-12 md:w-14 md:h-14 bg-white rounded-full flex items-center justify-center text-xs font-bold shadow-lg transition-all hover:scale-110 ring-2 ring-white/50">
                              <img v-if="player.player?.profile_photo_url" :src="player.player.profile_photo_url" class="w-full h-full rounded-full object-cover" />
                              <span v-else class="text-emerald-700">{{ getInitials(player.player?.first_name, player.player?.last_name) }}</span>
                            </div>
                            <div v-if="player.jersey_number" class="absolute -top-1 -right-1 w-5 h-5 bg-primary-500 text-white text-[10px] font-bold rounded-full flex items-center justify-center">
                              {{ player.jersey_number }}
                            </div>
                            <div class="absolute -bottom-8 left-1/2 -translate-x-1/2 opacity-0 group-hover:opacity-100 transition-opacity whitespace-nowrap bg-black/80 text-white text-xs px-2 py-1 rounded pointer-events-none z-10">
                              {{ player.player?.first_name }} {{ player.player?.last_name }}
                            </div>
                          </div>
                        </template>
                      </div>

                      <!-- Midfielders Row -->
                      <div class="flex justify-center gap-3">
                        <template v-for="(player, idx) in formationByPosition.midfielders.slice(selectedFormation === '4-2-3-1' ? 3 : 0)" :key="player.id">
                          <div
                            draggable="true"
                            @dragstart="startDrag(player)"
                            @dragend="endDrag"
                            class="relative group cursor-grab active:cursor-grabbing"
                          >
                            <div class="w-12 h-12 md:w-14 md:h-14 bg-white rounded-full flex items-center justify-center text-xs font-bold shadow-lg transition-all hover:scale-110 ring-2 ring-white/50">
                              <img v-if="player.player?.profile_photo_url" :src="player.player.profile_photo_url" class="w-full h-full rounded-full object-cover" />
                              <span v-else class="text-emerald-700">{{ getInitials(player.player?.first_name, player.player?.last_name) }}</span>
                            </div>
                            <div v-if="player.jersey_number" class="absolute -top-1 -right-1 w-5 h-5 bg-primary-500 text-white text-[10px] font-bold rounded-full flex items-center justify-center">
                              {{ player.jersey_number }}
                            </div>
                            <div class="absolute -bottom-8 left-1/2 -translate-x-1/2 opacity-0 group-hover:opacity-100 transition-opacity whitespace-nowrap bg-black/80 text-white text-xs px-2 py-1 rounded pointer-events-none z-10">
                              {{ player.player?.first_name }} {{ player.player?.last_name }}
                            </div>
                          </div>
                        </template>
                        <template v-for="n in Math.max(0, getFormationRow(1) - formationByPosition.midfielders.length)" :key="'mf-empty-' + n">
                          <div class="w-12 h-12 md:w-14 md:h-14 bg-white/20 rounded-full flex items-center justify-center text-white/50 text-xs border-2 border-dashed border-white/30">
                            MF
                          </div>
                        </template>
                      </div>

                      <!-- Defenders Row -->
                      <div class="flex justify-center gap-3">
                        <template v-for="player in formationByPosition.defenders" :key="player.id">
                          <div
                            draggable="true"
                            @dragstart="startDrag(player)"
                            @dragend="endDrag"
                            class="relative group cursor-grab active:cursor-grabbing"
                          >
                            <div class="w-12 h-12 md:w-14 md:h-14 bg-white rounded-full flex items-center justify-center text-xs font-bold shadow-lg transition-all hover:scale-110 ring-2 ring-white/50">
                              <img v-if="player.player?.profile_photo_url" :src="player.player.profile_photo_url" class="w-full h-full rounded-full object-cover" />
                              <span v-else class="text-emerald-700">{{ getInitials(player.player?.first_name, player.player?.last_name) }}</span>
                            </div>
                            <div v-if="player.jersey_number" class="absolute -top-1 -right-1 w-5 h-5 bg-primary-500 text-white text-[10px] font-bold rounded-full flex items-center justify-center">
                              {{ player.jersey_number }}
                            </div>
                            <div class="absolute -bottom-8 left-1/2 -translate-x-1/2 opacity-0 group-hover:opacity-100 transition-opacity whitespace-nowrap bg-black/80 text-white text-xs px-2 py-1 rounded pointer-events-none z-10">
                              {{ player.player?.first_name }} {{ player.player?.last_name }}
                            </div>
                          </div>
                        </template>
                        <template v-for="n in Math.max(0, getFormationRow(2) - formationByPosition.defenders.length)" :key="'df-empty-' + n">
                          <div class="w-12 h-12 md:w-14 md:h-14 bg-white/20 rounded-full flex items-center justify-center text-white/50 text-xs border-2 border-dashed border-white/30">
                            DF
                          </div>
                        </template>
                      </div>

                      <!-- Goalkeeper Row -->
                      <div class="flex justify-center">
                        <template v-for="player in formationByPosition.goalkeepers" :key="player.id">
                          <div
                            draggable="true"
                            @dragstart="startDrag(player)"
                            @dragend="endDrag"
                            class="relative group cursor-grab active:cursor-grabbing"
                          >
                            <div class="w-12 h-12 md:w-14 md:h-14 bg-amber-400 rounded-full flex items-center justify-center text-xs font-bold shadow-lg transition-all hover:scale-110 ring-2 ring-amber-300">
                              <img v-if="player.player?.profile_photo_url" :src="player.player.profile_photo_url" class="w-full h-full rounded-full object-cover" />
                              <span v-else class="text-amber-900">{{ getInitials(player.player?.first_name, player.player?.last_name) }}</span>
                            </div>
                            <div v-if="player.jersey_number" class="absolute -top-1 -right-1 w-5 h-5 bg-amber-600 text-white text-[10px] font-bold rounded-full flex items-center justify-center">
                              {{ player.jersey_number }}
                            </div>
                            <div class="absolute -bottom-8 left-1/2 -translate-x-1/2 opacity-0 group-hover:opacity-100 transition-opacity whitespace-nowrap bg-black/80 text-white text-xs px-2 py-1 rounded pointer-events-none z-10">
                              {{ player.player?.first_name }} {{ player.player?.last_name }}
                            </div>
                          </div>
                        </template>
                        <template v-if="formationByPosition.goalkeepers.length === 0">
                          <div class="w-12 h-12 md:w-14 md:h-14 bg-amber-400/30 rounded-full flex items-center justify-center text-white/50 text-xs border-2 border-dashed border-amber-400/50">
                            GK
                          </div>
                        </template>
                      </div>
                    </div>

                    <!-- Formation Label -->
                    <div class="absolute bottom-2 right-2 bg-black/30 text-white text-xs px-2 py-1 rounded-full">
                      {{ selectedFormation }}
                    </div>
                  </div>
                </div>

                <!-- Substitutes Bench & Staff -->
                <div class="space-y-4">
                  <!-- Substitutes Bench -->
                  <div>
                    <div class="flex items-center justify-between mb-3">
                      <h3 class="font-semibold text-neutral-900 flex items-center gap-2 text-sm">
                        <svg class="w-4 h-4 text-neutral-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0z" />
                        </svg>
                        Substitutes
                      </h3>
                      <span class="text-xs text-neutral-500">{{ substitutes.length }}</span>
                    </div>
                    
                    <div 
                      class="space-y-2 min-h-[80px] border-2 border-dashed border-neutral-200 rounded-xl p-2"
                      :class="{ 'border-primary-400 bg-primary-50': draggingPlayer }"
                      @dragover.prevent
                      @drop="handleBenchDrop"
                    >
                      <div
                        v-for="player in substitutes"
                        :key="player.id"
                        draggable="true"
                        @dragstart="startDrag(player)"
                        @dragend="endDrag"
                        class="flex items-center gap-3 p-2 bg-neutral-50 rounded-lg cursor-grab active:cursor-grabbing hover:bg-neutral-100 transition-all"
                      >
                        <div class="w-9 h-9 rounded-full overflow-hidden bg-gradient-to-br from-neutral-400 to-neutral-500 flex items-center justify-center text-white text-xs font-medium flex-shrink-0">
                          <img v-if="player.player?.profile_photo_url" :src="player.player.profile_photo_url" class="w-full h-full object-cover" />
                          <span v-else>{{ getInitials(player.player?.first_name, player.player?.last_name) }}</span>
                        </div>
                        <div class="flex-1 min-w-0">
                          <p class="font-medium text-neutral-900 text-sm truncate">{{ player.player?.first_name }} {{ player.player?.last_name }}</p>
                          <p class="text-xs text-neutral-500">{{ player.position_played || player.player?.position }}</p>
                        </div>
                      </div>
                      
                      <p v-if="substitutes.length === 0" class="text-center text-neutral-400 text-xs py-4">
                        Drag players here
                      </p>
                    </div>
                  </div>

                  <!-- Coaching Staff - simplified -->
                  <div>
                    <h3 class="font-semibold text-neutral-900 flex items-center gap-2 mb-3 text-sm">
                      <svg class="w-4 h-4 text-neutral-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                      </svg>
                      Staff
                    </h3>
                    
                    <div class="space-y-2">
                      <div class="flex items-center gap-3 p-2 border border-dashed border-neutral-200 rounded-lg">
                        <div class="w-8 h-8 rounded-full bg-neutral-100 flex items-center justify-center text-neutral-400">
                          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
                          </svg>
                        </div>
                        <p class="text-xs text-neutral-400">Add Head Coach</p>
                      </div>
                      <div class="flex items-center gap-3 p-2 border border-dashed border-neutral-200 rounded-lg">
                        <div class="w-8 h-8 rounded-full bg-neutral-100 flex items-center justify-center text-neutral-400">
                          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
                          </svg>
                        </div>
                        <p class="text-xs text-neutral-400">Add Assistant</p>
                      </div>
                    </div>
                  </div>

                  <!-- Match Stats - compact inline -->
                  <div class="grid grid-cols-3 gap-2 pt-2 border-t border-neutral-100">
                    <div class="text-center py-2">
                      <p class="text-lg font-bold text-primary-600">{{ starters.length }}</p>
                      <p class="text-[10px] text-neutral-500 uppercase">Starters</p>
                    </div>
                    <div class="text-center py-2">
                      <p class="text-lg font-bold text-primary-600">{{ substitutes.length }}</p>
                      <p class="text-[10px] text-neutral-500 uppercase">Bench</p>
                    </div>
                    <div class="text-center py-2">
                      <p class="text-lg font-bold text-primary-600">{{ totalMinutes }}</p>
                      <p class="text-[10px] text-neutral-500 uppercase">Minutes</p>
                    </div>
                  </div>
                </div>
              </div>

              <p class="text-center text-sm text-neutral-500 mt-4">
                Drag players between pitch and bench. Click a player to edit details.
              </p>
            </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { ApiResponse } from '~/types/index'

interface Player {
  id: string
  first_name: string
  last_name: string
  position: string
  profile_photo_url?: string
  academy?: { name: string }
}

interface RosterPlayer {
  id: string
  match_id: string
  player_id: string
  position_played?: string
  minutes_played?: number
  player?: Player
  is_starter: boolean
  jersey_number?: number
  formation_x?: number
  formation_y?: number
  subbed_in_for?: string
  subbed_in_at?: number
  subbed_out_at?: number
  _isNew?: boolean
  _isChanged?: boolean
}

interface Match {
  id: string
  title: string
  match_date: string
  location?: string
  home_team?: string
  away_team?: string
  tournament_id?: string
}

definePageMeta({
  layout: 'admin',
  middleware: 'admin',
})

const route = useRoute()
const config = useRuntimeConfig()
const authStore = useAuthStore()
const toast = useToast()

// Route params
const tournamentId = computed(() => route.params.id as string)
const matchId = computed(() => route.params.matchId as string)

// State
const loading = ref(true)
const loadError = ref<string | null>(null)
const saving = ref(false)
const searching = ref(false)
const addingPlayerId = ref<string | null>(null) // Track which player is being added
const removingPlayerId = ref<string | null>(null) // Track which player is being removed
const match = ref<Match | null>(null)
const rosterPlayers = ref<RosterPlayer[]>([])
const searchQuery = ref('')
const searchResults = ref<Player[]>([])
const tournamentPlayers = ref<Player[]>([])
const recentPlayers = ref<Player[]>([])
const viewMode = ref<'list' | 'formation'>('list')
const selectedFormation = ref('4-3-3')
const draggingPlayer = ref<RosterPlayer | null>(null)

// Available formations
const formations = [
  { value: '4-3-3', label: '4-3-3', rows: [1, 4, 3, 3] },
  { value: '4-4-2', label: '4-4-2', rows: [1, 4, 4, 2] },
  { value: '3-5-2', label: '3-5-2', rows: [1, 3, 5, 2] },
  { value: '4-2-3-1', label: '4-2-3-1', rows: [1, 4, 2, 3, 1] },
  { value: '5-3-2', label: '5-3-2', rows: [1, 5, 3, 2] },
  { value: '3-4-3', label: '3-4-3', rows: [1, 3, 4, 3] },
]

// Debounce search
let searchTimeout: ReturnType<typeof setTimeout>
const debouncedSearch = () => {
  clearTimeout(searchTimeout)
  searchTimeout = setTimeout(() => searchPlayers(), 300)
}

// Computed
const hasUnsavedChanges = computed(() => 
  rosterPlayers.value.some(rp => rp._isNew || rp._isChanged)
)

const totalMinutes = computed(() => 
  rosterPlayers.value.reduce((sum, rp) => sum + (rp.minutes_played || 0), 0)
)

// Split into starters and substitutes
const starters = computed(() => 
  rosterPlayers.value.filter(rp => rp.is_starter)
)

const substitutes = computed(() => 
  rosterPlayers.value.filter(rp => !rp.is_starter)
)

const positionGroups = computed(() => {
  const groups: Array<{ name: string; positions: string[]; players: RosterPlayer[] }> = [
    { name: 'Goalkeepers', positions: ['GK', 'Goalkeeper'], players: [] },
    { name: 'Defenders', positions: ['CB', 'LB', 'RB', 'LWB', 'RWB', 'Defender', 'Centre-Back', 'Full-Back'], players: [] },
    { name: 'Midfielders', positions: ['CM', 'CDM', 'CAM', 'LM', 'RM', 'Midfielder', 'Central Midfielder', 'Defensive Midfielder', 'Attacking Midfielder'], players: [] },
    { name: 'Forwards', positions: ['ST', 'CF', 'LW', 'RW', 'Forward', 'Striker', 'Winger'], players: [] },
  ]

  for (const rp of rosterPlayers.value) {
    const pos = (rp.position_played || rp.player?.position || '').toLowerCase()
    let assigned = false
    for (const group of groups) {
      if (group.positions.some(p => pos.includes(p.toLowerCase()))) {
        group.players.push(rp)
        assigned = true
        break
      }
    }
    if (!assigned) {
      const midfielders = groups[2]
      if (midfielders) {
        midfielders.players.push(rp) // Default to midfielders
      }
    }
  }

  return groups
})

const formationByPosition = computed(() => {
  // Only show starters on the pitch
  const startersOnly = starters.value
  
  const groups = {
    goalkeepers: [] as RosterPlayer[],
    defenders: [] as RosterPlayer[],
    midfielders: [] as RosterPlayer[],
    forwards: [] as RosterPlayer[],
  }

  for (const rp of startersOnly) {
    const pos = (rp.position_played || rp.player?.position || '').toLowerCase()
    if (pos.includes('gk') || pos.includes('goalkeeper')) {
      groups.goalkeepers.push(rp)
    } else if (pos.includes('cb') || pos.includes('lb') || pos.includes('rb') || pos.includes('defender') || pos.includes('back')) {
      groups.defenders.push(rp)
    } else if (pos.includes('st') || pos.includes('cf') || pos.includes('lw') || pos.includes('rw') || pos.includes('forward') || pos.includes('striker') || pos.includes('winger')) {
      groups.forwards.push(rp)
    } else {
      groups.midfielders.push(rp)
    }
  }

  return groups
})

// Get the number of players for each row based on formation
function getFormationRow(rowIndex: number): number {
  const formation = formations.find(f => f.value === selectedFormation.value)
  if (!formation) return 0
  
  // formations.rows = [GK, DEF, MID, FWD] - index 0 is GK which is always 1
  // rowIndex 0 = forwards (last in rows array)
  // rowIndex 1 = midfielders
  // rowIndex 2 = defenders  
  // For 4-2-3-1 we have [1, 4, 2, 3, 1] - special handling
  if (selectedFormation.value === '4-2-3-1') {
    switch (rowIndex) {
      case 0: return 1 // 1 striker
      case 1: return 2 // 2 defensive mids (shown in MF row after AM row)
      case 2: return 4 // 4 defenders
      default: return 0
    }
  }
  
  // Normal formations like 4-3-3, 4-4-2 etc.
  const rows = formation.rows // [GK, DEF, MID, FWD]
  switch (rowIndex) {
    case 0: return rows[3] || 0 // Forwards
    case 1: return rows[2] || 0 // Midfielders
    case 2: return rows[1] || 0 // Defenders
    default: return 0
  }
}

// Drag and drop handlers
function startDrag(player: RosterPlayer) {
  draggingPlayer.value = player
}

function endDrag() {
  draggingPlayer.value = null
}

async function handlePitchDrop(event: DragEvent) {
  event.preventDefault()
  if (!draggingPlayer.value) return

  const player = draggingPlayer.value
  
  // If player was a substitute, make them a starter
  if (!player.is_starter) {
    player.is_starter = true
    player._isChanged = true
    
    // Save the change
    await savePlayerStarterStatus(player)
  }
  
  draggingPlayer.value = null
}

async function handleBenchDrop(event: DragEvent) {
  event.preventDefault()
  if (!draggingPlayer.value) return

  const player = draggingPlayer.value
  
  // If player was a starter, make them a substitute
  if (player.is_starter) {
    player.is_starter = false
    player._isChanged = true
    
    // Save the change
    await savePlayerStarterStatus(player)
  }
  
  draggingPlayer.value = null
}

async function savePlayerStarterStatus(player: RosterPlayer) {
  try {
    await $fetch(`/admin/matches/${matchId.value}/players/${player.player_id}`, {
      baseURL: config.public.apiBase,
      method: 'PUT',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: {
        is_starter: player.is_starter,
        position_played: player.position_played,
        minutes_played: player.minutes_played,
        jersey_number: player.jersey_number,
      },
    })
    
    player._isChanged = false
    toast.success(
      player.is_starter ? 'Moved to Starting XI' : 'Moved to Bench',
      `${player.player?.first_name} ${player.player?.last_name}`
    )
  } catch (error) {
    console.error('Failed to update player starter status:', error)
    // Revert the change
    player.is_starter = !player.is_starter
    toast.error('Failed to Update', 'Could not save player status')
  }
}

// Methods
function getInitials(firstName?: string, lastName?: string): string {
  return `${firstName?.charAt(0) || ''}${lastName?.charAt(0) || ''}`.toUpperCase()
}

function formatDate(date?: string): string {
  if (!date) return ''
  return new Date(date).toLocaleDateString('en-US', { 
    weekday: 'short', 
    month: 'short', 
    day: 'numeric', 
    year: 'numeric' 
  })
}

function isPlayerInRoster(playerId: string): boolean {
  return rosterPlayers.value.some(rp => rp.player_id === playerId)
}

function markAsChanged(rp: RosterPlayer) {
  if (!rp._isNew) {
    rp._isChanged = true
  }
}

// API response type for players (flat structure from backend)
interface ApiPlayerInMatch {
  id: string
  player_id: string
  first_name: string
  last_name: string
  position: string
  profile_photo_url?: string
  position_played?: string
  minutes_played?: number
  goals: number
  assists: number
  highlight_count?: number
  is_starter: boolean
  jersey_number?: number
  formation_x?: number
  formation_y?: number
  subbed_in_for?: string
  subbed_in_at?: number
  subbed_out_at?: number
}

// Unified fetch - API returns both match and players in one response
async function fetchMatchAndRoster() {
  try {
    const response = await $fetch<ApiResponse<{ match: Match; players: ApiPlayerInMatch[] }>>(
      `/admin/matches/${matchId.value}`,
      {
        baseURL: config.public.apiBase,
        headers: { Authorization: `Bearer ${authStore.accessToken}` },
      }
    )
    if (response.success && response.data) {
      match.value = response.data.match
      // Transform flat API response to our RosterPlayer structure
      rosterPlayers.value = (response.data.players || []).map(p => ({
        id: p.id,
        match_id: matchId.value,
        player_id: p.player_id,
        position_played: p.position_played,
        minutes_played: p.minutes_played,
        goals: p.goals,
        assists: p.assists,
        is_starter: p.is_starter ?? true,
        jersey_number: p.jersey_number,
        formation_x: p.formation_x,
        formation_y: p.formation_y,
        subbed_in_for: p.subbed_in_for,
        subbed_in_at: p.subbed_in_at,
        subbed_out_at: p.subbed_out_at,
        player: {
          id: p.player_id,
          first_name: p.first_name,
          last_name: p.last_name,
          position: p.position,
          profile_photo_url: p.profile_photo_url,
        },
      }))
    }
  } catch (error) {
    console.error('Failed to fetch match and roster:', error)
    throw error // Re-throw so caller can handle
  }
}

async function fetchSuggestions() {
  try {
    // Fetch recent players (from all matches in this tournament)
    const response = await $fetch<ApiResponse<Player[]>>(`/admin/players?limit=20`, {
      baseURL: config.public.apiBase,
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
    })
    if (response.success && response.data) {
      recentPlayers.value = Array.isArray(response.data) ? response.data : (response.data as any).players || []
    }
  } catch (error) {
    console.error('Failed to fetch suggestions:', error)
  }
}

async function searchPlayers() {
  if (!searchQuery.value.trim()) {
    searchResults.value = []
    return
  }

  searching.value = true
  try {
    const response = await $fetch<ApiResponse<{ players: Player[] }>>(`/admin/players?search=${encodeURIComponent(searchQuery.value)}`, {
      baseURL: config.public.apiBase,
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
    })
    if (response.success && response.data) {
      searchResults.value = Array.isArray(response.data) ? response.data : (response.data as any).players || []
    }
  } catch (error) {
    console.error('Search failed:', error)
  } finally {
    searching.value = false
  }
}

async function quickAddPlayer(player: Player) {
  if (isPlayerInRoster(player.id)) return
  if (addingPlayerId.value) return // Prevent double-adds

  addingPlayerId.value = player.id

  // Add to local roster immediately (optimistic UI)
  const newRosterPlayer: RosterPlayer = {
    id: `temp-${Date.now()}`,
    match_id: matchId.value,
    player_id: player.id,
    position_played: player.position,
    minutes_played: undefined,
    is_starter: true, // Default to starter
    player: player,
    _isNew: true,
  }
  
  rosterPlayers.value.push(newRosterPlayer)

  // Actually save to backend
  try {
    const response = await $fetch<ApiResponse<RosterPlayer>>(`/admin/matches/${matchId.value}/players`, {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: {
        player_id: player.id,
        position_played: player.position,
        is_starter: true,
      },
    })
    
    if (response.success && response.data) {
      // Update with real ID from server
      const idx = rosterPlayers.value.findIndex(rp => rp.id === newRosterPlayer.id)
      if (idx !== -1) {
        rosterPlayers.value[idx] = { ...response.data, player, _isNew: false }
      }
      toast.success('Player Added', `${player.first_name} ${player.last_name} added to roster`)
    }
  } catch (error: any) {
    // Rollback on failure
    rosterPlayers.value = rosterPlayers.value.filter(rp => rp.id !== newRosterPlayer.id)
    toast.error('Failed to Add', error.data?.message || 'Could not add player')
  } finally {
    addingPlayerId.value = null
  }
}

async function removePlayer(rp: RosterPlayer) {
  if (removingPlayerId.value) return // Prevent multiple removals

  removingPlayerId.value = rp.player_id

  // Remove from local roster immediately
  const idx = rosterPlayers.value.findIndex(r => r.id === rp.id)
  if (idx === -1) {
    removingPlayerId.value = null
    return
  }

  const [removed] = rosterPlayers.value.splice(idx, 1)

  // If it was a new unsaved player, we're done
  if (rp._isNew) {
    toast.success('Player Removed', `${rp.player?.first_name} ${rp.player?.last_name} removed`)
    removingPlayerId.value = null
    return
  }

  // Delete from backend
  try {
    await $fetch(`/admin/matches/${matchId.value}/players/${rp.player_id}`, {
      baseURL: config.public.apiBase,
      method: 'DELETE',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
    })
    toast.success('Player Removed', `${rp.player?.first_name} ${rp.player?.last_name} removed`)
  } catch (error: any) {
    // Rollback on failure
    if (removed !== undefined) {
      rosterPlayers.value.splice(idx, 0, removed)
    }
    toast.error('Failed to Remove', error.data?.message || 'Could not remove player')
  } finally {
    removingPlayerId.value = null
  }
}

async function saveAllChanges() {
  saving.value = true
  
  const changedPlayers = rosterPlayers.value.filter(rp => rp._isChanged && !rp._isNew)
  
  if (changedPlayers.length === 0) {
    toast.info('No Changes', 'There are no pending changes to save')
    saving.value = false
    return
  }
  
  try {
    await Promise.all(changedPlayers.map(rp =>
      $fetch(`/admin/matches/${matchId.value}/players/${rp.player_id}`, {
        baseURL: config.public.apiBase,
        method: 'PUT',
        headers: { Authorization: `Bearer ${authStore.accessToken}` },
        body: {
          position_played: rp.position_played,
          minutes_played: rp.minutes_played,
        },
      })
    ))

    // Clear changed flags
    rosterPlayers.value.forEach(rp => {
      rp._isNew = false
      rp._isChanged = false
    })

    toast.success('Saved', `${changedPlayers.length} player${changedPlayers.length > 1 ? 's' : ''} updated`)
  } catch (error: any) {
    toast.error('Save Failed', error.data?.message || 'Some changes could not be saved')
  } finally {
    saving.value = false
  }
}

// Lifecycle
onMounted(async () => {
  try {
    await Promise.all([fetchMatchAndRoster(), fetchSuggestions()])
  } catch (error) {
    console.error('Failed to load roster page:', error)
    loadError.value = 'Failed to load roster. Please try refreshing the page.'
  } finally {
    loading.value = false
  }
})
</script>
