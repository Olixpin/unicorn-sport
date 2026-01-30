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
      <div class="bg-white rounded-2xl border border-neutral-200 p-4 sm:p-6 mb-6">
        <div class="flex flex-col gap-4">
          <div class="flex items-start justify-between gap-3">
            <h1 class="text-xl sm:text-2xl font-bold text-neutral-900">{{ match.title }}</h1>
            <button
              @click="showEditModal = true"
              class="flex-shrink-0 w-9 h-9 sm:w-auto sm:h-auto sm:px-4 sm:py-2 flex items-center justify-center gap-2 bg-neutral-100 text-neutral-700 rounded-lg hover:bg-neutral-200 transition-colors"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
              </svg>
              <span class="hidden sm:inline">Edit Match</span>
            </button>
          </div>
          
          <!-- Metadata - compact on mobile -->
          <div class="flex flex-wrap items-center gap-x-3 gap-y-1.5 text-sm text-neutral-600">
            <span class="flex items-center gap-1.5">
              <svg class="w-3.5 h-3.5 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
              </svg>
              {{ formatDate(match.match_date) }}
            </span>
            <span v-if="match.stage" class="flex items-center gap-1.5">
              <span class="text-neutral-300">•</span>
              {{ formatStage(match.stage) }}
            </span>
            <span v-if="match.location" class="flex items-center gap-1.5">
              <span class="text-neutral-300">•</span>
              {{ match.location }}
            </span>
          </div>
          
          <!-- Score -->
          <div v-if="match.home_team && match.away_team" class="text-base sm:text-lg font-semibold">
            {{ match.home_team }} 
            <span v-if="match.home_score !== null" class="text-primary-600">{{ match.home_score }}</span>
            <span class="text-neutral-400 mx-1.5">vs</span>
            <span v-if="match.away_score !== null" class="text-primary-600">{{ match.away_score }}</span>
            {{ match.away_team }}
          </div>
        </div>
      </div>

      <!-- Full Match Video Section -->
      <div class="bg-white rounded-2xl border border-neutral-200 p-4 sm:p-6 mb-6">
        <div class="flex items-center gap-2 mb-4">
          <svg class="w-5 h-5 text-amber-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
          </svg>
          <h2 class="text-base sm:text-lg font-semibold text-neutral-900">Full Match Video</h2>
          <span class="bg-amber-100 text-amber-700 text-[10px] sm:text-xs px-2 py-0.5 rounded-full font-medium">PAID</span>
        </div>

        <!-- Video Exists -->
        <div v-if="matchVideo" class="space-y-4">
          <!-- Video Thumbnail -->
          <div class="w-full aspect-video sm:aspect-auto sm:h-44 bg-neutral-100 rounded-xl overflow-hidden relative group">
            <img
              v-if="matchVideo.thumbnail_url"
              :src="matchVideo.thumbnail_url"
              alt="Video thumbnail"
              class="w-full h-full object-cover"
            />
            <div v-else class="w-full h-full flex items-center justify-center bg-gradient-to-br from-neutral-200 to-neutral-300">
              <svg class="w-10 h-10 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
              </svg>
            </div>
            <div class="absolute inset-0 bg-black/30 flex items-center justify-center">
              <button
                @click="previewVideo"
                class="w-11 h-11 bg-white/95 rounded-full flex items-center justify-center hover:bg-white hover:scale-105 transition-all shadow-lg"
              >
                <svg class="w-5 h-5 text-neutral-900 ml-0.5" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M8 5v14l11-7z" />
                </svg>
              </button>
            </div>
            <!-- Upload thumbnail button overlay -->
            <button
              @click="thumbnailInput?.click()"
              class="absolute bottom-2 right-2 p-1.5 bg-black/60 text-white rounded-lg opacity-0 group-hover:opacity-100 transition-opacity text-xs flex items-center gap-1"
            >
              <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
              </svg>
              {{ matchVideo.thumbnail_url ? 'Change' : 'Add' }}
            </button>
            <input
              ref="thumbnailInput"
              type="file"
              accept="image/*"
              class="hidden"
              @change="handleThumbnailUpload"
            />
          </div>
          
          <!-- Video Stats Grid -->
          <div class="grid grid-cols-2 sm:grid-cols-4 gap-3">
            <div class="bg-neutral-50 rounded-lg p-3">
              <span class="text-[11px] text-neutral-500 uppercase tracking-wide">Duration</span>
              <p class="font-semibold text-neutral-900 mt-0.5">{{ formatDuration(matchVideo.duration_seconds) }}</p>
            </div>
            <div class="bg-neutral-50 rounded-lg p-3">
              <span class="text-[11px] text-neutral-500 uppercase tracking-wide">File Size</span>
              <p class="font-semibold text-neutral-900 mt-0.5">{{ formatFileSize(matchVideo.file_size_bytes) }}</p>
            </div>
            <div class="bg-neutral-50 rounded-lg p-3">
              <span class="text-[11px] text-neutral-500 uppercase tracking-wide">Status</span>
              <p class="font-semibold text-neutral-900 mt-0.5 capitalize">{{ matchVideo.status }}</p>
            </div>
            <div class="bg-neutral-50 rounded-lg p-3">
              <span class="text-[11px] text-neutral-500 uppercase tracking-wide">Price</span>
              <p class="font-semibold text-neutral-900 mt-0.5">${{ (matchVideo.price_cents / 100).toFixed(2) }}</p>
            </div>
          </div>
          
          <!-- Action Buttons -->
          <div class="flex gap-2">
            <button
              @click="showReplaceVideoModal = true"
              class="flex-1 sm:flex-none px-4 py-2.5 bg-neutral-100 text-neutral-700 rounded-xl hover:bg-neutral-200 transition-colors text-sm font-medium"
            >
              Replace Video
            </button>
            <button
              @click="deleteMatchVideo"
              class="flex-1 sm:flex-none px-4 py-2.5 bg-rose-50 text-rose-600 rounded-xl hover:bg-rose-100 transition-colors text-sm font-medium border border-rose-100"
            >
              Delete Video
            </button>
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
      <div class="bg-white rounded-2xl border border-neutral-200 p-4 sm:p-6 mb-6">
        <div class="flex items-center justify-between gap-3 mb-4">
          <h2 class="text-base sm:text-lg font-semibold text-neutral-900 flex items-center gap-2">
            <svg class="w-5 h-5 text-primary-500 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
            </svg>
            <span>Players <span class="text-neutral-500">({{ players.length }})</span></span>
          </h2>
          <button
            @click="navigateToRoster"
            class="flex-shrink-0 w-9 h-9 sm:w-auto sm:h-auto sm:px-4 sm:py-2 flex items-center justify-center gap-2 bg-primary-500 text-white rounded-lg hover:bg-primary-600 transition-colors text-sm"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
            </svg>
            <span class="hidden sm:inline">Manage Roster</span>
          </button>
        </div>

        <!-- Players List -->
        <div v-if="players.length > 0" class="space-y-2">
          <div
            v-for="player in players"
            :key="player.player_id"
            class="p-3 sm:p-4 bg-neutral-50 rounded-xl hover:bg-neutral-100 transition-colors"
          >
            <!-- Mobile: stacked layout -->
            <div class="flex items-center gap-3">
              <div class="w-10 h-10 sm:w-12 sm:h-12 flex-shrink-0 rounded-full bg-gradient-to-br from-primary-500 to-emerald-600 flex items-center justify-center text-white text-sm sm:text-base font-semibold overflow-hidden">
                <img
                  v-if="player.profile_photo_url"
                  :src="player.profile_photo_url"
                  :alt="player.first_name"
                  class="w-full h-full object-cover"
                />
                <span v-else>{{ getInitials(player.first_name, player.last_name) }}</span>
              </div>
              
              <div class="flex-1 min-w-0">
                <p class="font-medium text-neutral-900 text-sm sm:text-base truncate">{{ player.first_name }} {{ player.last_name }}</p>
                <div class="flex items-center gap-2 text-xs sm:text-sm text-neutral-500">
                  <span>{{ player.position_played || player.position }}</span>
                  <span class="text-neutral-300">•</span>
                  <span class="flex items-center gap-1">
                    <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
                    </svg>
                    {{ player.highlight_count }}
                  </span>
                </div>
              </div>

              <!-- Action buttons - always visible -->
              <div class="flex items-center gap-1.5">
                <button
                  v-if="!isFutureMatch"
                  @click="openHighlightUpload(player)"
                  class="w-8 h-8 flex items-center justify-center bg-emerald-50 text-emerald-600 rounded-lg hover:bg-emerald-100 transition-colors"
                  title="Add Highlight"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                  </svg>
                </button>
                <button
                  @click="viewPlayerHighlights(player)"
                  class="w-8 h-8 flex items-center justify-center bg-blue-50 text-blue-600 rounded-lg hover:bg-blue-100 transition-colors"
                  title="View Highlights"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                  </svg>
                </button>
                <button
                  @click="removePlayer(player)"
                  class="w-8 h-8 flex items-center justify-center bg-rose-50 text-rose-600 rounded-lg hover:bg-rose-100 transition-colors"
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
        <div v-else class="text-center py-12">
          <div class="w-16 h-16 mx-auto mb-4 rounded-full bg-neutral-100 flex items-center justify-center">
            <svg class="w-8 h-8 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0z" />
            </svg>
          </div>
          <h3 class="text-lg font-medium text-neutral-900 mb-1">No players in roster</h3>
          <p class="text-neutral-500 mb-4">Add players to track their performance and highlights</p>
          <button
            @click="navigateToRoster"
            class="inline-flex items-center gap-2 px-4 py-2 bg-primary-500 text-white rounded-lg hover:bg-primary-600 transition-colors"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
            </svg>
            Add Players
          </button>
        </div>
      </div>

      <!-- Match Highlights Section -->
      <div data-highlights class="bg-white rounded-2xl border border-neutral-200 p-4 sm:p-6">
        <!-- Header -->
        <div class="flex flex-col gap-3 mb-4">
          <div class="flex items-center justify-between gap-3">
            <h2 class="text-base sm:text-lg font-semibold text-neutral-900 flex items-center gap-2">
              <svg class="w-5 h-5 text-emerald-500 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
              </svg>
              <span>Highlights</span>
              <span class="bg-emerald-100 text-emerald-700 text-[10px] sm:text-xs px-2 py-0.5 rounded-full font-medium">FREE</span>
            </h2>
            <!-- Add Highlight Button -->
            <button
              v-if="players.length > 0 && !isFutureMatch"
              @click="handleAddHighlight"
              class="flex-shrink-0 w-9 h-9 sm:w-auto sm:h-auto sm:px-4 sm:py-2 flex items-center justify-center gap-2 bg-emerald-500 text-white rounded-lg hover:bg-emerald-600 transition-colors text-sm font-medium"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
              </svg>
              <span class="hidden sm:inline">Add Highlight</span>
            </button>
          </div>
          
          <!-- Filters Row - Equal width on mobile -->
          <div class="grid grid-cols-2 gap-2">
            <!-- Player Filter -->
            <div class="relative">
              <select
                v-model="highlightPlayerFilter"
                class="w-full appearance-none pl-3 pr-8 py-2.5 border border-neutral-200 rounded-xl text-sm bg-white truncate"
              >
                <option value="">All Players</option>
                <option v-for="player in players" :key="player.id" :value="player.player_id">
                  {{ player.first_name }} {{ player.last_name?.charAt(0) }}.
                </option>
              </select>
              <div class="absolute inset-y-0 right-0 pr-2.5 flex items-center pointer-events-none">
                <svg class="w-4 h-4 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                </svg>
              </div>
            </div>
            <!-- Type Filter -->
            <div class="relative">
              <select
                v-model="highlightTypeFilter"
                class="w-full appearance-none pl-3 pr-8 py-2.5 border border-neutral-200 rounded-xl text-sm bg-white"
              >
                <option value="">All Types</option>
                <option v-for="type in highlightTypes" :key="type.value" :value="type.value">
                  {{ type.icon }} {{ type.label }}
                </option>
              </select>
              <div class="absolute inset-y-0 right-0 pr-2.5 flex items-center pointer-events-none">
                <svg class="w-4 h-4 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                </svg>
              </div>
            </div>
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
          <p class="font-medium text-neutral-700 mb-2">No highlights added yet</p>
          <p class="text-sm mb-4" v-if="players.length === 0">Add players to this match first, then upload their highlight clips</p>
          <p class="text-sm mb-4" v-else-if="isFutureMatch">You can add highlights after the match date</p>
          <p class="text-sm mb-4" v-else>Capture the best moments from this match</p>
          <button
            v-if="players.length > 0 && !isFutureMatch"
            @click="handleAddHighlight"
            class="inline-flex items-center gap-2 px-4 py-2 bg-emerald-500 text-white rounded-lg hover:bg-emerald-600 transition-colors text-sm font-medium"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
            </svg>
            Add First Highlight
          </button>
          <button
            v-else-if="players.length === 0"
            @click="navigateToRoster"
            class="inline-flex items-center gap-2 px-4 py-2 bg-primary-500 text-white rounded-lg hover:bg-primary-600 transition-colors text-sm font-medium"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
            </svg>
            Add Players to Match
          </button>
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
                  <svg class="w-5 h-5 text-emerald-500 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                  </svg>
                  <span class="text-sm font-medium truncate max-w-[200px]" :title="highlightFile.name">
                    {{ truncateFilename(highlightFile.name, 25) }}
                  </span>
                  <span class="text-xs text-neutral-500 flex-shrink-0">({{ formatFileSize(highlightFile.size) }})</span>
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

    <!-- Replace Video Modal -->
    <Teleport to="body">
      <div v-if="showReplaceVideoModal" class="fixed inset-0 z-50 overflow-y-auto">
        <div class="fixed inset-0 bg-black/50" @click="showReplaceVideoModal = false"></div>
        <div class="relative min-h-screen flex items-center justify-center p-4">
          <div class="relative bg-white rounded-2xl shadow-xl w-full max-w-lg p-6">
            <h3 class="text-xl font-semibold mb-4">Replace Match Video</h3>
            <p class="text-sm text-neutral-600 mb-4">
              The current video will be deleted and replaced with the new one.
            </p>

            <div
              @click="replaceVideoInput?.click()"
              @drop.prevent="handleReplaceVideoFileDrop"
              @dragover.prevent
              class="border-2 border-dashed border-neutral-300 rounded-xl p-8 text-center cursor-pointer hover:border-neutral-400 transition-colors"
            >
              <input
                ref="replaceVideoInput"
                type="file"
                accept="video/*"
                class="hidden"
                @change="handleReplaceVideoFileSelect"
              />
              <div v-if="!replaceVideoFile">
                <svg class="w-12 h-12 text-neutral-400 mx-auto mb-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
                </svg>
                <p class="text-neutral-900 font-medium mb-1">Select new video file</p>
                <p class="text-sm text-neutral-500">MP4, MOV, WebM (max 25GB)</p>
              </div>
              <div v-else class="flex flex-col items-center">
                <svg class="w-10 h-10 text-emerald-500 mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                </svg>
                <p class="font-medium">{{ replaceVideoFile.name }}</p>
                <p class="text-sm text-neutral-500">{{ formatFileSize(replaceVideoFile.size) }}</p>
              </div>
            </div>

            <!-- Upload Progress -->
            <div v-if="replaceVideoUploadProgress !== null" class="mt-4">
              <div class="flex items-center justify-between text-sm mb-1">
                <span>{{ replaceVideoStep }}</span>
                <span>{{ replaceVideoUploadProgress }}%</span>
              </div>
              <div class="h-3 bg-neutral-200 rounded-full overflow-hidden">
                <div
                  class="h-full bg-gradient-to-r from-primary-500 to-emerald-600 transition-all"
                  :style="{ width: `${replaceVideoUploadProgress}%` }"
                ></div>
              </div>
            </div>

            <div class="flex justify-end gap-3 mt-6">
              <button
                @click="closeReplaceVideoModal"
                :disabled="replacingVideo"
                class="px-4 py-2 border border-neutral-200 rounded-lg hover:bg-neutral-50 transition-colors disabled:opacity-50"
              >
                Cancel
              </button>
              <button
                @click="replaceMatchVideo"
                :disabled="!replaceVideoFile || replacingVideo"
                class="px-4 py-2 bg-amber-500 text-white rounded-lg hover:bg-amber-600 transition-colors disabled:opacity-50"
              >
                {{ replacingVideo ? 'Replacing...' : 'Replace Video' }}
              </button>
            </div>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- Edit Match Modal -->
    <Teleport to="body">
      <div v-if="showEditModal" class="fixed inset-0 z-50 overflow-y-auto">
        <div class="fixed inset-0 bg-black/50" @click="showEditModal = false"></div>
        <div class="relative min-h-screen flex items-center justify-center p-4">
          <div class="relative bg-white rounded-2xl shadow-xl w-full max-w-lg p-6 max-h-[90vh] overflow-y-auto">
            <h3 class="text-xl font-semibold mb-6">Edit Match</h3>

            <form @submit.prevent="updateMatch" class="space-y-5">
              <!-- Match Title -->
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-1">Match Title *</label>
                <input
                  v-model="editMatchForm.title"
                  type="text"
                  class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                  placeholder="e.g., Semi Final - Game 1"
                  required
                />
              </div>

              <!-- Teams Section -->
              <div class="space-y-3">
                <div class="flex items-center gap-2">
                  <svg class="w-4 h-4 text-primary-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0z" />
                  </svg>
                  <span class="text-sm font-medium text-neutral-700">Teams</span>
                </div>
                <div class="grid grid-cols-2 gap-4">
                  <div>
                    <label class="block text-xs text-neutral-500 mb-1">Home Team</label>
                    <input
                      v-model="editMatchForm.home_team"
                      type="text"
                      class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                      placeholder="e.g., Nigeria U-17"
                    />
                  </div>
                  <div>
                    <label class="block text-xs text-neutral-500 mb-1">Away Team</label>
                    <input
                      v-model="editMatchForm.away_team"
                      type="text"
                      class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                      placeholder="e.g., Ghana U-17"
                    />
                  </div>
                </div>
              </div>

              <!-- Date & Stage Row -->
              <div class="grid grid-cols-2 gap-4">
                <div>
                  <label class="block text-sm font-medium text-neutral-700 mb-1">Match Date *</label>
                  <input
                    v-model="editMatchForm.match_date"
                    type="date"
                    class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                    required
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-neutral-700 mb-1">Stage</label>
                  <select
                    v-model="editMatchForm.stage"
                    class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent bg-white"
                  >
                    <option value="">Select stage</option>
                    <option value="group">Group Stage</option>
                    <option value="round_of_16">Round of 16</option>
                    <option value="quarterfinal">Quarter Final</option>
                    <option value="semifinal">Semi Final</option>
                    <option value="third_place">3rd Place</option>
                    <option value="final">Final</option>
                  </select>
                </div>
              </div>

              <!-- Location -->
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-1">Location</label>
                <input
                  v-model="editMatchForm.location"
                  type="text"
                  class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                  placeholder="e.g., Teslim Balogun Stadium, Lagos"
                />
              </div>

              <!-- Status & Score Section -->
              <div class="space-y-3">
                <div>
                  <label class="block text-sm font-medium text-neutral-700 mb-1">Status</label>
                  <select
                    v-model="editMatchForm.status"
                    class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent bg-white"
                  >
                    <option value="scheduled">Scheduled</option>
                    <option value="in_progress">In Progress</option>
                    <option value="completed">Completed</option>
                    <option value="postponed">Postponed</option>
                    <option value="cancelled">Cancelled</option>
                  </select>
                </div>

                <!-- Score (only show when match is in progress or completed) -->
                <div v-if="editMatchForm.status === 'completed' || editMatchForm.status === 'in_progress'" class="p-4 bg-neutral-50 rounded-xl space-y-3">
                  <div class="flex items-center gap-2">
                    <svg class="w-4 h-4 text-emerald-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                    </svg>
                    <span class="text-sm font-medium text-neutral-700">Final Score</span>
                  </div>
                  <div class="grid grid-cols-2 gap-4">
                    <div>
                      <label class="block text-xs text-neutral-500 mb-1">{{ editMatchForm.home_team || 'Home' }}</label>
                      <input
                        v-model="editMatchForm.home_score"
                        type="number"
                        min="0"
                        class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent text-center text-lg font-semibold"
                        placeholder="0"
                      />
                    </div>
                    <div>
                      <label class="block text-xs text-neutral-500 mb-1">{{ editMatchForm.away_team || 'Away' }}</label>
                      <input
                        v-model="editMatchForm.away_score"
                        type="number"
                        min="0"
                        class="w-full px-4 py-3 border border-neutral-200 rounded-xl focus:ring-2 focus:ring-primary-500 focus:border-transparent text-center text-lg font-semibold"
                        placeholder="0"
                      />
                    </div>
                  </div>
                </div>
              </div>

              <!-- Modal Footer -->
              <div class="flex gap-3 pt-4 border-t border-neutral-100">
                <button
                  type="button"
                  @click="showEditModal = false"
                  class="w-auto px-6 py-3 border border-neutral-300 text-neutral-700 rounded-xl text-sm font-medium hover:bg-neutral-50 transition-colors"
                >
                  Cancel
                </button>
                <button
                  type="submit"
                  :disabled="updatingMatch"
                  class="flex-1 inline-flex items-center justify-center gap-2 px-4 py-3 bg-gradient-to-r from-primary-500 to-emerald-600 text-white rounded-xl text-sm font-semibold hover:from-primary-600 hover:to-emerald-700 transition-all disabled:opacity-50 disabled:cursor-not-allowed"
                >
                  <svg v-if="updatingMatch" class="w-4 h-4 animate-spin" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                  {{ updatingMatch ? 'Saving...' : 'Save Changes' }}
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- Video Preview Modal -->
    <Teleport to="body">
      <div v-if="showVideoPreviewModal" class="fixed inset-0 z-50 overflow-y-auto">
        <div class="fixed inset-0 bg-black/80" @click="showVideoPreviewModal = false"></div>
        <div class="relative min-h-screen flex items-center justify-center p-4">
          <div class="relative bg-black rounded-2xl shadow-xl w-full max-w-4xl overflow-hidden">
            <!-- Close button -->
            <button
              @click="showVideoPreviewModal = false"
              class="absolute top-4 right-4 z-10 p-2 bg-black/50 hover:bg-black/70 text-white rounded-full transition-colors"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>

            <!-- Video Player -->
            <div class="aspect-video bg-black relative">
              <video
                ref="videoPlayerRef"
                v-if="matchVideo?.video_url"
                :src="matchVideo.video_url"
                controls
                autoplay
                crossorigin="anonymous"
                class="w-full h-full"
                @loadeddata="onVideoLoaded"
              >
                Your browser does not support the video tag.
              </video>
              
              <!-- Capture Thumbnail Button (overlay) -->
              <button
                v-if="videoLoaded && !capturingThumbnail"
                @click="captureThumbnail"
                class="absolute bottom-16 left-4 px-3 py-2 bg-amber-500 hover:bg-amber-600 text-white rounded-lg text-sm font-medium flex items-center gap-2 transition-colors shadow-lg"
              >
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 9a2 2 0 012-2h.93a2 2 0 001.664-.89l.812-1.22A2 2 0 0110.07 4h3.86a2 2 0 011.664.89l.812 1.22A2 2 0 0018.07 7H19a2 2 0 012 2v9a2 2 0 01-2 2H5a2 2 0 01-2-2V9z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 13a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
                Capture as Thumbnail
              </button>
              <div
                v-if="capturingThumbnail"
                class="absolute bottom-16 left-4 px-3 py-2 bg-neutral-700 text-white rounded-lg text-sm flex items-center gap-2"
              >
                <svg class="w-4 h-4 animate-spin" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                Capturing...
              </div>
            </div>

            <!-- Hidden canvas for thumbnail capture -->
            <canvas ref="thumbnailCanvas" class="hidden"></canvas>

            <!-- Video Info -->
            <div class="p-4 bg-neutral-900 text-white">
              <div class="flex items-center justify-between">
                <div>
                  <h3 class="font-semibold">{{ match?.title }}</h3>
                  <div class="flex items-center gap-4 mt-2 text-sm text-neutral-400">
                    <span v-if="matchVideo?.duration_seconds">{{ formatDuration(matchVideo.duration_seconds) }}</span>
                    <span v-if="matchVideo?.file_size_bytes">{{ formatFileSize(matchVideo.file_size_bytes) }}</span>
                  </div>
                </div>
                <div class="text-xs text-neutral-500">
                  Pause video and click "Capture as Thumbnail" to set cover image
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- Highlight Preview Modal -->
    <Teleport to="body">
      <div v-if="showHighlightPreviewModal && selectedHighlight" class="fixed inset-0 z-50 overflow-y-auto">
        <div class="fixed inset-0 bg-black/80" @click="showHighlightPreviewModal = false"></div>
        <div class="relative min-h-screen flex items-center justify-center p-4">
          <div class="relative bg-black rounded-2xl shadow-xl w-full max-w-3xl overflow-hidden">
            <!-- Close button -->
            <button
              @click="showHighlightPreviewModal = false"
              class="absolute top-4 right-4 z-10 p-2 bg-black/50 hover:bg-black/70 text-white rounded-full transition-colors"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>

            <!-- Video Player -->
            <div class="aspect-video bg-black relative">
              <video
                v-if="selectedHighlight.video_url"
                :src="selectedHighlight.video_url"
                controls
                autoplay
                class="w-full h-full"
              >
                Your browser does not support the video tag.
              </video>
              <div v-else class="w-full h-full flex items-center justify-center text-neutral-500">
                <div class="text-center">
                  <svg class="w-16 h-16 mx-auto mb-2 opacity-50" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
                  </svg>
                  <p>Video not available</p>
                </div>
              </div>
            </div>

            <!-- Highlight Info -->
            <div class="p-4 bg-neutral-900 text-white">
              <div class="flex items-center justify-between">
                <div>
                  <div class="flex items-center gap-3 mb-2">
                    <span class="bg-emerald-500/20 text-emerald-400 px-2 py-1 rounded text-sm font-medium">
                      {{ getHighlightIcon(selectedHighlight.highlight_type) }} {{ selectedHighlight.highlight_type }}
                    </span>
                    <span v-if="selectedHighlight.duration_seconds" class="text-neutral-400 text-sm">
                      {{ formatDuration(selectedHighlight.duration_seconds) }}
                    </span>
                  </div>
                  <h3 class="font-semibold">
                    {{ selectedHighlight.player?.first_name }} {{ selectedHighlight.player?.last_name }}
                  </h3>
                  <p class="text-sm text-neutral-400 mt-1">
                    {{ match?.home_team }} vs {{ match?.away_team }}
                  </p>
                </div>
                <button
                  @click="deleteHighlight(selectedHighlight); showHighlightPreviewModal = false"
                  class="p-2 bg-rose-500/20 hover:bg-rose-500/30 text-rose-400 rounded-lg transition-colors"
                  title="Delete Highlight"
                >
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- Add Highlight Player Picker Modal -->
    <Teleport to="body">
      <div v-if="showAddHighlightPicker" class="fixed inset-0 z-50 overflow-y-auto">
        <div class="fixed inset-0 bg-black/50" @click="showAddHighlightPicker = false"></div>
        <div class="relative min-h-screen flex items-center justify-center p-4">
          <div class="relative bg-white rounded-2xl shadow-xl w-full max-w-md p-6">
            <div class="flex items-center justify-between mb-6">
              <h3 class="text-xl font-semibold text-neutral-900">Add Highlight</h3>
              <button
                @click="showAddHighlightPicker = false"
                class="p-2 hover:bg-neutral-100 rounded-lg transition-colors"
              >
                <svg class="w-5 h-5 text-neutral-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            </div>

            <p class="text-sm text-neutral-600 mb-4">Select a player to add a highlight for:</p>

            <!-- Player List -->
            <div class="space-y-2 max-h-80 overflow-y-auto">
              <button
                v-for="player in players"
                :key="player.id"
                @click="selectPlayerForHighlight(player)"
                class="w-full flex items-center gap-3 p-3 rounded-xl border border-neutral-200 hover:border-emerald-300 hover:bg-emerald-50 transition-all text-left"
              >
                <div class="w-10 h-10 rounded-full bg-neutral-200 flex items-center justify-center overflow-hidden flex-shrink-0">
                  <img
                    v-if="player.profile_photo_url"
                    :src="player.profile_photo_url"
                    :alt="player.first_name"
                    class="w-full h-full object-cover"
                  />
                  <span v-else class="text-sm font-medium text-neutral-500">
                    {{ player.first_name?.charAt(0) }}{{ player.last_name?.charAt(0) }}
                  </span>
                </div>
                <div class="flex-1 min-w-0">
                  <p class="font-medium text-neutral-900">{{ player.first_name }} {{ player.last_name }}</p>
                  <p class="text-sm text-neutral-500">{{ player.position }} • {{ player.highlight_count || 0 }} highlights</p>
                </div>
                <svg class="w-5 h-5 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
                </svg>
              </button>
            </div>

            <!-- No players state -->
            <div v-if="players.length === 0" class="text-center py-8 text-neutral-500">
              <p>No players in this match yet</p>
              <button
                @click="showAddHighlightPicker = false; navigateToRoster()"
                class="mt-2 text-primary-600 hover:text-primary-700 text-sm font-medium"
              >
                Add players first →
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
  video_url?: string
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
const { confirm } = useConfirm()

