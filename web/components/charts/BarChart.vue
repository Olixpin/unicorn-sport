<template>
  <svg class="w-full h-full" :viewBox="`0 0 ${width} ${height}`" preserveAspectRatio="none">
    <!-- Grid lines -->
    <g class="grid-lines">
      <line 
        v-for="i in 4" 
        :key="i"
        :x1="padding"
        :y1="padding + ((height - 2 * padding) / 4) * i"
        :x2="width - padding"
        :y2="padding + ((height - 2 * padding) / 4) * i"
        stroke="#E5E7EB"
        stroke-width="1"
      />
    </g>
    
    <!-- Bars -->
    <g>
      <rect
        v-for="(value, i) in data.values"
        :key="i"
        :x="getBarX(i)"
        :y="getBarY(value)"
        :width="barWidth"
        :height="getBarHeight(value)"
        :fill="color"
        rx="4"
        class="transition-all duration-300 hover:opacity-80"
      />
    </g>
    
    <!-- X-axis labels (show only some) -->
    <g v-if="showLabels" class="x-labels">
      <text
        v-for="(label, i) in visibleLabels"
        :key="label.index"
        :x="getBarX(label.index) + barWidth / 2"
        :y="height - 4"
        text-anchor="middle"
        class="fill-neutral-400 text-[8px]"
      >
        {{ label.text }}
      </text>
    </g>
  </svg>
</template>

<script setup lang="ts">
interface BarChartData {
  labels: string[]
  values: number[]
}

const props = withDefaults(defineProps<{
  data: BarChartData
  color?: string
  showLabels?: boolean
}>(), {
  color: '#3B82F6',
  showLabels: true
})

const width = 400
const height = 150
const padding = 20
const barGap = 4

import { computed } from 'vue'

const maxValue = computed(() => Math.max(...props.data.values, 1))

const barWidth = computed(() => {
  const availableWidth = width - 2 * padding
  const totalGaps = (props.data.values.length - 1) * barGap
  return (availableWidth - totalGaps) / props.data.values.length
})

function getBarX(index: number): number {
  return padding + index * (barWidth.value + barGap)
}

function getBarY(value: number): number {
  const chartHeight = height - 2 * padding - 10
  const normalized = value / maxValue.value
  return height - padding - 10 - (normalized * chartHeight)
}

function getBarHeight(value: number): number {
  const chartHeight = height - 2 * padding - 10
  return (value / maxValue.value) * chartHeight
}

const visibleLabels = computed(() => {
  const labels = props.data.labels
  const step = Math.ceil(labels.length / 6)
  return labels
    .map((text, index) => ({ text, index }))
    .filter((_, i) => i % step === 0)
})
</script>
