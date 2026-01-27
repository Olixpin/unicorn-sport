/**
 * useModal Composable
 * 
 * A composable for programmatic modal control. Provides a clean API for
 * opening, closing, and managing modal state throughout the application.
 * 
 * @example Basic Usage
 * ```ts
 * const { isOpen, open, close, toggle } = useModal()
 * 
 * // In template: <UiModal v-model="isOpen">...</UiModal>
 * ```
 * 
 * @example With Data
 * ```ts
 * const { isOpen, data, openWith, close } = useModal<{ videoId: string }>()
 * 
 * openWith({ videoId: '123' })
 * // data.value?.videoId === '123'
 * ```
 * 
 * @example Multiple Modals
 * ```ts
 * const confirmModal = useModal()
 * const videoModal = useModal<VideoData>()
 * const editModal = useModal<EditData>()
 * ```
 */

interface UseModalOptions {
  /** Initial open state */
  initialValue?: boolean
  /** Callback when modal opens */
  onOpen?: () => void
  /** Callback when modal closes */
  onClose?: () => void
}

interface UseModalReturn<T = undefined> {
  /** Reactive open state - use with v-model */
  isOpen: Ref<boolean>
  /** Data passed to the modal */
  data: Ref<T | undefined>
  /** Open the modal */
  open: () => void
  /** Open the modal with data */
  openWith: (payload: T) => void
  /** Close the modal */
  close: () => void
  /** Toggle the modal state */
  toggle: () => void
  /** Set data without opening */
  setData: (payload: T) => void
  /** Clear data */
  clearData: () => void
}

export function useModal<T = undefined>(options: UseModalOptions = {}): UseModalReturn<T> {
  const { initialValue = false, onOpen, onClose } = options

  const isOpen = ref(initialValue)
  const data = ref<T | undefined>(undefined) as Ref<T | undefined>

  function open() {
    isOpen.value = true
    onOpen?.()
  }

  function openWith(payload: T) {
    data.value = payload
    open()
  }

  function close() {
    isOpen.value = false
    onClose?.()
  }

  function toggle() {
    if (isOpen.value) {
      close()
    } else {
      open()
    }
  }

  function setData(payload: T) {
    data.value = payload
  }

  function clearData() {
    data.value = undefined
  }

  // Watch for external changes to isOpen
  watch(isOpen, (newValue, oldValue) => {
    if (newValue && !oldValue) {
      onOpen?.()
    } else if (!newValue && oldValue) {
      onClose?.()
    }
  })

  return {
    isOpen,
    data,
    open,
    openWith,
    close,
    toggle,
    setData,
    clearData,
  }
}

// =============================================================================
// MODAL MANAGER (for stacking modals)
// =============================================================================

interface ModalInstance {
  id: string
  close: () => void
}

const modalStack = ref<ModalInstance[]>([])

/**
 * useModalManager
 * 
 * Manages a stack of modals for proper z-index and closing behavior.
 * Useful when you have nested or multiple simultaneous modals.
 */
export function useModalManager() {
  function register(id: string, closeFn: () => void) {
    modalStack.value.push({ id, close: closeFn })
    return () => unregister(id)
  }

  function unregister(id: string) {
    const index = modalStack.value.findIndex((m) => m.id === id)
    if (index > -1) {
      modalStack.value.splice(index, 1)
    }
  }

  function closeTopmost() {
    const topModal = modalStack.value[modalStack.value.length - 1]
    if (topModal) {
      topModal.close()
    }
  }

  function closeAll() {
    // Close from top to bottom
    while (modalStack.value.length > 0) {
      closeTopmost()
    }
  }

  function getStackPosition(id: string): number {
    return modalStack.value.findIndex((m) => m.id === id)
  }

  function isTopmost(id: string): boolean {
    return modalStack.value[modalStack.value.length - 1]?.id === id
  }

  return {
    stack: readonly(modalStack),
    register,
    unregister,
    closeTopmost,
    closeAll,
    getStackPosition,
    isTopmost,
  }
}

// =============================================================================
// CONFIRM MODAL COMPOSABLE
// =============================================================================

interface ConfirmOptions {
  title: string
  message: string
  confirmText?: string
  cancelText?: string
  confirmVariant?: 'primary' | 'danger' | 'warning'
  icon?: 'warning' | 'danger' | 'info' | 'success' | 'trash' | 'question'
}

interface ConfirmModalState {
  isOpen: boolean
  options: ConfirmOptions | null
  resolve: ((value: boolean) => void) | null
}

const confirmState = reactive<ConfirmModalState>({
  isOpen: false,
  options: null,
  resolve: null,
})

/**
 * useConfirm
 * 
 * Programmatic confirm dialogs. Returns a promise that resolves to
 * true (confirmed) or false (cancelled).
 * 
 * @example
 * ```ts
 * const { confirm } = useConfirm()
 * 
 * const confirmed = await confirm({
 *   title: 'Delete Player',
 *   message: 'Are you sure you want to delete this player?',
 *   confirmVariant: 'danger'
 * })
 * 
 * if (confirmed) {
 *   // proceed with deletion
 * }
 * ```
 */
export function useConfirm() {
  function confirm(options: ConfirmOptions): Promise<boolean> {
    return new Promise((resolve) => {
      confirmState.options = {
        confirmText: 'Confirm',
        cancelText: 'Cancel',
        confirmVariant: 'primary',
        ...options,
      }
      confirmState.resolve = resolve
      confirmState.isOpen = true
    })
  }

  function handleConfirm() {
    confirmState.resolve?.(true)
    closeConfirm()
  }

  function handleCancel() {
    confirmState.resolve?.(false)
    closeConfirm()
  }

  function closeConfirm() {
    confirmState.isOpen = false
    confirmState.options = null
    confirmState.resolve = null
  }

  return {
    confirm,
    handleConfirm,
    handleCancel,
    state: readonly(confirmState),
  }
}
