# 🎯 Scout User Research: Personas, Jobs-to-be-Done & Algorithm Design

> **Purpose:** Deep understanding of our users before building  
> **Framework:** Google UX Research + Jobs-to-be-Done Theory  
> **Date:** January 29, 2026

---

## Part 1: Who Are Our Users?

### Primary Persona: The Professional Scout

```
┌─────────────────────────────────────────────────────────────────┐
│  👤 PERSONA: Marcus Thompson                                     │
│  Role: Youth Scout, Premier League Club                         │
│  Age: 38 | Location: London, UK | Experience: 12 years          │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  📋 RESPONSIBILITIES                                             │
│  • Identify talent for U18/U21/First team pipeline              │
│  • Watch 50+ players per week (live + video)                    │
│  • Write scouting reports for recruitment committee             │
│  • Build relationships with academies & agents                  │
│  • Travel to tournaments (Africa, Europe, South America)        │
│                                                                  │
│  🎯 SUCCESS METRICS (How his boss measures him)                  │
│  • Players signed who make it to first team                     │
│  • Cost efficiency (find gems before they're expensive)         │
│  • Speed (find players before competitors)                      │
│  • Accuracy (low rejection rate from technical staff)           │
│                                                                  │
│  😤 CURRENT FRUSTRATIONS                                         │
│  • "I spend 70% of my time on logistics, 30% on actual scouting"│
│  • "Videos from agents are always cherry-picked highlights"     │
│  • "I can't trust player ages - seen too many fake documents"   │
│  • "Flying to Africa for 1 tournament is expensive"             │
│  • "By the time I find a player, 5 other clubs know about him"  │
│                                                                  │
│  💰 BUDGET                                                       │
│  • Club pays for tools (Wyscout, InStat, TransferMarkt)         │
│  • Personal budget: Would pay $100-200/mo for good tool         │
│  • Club budget: $500-2000/mo for team tools                     │
│                                                                  │
│  🛠️ CURRENT TOOLS                                                │
│  • Wyscout ($$$) - European focus, weak on Africa               │
│  • InStat - Match analysis, not discovery                       │
│  • YouTube - Unreliable, no verification                        │
│  • WhatsApp groups - Agents sharing videos                      │
│  • Excel - Personal player database                             │
│  • Fly to tournaments - Expensive but necessary                 │
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

### Secondary Persona: Independent Scout / Agent

```
┌─────────────────────────────────────────────────────────────────┐
│  👤 PERSONA: David Mensah                                        │
│  Role: Independent Scout & Agent                                │
│  Age: 45 | Location: Accra, Ghana | Experience: 20 years        │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  📋 RESPONSIBILITIES                                             │
│  • Find talented players in local academies                     │
│  • Connect players with European clubs                          │
│  • Negotiate trials and transfers                               │
│  • Manage player documentation                                  │
│                                                                  │
│  🎯 SUCCESS METRICS                                              │
│  • Number of players placed in European clubs                   │
│  • Commission from transfers                                    │
│  • Reputation in the industry                                   │
│                                                                  │
│  😤 CURRENT FRUSTRATIONS                                         │
│  • "European scouts don't trust videos I send"                  │
│  • "I can't prove player ages to skeptical clubs"               │
│  • "No central platform to showcase my players"                 │
│  • "Communication with European clubs is fragmented"            │
│                                                                  │
│  💰 BUDGET                                                       │
│  • Limited personal budget: $50-100/mo maximum                  │
│  • Success-based: Would pay more if it leads to deals           │
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

### Tertiary Persona: Academy Director (Supply Side)

