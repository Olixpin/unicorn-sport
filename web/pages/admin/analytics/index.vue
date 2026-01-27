<template>
  <div class="max-w-7xl mx-auto">
    <!-- Page Header -->
    <div class="mb-8">
      <div class="flex flex-col lg:flex-row lg:items-center lg:justify-between gap-4">
        <div>
          <h1 class="font-display text-2xl lg:text-3xl font-bold text-neutral-900">
            Analytics
          </h1>
          <p class="mt-1 text-neutral-500">Platform growth metrics and insights</p>
        </div>
        <div class="flex items-center gap-3">
          <!-- Time Range Selector -->
          <select
            v-model="timeRange"
            @change="fetchAnalytics"
            class="px-4 py-2 bg-white border border-neutral-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-primary-500/20"
          >
            <option value="7">Last 7 days</option>
            <option value="30">Last 30 days</option>
            <option value="90">Last 90 days</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex justify-center py-20">
      <div class="w-8 h-8 border-2 border-primary-500 border-t-transparent rounded-full animate-spin"></div>
    </div>

    <div v-else class="space-y-6">
      <!-- Activity Charts Row -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- User & Player Signups Chart -->
        <div class="bg-white rounded-2xl border border-neutral-200 p-6">
          <h3 class="font-semibold text-neutral-900 mb-4">User & Player Signups</h3>
          <div class="h-64">
            <LineChart
              :data="signupsChartData"
              :colors="['#3B82F6', '#10B981']"
              :labels="['Users', 'Players']"
            />
          </div>
        </div>

        <!-- Video & Highlight Uploads Chart -->
        <div class="bg-white rounded-2xl border border-neutral-200 p-6">
          <h3 class="font-semibold text-neutral-900 mb-4">Content Uploads</h3>
          <div class="h-64">
            <LineChart
              :data="uploadsChartData"
              :colors="['#8B5CF6', '#F59E0B']"
              :labels="['Videos', 'Highlights']"
            />
          </div>
        </div>
      </div>

      <!-- Breakdown Charts Row -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- Subscription Tiers -->
        <div class="bg-white rounded-2xl border border-neutral-200 p-6">
          <h3 class="font-semibold text-neutral-900 mb-4">Subscription Breakdown</h3>
          <div class="h-48">
            <DonutChart :data="subscriptionData" />
          </div>
          <div class="mt-4 space-y-2">
            <div v-for="(item, index) in subscriptionData" :key="item.label" class="flex items-center justify-between text-sm">
              <div class="flex items-center gap-2">
                <span :class="['w-3 h-3 rounded-full', donutColors[index % donutColors.length]]"></span>
                <span class="text-neutral-600">{{ item.label }}</span>
              </div>
              <span class="font-medium text-neutral-900">{{ item.value }}</span>
            </div>
          </div>
        </div>

        <!-- Player Positions -->
        <div class="bg-white rounded-2xl border border-neutral-200 p-6">
          <h3 class="font-semibold text-neutral-900 mb-4">Player Positions</h3>
          <div class="space-y-3">
            <div v-for="item in positionData" :key="item.label" class="space-y-1">
              <div class="flex justify-between text-sm">
                <span class="text-neutral-600">{{ item.label }}</span>
                <span class="font-medium text-neutral-900">{{ item.value }}</span>
              </div>
              <div class="h-2 bg-neutral-100 rounded-full overflow-hidden">
                <div 
                  class="h-full bg-gradient-to-r from-primary-500 to-primary-400 rounded-full transition-all duration-500"
                  :style="{ width: `${(item.value / maxPositionValue) * 100}%` }"
                ></div>
              </div>
            </div>
          </div>
        </div>

        <!-- Player Countries -->
        <div class="bg-white rounded-2xl border border-neutral-200 p-6">
          <h3 class="font-semibold text-neutral-900 mb-4">Top Countries</h3>
          <div class="space-y-3">
            <div v-for="item in countryData" :key="item.label" class="space-y-1">
              <div class="flex justify-between text-sm">
                <span class="text-neutral-600">{{ item.label }}</span>
                <span class="font-medium text-neutral-900">{{ item.value }}</span>
              </div>
              <div class="h-2 bg-neutral-100 rounded-full overflow-hidden">
                <div 
                  class="h-full bg-gradient-to-r from-emerald-500 to-emerald-400 rounded-full transition-all duration-500"
                  :style="{ width: `${(item.value / maxCountryValue) * 100}%` }"
                ></div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Contact Requests Chart -->
      <div class="bg-white rounded-2xl border border-neutral-200 p-6">
        <h3 class="font-semibold text-neutral-900 mb-4">Contact Requests Over Time</h3>
        <div class="h-48">
          <BarChart :data="contactRequestsData" color="#EF4444" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { ApiResponse } from '~/types'

