# 🦄 Unicorn Sport - Product Strategy Summary

> **The Single Source of Truth for Product Direction**  
> **Last Updated:** January 29, 2026

---

## 🎯 Vision

**"The fastest way for scouts to discover, verify, and connect with African football talent."**

---

## 👤 Our User: The Scout

### Primary Need (Job-to-be-Done)
> "When I need to find talented young African players for my club, I want to quickly discover, evaluate, and connect with verified prospects, so I can sign players before competitors at a good value."

### Key Insights

1. **Trust is Everything** - Scouts have been burned by fake ages and unreliable sources
2. **Time is Money** - They spend 55% of time on admin, not actual scouting
3. **Video is the Hook** - Emotional "wow" moments drive interest
4. **Connection is the Sale** - Ability to contact players is why they pay
5. **Fresh Content Matters** - Stale database = cancelled subscription

---

## 🏗️ Product Architecture

### The Hybrid Discovery Model

```
┌─────────────────────────────────────────────────────────────────┐
│                                                                  │
│   🎬 DISCOVERY MODE              📊 RESEARCH MODE                │
│   (TikTok Video Feed)            (Grid/List View)               │
│                                                                  │
│   • Emotional engagement         • Logical evaluation           │
│   • "Wow" moments                • Filter & sort                │
│   • Quick scanning               • Compare players              │
│   • Passive browsing             • Detailed analysis            │
│                                                                  │
│                    ↓ SEAMLESS TRANSITION ↓                      │
│                                                                  │
│   ❤️ INTEREST                                                    │
│   Save player + Add notes                                       │
│                                                                  │
│                    ↓                                            │
│                                                                  │
│   📩 ACTION                                                      │
│   Contact academy + Track response                              │
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

### Why TikTok Feed Works (Keep It!)

| Benefit | Scout Value |
|---------|-------------|
| Immediate visual impact | See skills instantly |
| "Wow" factor | Creates excitement about talent |
| Low cognitive load | Passive discovery |
| Mobile-first | Works on-the-go |
| Unique differentiator | No one else has this |

### What We Add to Make It Complete

| Enhancement | Why |
|-------------|-----|
| Player info overlay | Context without leaving video |
| Quick-save button | One-tap to shortlist |
| View Profile button | Transition to research mode |
| Filter toggle | Narrow down while browsing |

---

## 📋 The Complete Scout Workflow

### Current Flow (What We Have)
```
Login → Dashboard (Video Feed) → Player Profile → Save → Contact
                                       ↑
                              Discover (Grid) 
```

### Optimized Flow (What We're Building)
```
Login → Discover (Default Entry)
           ├── Grid View (Research)
           ├── Video Feed (Discovery) ← Enhanced with overlays
           └── List View (Compare) ← New
                      ↓
              Player Profile
                      ↓
           Save + Notes ← Enhanced
                      ↓
           Compare (if multiple saved) ← New
                      ↓
           Contact Academy ← Enhanced with academy info
                      ↓
           Track in Activity Dashboard
```

---

## 🔢 Feature Prioritization

### P0: Launch Blockers (This Week)

| # | Feature | Impact | Effort |
|---|---------|--------|--------|
| 1 | Redirect auth users from homepage | High | Low |
| 2 | Default post-login to /discover | High | Low |
| 3 | Notes field on save | High | Medium |
| 4 | Display notes on saved page | High | Low |
| 5 | Academy filter on discover | High | Medium |
| 6 | Public academies API | High | Medium |

### P1: Core Experience (Next 2 Weeks)

| # | Feature | Impact | Effort |
|---|---------|--------|--------|
| 7 | Video overlay with player info | High | Medium |
| 8 | Quick-save from video feed | High | Low |
| 9 | Academy info on contact request | High | Low |
| 10 | Sorting options on discover | Medium | Low |
| 11 | Similar players on profile | Medium | High |
| 12 | Player comparison tool | Medium | High |

### P2: Enhancement (Month 2)

| # | Feature | Impact | Effort |
|---|---------|--------|--------|
| 13 | Academy profile pages | Medium | Medium |
| 14 | Scout profile/verification | Medium | High |
| 15 | Email notifications | Medium | High |
| 16 | PDF export | Low | Medium |
| 17 | Activity dashboard | Low | Medium |

---

## 💰 Subscription Strategy

### Recommended Tiers

| Tier | Price | Key Value | Target User |
|------|-------|-----------|-------------|
| **Free** | $0 | Browse + highlights + 5 saves | Curious browsers |
| **Scout** | $79/mo | Full matches + unlimited saves + notes | Active scouts |
| **Pro** | $149/mo | Contact players + comparison + PDF | Serious scouts |
| **Club** | $399/mo | Team access (5 seats) + collaboration | Organizations |

### Conversion Triggers

```
FREE → SCOUT
• Hit 5 saved players limit
• Try to watch full match
• Try to add notes

