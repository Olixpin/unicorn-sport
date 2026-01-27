<template>
  <div class="max-w-6xl mx-auto">
    <!-- Loading State -->
    <div v-if="loading" class="flex items-center justify-center min-h-[400px]">
      <div class="text-center">
        <svg class="w-12 h-12 animate-spin text-primary-600 mx-auto mb-4" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" />
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z" />
        </svg>
        <p class="text-neutral-600">Loading academy...</p>
      </div>
    </div>

    <template v-else>
      <!-- Page Header -->
      <div class="mb-8">
        <div class="flex items-center gap-4 mb-2">
          <NuxtLink 
            to="/admin/academies" 
            class="w-10 h-10 flex items-center justify-center rounded-xl bg-neutral-100 hover:bg-neutral-200 text-neutral-600 transition-colors"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
            </svg>
          </NuxtLink>
          <div class="flex-1">
            <h1 class="font-display text-2xl lg:text-3xl font-bold text-neutral-900">
              Edit Academy
            </h1>
            <p class="mt-1 text-neutral-600">Update academy information and settings.</p>
          </div>
          <button
            @click="showDeleteModal = true"
            class="hidden sm:flex items-center gap-2 px-4 py-2 text-red-600 hover:bg-red-50 rounded-xl transition-colors"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
            </svg>
            Delete
          </button>
        </div>
      </div>

      <form @submit.prevent="handleSubmit" class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- Main Content (2 columns) -->
        <div class="lg:col-span-2 space-y-6">
          <!-- Basic Information -->
          <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
            <div class="px-6 py-4 border-b border-neutral-200 bg-gradient-to-r from-violet-50 to-purple-50">
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-gradient-to-br from-violet-500 to-purple-600 rounded-xl flex items-center justify-center text-white shadow-lg shadow-violet-500/25">
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
                  </svg>
                </div>
                <div>
                  <h2 class="font-semibold text-neutral-900">Basic Information</h2>
                  <p class="text-sm text-neutral-500">Academy name and details</p>
                </div>
              </div>
            </div>

            <div class="p-6 space-y-5">
              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-2">Academy Name *</label>
                <input
                  v-model="form.name"
                  type="text"
                  placeholder="Lagos Football Academy"
                  :class="[
                    'w-full px-4 py-3 border rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all',
                    errors.name ? 'border-red-300 bg-red-50' : 'border-neutral-200 bg-neutral-50'
                  ]"
                />
                <p v-if="errors.name" class="mt-1.5 text-sm text-red-600">{{ errors.name }}</p>
              </div>

              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-2">Description</label>
                <textarea
                  v-model="form.description"
                  rows="4"
                  placeholder="Brief description of the academy, training philosophy, achievements..."
                  class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all resize-none"
                ></textarea>
              </div>

              <div class="grid grid-cols-1 sm:grid-cols-2 gap-5">
                <div>
                  <label class="block text-sm font-medium text-neutral-700 mb-2">Founded Year</label>
                  <input
                    v-model.number="form.founded_year"
                    type="number"
                    placeholder="2005"
                    min="1900"
                    :max="new Date().getFullYear()"
                    class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-neutral-700 mb-2">Logo URL</label>
                  <input
                    v-model="form.logo_url"
                    type="url"
                    placeholder="https://example.com/logo.png"
                    class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
                  />
                </div>
              </div>
            </div>
          </div>

          <!-- Contact Information -->
          <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
            <div class="px-6 py-4 border-b border-neutral-200 bg-gradient-to-r from-blue-50 to-indigo-50">
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-gradient-to-br from-blue-500 to-indigo-600 rounded-xl flex items-center justify-center text-white shadow-lg shadow-blue-500/25">
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                  </svg>
                </div>
                <div>
                  <h2 class="font-semibold text-neutral-900">Contact Information</h2>
                  <p class="text-sm text-neutral-500">How scouts can reach the academy</p>
                </div>
              </div>
            </div>

            <div class="p-6 space-y-5">
              <div class="grid grid-cols-1 sm:grid-cols-2 gap-5">
                <div>
                  <label class="block text-sm font-medium text-neutral-700 mb-2">Email</label>
                  <input
                    v-model="form.email"
                    type="email"
                    placeholder="info@academy.com"
                    :class="[
                      'w-full px-4 py-3 border rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all',
                      errors.email ? 'border-red-300 bg-red-50' : 'border-neutral-200 bg-neutral-50'
                    ]"
                  />
                  <p v-if="errors.email" class="mt-1.5 text-sm text-red-600">{{ errors.email }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-neutral-700 mb-2">Phone</label>
                  <input
                    v-model="form.phone"
                    type="tel"
                    placeholder="+234 800 000 0000"
                    class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
                  />
                </div>
              </div>

              <div>
                <label class="block text-sm font-medium text-neutral-700 mb-2">Website</label>
                <input
                  v-model="form.website"
                  type="url"
                  placeholder="https://www.academy.com"
                  :class="[
                    'w-full px-4 py-3 border rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all',
                    errors.website ? 'border-red-300 bg-red-50' : 'border-neutral-200 bg-neutral-50'
                  ]"
                />
                <p v-if="errors.website" class="mt-1.5 text-sm text-red-600">{{ errors.website }}</p>
              </div>
            </div>
          </div>

          <!-- Location -->
          <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
            <div class="px-6 py-4 border-b border-neutral-200 bg-gradient-to-r from-emerald-50 to-green-50">
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-gradient-to-br from-emerald-500 to-green-600 rounded-xl flex items-center justify-center text-white shadow-lg shadow-emerald-500/25">
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
                  </svg>
                </div>
                <div>
                  <h2 class="font-semibold text-neutral-900">Location</h2>
                  <p class="text-sm text-neutral-500">Where the academy is based</p>
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
                      :class="[
                        'appearance-none w-full px-4 py-3 pr-10 border rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all',
                        errors.country ? 'border-red-300 bg-red-50' : 'border-neutral-200 bg-neutral-50'
                      ]"
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
                  <p v-if="errors.country" class="mt-1.5 text-sm text-red-600">{{ errors.country }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-neutral-700 mb-2">State/Region</label>
                  <input
                    v-model="form.state"
                    type="text"
                    placeholder="Lagos State"
                    class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
                  />
                </div>
              </div>

              <div class="grid grid-cols-1 sm:grid-cols-2 gap-5">
                <div>
                  <label class="block text-sm font-medium text-neutral-700 mb-2">City</label>
                  <input
                    v-model="form.city"
                    type="text"
                    placeholder="Lagos"
                    class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-neutral-700 mb-2">Address</label>
                  <input
                    v-model="form.address"
                    type="text"
                    placeholder="123 Football Lane"
                    class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent transition-all"
                  />
                </div>
              </div>
            </div>
          </div>

          <!-- Status & Actions Section -->
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
                <div class="rounded-xl border-2 overflow-hidden" :class="form.is_verified ? 'border-emerald-200 bg-white' : 'border-neutral-200 bg-white'">
                  <!-- Status Header -->
                  <div class="px-4 py-3 flex items-center gap-3" :class="form.is_verified ? 'bg-emerald-50' : 'bg-neutral-50'">
                    <div class="w-8 h-8 rounded-full flex items-center justify-center" :class="form.is_verified ? 'bg-emerald-500' : 'bg-neutral-400'">
                      <svg v-if="form.is_verified" class="w-4 h-4 text-white" fill="currentColor" viewBox="0 0 20 20">
                        <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                      </svg>
                      <svg v-else class="w-4 h-4 text-white" fill="currentColor" viewBox="0 0 20 20">
                        <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z" clip-rule="evenodd" />
                      </svg>
                    </div>
                    <div class="flex-1">
                      <p class="text-sm font-semibold" :class="form.is_verified ? 'text-emerald-900' : 'text-neutral-900'">
                        {{ form.is_verified ? 'Verified' : 'Pending Review' }}
                      </p>
                    </div>
                    <span 
                      class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                      :class="form.is_verified ? 'bg-emerald-100 text-emerald-800' : 'bg-neutral-200 text-neutral-700'"
                    >
                      {{ form.is_verified ? 'Active' : 'Unverified' }}
                    </span>
                  </div>
                  
                  <!-- Status Body with Action -->
                  <div class="p-4 flex items-center justify-between gap-4">
                    <p class="text-sm text-neutral-600">
                      {{ form.is_verified 
                        ? 'Displays verification badge in search results and academy profile.' 
                        : 'Complete verification to display trust badge and improve visibility.' 
                      }}
                    </p>
                    <button
                      type="button"
                      @click="showVerifyModal = true"
                      class="shrink-0 inline-flex items-center gap-2 px-4 py-2 rounded-lg text-sm font-medium transition-colors"
                      :class="form.is_verified 
                        ? 'text-neutral-700 bg-neutral-100 hover:bg-neutral-200' 
                        : 'text-white bg-neutral-900 hover:bg-neutral-800'"
                    >
                      <svg v-if="!form.is_verified" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
                      </svg>
                      {{ form.is_verified ? 'Manage' : 'Verify Now' }}
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
                
                <NuxtLink to="/admin/academies" class="sm:w-auto">
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
          <!-- Logo Preview -->
          <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
            <div class="px-6 py-4 border-b border-neutral-200">
              <h3 class="font-semibold text-neutral-900">Logo Preview</h3>
            </div>
            <div class="p-6 flex flex-col items-center">
              <div class="w-24 h-24 rounded-xl bg-neutral-100 border-2 border-dashed border-neutral-300 flex items-center justify-center overflow-hidden">
                <img
                  v-if="form.logo_url"
                  :src="form.logo_url"
                  :alt="form.name || 'Academy Logo'"
                  class="w-full h-full object-cover"
                  @error="form.logo_url = ''"
                />
                <svg v-else class="w-10 h-10 text-neutral-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
              </div>
              <p class="mt-3 text-sm text-neutral-500 text-center">
                Enter a valid image URL to preview the academy logo
              </p>
            </div>
          </div>

          <!-- Quick Actions -->
          <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
            <div class="px-6 py-4 border-b border-neutral-200">
              <h3 class="font-semibold text-neutral-900">Quick Actions</h3>
            </div>
            <div class="p-4 space-y-2">
              <button
                type="button"
                @click="showDeleteModal = true"
                class="w-full flex items-center gap-3 px-4 py-3 text-left text-red-600 hover:bg-red-50 rounded-xl transition-colors"
              >
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                </svg>
                Delete Academy
              </button>
            </div>
          </div>

          <!-- Tips -->
          <div class="bg-gradient-to-br from-blue-50 to-indigo-50 rounded-2xl border border-blue-200 p-5">
            <div class="flex items-start gap-3">
              <div class="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center flex-shrink-0">
                <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </div>
              <div>
                <h4 class="text-sm font-semibold text-blue-900 mb-1">Editing Tips</h4>
                <ul class="text-xs text-blue-700 space-y-1">
                  <li>• Changes are saved immediately</li>
                  <li>• Verify contact info is current</li>
                  <li>• Keep logo URL accessible</li>
                </ul>
              </div>
            </div>
          </div>
        </div>
      </form>

      <!-- Mobile Fixed Action Bar -->
      <div class="fixed bottom-0 left-0 right-0 bg-white border-t border-neutral-200 p-4 shadow-lg z-50 lg:hidden">
        <div class="flex gap-3 max-w-md mx-auto">
          <NuxtLink to="/admin/academies" class="flex-1">
            <button
              type="button"
              class="w-full px-4 py-3 border border-neutral-300 text-neutral-700 rounded-xl text-sm font-medium hover:bg-neutral-50 transition-colors"
            >
              Cancel
            </button>
          </NuxtLink>
          <button
            type="button"
            @click="handleSubmit"
            :disabled="submitting"
            class="flex-1 inline-flex items-center justify-center gap-2 px-4 py-3 bg-gradient-to-r from-primary-600 to-emerald-600 text-white rounded-xl text-sm font-semibold hover:from-primary-700 hover:to-emerald-700 disabled:opacity-50 disabled:cursor-not-allowed transition-all shadow-lg shadow-primary-600/25"
          >
            <svg v-if="submitting" class="w-5 h-5 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" fill="none" />
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z" />
            </svg>
            {{ submitting ? 'Saving...' : 'Save Changes' }}
          </button>
        </div>
      </div>

      <!-- Spacer for mobile fixed bar -->
      <div class="lg:hidden h-20"></div>
    </template>

    <!-- Delete Confirmation Modal -->
    <Teleport to="body">
      <div v-if="showDeleteModal" class="fixed inset-0 z-50 flex items-center justify-center p-4">
        <div class="absolute inset-0 bg-black/50" @click="showDeleteModal = false"></div>
        <div class="relative bg-white rounded-2xl shadow-xl max-w-md w-full p-6">
          <div class="flex items-center gap-4 mb-4">
            <div class="w-12 h-12 bg-red-100 rounded-full flex items-center justify-center">
              <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
              </svg>
            </div>
            <div>
              <h3 class="text-lg font-semibold text-neutral-900">Delete Academy</h3>
              <p class="text-sm text-neutral-500">This action cannot be undone</p>
            </div>
          </div>
          <p class="text-neutral-600 mb-6">
            Are you sure you want to delete <strong>{{ form.name }}</strong>? All data associated with this academy will be permanently removed.
          </p>
          <div class="flex gap-3">
            <button
              @click="showDeleteModal = false"
              class="flex-1 px-4 py-3 border border-neutral-300 text-neutral-700 rounded-xl font-medium hover:bg-neutral-50 transition-colors"
            >
              Cancel
            </button>
            <button
              @click="handleDelete"
              :disabled="deleting"
              class="flex-1 px-4 py-3 bg-red-600 text-white rounded-xl font-medium hover:bg-red-700 disabled:opacity-50 transition-colors"
            >
              {{ deleting ? 'Deleting...' : 'Delete' }}
            </button>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- Verification Modal (handles both verify and revoke) -->
    <AdminAcademyVerificationModal
      v-model="showVerifyModal"
      :academy="academyForModal"
      :mode="form.is_verified ? 'revoke' : 'verify'"
      :loading="verifying"
      @confirm="handleVerificationConfirm"
      @close="showVerifyModal = false"
    />
  </div>
</template>

<script setup lang="ts">
import { academySchema, type AcademyFormData, type Academy } from '~/schemas/academy'
import { AFRICAN_COUNTRIES } from '~/schemas/player'
import type { ApiResponse } from '~/types/index'

definePageMeta({
  layout: 'admin',
  middleware: 'admin',
})

const route = useRoute()
const router = useRouter()
const toast = useToast()
const config = useRuntimeConfig()
const authStore = useAuthStore()

const academyId = route.params.id as string
const loading = ref(true)
const submitting = ref(false)
const deleting = ref(false)
const verifying = ref(false)
const showDeleteModal = ref(false)
const showVerifyModal = ref(false)

// Computed property to convert form to Academy object for the modal
const academyForModal = computed<Academy>(() => ({
  id: academyId,
  name: form.name || '',
  description: form.description,
  country: form.country || '',
  state: form.state,
  city: form.city,
  address: form.address,
  phone: form.phone,
  email: form.email,
  website: form.website,
  logo_url: form.logo_url,
  founded_year: form.founded_year,
  is_verified: form.is_verified,
  created_at: '',
  updated_at: '',
}))

const form = reactive<Partial<AcademyFormData> & { is_verified?: boolean }>({
  name: '',
  description: '',
  country: '',
  state: '',
  city: '',
  address: '',
  phone: '',
  email: '',
  website: '',
  founded_year: undefined,
  logo_url: '',
  is_verified: false,
})

const errors = reactive<Record<string, string>>({})

async function fetchAcademy() {
  try {
    const response = await $fetch<ApiResponse<Academy>>(`/admin/academies/${academyId}`, {
      baseURL: config.public.apiBase,
      headers: {
        Authorization: `Bearer ${authStore.accessToken}`,
      },
    })

    if (response.success && response.data) {
      const academy = response.data
      form.name = academy.name
      form.description = academy.description || ''
      form.country = academy.country
      form.state = academy.state || ''
      form.city = academy.city || ''
      form.address = academy.address || ''
      form.phone = academy.phone || ''
      form.email = academy.email || ''
      form.website = academy.website || ''
      form.logo_url = academy.logo_url || ''
      form.founded_year = academy.founded_year
      form.is_verified = academy.is_verified
    }
  } catch (error) {
    console.error('Failed to fetch academy:', error)
    toast.error('Error', 'Failed to load academy data')
    router.push('/admin/academies')
  } finally {
    loading.value = false
  }
}

function validateForm(): boolean {
  Object.keys(errors).forEach((key) => delete errors[key])

  const result = academySchema.safeParse(form)

  if (!result.success) {
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
    const response = await $fetch<ApiResponse<Academy>>(`/admin/academies/${academyId}`, {
      baseURL: config.public.apiBase,
      method: 'PUT',
      headers: {
        Authorization: `Bearer ${authStore.accessToken}`,
      },
      body: form,
    })

    if (response.success) {
      toast.success('Academy Updated', 'The academy has been successfully updated')
      router.push('/admin/academies')
    } else {
      toast.error('Error', response.message || 'Failed to update academy')
    }
  } catch (error: unknown) {
    const err = error as { data?: { message?: string } }
    toast.error('Error', err.data?.message || 'An unexpected error occurred')
  } finally {
    submitting.value = false
  }
}

async function handleDelete() {
  deleting.value = true

  try {
    const response = await $fetch<ApiResponse<null>>(`/admin/academies/${academyId}`, {
      baseURL: config.public.apiBase,
      method: 'DELETE',
      headers: {
        Authorization: `Bearer ${authStore.accessToken}`,
      },
    })

    if (response.success) {
      toast.success('Academy Deleted', 'The academy has been permanently deleted')
      router.push('/admin/academies')
    } else {
      toast.error('Error', response.message || 'Failed to delete academy')
    }
  } catch (error: unknown) {
    const err = error as { data?: { message?: string } }
    toast.error('Error', err.data?.message || 'An unexpected error occurred')
  } finally {
    deleting.value = false
    showDeleteModal.value = false
  }
}

async function handleVerificationConfirm(_data: { academy: Academy; reason?: string }) {
  verifying.value = true
  const isVerify = !form.is_verified

  try {
    const response = await $fetch<ApiResponse<Academy>>(`/admin/academies/${academyId}`, {
      baseURL: config.public.apiBase,
      method: 'PUT',
      headers: {
        Authorization: `Bearer ${authStore.accessToken}`,
      },
      body: {
        ...form,
        is_verified: isVerify,
      },
    })

    if (response.success) {
      form.is_verified = isVerify
      toast.success(
        isVerify ? 'Academy Verified' : 'Verification Revoked',
        isVerify
          ? 'The academy is now verified and will display a verification badge'
          : 'The academy verification has been revoked'
      )
    } else {
      toast.error('Error', response.message || `Failed to ${isVerify ? 'verify' : 'revoke'} academy`)
    }
  } catch (error: unknown) {
    const err = error as { data?: { message?: string } }
    toast.error('Error', err.data?.message || 'An unexpected error occurred')
  } finally {
    verifying.value = false
    showVerifyModal.value = false
  }
}

onMounted(() => {
  fetchAcademy()
})
</script>
