export interface Subscription {
  id: string
  user_id: string
  tier: SubscriptionTier
  status: 'active' | 'cancelled' | 'past_due' | 'inactive'
  current_period_start?: string
  current_period_end?: string
  cancel_at_period_end: boolean
  cancelled_at?: string
  created_at: string
  updated_at: string
}

export type SubscriptionTier = 'free' | 'scout' | 'pro' | 'club'

export interface SubscriptionPlan {
  id: string
  name: string
  tier: SubscriptionTier
  price_monthly: number
  price_yearly: number
  features: string[]
  is_popular?: boolean
}

export const SUBSCRIPTION_TIERS: Record<SubscriptionTier, SubscriptionPlan> = {
  free: {
    id: 'free',
    name: 'Free',
    tier: 'free',
    price_monthly: 0,
    price_yearly: 0,
    features: [
      'Browse player profiles',
      'Watch highlight videos',
      'Basic search filters',
    ],
  },
  scout: {
    id: 'scout',
    name: 'Scout',
    tier: 'scout',
    price_monthly: 99,
    price_yearly: 990,
    features: [
      'Everything in Free',
      'Watch full match recordings',
      'Save players to favorites',
      'Advanced search filters',
      'Email notifications for new players',
    ],
  },
  pro: {
    id: 'pro',
    name: 'Pro',
    tier: 'pro',
    price_monthly: 299,
    price_yearly: 2990,
    is_popular: true,
    features: [
      'Everything in Scout',
      'Contact players via academy',
      'Priority support',
      'Export player data (PDF)',
      'Tournament coverage alerts',
    ],
  },
  club: {
    id: 'club',
    name: 'Club',
    tier: 'club',
    price_monthly: 999,
    price_yearly: 9990,
    features: [
      'Everything in Pro',
      'API access for integration',
      'Dedicated account manager',
      'Custom reports',
      'Early access to new features',
      'Unlimited contact requests',
    ],
  },
}
