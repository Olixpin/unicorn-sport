<template>
  <div class="min-h-screen bg-neutral-50">
    <!-- Header -->
    <div class="bg-white border-b border-neutral-200 sticky top-0 z-10">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-4">
            <NuxtLink
              :to="`/admin/tournaments/${match?.tournament_id}`"
              class="flex items-center gap-2 text-neutral-600 hover:text-neutral-900 transition-colors"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
              </svg>
              Back to Tournament
            </NuxtLink>
          </div>
          <div class="flex items-center gap-3">
            <span v-if="isFutureMatch" class="px-3 py-1 bg-blue-100 text-blue-700 rounded-full text-sm font-medium flex items-center gap-1">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
              Upcoming
            </span>
            <span :class="getStatusClass(match?.status)" class="px-3 py-1 rounded-full text-sm font-medium">
              {{ match?.status || 'Loading...' }}
            </span>
          </div>
        </div>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex justify-center items-center h-64">
      <div class="animate-spin rounded-full h-12 w-12 border-4 border-primary-500 border-t-transparent"></div>
    </div>

    <!-- Match Content -->
    <div v-else-if="match" class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <!-- Match Header -->
      <div class="bg-white rounded-2xl border border-neutral-200 p-6 mb-6">
        <div class="flex flex-col md:flex-row md:items-start md:justify-between gap-4">
          <div>
            <h1 class="text-2xl font-bold text-neutral-900">{{ match.title }}</h1>
            <div class="flex items-center gap-4 mt-2 text-neutral-600">
              <span class="flex items-center gap-1">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
                {{ formatDate(match.match_date) }}
              </span>
              <span v-if="match.stage" class="flex items-center gap-1">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z" />
                </svg>
                {{ formatStage(match.stage) }}
              </span>
              <span v-if="match.location" class="flex items-center gap-1">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                </svg>
                {{ match.location }}
              </span>
            </div>
            <div v-if="match.home_team && match.away_team" class="mt-3 text-lg font-medium">
              {{ match.home_team }} 
              <span v-if="match.home_score !== null" class="text-primary-600">{{ match.home_score }}</span>
              <span class="text-neutral-400 mx-2">vs</span>
              <span v-if="match.away_score !== null" class="text-primary-600">{{ match.away_score }}</span>
              {{ match.away_team }}
            </div>
          </div>
          <button
            @click="showEditModal = true"
            class="flex items-center gap-2 px-4 py-2 bg-neutral-100 text-neutral-700 rounded-lg hover:bg-neutral-200 transition-colors"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
            </svg>
            Edit Match
          </button>
        </div>
      </div>

      <!-- Full Match Video Section -->
      <div class="bg-white rounded-2xl border border-neutral-200 p-6 mb-6">
        <div class="flex items-center justify-between mb-4">
          <h2 class="text-lg font-semibold text-neutral-900 flex items-center gap-2">
            <svg class="w-5 h-5 text-amber-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
            </svg>
            Full Match Video
            <span class="bg-amber-100 text-amber-700 text-xs px-2 py-0.5 rounded-full">PAID</span>
          </h2>
        </div>

        <!-- Video Exists -->
        <div v-if="matchVideo" class="flex flex-col md:flex-row gap-6">
          <div class="w-full md:w-72 h-40 bg-neutral-100 rounded-xl overflow-hidden relative">
            <img
              v-if="matchVideo.thumbnail_url"
              :src="matchVideo.thumbnail_url"
              alt="Video thumbnail"
              class="w-full h-full object-cover"
            />
            <div v-else class="w-full h-full flex items-center justify-center">
              <svg class="w-12 h-12 text-neutral-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
              </svg>
            </div>
            <div class="absolute inset-0 bg-black/40 flex items-center justify-center">
              <button
                @click="previewVideo"
                class="w-12 h-12 bg-white/90 rounded-full flex items-center justify-center hover:bg-white transition-colors"
              >
                <svg class="w-5 h-5 text-neutral-900" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M8 5v14l11-7z" />
                </svg>
              </button>
            </div>
          </div>
          <div class="flex-1">
            <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
              <div>
                <span class="text-xs text-neutral-500">Duration</span>
                <p class="font-medium">{{ formatDuration(matchVideo.duration_seconds) }}</p>
              </div>
              <div>
                <span class="text-xs text-neutral-500">File Size</span>
                <p class="font-medium">{{ formatFileSize(matchVideo.file_size_bytes) }}</p>
              </div>
              <div>
                <span class="text-xs text-neutral-500">Status</span>
                <p class="font-medium capitalize">{{ matchVideo.status }}</p>
              </div>
              <div>
                <span class="text-xs text-neutral-500">Price</span>
                <p class="font-medium">${{ (matchVideo.price_cents / 100).toFixed(2) }}</p>
              </div>
            </div>
            <div class="flex gap-3 mt-4">
              <button
                @click="showReplaceVideoModal = true"
                class="px-4 py-2 bg-neutral-100 text-neutral-700 rounded-lg hover:bg-neutral-200 transition-colors text-sm"
              >
                Replace Video
              </button>
              <button
                @click="deleteMatchVideo"
                class="px-4 py-2 bg-rose-50 text-rose-600 rounded-lg hover:bg-rose-100 transition-colors text-sm"
              >
                Delete Video
              </button>
            </div>
          </div>
        </div>

        <!-- No Video - Upload -->
        <div v-else class="border-2 border-dashed border-neutral-300 rounded-xl p-8 text-center">
          <div class="flex flex-col items-center">
            <div class="w-16 h-16 bg-amber-50 rounded-2xl flex items-center justify-center mb-4">
              <svg class="w-8 h-8 text-amber-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
              </svg>
            </div>
            
            <!-- Future match warning -->
            <template v-if="isFutureMatch">
              <p class="text-neutral-900 font-medium mb-1">Match hasn't happened yet</p>
              <p class="text-sm text-neutral-500 mb-2">Scheduled for {{ formatDate(match.match_date) }}</p>
              <div class="flex items-center gap-2 px-4 py-2 bg-blue-50 text-blue-700 rounded-lg text-sm">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                Video upload available after match date
              </div>
            </template>
            
            <!-- Normal upload state -->
            <template v-else>
              <p class="text-neutral-900 font-medium mb-1">No video uploaded yet</p>
              <p class="text-sm text-neutral-500 mb-4">Upload the full match recording (max 25GB)</p>
              <button
                @click="showUploadVideoModal = true"
                class="px-6 py-2.5 bg-gradient-to-r from-primary-500 to-emerald-600 text-white rounded-xl font-medium hover:from-primary-600 hover:to-emerald-700 transition-all"
              >
                Upload Full Match Video
              </button>
            </template>
          </div>
        </div>
      </div>

      <!-- Match Players Section -->
      <div class="bg-white rounded-2xl border border-neutral-200 p-6 mb-6">
        <div class="flex items-center justify-between mb-4">
          <h2 class="text-lg font-semibold text-neutral-900 flex items-center gap-2">
            <svg class="w-5 h-5 text-primary-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
            </svg>
            Match Players ({{ players.length }})
          </h2>
          <button
            @click="showAddPlayerModal = true"
            class="flex items-center gap-2 px-4 py-2 bg-primary-500 text-white rounded-lg hover:bg-primary-600 transition-colors text-sm"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
            </svg>
            Add Player
          </button>
        </div>

        <!-- Players List -->
        <div v-if="players.length > 0" class="space-y-3">
          <div
            v-for="player in players"
            :key="player.player_id"
            class="flex items-center justify-between p-4 bg-neutral-50 rounded-xl hover:bg-neutral-100 transition-colors"
          >
            <div class="flex items-center gap-4">
              <div class="w-12 h-12 rounded-full bg-gradient-to-br from-primary-500 to-emerald-600 flex items-center justify-center text-white font-semibold overflow-hidden">
                <img
                  v-if="player.profile_photo_url"
                  :src="player.profile_photo_url"
                  :alt="player.first_name"
                  class="w-full h-full object-cover"
                />
                <span v-else>{{ getInitials(player.first_name, player.last_name) }}</span>
              </div>
              <div>
                <p class="font-medium text-neutral-900">{{ player.first_name }} {{ player.last_name }}</p>
                <p class="text-sm text-neutral-500">
                  {{ player.position_played || player.position }}
                  <span v-if="player.minutes_played" class="ml-2">• {{ player.minutes_played }} min</span>
                </p>
              </div>
            </div>

            <div class="flex items-center gap-6">
              <!-- Stats -->
              <div class="flex items-center gap-4 text-sm">
                <span v-if="player.goals > 0" class="flex items-center gap-1 text-amber-600 font-medium">
                  ⚽ {{ player.goals }}
                </span>
                <span v-if="player.assists > 0" class="flex items-center gap-1 text-blue-600 font-medium">
                  🎯 {{ player.assists }}
                </span>
                <span class="flex items-center gap-1 text-neutral-600">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
                  </svg>
                  {{ player.highlight_count }} highlights
                </span>
              </div>

              <!-- Actions -->
              <div class="flex items-center gap-2">
                <button
                  v-if="!isFutureMatch"
                  @click="openHighlightUpload(player)"
                  class="p-2 bg-emerald-50 text-emerald-600 rounded-lg hover:bg-emerald-100 transition-colors"
                  title="Add Highlight"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                  </svg>
                </button>
                <span
                  v-else
                  class="p-2 bg-neutral-100 text-neutral-400 rounded-lg cursor-not-allowed"
                  title="Available after match date"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                  </svg>
                </span>
                <button
                  @click="viewPlayerHighlights(player)"
                  class="p-2 bg-blue-50 text-blue-600 rounded-lg hover:bg-blue-100 transition-colors"
                  title="View Highlights"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                  </svg>
                </button>
                <button
                  @click="removePlayer(player)"
                  class="p-2 bg-rose-50 text-rose-600 rounded-lg hover:bg-rose-100 transition-colors"
                  title="Remove from Match"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- Empty State -->
        <div v-else class="text-center py-8 text-neutral-500">
          <p>No players added to this match yet</p>
          <button
            @click="showAddPlayerModal = true"
            class="mt-2 text-primary-600 hover:text-primary-700 font-medium"
          >
            Add the first player
          </button>
        </div>
      </div>

      <!-- Match Highlights Section -->
      <div class="bg-white rounded-2xl border border-neutral-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h2 class="text-lg font-semibold text-neutral-900 flex items-center gap-2">
            <svg class="w-5 h-5 text-emerald-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
            </svg>
            Player Highlights
            <span class="bg-emerald-100 text-emerald-700 text-xs px-2 py-0.5 rounded-full">FREE</span>
          </h2>
          <div class="flex items-center gap-3">
            <select
              v-model="highlightTypeFilter"
              class="px-3 py-2 border border-neutral-200 rounded-lg text-sm bg-white"
            >
              <option value="">All Types</option>
              <option v-for="type in highlightTypes" :key="type.value" :value="type.value">
                {{ type.icon }} {{ type.label }}
              </option>
            </select>
          </div>
        </div>

        <!-- Highlights Grid -->
        <div v-if="filteredHighlights.length > 0" class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
          <div
            v-for="highlight in filteredHighlights"
            :key="highlight.id"
            class="group relative bg-neutral-100 rounded-xl overflow-hidden aspect-video cursor-pointer"
            @click="viewHighlight(highlight)"
          >
            <img
              v-if="highlight.thumbnail_url"
              :src="highlight.thumbnail_url"
              alt="Highlight thumbnail"
              class="w-full h-full object-cover"
            />
            <div class="absolute inset-0 bg-gradient-to-t from-black/70 to-transparent flex flex-col justify-end p-3">
              <div class="flex items-center justify-between">
                <span class="text-white text-xs font-medium bg-white/20 px-2 py-0.5 rounded">
                  {{ getHighlightIcon(highlight.highlight_type) }} {{ highlight.highlight_type }}
                </span>
                <span class="text-white text-xs">{{ formatDuration(highlight.duration_seconds) }}</span>
              </div>
              <p class="text-white text-sm font-medium truncate mt-1">
                {{ highlight.player?.first_name }} {{ highlight.player?.last_name }}
              </p>
            </div>
            <!-- Delete button on hover -->
            <button
              @click.stop="deleteHighlight(highlight)"
              class="absolute top-2 right-2 p-1.5 bg-rose-500 text-white rounded-lg opacity-0 group-hover:opacity-100 transition-opacity"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
        </div>

        <!-- Empty State -->
        <div v-else class="text-center py-12 text-neutral-500">
          <div class="w-16 h-16 bg-neutral-100 rounded-2xl flex items-center justify-center mx-auto mb-4">
            <svg class="w-8 h-8 text-neutral-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
            </svg>
          </div>
          <p class="mb-2">No highlights added yet</p>
          <p class="text-sm">Add players to this match, then upload their highlight clips</p>
        </div>
      </div>
    </div>

    <!-- Add Player Modal -->
    <Teleport to="body">
      <div v-if="showAddPlayerModal" class="fixed inset-0 z-50 overflow-y-auto">
        <div class="fixed inset-0 bg-black/50" @click="showAddPlayerModal = false"></div>
        <div class="relative min-h-screen flex items-center justify-center p-4">
          <div class="relative bg-white rounded-2xl shadow-xl w-full max-w-lg p-6">
            <h3 class="text-xl font-semibold mb-4">Add Player to Match</h3>

            <!-- Search -->
            <div class="mb-4">
              <input
                v-model="playerSearch"
                type="text"
                placeholder="Search players by name..."
                class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                @input="searchPlayers"
              />
            </div>

            <!-- Search Results -->
            <div v-if="playerSearchResults.length > 0" class="max-h-60 overflow-y-auto mb-4 space-y-2">
              <button
                v-for="p in playerSearchResults"
                :key="p.id"
                @click="selectPlayer(p)"
                class="w-full flex items-center gap-3 p-3 bg-neutral-50 rounded-lg hover:bg-neutral-100 transition-colors text-left"
              >
                <div class="w-10 h-10 bg-gradient-to-br from-primary-500 to-emerald-600 rounded-full flex items-center justify-center text-white text-sm font-medium">
                  {{ getInitials(p.first_name, p.last_name) }}
                </div>
                <div>
                  <p class="font-medium">{{ p.first_name }} {{ p.last_name }}</p>
                  <p class="text-xs text-neutral-500">{{ p.position }} • {{ p.academy?.name || 'No academy' }}</p>
                </div>
              </button>
            </div>

            <!-- Selected Player Form -->
            <div v-if="selectedPlayer" class="border-t border-neutral-200 pt-4 mt-4">
              <div class="flex items-center gap-3 mb-4">
                <div class="w-10 h-10 bg-gradient-to-br from-primary-500 to-emerald-600 rounded-full flex items-center justify-center text-white text-sm font-medium">
                  {{ getInitials(selectedPlayer.first_name, selectedPlayer.last_name) }}
                </div>
                <div>
                  <p class="font-medium">{{ selectedPlayer.first_name }} {{ selectedPlayer.last_name }}</p>
                  <p class="text-xs text-neutral-500">{{ selectedPlayer.position }}</p>
                </div>
              </div>

              <div class="grid grid-cols-2 gap-4">
                <div>
                  <label class="block text-sm font-medium text-neutral-700 mb-1">Position Played</label>
                  <input
                    v-model="addPlayerForm.position_played"
                    type="text"
                    class="w-full px-3 py-2 border border-neutral-200 rounded-lg"
                    :placeholder="selectedPlayer.position"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-neutral-700 mb-1">Minutes Played</label>
                  <input
                    v-model="addPlayerForm.minutes_played"
                    type="number"
                    class="w-full px-3 py-2 border border-neutral-200 rounded-lg"
                    placeholder="90"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-neutral-700 mb-1">Goals</label>
                  <input
                    v-model="addPlayerForm.goals"
                    type="number"
                    class="w-full px-3 py-2 border border-neutral-200 rounded-lg"
                    placeholder="0"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-neutral-700 mb-1">Assists</label>
                  <input
                    v-model="addPlayerForm.assists"
                    type="number"
                    class="w-full px-3 py-2 border border-neutral-200 rounded-lg"
                    placeholder="0"
                  />
                </div>
              </div>
            </div>

            <div class="flex justify-end gap-3 mt-6">
              <button
                @click="showAddPlayerModal = false"
                class="px-4 py-2 border border-neutral-200 rounded-lg hover:bg-neutral-50 transition-colors"
              >
                Cancel
              </button>
              <button
                @click="addPlayerToMatch"
                :disabled="!selectedPlayer || addingPlayer"
                class="px-4 py-2 bg-primary-500 text-white rounded-lg hover:bg-primary-600 transition-colors disabled:opacity-50"
              >
                {{ addingPlayer ? 'Adding...' : 'Add Player' }}
              </button>
            </div>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- Upload Highlight Modal -->
    <Teleport to="body">
      <div v-if="showHighlightUploadModal" class="fixed inset-0 z-50 overflow-y-auto">
        <div class="fixed inset-0 bg-black/50" @click="showHighlightUploadModal = false"></div>
        <div class="relative min-h-screen flex items-center justify-center p-4">
          <div class="relative bg-white rounded-2xl shadow-xl w-full max-w-lg p-6">
            <h3 class="text-xl font-semibold mb-2">Add Highlight</h3>
            <p class="text-neutral-500 mb-4">
              Upload a highlight clip for {{ selectedPlayerForHighlight?.first_name }} {{ selectedPlayerForHighlight?.last_name }}
            </p>

            <!-- Highlight Type -->
            <div class="mb-4">
              <label class="block text-sm font-medium text-neutral-700 mb-2">Highlight Type *</label>
              <div class="grid grid-cols-3 gap-2">
                <button
                  v-for="type in highlightTypes"
                  :key="type.value"
                  @click="highlightForm.highlight_type = type.value"
                  :class="[
                    'p-2 border rounded-lg text-sm text-center transition-colors',
                    highlightForm.highlight_type === type.value
                      ? 'border-primary-500 bg-primary-50 text-primary-700'
                      : 'border-neutral-200 hover:bg-neutral-50'
                  ]"
                >
                  <span class="text-lg block mb-0.5">{{ type.icon }}</span>
                  {{ type.label }}
                </button>
              </div>
            </div>

            <!-- File Upload -->
            <div class="mb-4">
              <label class="block text-sm font-medium text-neutral-700 mb-2">Video File *</label>
              <div
                @click="highlightInput?.click()"
                class="border-2 border-dashed border-neutral-300 rounded-xl p-6 text-center cursor-pointer hover:border-neutral-400 transition-colors"
              >
                <input
                  ref="highlightInput"
                  type="file"
                  accept="video/*"
                  class="hidden"
                  @change="handleHighlightFileSelect"
                />
                <div v-if="!highlightFile">
                  <svg class="w-8 h-8 text-neutral-400 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
                  </svg>
                  <p class="text-sm text-neutral-500">Click to upload video (max 1GB)</p>
                </div>
                <div v-else class="flex items-center justify-center gap-2">
                  <svg class="w-5 h-5 text-emerald-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                  </svg>
                  <span class="text-sm font-medium">{{ highlightFile.name }}</span>
                  <span class="text-xs text-neutral-500">({{ formatFileSize(highlightFile.size) }})</span>
                </div>
              </div>
            </div>

            <!-- Upload Progress -->
            <div v-if="highlightUploadProgress !== null" class="mb-4">
              <div class="flex items-center justify-between text-sm mb-1">
                <span>Uploading...</span>
                <span>{{ highlightUploadProgress }}%</span>
              </div>
              <div class="h-2 bg-neutral-200 rounded-full overflow-hidden">
                <div
                  class="h-full bg-gradient-to-r from-primary-500 to-emerald-600 transition-all"
                  :style="{ width: `${highlightUploadProgress}%` }"
                ></div>
              </div>
            </div>

            <!-- Optional Fields -->
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-1">Title (optional)</label>
                <input
                  v-model="highlightForm.title"
                  type="text"
                  class="w-full px-3 py-2 border border-neutral-200 rounded-lg"
                  placeholder="e.g., Amazing solo goal"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-1">Timestamp in Match (minutes)</label>
                <input
                  v-model="highlightForm.timestamp_minutes"
                  type="number"
                  class="w-full px-3 py-2 border border-neutral-200 rounded-lg"
                  placeholder="e.g., 45 for 45th minute"
                />
              </div>
            </div>

            <div class="flex justify-end gap-3 mt-6">
              <button
                @click="closeHighlightUpload"
                class="px-4 py-2 border border-neutral-200 rounded-lg hover:bg-neutral-50 transition-colors"
              >
                Cancel
              </button>
              <button
                @click="uploadHighlight"
                :disabled="!canUploadHighlight || uploadingHighlight"
                class="px-4 py-2 bg-primary-500 text-white rounded-lg hover:bg-primary-600 transition-colors disabled:opacity-50"
              >
                {{ uploadingHighlight ? 'Uploading...' : 'Upload Highlight' }}
              </button>
            </div>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- Upload Video Modal (for full match) -->
    <Teleport to="body">
      <div v-if="showUploadVideoModal" class="fixed inset-0 z-50 overflow-y-auto">
        <div class="fixed inset-0 bg-black/50" @click="showUploadVideoModal = false"></div>
        <div class="relative min-h-screen flex items-center justify-center p-4">
          <div class="relative bg-white rounded-2xl shadow-xl w-full max-w-lg p-6">
            <h3 class="text-xl font-semibold mb-4">Upload Full Match Video</h3>

            <div
              @click="videoInput?.click()"
              @drop.prevent="handleVideoFileDrop"
              @dragover.prevent
              class="border-2 border-dashed border-neutral-300 rounded-xl p-8 text-center cursor-pointer hover:border-neutral-400 transition-colors"
            >
              <input
                ref="videoInput"
                type="file"
                accept="video/*"
                class="hidden"
                @change="handleVideoFileSelect"
              />
              <div v-if="!matchVideoFile">
                <svg class="w-12 h-12 text-neutral-400 mx-auto mb-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
                </svg>
                <p class="text-neutral-900 font-medium mb-1">Drag & drop or click to upload</p>
                <p class="text-sm text-neutral-500">MP4, MOV, WebM (max 25GB)</p>
              </div>
              <div v-else class="flex flex-col items-center">
                <svg class="w-10 h-10 text-emerald-500 mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                </svg>
                <p class="font-medium">{{ matchVideoFile.name }}</p>
                <p class="text-sm text-neutral-500">{{ formatFileSize(matchVideoFile.size) }}</p>
              </div>
            </div>

            <!-- Upload Progress -->
            <div v-if="matchVideoUploadProgress !== null" class="mt-4">
              <div class="flex items-center justify-between text-sm mb-1">
                <span>Uploading...</span>
                <span>{{ matchVideoUploadProgress }}%</span>
              </div>
              <div class="h-3 bg-neutral-200 rounded-full overflow-hidden">
                <div
                  class="h-full bg-gradient-to-r from-primary-500 to-emerald-600 transition-all"
                  :style="{ width: `${matchVideoUploadProgress}%` }"
                ></div>
              </div>
            </div>

            <div class="flex justify-end gap-3 mt-6">
              <button
                @click="showUploadVideoModal = false"
                class="px-4 py-2 border border-neutral-200 rounded-lg hover:bg-neutral-50 transition-colors"
              >
                Cancel
              </button>
              <button
                @click="uploadMatchVideo"
                :disabled="!matchVideoFile || uploadingMatchVideo"
                class="px-4 py-2 bg-primary-500 text-white rounded-lg hover:bg-primary-600 transition-colors disabled:opacity-50"
              >
                {{ uploadingMatchVideo ? 'Uploading...' : 'Upload Video' }}
              </button>
            </div>
          </div>
        </div>
      </div>
    </Teleport>
  </div>
