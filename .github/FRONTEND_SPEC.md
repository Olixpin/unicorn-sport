# Unicorn Sport - Frontend Specification

## 🎯 Overview

This document provides the complete blueprint for building the Unicorn Sport frontend. The application uses **Nuxt 3** with **Vue 3 Composition API**, following Google-scale engineering principles: type safety, component isolation, and clean architecture.

---

## 🛠️ Technology Stack

| Layer | Technology | Rationale |
|-------|------------|-----------|
| **Framework** | Nuxt 3 | SSR for SEO, file-based routing, auto-imports |
| **UI Library** | Vue 3 Composition API | Type-safe, composable, performant |
| **Styling** | Tailwind CSS 3 | Utility-first, consistent design system |
| **State** | Pinia | Official Vue state management, TypeScript native |
| **HTTP** | ofetch (built-in) | Nuxt's native fetch with interceptors |
| **Forms** | VeeValidate + Zod | Schema-based validation |
| **Icons** | Heroicons | Consistent, accessible SVG icons |
| **Types** | TypeScript (strict) | Catch errors at compile time |

---

## 📁 Project Structure

```
web/
├── nuxt.config.ts              # Nuxt configuration
├── app.vue                     # Root component
├── tailwind.config.ts          # Tailwind configuration
├── tsconfig.json               # TypeScript config
│
├── assets/
│   └── css/
│       └── main.css            # Global styles, Tailwind imports
│
├── components/
│   ├── ui/                     # Atomic UI components
│   │   ├── UButton.vue
│   │   ├── UInput.vue
│   │   ├── UCard.vue
│   │   ├── UModal.vue
│   │   ├── UBadge.vue
│   │   ├── UAvatar.vue
│   │   ├── USpinner.vue
│   │   └── UAlert.vue
│   │
│   ├── layout/                 # Layout components
│   │   ├── AppHeader.vue
│   │   ├── AppFooter.vue
│   │   ├── AppSidebar.vue
│   │   └── AppNavbar.vue
│   │
│   ├── player/                 # Player-related components
│   │   ├── PlayerCard.vue      # Card in search results
│   │   ├── PlayerGrid.vue      # Grid of player cards
│   │   ├── PlayerStats.vue     # Physical stats display
│   │   ├── PlayerBio.vue       # Bio section
│   │   └── PlayerVideos.vue    # Video gallery
│   │
│   ├── video/                  # Video components
│   │   ├── VideoPlayer.vue     # Main video player
│   │   ├── VideoThumbnail.vue  # Thumbnail with play overlay
│   │   └── VideoGrid.vue       # Grid of videos
│   │
│   ├── search/                 # Search components
│   │   ├── SearchBar.vue       # Main search input
│   │   ├── SearchFilters.vue   # Position, age, country filters
│   │   └── SearchResults.vue   # Results container
│   │
│   ├── subscription/           # Subscription components
│   │   ├── PricingCard.vue     # Individual plan card
│   │   ├── PricingTable.vue    # All plans comparison
│   │   └── SubscriptionBadge.vue # Current tier badge
│   │
│   └── admin/                  # Admin-only components
│       ├── PlayerForm.vue      # Create/edit player
│       ├── VideoUploader.vue   # Upload interface
│       ├── EventForm.vue       # Create tournament/event
│       └── DataTable.vue       # Generic data table
│
├── composables/                # Reusable logic (auto-imported)
│   ├── useAuth.ts              # Authentication state & methods
│   ├── useApi.ts               # API client with auth headers
│   ├── usePlayer.ts            # Player data fetching
│   ├── useSearch.ts            # Search state & methods
│   ├── useSubscription.ts      # Subscription checks
│   └── useToast.ts             # Toast notifications
│
├── layouts/
│   ├── default.vue             # Public pages layout
│   ├── auth.vue                # Login/register layout (minimal)
│   ├── dashboard.vue           # Scout dashboard layout
│   └── admin.vue               # Admin panel layout
│
├── middleware/
│   ├── auth.ts                 # Require authentication
│   ├── guest.ts                # Redirect if authenticated
│   ├── admin.ts                # Require admin role
│   └── subscription.ts         # Require active subscription
│
├── pages/
│   ├── index.vue               # Landing page (public)
│   ├── discover.vue            # Player search/browse (public)
│   ├── players/
│   │   └── [id].vue            # Player profile page
│   ├── pricing.vue             # Subscription plans (public)
│   │
│   ├── auth/
│   │   ├── login.vue           # Scout login
│   │   ├── register.vue        # Scout registration
│   │   ├── forgot-password.vue
│   │   └── reset-password.vue
│   │
│   ├── dashboard/              # Scout dashboard (auth required)
│   │   ├── index.vue           # Dashboard home
│   │   ├── saved.vue           # Saved players
│   │   ├── contacts.vue        # Contact requests sent
│   │   └── settings.vue        # Account settings
│   │
│   └── admin/                  # Admin panel (admin role required)
│       ├── index.vue           # Admin dashboard
│       ├── players/
│       │   ├── index.vue       # Player list
│       │   ├── create.vue      # Create player
│       │   └── [id]/
│       │       ├── edit.vue    # Edit player
│       │       └── videos.vue  # Manage player videos
│       ├── events/
│       │   ├── index.vue       # Event list
│       │   └── create.vue      # Create event
│       ├── videos/
│       │   ├── index.vue       # All videos
│       │   └── upload.vue      # Upload videos
│       └── users/
│           └── index.vue       # User management
│
├── plugins/
│   └── api.ts                  # API client initialization
│
├── server/
│   └── api/                    # API routes (if needed for BFF)
│
├── stores/
│   ├── auth.ts                 # Auth state (user, tokens)
│   ├── player.ts               # Current player being viewed
│   ├── search.ts               # Search filters & results
│   └── subscription.ts         # Current subscription state
│
├── types/
│   ├── api.ts                  # API response types
│   ├── player.ts               # Player interfaces
│   ├── user.ts                 # User interfaces
│   ├── video.ts                # Video interfaces
│   └── subscription.ts         # Subscription interfaces
│
└── utils/
    ├── constants.ts            # App constants
    ├── formatters.ts           # Date, currency formatters
    └── validators.ts           # Zod schemas
```

