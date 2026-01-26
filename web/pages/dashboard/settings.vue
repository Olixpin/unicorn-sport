<template>
  <div>
    <!-- Header -->
    <div class="mb-8">
      <h1 class="font-display text-2xl lg:text-3xl font-bold text-neutral-900">
        Account Settings
      </h1>
      <p class="mt-1 text-neutral-500">Manage your account preferences and subscription</p>
    </div>

    <div class="space-y-6">
      <!-- Profile Section -->
      <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
        <div class="px-6 py-4 border-b border-neutral-100">
          <h2 class="font-semibold text-neutral-900">Profile Information</h2>
        </div>

        <form @submit.prevent="updateProfile" class="p-6 space-y-5">
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-5">
            <div>
              <label class="block text-sm font-medium text-neutral-700 mb-2">First Name</label>
              <input
                v-model="profile.firstName"
                type="text"
                placeholder="John"
                class="w-full px-4 py-3 rounded-xl border border-neutral-200 focus:border-primary-500 focus:ring-4 focus:ring-primary-500/10 transition-all duration-200 outline-none"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-neutral-700 mb-2">Last Name</label>
              <input
                v-model="profile.lastName"
                type="text"
                placeholder="Doe"
                class="w-full px-4 py-3 rounded-xl border border-neutral-200 focus:border-primary-500 focus:ring-4 focus:ring-primary-500/10 transition-all duration-200 outline-none"
              />
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-neutral-700 mb-2">Email Address</label>
            <input
              v-model="profile.email"
              type="email"
              disabled
              class="w-full px-4 py-3 rounded-xl border border-neutral-200 bg-neutral-50 text-neutral-500 cursor-not-allowed"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-neutral-700 mb-2">Organization</label>
            <input
              v-model="profile.organization"
              type="text"
              placeholder="Your club or agency"
              class="w-full px-4 py-3 rounded-xl border border-neutral-200 focus:border-primary-500 focus:ring-4 focus:ring-primary-500/10 transition-all duration-200 outline-none"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-neutral-700 mb-2">Role</label>
            <input
              v-model="profile.role"
              type="text"
              placeholder="Scout, Coach, Agent, etc."
              class="w-full px-4 py-3 rounded-xl border border-neutral-200 focus:border-primary-500 focus:ring-4 focus:ring-primary-500/10 transition-all duration-200 outline-none"
            />
          </div>

          <div class="flex justify-end pt-2">
            <button 
              type="submit" 
              :disabled="saving"
              class="px-6 py-3 bg-primary-500 text-white font-semibold rounded-xl hover:bg-primary-600 transition-colors disabled:opacity-50 disabled:cursor-not-allowed flex items-center gap-2"
            >
              <svg v-if="saving" class="w-5 h-5 animate-spin" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
              </svg>
              Save Changes
            </button>
          </div>
        </form>
      </div>

      <!-- Subscription Section -->
      <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
        <div class="px-6 py-4 border-b border-neutral-100 flex items-center justify-between">
          <h2 class="font-semibold text-neutral-900">Subscription</h2>
          <span 
            :class="[
              'px-3 py-1 text-xs font-bold rounded-full',
              subscriptionStore.tier === 'club' ? 'bg-primary-100 text-primary-700' : '',
              subscriptionStore.tier === 'pro' ? 'bg-amber-100 text-amber-700' : '',
              subscriptionStore.tier === 'scout' ? 'bg-blue-100 text-blue-700' : '',
              subscriptionStore.tier === 'free' ? 'bg-neutral-100 text-neutral-600' : ''
            ]"
          >
            {{ subscriptionStore.tier?.toUpperCase() || 'FREE' }}
          </span>
        </div>

        <div class="p-6 space-y-0 divide-y divide-neutral-100">
          <div class="flex items-center justify-between py-4 first:pt-0">
            <div>
              <p class="font-medium text-neutral-900">Current Plan</p>
              <p class="text-sm text-neutral-500 mt-0.5">{{ getPlanDescription() }}</p>
            </div>
            <NuxtLink to="/pricing">
              <button class="px-5 py-2.5 border-2 border-neutral-200 text-neutral-700 font-semibold rounded-xl hover:bg-neutral-50 hover:border-neutral-300 transition-all duration-200 text-sm">
                {{ subscriptionStore.tier === 'free' ? 'Upgrade' : 'Change Plan' }}
              </button>
            </NuxtLink>
          </div>

          <div v-if="subscriptionStore.tier !== 'free'" class="flex items-center justify-between py-4">
            <div>
              <p class="font-medium text-neutral-900">Billing Period</p>
              <p class="text-sm text-neutral-500 mt-0.5">Next billing: {{ nextBillingDate }}</p>
            </div>
            <button 
              @click="manageBilling"
              class="text-sm font-medium text-primary-600 hover:text-primary-700"
            >
              Manage Billing →
            </button>
          </div>

          <div class="flex items-center justify-between py-4 last:pb-0">
            <div>
              <p class="font-medium text-neutral-900">Video Downloads</p>
              <p class="text-sm text-neutral-500 mt-0.5">
                {{ downloadsUsed }} / {{ downloadsLimit }} this month
              </p>
            </div>
            <div class="w-32">
              <div class="bg-neutral-200 rounded-full h-2 overflow-hidden">
                <div 
                  class="bg-primary-500 h-2 rounded-full transition-all duration-500" 
                  :style="{ width: `${Math.min((downloadsUsed / downloadsLimit) * 100, 100)}%` }"
                ></div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Notifications Section -->
      <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
        <div class="px-6 py-4 border-b border-neutral-100">
          <h2 class="font-semibold text-neutral-900">Notifications</h2>
        </div>

        <div class="p-6 space-y-0 divide-y divide-neutral-100">
          <div class="flex items-center justify-between py-4 first:pt-0">
            <div>
              <p class="font-medium text-neutral-900">New Player Alerts</p>
              <p class="text-sm text-neutral-500 mt-0.5">Get notified when new players match your criteria</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input 
                type="checkbox" 
                v-model="notifications.newPlayers" 
                class="sr-only peer"
              />
              <div class="w-12 h-7 bg-neutral-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-primary-500/20 rounded-full peer peer-checked:after:translate-x-5 peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-neutral-300 after:border after:rounded-full after:h-6 after:w-6 after:transition-all after:shadow-sm peer-checked:bg-primary-500"></div>
            </label>
          </div>

          <div class="flex items-center justify-between py-4">
            <div>
              <p class="font-medium text-neutral-900">Contact Request Updates</p>
              <p class="text-sm text-neutral-500 mt-0.5">Updates on your contact requests</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input 
                type="checkbox" 
                v-model="notifications.contactUpdates" 
                class="sr-only peer"
              />
              <div class="w-12 h-7 bg-neutral-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-primary-500/20 rounded-full peer peer-checked:after:translate-x-5 peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-neutral-300 after:border after:rounded-full after:h-6 after:w-6 after:transition-all after:shadow-sm peer-checked:bg-primary-500"></div>
            </label>
          </div>

          <div class="flex items-center justify-between py-4 last:pb-0">
            <div>
              <p class="font-medium text-neutral-900">Weekly Digest</p>
              <p class="text-sm text-neutral-500 mt-0.5">Weekly summary of new players and activity</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input 
                type="checkbox" 
                v-model="notifications.weeklyDigest" 
                class="sr-only peer"
              />
              <div class="w-12 h-7 bg-neutral-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-primary-500/20 rounded-full peer peer-checked:after:translate-x-5 peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-neutral-300 after:border after:rounded-full after:h-6 after:w-6 after:transition-all after:shadow-sm peer-checked:bg-primary-500"></div>
            </label>
          </div>
        </div>

        <div class="px-6 pb-6 flex justify-end">
          <button 
            @click="saveNotifications" 
            :disabled="savingNotifications"
            class="px-5 py-2.5 border-2 border-neutral-200 text-neutral-700 font-semibold rounded-xl hover:bg-neutral-50 hover:border-neutral-300 transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed flex items-center gap-2 text-sm"
          >
            <svg v-if="savingNotifications" class="w-4 h-4 animate-spin" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
            </svg>
            Save Preferences
          </button>
        </div>
      </div>

      <!-- Password Section -->
      <div class="bg-white rounded-2xl border border-neutral-200 shadow-sm overflow-hidden">
        <div class="px-6 py-4 border-b border-neutral-100">
          <h2 class="font-semibold text-neutral-900">Change Password</h2>
        </div>

        <form @submit.prevent="changePassword" class="p-6 space-y-5">
          <div>
            <label class="block text-sm font-medium text-neutral-700 mb-2">Current Password</label>
            <input
              v-model="passwordForm.current"
              type="password"
              placeholder="Enter current password"
              class="w-full px-4 py-3 rounded-xl border border-neutral-200 focus:border-primary-500 focus:ring-4 focus:ring-primary-500/10 transition-all duration-200 outline-none"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-neutral-700 mb-2">New Password</label>
            <input
              v-model="passwordForm.new"
              type="password"
              placeholder="Enter new password"
              class="w-full px-4 py-3 rounded-xl border border-neutral-200 focus:border-primary-500 focus:ring-4 focus:ring-primary-500/10 transition-all duration-200 outline-none"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-neutral-700 mb-2">Confirm New Password</label>
            <input
              v-model="passwordForm.confirm"
              type="password"
              placeholder="Confirm new password"
              :class="[
                'w-full px-4 py-3 rounded-xl border transition-all duration-200 outline-none',
                passwordError ? 'border-red-300 focus:border-red-500 focus:ring-4 focus:ring-red-500/10' : 'border-neutral-200 focus:border-primary-500 focus:ring-4 focus:ring-primary-500/10'
              ]"
            />
            <p v-if="passwordError" class="mt-2 text-sm text-red-600">{{ passwordError }}</p>
          </div>

          <div class="flex justify-end pt-2">
            <button 
              type="submit" 
              :disabled="changingPassword"
              class="px-6 py-3 bg-primary-500 text-white font-semibold rounded-xl hover:bg-primary-600 transition-colors disabled:opacity-50 disabled:cursor-not-allowed flex items-center gap-2"
            >
              <svg v-if="changingPassword" class="w-5 h-5 animate-spin" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
              </svg>
              Update Password
            </button>
          </div>
        </form>
      </div>

      <!-- Danger Zone -->
      <div class="bg-white rounded-2xl border border-red-200 shadow-sm overflow-hidden">
        <div class="px-6 py-4 border-b border-red-100 bg-red-50">
          <h2 class="font-semibold text-red-700">Danger Zone</h2>
        </div>

        <div class="p-6 flex items-center justify-between">
          <div>
            <p class="font-medium text-neutral-900">Delete Account</p>
            <p class="text-sm text-neutral-500 mt-0.5">Permanently delete your account and all data</p>
          </div>
          <button 
            @click="showDeleteModal = true"
            class="px-5 py-2.5 border-2 border-red-200 text-red-600 font-semibold rounded-xl hover:bg-red-50 hover:border-red-300 transition-all duration-200 text-sm"
          >
            Delete Account
          </button>
        </div>
      </div>
    </div>

    <!-- Delete Account Modal -->
    <Teleport to="body">
      <div v-if="showDeleteModal" class="fixed inset-0 z-50 overflow-y-auto">
        <div class="flex min-h-full items-center justify-center p-4">
          <div class="fixed inset-0 bg-black/50 backdrop-blur-sm" @click="showDeleteModal = false"></div>
          
          <div class="relative bg-white rounded-2xl shadow-2xl max-w-md w-full p-6 animate-scale-in">
            <div class="w-14 h-14 mx-auto bg-red-100 rounded-2xl flex items-center justify-center mb-5">
              <svg class="w-7 h-7 text-red-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
              </svg>
            </div>
            <h3 class="text-xl font-bold text-neutral-900 text-center mb-2">Delete Account</h3>
            <p class="text-neutral-600 text-center mb-6">
              Are you sure? This action cannot be undone and all your data will be permanently deleted.
            </p>
            
            <div class="mb-6">
              <label class="block text-sm font-medium text-neutral-700 mb-2">Type 'DELETE' to confirm</label>
              <input
                v-model="deleteConfirmation"
                type="text"
                placeholder="DELETE"
                class="w-full px-4 py-3 rounded-xl border border-neutral-200 focus:border-red-500 focus:ring-4 focus:ring-red-500/10 transition-all duration-200 outline-none"
              />
            </div>

            <div class="flex gap-3">
              <button 
                @click="showDeleteModal = false"
                class="flex-1 py-3 border-2 border-neutral-200 text-neutral-700 font-semibold rounded-xl hover:bg-neutral-50 transition-colors"
              >
                Cancel
              </button>
              <button 
                @click="deleteAccount"
                :disabled="deleteConfirmation !== 'DELETE' || deleting"
                class="flex-1 py-3 bg-red-600 text-white font-semibold rounded-xl hover:bg-red-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center gap-2"
              >
                <svg v-if="deleting" class="w-5 h-5 animate-spin" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
                </svg>
                Delete
              </button>
            </div>
          </div>
        </div>
      </div>
    </Teleport>
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  layout: 'dashboard',
  middleware: 'auth',
})