</template>

<script setup lang="ts">
import type { ApiResponse } from '~/types/index'

interface Match {
  id: string
  title: string
  description?: string
  match_date: string
  location?: string
  stage?: string
  home_team?: string
  away_team?: string
  home_score?: number
  away_score?: number
  status: string
  tournament_id?: string
  video?: MatchVideo
}

interface MatchVideo {
  id: string
  match_id: string
  video_url: string
  thumbnail_url?: string
  duration_seconds?: number
  file_size_bytes?: number
  status: string
  price_cents: number
}

interface MatchPlayer {
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
  highlight_count: number
}

interface PlayerHighlight {
  id: string
  player_id: string
  highlight_type: string
  title?: string
  thumbnail_url?: string
  duration_seconds?: number
  view_count: number
  player?: {
    first_name: string
    last_name: string
  }
}

interface Player {
  id: string
  first_name: string
  last_name: string
  position: string
  academy?: { name: string }
}

definePageMeta({
  layout: 'admin',
  middleware: 'admin',
})

const route = useRoute()
const router = useRouter()
const config = useRuntimeConfig()
const authStore = useAuthStore()
const toast = useToast()

const matchId = route.params.matchId as string

// State
const loading = ref(true)
const match = ref<Match | null>(null)
const matchVideo = ref<MatchVideo | null>(null)
const players = ref<MatchPlayer[]>([])
const highlights = ref<PlayerHighlight[]>([])
const highlightTypeFilter = ref('')

