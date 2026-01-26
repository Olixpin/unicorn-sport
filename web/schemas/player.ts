import { z } from 'zod'

// Player create/edit validation schema
export const playerSchema = z.object({
  first_name: z
    .string()
    .min(2, 'First name must be at least 2 characters')
    .max(50, 'First name must be less than 50 characters'),
  last_name: z
    .string()
    .min(2, 'Last name must be at least 2 characters')
    .max(50, 'Last name must be less than 50 characters'),
  date_of_birth: z
    .string()
    .regex(/^\d{4}-\d{2}-\d{2}$/, 'Date must be in YYYY-MM-DD format'),
  position: z.enum([
    'Goalkeeper',
    'Defender',
    'Midfielder',
    'Forward',
    'Center Back',
    'Full Back',
    'Wing Back',
    'Defensive Midfielder',
    'Central Midfielder',
    'Attacking Midfielder',
    'Winger',
    'Striker',
  ], { errorMap: () => ({ message: 'Please select a valid position' }) }),
  preferred_foot: z.enum(['left', 'right', 'both']).optional(),
  height_cm: z.number().min(100).max(250).optional(),
  weight_kg: z.number().min(30).max(150).optional(),
  country: z
    .string()
    .min(2, 'Country is required'),
  state: z.string().optional(),
  city: z.string().optional(),
  school_name: z.string().optional(),
  academy_id: z.string().uuid().optional(),
})

export type PlayerFormData = z.infer<typeof playerSchema>

// Positions list for dropdowns
export const PLAYER_POSITIONS = [
  'Goalkeeper',
  'Defender',
  'Center Back',
  'Full Back',
  'Wing Back',
  'Midfielder',
  'Defensive Midfielder',
  'Central Midfielder',
  'Attacking Midfielder',
  'Forward',
  'Winger',
  'Striker',
] as const

// African countries for dropdown
export const AFRICAN_COUNTRIES = [
  'Algeria', 'Angola', 'Benin', 'Botswana', 'Burkina Faso', 'Burundi',
  'Cameroon', 'Cape Verde', 'Central African Republic', 'Chad', 'Comoros',
  'Democratic Republic of the Congo', 'Republic of the Congo', 'Djibouti',
  'Egypt', 'Equatorial Guinea', 'Eritrea', 'Eswatini', 'Ethiopia',
  'Gabon', 'Gambia', 'Ghana', 'Guinea', 'Guinea-Bissau', 'Ivory Coast',
  'Kenya', 'Lesotho', 'Liberia', 'Libya', 'Madagascar', 'Malawi', 'Mali',
  'Mauritania', 'Mauritius', 'Morocco', 'Mozambique', 'Namibia', 'Niger',
  'Nigeria', 'Rwanda', 'São Tomé and Príncipe', 'Senegal', 'Seychelles',
  'Sierra Leone', 'Somalia', 'South Africa', 'South Sudan', 'Sudan',
  'Tanzania', 'Togo', 'Tunisia', 'Uganda', 'Zambia', 'Zimbabwe',
] as const
