# Admin Data Quality Guide

## 🚨 Current Issues Blocking Scout Experience

### 1. Player Photos Are Wrong
**Problem:** Profile photos contain logos and promotional images instead of player photos.

| Player | Current Photo | Should Be |
|--------|---------------|-----------|
| Ayomide C. | Fiilar logo (`fi-icon2.png`) | Professional headshot or action shot |
| Jeffrey O. | "Core Team Applications" poster | Professional headshot or action shot |
| Abdulrasaq N. | No photo | Professional headshot or action shot |

**How to Fix:**
1. Go to Admin → Players → [Player Name]
2. Upload a proper photo:
   - **Best:** Action shot from a match/training (scouts love seeing players in action)
   - **Good:** Professional headshot with neutral background
   - **Acceptable:** Clear candid photo showing the player's face
3. Delete the incorrect logo/poster files

### 2. No Videos Linked to Players
**Problem:** `video_count: 0` for all players. Videos exist in the system but aren't linked to player profiles.

**How to Fix:**
1. Go to Admin → Videos
2. For each video, click "Edit" and:
   - Select the player(s) featured in that video
   - Ensure video has a thumbnail (auto-generated or uploaded)
3. Alternatively, go to Admin → Players → [Player Name] → Videos tab
4. Click "Link Video" and select relevant highlights

**Why This Matters:**
- When videos are linked, scouts see "▶ 2 videos" on player cards
- Video thumbnails become the primary card image (action shots!)
- Scouts can quickly identify who has content worth watching

### 3. Incorrect Player Data
**Problem:** Some player data appears incorrect or suspicious.

| Player | Issue | Likely Correct Value |
|--------|-------|---------------------|
| Ayomide C. | Height: 120cm (4 feet) | ~160-170cm for 14yo goalkeeper |
| Abdulrasaq N. | Age: missing | Check date of birth |

**How to Fix:**
1. Go to Admin → Players → [Player Name]
2. Verify and correct:
   - Date of birth (auto-calculates age)
   - Height in centimeters
   - Weight in kilograms
   - Preferred foot

---

## ✅ Data Quality Checklist (Per Player)

Before marking a player as "verified", ensure:

- [ ] **Profile Photo** - Clear photo of the player (not a logo)
- [ ] **Full Name** - First and last name correct
- [ ] **Date of Birth** - Valid date, age calculates correctly
- [ ] **Position** - Accurate primary position
- [ ] **Height** - Realistic for age (14yo: 150-180cm typical)
- [ ] **Country/State** - Correct location
- [ ] **Academy** - Linked if applicable
- [ ] **At Least 1 Video** - Linked highlight showing player in action

---

## 📊 Ideal Player Card Data

For the best scout experience, each player should have:

```
┌─────────────────────────────────────┐
│ [ACTION SHOT from video highlight] │  ← video_thumbnail_url
│ ▶ 2 videos                     GK  │  ← video_count > 0
│ ────────────────────────────────── │
│ Ayomide C.                    ✓    │  ← verified
│ 🇳🇬 Nigeria                         │
│ 🏫 Lagos Football Academy          │  ← academy linked
│ 14y • 170cm • Both                 │  ← accurate stats
└─────────────────────────────────────┘
```

---

## 🎯 Priority Actions

1. **Immediate:** Fix the 3 test players' photos
2. **Immediate:** Link at least 1 video to each player
3. **This Week:** Verify all player data accuracy
4. **Ongoing:** Ensure new players meet quality checklist before verification

---

## Video Thumbnail Best Practices

When uploading videos, the system can auto-generate thumbnails, but you can also upload custom ones:

**Good thumbnails show:**
- Player actively playing (kicking, saving, dribbling)
- Clear view of the player
- Good lighting, not blurry

**Bad thumbnails:**
- Crowd shots
- Blurry or dark images
- Multiple players where subject is unclear
- Logos or text overlays