---

## 🎨 Design System

### Color Palette

```css
/* tailwind.config.ts */
colors: {
  primary: {
    50: '#f0fdf4',
    100: '#dcfce7',
    500: '#22c55e',   /* Main green - African football feel */
    600: '#16a34a',
    700: '#15803d',
  },
  secondary: {
    500: '#eab308',   /* Gold accent */
    600: '#ca8a04',
  },
  neutral: {
    50: '#fafafa',
    100: '#f5f5f5',
    800: '#262626',
    900: '#171717',
  }
}
```

### Typography

```css
fontFamily: {
  sans: ['Inter', 'system-ui', 'sans-serif'],
  display: ['Plus Jakarta Sans', 'Inter', 'sans-serif'],
}
```

### Component Sizes

| Size | Padding | Font | Use Case |
|------|---------|------|----------|
| `sm` | `px-3 py-1.5` | `text-sm` | Secondary actions |
| `md` | `px-4 py-2` | `text-base` | Default |
| `lg` | `px-6 py-3` | `text-lg` | Primary CTAs |

---

## 📄 Page Specifications

### 1. Landing Page (`/`)

**Purpose:** Convert visitors to registered scouts

**Sections:**
```
┌─────────────────────────────────────────────────────────────────────────────┐
│  HEADER                                                                      │
│  Logo | Discover | Pricing | Login | Register (CTA)                        │
├─────────────────────────────────────────────────────────────────────────────┤
│  HERO SECTION                                                                │
│  "Discover Africa's Next Football Stars"                                    │
│  Subtitle + CTA buttons (Browse Players | Start Free Trial)                │
│  Background: Video montage or dynamic player images                         │
├─────────────────────────────────────────────────────────────────────────────┤
│  FEATURED PLAYERS (3-4 cards)                                               │
│  Verified, high-potential players                                           │
│  [PlayerCard] [PlayerCard] [PlayerCard]                                     │
├─────────────────────────────────────────────────────────────────────────────┤
│  HOW IT WORKS                                                                │
│  1. Browse Players → 2. Watch Highlights → 3. Contact via Academy          │
├─────────────────────────────────────────────────────────────────────────────┤
│  STATS/TRUST SECTION                                                        │
│  "500+ Verified Players" | "50+ Partner Academies" | "Document-Verified"   │
├─────────────────────────────────────────────────────────────────────────────┤
│  PRICING PREVIEW                                                            │
│  [PricingCard: Free] [PricingCard: Scout] [PricingCard: Pro]               │
├─────────────────────────────────────────────────────────────────────────────┤
│  FOOTER                                                                      │
│  Links | Social | © 2026 Unicorn Sport                                      │
└─────────────────────────────────────────────────────────────────────────────┘
```

