<template>
  <Teleport to="body">
    <Transition
      enter-active-class="transition-opacity duration-200 ease-out"
      enter-from-class="opacity-0"
      enter-to-class="opacity-100"
      leave-active-class="transition-opacity duration-150 ease-in"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
      @after-leave="onAfterLeave"
    >
      <div
        v-if="modelValue"
        ref="modalRef"
        class="fixed inset-0 z-[100] overflow-y-auto"
        role="dialog"
        :aria-modal="true"
        :aria-labelledby="titleId"
        :aria-describedby="descriptionId"
      >
        <!-- Backdrop -->
        <div
          class="fixed inset-0 bg-black/60 backdrop-blur-sm"
          :class="backdropClass"
          @click="onBackdropClick"
          aria-hidden="true"
        />

        <!-- Modal Container - Centers the modal -->
        <div
          class="flex min-h-full items-center justify-center p-4"
          :class="containerClass"
        >
          <!-- Modal Panel -->
          <Transition
            enter-active-class="transition-all duration-300 ease-out"
            enter-from-class="opacity-0 scale-95 translate-y-4"
            enter-to-class="opacity-100 scale-100 translate-y-0"
            leave-active-class="transition-all duration-200 ease-in"
            leave-from-class="opacity-100 scale-100 translate-y-0"
            leave-to-class="opacity-0 scale-95 translate-y-4"
          >
            <div
              v-if="modelValue"
              ref="panelRef"
              class="relative w-full transform overflow-hidden rounded-2xl bg-white shadow-2xl transition-all"
              :class="[sizeClasses, panelClass]"
              @click.stop
            >
              <!-- Close Button (Optional) -->
              <button
                v-if="showCloseButton"
                type="button"
                class="absolute right-4 top-4 z-10 rounded-full p-2 text-neutral-400 transition-colors hover:bg-neutral-100 hover:text-neutral-600 focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2"
                :class="closeButtonClass"
                @click="close"
                :aria-label="closeAriaLabel"
              >
                <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>

              <!-- Header Slot -->
              <div v-if="$slots.header || title" class="border-b border-neutral-200 px-6 py-4" :class="headerClass">
                <slot name="header">
                  <h3 :id="titleId" class="text-lg font-semibold text-neutral-900">
                    {{ title }}
                  </h3>
                  <p v-if="description" :id="descriptionId" class="mt-1 text-sm text-neutral-500">
                    {{ description }}
                  </p>
                </slot>
              </div>

              <!-- Body Slot -->
              <div class="relative" :class="[bodyClass, { 'px-6 py-4': !noPadding }]">
                <slot />
              </div>

              <!-- Footer Slot -->
              <div v-if="$slots.footer" class="border-t border-neutral-200 px-6 py-4" :class="footerClass">
                <slot name="footer" :close="close" />
              </div>
            </div>
          </Transition>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted, watch, nextTick } from 'vue'

// =============================================================================
// TYPES
// =============================================================================

type ModalSize = 'xs' | 'sm' | 'md' | 'lg' | 'xl' | '2xl' | 'full' | 'video'

interface ModalProps {
  modelValue: boolean
  title?: string
  description?: string
  size?: ModalSize
  persistent?: boolean // Prevent closing on backdrop click
  showCloseButton?: boolean
  closeOnEscape?: boolean
  lockScroll?: boolean
  initialFocus?: HTMLElement | null
  closeAriaLabel?: string
  // Custom classes
  backdropClass?: string
  containerClass?: string
  panelClass?: string
  headerClass?: string
  bodyClass?: string
  footerClass?: string
  closeButtonClass?: string
  noPadding?: boolean
}

interface ModalEmits {
  (e: 'update:modelValue', value: boolean): void
  (e: 'close'): void
  (e: 'closed'): void // After animation completes
  (e: 'open'): void
}

// =============================================================================
// PROPS & EMITS
// =============================================================================

const props = withDefaults(defineProps<ModalProps>(), {
  size: 'md',
  persistent: false,
  showCloseButton: true,
  closeOnEscape: true,
  lockScroll: true,
  closeAriaLabel: 'Close modal',
  noPadding: false,
})

