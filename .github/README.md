# Unicorn Sport - Project Documentation

## 🎯 Mission Statement

**Unicorn Sport** is a professional talent discovery and showcase platform operated by a scouting academy. We record, produce, and distribute verified African youth football talent content to professional scouts and clubs worldwide.

### The Problem We Solve

1. **Age Fraud Epidemic** - Age falsification is rampant in African football, with players routinely misrepresenting their ages to appear younger or meet eligibility requirements
2. **Talent Discovery Gap** - Exceptional athletes in rural and underserved areas lack visibility to international scouts
3. **Trust Deficit** - Scouts and clubs hesitate to invest in African talent due to verification challenges
4. **Quality Content Gap** - Professional scouts need high-quality video content, not amateur uploads

### Our Solution

A professional academy-operated platform that:
- **Runs talent showcase events** where students come to play organized matches
- **Records professional content** using our media team's equipment
- **Creates player highlights** extracted from full match recordings
- **Verifies student identity** through document verification at registration
- **Provides subscription access** for scouts to view highlights and contact players

### How We're Different

We are NOT a user-generated content platform. We are:
- **The content creators** - Our media team records and produces all videos
- **The verification authority** - We verify player data at the source
- **The gatekeepers** - Scouts must subscribe to access highlights

---

## 📚 Documentation Index

### 🚨 P0 - Critical (Must Build First)

| Document | Description |
|----------|-------------|
| [ADMIN_PORTAL.md](./ADMIN_PORTAL.md) | **🚨 P0** - Create player profiles, upload videos, manage content |
| [SUBSCRIPTION_SERVICE.md](./SUBSCRIPTION_SERVICE.md) | **🚨 P0** - Stripe integration, highlight paywall |

### 🔹 P1 - Core Features

| Document | Description |
|----------|-------------|
| [VERIFICATION_SYSTEM.md](./VERIFICATION_SYSTEM.md) | **🔹 P1** - NIN/Passport age verification process |

### 🔸 P2 - Partnerships & Enhancements

| Document | Description |
|----------|-------------|
| [SCHOOL_PORTAL.md](./SCHOOL_PORTAL.md) | **🔸 P2** - School referral partnerships |

### 📖 Platform Documentation

| Document | Description |
|----------|-------------|
| [BUSINESS_MODEL.md](./BUSINESS_MODEL.md) | Academy operation and revenue model |
| [CONTENT_MANAGEMENT.md](./CONTENT_MANAGEMENT.md) | Video storage, delivery, and access control |
| [VISION.md](./VISION.md) | Project vision, goals, and success metrics |
| [ARCHITECTURE.md](./ARCHITECTURE.md) | Technical architecture and system design |
| [DATA_MODEL.md](./DATA_MODEL.md) | Database schema and data relationships |
| [USER_JOURNEYS.md](./USER_JOURNEYS.md) | User flows for players, scouts, and admins |
| [API_REFERENCE.md](./API_REFERENCE.md) | API endpoints and integration guide |
| [DEPLOYMENT.md](./DEPLOYMENT.md) | Infrastructure and deployment guide |
| [SECURITY.md](./SECURITY.md) | Security measures and compliance |
| [FRONTEND_SPEC.md](./FRONTEND_SPEC.md) | Complete frontend blueprint |
| **[DEVELOPMENT_SEQUENCE.md](./DEVELOPMENT_SEQUENCE.md)** | **🚀 Build order and implementation guide** |
| **[ENVIRONMENT_SETUP.md](./ENVIRONMENT_SETUP.md)** | **🌍 Local/Staging/Production configuration** |

---

## 🏛️ Platform Overview

### How It Works

