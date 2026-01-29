# 🔍 Feature Gap Analysis: Current State vs. Ideal State

> **Purpose:** Identify exactly what exists and what's missing  
> **Date:** January 29, 2026

---

## Current Features Audit

### ✅ What We Have (Backend + Frontend)

| Feature | Backend | Frontend | Notes |
|---------|---------|----------|-------|
| **Player Management** | ✅ | ✅ | Full CRUD, admin only |
| **Player Listing (Public)** | ✅ | ✅ | With filters |
| **Player Profiles** | ✅ | ✅ | Good detail |
| **Academy Management** | ✅ | ✅ | Admin only |
| **Video Upload** | ✅ | ✅ | Admin only |
| **User Auth (Login/Register)** | ✅ | ✅ | Complete |
| **Subscription Tiers** | ✅ | ✅ | Stripe integrated |
| **Save Players** | ✅ | ✅ | Works |
| **Contact Requests** | ✅ | ✅ | Basic |
| **Verification Status** | ✅ | ✅ | Badge shows |
| **Video Feed (TikTok style)** | ✅ | ✅ | Dashboard |
| **Discover Page (Grid)** | ✅ | ✅ | Good filters |
| **Position Filter** | ✅ | ✅ | Works |
| **Country Filter** | ✅ | ✅ | Works |
| **Age Filter** | ✅ | ✅ | Min/Max |
| **Search by Name** | ✅ | ✅ | Works |

### ⚠️ Exists but Needs Enhancement

| Feature | Issue | Enhancement Needed |
|---------|-------|-------------------|
| **Save Players** | No notes field in UI | Add notes input |
| **Contact Requests** | No academy info shown | Show academy details |
| **Video Feed** | No overlay info | Add player info overlay |
| **Video Feed** | No quick save | Add heart button |
| **Saved Players Page** | Notes exist in DB but not shown | Display notes |
| **Player Profile** | No similar players | Add recommendations |

### ❌ Missing Features

| Feature | Priority | Complexity | Required For |
|---------|----------|------------|--------------|
| **Academy Filter (Discover)** | P0 | Low | Better research |
| **Redirect Auth Users from Home** | P0 | Low | Better flow |
| **Player Comparison Tool** | P1 | Medium | Decision making |
| **Sorting Options** | P1 | Low | Research efficiency |
| **Scout Profile/Verification** | P1 | Medium | Trust |
| **PDF Export** | P2 | Medium | Sharing with team |
| **Email Notifications** | P2 | High | Engagement |
| **Academy Public Pages** | P2 | Medium | Discovery |
| **Search History** | P2 | Low | Convenience |
| **Activity Dashboard** | P2 | Medium | Overview |
| **Team/Shared Shortlists** | P3 | High | Club tier |

---

## API Gap Analysis

### Current API Endpoints

```
AUTH
  POST /auth/login
  POST /auth/register  
  POST /auth/forgot-password
  POST /auth/refresh
  POST /auth/logout

PLAYERS (Public)
  GET  /players           ← List with filters
  GET  /players/:id       ← Single player

PLAYERS (Admin)
  POST   /admin/players
  PUT    /admin/players/:id
  DELETE /admin/players/:id
  POST   /admin/players/:id/verify

ACADEMIES (Admin Only - ISSUE!)
  GET    /admin/academies      ← Should have public version
  GET    /admin/academies/:id
  POST   /admin/academies
  PUT    /admin/academies/:id
  DELETE /admin/academies/:id

SAVED PLAYERS
  GET    /users/me/saved-players
  POST   /users/me/saved-players
  DELETE /users/me/saved-players/:id

CONTACTS
  GET    /users/me/contacts
  POST   /contacts

SUBSCRIPTIONS
  GET  /subscriptions/plans
  POST /subscriptions/subscribe
  POST /subscriptions/cancel
```

### Missing API Endpoints

```
ACADEMIES (Public - NEEDED)
  GET /academies              ← List public academies for filter
  GET /academies/:id          ← Public academy profile

SAVED PLAYERS
  PUT /users/me/saved-players/:id   ← Update notes

PLAYERS (Enhanced)
  GET /players/:id/similar    ← Similar players

HIGHLIGHTS/VIDEOS (For Feed)
  GET /highlights             ← For video feed with player info

NOTIFICATIONS
  GET  /users/me/notifications
  POST /users/me/notifications/:id/read

SCOUT PROFILE
  GET  /users/me/profile
  PUT  /users/me/profile
  POST /users/me/verify-scout
```

---

## Frontend Component Gap Analysis

### Pages Needed