const emit = defineEmits<ModalEmits>()

// =============================================================================
// REFS
// =============================================================================

const modalRef = ref<HTMLElement | null>(null)
const panelRef = ref<HTMLElement | null>(null)
const previousActiveElement = ref<HTMLElement | null>(null)

// Generate unique IDs for accessibility
const titleId = `modal-title-${Math.random().toString(36).substr(2, 9)}`
const descriptionId = `modal-desc-${Math.random().toString(36).substr(2, 9)}`

// =============================================================================
// COMPUTED
// =============================================================================

const sizeClasses = computed(() => {
  const sizes: Record<ModalSize, string> = {
    xs: 'max-w-xs',
    sm: 'max-w-sm',
    md: 'max-w-md',
    lg: 'max-w-lg',
    xl: 'max-w-xl',
    '2xl': 'max-w-2xl',
    full: 'max-w-[95vw] max-h-[95vh]',
    video: 'max-w-4xl',
  }
  return sizes[props.size]
})

// =============================================================================
// METHODS
// =============================================================================

function close() {
  emit('update:modelValue', false)
  emit('close')
}

function onBackdropClick() {
  if (!props.persistent) {
    close()
  }
}

function onAfterLeave() {
  emit('closed')
  // Restore focus to the element that was focused before the modal opened
  if (previousActiveElement.value && typeof previousActiveElement.value.focus === 'function') {
    nextTick(() => {
      previousActiveElement.value?.focus()
    })
  }
}

function handleKeydown(event: KeyboardEvent) {
  if (event.key === 'Escape' && props.closeOnEscape && props.modelValue) {
    event.preventDefault()
    event.stopPropagation()
    close()
  }

  // Focus trap - Tab key handling
  if (event.key === 'Tab' && props.modelValue && panelRef.value) {
    const focusableElements = panelRef.value.querySelectorAll<HTMLElement>(
      'button, [href], input, select, textarea, [tabindex]:not([tabindex="-1"])'
    )
    const firstElement = focusableElements[0]
    const lastElement = focusableElements[focusableElements.length - 1]

    if (event.shiftKey && document.activeElement === firstElement) {
      event.preventDefault()
      lastElement?.focus()
    } else if (!event.shiftKey && document.activeElement === lastElement) {
      event.preventDefault()
      firstElement?.focus()
    }
  }
}

function lockBodyScroll() {
  if (props.lockScroll) {
    document.body.style.overflow = 'hidden'
    document.body.style.paddingRight = `${window.innerWidth - document.documentElement.clientWidth}px`
  }
}

function unlockBodyScroll() {
  if (props.lockScroll) {
    document.body.style.overflow = ''
    document.body.style.paddingRight = ''
  }
}

// =============================================================================
// WATCHERS
// =============================================================================

watch(() => props.modelValue, (isOpen) => {
  if (isOpen) {
    // Store the current active element to restore focus later
    previousActiveElement.value = document.activeElement as HTMLElement
    
    // Lock scroll
    lockBodyScroll()
    
    // Emit open event
    emit('open')
    
    // Focus management
    nextTick(() => {
      if (props.initialFocus) {
        props.initialFocus.focus()
      } else if (panelRef.value) {
        // Focus the first focusable element or the panel itself
        const firstFocusable = panelRef.value.querySelector<HTMLElement>(
          'button, [href], input, select, textarea, [tabindex]:not([tabindex="-1"])'
        )
        if (firstFocusable) {
          firstFocusable.focus()
        } else {
          panelRef.value.setAttribute('tabindex', '-1')
          panelRef.value.focus()
        }
      }
    })
  } else {
    // Unlock scroll
    unlockBodyScroll()
  }
})

// =============================================================================
// LIFECYCLE
// =============================================================================

onMounted(() => {
  document.addEventListener('keydown', handleKeydown)
  
  // If modal is open on mount, lock scroll
  if (props.modelValue) {
    lockBodyScroll()
  }
})

onUnmounted(() => {
  document.removeEventListener('keydown', handleKeydown)
  unlockBodyScroll()
})
</script>