**Data Requirements:**
- Featured players (cached, public API)
- Stats counters (cached)

---

### 2. Discover Page (`/discover`)

**Purpose:** Search and browse player profiles

**Layout:**
```
┌─────────────────────────────────────────────────────────────────────────────┐
│  SEARCH BAR                                                                  │
│  [🔍 Search by name, position, country...]                                  │
├─────────────────────────────────────────────────────────────────────────────┤
│  FILTERS (collapsible on mobile)                                            │
│  Position: [All] [GK] [DEF] [MID] [FWD]                                    │
│  Age: [14-16] [17-18] [19-21]                                              │
│  Country: [Nigeria] [Ghana] [Kenya] [...]                                  │
│  Verified: [✓] Only verified players                                       │
├─────────────────────────────────────────────────────────────────────────────┤
│  RESULTS HEADER                                                              │
│  "342 players found" | Sort: [Most Recent] [Name A-Z]                       │
├─────────────────────────────────────────────────────────────────────────────┤
│  PLAYER GRID (responsive)                                                    │
│  ┌────────┐ ┌────────┐ ┌────────┐ ┌────────┐                               │
│  │ Player │ │ Player │ │ Player │ │ Player │                               │
│  │  Card  │ │  Card  │ │  Card  │ │  Card  │                               │
│  └────────┘ └────────┘ └────────┘ └────────┘                               │
│  [Load More] or Infinite Scroll                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

**PlayerCard Display:**
```
┌─────────────────────────┐
│  [Profile Photo]        │
│                         │
│  James O. ✓             │  ← First name + last initial + verified badge
│  Striker • 17 years     │  ← Position + Age
│  Lagos, Nigeria 🇳🇬      │  ← Location + flag
│                         │
│  [▶ Watch Highlights]   │  ← CTA (links to profile)
└─────────────────────────┘
```

**Data Requirements:**
- GET `/api/v1/players?position=&country=&age_min=&age_max=&verified=&page=&limit=`
- Returns: `PlayerPublicDTO[]` with pagination

---

### 3. Player Profile (`/players/[id]`)

**Purpose:** Showcase individual player with videos

**Layout:**
```
┌─────────────────────────────────────────────────────────────────────────────┐
│  HERO SECTION                                                                │
│  ┌──────────────┐  James Okonkwo ✓                                          │
│  │              │  Striker | Right-footed                                   │
│  │   [Photo]    │  Age: 17 | Height: 178cm | Weight: 72kg                   │
│  │              │  Lagos, Nigeria                                            │
│  └──────────────┘  Academy: Lagos Football Academy                          │
│                    Tournament: Nationals 2025                                │
├─────────────────────────────────────────────────────────────────────────────┤
│  ACTION BUTTONS (subscription-gated)                                        │
│  [💾 Save Player] [📧 Request Contact] ← Pro+ only                         │
├─────────────────────────────────────────────────────────────────────────────┤
│  HIGHLIGHTS (FREE - visible to all)                                         │
│  ┌────────────┐ ┌────────────┐ ┌────────────┐                               │
│  │ ▶ 0:45     │ │ ▶ 0:32     │ │ ▶ 0:28     │                               │
│  │ Goal vs... │ │ Assist...  │ │ Dribble... │                               │
│  └────────────┘ └────────────┘ └────────────┘                               │
├─────────────────────────────────────────────────────────────────────────────┤
│  FULL MATCHES (PAID - Scout+ tier)                                          │
│  🔒 Upgrade to Scout tier to watch full matches                             │
│  [View Pricing]                                                              │
│  ┌────────────┐ ┌────────────┐  ← Blurred/locked thumbnails                │
│  │ 🔒 Match 1 │ │ 🔒 Match 2 │                                              │
│  └────────────┘ └────────────┘                                              │
├─────────────────────────────────────────────────────────────────────────────┤
│  CONTACT SECTION (PRO+ only)                                                │
│  "Interested in this player?"                                               │
│  [📧 Submit Contact Request] → Opens modal form                            │
└─────────────────────────────────────────────────────────────────────────────┘
```

**Access Control Logic:**
```typescript
// Visibility based on subscription tier
const canViewHighlights = true; // Everyone
const canViewFullMatches = user?.tier in ['scout', 'pro', 'club'];
const canSavePlayer = user?.tier in ['scout', 'pro', 'club'];
const canRequestContact = user?.tier in ['pro', 'club'];
```

**Data Requirements:**
- GET `/api/v1/players/:id` → `PlayerDetailDTO`
- GET `/api/v1/players/:id/videos?type=highlight` → `Video[]`
- GET `/api/v1/players/:id/videos?type=full_match` → `Video[]` (requires subscription check)

---

### 4. Pricing Page (`/pricing`)

**Purpose:** Convert free users to paid subscriptions

**Layout:**
```
┌─────────────────────────────────────────────────────────────────────────────┐
│  "Choose Your Plan"                                                          │
│  Subtitle: "Discover Africa's next football superstars"                     │
├─────────────────────────────────────────────────────────────────────────────┤
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐     │
│  │    FREE      │  │    SCOUT     │  │     PRO      │  │    CLUB      │     │
│  │              │  │   $99/mo     │  │   $299/mo    │  │   $999/mo    │     │
│  │              │  │              │  │  ⭐ Popular  │  │              │     │
│  │ • Browse     │  │ • Everything │  │ • Everything │  │ • Everything │     │
│  │ • Highlights │  │   in Free    │  │   in Scout   │  │   in Pro     │     │
│  │              │  │ • Full       │  │ • Contact    │  │ • API Access │     │
│  │              │  │   matches    │  │   requests   │  │ • Dedicated  │     │
│  │              │  │ • Save       │  │ • Priority   │  │   support    │     │
│  │              │  │   players    │  │   support    │  │              │     │
│  │              │  │              │  │              │  │              │     │
│  │  [Browse]    │  │ [Subscribe]  │  │ [Subscribe]  │  │ [Contact Us] │     │
│  └──────────────┘  └──────────────┘  └──────────────┘  └──────────────┘     │
├─────────────────────────────────────────────────────────────────────────────┤
│  FAQ Section                                                                 │
│  • How does billing work?                                                   │
│  • Can I cancel anytime?                                                    │
│  • How do contact requests work?                                            │
└─────────────────────────────────────────────────────────────────────────────┘
```

**Stripe Integration:**
- Click "Subscribe" → Create Stripe Checkout Session → Redirect to Stripe
- Webhook updates subscription status in DB

---

### 5. Scout Dashboard (`/dashboard`)

**Purpose:** Central hub for authenticated scouts

**Layout:**
```
┌─────────────────────────────────────────────────────────────────────────────┐
│  SIDEBAR                │  MAIN CONTENT                                     │
│  ┌───────────────────┐  │  ┌─────────────────────────────────────────────┐  │
│  │ 👤 Scout Name     │  │  │  Welcome back, John!                        │  │
│  │ Pro Tier ⭐       │  │  │                                             │  │
│  ├───────────────────┤  │  │  Quick Stats:                               │  │
│  │ 🏠 Dashboard      │  │  │  [12 Saved] [3 Contacts] [5 New Players]   │  │
│  │ 🔍 Discover       │  │  │                                             │  │
│  │ 💾 Saved Players  │  │  ├─────────────────────────────────────────────┤  │
│  │ 📧 My Contacts    │  │  │  Recently Viewed Players                    │  │
│  │ ⚙️ Settings       │  │  │  [Card] [Card] [Card]                       │  │
│  ├───────────────────┤  │  │                                             │  │
│  │ 🚪 Logout         │  │  │  New Players This Week                      │  │
│  └───────────────────┘  │  │  [Card] [Card] [Card]                       │  │
│                         │  └─────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