const tournamentId = route.params.id as string
const matchId = route.params.matchId as string

// State
const loading = ref(true)
const match = ref<Match | null>(null)
const matchVideo = ref<MatchVideo | null>(null)
const players = ref<MatchPlayer[]>([])
const highlights = ref<PlayerHighlight[]>([])
const highlightTypeFilter = ref('')
const highlightPlayerFilter = ref('')

// Modals
const showEditModal = ref(false)
const showAddPlayerModal = ref(false)
const showHighlightUploadModal = ref(false)
const showUploadVideoModal = ref(false)
const showReplaceVideoModal = ref(false)
const showHighlightPreviewModal = ref(false)
const showAddHighlightPicker = ref(false)
const selectedHighlight = ref<PlayerHighlight | null>(null)
const updatingMatch = ref(false)

// Edit match form
const editMatchForm = reactive({
  title: '',
  home_team: '',
  away_team: '',
  home_score: null as number | null,
  away_score: null as number | null,
  match_date: '',
  location: '',
  stage: '',
  status: 'scheduled',
})

// Initialize edit form when modal opens
watch(showEditModal, (isOpen) => {
  if (isOpen && match.value) {
    editMatchForm.title = match.value.title || ''
    editMatchForm.home_team = match.value.home_team || ''
    editMatchForm.away_team = match.value.away_team || ''
    editMatchForm.home_score = match.value.home_score ?? null
    editMatchForm.away_score = match.value.away_score ?? null
    editMatchForm.match_date = match.value.match_date?.split('T')[0] || ''
    editMatchForm.location = match.value.location || ''
    editMatchForm.stage = match.value.stage || ''
    editMatchForm.status = match.value.status || 'scheduled'
  }
})

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