const authStore = useAuthStore()
const subscriptionStore = useSubscriptionStore()
const api = useApi()
const router = useRouter()

const profile = reactive({
  firstName: '',
  lastName: '',
  email: authStore.user?.email || '',
  organization: '',
  role: '',
})

const notifications = reactive({
  newPlayers: true,
  contactUpdates: true,
  weeklyDigest: false,
})

const passwordForm = reactive({
  current: '',
  new: '',
  confirm: '',
})

const saving = ref(false)
const savingNotifications = ref(false)
const changingPassword = ref(false)
const showDeleteModal = ref(false)
const deleteConfirmation = ref('')
const deleting = ref(false)
const passwordError = ref('')

const downloadsUsed = ref(5)
const downloadsLimit = computed(() => {
  switch (subscriptionStore.tier) {
    case 'scout': return 10
    case 'pro': return 50
    case 'club': return 999
    default: return 0
  }
})

const nextBillingDate = computed(() => {
  const date = new Date()
  date.setMonth(date.getMonth() + 1)
  return date.toLocaleDateString('en-US', { month: 'long', day: 'numeric', year: 'numeric' })
})

const subscriptionBadgeVariant = computed(() => {
  switch (subscriptionStore.tier) {
    case 'club': return 'success'
    case 'pro': return 'warning'
    case 'scout': return 'default'
    default: return 'default'
  }
})