### 6. Admin Panel (`/admin`)

**Purpose:** Content management for academy staff

**Key Pages:**

| Page | Purpose | Key Actions |
|------|---------|-------------|
| `/admin` | Dashboard | Stats, recent activity |
| `/admin/players` | Player list | Search, filter, view, edit |
| `/admin/players/create` | Add player | Form with validation |
| `/admin/players/[id]/edit` | Edit player | Update profile, verify |
| `/admin/players/[id]/videos` | Manage videos | Upload, link, delete |
| `/admin/events` | Tournament list | View past/upcoming |
| `/admin/events/create` | Create event | Form with date, location |
| `/admin/videos/upload` | Bulk upload | Drag-drop, progress tracking |
| `/admin/users` | User management | View scouts, manage access |

**Admin Player Form:**
```
┌─────────────────────────────────────────────────────────────────────────────┐
│  Create Player                                                               │
├─────────────────────────────────────────────────────────────────────────────┤
│  Basic Information                                                           │
│  ┌─────────────────┐ ┌─────────────────┐ ┌─────────────────┐                │
│  │ First Name *    │ │ Last Name *     │ │ Date of Birth * │                │
│  └─────────────────┘ └─────────────────┘ └─────────────────┘                │
│                                                                              │
│  ┌─────────────────┐ ┌─────────────────┐ ┌─────────────────┐                │
│  │ Position *      │ │ Preferred Foot  │ │ Country *       │                │
│  │ [Dropdown]      │ │ [Dropdown]      │ │ [Dropdown]      │                │
│  └─────────────────┘ └─────────────────┘ └─────────────────┘                │
│                                                                              │
│  Physical Stats                                                              │
│  ┌─────────────────┐ ┌─────────────────┐                                    │
│  │ Height (cm)     │ │ Weight (kg)     │                                    │
│  └─────────────────┘ └─────────────────┘                                    │
│                                                                              │
│  Tournament Information                                                      │
│  ┌─────────────────┐ ┌─────────────────┐                                    │
│  │ Tournament Name │ │ Tournament Year │                                    │
│  └─────────────────┘ └─────────────────┘                                    │
│                                                                              │
│  Verification                                                                │
│  ┌─────────────────────────────────────────────────────────────────────┐    │
│  │ Verification Document (NIN/Passport scan) *                         │    │
│  │ [📎 Upload File] or [Drag and drop]                                 │    │
│  └─────────────────────────────────────────────────────────────────────┘    │
│  ☑️ I confirm this player's identity was verified in person                 │
│                                                                              │
│  Profile Photo                                                               │
│  ┌──────────────┐                                                            │
│  │ [Upload]     │                                                            │
│  └──────────────┘                                                            │
│                                                                              │
│  [Cancel]                                                [Create Player]    │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 🔌 API Integration Layer

### API Client Setup

```typescript
// composables/useApi.ts
export const useApi = () => {
  const config = useRuntimeConfig()
  const auth = useAuthStore()
  
  const api = $fetch.create({
    baseURL: config.public.apiBase,
    headers: {
      'Content-Type': 'application/json',
    },
    onRequest({ options }) {
      if (auth.accessToken) {
        options.headers.set('Authorization', `Bearer ${auth.accessToken}`)
      }
    },
    onResponseError({ response }) {
      if (response.status === 401) {
        auth.logout()
        navigateTo('/auth/login')
      }
    },
  })
  
  return { api }
}
```

### API Endpoints Used

| Endpoint | Method | Auth | Description |
|----------|--------|------|-------------|
| `/auth/register` | POST | No | Scout registration |
| `/auth/login` | POST | No | Login, returns JWT |
| `/auth/refresh` | POST | No | Refresh access token |
| `/auth/me` | GET | Yes | Current user info |
| `/players` | GET | No | List/search players |
| `/players/:id` | GET | No | Player detail (masked) |
| `/players/:id/videos` | GET | Tier | Player's videos |
| `/subscriptions/checkout` | POST | Yes | Create Stripe session |
| `/subscriptions/me` | GET | Yes | Current subscription |
| `/contact-requests` | POST | Pro+ | Submit contact request |
| `/saved-players` | GET/POST/DELETE | Scout+ | Manage saved players |

### Admin-Only Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/admin/players` | POST | Create player |
| `/admin/players/:id` | PUT | Update player |
| `/admin/players/:id` | DELETE | Delete player |
| `/admin/videos/upload-url` | POST | Get presigned S3 URL |
| `/admin/videos` | POST | Create video record after upload |
| `/admin/events` | POST/GET | Manage events |