definePageMeta({
  layout: 'admin',
  middleware: 'admin',
})

const api = useApi()

const loading = ref(true)
const timeRange = ref('30')

// Raw data from API
const analyticsData = ref<{
  user_signups: Array<{ date: string; count: number }>
  player_signups: Array<{ date: string; count: number }>
  video_uploads: Array<{ date: string; count: number }>
  highlight_uploads: Array<{ date: string; count: number }>
  contact_requests: Array<{ date: string; count: number }>
  subscription_tiers: Array<{ subscription_tier: string; count: number }>
  player_positions: Array<{ subscription_tier: string; count: number }>
  player_countries: Array<{ subscription_tier: string; count: number }>
}>({
  user_signups: [],
  player_signups: [],
  video_uploads: [],
  highlight_uploads: [],
  contact_requests: [],
  subscription_tiers: [],
  player_positions: [],
  player_countries: [],
})

const donutColors = [
  'bg-blue-500',
  'bg-emerald-500',
  'bg-amber-500',
  'bg-purple-500',
  'bg-pink-500',
]

// Computed chart data
const signupsChartData = computed(() => {
  const dates = new Set([
    ...analyticsData.value.user_signups.map(d => d.date),
    ...analyticsData.value.player_signups.map(d => d.date)
  ])
  const sortedDates = Array.from(dates).sort()
  
  const userMap = new Map(analyticsData.value.user_signups.map(d => [d.date, d.count]))
  const playerMap = new Map(analyticsData.value.player_signups.map(d => [d.date, d.count]))
  
  return {
    labels: sortedDates.map(d => formatDateLabel(d)),
    datasets: [
      sortedDates.map(d => userMap.get(d) || 0),
      sortedDates.map(d => playerMap.get(d) || 0)
    ]
  }
})

const uploadsChartData = computed(() => {
  const dates = new Set([
    ...analyticsData.value.video_uploads.map(d => d.date),
    ...analyticsData.value.highlight_uploads.map(d => d.date)
  ])
  const sortedDates = Array.from(dates).sort()
  
  const videoMap = new Map(analyticsData.value.video_uploads.map(d => [d.date, d.count]))
  const highlightMap = new Map(analyticsData.value.highlight_uploads.map(d => [d.date, d.count]))
  
  return {
    labels: sortedDates.map(d => formatDateLabel(d)),
    datasets: [
      sortedDates.map(d => videoMap.get(d) || 0),
      sortedDates.map(d => highlightMap.get(d) || 0)
    ]
  }
})

const contactRequestsData = computed(() => {
  const sorted = [...analyticsData.value.contact_requests].sort((a, b) => a.date.localeCompare(b.date))
  return {
    labels: sorted.map(d => formatDateLabel(d.date)),
    values: sorted.map(d => d.count)
  }
})

const subscriptionData = computed(() => {
  return analyticsData.value.subscription_tiers.map(t => ({
    label: t.subscription_tier || 'Free',
    value: t.count
  }))
})

const positionData = computed(() => {
  return analyticsData.value.player_positions.slice(0, 6).map(p => ({
    label: p.subscription_tier,
    value: p.count
  }))
})

const maxPositionValue = computed(() => {
  return Math.max(...positionData.value.map(p => p.value), 1)
})

const countryData = computed(() => {
  return analyticsData.value.player_countries.slice(0, 6).map(c => ({
    label: c.subscription_tier,
    value: c.count
  }))
})

const maxCountryValue = computed(() => {
  return Math.max(...countryData.value.map(c => c.value), 1)
})

function formatDateLabel(dateStr: string): string {
  const date = new Date(dateStr)
  return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric' })
}

async function fetchAnalytics() {
  loading.value = true
  try {
    const response = await api.get<ApiResponse<typeof analyticsData.value>>(
      `/admin/analytics?days=${timeRange.value}`,
      {},
      true
    )
    if (response.success && response.data) {
      analyticsData.value = response.data
    }
  } catch (error) {
    console.error('Failed to fetch analytics:', error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchAnalytics()
})
</script>
