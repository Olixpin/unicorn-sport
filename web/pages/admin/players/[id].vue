<template>
  <div class="max-w-6xl mx-auto">
    <!-- Page Header -->
    <div class="mb-8">
      <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
        <div class="flex items-center gap-4">
          <NuxtLink 
            to="/admin/players" 
            class="w-10 h-10 flex items-center justify-center rounded-xl bg-neutral-100 hover:bg-neutral-200 text-neutral-600 transition-colors"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
            </svg>
          </NuxtLink>
          <div>
            <h1 class="font-display text-2xl lg:text-3xl font-bold text-neutral-900">
              Edit Player
            </h1>
            <p class="mt-1 text-neutral-600">Update player information and verification status.</p>
          </div>
        </div>
        <div v-if="player" class="flex items-center gap-3">
          <span :class="getVerificationBadgeClass(player.verification_status === 'verified')">
            <span class="w-2 h-2 rounded-full" :class="player.verification_status === 'verified' ? 'bg-emerald-500' : 'bg-amber-500'"></span>
            {{ player.verification_status === 'verified' ? 'Verified Player' : 'Pending Verification' }}
          </span>
        </div>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex flex-col items-center justify-center py-20">
      <div class="w-12 h-12 border-4 border-primary-200 border-t-primary-600 rounded-full animate-spin"></div>
      <p class="mt-4 text-neutral-600">Loading player data...</p>
    </div>

    <!-- Not Found -->
    <div v-else-if="!player" class="bg-white rounded-2xl border border-neutral-200 px-6 py-16 text-center">
      <div class="w-16 h-16 bg-neutral-100 rounded-2xl flex items-center justify-center mx-auto mb-4">
        <svg class="w-8 h-8 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
        </svg>
      </div>
      <h3 class="font-semibold text-neutral-900 mb-1">Player not found</h3>
      <p class="text-neutral-500 text-sm mb-6">The player you're looking for doesn't exist or has been removed.</p>
      <NuxtLink to="/admin/players">
        <button class="inline-flex items-center gap-2 px-5 py-2.5 bg-gradient-to-r from-primary-600 to-emerald-600 text-white rounded-xl text-sm font-semibold hover:from-primary-700 hover:to-emerald-700 transition-all shadow-lg shadow-primary-600/25">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
          </svg>
          Back to Players
        </button>
      </NuxtLink>
    </div>

    <!-- Edit Form -->
    <form v-else @submit.prevent="handleSubmit" class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <!-- Main Content (2 columns) -->
      <div class="lg:col-span-2 space-y-6">
        <!-- Basic Information -->
        <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
          <div class="px-6 py-4 border-b border-neutral-200 bg-gradient-to-r from-primary-50 to-emerald-50">
            <div class="flex items-center gap-3">
              <div class="w-10 h-10 bg-gradient-to-br from-primary-500 to-emerald-600 rounded-xl flex items-center justify-center text-white shadow-lg shadow-primary-500/25">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                </svg>
              </div>
              <div>
                <h2 class="font-semibold text-neutral-900">Basic Information</h2>
                <p class="text-sm text-neutral-500">Player identity and personal details</p>
              </div>
            </div>
          </div>

          <div class="p-6 space-y-5">
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-5">
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-2">First Name *</label>
                <input
                  v-model="form.first_name"
                  type="text"
                  placeholder="John"
                  :class="[
                    'w-full px-4 py-3 border rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all',
                    errors.first_name ? 'border-red-300 bg-red-50' : 'border-neutral-200 bg-neutral-50'
                  ]"
                />
                <p v-if="errors.first_name" class="mt-1.5 text-sm text-red-600">{{ errors.first_name }}</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-2">Last Name *</label>
                <input
                  v-model="form.last_name"
                  type="text"
                  placeholder="Doe"
                  :class="[
                    'w-full px-4 py-3 border rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all',
                    errors.last_name ? 'border-red-300 bg-red-50' : 'border-neutral-200 bg-neutral-50'
                  ]"
                />
                <p v-if="errors.last_name" class="mt-1.5 text-sm text-red-600">{{ errors.last_name }}</p>
              </div>
            </div>

            <div class="grid grid-cols-1 sm:grid-cols-2 gap-5">
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-2">Date of Birth *</label>
                <input
                  v-model="form.date_of_birth"
                  type="date"
                  :class="[
                    'w-full px-4 py-3 border rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all',
                    errors.date_of_birth ? 'border-red-300 bg-red-50' : 'border-neutral-200 bg-neutral-50'
                  ]"
                />
                <p v-if="errors.date_of_birth" class="mt-1.5 text-sm text-red-600">{{ errors.date_of_birth }}</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-2">Position *</label>
                <div class="relative">
                  <select
                    v-model="form.position"
                    class="appearance-none w-full px-4 py-3 pr-10 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
                  >
                    <option value="">Select position</option>
                    <option v-for="pos in PLAYER_POSITIONS" :key="pos" :value="pos">
                      {{ pos }}
                    </option>
                  </select>
                  <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                    <svg class="w-5 h-5 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                    </svg>
                  </div>
                </div>
              </div>
            </div>

            <div class="grid grid-cols-1 sm:grid-cols-3 gap-5">
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-2">Preferred Foot</label>
                <div class="relative">
                  <select
                    v-model="form.preferred_foot"
                    class="appearance-none w-full px-4 py-3 pr-10 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
                  >
                    <option value="">Select</option>
                    <option value="right">Right</option>
                    <option value="left">Left</option>
                    <option value="both">Both</option>
                  </select>
                  <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                    <svg class="w-5 h-5 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                    </svg>
                  </div>
                </div>
              </div>
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-2">Height (cm)</label>
                <input
                  v-model.number="form.height_cm"
                  type="number"
                  placeholder="175"
                  min="100"
                  max="250"
                  class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-2">Weight (kg)</label>
                <input
                  v-model.number="form.weight_kg"
                  type="number"
                  placeholder="70"
                  min="30"
                  max="150"
                  class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
                />
              </div>
            </div>
          </div>
        </div>

        <!-- Location -->
        <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
          <div class="px-6 py-4 border-b border-neutral-200 bg-gradient-to-r from-blue-50 to-indigo-50">
            <div class="flex items-center gap-3">
              <div class="w-10 h-10 bg-gradient-to-br from-blue-500 to-indigo-600 rounded-xl flex items-center justify-center text-white shadow-lg shadow-blue-500/25">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
              </div>
              <div>
                <h2 class="font-semibold text-neutral-900">Location</h2>
                <p class="text-sm text-neutral-500">Where the player is based</p>
              </div>
            </div>
          </div>

          <div class="p-6 space-y-5">
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-5">
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-2">Country *</label>
                <div class="relative">
                  <select
                    v-model="form.country"
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
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-2">State/Region</label>
                <!-- Dropdown when country has predefined states -->
                <div v-if="countryHasLocationData && availableStates.length > 0" class="relative">
                  <select
                    v-model="form.state"
                    :disabled="!form.country"
                    class="appearance-none w-full px-4 py-3 pr-10 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all disabled:bg-neutral-100 disabled:cursor-not-allowed"
                  >
                    <option value="">Select state/region</option>
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
                <!-- Text input fallback for countries without predefined data -->
                <input
                  v-else
                  v-model="form.state"
                  type="text"
                  :disabled="!form.country"
                  placeholder="Enter state/region"
                  class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all disabled:bg-neutral-100 disabled:cursor-not-allowed"
                />
              </div>
            </div>

            <div class="grid grid-cols-1 sm:grid-cols-2 gap-5">
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-2">City</label>
                <!-- Dropdown when state has predefined cities -->
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
                <!-- Text input fallback -->
                <input
                  v-else
                  v-model="form.city"
                  type="text"
                  :disabled="!form.country"
                  placeholder="Enter city"
                  class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all disabled:bg-neutral-100 disabled:cursor-not-allowed"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-2">Academy</label>
                <div class="relative">
                  <select
                    v-model="form.academy_id"
                    :disabled="loadingAcademies || !form.country"
                    :class="[
                      'appearance-none w-full px-4 py-3 pr-10 border rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all',
                      loadingAcademies || !form.country ? 'bg-neutral-100 cursor-not-allowed' : 'bg-neutral-50 border-neutral-200'
                    ]"
                  >
                    <option value="" disabled v-if="loadingAcademies">Loading...</option>
                    <option value="" v-else-if="!form.country">Select a country first</option>
                    <option value="" v-else-if="filteredAcademies.length === 0">No academies in {{ form.country }}</option>
                    <option value="">No academy assigned</option>
                    <option v-for="academy in filteredAcademies" :key="academy.id" :value="academy.id">
                      {{ academy.name }}
                    </option>
                  </select>
                  <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                    <svg v-if="loadingAcademies" class="w-5 h-5 text-neutral-400 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" fill="none" />
                      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z" />
                    </svg>
                    <svg v-else class="w-5 h-5 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                    </svg>
                  </div>
                </div>
                <div class="mt-1.5 flex items-center justify-between">
                  <p class="text-xs text-neutral-500">Link player to a registered academy</p>
                  <NuxtLink 
                    to="/admin/academies/new" 
                    class="text-xs font-medium text-primary-600 hover:text-primary-700 flex items-center gap-1 hover:underline"
                  >
                    <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                    </svg>
                    Add new
                  </NuxtLink>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Status & Save -->
        <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
          <div class="px-6 py-4 border-b border-neutral-200 bg-gradient-to-r from-amber-50 to-yellow-50">
            <div class="flex items-center gap-3">
              <div class="w-10 h-10 bg-gradient-to-br from-amber-500 to-yellow-600 rounded-xl flex items-center justify-center text-white shadow-lg shadow-amber-500/25">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </div>
              <div>
                <h2 class="font-semibold text-neutral-900">Status & Save</h2>
                <p class="text-sm text-neutral-500">Update verification status and save changes</p>
              </div>
            </div>
          </div>

          <div class="p-6 space-y-6">
            <!-- Verification Status -->
            <div>
              <label class="block text-sm font-medium text-neutral-700 mb-3">Verification Status</label>
              
              <!-- Status Card -->
              <div class="rounded-xl border-2 overflow-hidden" :class="form.verification_status === 'verified' ? 'border-emerald-200 bg-white' : 'border-neutral-200 bg-white'">
                <!-- Status Header -->
                <div class="px-4 py-3 flex items-center gap-3" :class="form.verification_status === 'verified' ? 'bg-emerald-50' : 'bg-neutral-50'">
                  <div class="w-8 h-8 rounded-full flex items-center justify-center" :class="form.verification_status === 'verified' ? 'bg-emerald-500' : 'bg-neutral-400'">
                    <svg v-if="form.verification_status === 'verified'" class="w-4 h-4 text-white" fill="currentColor" viewBox="0 0 20 20">
                      <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                    </svg>
                    <svg v-else class="w-4 h-4 text-white" fill="currentColor" viewBox="0 0 20 20">
                      <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z" clip-rule="evenodd" />
                    </svg>
                  </div>
                  <div class="flex-1">
                    <p class="text-sm font-semibold" :class="form.verification_status === 'verified' ? 'text-emerald-900' : 'text-neutral-900'">
                      {{ form.verification_status === 'verified' ? 'Verified' : 'Pending Review' }}
                    </p>
                  </div>
                  <span 
                    class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                    :class="form.verification_status === 'verified' ? 'bg-emerald-100 text-emerald-800' : 'bg-neutral-200 text-neutral-700'"
                  >
                    {{ form.verification_status === 'verified' ? 'Active' : 'Unverified' }}
                  </span>
                </div>
                
                <!-- Status Body with Action -->
                <div class="p-4 flex items-center justify-between gap-4">
                  <p class="text-sm text-neutral-600">
                    {{ form.verification_status === 'verified' 
                      ? 'Player displays verification badge and is visible to scouts.' 
                      : 'Complete verification to display trust badge and improve visibility.' 
                    }}
                  </p>
                  <button
                    type="button"
                    @click="form.verification_status = form.verification_status === 'verified' ? 'pending' : 'verified'"
                    class="shrink-0 inline-flex items-center gap-2 px-4 py-2 rounded-lg text-sm font-medium transition-colors"
                    :class="form.verification_status === 'verified' 
                      ? 'text-neutral-700 bg-neutral-100 hover:bg-neutral-200' 
                      : 'text-white bg-neutral-900 hover:bg-neutral-800'"
                  >
                    <svg v-if="form.verification_status !== 'verified'" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
                    </svg>
                    {{ form.verification_status === 'verified' ? 'Revoke' : 'Verify Now' }}
                  </button>
                </div>
              </div>
            </div>

            <!-- Action Buttons -->
            <div class="flex flex-col sm:flex-row gap-3 pt-2">
              <button
                type="submit"
                :disabled="submitting"
                class="flex-1 inline-flex items-center justify-center gap-2 px-6 py-4 bg-gradient-to-r from-primary-600 to-emerald-600 text-white rounded-xl text-base font-semibold hover:from-primary-700 hover:to-emerald-700 disabled:opacity-50 disabled:cursor-not-allowed transition-all shadow-lg shadow-primary-600/25"
              >
                <svg v-if="submitting" class="w-5 h-5 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" fill="none" />
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z" />
                </svg>
                <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                </svg>
                {{ submitting ? 'Saving...' : 'Save Changes' }}
              </button>
              
              <NuxtLink to="/admin/players" class="sm:w-auto">
                <button
                  type="button"
                  class="w-full px-6 py-4 border border-neutral-300 text-neutral-700 rounded-xl text-base font-medium hover:bg-neutral-50 transition-colors"
                >
                  Cancel
                </button>
              </NuxtLink>
            </div>
          </div>
        </div>
      </div>

      <!-- Sidebar -->
      <div class="hidden lg:block space-y-6 lg:sticky lg:top-6 lg:self-start">
        <!-- Profile Photo Upload -->
        <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
          <div class="px-6 py-4 border-b border-neutral-200 bg-gradient-to-r from-blue-50 to-indigo-50">
            <div class="flex items-center gap-3">
              <div class="w-10 h-10 bg-gradient-to-br from-blue-500 to-indigo-600 rounded-xl flex items-center justify-center text-white shadow-lg shadow-blue-500/25">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
              </div>
              <div>
                <h2 class="font-semibold text-neutral-900">Profile Photo</h2>
                <p class="text-sm text-neutral-500">Player headshot or action photo</p>
              </div>
            </div>
          </div>
          <div class="p-6">
            <!-- Current Photo Preview -->
            <div class="relative w-full aspect-square rounded-2xl overflow-hidden bg-neutral-100 mb-4">
              <img 
                v-if="profilePhotoUrl"
                :src="profilePhotoUrl"
                :alt="`${form.first_name} ${form.last_name}`"
                class="w-full h-full object-cover"
              />
              <div v-else class="w-full h-full flex flex-col items-center justify-center text-neutral-400">
                <svg class="w-16 h-16 mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                </svg>
                <span class="text-sm">No photo</span>
              </div>
              <!-- Upload Progress Overlay -->
              <div v-if="uploadingPhoto" class="absolute inset-0 bg-black/50 flex items-center justify-center">
                <div class="text-center text-white">
                  <svg class="w-8 h-8 animate-spin mx-auto mb-2" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
                  </svg>
                  <span class="text-sm">Uploading...</span>
                </div>
              </div>
            </div>

            <!-- Upload Button -->
            <label class="block">
              <input
                type="file"
                accept="image/jpeg,image/png,image/webp"
                class="hidden"
                @change="handlePhotoUpload"
                :disabled="uploadingPhoto"
              />
              <div class="flex items-center justify-center gap-2 px-4 py-3 border-2 border-dashed border-neutral-300 rounded-xl cursor-pointer hover:border-primary-400 hover:bg-primary-50/50 transition-all">
                <svg class="w-5 h-5 text-neutral-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-8l-4-4m0 0L8 8m4-4v12" />
                </svg>
                <span class="text-sm font-medium text-neutral-700">{{ profilePhotoUrl ? 'Change Photo' : 'Upload Photo' }}</span>
              </div>
            </label>
            <p class="mt-2 text-xs text-neutral-500 text-center">JPEG, PNG or WebP. Max 5MB.</p>
          </div>
        </div>

        <!-- Player Preview -->
        <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
          <div class="px-6 py-4 border-b border-neutral-200">
            <h3 class="font-semibold text-neutral-900">Player Preview</h3>
          </div>
          <div class="p-6 flex flex-col items-center">
            <div class="w-20 h-20 rounded-2xl flex items-center justify-center text-white text-2xl font-bold shadow-lg overflow-hidden" :class="profilePhotoUrl ? '' : 'bg-gradient-to-br from-primary-500 to-emerald-600 shadow-primary-500/25'">
              <img v-if="profilePhotoUrl" :src="profilePhotoUrl" class="w-full h-full object-cover" />
              <span v-else>{{ getInitials(form.first_name, form.last_name) }}</span>
            </div>
            <h4 class="mt-4 font-semibold text-neutral-900">
              {{ form.first_name || 'First' }} {{ form.last_name || 'Last' }}
            </h4>
            <p class="text-sm text-neutral-500">{{ form.position || 'Position' }}</p>
            <div class="mt-3 flex items-center gap-2 text-xs text-neutral-400">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
              </svg>
              {{ form.city || 'City' }}, {{ form.country || 'Country' }}
            </div>
          </div>
        </div>

        <!-- Danger Zone -->
        <div class="bg-gradient-to-br from-rose-50 to-red-50 rounded-2xl border border-rose-200 overflow-hidden">
          <div class="px-6 py-4 border-b border-rose-200">
            <div class="flex items-center gap-3">
              <div class="w-8 h-8 bg-rose-100 rounded-lg flex items-center justify-center">
                <svg class="w-4 h-4 text-rose-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                </svg>
              </div>
              <h3 class="font-semibold text-rose-900">Danger Zone</h3>
            </div>
          </div>
          <div class="p-6">
            <p class="text-sm text-rose-700 mb-4">
              Permanently delete this player and all associated data. This action cannot be undone.
            </p>
            <button
              type="button"
              @click="confirmDelete"
              class="w-full px-4 py-2.5 bg-rose-600 text-white text-sm font-medium rounded-xl hover:bg-rose-700 transition-colors"
            >
              Delete Player
            </button>
          </div>
        </div>
      </div>
    </form>

    <!-- Delete Confirmation Modal -->
    <Teleport to="body">
      <div 
        v-if="showDeleteModal" 
        class="fixed inset-0 z-50 flex items-center justify-center p-4"
        @click.self="showDeleteModal = false"
      >
        <div class="absolute inset-0 bg-black/50 backdrop-blur-sm"></div>
        <div class="relative bg-white rounded-2xl shadow-2xl max-w-md w-full p-6">
          <div class="flex items-start gap-4 mb-6">
            <div class="w-12 h-12 bg-rose-100 rounded-xl flex items-center justify-center flex-shrink-0">
              <svg class="w-6 h-6 text-rose-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
              </svg>
            </div>
            <div>
              <h3 class="text-lg font-semibold text-neutral-900">Delete Player</h3>
              <p class="mt-2 text-neutral-600">
                Are you sure you want to delete <span class="font-medium">{{ player?.first_name }} {{ player?.last_name }}</span>? 
                This action cannot be undone and all associated data will be permanently removed.
              </p>
            </div>
          </div>

          <div class="flex justify-end gap-3">
            <button
              type="button"
              @click="showDeleteModal = false"
              class="px-4 py-2.5 border border-neutral-300 text-neutral-700 rounded-xl text-sm font-medium hover:bg-neutral-50 transition-colors"
            >
              Cancel
            </button>
            <button
              type="button"
              @click="deletePlayer"
              :disabled="deleting"
              class="inline-flex items-center gap-2 px-4 py-2.5 bg-rose-600 text-white rounded-xl text-sm font-medium hover:bg-rose-700 disabled:opacity-50 transition-colors"
            >
              <svg v-if="deleting" class="w-4 h-4 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" fill="none" />
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z" />
              </svg>
              {{ deleting ? 'Deleting...' : 'Delete Player' }}
            </button>
          </div>
        </div>
      </div>
    </Teleport>
  </div>