---

## 🗃️ State Management (Pinia Stores)

### Auth Store

```typescript
// stores/auth.ts
interface AuthState {
  user: User | null
  accessToken: string | null
  refreshToken: string | null
}

// Actions
- login(email, password): Promise<void>
- register(data): Promise<void>
- logout(): void
- refreshAccessToken(): Promise<void>
- fetchCurrentUser(): Promise<void>

// Getters
- isAuthenticated: boolean
- isAdmin: boolean
- userTier: 'free' | 'scout' | 'pro' | 'club'
```

### Search Store

```typescript
// stores/search.ts
interface SearchState {
  query: string
  filters: {
    position: string | null
    country: string | null
    ageMin: number | null
    ageMax: number | null
    verifiedOnly: boolean
  }
  results: Player[]
  pagination: { page: number, totalPages: number, total: number }
  loading: boolean
}

// Actions
- search(): Promise<void>
- setFilter(key, value): void
- resetFilters(): void
- loadMore(): Promise<void>
```

### Subscription Store

```typescript
// stores/subscription.ts
interface SubscriptionState {
  subscription: Subscription | null
  loading: boolean
}

// Actions
- fetchSubscription(): Promise<void>
- createCheckoutSession(priceId): Promise<string> // Returns Stripe URL

// Getters
- tier: 'free' | 'scout' | 'pro' | 'club'
- canAccessFullMatches: boolean
- canRequestContact: boolean
- canSavePlayers: boolean
```