// Modals
const showEditModal = ref(false)
const showAddPlayerModal = ref(false)
const showHighlightUploadModal = ref(false)
const showUploadVideoModal = ref(false)
const showReplaceVideoModal = ref(false)

// Player search
const playerSearch = ref('')
const playerSearchResults = ref<Player[]>([])
const selectedPlayer = ref<Player | null>(null)
const addPlayerForm = reactive({
  position_played: '',
  minutes_played: null as number | null,
  goals: 0,
  assists: 0,
})
const addingPlayer = ref(false)

// Highlight upload
const selectedPlayerForHighlight = ref<MatchPlayer | null>(null)
const highlightFile = ref<File | null>(null)
const highlightForm = reactive({
  highlight_type: 'goal',
  title: '',
  timestamp_minutes: null as number | null,
})
const highlightUploadProgress = ref<number | null>(null)
const uploadingHighlight = ref(false)
const highlightInput = ref<HTMLInputElement | null>(null)

// Match video upload
const matchVideoFile = ref<File | null>(null)
const matchVideoUploadProgress = ref<number | null>(null)
const uploadingMatchVideo = ref(false)
const videoInput = ref<HTMLInputElement | null>(null)

// Highlight types
const highlightTypes = [
  { value: 'goal', label: 'Goal', icon: '⚽' },
  { value: 'assist', label: 'Assist', icon: '🎯' },
  { value: 'dribbling', label: 'Dribbling', icon: '🦵' },
  { value: 'defending', label: 'Defending', icon: '🛡️' },
  { value: 'tackling', label: 'Tackling', icon: '🦶' },
  { value: 'passing', label: 'Passing', icon: '➡️' },
  { value: 'shooting', label: 'Shooting', icon: '🎯' },
  { value: 'heading', label: 'Heading', icon: '🗣️' },
  { value: 'speed', label: 'Speed', icon: '⚡' },
  { value: 'save', label: 'Save', icon: '🧤' },
]

