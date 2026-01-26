# Unicorn Sport - Verification System

## 🎯 Overview

The **Verification System** is the heart of Unicorn Sport's value proposition. Age verification happens **BEFORE** tournament participation using official government documents (NIN or International Passport). Only verified players are approved to compete.

---

## 🔑 Core Philosophy

> **"Verify first, compete second."**

Unlike platforms that verify after registration, we verify age BEFORE players participate in tournaments. This ensures every player on our platform has been verified from the start.

---

## 📋 How Verification Works

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         VERIFICATION PROCESS                                 │
└─────────────────────────────────────────────────────────────────────────────┘

Step 1: Player/School Requests Tournament Participation
┌─────────────────────────────────────────────────────────────────────────────┐
│  • Player or school contacts academy                                         │
│  • Requests to participate in upcoming tournament                           │
└─────────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
Step 2: Document Submission
┌─────────────────────────────────────────────────────────────────────────────┐
│  School/Player presents official documents:                                  │
│                                                                              │
│  ACCEPTED DOCUMENTS:                                                         │
│  ┌─────────────────────────────────────────────────────────────────────┐    │
│  │  ✓ NIN (National Identification Number) - Nigeria                   │    │
│  │  ✓ International Passport                                           │    │
│  └─────────────────────────────────────────────────────────────────────┘    │
│                                                                              │
│  Document must show:                                                         │
│  • Full legal name                                                          │
│  • Date of birth                                                            │
│  • Photo ID                                                                 │
└─────────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
Step 3: Academy Verifies Age
┌─────────────────────────────────────────────────────────────────────────────┐
│  Academy staff manually verifies:                                            │
│  • Document is authentic (not forged)                                       │
│  • Date of birth matches tournament age requirements                        │
│  • Photo matches the player                                                 │
│                                                                              │
│  Result:                                                                     │
│  ✅ APPROVED - Player cleared to participate                                │
│  ❌ REJECTED - Age doesn't meet requirements or document issues            │
└─────────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
Step 4: Player Participates (Only if Verified)
┌─────────────────────────────────────────────────────────────────────────────┐
│  • Verified player competes in tournament                                   │
│  • Media team records matches                                               │
│  • Performance data collected                                               │
└─────────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
Step 5: Profile Created with Verified Status
┌─────────────────────────────────────────────────────────────────────────────┐
│  Profile automatically shows:                                                │
│  ✅ "Age Verified" badge                                                    │
│  ✅ Verification method (NIN/Passport)                                      │
│  ✅ Tournament participation proof                                          │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 📊 Verification Status

Since all players are verified BEFORE participating, profiles have simple status:

| Status | Meaning | Display |
|--------|---------|---------|
| **✅ Verified** | Age confirmed via NIN/Passport before tournament | Green badge |
| **Document Type** | NIN or International Passport | Shown on profile |

> **Note**: There are no "unverified" players on the platform. If a player wasn't verified, they couldn't participate and therefore don't have a profile.

---

## 🏫 School Partnership Model

While primary verification is document-based (NIN/Passport), school partnerships add value:

| School Role | Benefit |
|-------------|---------|
| **Referral Source** | Schools send students to participate in tournaments |
| **Group Registration** | Schools can register multiple students at once |
| **Additional Credibility** | "Attends [School Name]" adds trust |
| **Future Verification** | Schools can provide enrollment confirmation if needed |

---

## 🏫 School Partnership Model (P2 - Phase 2)

> **Note:** School portal is a P2 feature. For MVP, schools are referral partners only.
> Primary verification is via NIN/Passport presented in person before tournament.

### MVP Approach

For MVP, schools function as **referral partners**:
1. Schools send students to participate in tournaments
2. Academy verifies age via NIN/Passport on-site
3. School name recorded on player profile for credibility

### Future School Portal (P2)

When implemented, schools will be able to:

| Feature | Description |
|---------|-------------|
| **Student Roster** | View all linked students |
| **Bulk Registration** | Register multiple students for tournaments |
| **Scout Notifications** | Get notified when scouts view their students |
| **Success Tracking** | Track student recruitment outcomes |

---

## 📋 Verification Data Points

### What We Verify

