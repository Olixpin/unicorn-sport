# Unicorn Sport - API Reference

## 📡 Overview

Unicorn Sport exposes a RESTful API from a **single Go binary**. All endpoints use JSON for request/response bodies and JWT for authentication.

> **MVP Architecture:** Single API server with modular internal structure. 
> All endpoints served from one base URL.

---

## 🔐 Authentication

### Base URL
```
Production: https://api.unicornsport.africa/v1
Development: http://localhost:8080/api/v1
```

### Authentication Header

```http
Authorization: Bearer <access_token>
```

### Token Lifecycle
- **Access Token**: Valid for 15 minutes
- **Refresh Token**: Valid for 7 days
- Use refresh endpoint to get new access token

---

## 📋 Auth Endpoints

### Register Scout (Self-Registration)

> **Note:** Only scouts can self-register. Players are created by admin.

```http
POST /api/v1/auth/register
```

**Request Body:**
```json
{
  "email": "scout@agency.com",
  "password": "securePassword123!",
  "organization_name": "Premier Talent Agency",
  "organization_type": "agency",
  "country": "United Kingdom"
}
```

**Response (201 Created):**
```json
{
  "success": true,
  "data": {
    "user": {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "email": "scout@agency.com",
      "role": "scout",
      "email_verified": false,
      "created_at": "2026-01-25T10:30:00Z"
    },
    "subscription": {
      "tier": "free",
      "status": "active"
    },
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "refresh_token": "dGhpcyBpcyBhIHJlZnJlc2g...",
    "expires_in": 900
  }
}
```

**Validation:**
- `email`: Required, valid email format, unique
- `password`: Required, min 8 chars, 1 uppercase, 1 number
- `organization_name`: Required for scouts

---

### Login

```http
POST /api/v1/auth/login
```

**Request Body:**
```json
{
  "email": "scout@agency.com",
  "password": "securePassword123!"
}
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "user": {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "email": "scout@agency.com",
      "role": "scout"
    },
    "subscription": {
      "tier": "scout",
      "status": "active",
      "expires_at": "2026-02-25T10:30:00Z"
    },
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "refresh_token": "dGhpcyBpcyBhIHJlZnJlc2g...",
    "expires_in": 900
  }
}
```

**Error Response (401 Unauthorized):**
```json
{
  "success": false,
  "error": {
    "code": "INVALID_CREDENTIALS",
    "message": "Invalid email or password"
  }
}
```

---

### Refresh Token

```http
POST /api/v1/auth/refresh
```

**Request Body:**
```json
{
  "refresh_token": "dGhpcyBpcyBhIHJlZnJlc2g..."
}
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "refresh_token": "bmV3IHJlZnJlc2ggdG9rZW4...",
    "expires_in": 900
  }
}
```

---

### Logout

```http
POST /api/v1/auth/logout
Authorization: Bearer <access_token>
```

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Successfully logged out"
}
```

---

### Get Current User

```http
GET /api/v1/auth/me
Authorization: Bearer <access_token>
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "email": "player@example.com",
    "role": "player",
    "email_verified": true,
    "phone": "+234801234567",
    "phone_verified": false,
    "status": "active",
    "created_at": "2026-01-25T10:30:00Z",
    "last_login_at": "2026-01-25T14:00:00Z"
  }
}
```

---

### Verify Email

```http
POST /api/v1/auth/verify-email
```

**Request Body:**
```json
{
  "code": "123456"
}
```

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Email verified successfully"
}
```

---

### Forgot Password

```http
POST /api/v1/auth/forgot-password
```

**Request Body:**
```json
{
  "email": "player@example.com"
}
```

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Password reset instructions sent to email"
}
```

---

### Reset Password

```http
POST /api/v1/auth/reset-password
```

**Request Body:**
```json
{
  "code": "123456",
  "new_password": "newSecurePassword123!"
}
```

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Password reset successfully"
}
```

---

## 👤 Player Endpoints

> **Access Control:** Players are created by admin only. Players can view their own profile.

### List Players (Public)

```http
GET /api/v1/players
```