const filteredHighlights = computed(() => {
  if (!highlightTypeFilter.value) return highlights.value
  return highlights.value.filter(h => h.highlight_type === highlightTypeFilter.value)
})

const canUploadHighlight = computed(() => {
  return highlightFile.value && highlightForm.highlight_type
})

// Check if match is in the future - restrict video/highlight uploads
const isFutureMatch = computed(() => {
  if (!match.value?.match_date) return false
  const matchDate = new Date(match.value.match_date)
  const now = new Date()
  // Consider match as "past" if it's today or earlier
  matchDate.setHours(23, 59, 59, 999)
  return matchDate > now
})

// Fetch data
async function fetchMatch() {
  try {
    const response = await $fetch<ApiResponse<{ match: Match; players: MatchPlayer[] }>>(
      `/admin/matches/${matchId}`,
      {
        baseURL: config.public.apiBase,
        headers: { Authorization: `Bearer ${authStore.accessToken}` },
      }
    )
    if (response.success && response.data) {
      match.value = response.data.match
      matchVideo.value = response.data.match.video || null
      players.value = response.data.players || []
    }
  } catch (error) {
    console.error('Failed to fetch match:', error)
    toast.error('Error', 'Failed to load match')
  }
}

async function fetchHighlights() {
  try {
    const response = await $fetch<ApiResponse<PlayerHighlight[]>>(
      `/admin/matches/${matchId}/highlights`,
      {
        baseURL: config.public.apiBase,
        headers: { Authorization: `Bearer ${authStore.accessToken}` },
      }
    )
    if (response.success && response.data) {
      // Flatten the grouped response
      const allHighlights: PlayerHighlight[] = []
      for (const group of response.data) {
        if (Array.isArray((group as any).highlights)) {
          for (const h of (group as any).highlights) {
            allHighlights.push({
              ...h,
              player: { first_name: (group as any).player_name?.split(' ')[0] || '', last_name: (group as any).player_name?.split(' ')[1] || '' }
            })
          }
        }
      }
      highlights.value = allHighlights
    }
  } catch (error) {
    console.error('Failed to fetch highlights:', error)
  }
}