```
┌─────────────────────────────────────────────────────────────────┐
│  👤 PERSONA: Coach Emmanuel Okonkwo                              │
│  Role: Academy Director                                         │
│  Age: 52 | Location: Lagos, Nigeria                             │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  📋 RESPONSIBILITIES                                             │
│  • Develop young players (ages 10-18)                           │
│  • Place graduates in professional clubs                        │
│  • Generate revenue for academy sustainability                  │
│                                                                  │
│  🎯 SUCCESS METRICS                                              │
│  • Players signed by professional clubs                         │
│  • Transfer fees / training compensation                        │
│  • Academy reputation                                           │
│                                                                  │
│  😤 CURRENT FRUSTRATIONS                                         │
│  • "European scouts don't know we exist"                        │
│  • "No way to verify our legitimacy to scouts"                  │
│  • "Agents take advantage of our players"                       │
│  • "We have great talent but no visibility"                     │
│                                                                  │
│  💡 WHAT THEY WANT FROM US                                       │
│  • Visibility to European scouts                                │
│  • Verified academy badge (legitimacy)                          │
│  • Direct communication with scouts (cut out middlemen)         │
│  • Fair compensation when players are signed                    │
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

---

## Part 2: Jobs-to-be-Done (JTBD) Framework

### Core Job Statement

> **When** I need to find talented young African players for my club,  
> **I want to** quickly discover, evaluate, and connect with verified prospects,  
> **So that** I can sign players before competitors at a good value.

### Functional Jobs (What they need to DO)

| Job # | Job Statement | Importance | Satisfaction (Current Tools) |
|-------|--------------|------------|------------------------------|
| F1 | Discover new players I haven't seen before | 🔴 Critical | ⭐⭐ Low |
| F2 | Filter players by position, age, physical attributes | 🔴 Critical | ⭐⭐⭐ Medium |
| F3 | Watch video evidence of player ability | 🔴 Critical | ⭐⭐ Low (trust issues) |
| F4 | Verify player age and identity | 🔴 Critical | ⭐ Very Low |
| F5 | Compare multiple players side-by-side | 🟡 Important | ⭐⭐ Low |
| F6 | Save players to a shortlist | 🟡 Important | ⭐⭐⭐ Medium (use Excel) |
| F7 | Contact player's academy/representative | 🔴 Critical | ⭐⭐ Low (fragmented) |
| F8 | Track contact request status | 🟡 Important | ⭐ Very Low |
| F9 | Share player profiles with colleagues | 🟡 Important | ⭐⭐ Low |
| F10 | Export reports for recruitment meetings | 🟢 Nice-to-have | ⭐ Very Low |

### Emotional Jobs (How they want to FEEL)

| Job # | Emotional Need | Current State |
|-------|---------------|---------------|
| E1 | **Confident** that player info is accurate | 😰 Skeptical - too many fakes |
| E2 | **Ahead** of competition (found player first) | 😐 Neutral - same sources as everyone |
| E3 | **Efficient** - not wasting time on logistics | 😤 Frustrated - too much admin |
| E4 | **Excited** when discovering exceptional talent | 😊 This is why they do the job |
| E5 | **Trusted** by their club for recommendations | 😰 Nervous - one bad rec hurts reputation |
| E6 | **Connected** to a reliable network | 😐 Fragmented - WhatsApp chaos |

### Social Jobs (How they want to be PERCEIVED)

| Job # | Social Need |
|-------|-------------|
| S1 | Seen as having a "good eye" for talent |
| S2 | Known for finding players early |
| S3 | Respected by academies and agents |
| S4 | Valued by their club for unique access |

---

## Part 3: The Scout's Real Workflow (Current State)

### Step-by-Step: How Scouts Find Players Today

```
┌─────────────────────────────────────────────────────────────────┐
│                    CURRENT SCOUT WORKFLOW                        │
│                    (Without Unicorn Sport)                       │
└─────────────────────────────────────────────────────────────────┘

1. HEAR ABOUT A PLAYER
   └─> WhatsApp from agent
   └─> Tip from local contact
   └─> Tournament attendance
   └─> YouTube algorithm
   
2. REQUEST VIDEO
   └─> Ask agent/academy for footage
   └─> Wait 1-7 days
   └─> Receive cherry-picked highlights (biased)
   
3. EVALUATE VIDEO
   └─> Watch on phone/laptop
   └─> No context (opposition level? age group?)
   └─> Can't verify if player is really that age
   
4. VERIFY PLAYER
   └─> Ask for documents (birth cert, passport)
   └─> Documents often unreliable
   └─> MRI bone scan for serious interest ($$$)
   
5. COMPARE OPTIONS
   └─> Manual Excel spreadsheet
   └─> Memory-based ("I think Player A was faster")
   └─> No side-by-side video comparison
   
6. SHORTLIST & REPORT
   └─> Write report in Word
   └─> Share via email
   └─> Present to recruitment committee
   
7. CONTACT ACADEMY
   └─> Find contact (often difficult)
   └─> Cold email/WhatsApp
   └─> No response tracking
   └─> Multiple intermediaries
   
8. ARRANGE TRIAL
   └─> Negotiate directly
   └─> Handle visa, travel, accommodation
   └─> Hope player shows up
   
                         ⏱️ TOTAL TIME: 2-4 weeks per player
                         💰 TOTAL COST: $500-5000 (travel, admin)
                         😤 FRUSTRATION: HIGH
```

### What Scouts ACTUALLY Spend Time On

| Activity | Time Spent | Value Added |
|----------|------------|-------------|
| Watching videos | 20% | ✅ High |
| Searching for videos | 25% | ❌ Low |
| Verifying player info | 15% | ⚠️ Medium (but necessary) |
| Communication/logistics | 30% | ❌ Low |
| Writing reports | 10% | ✅ High |

**Insight:** 55% of scout time is low-value administrative work. Our platform should eliminate this.

---

## Part 4: The Ideal Workflow (With Unicorn Sport)

```
┌─────────────────────────────────────────────────────────────────┐
│                     IDEAL SCOUT WORKFLOW                         │
│                     (With Unicorn Sport)                         │
└─────────────────────────────────────────────────────────────────┘

