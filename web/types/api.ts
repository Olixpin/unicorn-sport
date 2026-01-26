// API Response Types
export interface ApiResponse<T> {
  success: boolean
  data: T
  message?: string
}

export interface ApiError {
  success: false
  message: string
  errors?: Record<string, string[]>
}

export interface PaginatedResponse<T> {
  success: boolean
  data: {
    items: T[]
    pagination: Pagination
  }
}

export interface Pagination {
  page: number
  limit: number
  total: number
  total_pages: number
}