| Data Point | Source | Method | Trust Impact |
|------------|--------|--------|--------------|
| **Name** | Student claim | School confirmation | ⭐⭐ |
| **Date of Birth** | Birth certificate | School records + Document | ⭐⭐⭐⭐ |
| **Current Age** | Calculated | DOB verification | ⭐⭐⭐⭐ |
| **School Enrollment** | School records | School admin confirmation | ⭐⭐⭐ |
| **Grade/Class** | School records | School admin confirmation | ⭐⭐ |
| **Height** | Self-reported | School measurement (optional) | ⭐ |
| **Weight** | Self-reported | School measurement (optional) | ⭐ |
| **Playing Position** | Self-reported | Video evidence | ⭐ |
| **Location** | Self-reported | School address verification | ⭐⭐ |

### Document Types Accepted

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                      ACCEPTED VERIFICATION DOCUMENTS                         │
└─────────────────────────────────────────────────────────────────────────────┘

Primary (Required for Level 3):
├── Birth Certificate
│   ├── Government-issued original
│   ├── Scanned copy (high resolution)
│   └── School administrator attestation
│
├── School Enrollment Record
│   ├── Current enrollment letter
│   ├── Student ID card
│   └── Class register entry
│
└── Photo ID
    ├── National ID (if available)
    ├── School ID with photo
    └── Passport (if available)

Secondary (Strengthens verification):
├── Medical Records
│   ├── Hospital birth records
│   └── Growth chart records
│
├── Immunization Records
│   └── With date of birth listed
│
├── Baptism/Naming Certificate
│   └── Religious institution records
│
└── Previous School Records
    └── Enrollment history
```

---

## 🔄 Verification Workflow

### Player Verification Request Flow

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                      VERIFICATION REQUEST WORKFLOW                           │
└─────────────────────────────────────────────────────────────────────────────┘

Player                    Platform                    School Admin
  │                          │                             │
  │  1. Request Verification │                             │
  │  (selects school)        │                             │
  │─────────────────────────▶│                             │
  │                          │                             │
  │                          │  2. Create verification     │
  │                          │     request record          │
  │                          │                             │
  │                          │  3. Notify school admin     │
  │                          │─────────────────────────────▶│
  │                          │                             │
  │                          │                             │  4. Review request
  │                          │                             │     • Check enrollment
  │                          │                             │     • Verify DOB
  │                          │                             │     • Check records
  │                          │                             │
  │                          │  5. Submit decision         │
  │                          │◀─────────────────────────────│
  │                          │     (approve/deny/flag)     │
  │                          │                             │
  │                          │  6. If approved:            │
  │                          │     • Update trust level    │
  │                          │     • Add verification badge│
  │                          │     • Log verification      │
  │                          │                             │
  │  7. Notification         │                             │
  │◀─────────────────────────│                             │
  │  (verification result)   │                             │
  │                          │                             │
```

### Verification States

```
┌────────────┐    ┌────────────┐    ┌────────────┐    ┌────────────┐
│  PENDING   │───▶│  REVIEW    │───▶│  APPROVED  │    │  REJECTED  │
│            │    │            │    │            │    │            │
│ Request    │    │ School is  │    │ Verified   │    │ Failed     │
│ submitted  │    │ reviewing  │    │ by school  │    │ checks     │
└────────────┘    └─────┬──────┘    └────────────┘    └────────────┘
                        │                                    ▲
                        │           ┌────────────┐           │
                        └──────────▶│  FLAGGED   │───────────┘
                                    │            │
                                    │ Needs more │
                                    │ documents  │
                                    └────────────┘
```

---

## 🛡️ Fraud Prevention

### Detection Mechanisms

| Mechanism | Description | Action |
|-----------|-------------|--------|
| **Duplicate Detection** | Same person with multiple accounts | Account merge or suspension |
| **Document Review** | Manual document authenticity check | Admin review queue |
| **Cross-Reference** | Same DOB/name across accounts | Flag for investigation |
| **Photo Matching** | Visual check across profiles | Alert and review |
| **School Anomalies** | Too many students from small school | School audit |
| **School Verification** | School confirms student enrollment & age | Trust level upgrade |