**Query Parameters:**
| Param | Type | Description |
|-------|------|-------------|
| position | string | Filter by position (e.g., "midfielder") |
| country | string | Filter by country |
| age_min | int | Minimum age |
| age_max | int | Maximum age |
| page | int | Page number (default: 1) |
| limit | int | Results per page (default: 20, max: 100) |

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "players": [
      {
        "id": "660e8400-e29b-41d4-a716-446655440001",
        "first_name": "James",
        "last_name_init": "O.",
        "age": 17,
        "position": "Midfielder",
        "country": "Nigeria",
        "thumbnail_url": "https://storage.unicornsport.africa/profiles/...",
        "is_verified": true
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 20,
      "total": 156,
      "total_pages": 8
    }
  }
}
```

---

### Get Player Detail

```http
GET /api/v1/players/{id}
Authorization: Bearer <access_token>
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "id": "660e8400-e29b-41d4-a716-446655440001",
    "first_name": "James",
    "last_name_init": "O.",
    "age": 17,
    "position": "Midfielder",
    "preferred_foot": "right",
    "height_cm": 175,
    "weight_kg": 68,
    "country": "Nigeria",
    "state": "Lagos",
    "school_name": "Lagos High School",
    "tournament_name": "Unicorn Cup 2025",
    "profile_photo_url": "https://...",
    "is_verified": true,
    "highlight_videos": [
      {
        "id": "vid-001",
        "title": "James O. Highlights - Unicorn Cup 2025",
        "thumbnail_url": "https://...",
        "duration_seconds": 180
      }
    ],
    "full_match_videos": []
  }
}
```

> **Note:** `full_match_videos` only populated for Scout tier and above.
> Free tier sees empty array with upgrade prompt.

---

### Create Player (Admin Only)

```http
POST /api/v1/admin/players
Authorization: Bearer <admin_token>
```

**Request Body:**
```json
{
  "first_name": "James",
  "last_name": "Okoro",
  "date_of_birth": "2009-03-15",
  "position": "Midfielder",
  "preferred_foot": "right",
  "height_cm": 175,
  "weight_kg": 68,
  "country": "Nigeria",
  "state": "Lagos",
  "city": "Ikeja",
  "school_name": "Lagos High School",
  "tournament_id": "tourn-001",
  "verification_status": "verified",
  "verification_doc_type": "nin"
}
```

**Response (201 Created):**
```json
{
  "success": true,
  "data": {
    "id": "660e8400-e29b-41d4-a716-446655440001",
    "credentials": {
      "email": "james.okoro.2009@unicornsport.africa",
      "temp_password": "TempPass123!"
    }
  }
}
```

---

### Update Player (Admin Only)

```http
PUT /api/v1/admin/players/:id
Authorization: Bearer <admin_token>
```

**Request Body:**
```json
{
  "first_name": "James",
  "last_name": "Okoro",
  "position": "Central Midfielder",
  "height_cm": 177,
  "weight_kg": 70,
  "verification_status": "verified"
}
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "id": "660e8400-e29b-41d4-a716-446655440001",
    "first_name": "James",
    "last_name": "Okoro",
    "updated_at": "2026-01-26T14:00:00Z"
  }
}
```

---

### Delete Player (Admin Only)

```http
DELETE /api/v1/admin/players/:id
Authorization: Bearer <admin_token>
```

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Player deleted successfully"
}
```

> **Note:** Soft delete - sets `deleted_at` timestamp. Associated videos remain but are unlinked.

---

## 🏆 Tournament Endpoints (Admin)

### List Tournaments

```http
GET /api/v1/admin/events
Authorization: Bearer <admin_token>
```

