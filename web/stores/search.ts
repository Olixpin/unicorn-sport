import { defineStore } from 'pinia'
import type { Player, PlayerFilters, Pagination, ApiResponse } from '~/types/index'

interface SearchState {
  query: string
  filters: PlayerFilters
  results: Player[]
  pagination: Pagination
  loading: boolean
}

export const useSearchStore = defineStore('search', {
  state: (): SearchState => ({
    query: '',
    filters: {
      position: undefined,
      country: undefined,
      age_min: undefined,
      age_max: undefined,
      verified_only: false,
      page: 1,
      limit: 20,
    },
    results: [],
    pagination: {
      page: 1,
      limit: 20,
      total: 0,
      total_pages: 0,
    },
    loading: false,
  }),

  actions: {
    async search(): Promise<void> {
      this.loading = true
      try {
        const config = useRuntimeConfig()
        const params: Record<string, unknown> = {
          page: this.filters.page || 1,
          limit: this.filters.limit || 20,
        }

        if (this.query) params.q = this.query
        if (this.filters.position) params.position = this.filters.position
        if (this.filters.country) params.country = this.filters.country
        if (this.filters.age_min) params.min_age = this.filters.age_min
        if (this.filters.age_max) params.max_age = this.filters.age_max
        if (this.filters.verified_only) params.verified = true
        if (this.filters.tournament_id) params.tournament_id = this.filters.tournament_id

        const response = await $fetch<ApiResponse<{ players: Player[]; pagination: Pagination }>>('/players', {
          baseURL: config.public.apiBase,
          method: 'GET',
          query: params,
        })
        
        if (response.success && response.data) {
          this.results = response.data.players || []
          this.pagination = response.data.pagination || this.pagination
        }
      } catch (error) {
        console.error('Search failed:', error)
        this.results = []
      } finally {
        this.loading = false
      }
    },

    setFilter(key: keyof PlayerFilters, value: unknown): void {
      (this.filters as Record<string, unknown>)[key] = value
      this.filters.page = 1 // Reset to first page when filter changes
    },

    setQuery(query: string): void {
      this.query = query
      this.filters.page = 1
    },

    resetFilters(): void {
      this.query = ''
      this.filters = {
        position: undefined,
        country: undefined,
        age_min: undefined,
        age_max: undefined,
        verified_only: false,
        page: 1,
        limit: 20,
      }
    },

    async loadMore(): Promise<void> {
      if (this.pagination.page >= this.pagination.total_pages) return
      this.filters.page = (this.filters.page || 1) + 1
      
      this.loading = true
      try {
        const config = useRuntimeConfig()
        const params: Record<string, unknown> = {
          page: this.filters.page,
          limit: this.filters.limit || 20,
        }

        if (this.query) params.q = this.query
        if (this.filters.position) params.position = this.filters.position
        if (this.filters.country) params.country = this.filters.country

        const response = await $fetch<ApiResponse<{ players: Player[]; pagination: Pagination }>>('/players', {
          baseURL: config.public.apiBase,
          method: 'GET',
          query: params,
        })
        
        if (response.success && response.data) {
          this.results = [...this.results, ...(response.data.players || [])]
          this.pagination = response.data.pagination || this.pagination
        }
      } catch (error) {
        console.error('Load more failed:', error)
      } finally {
        this.loading = false
      }
    },
  },
})