</template>

<script setup lang="ts">
import { PLAYER_POSITIONS } from '~/schemas/player'
import { AFRICAN_COUNTRIES, getStatesForCountry, getCitiesForState, hasLocationData } from '~/data/locations'
import type { Player, ApiResponse } from '~/types/index'

definePageMeta({
  layout: 'admin',
  middleware: 'admin',
})

const route = useRoute()
const router = useRouter()
const toast = useToast()
const config = useRuntimeConfig()
const authStore = useAuthStore()

const playerId = route.params.id as string

const player = ref<Player | null>(null)
const loading = ref(true)
const submitting = ref(false)
const deleting = ref(false)
const showDeleteModal = ref(false)
const errors = reactive<Record<string, string>>({})

// Profile photo state
const profilePhotoUrl = ref<string | null>(null)
const uploadingPhoto = ref(false)

const form = reactive({
  first_name: '',
  last_name: '',
  date_of_birth: '',
  position: '',
  preferred_foot: '' as '' | 'left' | 'right' | 'both',
  height_cm: undefined as number | undefined,
  weight_kg: undefined as number | undefined,
  country: '',
  state: '',
  city: '',
  academy_id: '' as string,
  school_name: '',
  verification_status: 'pending' as 'pending' | 'verified' | 'rejected',
})