1. DISCOVERY (Passive + Active)
   ├─> [ACTIVE] Open Discover → Filter by criteria → Browse
   ├─> [PASSIVE] Video Feed → Swipe → See "wow" moments
   └─> [PASSIVE] Email digest: "New players matching your criteria"
   
   ⏱️ Time: 5-15 minutes
   
2. QUICK EVALUATE  
   ├─> See player card: Age ✓ Position ✓ Academy ✓ Verified ✓
   ├─> Watch 30-60 second highlight
   └─> Instant gut check: "Interested or skip"
   
   ⏱️ Time: 30 seconds per player
   
3. DEEP EVALUATE (for interested players)
   ├─> Full player profile
   ├─> Multiple videos (highlights + full matches)
   ├─> Physical stats, playing history
   ├─> Academy reputation
   └─> Verification status (age confirmed)
   
   ⏱️ Time: 5-10 minutes per player
   
4. COMPARE & DECIDE
   ├─> Select 2-4 players
   ├─> Side-by-side comparison view
   ├─> Stats table + video snippets
   └─> Make shortlist decision
   
   ⏱️ Time: 10 minutes
   
5. SAVE & ANNOTATE
   ├─> Save to shortlist with one click
   ├─> Add personal notes
   ├─> Categorize (Watch list / Contact / Priority)
   └─> Set reminder
   
   ⏱️ Time: 1 minute
   
6. SHARE & REPORT
   ├─> Export PDF player profile
   ├─> Share link with colleague
   ├─> Download comparison report
   └─> Present at recruitment meeting
   
   ⏱️ Time: 2 minutes
   
7. CONTACT
   ├─> See academy info + contact history
   ├─> Send structured inquiry
   ├─> Track status in dashboard
   └─> Receive response (expected: 48-72 hrs)
   
   ⏱️ Time: 2 minutes
   
                         ⏱️ TOTAL TIME: 30-60 minutes per player
                         💰 TOTAL COST: $99-399/month subscription
                         😊 FRUSTRATION: LOW
```

---

## Part 5: Critical Success Factors

### What MUST Work for Scouts to Pay

| Factor | Why It's Critical | How We Deliver |
|--------|-------------------|----------------|
| **Trust** | "Is this player really 17?" | Verification badges, academy partnerships, document verification |
| **Quality** | "Are these actually good players?" | Curated academies, tournament footage, no agent-filtered content |
| **Speed** | "I need to find players before competitors" | Real-time uploads, notifications, efficient filters |
| **Completeness** | "I need to see full matches, not just highlights" | Full match archive for paid tiers |
| **Connection** | "I need to actually reach these players" | Direct academy contact, response tracking |
| **Exclusivity** | "Am I seeing players others don't?" | Focus on under-represented African regions |

### Dealbreakers (If We Fail Here, They Leave)

| ❌ Dealbreaker | Impact |
|---------------|--------|
| Fake/wrong player ages | Destroys all trust, never returns |
| Can't contact academy | Platform is useless for actual deals |
| Only cherry-picked highlights | Same as YouTube, no value add |
| Cluttered/slow interface | Scouts are busy, will abandon |
| No new players (stale content) | Why pay monthly? |

---

## Part 6: Algorithm & Feed Design

### Discovery Algorithm: What to Show

#### For TikTok-Style Video Feed

The feed should optimize for: **Engagement + Discovery + Diversity**

```python
# Conceptual feed algorithm

def generate_feed(scout):
    feed = []
    
    # 40% - Personalized based on history
    # "More like what you've saved/watched"
    feed += get_similar_to_saved(scout, limit=4)
    
    # 30% - New & Trending
    # "Hot this week" - recently uploaded, high engagement
    feed += get_new_trending(limit=3)
    
    # 20% - Discovery (expand horizons)  
    # Players outside their usual filters
    feed += get_discovery_picks(scout, limit=2)
    
    # 10% - Random quality
    # Prevent filter bubble
    feed += get_random_verified(limit=1)
    
    return shuffle_with_smart_ordering(feed)
