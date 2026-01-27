<template>
  <svg class="w-full h-full" :viewBox="`0 0 ${width} ${height}`" preserveAspectRatio="none">
    <!-- Grid lines -->
    <g class="grid-lines">
      <line 
        v-for="i in 5" 
        :key="i"
        :x1="padding"
        :y1="padding + ((height - 2 * padding) / 5) * (i - 1)"
        :x2="width - padding"
        :y2="padding + ((height - 2 * padding) / 5) * (i - 1)"
        stroke="#E5E7EB"
        stroke-width="1"
      />
    </g>
    
    <!-- Lines for each dataset -->
    <g v-for="(dataset, datasetIndex) in data.datasets" :key="datasetIndex">
      <!-- Line path -->
      <path
        :d="getLinePath(dataset)"
        fill="none"
        :stroke="colors[datasetIndex]"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
      <!-- Area fill -->
      <path
        :d="getAreaPath(dataset)"
        :fill="`url(#gradient-${datasetIndex})`"
        opacity="0.1"
      />
      <!-- Dots -->
      <circle
        v-for="(value, i) in dataset"
        :key="i"
        :cx="getX(i)"
        :cy="getY(value)"
        r="3"
        :fill="colors[datasetIndex]"
      />
    </g>
    
    <!-- Gradients -->
    <defs>
      <linearGradient 
        v-for="(color, i) in colors" 
        :key="i"
        :id="`gradient-${i}`" 
        x1="0%" 
        y1="0%" 
        x2="0%" 
        y2="100%"
      >
        <stop offset="0%" :stop-color="color" />
        <stop offset="100%" :stop-color="color" stop-opacity="0" />
      </linearGradient>
    </defs>
  </svg>
  
  <!-- Legend -->
  <div v-if="labels?.length" class="flex justify-center gap-4 mt-2">
    <div v-for="(label, i) in labels" :key="label" class="flex items-center gap-1.5 text-sm">
      <span class="w-3 h-3 rounded-full" :style="{ backgroundColor: colors[i] }"></span>
      <span class="text-neutral-600">{{ label }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
interface ChartData {
  labels: string[]
  datasets: number[][]
}

const props = withDefaults(defineProps<{
  data: ChartData
  colors?: string[]
  labels?: string[]
}>(), {
  colors: () => ['#3B82F6', '#10B981'],
  labels: () => []
})

const width = 400
const height = 200
const padding = 20

const maxValue = computed(() => {
  const allValues = props.data.datasets.flat()
  return Math.max(...allValues, 1)
})

function getX(index: number): number {
  const dataLength = props.data.labels.length
  if (dataLength <= 1) return width / 2
  const step = (width - 2 * padding) / (dataLength - 1)
  return padding + index * step
}

function getY(value: number): number {
  const chartHeight = height - 2 * padding
  const normalized = value / maxValue.value
  return height - padding - (normalized * chartHeight)
}

function getLinePath(dataset: number[]): string {
  if (dataset.length === 0) return ''
  
  const points = dataset.map((value, i) => `${getX(i)},${getY(value)}`)
  return `M ${points.join(' L ')}`
}

function getAreaPath(dataset: number[]): string {
  if (dataset.length === 0) return ''
  
  const points = dataset.map((value, i) => `${getX(i)},${getY(value)}`)
  const startX = getX(0)
  const endX = getX(dataset.length - 1)
  const bottomY = height - padding
  
  return `M ${startX},${bottomY} L ${points.join(' L ')} L ${endX},${bottomY} Z`
}
</script>