// Academies list for dropdown
const academies = ref<Array<{ id: string; name: string; country: string; is_verified: boolean }>>([])
const allAcademies = ref<Array<{ id: string; name: string; country: string; is_verified: boolean }>>([])
const loadingAcademies = ref(false)

// Computed states based on selected country
const availableStates = computed(() => {
  if (!form.country) return []
  return getStatesForCountry(form.country)
})

// Computed cities based on selected state
const availableCities = computed(() => {
  if (!form.country || !form.state) return []
  return getCitiesForState(form.country, form.state)
})

// Filter academies by selected country
const filteredAcademies = computed(() => {
  if (!form.country) return allAcademies.value
  return allAcademies.value.filter(a => a.country === form.country)
})

// Check if country has predefined location data
const countryHasLocationData = computed(() => hasLocationData(form.country))

// Track if initial load is complete to avoid resetting on load
const initialLoadComplete = ref(false)

// Watch country changes to reset state and city (only after initial load)
watch(() => form.country, (newVal, oldVal) => {
  if (initialLoadComplete.value && oldVal !== newVal) {
    form.state = ''
    form.city = ''
    form.academy_id = ''
  }
})

// Watch state changes to reset city (only after initial load)
watch(() => form.state, (newVal, oldVal) => {
  if (initialLoadComplete.value && oldVal !== newVal) {
    form.city = ''
  }
})

