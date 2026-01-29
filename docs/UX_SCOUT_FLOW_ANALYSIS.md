# 🦄 Unicorn Sport - Scout UX Flow Analysis

> **Document Purpose:** Comprehensive UX audit and recommendations for the scout user experience  
> **Date:** January 29, 2026  
> **Status:** Action Required

---

## 📋 Table of Contents

1. [Executive Summary](#executive-summary)
2. [Current User Flow Map](#current-user-flow-map)
3. [Critical Issues (P0)](#critical-issues-p0)
4. [High Priority Issues (P1)](#high-priority-issues-p1)
5. [Medium Priority Issues (P2)](#medium-priority-issues-p2)
6. [TikTok Video Feed - Analysis & Recommendations](#tiktok-video-feed---analysis--recommendations)
7. [Missing Features](#missing-features)
8. [Subscription Tier Recommendations](#subscription-tier-recommendations)
9. [Implementation Checklist](#implementation-checklist)
10. [What We're Doing Right](#what-were-doing-right)

---

## Executive Summary

### The Core Challenge
We're designing for **entertainment patterns (TikTok)** when scouts need **productivity tools (LinkedIn Recruiter)**. However, the TikTok-style video feed has merit for immediate visual impact.

### Recommended Approach: Hybrid Model
- **TikTok for Discovery:** Quick visual scanning, emotional "wow" moments
- **LinkedIn for Action:** Research, compare, shortlist, contact

### Key Metrics to Optimize
1. Time to find relevant player: < 2 minutes
2. Contact request success rate: > 70%
3. Scout return rate: Weekly active usage
4. Conversion: Free → Paid tier

---

## Current User Flow Map

```
┌─────────────────────────────────────────────────────────────────┐
│                        CURRENT FLOW                              │
└─────────────────────────────────────────────────────────────────┘

[Homepage] ──login──> [Dashboard/Video Feed] ──> [Discover]
     │                        │                       │
     │                        ▼                       ▼
     │                  [Player Profile] ◄──── [Player Card]
     │                        │
     │                        ▼
     │              [Save] or [Contact Request]
     │                        │
     │                        ▼
     │                  [Saved Players] ──> [Contacts Page]
     │
     └──── ⚠️ ISSUE: Authenticated users still see marketing page

┌─────────────────────────────────────────────────────────────────┐
│                      RECOMMENDED FLOW                            │
└─────────────────────────────────────────────────────────────────┘

[Homepage] ──if authenticated──> [Discover Page] (default entry)
                                       │
                        ┌──────────────┼──────────────┐
                        ▼              ▼              ▼
                 [Video Feed]   [Grid View]    [List View]
                 "Wow moments"   "Research"    "Compare"
                        │              │              │
                        └──────────────┼──────────────┘
                                       ▼
                              [Player Profile]
                                       │
                        ┌──────────────┼──────────────┐
                        ▼              ▼              ▼
                    [Save+Notes]  [Compare]    [Contact]
                        │              │              │
                        └──────────────┼──────────────┘
                                       ▼
                              [Activity Dashboard]
                              (track all actions)
```

---

## Critical Issues (P0)

### 🚨 P0-1: Homepage Shows Marketing to Authenticated Users

**Problem:**  
After login, scouts still see "Start Free Trial" and "Explore Players" CTAs. This creates cognitive dissonance and wastes time.

**Impact:** High - Confusion, wasted clicks, poor first impression post-login

**Solution:**
```typescript
// web/middleware/redirect-authenticated.ts
export default defineNuxtRouteMiddleware(async () => {
  const authStore = useAuthStore()
  if (!authStore.initialized) {
    await authStore.initFromStorage()
  }
  if (authStore.isAuthenticated) {
    return navigateTo('/discover')
  }
})
```

Apply to `pages/index.vue`:
```typescript
definePageMeta({
  middleware: 'redirect-authenticated'
})
```

**Status:** [ ] Not Started

---

### 🚨 P0-2: Wrong Default Entry Point Post-Login

**Problem:**  
Dashboard (video feed) is the default, but scouts need to filter/search first.

**Impact:** High - Scouts can't immediately start researching

**Current:** Login → `/dashboard` (video feed)  
**Recommended:** Login → `/discover` (search/filter page)

**Solution:**
Update `auth.ts` store login redirect:
```typescript
// After successful login
router.push('/discover')
```

**Status:** [ ] Not Started

---

### 🚨 P0-3: No Notes Field for Saved Players

**Problem:**  
Scouts save players but can't record WHY they saved them or any observations.

**Impact:** High - Saved list becomes useless without context

**Solution:**
1. Add `notes` field to SavedPlayer API/DB (already exists in schema)
2. Add notes input in UI when saving
3. Show/edit notes in Saved Players list

**Files to update:**
- `web/pages/players/[id].vue` - Add notes modal when saving
- `web/pages/dashboard/saved.vue` - Show and edit notes
- Backend already supports notes field

**Status:** [ ] Not Started

---

### 🚨 P0-4: Contact Request Lacks Academy Context

**Problem:**  
Scouts request contact but don't know:
- Which academy receives the request
- Expected response time
- Success rates

**Impact:** High - Low confidence in the system, hesitation to pay for Pro tier

**Solution:**
Show on player profile before contact:
```vue
<div class="academy-contact-card">
  <h4>Contact through: {{ player.academy_name }}</h4>
  <p>Average response time: 48-72 hours</p>
  <p>Response rate: 85%</p>
</div>
```

**Status:** [ ] Not Started

---

### 🚨 P0-5: No Academy Filter on Discover Page

**Problem:**  
Scouts often want players from specific academies (reputation matters).

**Impact:** Medium-High - Missing key research capability

**Solution:**
Add academy dropdown filter alongside position/country/age filters.

**Status:** [ ] Not Started

---

## High Priority Issues (P1)

### P1-1: No Player Comparison Tool

**Problem:**  
Scouts can't compare 2-3 players side-by-side (age, height, position, videos).

**Impact:** Medium - Scouts do this manually, poor experience

**Solution:**
1. Add "Compare" checkbox on player cards
2. Floating "Compare (2)" button when 2+ selected
3. Comparison modal with stats table + video thumbnails

**Status:** [ ] Not Started

---

### P1-2: No Sorting Options on Discover

**Problem:**  
Can't sort players by: Newest, Age, Height, etc.

**Impact:** Medium - Limits research efficiency

**Solution:**
Add sort dropdown: "Sort by: Newest | Age (low-high) | Age (high-low) | Position"

**Status:** [ ] Not Started

---

### P1-3: No Scout Profile/Verification

**Problem:**  
Anyone can sign up as "scout" - no trust verification.

**Risk:**
- Players/academies receive spam
- No accountability
- Low response rates

**Solution:**
Scout registration should collect:
- Organization name
- Professional email (club domain)
- LinkedIn profile (optional)
- License/credential (optional)

Display verified badge to academies.

**Status:** [ ] Not Started

---

### P1-4: No Similar Players Suggestion

**Problem:**  
After viewing a player, no "You might also like" suggestions.

**Impact:** Medium - Missed discovery opportunities

**Solution:**
On player profile, show:
- Same position, similar age
- Same academy
- Same country
- "Scouts who viewed this also viewed..."

**Status:** [ ] Not Started

---

### P1-5: Contact Request Flow Incomplete

**Problem:**  
Status shows "pending, forwarded, responded" but scouts can't see the actual response.

**Impact:** High - Broken communication loop

**Solution:**
1. Enable two-way messaging
2. Show response message from academy
3. Add follow-up mechanism

**Status:** [ ] Not Started

---

## Medium Priority Issues (P2)

### P2-1: No Email Notifications

**Problem:**  
No alerts for: new players matching criteria, contact responses, etc.

**Solution:**
- Weekly digest of new talent
- Instant notification for contact responses
- Activity summary

**Status:** [ ] Not Started

---

### P2-2: No Export/Download Functionality

**Problem:**  
Scouts need to share with managers/colleagues. Can't export:
- Shortlist
- Player profiles
- Comparison reports

**Solution:**
- PDF player profile export
- CSV shortlist export
- Shareable links

**Status:** [ ] Not Started

---

### P2-3: No Academy Profile Pages

**Problem:**  
Can't browse by academy → see all their players.

**Solution:**
Create `/academies/[id]` page with:
- Academy info
- Player roster
- Contact information

**Status:** [ ] Not Started

---

### P2-4: No Search History

**Problem:**  
"What did I search for last week?" - No way to know.

**Solution:**
- Recent searches dropdown
- Saved search filters

**Status:** [ ] Not Started

---

### P2-5: No Activity Dashboard

**Problem:**  
Dashboard is video feed, not activity overview.

**Solution:**
Add real dashboard with:
- Recent activity
- Saved players count
- Pending contact requests
- New matches for saved searches

**Status:** [ ] Not Started

---

## TikTok Video Feed - Analysis & Recommendations

### ✅ Why the TikTok Approach Works

| Benefit | Explanation |
|---------|-------------|
| **Immediate Impact** | Scouts see player skills instantly - the "wow" moment |
| **Emotional Engagement** | Videos create excitement about talent |
| **Quick Scanning** | Swipe through many players fast |
| **Mobile-First** | Works great for on-the-go browsing |
| **Differentiation** | Unique in the market vs. spreadsheet tools |

### ⚠️ Where It Falls Short

| Issue | Impact |
|-------|--------|
| **No Memory** | Can't remember who they saw 10 swipes ago |
| **No Context Overlay** | Missing: age, academy, position on video |
| **Can't Compare** | Viewing sequentially, not side-by-side |
| **Can't Filter While Watching** | Need to exit feed to apply filters |
| **No Quick-Save** | Must stop, go to profile, then save |

### 🎯 Recommended Hybrid Approach

**Keep the video feed BUT enhance it:**

#### Enhancement 1: Rich Video Overlay
```
┌────────────────────────────────────────┐
│                                        │
│           [VIDEO PLAYING]              │
│                                        │
│  ┌─────────────────────────────────┐   │
│  │ 🏃 John M. • Striker • 17yo     │   │
│  │ 📍 Lagos, Nigeria               │   │
│  │ 🏟️ Rising Stars Academy         │   │
│  └─────────────────────────────────┘   │
│                                        │
│  [❤️ Save]  [👤 Profile]  [📩 Contact] │
└────────────────────────────────────────┘
```

#### Enhancement 2: Quick Actions While Watching
- ❤️ Double-tap to save (like Instagram)
- Swipe right = Save to shortlist
- Long-press = Quick profile preview
- Button overlay for "View Full Profile"

#### Enhancement 3: Filter Toggle in Feed
Add a minimal filter bar at top of video feed:
```
[All] [Strikers] [Midfielders] [18-21] [Nigeria ▼]
```

#### Enhancement 4: Make Feed ONE of Multiple Views

```
[📹 Feed]  [📊 Grid]  [📋 List]
     ↓          ↓          ↓
  TikTok    Card View   Table View
  Style     + Filters   + Sort/Export
```

Let scouts choose their preferred workflow.

---

## Missing Features

### Critical (Must Have for Launch)
- [ ] Notes on saved players
- [ ] Academy filter
- [ ] Video overlay with player info
- [ ] Quick-save from video feed
- [ ] Contact request academy details

### Important (Soon After Launch)
- [ ] Player comparison tool
- [ ] Sorting options
- [ ] Similar players suggestions
- [ ] Scout verification/profile
- [ ] Two-way messaging

### Nice to Have (Future Roadmap)
- [ ] PDF export
- [ ] Email notifications
- [ ] Academy profile pages
- [ ] Search history
- [ ] Team collaboration (Club tier)
- [ ] API access for integrations

---

## Subscription Tier Recommendations

### Current Tiers (Issues)
| Tier | Price | Problem |
|------|-------|---------|
| Free | $0 | Can't save = frustrating trial |
| Scout | $99/mo | Can't contact = incomplete workflow |
| Pro | $199/mo | Expensive for individual scouts |
| Club | $399/mo | Value unclear |

### Recommended Tiers

#### 🆓 EXPLORER (Free)
- Browse all players
- Watch highlights only
- Save up to 5 players (with notes)
- Basic filters (position, country)

#### 🔍 SCOUT ($79/mo or $59/mo annual)
- Everything in Free
- Full match recordings
- Unlimited saves with notes
- Advanced filters (academy, tournament, age range)
- Download player reports (PDF)
- Email notifications for new talent

#### ⭐ PRO ($149/mo or $119/mo annual)
- Everything in Scout
- Contact players (unlimited)
- Comparison tools (up to 5 players)
- Priority support
- Analytics dashboard (views, response rates)

#### 🏢 CLUB ($399/mo)
- Everything in Pro
- 5 scout accounts
- Shared shortlists
- Team collaboration tools
- Dedicated account manager
- Custom API integrations

---

## Implementation Checklist

### Phase 1: P0 Fixes (This Week)
- [ ] Add redirect-authenticated middleware to homepage
- [ ] Change post-login redirect to /discover
- [ ] Add notes field to save player flow
- [ ] Show academy info on contact request
- [ ] Add academy filter to discover

### Phase 2: Video Feed Enhancements (Next Week)
- [ ] Add player info overlay on videos
- [ ] Add quick-save button (heart icon)
- [ ] Add "View Profile" button overlay
- [ ] Add minimal filter bar to feed

### Phase 3: Research Tools (Week 3)
- [ ] Player comparison feature
- [ ] Sorting options on discover
- [ ] Similar players suggestions
- [ ] Improve saved players page with notes display

### Phase 4: Communication (Week 4)
- [ ] Enhanced contact request flow
- [ ] Response display
- [ ] Scout profile/verification

### Phase 5: Polish (Month 2)
- [ ] Email notifications
- [ ] Export functionality
- [ ] Academy profile pages
- [ ] Activity dashboard

---

## What We're Doing Right

### ✅ Strengths to Preserve

| Feature | Why It Works |
|---------|--------------|
| **Verified Players Only** | Huge trust builder - differentiator |
| **Clean, Modern Design** | Professional, credible platform |
| **Academy Linkage** | Critical for contact and trust |
| **HD Video Quality** | Shows player skills clearly |
| **Mobile-Responsive** | Accessibility for on-the-go |
| **Subscription Model** | Sustainable business model |
| **TikTok-Style Feed** | Unique, engaging discovery method |
| **African Focus** | Niche market with real demand |

### 🎯 Core Value Proposition
> "The fastest way to discover and connect with verified African football talent"

Keep this front and center in all UX decisions.

---

## Summary: The Golden Rule

**Scouts have TWO modes:**

1. **Discovery Mode** 🎬
   - "Show me impressive players"
   - TikTok feed is PERFECT here
   - Emotional, visual, quick

2. **Research Mode** 📊
   - "Help me evaluate and decide"
   - Need filters, comparison, data
   - Logical, detailed, thorough

**Our job:** Support BOTH modes seamlessly.

```
DISCOVERY (Video Feed) → INTEREST (Quick Save) → RESEARCH (Profile/Compare) → ACTION (Contact)
```

Don't force one mode. Let scouts flow naturally between them.

---

## Questions for Discussion

1. Should video feed show ALL players or respect filters?
2. Do we need separate "Hot This Week" curated feed?
3. Should free tier allow 5 saves or 0 saves?
4. Do we verify scouts before they can contact?
5. What's our response time SLA with academies?

---

*Document maintained by: Engineering Team*  
*Last updated: January 29, 2026*
