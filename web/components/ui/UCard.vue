<template>
  <div :class="['card', { 'card-hover cursor-pointer': hoverable }]">
    <div v-if="$slots.header" class="px-6 py-4 border-b border-neutral-200">
      <slot name="header" />
    </div>
    <div :class="['px-6 py-4', bodyClass]">
      <slot />
    </div>
    <div v-if="$slots.footer" class="px-6 py-4 border-t border-neutral-200 bg-neutral-50">
      <slot name="footer" />
    </div>
  </div>
</template>

<script setup lang="ts">
interface Props {
  hoverable?: boolean
  bodyClass?: string
}

withDefaults(defineProps<Props>(), {
  hoverable: false,
  bodyClass: '',
})
</script>