function getInitials(firstName?: string, lastName?: string): string {
  const first = firstName?.charAt(0)?.toUpperCase() || ''
  const last = lastName?.charAt(0)?.toUpperCase() || ''
  return first + last || '?'
}

function getVerificationBadgeClass(isVerified?: boolean): string {
  const baseClass = 'inline-flex items-center gap-2 px-3 py-1.5 rounded-full text-sm font-medium'
  return isVerified
    ? `${baseClass} bg-emerald-50 text-emerald-700 border border-emerald-200`
    : `${baseClass} bg-amber-50 text-amber-700 border border-amber-200`
}

// Fetch player data
onMounted(async () => {
  // Fetch academies first
  await fetchAcademies()
  
  try {
    const response = await $fetch<ApiResponse<{ player: Player }>>(`/admin/players/${playerId}`, {
      baseURL: config.public.apiBase,
      headers: {
        Authorization: `Bearer ${authStore.accessToken}`,
      },
    })

    if (response.success && response.data?.player) {
      player.value = response.data.player
      // Set profile photo URL
      profilePhotoUrl.value = player.value.profile_photo_url || null
      // Populate form
      Object.assign(form, {
        first_name: player.value.first_name,
        last_name: player.value.last_name,
        date_of_birth: player.value.date_of_birth.split('T')[0],
        position: player.value.position,
        preferred_foot: player.value.preferred_foot || '',
        height_cm: player.value.height_cm,
        weight_kg: player.value.weight_kg,
        country: player.value.country,
        state: player.value.state || '',
        city: player.value.city || '',
        academy_id: player.value.academy_id || '',
        school_name: player.value.school_name || '',
        verification_status: player.value.verification_status || 'pending',
      })
      // Mark initial load complete after a tick to avoid watchers resetting values
      nextTick(() => {
        initialLoadComplete.value = true
      })
    }
  } catch (error) {
    console.error('Failed to fetch player:', error)
    toast.error('Error', 'Failed to load player data')
  } finally {
    loading.value = false
  }
})