async function searchPlayers() {
  if (playerSearch.value.length < 2) {
    playerSearchResults.value = []
    return
  }
  try {
    const response = await $fetch<ApiResponse<{ players: Player[] }>>(
      `/admin/players?q=${encodeURIComponent(playerSearch.value)}&limit=10`,
      {
        baseURL: config.public.apiBase,
        headers: { Authorization: `Bearer ${authStore.accessToken}` },
      }
    )
    if (response.success && response.data) {
      // Filter out players already in match
      const existingIds = new Set(players.value.map(p => p.player_id))
      playerSearchResults.value = (response.data.players || []).filter(p => !existingIds.has(p.id))
    }
  } catch (error) {
    console.error('Failed to search players:', error)
  }
}

function selectPlayer(player: Player) {
  selectedPlayer.value = player
  addPlayerForm.position_played = player.position
}

async function addPlayerToMatch() {
  if (!selectedPlayer.value) return
  addingPlayer.value = true
  try {
    await $fetch(`/admin/matches/${matchId}/players`, {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: {
        player_id: selectedPlayer.value.id,
        position_played: addPlayerForm.position_played,
        minutes_played: addPlayerForm.minutes_played,
        goals: addPlayerForm.goals || 0,
        assists: addPlayerForm.assists || 0,
      },
    })
    toast.success('Player Added', 'Player has been added to the match')
    showAddPlayerModal.value = false
    selectedPlayer.value = null
    playerSearch.value = ''
    playerSearchResults.value = []
    Object.assign(addPlayerForm, { position_played: '', minutes_played: null, goals: 0, assists: 0 })
    fetchMatch()
  } catch (error: any) {
    toast.error('Error', error.data?.message || 'Failed to add player')
  } finally {
    addingPlayer.value = false
  }
}

