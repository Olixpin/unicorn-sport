<template>
  <Teleport to="body">
    <div class="fixed bottom-6 right-6 z-[9999] flex flex-col gap-3 max-w-sm pointer-events-none">
      <TransitionGroup name="toast">
        <div
          v-for="toast in toasts"
          :key="toast.id"
          :class="[
            'pointer-events-auto relative flex items-start gap-3 p-4 rounded-2xl shadow-2xl border backdrop-blur-sm',
            variantClasses[toast.type],
          ]"
        >
          <!-- Icon with gradient background -->
          <div :class="['flex-shrink-0 w-10 h-10 rounded-xl flex items-center justify-center', iconBgClasses[toast.type]]">
            <!-- Success Icon -->
            <svg v-if="toast.type === 'success'" class="w-5 h-5 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M5 13l4 4L19 7" />
            </svg>
            <!-- Error Icon -->
            <svg v-else-if="toast.type === 'error'" class="w-5 h-5 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M6 18L18 6M6 6l12 12" />
            </svg>
            <!-- Warning Icon -->
            <svg v-else-if="toast.type === 'warning'" class="w-5 h-5 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
            </svg>
            <!-- Info Icon -->
            <svg v-else class="w-5 h-5 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          
          <!-- Content -->
          <div class="flex-1 min-w-0 pt-0.5">
            <p class="text-sm font-semibold text-neutral-900">{{ toast.title }}</p>
            <p v-if="toast.message" class="mt-0.5 text-sm text-neutral-500 line-clamp-2">{{ toast.message }}</p>
          </div>
          
          <!-- Close Button -->
          <button
            @click="remove(toast.id)"
            class="flex-shrink-0 w-8 h-8 flex items-center justify-center rounded-lg text-neutral-400 hover:text-neutral-600 hover:bg-neutral-100 transition-colors"
          >
            <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
          
          <!-- Progress bar -->
          <div class="absolute bottom-0 left-0 right-0 h-1 rounded-b-2xl overflow-hidden">
            <div 
              :class="['h-full animate-shrink', progressClasses[toast.type]]"
              :style="{ animationDuration: '5s' }"
            ></div>
          </div>
        </div>
      </TransitionGroup>
    </div>
  </Teleport>
</template>

<script setup lang="ts">
const { toasts, remove } = useToast()

const variantClasses: Record<string, string> = {
  success: 'bg-white/95 border-emerald-200/50',
  error: 'bg-white/95 border-red-200/50',
  warning: 'bg-white/95 border-amber-200/50',
  info: 'bg-white/95 border-blue-200/50',
}

const iconBgClasses: Record<string, string> = {
  success: 'bg-gradient-to-br from-emerald-500 to-green-600 shadow-lg shadow-emerald-500/30',
  error: 'bg-gradient-to-br from-red-500 to-rose-600 shadow-lg shadow-red-500/30',
  warning: 'bg-gradient-to-br from-amber-500 to-orange-600 shadow-lg shadow-amber-500/30',
  info: 'bg-gradient-to-br from-blue-500 to-indigo-600 shadow-lg shadow-blue-500/30',
}

const progressClasses: Record<string, string> = {
  success: 'bg-gradient-to-r from-emerald-500 to-green-500',
  error: 'bg-gradient-to-r from-red-500 to-rose-500',
  warning: 'bg-gradient-to-r from-amber-500 to-orange-500',
  info: 'bg-gradient-to-r from-blue-500 to-indigo-500',
}
</script>

<style scoped>
.toast-enter-active,
.toast-leave-active {
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

.toast-enter-from {
  opacity: 0;
  transform: translateX(100%) scale(0.95);
}

.toast-leave-to {
  opacity: 0;
  transform: translateX(100%) scale(0.95);
}

.toast-move {
  transition: transform 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

@keyframes shrink {
  from {
    width: 100%;
  }
  to {
    width: 0%;
  }
}

.animate-shrink {
  animation: shrink linear forwards;
}
</style>