| Page | Exists | Status |
|------|--------|--------|
| `/` | ✅ | Needs auth redirect |
| `/discover` | ✅ | Needs academy filter, sorting |
| `/dashboard` | ✅ | Needs video overlay enhancement |
| `/dashboard/saved` | ✅ | Needs notes display |
| `/dashboard/contacts` | ✅ | Works |
| `/dashboard/settings` | ✅ | Works |
| `/players/[id]` | ✅ | Needs similar players |
| `/pricing` | ✅ | Works |
| `/auth/*` | ✅ | Works |
| `/academies` | ❌ | NEW - Academy directory |
| `/academies/[id]` | ❌ | NEW - Academy profile |
| `/compare` | ❌ | NEW - Comparison view |

### Components Needed

| Component | Exists | Notes |
|-----------|--------|-------|
| `PlayerCard.vue` | ✅ | Good |
| `VideoPlayerModal.vue` | ✅ | Good |
| `PlayerCompare.vue` | ❌ | NEW |
| `AcademyCard.vue` | ❌ | NEW |
| `SaveWithNotes.vue` | ❌ | NEW - Modal for save + notes |
| `VideoOverlay.vue` | ❌ | NEW - Info overlay for feed |
| `NotificationBell.vue` | ❌ | Future |

---

## Data Model Completeness

### Player Model - ✅ Complete
```typescript
interface Player {
  id: string
  user_id?: string
  first_name: string
  last_name: string
  date_of_birth: string
  position: string
  preferred_foot?: 'left' | 'right' | 'both'
  height_cm?: number
  weight_kg?: number
  country: string
  state?: string
  city?: string
  school_name?: string
  verification_status: 'pending' | 'verified' | 'rejected'
  profile_photo_url?: string
  thumbnail_url?: string
  tournament_id?: string
  tournament_year?: number
  academy_id?: string          // ✅ Has academy link
  academy_name?: string        // ✅ Has academy name
  video_count?: number
  created_at: string
  updated_at: string
  videos?: Video[]
}
```

### SavedPlayer Model - ⚠️ Has notes but UI doesn't use it
```typescript
interface SavedPlayer {
  id: string
  user_id: string
  player_id: string
  notes?: string               // ✅ Exists in backend
  created_at: string
  player?: Player
}
```

### Academy Model - ✅ Complete but needs public API
```typescript
interface Academy {
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
```

---

## Implementation Priority Queue

### 🔴 PHASE 1: Quick Wins (This Week)

| Task | File(s) | Effort |
|------|---------|--------|
| 1. Auth redirect from homepage | `middleware/redirect-authenticated.ts`, `pages/index.vue` | 30 min |
| 2. Change post-login to /discover | `stores/auth.ts` | 15 min |
| 3. Add notes modal when saving | `pages/players/[id].vue` | 2 hrs |
| 4. Show notes on saved page | `pages/dashboard/saved.vue` | 1 hr |
| 5. Add academy filter to discover | `pages/discover.vue` | 2 hrs |
| 6. Create public academies API | `backend` | 2 hrs |

### 🟡 PHASE 2: Video Feed Enhancement (Next Week)

| Task | File(s) | Effort |
|------|---------|--------|
| 1. Create VideoOverlay component | `components/VideoOverlay.vue` | 3 hrs |
| 2. Add overlay to dashboard feed | `pages/dashboard/index.vue` | 2 hrs |
| 3. Add quick-save button | `pages/dashboard/index.vue` | 1 hr |
| 4. Add "View Profile" button | `pages/dashboard/index.vue` | 1 hr |
| 5. Show academy info on contact | `pages/players/[id].vue` | 2 hrs |

### 🟢 PHASE 3: Research Tools (Week 3)

| Task | File(s) | Effort |
|------|---------|--------|
| 1. Add sorting dropdown | `pages/discover.vue` | 2 hrs |
| 2. Create comparison page | `pages/compare.vue` | 8 hrs |
| 3. Add compare checkboxes | `components/PlayerCard.vue` | 2 hrs |
| 4. Similar players section | `pages/players/[id].vue` | 4 hrs |

### 🔵 PHASE 4: Academy & Trust (Week 4)

| Task | File(s) | Effort |
|------|---------|--------|
| 1. Academy listing page | `pages/academies/index.vue` | 4 hrs |
| 2. Academy profile page | `pages/academies/[id].vue` | 6 hrs |
| 3. Scout profile section | `pages/dashboard/settings.vue` | 4 hrs |

---

## Definition of Done Checklist

Before marking any feature complete:

### Functionality
- [ ] Works on desktop (Chrome, Firefox, Safari)
- [ ] Works on mobile (iOS Safari, Android Chrome)
- [ ] Handles loading states
- [ ] Handles error states
- [ ] Handles empty states

### UX
- [ ] Follows existing design system
- [ ] Has appropriate feedback (toasts, transitions)
- [ ] Is accessible (keyboard navigation, screen reader)
- [ ] Matches scout workflow expectations

### Quality
- [ ] No TypeScript errors
- [ ] No console errors
- [ ] Tested happy path
- [ ] Tested edge cases

---

*This document tracks what exists vs what we need.*  
*Update as features are completed.*
