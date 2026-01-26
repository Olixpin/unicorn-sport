# 🌍 Environment Configuration Guide

## Overview
Unicorn Sport requires three environments: **Local**, **Staging**, and **Production**.

---

## 🏠 LOCAL ENVIRONMENT (Development)

### Purpose
- Local development and testing
- Fast iteration and debugging
- Isolated from production data

### Tech Stack
- **Backend:** Go binary (hot reload)
- **Database:** PostgreSQL in Docker
- **Frontend:** Nuxt 3 dev server
- **Storage:** Local file system (simulated S3)
- **Payments:** Stripe test mode

### Setup Commands
```bash
# Backend setup
cd backend
make dev  # Starts Go server + PostgreSQL

# Frontend setup
cd frontend
npm install
npm run dev  # Runs on http://localhost:3000

# Database
docker-compose up -d postgres
psql -h localhost -U unicorn -d unicorn < init-databases.sql
```

### Environment Variables (`.env.local`)
```bash
# Database
DATABASE_URL=postgres://unicorn:localdev@localhost:5432/unicorn?sslmode=disable

# Authentication
JWT_SECRET=dev-jwt-secret-not-for-production-256-bits-long
JWT_REFRESH_SECRET=dev-refresh-secret-256-bits-long

# AWS (LocalStack for S3 simulation)
AWS_ACCESS_KEY_ID=test
AWS_SECRET_ACCESS_KEY=test
AWS_S3_BUCKET=unicornsport-dev
AWS_REGION=us-east-1
AWS_ENDPOINT=http://localhost:4566  # LocalStack

# Stripe (Test Mode)
STRIPE_SECRET_KEY=sk_test_your_stripe_test_key
STRIPE_WEBHOOK_SECRET=whsec_your_test_webhook_secret
STRIPE_PUBLISHABLE_KEY=pk_test_your_publishable_key

# Application
APP_ENV=local
API_BASE_URL=http://localhost:8080/api/v1
FRONTEND_URL=http://localhost:3000

# CDN (Mock for local)
CLOUDFRONT_URL=http://localhost:8080/static
```

---

## 🧪 STAGING ENVIRONMENT (Pre-Production)

### Purpose
- Pre-production testing
- Integration testing
- Performance testing
- User acceptance testing

### Tech Stack
- **Infrastructure:** AWS EC2/ECS or DigitalOcean
- **Database:** AWS RDS PostgreSQL
- **Storage:** AWS S3 (real bucket)
- **CDN:** CloudFront
- **Payments:** Stripe test mode
- **Monitoring:** Basic CloudWatch

### Infrastructure Setup
```bash
# AWS EC2 t3.medium instance ($30-50/month)
# Ubuntu 22.04 LTS
# Docker + Docker Compose

# Clone and deploy
git clone https://github.com/your-org/unicornsport.git
cd unicornsport
docker-compose -f docker-compose.staging.yml up -d
```

### Environment Variables (`.env.staging`)
```bash
# Database
DATABASE_URL=postgres://unicorn:staging_password@staging-db.unicornsport.internal:5432/unicorn?sslmode=require

# Authentication
JWT_SECRET=staging-jwt-secret-256-bits-long-unique
JWT_REFRESH_SECRET=staging-refresh-secret-256-bits-long-unique

# AWS
AWS_ACCESS_KEY_ID=AKIA...staging_key
AWS_SECRET_ACCESS_KEY=staging_secret_key
AWS_S3_BUCKET=unicornsport-staging
AWS_REGION=us-east-1

# Stripe (Test Mode)
STRIPE_SECRET_KEY=sk_test_staging_key
STRIPE_WEBHOOK_SECRET=whsec_staging_webhook
STRIPE_PUBLISHABLE_KEY=pk_test_staging_publishable

# Application
APP_ENV=staging
API_BASE_URL=https://staging-api.unicornsport.africa/api/v1
FRONTEND_URL=https://staging.unicornsport.africa

# CDN
CLOUDFRONT_URL=https://staging-cdn.unicornsport.africa
```

---

## 🚀 PRODUCTION ENVIRONMENT (Live)