---

## 🔐 Authentication & Authorization

### Route Protection

```typescript
// middleware/auth.ts
export default defineNuxtRouteMiddleware((to) => {
  const auth = useAuthStore()
  
  if (!auth.isAuthenticated) {
    return navigateTo(`/auth/login?redirect=${to.fullPath}`)
  }
})

// middleware/admin.ts
export default defineNuxtRouteMiddleware(() => {
  const auth = useAuthStore()
  
  if (!auth.isAdmin) {
    return navigateTo('/dashboard')
  }
})

// middleware/subscription.ts
export default defineNuxtRouteMiddleware((to) => {
  const sub = useSubscriptionStore()
  
  if (to.meta.requiredTier && !sub.hasTier(to.meta.requiredTier)) {
    return navigateTo('/pricing')
  }
})
```

### Page Middleware Assignment

```vue
<!-- pages/dashboard/index.vue -->
<script setup>
definePageMeta({
  layout: 'dashboard',
  middleware: 'auth'
})
</script>

<!-- pages/admin/players/create.vue -->
<script setup>
definePageMeta({
  layout: 'admin',
  middleware: ['auth', 'admin']
})
</script>
```

---

## 📱 Responsive Breakpoints

| Breakpoint | Width | Layout |
|------------|-------|--------|
| `sm` | 640px | Mobile |
| `md` | 768px | Tablet |
| `lg` | 1024px | Desktop |
| `xl` | 1280px | Large desktop |