// Replace video
const replaceVideoFile = ref<File | null>(null)
const replaceVideoUploadProgress = ref<number | null>(null)
const replaceVideoStep = ref('Uploading...')
const replacingVideo = ref(false)
const replaceVideoInput = ref<HTMLInputElement | null>(null)

// Thumbnail upload
const thumbnailInput = ref<HTMLInputElement | null>(null)
const uploadingThumbnail = ref(false)

// Video player for thumbnail capture
const videoPlayerRef = ref<HTMLVideoElement | null>(null)
const thumbnailCanvas = ref<HTMLCanvasElement | null>(null)
const videoLoaded = ref(false)
const capturingThumbnail = ref(false)

function onVideoLoaded() {
  videoLoaded.value = true
}

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
  let result = highlights.value
  
  // Filter by player
  if (highlightPlayerFilter.value) {
    result = result.filter(h => h.player_id === highlightPlayerFilter.value)
  }
  
  // Filter by type
  if (highlightTypeFilter.value) {
    result = result.filter(h => h.highlight_type === highlightTypeFilter.value)
  }
  
  return result
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

// Navigation
function navigateToRoster() {
  const rosterPath = `/admin/tournaments/${tournamentId}/matches/${matchId}/roster`
  console.log('Navigating to roster:', rosterPath)
  router.push(rosterPath)
}

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

