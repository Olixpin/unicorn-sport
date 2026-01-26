import { z } from 'zod'

// Academy create/edit validation schema
export const academySchema = z.object({
  name: z
    .string()
    .min(2, 'Academy name must be at least 2 characters')
    .max(100, 'Academy name must be less than 100 characters'),
  description: z.string().optional(),
  country: z.string().min(2, 'Country is required'),
  state: z.string().optional(),
  city: z.string().optional(),
  address: z.string().optional(),
  phone: z.string().optional(),
  email: z.string().email('Invalid email address').optional().or(z.literal('')),
  website: z.string().url('Invalid URL').optional().or(z.literal('')),
  founded_year: z.number().min(1900).max(new Date().getFullYear()).optional(),
  logo_url: z.string().url().optional().or(z.literal('')),
})

export type AcademyFormData = z.infer<typeof academySchema>

// Academy interface
export interface Academy {
  id: string
  name: string
  description?: string
  country: string
  state?: string
  city?: string
  address?: string
  phone?: string
  email?: string
  website?: string
  founded_year?: number
  logo_url?: string
  player_count?: number
  is_verified?: boolean
  created_at: string
  updated_at: string
}