```

#### Smart Ordering Rules

1. **Never show same academy back-to-back** (variety)
2. **Never show same position 3x in a row** (variety)
3. **Prioritize verified players** (trust)
4. **Boost recently uploaded** (freshness)
5. **Boost players with full match available** (depth)

### What to Show on Video Overlay

**Essential (Always Visible):**
```
┌────────────────────────────────┐
│ [Verified ✓]                   │
│                                │
│ NAME: John Mensah              │
│ AGE: 17 • POSITION: Striker    │
│ ACADEMY: Rising Stars FC       │
│ COUNTRY: 🇳🇬 Nigeria            │
│                                │
│ [❤️ Save] [👤 Profile] [📤 Share]│
└────────────────────────────────┘
```

**Secondary (Available on tap):**
- Height/Weight
- Preferred foot
- Tournament (video context)
- Video date

### Notification & Alert Strategy

| Trigger | Notification | Frequency |
|---------|-------------|-----------|
| New player matches saved search | Push + Email | Real-time |
| Contact request response | Push + Email | Real-time |
| Weekly digest of new players | Email | Weekly (Sunday) |
| Player you saved has new video | Push | Real-time |
| Academy you follow adds player | Push | Real-time |

---

## Part 7: Competitive Advantage

### Why Unicorn Sport vs. Alternatives

| Alternative | Their Strength | Our Advantage |
|-------------|---------------|---------------|
| **Wyscout** | Comprehensive European data | Africa-focused, verified ages |
| **YouTube** | Free, massive content | Curated, verified, contactable |
| **Agents** | Personal relationships | Unbiased, complete access |
| **Travel to tournaments** | See players live | Pre-filter before expensive travel |

### Our Unique Value Proposition

> **"The only platform where you can discover, verify, and contact African football talent in one place."**

Core differentiators:
1. **Verified ages** - MRI/document verification partnerships
2. **Academy partnerships** - Direct, not through agents
3. **Complete workflow** - Discover → Evaluate → Contact
4. **Africa specialist** - Depth no one else has

---

## Part 8: Feature Priority Matrix

### Must-Have (MVP - Without These, Don't Launch)

| Feature | JTBD Addressed | Implementation Complexity |
|---------|---------------|--------------------------|
| Video player with highlights | F1, F3, E4 | Medium |
| Player profiles with stats | F2, F3 | Low |
| Verification badges | F4, E1, E5 | Medium |
| Save to shortlist | F6 | Low |
| Contact academy | F7, E6 | Medium |
| Position/Age/Country filters | F2 | Low |
| Search by name | F2 | Low |

### Should-Have (Launch +30 Days)

| Feature | JTBD Addressed | Implementation Complexity |
|---------|---------------|--------------------------|
| Notes on saved players | F6, E5 | Low |
| Full match videos | F3 | Medium |
| Contact status tracking | F8 | Medium |
| Player comparison | F5 | Medium |
| Academy profiles | F7 | Medium |
| Email notifications | F1 | Medium |

### Nice-to-Have (Launch +90 Days)

| Feature | JTBD Addressed | Implementation Complexity |
|---------|---------------|--------------------------|
| PDF export | F9, F10 | Medium |
| Team collaboration | F9, S4 | High |
| Similar players suggestions | F1 | High (ML) |
| Advanced analytics | E5 | High |
| API access | F9 | High |

---

## Part 9: Metrics to Track

### North Star Metric
> **Contact Requests per Active Scout per Month**

This captures: Discovery → Interest → Action

### Supporting Metrics

| Metric | What It Tells Us | Target |
|--------|-----------------|--------|
| Time to first save | Is discovery working? | < 5 min |
| Videos watched per session | Is content engaging? | > 10 |
| Save rate | Are players quality? | > 5% of views |
| Contact request rate | Is workflow complete? | > 20% of saves |
| Contact response rate | Are academies engaged? | > 70% in 72hrs |
| Weekly return rate | Is there ongoing value? | > 60% |
| Free → Paid conversion | Is value clear? | > 10% |
| Churn rate | Are we delivering value? | < 5% monthly |

---

## Part 10: Open Questions to Resolve

### Business Model Questions
1. Do we charge academies to be listed? (Marketplace model)
2. Do we take commission on successful placements?
3. How do we verify academies are legitimate?

### Product Questions
4. Do we show unverified players with a warning, or hide them completely?
5. Do free users see the same players as paid users?
6. How do we handle player privacy (full name vs. abbreviated)?
7. Do we allow agents on the platform, or academies only?

### Technical Questions
8. How do we verify player ages? (Partner with verification service?)
9. How do we ensure video quality standards?
10. How do we prevent video piracy/downloads?

---

## Summary: The Golden Insights

### 1. Trust is Everything
Scouts have been burned by fake ages, doctored videos, and unreliable agents. Every feature must reinforce trust.

### 2. Save Their Time
Scouts spend 55% of time on admin. We win by eliminating friction between "interesting player" and "contacted academy."

### 3. Video is the Hook, Data is the Sale
TikTok feed creates excitement. Rich profiles and verification close the deal.

### 4. Connection is the Ultimate Value
Seeing players is table stakes. Connecting with them is why they pay.

### 5. Fresh Content is Non-Negotiable
Stale player database = cancelled subscription. We need constant new content.

---

*This document should be reviewed before any major UX decision.*  
*Last updated: January 29, 2026*