async function updateMatch() {
  updatingMatch.value = true
  try {
    const response = await $fetch<ApiResponse<Match>>(`/admin/matches/${matchId}`, {
      baseURL: config.public.apiBase,
      method: 'PUT',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: {
        title: editMatchForm.title,
        home_team: editMatchForm.home_team || undefined,
        away_team: editMatchForm.away_team || undefined,
        home_score: editMatchForm.home_score,
        away_score: editMatchForm.away_score,
        match_date: editMatchForm.match_date,
        location: editMatchForm.location || undefined,
        stage: editMatchForm.stage || undefined,
        status: editMatchForm.status,
      },
    })
    if (response.success) {
      toast.success('Match Updated', 'Match details have been saved')
      showEditModal.value = false
      await fetchMatch()
    }
  } catch (error: any) {
    toast.error('Error', error.data?.message || 'Failed to update match')
  } finally {
    updatingMatch.value = false
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
              // Map stream_url to video_url for the frontend
              video_url: h.stream_url || h.video_url,
              player_id: (group as any).player_id,
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
  const confirmed = await confirm({
    title: 'Remove Player',
    message: `Are you sure you want to remove ${player.first_name} ${player.last_name} from this match? This action cannot be undone.`,
    confirmText: 'Remove',
    cancelText: 'Cancel',
    confirmVariant: 'danger',
    icon: 'trash',
  })
  
  if (!confirmed) return
  
  try {
    await $fetch(`/admin/matches/${matchId}/players/${player.player_id}`, {
      baseURL: config.public.apiBase,
      method: 'DELETE',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
    })
    toast.success('Player Removed', `${player.first_name} ${player.last_name} has been removed from the match`)
    fetchMatch()
  } catch (error) {
    toast.error('Error', 'Failed to remove player')
  }
}

function openHighlightUpload(player: MatchPlayer) {
  selectedPlayerForHighlight.value = player
  showHighlightUploadModal.value = true
}

// Smart add highlight - if player is already filtered, skip picker
function handleAddHighlight() {
  if (highlightPlayerFilter.value) {
    // Find the player from the filter
    const filteredPlayer = players.value.find(p => p.player_id === highlightPlayerFilter.value)
    if (filteredPlayer) {
      openHighlightUpload(filteredPlayer)
      return
    }
  }
  // Otherwise show the player picker
  showAddHighlightPicker.value = true
}

function selectPlayerForHighlight(player: MatchPlayer) {
  showAddHighlightPicker.value = false
  openHighlightUpload(player)
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
    // 0. Auto-generate thumbnail from video
    highlightUploadProgress.value = 0
    const thumbnailBlob = await generateThumbnailFromFile(highlightFile.value, 1)

    // 1. Init upload
    highlightUploadProgress.value = 5
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
          // Progress from 10% to 80%
          highlightUploadProgress.value = 10 + Math.round((e.loaded / e.total) * 70)
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
    highlightUploadProgress.value = 85
    const createResponse = await $fetch<ApiResponse<{ id: string }>>(`/admin/highlights`, {
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

    // 4. Upload auto-generated thumbnail (if available)
    if (thumbnailBlob && createResponse.success && createResponse.data?.id) {
      highlightUploadProgress.value = 90
      await uploadHighlightThumbnail(createResponse.data.id, thumbnailBlob)
    }

    highlightUploadProgress.value = 100
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

// Upload thumbnail for a highlight
async function uploadHighlightThumbnail(highlightId: string, thumbnailBlob: Blob): Promise<boolean> {
  try {
    const fileName = `thumbnail_${Date.now()}.jpg`
    
    // Get presigned URL for thumbnail upload
    const initResponse = await $fetch<ApiResponse<{ upload_url: string; s3_key: string }>>(`/admin/highlights/${highlightId}/thumbnail/upload`, {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: {
        file_name: fileName,
        file_size: thumbnailBlob.size,
        content_type: 'image/jpeg',
      },
    })

    if (!initResponse.success || !initResponse.data?.upload_url) {
      return false
    }

    // Upload to S3
    const uploadRes = await fetch(initResponse.data.upload_url, {
      method: 'PUT',
      headers: { 'Content-Type': 'image/jpeg' },
      body: thumbnailBlob,
    })

    if (!uploadRes.ok) return false

    // Save thumbnail URL
    await $fetch(`/admin/highlights/${highlightId}/thumbnail`, {
      baseURL: config.public.apiBase,
      method: 'PUT',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: { s3_key: initResponse.data.s3_key },
    })

    return true
  } catch (error) {
    console.error('Failed to upload highlight thumbnail:', error)
    return false
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

// Auto-generate thumbnail from video file at specified time (in seconds)
async function generateThumbnailFromFile(videoFile: File, timeInSeconds: number = 2): Promise<Blob | null> {
  return new Promise((resolve) => {
    const video = document.createElement('video')
    video.preload = 'auto'
    video.muted = true
    video.playsInline = true
    video.crossOrigin = 'anonymous'
    
    const url = URL.createObjectURL(videoFile)
    video.src = url
    
    let resolved = false
    
    const cleanup = () => {
      if (!resolved) {
        resolved = true
        URL.revokeObjectURL(url)
      }
    }
    
    video.onloadeddata = () => {
      // Seek to the specified time or 10% of video if too short
      const seekTime = Math.min(timeInSeconds, video.duration * 0.1)
      video.currentTime = seekTime
    }
    
    video.onseeked = () => {
      if (video.videoWidth === 0 || video.videoHeight === 0) {
        cleanup()
        resolve(null)
        return
      }
      
      const canvas = document.createElement('canvas')
      canvas.width = video.videoWidth
      canvas.height = video.videoHeight
      
      const ctx = canvas.getContext('2d')
      if (!ctx) {
        console.error('Failed to get canvas context')
        cleanup()
        resolve(null)
        return
      }
      
      ctx.drawImage(video, 0, 0, canvas.width, canvas.height)
      
      canvas.toBlob((blob) => {
        cleanup()
        resolve(blob)
      }, 'image/jpeg', 0.85)
    }
    
    video.onerror = (e) => {
      console.error('Video load error:', e)
      cleanup()
      resolve(null)
    }
    
    // Start loading
    video.load()
    
    // Timeout fallback (15 seconds for large files)
    setTimeout(() => {
      if (!resolved) {
        console.warn('Thumbnail generation timed out')
        cleanup()
        resolve(null)
      }
    }, 15000)
  })
}

// Upload thumbnail blob to S3
async function uploadThumbnailBlob(thumbnailBlob: Blob): Promise<boolean> {
  try {
    const fileName = `thumbnail_${Date.now()}.jpg`
    
    // Get presigned URL for thumbnail upload
    const initResponse = await $fetch<ApiResponse<{ upload_url: string; s3_key: string }>>(`/admin/matches/${matchId}/video/thumbnail/upload`, {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: {
        file_name: fileName,
        file_size: thumbnailBlob.size,
        content_type: 'image/jpeg',
      },
    })

    if (!initResponse.success || !initResponse.data?.upload_url) {
      return false
    }

    // Upload to S3
    const uploadRes = await fetch(initResponse.data.upload_url, {
      method: 'PUT',
      headers: { 'Content-Type': 'image/jpeg' },
      body: thumbnailBlob,
    })

    if (!uploadRes.ok) return false

    // Save thumbnail URL
    await $fetch(`/admin/matches/${matchId}/video/thumbnail`, {
      baseURL: config.public.apiBase,
      method: 'PUT',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: { s3_key: initResponse.data.s3_key },
    })

    return true
  } catch (error) {
    console.error('Failed to upload thumbnail:', error)
    return false
  }
}

async function uploadMatchVideo() {
  if (!matchVideoFile.value) return
  uploadingMatchVideo.value = true

  try {
    // 0. Auto-generate thumbnail from video
    matchVideoUploadProgress.value = 0
    const thumbnailBlob = await generateThumbnailFromFile(matchVideoFile.value, 2)
    
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

    // 4. Upload auto-generated thumbnail (if available)
    if (thumbnailBlob) {
      await uploadThumbnailBlob(thumbnailBlob)
    }

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
  const confirmed = await confirm({
    title: 'Delete Video',
    message: 'Are you sure you want to delete this match video? This action cannot be undone.',
    confirmText: 'Delete',
    cancelText: 'Cancel',
    confirmVariant: 'danger',
    icon: 'trash',
  })
  
  if (!confirmed) return
  
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

// Replace video handlers
function handleReplaceVideoFileSelect(e: Event) {
  const input = e.target as HTMLInputElement
  if (input.files && input.files[0]) {
    replaceVideoFile.value = input.files[0]
  }
}

function handleReplaceVideoFileDrop(e: DragEvent) {
  if (e.dataTransfer?.files && e.dataTransfer.files[0]) {
    replaceVideoFile.value = e.dataTransfer.files[0]
  }
}

function closeReplaceVideoModal() {
  showReplaceVideoModal.value = false
  replaceVideoFile.value = null
  replaceVideoUploadProgress.value = null
  replaceVideoStep.value = 'Uploading...'
}

async function replaceMatchVideo() {
  if (!replaceVideoFile.value) return
  replacingVideo.value = true

  try {
    // Step 0: Generate thumbnail from new video
    replaceVideoStep.value = 'Generating thumbnail...'
    replaceVideoUploadProgress.value = 5
    const thumbnailBlob = await generateThumbnailFromFile(replaceVideoFile.value, 2)

    // Step 1: Delete old video
    replaceVideoStep.value = 'Removing old video...'
    replaceVideoUploadProgress.value = 10
    
    await $fetch(`/admin/matches/${matchId}/video`, {
      baseURL: config.public.apiBase,
      method: 'DELETE',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
    })

    // Step 2: Init new upload
    replaceVideoStep.value = 'Preparing upload...'
    replaceVideoUploadProgress.value = 20

    const initResponse = await $fetch<ApiResponse<{ session_id: string; upload_url?: string; s3_key: string; upload_method: string }>>(`/admin/matches/${matchId}/video/upload`, {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: {
        file_name: replaceVideoFile.value.name,
        file_size: replaceVideoFile.value.size,
        content_type: replaceVideoFile.value.type,
      },
    })

    if (!initResponse.success || !initResponse.data) {
      throw new Error('Failed to initialize upload')
    }

    // Step 3: Upload to S3
    replaceVideoStep.value = 'Uploading new video...'
    
    if (initResponse.data.upload_url) {
      await new Promise<void>((resolve, reject) => {
        const xhr = new XMLHttpRequest()
        xhr.upload.onprogress = (e) => {
          if (e.lengthComputable) {
            // Progress from 25% to 85%
            replaceVideoUploadProgress.value = 25 + Math.round((e.loaded / e.total) * 60)
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
        xhr.setRequestHeader('Content-Type', replaceVideoFile.value!.type)
        xhr.send(replaceVideoFile.value)
      })
    }

    // Step 4: Save video record
    replaceVideoStep.value = 'Saving video...'
    replaceVideoUploadProgress.value = 90

    await $fetch(`/admin/matches/${matchId}/video`, {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: {
        s3_key: initResponse.data.s3_key,
        file_size_bytes: replaceVideoFile.value.size,
      },
    })

    // Step 5: Upload auto-generated thumbnail
    if (thumbnailBlob) {
      replaceVideoStep.value = 'Uploading thumbnail...'
      replaceVideoUploadProgress.value = 95
      await uploadThumbnailBlob(thumbnailBlob)
    }

    replaceVideoUploadProgress.value = 100
    toast.success('Video Replaced', 'Match video has been replaced successfully')
    closeReplaceVideoModal()
    fetchMatch()
  } catch (error: any) {
    toast.error('Replace Failed', error.message || 'Failed to replace video')
  } finally {
    replacingVideo.value = false
  }
}

// Thumbnail upload handler
async function handleThumbnailUpload(e: Event) {
  const input = e.target as HTMLInputElement
  if (!input.files || !input.files[0]) return
  
  const file = input.files[0]
  if (!file.type.startsWith('image/')) {
    toast.error('Invalid File', 'Please select an image file')
    return
  }

  uploadingThumbnail.value = true
  try {
    // Get presigned URL for thumbnail upload
    const initResponse = await $fetch<ApiResponse<{ upload_url: string; s3_key: string }>>(`/admin/matches/${matchId}/video/thumbnail/upload`, {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: {
        file_name: file.name,
        file_size: file.size,
        content_type: file.type,
      },
    })

    if (!initResponse.success || !initResponse.data?.upload_url) {
      throw new Error('Failed to initialize thumbnail upload')
    }

    // Upload to S3
    await fetch(initResponse.data.upload_url, {
      method: 'PUT',
      headers: { 'Content-Type': file.type },
      body: file,
    })

    // Save thumbnail URL
    await $fetch(`/admin/matches/${matchId}/video/thumbnail`, {
      baseURL: config.public.apiBase,
      method: 'PUT',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: { s3_key: initResponse.data.s3_key },
    })

    toast.success('Thumbnail Updated', 'Video thumbnail has been updated')
    fetchMatch()
  } catch (error: any) {
    toast.error('Upload Failed', error.message || 'Failed to upload thumbnail')
  } finally {
    uploadingThumbnail.value = false
    input.value = '' // Reset input
  }
}

// Capture thumbnail from video using canvas
async function captureThumbnail() {
  const video = videoPlayerRef.value
  const canvas = thumbnailCanvas.value
  
  if (!video || !canvas) {
    toast.error('Error', 'Video player not ready')
    return
  }

  // Pause video for capture
  video.pause()
  
  capturingThumbnail.value = true
  
  try {
    // Set canvas dimensions to match video
    canvas.width = video.videoWidth
    canvas.height = video.videoHeight
    
    // Draw current frame to canvas
    const ctx = canvas.getContext('2d')
    if (!ctx) {
      throw new Error('Canvas context not available')
    }
    ctx.drawImage(video, 0, 0, canvas.width, canvas.height)
    
    // Convert canvas to blob
    const blob = await new Promise<Blob>((resolve, reject) => {
      canvas.toBlob((b) => {
        if (b) resolve(b)
        else reject(new Error('Failed to create image'))
      }, 'image/jpeg', 0.9)
    })
    
    // Create file from blob
    const fileName = `thumbnail_${Date.now()}.jpg`
    const file = new File([blob], fileName, { type: 'image/jpeg' })
    
    // Get presigned URL for thumbnail upload
    const initResponse = await $fetch<ApiResponse<{ upload_url: string; s3_key: string }>>(`/admin/matches/${matchId}/video/thumbnail/upload`, {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: {
        file_name: fileName,
        file_size: file.size,
        content_type: 'image/jpeg',
      },
    })

    if (!initResponse.success || !initResponse.data?.upload_url) {
      throw new Error('Failed to initialize thumbnail upload')
    }

    // Upload to S3
    await fetch(initResponse.data.upload_url, {
      method: 'PUT',
      headers: { 'Content-Type': 'image/jpeg' },
      body: file,
    })

    // Save thumbnail URL
    await $fetch(`/admin/matches/${matchId}/video/thumbnail`, {
      baseURL: config.public.apiBase,
      method: 'PUT',
      headers: { Authorization: `Bearer ${authStore.accessToken}` },
      body: { s3_key: initResponse.data.s3_key },
    })

    toast.success('Thumbnail Captured', 'Video thumbnail has been set from current frame')
    showVideoPreviewModal.value = false
    fetchMatch()
  } catch (error: any) {
    console.error('Thumbnail capture error:', error)
    toast.error('Capture Failed', error.message || 'Failed to capture thumbnail. This may be due to CORS restrictions.')
  } finally {
    capturingThumbnail.value = false
  }
}

async function deleteHighlight(highlight: PlayerHighlight) {
  const confirmed = await confirm({
    title: 'Delete Highlight',
    message: 'Are you sure you want to delete this highlight? This action cannot be undone.',
    confirmText: 'Delete',
    cancelText: 'Cancel',
    confirmVariant: 'danger',
    icon: 'trash',
  })
  
  if (!confirmed) return
  
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
  // Set the player filter to show only this player's highlights
  highlightPlayerFilter.value = player.player_id
  highlightTypeFilter.value = ''
  // Scroll to highlights section
  nextTick(() => {
    document.querySelector('[data-highlights]')?.scrollIntoView({ behavior: 'smooth' })
  })
}

function viewHighlight(highlight: PlayerHighlight) {
  selectedHighlight.value = highlight
  showHighlightPreviewModal.value = true
}

// Video preview modal state
const showVideoPreviewModal = ref(false)

// Reset video state when modal closes
watch(showVideoPreviewModal, (isOpen) => {
  if (!isOpen) {
    videoLoaded.value = false
  }
})

function previewVideo() {
  if (matchVideo.value?.video_url) {
    showVideoPreviewModal.value = true
  } else {
    toast.error('Error', 'Video URL not available')
  }
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

// Truncate long filenames while preserving extension
function truncateFilename(filename: string, maxLength: number): string {
  if (filename.length <= maxLength) return filename
  const ext = filename.slice(filename.lastIndexOf('.'))
  const name = filename.slice(0, filename.lastIndexOf('.'))
  const truncatedName = name.slice(0, maxLength - ext.length - 3) + '...'
  return truncatedName + ext
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