function getPlanDescription(): string {
  switch (subscriptionStore.tier) {
    case 'club': return 'Unlimited access to all players and features'
    case 'pro': return 'Full match videos and priority support'
    case 'scout': return 'Extended player information and contact requests'
    default: return 'Basic access to player highlights'
  }
}

async function updateProfile() {
  saving.value = true
  try {
    // API call to update profile
    await api.put('/profile', profile, true)
  } catch (error) {
    console.error('Failed to update profile:', error)
  } finally {
    saving.value = false
  }
}

async function saveNotifications() {
  savingNotifications.value = true
  try {
    await api.put('/profile/notifications', notifications, true)
  } catch (error) {
    console.error('Failed to save notifications:', error)
  } finally {
    savingNotifications.value = false
  }
}

async function changePassword() {
  if (passwordForm.new !== passwordForm.confirm) {
    passwordError.value = 'Passwords do not match'
    return
  }
  
  passwordError.value = ''
  changingPassword.value = true
  
  try {
    await api.put('/auth/password', {
      current_password: passwordForm.current,
      new_password: passwordForm.new,
    }, true)
    
    passwordForm.current = ''
    passwordForm.new = ''
    passwordForm.confirm = ''
  } catch (error) {
    console.error('Failed to change password:', error)
  } finally {
    changingPassword.value = false
  }
}

function manageBilling() {
  // Redirect to Stripe customer portal
  window.open('https://billing.stripe.com/p/login/xxx', '_blank')
}

async function deleteAccount() {
  deleting.value = true
  try {
    await api.delete('/account', true)
    authStore.logout()
    router.push('/')
  } catch (error) {
    console.error('Failed to delete account:', error)
  } finally {
    deleting.value = false
    showDeleteModal.value = false
  }
}
</script>