**Query Parameters:**
| Param | Type | Description |
|-------|------|-------------|
| status | string | "upcoming", "ongoing", "completed" |
| year | int | Filter by year |
| page | int | Page number |
| limit | int | Results per page |

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "events": [
      {
        "id": "tourn-001",
        "name": "Unicorn Cup 2026",
        "location": "Lagos, Nigeria",
        "start_date": "2026-06-01",
        "end_date": "2026-06-15",
        "status": "upcoming",
        "player_count": 0
      }
    ],
    "pagination": { "page": 1, "limit": 20, "total": 5 }
  }
}
```

---

### Create Tournament (Admin Only)

```http
POST /api/v1/admin/events
Authorization: Bearer <admin_token>
```

**Request Body:**
```json
{
  "name": "Unicorn Cup 2026",
  "location": "Lagos, Nigeria",
  "country": "Nigeria",
  "start_date": "2026-06-01",
  "end_date": "2026-06-15",
  "description": "Annual youth football tournament"
}
```

**Response (201 Created):**
```json
{
  "success": true,
  "data": {
    "id": "tourn-002",
    "name": "Unicorn Cup 2026",
    "created_at": "2026-01-25T10:00:00Z"
  }
}
```

---

### Update Tournament (Admin Only)

```http
PUT /api/v1/admin/events/:id
Authorization: Bearer <admin_token>
```

**Request Body:**
```json
{
  "name": "Unicorn Cup 2026 - Lagos",
  "end_date": "2026-06-20"
}
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "id": "tourn-002",
    "name": "Unicorn Cup 2026 - Lagos",
    "updated_at": "2026-01-26T12:00:00Z"
  }
}
```

---

## 👥 User Management (Admin)

### List Users (Admin Only)

```http
GET /api/v1/admin/users
Authorization: Bearer <admin_token>
```

**Query Parameters:**
| Param | Type | Description |
|-------|------|-------------|
| role | string | "scout", "admin", "player" |
| subscription_tier | string | "free", "scout", "pro", "club" |
| page | int | Page number |
| limit | int | Results per page |

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "users": [
      {
        "id": "user-001",
        "email": "scout@agency.com",
        "role": "scout",
        "organization_name": "Premier Talent Agency",
        "subscription_tier": "pro",
        "created_at": "2026-01-01T00:00:00Z",
        "last_login": "2026-01-25T09:00:00Z"
      }
    ],
    "pagination": { "page": 1, "limit": 20, "total": 150 }
  }
}
```

---

### Update User Status (Admin Only)

```http
PUT /api/v1/admin/users/:id
Authorization: Bearer <admin_token>
```

**Request Body:**
```json
{
  "is_active": false,
  "role": "scout"
}
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "id": "user-001",
    "is_active": false,
    "updated_at": "2026-01-25T12:00:00Z"
  }
}
```

---

## 📹 Video Endpoints

### List Highlights (FREE - Everyone)

```http
GET /api/v1/videos/highlights
```