### Purpose
- Live user traffic
- Real revenue generation
- High availability and performance
- Full monitoring and alerting

### Tech Stack
- **Infrastructure:** AWS EKS (Kubernetes)
- **Database:** AWS RDS PostgreSQL (Multi-AZ)
- **Storage:** AWS S3 (versioned, encrypted)
- **CDN:** CloudFront with WAF
- **Payments:** Stripe live mode
- **Monitoring:** CloudWatch + DataDog
- **Backup:** Automated RDS snapshots

### Infrastructure Setup
```yaml
# Kubernetes manifests for production
apiVersion: apps/v1
kind: Deployment
metadata:
  name: unicornsport-api
spec:
  replicas: 3
  selector:
    matchLabels:
      app: unicornsport-api
  template:
    metadata:
      labels:
        app: unicornsport-api
    spec:
      containers:
      - name: api
        image: your-registry/unicornsport:latest
        envFrom:
        - configMapRef:
            name: unicornsport-config
        - secretRef:
            name: unicornsport-secrets
        resources:
          requests:
            memory: "512Mi"
            cpu: "250m"
          limits:
            memory: "1Gi"
            cpu: "500m"
```

### Environment Variables (`.env.production`)
```bash
# Database
DATABASE_URL=postgres://unicorn:prod_password@prod-db.unicornsport.internal:5432/unicorn?sslmode=require

# Authentication
JWT_SECRET=prod-jwt-secret-256-bits-long-unique-generated
JWT_REFRESH_SECRET=prod-refresh-secret-256-bits-long-unique-generated

# AWS
AWS_ACCESS_KEY_ID=AKIA...production_key
AWS_SECRET_ACCESS_KEY=production_secret_key
AWS_S3_BUCKET=unicornsport-media
AWS_REGION=us-east-1

# Stripe (Live Mode)
STRIPE_SECRET_KEY=sk_live_production_key
STRIPE_WEBHOOK_SECRET=whsec_production_webhook
STRIPE_PUBLISHABLE_KEY=pk_live_production_publishable

# Application
APP_ENV=production
API_BASE_URL=https://api.unicornsport.africa/api/v1
FRONTEND_URL=https://unicornsport.africa

# CDN
CLOUDFRONT_URL=https://cdn.unicornsport.africa

# Monitoring
DATADOG_API_KEY=your_datadog_key
CLOUDWATCH_LOG_GROUP=/unicornsport/production
```

---

## 🔄 Environment Promotion Strategy

### Local → Staging
```bash
# Build and tag
docker build -t unicornsport:staging .
docker tag unicornsport:staging your-registry/unicornsport:staging

# Push to registry
docker push your-registry/unicornsport:staging

# Deploy to staging
kubectl set image deployment/unicornsport-api api=your-registry/unicornsport:staging
```

### Staging → Production
```bash
# Tag for production
docker tag unicornsport:staging your-registry/unicornsport:v1.0.0
docker push your-registry/unicornsport:v1.0.0

# Blue-green deployment
kubectl set image deployment/unicornsport-api-blue api=your-registry/unicornsport:v1.0.0
# Test blue environment
# Switch traffic to blue
kubectl delete deployment unicornsport-api-green
kubectl rename deployment unicornsport-api-blue unicornsport-api
```

---

## 🔒 Security Considerations

### Secrets Management
- **Local:** Plain text `.env` files (never commit)
- **Staging:** AWS Secrets Manager or Kubernetes secrets
- **Production:** AWS Secrets Manager with rotation

### Access Control
- **Local:** Developer access only
- **Staging:** Team access for testing
- **Production:** Minimal access, automated deployments

### Backup Strategy
- **Database:** Daily automated backups
- **Files:** S3 versioning enabled
- **Code:** Git tags for each release

---

## 📊 Monitoring & Alerting

### Production Metrics
- API response times (< 200ms)
- Error rates (< 1%)
- Database connection pool usage
- S3 upload/download success rates
- Stripe webhook processing

### Alerts
- API down (PagerDuty)
- High error rates
- Database connection issues
- Payment processing failures
- Storage quota exceeded</content>
<parameter name="filePath">/Users/user/scouttalent-platform/project/ENVIRONMENT_SETUP.md