<template>
  <svg class="w-full h-full" :viewBox="`0 0 ${size} ${size}`">
    <!-- Background circle -->
    <circle
      :cx="center"
      :cy="center"
      :r="radius"
      fill="none"
      stroke="#F3F4F6"
      :stroke-width="strokeWidth"
    />
    
    <!-- Segments -->
    <circle
      v-for="(segment, i) in segments"
      :key="i"
      :cx="center"
      :cy="center"
      :r="radius"
      fill="none"
      :stroke="segmentColors[i % segmentColors.length]"
      :stroke-width="strokeWidth"
      :stroke-dasharray="`${segment.length} ${circumference}`"
      :stroke-dashoffset="-segment.offset"
      stroke-linecap="round"
      class="transition-all duration-500"
    />
    
    <!-- Center text -->
    <text
      :x="center"
      :y="center"
      text-anchor="middle"
      dominant-baseline="middle"
      class="fill-neutral-900 font-bold text-xl"
    >
      {{ total }}
    </text>
    <text
      :x="center"
      :y="center + 18"
      text-anchor="middle"
      dominant-baseline="middle"
      class="fill-neutral-500 text-xs"
    >
      Total
    </text>
  </svg>
</template>

<script setup lang="ts">
interface DataItem {
  label: string
  value: number
}

const props = defineProps<{
  data: DataItem[]
}>()

const size = 160
const center = size / 2
const strokeWidth = 24
const radius = (size - strokeWidth) / 2
const circumference = 2 * Math.PI * radius

const segmentColors = [
  '#3B82F6',
  '#10B981',
  '#F59E0B',
  '#8B5CF6',
  '#EC4899',
]

const total = computed(() => props.data.reduce((sum, item) => sum + item.value, 0))

const segments = computed(() => {
  if (total.value === 0) return []
  
  let offset = circumference * 0.25 // Start from top
  return props.data.map((item) => {
    const length = (item.value / total.value) * circumference
    const segment = { length, offset }
    offset -= length
    return segment
  })
})
</script>