async function removePlayer(player: MatchPlayer) {
  if (!confirm(`Remove ${player.first_name} ${player.last_name} from this match?`)) return
  try {
    await $fetch(`/admin/matches/${matchId}/players/${player.player_id}`, {
      baseURL: config.public.apiBase,
      method: 'DELETE',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
    })
    toast.success('Player Removed', 'Player has been removed from the match')
    fetchMatch()
  } catch (error) {
    toast.error('Error', 'Failed to remove player')
  }
}

function openHighlightUpload(player: MatchPlayer) {
  selectedPlayerForHighlight.value = player
  showHighlightUploadModal.value = true
}

function closeHighlightUpload() {
  showHighlightUploadModal.value = false
  selectedPlayerForHighlight.value = null
  highlightFile.value = null
  highlightForm.highlight_type = 'goal'
  highlightForm.title = ''
  highlightForm.timestamp_minutes = null
  highlightUploadProgress.value = null
}

function handleHighlightFileSelect(e: Event) {
  const input = e.target as HTMLInputElement
  if (input.files && input.files[0]) {
    highlightFile.value = input.files[0]
  }
}

async function uploadHighlight() {
  if (!highlightFile.value || !selectedPlayerForHighlight.value) return
  uploadingHighlight.value = true
  
  try {
    // 1. Init upload
    const initResponse = await $fetch<ApiResponse<{ upload_url: string; s3_key: string }>>(`/admin/highlights/upload/init`, {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: {
        player_id: selectedPlayerForHighlight.value.player_id,
        match_id: matchId,
        file_name: highlightFile.value.name,
        file_size: highlightFile.value.size,
        content_type: highlightFile.value.type,
      },
    })

    if (!initResponse.success || !initResponse.data) {
      throw new Error('Failed to initialize upload')
    }

    // 2. Upload to S3
    await new Promise<void>((resolve, reject) => {
      const xhr = new XMLHttpRequest()
      xhr.upload.onprogress = (e) => {
        if (e.lengthComputable) {
          highlightUploadProgress.value = Math.round((e.loaded / e.total) * 100)
        }
      }
      xhr.onload = () => {
        if (xhr.status >= 200 && xhr.status < 300) {
          resolve()
        } else {
          reject(new Error('Upload failed'))
        }
      }
      xhr.onerror = () => reject(new Error('Upload failed'))
      xhr.open('PUT', initResponse.data!.upload_url)
      xhr.setRequestHeader('Content-Type', highlightFile.value!.type)
      xhr.send(highlightFile.value)
    })

    // 3. Create highlight record
    await $fetch(`/admin/highlights`, {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: {
        player_id: selectedPlayerForHighlight.value.player_id,
        match_id: matchId,
        highlight_type: highlightForm.highlight_type,
        s3_key: initResponse.data.s3_key,
        title: highlightForm.title || undefined,
        timestamp_in_match: highlightForm.timestamp_minutes ? highlightForm.timestamp_minutes * 60 : undefined,
        file_size_bytes: highlightFile.value.size,
      },
    })

    toast.success('Highlight Uploaded', 'Highlight has been added successfully')
    closeHighlightUpload()
    fetchMatch()
    fetchHighlights()
  } catch (error: any) {
    toast.error('Upload Failed', error.message || 'Failed to upload highlight')
  } finally {
    uploadingHighlight.value = false
    highlightUploadProgress.value = null
  }
}

