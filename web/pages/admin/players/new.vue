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
              Add New Player
            </h1>
            <p class="mt-1 text-neutral-600">Create a new player profile in the system.</p>
          </div>
        </div>
        
        <!-- Progress Indicator -->
        <div class="flex items-center gap-2 px-4 py-2 bg-neutral-100 rounded-xl">
          <div class="flex items-center gap-1.5">
            <div :class="[
              'w-2.5 h-2.5 rounded-full transition-colors',
              form.first_name && form.last_name ? 'bg-emerald-500' : 'bg-neutral-300'
            ]"></div>
            <span class="text-xs font-medium text-neutral-600">Identity</span>
          </div>
          <div class="w-4 h-px bg-neutral-300"></div>
          <div class="flex items-center gap-1.5">
            <div :class="[
              'w-2.5 h-2.5 rounded-full transition-colors',
              form.position && form.date_of_birth ? 'bg-emerald-500' : 'bg-neutral-300'
            ]"></div>
            <span class="text-xs font-medium text-neutral-600">Details</span>
          </div>
          <div class="w-4 h-px bg-neutral-300"></div>
          <div class="flex items-center gap-1.5">
            <div :class="[
              'w-2.5 h-2.5 rounded-full transition-colors',
              form.country ? 'bg-emerald-500' : 'bg-neutral-300'
            ]"></div>
            <span class="text-xs font-medium text-neutral-600">Location</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Form -->
    <form @submit.prevent="handleSubmit">
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- Main Content -->
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
            </div>
          </div>

          <!-- Academy Assignment -->
          <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
            <div class="px-6 py-4 border-b border-neutral-200 bg-gradient-to-r from-amber-50 to-orange-50">
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-gradient-to-br from-amber-500 to-orange-600 rounded-xl flex items-center justify-center text-white shadow-lg shadow-amber-500/25">
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
                  </svg>
                </div>
                <div>
                  <h2 class="font-semibold text-neutral-900">Academy Assignment</h2>
                  <p class="text-sm text-neutral-500">Optionally assign to a partner academy</p>
                </div>
              </div>
            </div>

            <div class="p-6">
              <div class="relative">
                <select
                  v-model="form.academy_id"
                  :disabled="loadingAcademies || !form.country"
                  :class="[
                    'appearance-none w-full px-4 py-3 pr-10 border rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all',
                    loadingAcademies || !form.country ? 'bg-neutral-100 cursor-not-allowed' : 'bg-neutral-50'
                  ]"
                >
                  <option value="" disabled v-if="loadingAcademies">Loading academies...</option>
                  <option value="" v-else-if="!form.country">Select a country first</option>
                  <option value="" v-else-if="filteredAcademies.length === 0">No academies in {{ form.country }}</option>
                  <option value="">No academy assigned</option>
                  <option v-for="academy in filteredAcademies" :key="academy.id" :value="academy.id">
                    {{ academy.name }}
                  </option>
                </select>
                <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                  <svg v-if="loadingAcademies" class="w-5 h-5 text-neutral-400 animate-spin" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                  <svg v-else class="w-5 h-5 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                  </svg>
                </div>
              </div>
              
              <!-- Create new academy link -->
              <div class="mt-3 flex items-center justify-between">
                <p class="text-xs text-neutral-500 flex items-center gap-1.5">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                  Links player to an academy for tracking.
                </p>
                <NuxtLink 
                  to="/admin/academies/new" 
                  class="text-xs font-medium text-primary-600 hover:text-primary-700 flex items-center gap-1 hover:underline"
                >
                  <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                  </svg>
                  Create new academy
                </NuxtLink>
              </div>
            </div>
          </div>

          <!-- Status & Save Section -->
          <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
            <div class="px-6 py-4 border-b border-neutral-200 bg-gradient-to-r from-emerald-50 to-green-50">
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-gradient-to-br from-emerald-500 to-green-600 rounded-xl flex items-center justify-center text-white shadow-lg shadow-emerald-500/25">
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>
                <div>
                  <h2 class="font-semibold text-neutral-900">Status & Save</h2>
                  <p class="text-sm text-neutral-500">Set verification status and create player</p>
                </div>
              </div>
            </div>

            <div class="p-6 space-y-6">
              <!-- Verification Toggle -->
              <label class="flex items-center gap-4 cursor-pointer p-4 bg-neutral-50 rounded-xl border border-neutral-200 hover:border-primary-300 transition-colors">
                <div class="relative">
                  <input
                    v-model="form.is_verified"
                    type="checkbox"
                    class="sr-only peer"
                  />
                  <div class="w-12 h-7 bg-neutral-300 rounded-full peer-checked:bg-emerald-500 transition-colors"></div>
                  <div class="absolute left-1 top-1 w-5 h-5 bg-white rounded-full shadow peer-checked:translate-x-5 transition-transform"></div>
                </div>
                <div class="flex-1">
                  <span class="text-sm font-semibold text-neutral-900">Mark as Verified</span>
                  <p class="text-xs text-neutral-500 mt-0.5">Pre-approve this player for immediate visibility to scouts</p>
                </div>
                <div class="w-8 h-8 rounded-full flex items-center justify-center" :class="form.is_verified ? 'bg-emerald-100 text-emerald-600' : 'bg-neutral-200 text-neutral-400'">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                  </svg>
                </div>
              </label>

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
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                  </svg>
                  {{ submitting ? 'Creating Player...' : 'Create Player' }}
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

        <!-- Sidebar - Hidden on mobile, shown on desktop -->
        <div class="hidden lg:block space-y-6 lg:sticky lg:top-6 lg:self-start">
          <!-- Player Preview Card -->
          <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
            <div class="px-6 py-4 border-b border-neutral-200">
              <h3 class="font-semibold text-neutral-900">Player Preview</h3>
            </div>
            <div class="p-6 flex flex-col items-center">
              <!-- Avatar with initials -->
              <div class="w-20 h-20 bg-gradient-to-br from-primary-500 to-emerald-600 rounded-2xl flex items-center justify-center text-white text-2xl font-bold shadow-lg shadow-primary-500/25">
                {{ getInitials(form.first_name, form.last_name) }}
              </div>
              <h4 class="mt-4 font-semibold text-neutral-900 text-center">
                {{ form.first_name || 'First' }} {{ form.last_name || 'Last' }}
              </h4>
              <p class="text-sm text-neutral-500">{{ form.position || 'Position' }}</p>
              
              <!-- Stats Preview -->
              <div class="mt-4 w-full grid grid-cols-3 gap-3">
                <div class="text-center p-2 bg-neutral-50 rounded-lg">
                  <p class="text-xs text-neutral-500">Age</p>
                  <p class="font-semibold text-neutral-900">{{ playerAge || '—' }}</p>
                </div>
                <div class="text-center p-2 bg-neutral-50 rounded-lg">
                  <p class="text-xs text-neutral-500">Height</p>
                  <p class="font-semibold text-neutral-900">{{ form.height_cm || '—' }}</p>
                </div>
                <div class="text-center p-2 bg-neutral-50 rounded-lg">
                  <p class="text-xs text-neutral-500">Weight</p>
                  <p class="font-semibold text-neutral-900">{{ form.weight_kg || '—' }}</p>
                </div>
              </div>

              <!-- Location Preview -->
              <div class="mt-4 flex items-center gap-2 text-xs text-neutral-400">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                </svg>
                <span>{{ form.city || 'City' }}, {{ form.country || 'Country' }}</span>
              </div>

              <!-- Foot indicator -->
              <div v-if="form.preferred_foot" class="mt-3 flex items-center gap-1.5 px-3 py-1 bg-neutral-100 rounded-full">
                <svg class="w-4 h-4 text-neutral-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 17l-4 4m0 0l-4-4m4 4V3" />
                </svg>
                <span class="text-xs font-medium text-neutral-600 capitalize">{{ form.preferred_foot }} Foot</span>
              </div>
            </div>
          </div>

          <!-- Tips Card -->
          <div class="bg-gradient-to-br from-violet-50 to-purple-50 rounded-2xl border border-violet-100 p-5">
            <div class="flex items-start gap-3">
              <div class="w-8 h-8 bg-violet-100 rounded-lg flex items-center justify-center flex-shrink-0">
                <svg class="w-4 h-4 text-violet-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z" />
                </svg>
              </div>
              <div>
                <h4 class="font-medium text-violet-900 mb-1">Tips for Better Profiles</h4>
                <ul class="text-xs text-violet-700 space-y-1.5">
                  <li class="flex items-start gap-1.5">
                    <span class="text-violet-400 mt-0.5">•</span>
                    Complete all required fields for visibility
                  </li>
                  <li class="flex items-start gap-1.5">
                    <span class="text-violet-400 mt-0.5">•</span>
                    Accurate height and weight attract scouts
                  </li>
                  <li class="flex items-start gap-1.5">
                    <span class="text-violet-400 mt-0.5">•</span>
                    Verify players after document review
                  </li>
                </ul>
              </div>
            </div>
          </div>

          <!-- Keyboard Shortcuts -->
          <div class="text-center text-xs text-neutral-400 space-y-1">
            <p><kbd class="px-1.5 py-0.5 bg-neutral-100 rounded text-neutral-500">⌘</kbd> + <kbd class="px-1.5 py-0.5 bg-neutral-100 rounded text-neutral-500">Enter</kbd> to submit</p>
            <p><kbd class="px-1.5 py-0.5 bg-neutral-100 rounded text-neutral-500">Esc</kbd> to cancel</p>
          </div>
        </div>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { playerSchema, PLAYER_POSITIONS } from '~/schemas/player'