SCOUT → PRO
• Try to contact player
• Need comparison tool
• Need to share reports

PRO → CLUB
• Need colleague access
• Need shared shortlists
• Need API integration
```

---

## 📊 Success Metrics

### North Star Metric
> **Contact Requests per Active Scout per Month**

### Supporting Metrics

| Metric | Target | Why It Matters |
|--------|--------|----------------|
| Time to first save | < 5 min | Discovery works |
| Save rate | > 5% of views | Content is quality |
| Contact rate | > 20% of saves | Workflow complete |
| Response rate | > 70% in 72 hrs | Academies engaged |
| Weekly return | > 60% | Ongoing value |
| Free → Paid | > 10% | Clear value prop |
| Monthly churn | < 5% | Delivering value |

---

## 🏆 Competitive Advantage

### Our Moat

1. **Africa Focus** - Deep where others are shallow
2. **Verified Ages** - Trust no one else provides
3. **Academy Direct** - Cut out unreliable agents
4. **Complete Workflow** - Discover → Evaluate → Contact
5. **Video-First** - Engaging, modern experience

### Positioning Statement

> **For professional scouts** who need to discover African talent,  
> **Unicorn Sport** is a scouting platform that  
> **provides verified players, HD video, and direct academy contact**  
> **Unlike** Wyscout (European focus), YouTube (unverified), or agents (biased),  
> **We** are the trusted source for African football talent.

---

## 🚀 30-Day Roadmap

### Week 1: Foundation Fixes
- [ ] Auth redirect from homepage
- [ ] Post-login → /discover
- [ ] Notes on save
- [ ] Academy filter
- [ ] Public academies API

### Week 2: Video Feed Enhancement
- [ ] Player info overlay
- [ ] Quick-save button
- [ ] View Profile button
- [ ] Academy info on contact modal

### Week 3: Research Tools
- [ ] Sorting options
- [ ] Player comparison (2-3 players)
- [ ] Similar players section

### Week 4: Trust & Polish
- [ ] Academy profile pages
- [ ] Scout profile section
- [ ] Contact request improvements
- [ ] Testing & bug fixes

---

## 📚 Related Documents

1. **[UX_SCOUT_FLOW_ANALYSIS.md](./UX_SCOUT_FLOW_ANALYSIS.md)** - Detailed UX audit and issues
2. **[SCOUT_PERSONA_AND_JTBD.md](./SCOUT_PERSONA_AND_JTBD.md)** - User research and jobs-to-be-done
3. **[FEATURE_GAP_ANALYSIS.md](./FEATURE_GAP_ANALYSIS.md)** - What exists vs what's needed

---

## ❓ Open Questions

1. **Verification Partnership** - Who verifies player ages? MRI clinics?
2. **Academy Onboarding** - How do we get academies on the platform?
3. **Content Pipeline** - How often do we need new players/videos?
4. **Agent Policy** - Do we allow agents, or academies only?
5. **Geographic Priority** - Which African countries first?

---

## ✅ Next Action

**Start with P0 features (Week 1):**

1. Create `middleware/redirect-authenticated.ts`
2. Update auth store login redirect
3. Add notes modal to player save
4. Display notes on saved players page
5. Add academy filter dropdown
6. Create public academies endpoint

**Ready to implement? Start here ↑**

---

*This is the master strategy document. All product decisions should align with this.*  
*Review and update weekly.*