function handleVideoFileSelect(e: Event) {
  const input = e.target as HTMLInputElement
  if (input.files && input.files[0]) {
    matchVideoFile.value = input.files[0]
  }
}

function handleVideoFileDrop(e: DragEvent) {
  if (e.dataTransfer?.files && e.dataTransfer.files[0]) {
    matchVideoFile.value = e.dataTransfer.files[0]
  }
}

async function uploadMatchVideo() {
  if (!matchVideoFile.value) return
  uploadingMatchVideo.value = true

  try {
    // 1. Init upload
    const initResponse = await $fetch<ApiResponse<{ session_id: string; upload_url?: string; s3_key: string; upload_method: string }>>(`/admin/matches/${matchId}/video/upload`, {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: {
        file_name: matchVideoFile.value.name,
        file_size: matchVideoFile.value.size,
        content_type: matchVideoFile.value.type,
      },
    })

    if (!initResponse.success || !initResponse.data) {
      throw new Error('Failed to initialize upload')
    }

    // 2. Upload to S3 (direct for now, multipart not implemented in UI)
    if (initResponse.data.upload_url) {
      await new Promise<void>((resolve, reject) => {
        const xhr = new XMLHttpRequest()
        xhr.upload.onprogress = (e) => {
          if (e.lengthComputable) {
            matchVideoUploadProgress.value = Math.round((e.loaded / e.total) * 100)
          }
        }
        xhr.onload = () => {
          if (xhr.status >= 200 && xhr.status < 300) {
            resolve()
          } else {
            reject(new Error('Upload failed'))
          }
        }
        xhr.onerror = () => reject(new Error('Upload failed'))
        xhr.open('PUT', initResponse.data!.upload_url!)
        xhr.setRequestHeader('Content-Type', matchVideoFile.value!.type)
        xhr.send(matchVideoFile.value)
      })
    }

    // 3. Save video record
    await $fetch(`/admin/matches/${matchId}/video`, {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: {
        s3_key: initResponse.data.s3_key,
        file_size_bytes: matchVideoFile.value.size,
      },
    })

    toast.success('Video Uploaded', 'Full match video has been uploaded successfully')
    showUploadVideoModal.value = false
    matchVideoFile.value = null
    matchVideoUploadProgress.value = null
    fetchMatch()
  } catch (error: any) {
    toast.error('Upload Failed', error.message || 'Failed to upload video')
  } finally {
    uploadingMatchVideo.value = false
  }
}

