import { ref } from 'vue'

export interface Toast {
  id: string
  type: 'success' | 'error' | 'warning' | 'info'
  title: string
  message?: string
  duration?: number
}

const toasts = ref<Toast[]>([])

export const useToast = () => {
  const show = (toast: Omit<Toast, 'id'>) => {
    const id = Math.random().toString(36).substring(2, 9)
    const newToast: Toast = { ...toast, id }
    toasts.value.push(newToast)

    // Auto remove after duration
    setTimeout(() => {
      remove(id)
    }, toast.duration || 5000)

    return id
  }

  const remove = (id: string) => {
    const index = toasts.value.findIndex((t) => t.id === id)
    if (index > -1) {
      toasts.value.splice(index, 1)
    }
  }

  const success = (title: string, message?: string) =>
    show({ type: 'success', title, message })

  const error = (title: string, message?: string) =>
    show({ type: 'error', title, message })

  const warning = (title: string, message?: string) =>
    show({ type: 'warning', title, message })

  const info = (title: string, message?: string) =>
    show({ type: 'info', title, message })

  return {
    toasts,
    show,
    remove,
    success,
    error,
    warning,
    info,
  }
}