async function fetchAcademies() {
  loadingAcademies.value = true
  try {
    const response = await $fetch<ApiResponse<{ academies: Array<{ id: string; name: string; country: string; is_verified: boolean }> }>>('/admin/academies', {
      baseURL: config.public.apiBase,
      headers: {
        Authorization: `Bearer ${authStore.accessToken}`,
      },
      params: {
        per_page: 200, // Get all academies for dropdown
      },
    })
    if (response.success && response.data?.academies) {
      allAcademies.value = response.data.academies
      academies.value = response.data.academies
    }
  } catch (error) {
    console.error('Failed to fetch academies:', error)
  } finally {
    loadingAcademies.value = false
  }
}

async function handleSubmit() {
  submitting.value = true

  try {
    const payload = {
      ...form,
      preferred_foot: form.preferred_foot || undefined,
      academy_id: form.academy_id || undefined,
    }

    const response = await $fetch<ApiResponse<{ player: Player }>>(`/admin/players/${playerId}`, {
      baseURL: config.public.apiBase,
      method: 'PUT',
      headers: {
        Authorization: `Bearer ${authStore.accessToken}`,
      },
      body: payload,
    })

    if (response.success) {
      toast.success('Player Updated', 'The player has been successfully updated')
      router.push('/admin/players')
    } else {
      toast.error('Error', response.message || 'Failed to update player')
    }
  } catch (error: unknown) {
    const err = error as { data?: { message?: string } }
    toast.error('Error', err.data?.message || 'An unexpected error occurred')
  } finally {
    submitting.value = false
  }
}