async function deleteMatchVideo() {
  if (!confirm('Are you sure you want to delete this video?')) return
  try {
    await $fetch(`/admin/matches/${matchId}/video`, {
      baseURL: config.public.apiBase,
      method: 'DELETE',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
    })
    toast.success('Video Deleted', 'Match video has been deleted')
    matchVideo.value = null
    fetchMatch()
  } catch (error) {
    toast.error('Error', 'Failed to delete video')
  }
}

async function deleteHighlight(highlight: PlayerHighlight) {
  if (!confirm('Are you sure you want to delete this highlight?')) return
  try {
    await $fetch(`/admin/highlights/${highlight.id}`, {
      baseURL: config.public.apiBase,
      method: 'DELETE',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
    })
    toast.success('Highlight Deleted', 'Highlight has been deleted')
    fetchHighlights()
    fetchMatch()
  } catch (error) {
    toast.error('Error', 'Failed to delete highlight')
  }
}

function viewPlayerHighlights(player: MatchPlayer) {
  highlightTypeFilter.value = ''
  // Scroll to highlights section
  document.querySelector('[data-highlights]')?.scrollIntoView({ behavior: 'smooth' })
}

function viewHighlight(highlight: PlayerHighlight) {
  // Could open a modal to preview
  console.log('View highlight:', highlight)
}

function previewVideo() {
  // Could open video player modal
  console.log('Preview video')
}

// Helpers
function getInitials(firstName: string, lastName: string): string {
  return `${firstName?.[0] || ''}${lastName?.[0] || ''}`.toUpperCase()
}

function formatDate(date: string): string {
  return new Date(date).toLocaleDateString('en-US', { year: 'numeric', month: 'short', day: 'numeric' })
}

function formatStage(stage: string): string {
  const stages: Record<string, string> = {
    group: 'Group Stage',
    round_of_16: 'Round of 16',
    quarterfinal: 'Quarter Final',
    semifinal: 'Semi Final',
    final: 'Final',
  }
  return stages[stage] || stage
}

function formatDuration(seconds?: number): string {
  if (!seconds) return '0:00'
  const hrs = Math.floor(seconds / 3600)
  const mins = Math.floor((seconds % 3600) / 60)
  const secs = seconds % 60
  if (hrs > 0) {
    return `${hrs}:${mins.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}`
  }
  return `${mins}:${secs.toString().padStart(2, '0')}`
}

function formatFileSize(bytes?: number): string {
  if (!bytes) return '0 B'
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(1024))
  return `${(bytes / Math.pow(1024, i)).toFixed(1)} ${sizes[i]}`
}

function getStatusClass(status?: string): string {
  const classes: Record<string, string> = {
    scheduled: 'bg-blue-100 text-blue-700',
    in_progress: 'bg-amber-100 text-amber-700',
    completed: 'bg-emerald-100 text-emerald-700',
    cancelled: 'bg-neutral-100 text-neutral-700',
  }
  const key = status || 'scheduled'
  return classes[key] || classes.scheduled!
}

function getHighlightIcon(type: string): string {
  return highlightTypes.find(t => t.value === type)?.icon || '✨'
}

// Initialize
onMounted(async () => {
  loading.value = true
  await Promise.all([fetchMatch(), fetchHighlights()])
  loading.value = false
})
</script>