```
┌─────────────────────────────────────────────────────────────────────┐
│                    UNICORN SPORT COMPLETE WORKFLOW                   │
├─────────────────────────────────────────────────────────────────────┤
│                                                                      │
│  1. VERIFY FIRST        2. TOURNAMENT            3. RECORD          │
│  ┌──────────────┐       ┌──────────────┐       ┌──────────────┐     │
│  │ NIN/Passport │       │   Players    │       │  Media team  │     │
│  │ checked by   │──────▶│   compete    │──────▶│   records    │     │
│  │ academy      │       │ (age-verified│       │   matches    │     │
│  │              │       │   only)      │       │              │     │
│  └──────────────┘       └──────────────┘       └──────────────┘     │
│                                                        │             │
│                                                        ▼             │
│  4. CREATE PROFILES     5. DISTRIBUTE           6. MONETIZE         │
│  ┌──────────────┐       ┌──────────────┐       ┌──────────────┐     │
│  │  Academy     │       │ Highlights   │       │   Scouts     │     │
│  │  creates     │──────▶│ → FREE       │──────▶│   pay for    │     │
│  │  player      │       │ Full Matches │       │   full match │     │
│  │  profiles    │       │ → PAID       │       │   context    │     │
│  └──────────────┘       └──────────────┘       └──────────────┘     │
│                                                                      │
└─────────────────────────────────────────────────────────────────────┘
```

### Core Features

| Priority | Feature | Description | Status |
|----------|---------|-------------|--------|
| - | 🎬 Match Recording | Full match video storage | ✅ Built |
| - | ✂️ Player Highlights | Individual player clips | ✅ Built |
| - | 👤 Profile System | Academy-created player profiles | ✅ Built |
| - | 🔐 Authentication | Secure user authentication | ✅ Built |
| - | 🔍 Discovery | Search and find talent | ✅ Built |
| **P0** | 👨‍💼 **Admin Portal** | Create profiles, upload content | 🚨 **CRITICAL** |
| **P0** | 💳 **Subscriptions** | Scout paywall for full matches | 🚨 **CRITICAL** |
| **P1** | 🎬 Video Player | Secure streaming | 📋 Planned |
| **P2** | 📧 Contact System | Scout-player messaging | 📋 Planned |
| **P2** | 📊 Analytics | View tracking & reporting | 📋 Planned |

> **🚨 P0 = Must build FIRST. Without Admin Portal, can't create profiles or upload. Without Subscriptions, no revenue.**

### Content Distribution Strategy

```
┌──────────────────────────────────────────────────────────────────────────┐
│                         CONTENT MODEL                                     │
│                                                                           │
│   YOUTUBE (Marketing Only)                                                │
│   ┌─────────────────────────────────────────────────────────────────┐    │
│   │ • Sample matches (1 of many) - just a taste                     │    │
│   │ • Attracts scouts to the platform                               │    │
│   │ • Link: "Want more? Visit unicornsport.africa"                  │    │
│   └─────────────────────────────────────────────────────────────────┘    │
│                                                                           │
│   PLATFORM - FREE                    PLATFORM - PAID                      │
│   ┌─────────────────────┐            ┌─────────────────────┐             │
│   │ • Watch HIGHLIGHTS  │            │ • Browse profiles   │             │
│   │ • See thumbnails    │            │ • Watch FULL MATCHES│             │
│   │                     │            │ • Access analytics  │             │
│   │ (Flashy but can     │            │ • Contact players   │             │
│   │  be deceptive)      │            │ (See the REAL       │             │
│   │                     │            │  performance)       │             │
│   └─────────────────────┘            └─────────────────────┘             │
│                                                                           │
│   💡 WHY PAY? Highlights can be deceptive. Full matches show truth.      │
│                                                                           │
└──────────────────────────────────────────────────────────────────────────┘
```
│   │         │                   │         │                │         │  │
│   └────┬────┘                   └─────────┘                └─────────┘  │
│        │                             │                          ▲        │
│        ▼                             │                          │        │
│   ┌─────────┐                        │                     ┌─────────┐  │
│   │ Creates │                        │                     │Subscribe│  │
│   │ player  │────────────────────────┘                     │ to view │  │
│   │highlights│    (Direct upload,                          │ videos  │  │
│   └─────────┘     no transcoding)                          └─────────┘  │
│                                                                           │
└──────────────────────────────────────────────────────────────────────────┘
```

---

## 🚀 Quick Links

- **Backend Services**: [/backend/README.md](../backend/README.md)
- **Frontend App**: [/frontend/README.md](../frontend/README.md)
- **API Testing**: [Postman Collection](./postman/)

---

## 👥 Team & Contacts

*Add team information here*

---

## 📄 License

Proprietary - All Rights Reserved