### Mobile-First Approach

```vue
<!-- Example: Player grid -->
<div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
  <PlayerCard v-for="player in players" :key="player.id" :player="player" />
</div>
```

---

## ⚡ Performance Optimizations

### 1. Image Optimization

```vue
<!-- Use Nuxt Image for automatic optimization -->
<NuxtImg 
  :src="player.photoUrl" 
  :alt="player.firstName"
  width="300"
  height="400"
  loading="lazy"
  format="webp"
/>
```

### 2. Code Splitting

- Automatic route-based splitting (Nuxt default)
- Lazy-load heavy components:

```vue
<LazyVideoPlayer v-if="showVideo" :src="videoUrl" />
```

### 3. API Response Caching

```typescript
// Use Nuxt's built-in caching
const { data: featuredPlayers } = await useFetch('/api/players/featured', {
  key: 'featured-players',
  getCachedData(key) {
    return nuxtApp.payload.data[key] || nuxtApp.static.data[key]
  }
})
```

---

## 🧪 Testing Strategy

### Unit Tests (Vitest)

- Component tests for UI components
- Composable tests for business logic
- Store tests for state management

### E2E Tests (Playwright)

| Test | Path | Assertions |
|------|------|------------|
| Landing page loads | `/` | Hero visible, featured players load |
| Search works | `/discover` | Filters apply, results update |
| Login flow | `/auth/login` | Redirect to dashboard |
| Subscription gate | `/players/[id]` | Full matches locked for free users |
| Admin create player | `/admin/players/create` | Form submits, player appears |

---

## 📦 Environment Variables

```bash
# .env
NUXT_PUBLIC_API_BASE=http://localhost:8080/api/v1
NUXT_PUBLIC_SITE_URL=http://localhost:3000
NUXT_PUBLIC_STRIPE_PUBLISHABLE_KEY=pk_test_xxxxx
```

---

## 🚀 Build & Deployment

### Development

```bash
npm run dev  # Start dev server on localhost:3000
```

### Production Build

```bash
npm run build    # Generate .output folder
npm run preview  # Preview production build locally
```

### Deployment Target

- **SSR Mode**: `target: 'server'` (default)
- **Hosting**: Any Node.js host (Vercel, Railway, AWS ECS)
- **Static Assets**: Served via CloudFront CDN

---

## ✅ Implementation Checklist

### Phase 1: Foundation
```
□ Project setup (Nuxt 3, Tailwind, TypeScript)
□ UI component library (UButton, UInput, UCard, etc.)
□ Layout components (Header, Footer, Sidebar)
□ API client setup
□ Auth store + login/register pages
```

### Phase 2: Public Pages
```
□ Landing page
□ Discover page with search/filters
□ Player profile page
□ Pricing page
```

### Phase 3: Scout Dashboard
```
□ Dashboard layout
□ Saved players page
□ Contact requests page
□ Settings page
□ Stripe checkout integration
```

### Phase 4: Admin Panel
```
□ Admin layout + navigation
□ Player CRUD
□ Video upload with S3 presigned URLs
□ Event management
□ User management
```

### Phase 5: Polish
```
□ Loading states
□ Error handling
□ Toast notifications
□ Mobile responsiveness
□ SEO meta tags
□ Analytics integration
```