import { AFRICAN_COUNTRIES, getStatesForCountry, getCitiesForState, hasLocationData } from '~/data/locations'
import type { ApiResponse } from '~/types/index'

interface Academy {
  id: string
  name: string
  country: string
}

definePageMeta({
  layout: 'admin',
  middleware: 'admin',
})

const router = useRouter()
const toast = useToast()
const config = useRuntimeConfig()

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
  academy_id: '',
  is_verified: false,
})

const errors = reactive<Record<string, string>>({})
const submitting = ref(false)
const loadingAcademies = ref(true)
const academies = ref<Academy[]>([])
const allAcademies = ref<Academy[]>([])

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

// Watch country changes to reset state and city
watch(() => form.country, () => {
  form.state = ''
  form.city = ''
  form.academy_id = '' // Reset academy when country changes
})

// Watch state changes to reset city
watch(() => form.state, () => {
  form.city = ''
})

// Computed player age from DOB
const playerAge = computed(() => {
  if (!form.date_of_birth) return null
  const dob = new Date(form.date_of_birth)
  const today = new Date()
  let age = today.getFullYear() - dob.getFullYear()
  const monthDiff = today.getMonth() - dob.getMonth()
  if (monthDiff < 0 || (monthDiff === 0 && today.getDate() < dob.getDate())) {
    age--
  }
  return age
})

