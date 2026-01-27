<template>
  <Teleport to="body">
    <Transition
      enter-active-class="transition-opacity duration-200 ease-out"
      enter-from-class="opacity-0"
      enter-to-class="opacity-100"
      leave-active-class="transition-opacity duration-150 ease-in"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
    >
      <div
        v-if="confirmState.isOpen"
        class="fixed inset-0 z-[200] overflow-y-auto"
        role="alertdialog"
        aria-modal="true"
        aria-labelledby="confirm-title"
        aria-describedby="confirm-message"
      >
        <!-- Backdrop -->
        <div
          class="fixed inset-0 bg-black/60 backdrop-blur-sm"
          @click="handleCancel"
        />

        <!-- Dialog -->
        <div class="flex min-h-full items-center justify-center p-4">
          <Transition
            enter-active-class="transition-all duration-300 ease-out"
            enter-from-class="opacity-0 scale-95 translate-y-4"
            enter-to-class="opacity-100 scale-100 translate-y-0"
            leave-active-class="transition-all duration-200 ease-in"
            leave-from-class="opacity-100 scale-100 translate-y-0"
            leave-to-class="opacity-0 scale-95 translate-y-4"
          >
            <div
              v-if="confirmState.isOpen"
              class="relative w-full max-w-md transform overflow-hidden rounded-2xl bg-white p-6 shadow-2xl"
              @click.stop
            >
              <!-- Icon -->
              <div class="flex justify-center mb-4">
                <div
                  class="w-14 h-14 rounded-full flex items-center justify-center"
                  :class="iconBgClass"
                >
                  <!-- Trash Icon -->
                  <svg v-if="currentOptions.icon === 'trash'" class="w-7 h-7" :class="iconClass" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                  </svg>
                  <!-- Danger Icon -->
                  <svg v-else-if="currentOptions.icon === 'danger'" class="w-7 h-7" :class="iconClass" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                  </svg>
                  <!-- Warning Icon -->
                  <svg v-else-if="currentOptions.icon === 'warning'" class="w-7 h-7" :class="iconClass" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                  </svg>
                  <!-- Info Icon -->
                  <svg v-else-if="currentOptions.icon === 'info'" class="w-7 h-7" :class="iconClass" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                  <!-- Question Icon (default) -->
                  <svg v-else class="w-7 h-7" :class="iconClass" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>
              </div>

              <!-- Title -->
              <h3
                id="confirm-title"
                class="text-lg font-semibold text-neutral-900 text-center mb-2"
              >
                {{ currentOptions.title }}
              </h3>

              <!-- Message -->
              <p
                id="confirm-message"
                class="text-sm text-neutral-600 text-center mb-6"
              >
                {{ currentOptions.message }}
              </p>

              <!-- Actions -->
              <div class="flex gap-3">
                <button
                  type="button"
                  class="flex-1 px-4 py-2.5 text-sm font-medium text-neutral-700 bg-neutral-100 rounded-xl hover:bg-neutral-200 transition-colors focus:outline-none focus:ring-2 focus:ring-neutral-300"
                  @click="handleCancel"
                >
                  {{ currentOptions.cancelText || 'Cancel' }}
                </button>
                <button
                  type="button"
                  class="flex-1 px-4 py-2.5 text-sm font-medium text-white rounded-xl transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2"
                  :class="confirmButtonClass"
                  @click="handleConfirm"
                >
                  {{ currentOptions.confirmText || 'Confirm' }}
                </button>
              </div>
            </div>
          </Transition>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useConfirm } from '~/composables/useModal'

const { state: confirmState, handleConfirm, handleCancel } = useConfirm()

// Get current options or defaults
const currentOptions = computed(() => confirmState.options || {
  title: '',
  message: '',
  confirmText: 'Confirm',
  cancelText: 'Cancel',
  confirmVariant: 'danger',
  icon: 'warning',
})

const iconBgClass = computed(() => {
  switch (currentOptions.value.confirmVariant) {
    case 'danger':
      return 'bg-rose-100'
    case 'warning':
      return 'bg-amber-100'
    default:
      return 'bg-rose-100'
  }
})

const iconClass = computed(() => {
  switch (currentOptions.value.confirmVariant) {
    case 'danger':
      return 'text-rose-600'
    case 'warning':
      return 'text-amber-600'
    default:
      return 'text-rose-600'
  }
})

const confirmButtonClass = computed(() => {
  switch (currentOptions.value.confirmVariant) {
    case 'danger':
      return 'bg-rose-500 hover:bg-rose-600 focus:ring-rose-500'
    case 'warning':
      return 'bg-amber-500 hover:bg-amber-600 focus:ring-amber-500'
    default:
      return 'bg-primary-500 hover:bg-primary-600 focus:ring-primary-500'
  }
})
</script>
