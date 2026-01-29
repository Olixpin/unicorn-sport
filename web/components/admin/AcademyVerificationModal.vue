<template>
  <UiModal
    v-model="isOpen"
    :size="'md'"
    @close="handleClose"
  >
    <template #header>
      <div class="flex items-center gap-4">
        <div 
          class="w-12 h-12 rounded-full flex items-center justify-center"
          :class="isVerifyMode ? 'bg-emerald-100' : 'bg-amber-100'"
        >
          <!-- Verify Icon -->
          <svg v-if="isVerifyMode" class="w-6 h-6 text-emerald-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z" />
          </svg>
          <!-- Revoke Icon -->
          <svg v-else class="w-6 h-6 text-amber-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
          </svg>
        </div>
        <div>
          <h3 class="text-lg font-bold text-neutral-900">
            {{ isVerifyMode ? 'Verify Academy' : 'Revoke Verification' }}
          </h3>
          <p class="text-sm text-neutral-500">
            {{ isVerifyMode ? 'Review before verifying' : 'This will remove the verified badge' }}
          </p>
        </div>
      </div>
    </template>

    <!-- Academy Info Card -->
    <AdminAcademyInfoCard v-if="academy" :academy="academy" class="mb-6" />

    <!-- Verify Mode Content -->
    <template v-if="isVerifyMode">
      <!-- Verification Checklist -->
      <div class="bg-neutral-50 rounded-xl p-4 mb-6 space-y-3">
        <p class="text-sm font-medium text-neutral-700 mb-3">By verifying this academy, you confirm:</p>
        <label class="flex items-start gap-3 cursor-pointer">
          <input 
            type="checkbox" 
            v-model="checks.legitimacy" 
            class="mt-0.5 w-4 h-4 text-emerald-600 border-neutral-300 rounded focus:ring-emerald-500" 
          />
          <span class="text-sm text-neutral-600">The academy is a legitimate, registered organization</span>
        </label>
        <label class="flex items-start gap-3 cursor-pointer">
          <input 
            type="checkbox" 
            v-model="checks.contact" 
            class="mt-0.5 w-4 h-4 text-emerald-600 border-neutral-300 rounded focus:ring-emerald-500" 
          />
          <span class="text-sm text-neutral-600">Contact information has been verified</span>
        </label>
        <label class="flex items-start gap-3 cursor-pointer">
          <input 
            type="checkbox" 
            v-model="checks.location" 
            class="mt-0.5 w-4 h-4 text-emerald-600 border-neutral-300 rounded focus:ring-emerald-500" 
          />
          <span class="text-sm text-neutral-600">Location and address are accurate</span>
        </label>
      </div>

      <!-- Impact Notice -->
      <div class="bg-emerald-50 border border-emerald-200 rounded-xl p-4 mb-6">
        <p class="text-sm text-emerald-800"><strong>What happens when verified:</strong></p>
        <ul class="text-sm text-emerald-700 mt-2 space-y-1">
          <li>• Academy displays a verified badge</li>
          <li>• Prioritized in search results</li>
          <li>• Increased visibility to scouts</li>
        </ul>
      </div>
    </template>

    <!-- Revoke Mode Content -->
    <template v-else>
      <!-- Warning Notice -->
      <div class="bg-amber-50 border border-amber-200 rounded-xl p-4 mb-6">
        <p class="text-sm text-amber-800"><strong>What happens when revoked:</strong></p>
        <ul class="text-sm text-amber-700 mt-2 space-y-1">
          <li>• Verified badge will be removed</li>
          <li>• Academy loses search priority</li>
          <li>• Action will be logged for audit</li>
        </ul>
      </div>

      <!-- Reason Required -->
      <div class="mb-6">
        <label class="block text-sm font-medium text-neutral-700 mb-2">Reason for Revocation *</label>
        <textarea
          v-model="reason"
          rows="3"
          placeholder="Explain why verification is being revoked..."
          class="w-full px-4 py-3 bg-neutral-50 border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-amber-500 focus:border-transparent resize-none"
        ></textarea>
      </div>
    </template>

    <!-- Action Buttons -->
    <template #footer>
      <div class="flex gap-3">
        <button
          @click="handleClose"
          class="flex-1 px-4 py-3 border border-neutral-300 text-neutral-700 rounded-xl font-medium hover:bg-neutral-50 transition-colors"
        >
          Cancel
        </button>
        <button
          @click="handleConfirm"
          :disabled="loading || !canConfirm"
          class="flex-1 px-4 py-3 rounded-xl font-medium transition-colors"
          :class="confirmButtonClasses"
        >
          {{ confirmButtonText }}
        </button>
      </div>
    </template>
  </UiModal>
</template>

<script setup lang="ts">
import type { Academy } from '~/schemas/academy'

interface Props {
  modelValue: boolean
  academy: Academy | null
  mode: 'verify' | 'revoke'
  loading?: boolean
}

interface Emits {
  (e: 'update:modelValue', value: boolean): void
  (e: 'confirm', data: { academy: Academy; reason?: string }): void
  (e: 'close'): void
}

const props = withDefaults(defineProps<Props>(), {
  loading: false,
})

const emit = defineEmits<Emits>()

const isOpen = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value),
})

const isVerifyMode = computed(() => props.mode === 'verify')

// Verification checklist state
const checks = reactive({
  legitimacy: false,
  contact: false,
  location: false,
})

// Revoke reason
const reason = ref('')

const allChecksComplete = computed(() => 
  checks.legitimacy && checks.contact && checks.location
)

const canConfirm = computed(() => {
  if (isVerifyMode.value) {
    return allChecksComplete.value
  }
  return reason.value.trim().length > 0
})

const confirmButtonClasses = computed(() => {
  if (props.loading || !canConfirm.value) {
    return 'bg-neutral-200 text-neutral-400 cursor-not-allowed'
  }
  return isVerifyMode.value 
    ? 'bg-emerald-600 text-white hover:bg-emerald-700'
    : 'bg-amber-600 text-white hover:bg-amber-700'
})

const confirmButtonText = computed(() => {
  if (props.loading) {
    return isVerifyMode.value ? 'Verifying...' : 'Revoking...'
  }
  return isVerifyMode.value ? 'Confirm Verification' : 'Revoke Verification'
})

function resetState() {
  checks.legitimacy = false
  checks.contact = false
  checks.location = false
  reason.value = ''
}

function handleClose() {
  resetState()
  emit('close')
  emit('update:modelValue', false)
}

function handleConfirm() {
  if (!props.academy || !canConfirm.value) return
  
  emit('confirm', {
    academy: props.academy,
    reason: isVerifyMode.value ? undefined : reason.value,
  })
}

// Reset state when modal opens with new academy
watch(() => props.modelValue, (newVal) => {
  if (newVal) {
    resetState()
  }
})
</script>