**Query Parameters:**
| Param | Type | Description |
|-------|------|-------------|
| player_id | uuid | Filter by player |
| tournament_id | uuid | Filter by tournament |
| page | int | Page number |
| limit | int | Results per page |

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "videos": [
      {
        "id": "vid-001",
        "video_type": "highlight",
        "title": "James O. Highlights - Unicorn Cup 2025",
        "thumbnail_url": "https://...",
        "duration_seconds": 180,
        "blob_url": "https://storage.unicornsport.africa/highlights/..."
      }
    ],
    "pagination": { "page": 1, "limit": 20, "total": 45 }
  }
}
```

---

### List Full Matches (PAID - Scout+ Only)

```http
GET /api/v1/videos/full-matches
Authorization: Bearer <access_token>
```

**Response (200 OK) - For Scout+ tier:**
```json
{
  "success": true,
  "data": {
    "videos": [
      {
        "id": "match-001",
        "video_type": "full_match",
        "title": "Unicorn Cup 2025 - Semi Final",
        "thumbnail_url": "https://...",
        "duration_seconds": 5400,
        "blob_url": "https://storage.unicornsport.africa/matches/..."
      }
    ]
  }
}
```

**Response (402 Payment Required) - For Free tier:**
```json
{
  "success": false,
  "error": {
    "code": "SUBSCRIPTION_REQUIRED",
    "message": "Upgrade to Scout tier to access full matches",
    "upgrade_url": "/subscribe"
  }
}
```

---

### Get Video Upload URL (Admin Only)

> **Recommended for large files.** Returns a presigned S3 URL for direct browser upload.

```http
POST /api/v1/admin/videos/upload-url
Authorization: Bearer <admin_token>
```

**Request Body:**
```json
{
  "filename": "match_2026_final.mp4",
  "content_type": "video/mp4",
  "file_size": 524288000
}
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "upload_url": "https://unicornsport-videos.s3.amazonaws.com/uploads/abc123?X-Amz-Signature=...",
    "upload_id": "upload-550e8400-e29b-41d4-a716-446655440000",
    "expires_in": 3600,
    "max_file_size": 536870912
  }
}
```

**Validation:**
- `file_size`: Max 512MB per video
- `content_type`: Must be "video/mp4", "video/webm", or "video/quicktime"

---

### Create Video Record (Admin Only)

> **After presigned upload completes**, create the video record.

```http
POST /api/v1/admin/videos
Authorization: Bearer <admin_token>
```

**Request Body:**
```json
{
  "upload_id": "upload-550e8400-e29b-41d4-a716-446655440000",
  "title": "U17 Finals - Full Match",
  "video_type": "full_match",
  "player_ids": ["player-uuid-1", "player-uuid-2"],
  "tournament_id": "tournament-uuid"
}
```

**Response (201 Created):**
```json
{
  "success": true,
  "data": {
    "id": "vid-002",
    "s3_key": "videos/2026/01/vid-002.mp4",
    "cdn_url": "https://cdn.unicornsport.africa/videos/2026/01/vid-002.mp4",
    "status": "processing",
    "created_at": "2026-01-25T10:30:00Z"
  }
}
```

**Video Statuses:**
- `processing`: Video uploaded, transcoding in progress
- `ready`: Transcoding complete, video playable
- `failed`: Processing failed, retry or re-upload

---

## 💳 Subscription Endpoints

### Get Current Subscription

```http
GET /api/v1/subscriptions/me
Authorization: Bearer <access_token>
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "tier": "scout",
    "status": "active",
    "current_period_start": "2026-01-01T00:00:00Z",
    "current_period_end": "2026-02-01T00:00:00Z"
  }
}
```

---

### Create Checkout Session (Upgrade)

```http
POST /api/v1/subscriptions/checkout
Authorization: Bearer <access_token>
```

**Request Body:**
```json
{
  "tier": "pro",
  "success_url": "https://unicornsport.africa/subscribe/success",
  "cancel_url": "https://unicornsport.africa/subscribe/cancel"
}
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "checkout_url": "https://checkout.stripe.com/..."
  }
}
```

---

### Stripe Webhook

```http
POST /api/v1/webhooks/stripe
Stripe-Signature: <signature>
```

Handles: `checkout.session.completed`, `invoice.paid`, `customer.subscription.updated`, `customer.subscription.deleted`

---

## 📧 Contact Endpoints

### Create Contact Request (Pro+ Only)

```http
POST /api/v1/contact-requests
Authorization: Bearer <access_token>
```

**Request Body:**
```json
{
  "player_id": "660e8400-e29b-41d4-a716-446655440001",
  "message": "Interested in discussing opportunities for this player."
}
```

**Response (201 Created):**
```json
{
  "success": true,
  "data": {
    "id": "req-001",
    "status": "pending",
    "message": "Your request has been submitted. Our team will forward it to the player."
  }
}
```

**Response (402 Payment Required) - If not Pro+:**
```json
{
  "success": false,
  "error": {
    "code": "SUBSCRIPTION_REQUIRED",
    "message": "Upgrade to Pro tier to contact players"
  }
}
```

---

## 💾 Saved Players Endpoints (Scout+ Only)

### List Saved Players

```http
GET /api/v1/saved-players
Authorization: Bearer <access_token>
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "saved_players": [
      {
        "id": "save-001",
        "player": {
          "id": "660e8400-e29b-41d4-a716-446655440001",
          "first_name": "James",
          "last_name_init": "O.",
          "age": 17,
          "position": "Midfielder",
          "thumbnail_url": "https://..."
        },
        "notes": "Strong dribbling skills",
        "saved_at": "2026-01-20T14:30:00Z"
      }
    ]
  }
}
```

---

### Save Player

```http
POST /api/v1/saved-players
Authorization: Bearer <access_token>
```

**Request Body:**
```json
{
  "player_id": "660e8400-e29b-41d4-a716-446655440001",
  "notes": "Strong dribbling skills, worth following"
}
```

**Response (201 Created):**
```json
{
  "success": true,
  "data": {
    "id": "save-001",
    "player_id": "660e8400-e29b-41d4-a716-446655440001",
    "saved_at": "2026-01-20T14:30:00Z"
  }
}
```

---

### Remove Saved Player

```http
DELETE /api/v1/saved-players/{player_id}
Authorization: Bearer <access_token>
```

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Player removed from saved list"
}
```