// Generate initials from name
function getInitials(firstName: string, lastName: string): string {
  const f = firstName?.charAt(0)?.toUpperCase() || ''
  const l = lastName?.charAt(0)?.toUpperCase() || ''
  return f + l || '??'
}

// Fetch academies for dropdown
onMounted(async () => {
  loadingAcademies.value = true
  try {
    const response = await $fetch<ApiResponse<{ academies: Academy[] }>>('/admin/academies', {
      baseURL: config.public.apiBase,
      headers: {
        Authorization: `Bearer ${useAuthStore().accessToken}`,
      },
    })
    if (response.success && response.data) {
      allAcademies.value = response.data.academies || []
      academies.value = allAcademies.value
    }
  } catch (error) {
    console.error('Failed to fetch academies:', error)
  } finally {
    loadingAcademies.value = false
  }

  // Add keyboard shortcuts
  window.addEventListener('keydown', handleKeyDown)
})

onUnmounted(() => {
  window.removeEventListener('keydown', handleKeyDown)
})

function handleKeyDown(e: KeyboardEvent) {
  // Cmd/Ctrl + Enter to submit
  if ((e.metaKey || e.ctrlKey) && e.key === 'Enter') {
    e.preventDefault()
    handleSubmit()
  }
  // Escape to cancel
  if (e.key === 'Escape') {
    router.push('/admin/players')
  }
}

function validateForm(): boolean {
  // Clear previous errors
  Object.keys(errors).forEach((key) => delete errors[key])

  // Build validation object, converting empty strings to undefined for optional fields
  const validationData = {
    first_name: form.first_name,
    last_name: form.last_name,
    date_of_birth: form.date_of_birth,
    position: form.position || undefined,
    preferred_foot: form.preferred_foot || undefined,
    height_cm: form.height_cm || undefined,
    weight_kg: form.weight_kg || undefined,
    country: form.country,
    state: form.state || undefined,
    city: form.city || undefined,
    academy_id: form.academy_id || undefined,
  }

  const result = playerSchema.safeParse(validationData)

  if (!result.success) {
    console.log('Validation errors:', result.error.errors)
    result.error.errors.forEach((err) => {
      const field = err.path[0] as string
      errors[field] = err.message
    })
    return false
  }

  return true
}

async function handleSubmit() {
  if (!validateForm()) {
    toast.error('Validation Error', 'Please fix the errors in the form')
    return
  }

  submitting.value = true

  try {
    const authStore = useAuthStore()
    
    // Build payload, converting empty strings to undefined
    const payload: Record<string, unknown> = {
      first_name: form.first_name,
      last_name: form.last_name,
      date_of_birth: form.date_of_birth,
      position: form.position,
      country: form.country,
      is_verified: form.is_verified,
    }
    
    // Only include optional fields if they have values
    if (form.preferred_foot) payload.preferred_foot = form.preferred_foot
    if (form.height_cm) payload.height_cm = form.height_cm
    if (form.weight_kg) payload.weight_kg = form.weight_kg
    if (form.state) payload.state = form.state
    if (form.city) payload.city = form.city
    if (form.academy_id) payload.academy_id = form.academy_id

    console.log('Submitting payload:', payload)

    const response = await $fetch<ApiResponse<{ player: { id: string } }>>('/admin/players', {
      baseURL: config.public.apiBase,
      method: 'POST',
      headers: {
        Authorization: `Bearer ${authStore.accessToken}`,
      },
      body: payload,
    })

    if (response.success) {
      toast.success('Player Created', 'The player has been successfully created')
      router.push('/admin/players')
    } else {
      toast.error('Error', response.message || 'Failed to create player')
    }
  } catch (error: unknown) {
    console.error('Submit error:', error)
    const err = error as { data?: { message?: string } }
    toast.error('Error', err.data?.message || 'An unexpected error occurred')
  } finally {
    submitting.value = false
  }
}
</script>