### Reporting System

```
Players/Scouts can report:
├── Suspected age fraud
├── Fake documents
├── Impersonation
├── Multiple accounts
└── False claims

Reports trigger:
├── Investigation queue
├── Account flag
├── School notification (if linked)
└── Potential suspension
```

---

## 📊 Data Model (Verification)

### Proposed Schema Extensions

```sql
-- Verification requests table
CREATE TABLE verification_requests (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    profile_id UUID NOT NULL REFERENCES profiles(id),
    school_id UUID NOT NULL REFERENCES schools(id),
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    requested_at TIMESTAMP NOT NULL DEFAULT NOW(),
    reviewed_at TIMESTAMP,
    reviewed_by UUID,  -- School admin user_id
    result JSONB,
    notes TEXT,
    
    CONSTRAINT status_check CHECK (status IN 
        ('pending', 'in_review', 'approved', 'rejected', 'flagged'))
);

-- Schools table (verification partners)
CREATE TABLE schools (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(200) NOT NULL,
    registration_number VARCHAR(100),
    country VARCHAR(100) NOT NULL,
    region VARCHAR(100),
    city VARCHAR(100),
    address TEXT,
    type VARCHAR(50) NOT NULL,  -- 'secondary', 'academy', 'club'
    verified BOOLEAN DEFAULT FALSE,
    verified_at TIMESTAMP,
    admin_user_id UUID REFERENCES users(id),
    contact_email VARCHAR(255),
    contact_phone VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    
    CONSTRAINT type_check CHECK (type IN 
        ('secondary', 'primary', 'academy', 'club', 'university'))
);

-- Verification documents table
CREATE TABLE verification_documents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    request_id UUID NOT NULL REFERENCES verification_requests(id),
    document_type VARCHAR(50) NOT NULL,
    blob_url TEXT NOT NULL,
    uploaded_by UUID NOT NULL REFERENCES users(id),
    verified BOOLEAN DEFAULT FALSE,
    verification_notes TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    
    CONSTRAINT doc_type_check CHECK (document_type IN 
        ('birth_certificate', 'school_enrollment', 'national_id', 
         'passport', 'medical_record', 'school_id'))
);

-- School-student links
CREATE TABLE school_students (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    school_id UUID NOT NULL REFERENCES schools(id),
    profile_id UUID NOT NULL REFERENCES profiles(id),
    enrollment_date DATE,
    current_grade VARCHAR(20),
    verified BOOLEAN DEFAULT FALSE,
    verified_at TIMESTAMP,
    verified_by UUID,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    
    UNIQUE(school_id, profile_id)
);
```

---

## 🎯 Implementation Phases

### Phase 1: Foundation (Current Sprint)
- [ ] Design verification request API
- [ ] Create schools table and CRUD
- [ ] Build school registration flow
- [ ] Implement basic trust level system

### Phase 2: School Portal
- [ ] Build school admin dashboard
- [ ] Create student linking workflow
- [ ] Implement verification queue
- [ ] Add document upload capability

### Phase 3: Document Verification
- [ ] Integrate document storage
- [ ] Build document viewer
- [ ] Add manual review tools
- [ ] Create audit trail

### Phase 4: Advanced Features
- [ ] AI document analysis
- [ ] Face matching (optional)
- [ ] MRI verification partnership
- [ ] Federation integration

---

## 📈 Success Metrics

| Metric | Target | Measurement |
|--------|--------|-------------|
| Verification request turnaround | < 48 hours | Time from request to decision |
| Verification approval rate | > 70% | Approved / Total requests |
| School partner satisfaction | > 4.5/5 | Quarterly surveys |
| Fraud detection rate | > 95% | Fraud caught / Total fraud |
| False rejection rate | < 5% | Legitimate rejections appealed |

---

## 🔒 Privacy & Compliance

### Data Protection

- All documents encrypted at rest (AES-256)
- Documents never shared with scouts (only verification status)
- GDPR-compliant data handling
- Right to deletion supported
- Audit logs for all access

### Consent Management

- Explicit consent required for verification
- Parent/guardian consent for minors
- Clear data usage explanation
- Opt-out available (affects trust level)