function confirmDelete() {
  showDeleteModal.value = true
}

async function deletePlayer() {
  deleting.value = true

  try {
    const response = await $fetch<ApiResponse<null>>(`/admin/players/${playerId}`, {
      baseURL: config.public.apiBase,
      method: 'DELETE',
      headers: {
        Authorization: `Bearer ${authStore.accessToken}`,
      },
    })

    if (response.success) {
      toast.success('Player Deleted', 'The player has been successfully deleted')
      router.push('/admin/players')
    } else {
      toast.error('Error', response.message || 'Failed to delete player')
    }
  } catch (error: unknown) {
    const err = error as { data?: { message?: string } }
    toast.error('Error', err.data?.message || 'An unexpected error occurred')
  } finally {
    deleting.value = false
    showDeleteModal.value = false
  }
}

// Handle profile photo upload
async function handlePhotoUpload(event: Event) {
  const input = event.target as HTMLInputElement
  const file = input.files?.[0]
  if (!file) return

  // Validate file type
  const allowedTypes = ['image/jpeg', 'image/png', 'image/webp']
  if (!allowedTypes.includes(file.type)) {
    toast.error('Invalid File', 'Please upload a JPEG, PNG, or WebP image')
    return
  }

  // Validate file size (5MB max)
  if (file.size > 5 * 1024 * 1024) {
    toast.error('File Too Large', 'Maximum file size is 5MB')
    return
  }

  uploadingPhoto.value = true

  try {
    // Step 1: Get presigned URL
    const initResponse = await $fetch<ApiResponse<{ upload_url: string; s3_key: string; session_id: string }>>('/admin/upload/init', {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: {
        Authorization: `Bearer ${authStore.accessToken}`,
      },
      body: {
        upload_type: 'profile_photo',
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
    await fetch(upload_url, {
      method: 'PUT',
      headers: {
        'Content-Type': file.type,
      },
      body: file,
    })

    // Step 3: Update player with new photo URL
    const s3Url = `s3://unicorn-sport-media/${s3_key}`
    const updateResponse = await $fetch<ApiResponse<{ player: Player }>>(`/admin/players/${playerId}`, {
      baseURL: config.public.apiBase,
      method: 'PUT',
      headers: {
        Authorization: `Bearer ${authStore.accessToken}`,
      },
      body: {
        profile_photo_url: s3Url,
      },
    })

    if (updateResponse.success && updateResponse.data?.player) {
      // Generate a presigned URL for display (or use the direct URL if available)
      profilePhotoUrl.value = updateResponse.data.player.profile_photo_url || null
      toast.success('Photo Uploaded', 'Profile photo has been updated successfully')
    }
  } catch (error) {
    console.error('Photo upload failed:', error)
    toast.error('Upload Failed', 'Failed to upload profile photo')
  } finally {
    uploadingPhoto.value = false
    // Reset the input
    input.value = ''
  }
}
</script>