---

## 🔍 Search Endpoints

### Search Players

```http
GET /api/v1/search/players
```

**Query Parameters:**
| Param | Type | Description |
|-------|------|-------------|
| q | string | Full-text search query |
| position | string | Filter by position |
| country | string | Filter by country |
| age_min | int | Minimum age |
| age_max | int | Maximum age |
| page | int | Page number (default: 1) |
| limit | int | Results per page (default: 20) |

**Example:**
```http
GET /api/v1/search/players?position=midfielder&country=Nigeria&age_min=14&age_max=18
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "players": [
      {
        "id": "660e8400-e29b-41d4-a716-446655440001",
        "first_name": "James",
        "last_name_init": "O.",
        "age": 17,
        "position": "Midfielder",
        "country": "Nigeria",
        "thumbnail_url": "https://...",
        "is_verified": true,
        "has_highlights": true
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 20,
      "total": 47
    }
  }
}
```

---

### Get Featured Players

```http
GET /api/v1/players/featured
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "players": [
      {
        "id": "660e8400-e29b-41d4-a716-446655440001",
        "first_name": "James",
        "last_name_init": "O.",
        "position": "Midfielder",
        "age": 17,
        "country": "Nigeria",
        "thumbnail_url": "https://...",
        "is_verified": true
      }
    ]
  }
}
```

---

## ❌ Error Responses

### Standard Error Format

```json
{
  "success": false,
  "error": {
    "code": "ERROR_CODE",
    "message": "Human readable message",
    "details": {
      "field": "Additional context"
    }
  }
}
```

### Common Error Codes

| Code | HTTP Status | Description |
|------|-------------|-------------|
| `UNAUTHORIZED` | 401 | Missing or invalid token |
| `FORBIDDEN` | 403 | Insufficient permissions |
| `NOT_FOUND` | 404 | Resource not found |
| `VALIDATION_ERROR` | 400 | Invalid request data |
| `DUPLICATE_ENTRY` | 409 | Resource already exists |
| `RATE_LIMITED` | 429 | Too many requests |
| `INTERNAL_ERROR` | 500 | Server error |

---

## 📊 Rate Limiting

| Endpoint Type | Rate Limit |
|---------------|------------|
| Auth endpoints | 10 requests/minute |
| Read endpoints | 100 requests/minute |
| Write endpoints | 30 requests/minute |
| Upload endpoints | 10 requests/minute |

**Rate Limit Headers:**
```http
X-RateLimit-Limit: 100
X-RateLimit-Remaining: 95
X-RateLimit-Reset: 1706180400
```

---

## 🔗 Webhooks (Planned)

### Available Events

| Event | Description |
|-------|-------------|
| `highlight.linked` | New highlight linked to player profile |
| `profile.verified` | Profile verification status changed |
| `scout.contact` | Scout contacted player |
| `placement.success` | Successful player placement |

### Webhook Payload

```json
{
  "event": "highlight.linked",
  "timestamp": "2026-01-25T10:45:00Z",
  "data": {
    "highlight_id": "770e8400-e29b-41d4-a716-446655440002",
    "profile_id": "660e8400-e29b-41d4-a716-446655440001",
    "match_id": "880e8400-e29b-41d4-a716-446655440003",
    "title": "James Okoro - Skills Highlight"
  }
}
```
