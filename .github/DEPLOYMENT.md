# Unicorn Sport - Deployment Guide

## 🚀 Overview

This document covers the deployment architecture and procedures for Unicorn Sport. The MVP uses a **containerized single-service deployment** for simplicity and speed.

---

## 🌍 Environment Overview

| Environment | Purpose | URL | Auto-Deploy |
|-------------|---------|-----|-------------|
| **Development** | Local development | localhost:8080 | Manual |
| **Staging** | Pre-production testing | staging.unicornsport.africa | On PR merge |
| **Production** | Live users | api.unicornsport.africa | On release tag |

---

## 🏗️ MVP Infrastructure (Current)

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                       MVP INFRASTRUCTURE                                     │
│                  (Simple, Fast, Cost-Effective)                              │
└─────────────────────────────────────────────────────────────────────────────┘

                         ┌─────────────────────┐
                         │     Cloudflare      │
                         │  • DNS management   │
                         │  • SSL termination  │
                         │  • DDoS protection  │
                         │  • CDN caching      │
                         └──────────┬──────────┘
                                    │
┌───────────────────────────────────┼─────────────────────────────────────────┐
│                    AWS / VPS       │                                          │
│                                   ▼                                          │
│  ┌────────────────────────────────────────────────────────────────────────┐ │
│  │                     DOCKER HOST (Single Server)                         │ │
│  │                                                                         │ │
│  │  ┌─────────────────────────────────────────────────────────────────┐   │ │
│  │  │                    docker-compose                                │   │ │
│  │  │                                                                  │   │ │
│  │  │   ┌─────────────────┐   ┌─────────────────┐                     │   │ │
│  │  │   │   API Server    │   │   PostgreSQL    │                     │   │ │
│  │  │   │   (Go Binary)   │   │      :5432      │                     │   │ │
│  │  │   │     :8080       │   │                 │                     │   │ │
│  │  │   │                 │   │  • Single DB    │                     │   │ │
│  │  │   │  • Auth         │   │  • All tables   │                     │   │ │
│  │  │   │  • Profiles     │   │  • Backups      │                     │   │ │
│  │  │   │  • Media        │   │                 │                     │   │ │
│  │  │   │  • Admin        │   └─────────────────┘                     │   │ │
│  │  │   │  • Subscriptions│                                           │   │ │
│  │  │   │  • Search       │                                           │   │ │
│  │  │   └─────────────────┘                                           │   │ │
│  │  │                                                                  │   │ │
│  │  └─────────────────────────────────────────────────────────────────┘   │ │
│  │                                                                         │ │
│  └────────────────────────────────────────────────────────────────────────┘ │
│                                    │                                          │
└────────────────────────────────────┼────────────────────────────────────────┘
                                     │
                    ┌────────────────┴────────────────┐
                    │                                 │
          ┌─────────▼─────────┐           ┌──────────▼──────────┐
          │       AWS S3      │           │    Stripe API       │
          │                   │           │                     │
          │  • Videos         │           │  • Subscriptions    │
          │  • Thumbnails     │           │  • Webhooks         │
          │  • Documents      │           │  • Payments         │
          └───────────────────┘           └─────────────────────┘
```

### MVP Hosting Options

| Option | Monthly Cost | Recommendation |
|--------|--------------|----------------|
| **AWS EC2 t3.medium** | $30-50 | ✅ Recommended - simple, scales |
| **AWS ECS Fargate** | $40-80 | Container-native, serverless |
| **DigitalOcean Droplet** | $20-50 | Good for cost savings |
| **Hetzner VPS** | €10-20 | Best value in Europe |

### Why NOT Kubernetes for MVP?

| Kubernetes | Docker Compose |
|------------|----------------|
| Weeks to set up properly | Hours to deploy |
| $100+/month minimum | $20-50/month |
| Team of 3+ to operate | Solo developer friendly |
| Overkill for < 10k users | Perfect for MVP scale |

---

## 📦 Container Configuration

### Dockerfile

```dockerfile
# Multi-stage build for minimal image size
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server ./cmd/server

# Runtime stage
FROM alpine:3.19

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/migrations ./migrations

EXPOSE 8080

CMD ["./server"]
```

### docker-compose.yml (Production)

```yaml
version: '3.8'

services:
  api:
    image: ghcr.io/unicornsport/api:${VERSION:-latest}
    ports:
      - "8080:8080"
    environment:
      - DATABASE_URL=postgres://unicorn:${DB_PASSWORD}@db:5432/unicorn?sslmode=disable
      - JWT_SECRET=${JWT_SECRET}
      - AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID}
      - AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY}
      - AWS_S3_BUCKET=${AWS_S3_BUCKET}
      - AWS_REGION=${AWS_REGION}
      - STRIPE_SECRET_KEY=${STRIPE_SECRET_KEY}
      - STRIPE_WEBHOOK_SECRET=${STRIPE_WEBHOOK_SECRET}
      - GIN_MODE=release
    depends_on:
      db:
        condition: service_healthy
    restart: unless-stopped
    healthcheck:
      test: ["CMD", "wget", "-q", "--spider", "http://localhost:8080/health"]
      interval: 30s
      timeout: 10s
      retries: 3

  db:
    image: postgres:16-alpine
    volumes:
      - postgres_data:/var/lib/postgresql/data
      - ./backups:/backups
    environment:
      - POSTGRES_DB=unicorn
      - POSTGRES_USER=unicorn
      - POSTGRES_PASSWORD=${DB_PASSWORD}
    restart: unless-stopped
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U unicorn"]
      interval: 10s
      timeout: 5s
      retries: 5

volumes:
  postgres_data:
```

### docker-compose.yml (Development)

```yaml
version: '3.8'

services:
  api:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    environment:
      - DATABASE_URL=postgres://unicorn:localdev@db:5432/unicorn?sslmode=disable
      - JWT_SECRET=dev-secret-not-for-production
      - GIN_MODE=debug
    volumes:
      - .:/app  # Hot reload
    depends_on:
      - db

  db:
    image: postgres:16-alpine
    ports:
      - "5432:5432"
    environment:
      - POSTGRES_DB=unicorn
      - POSTGRES_USER=unicorn
      - POSTGRES_PASSWORD=localdev
    volumes:
      - postgres_dev:/var/lib/postgresql/data

volumes:
  postgres_dev:
```

---

## 🚢 Deployment Procedures

### Initial Setup (One-Time)

```bash
# 1. Provision server (AWS EC2 or VPS)

# 2. Install Docker (if VPS)
curl -fsSL https://get.docker.com | sh

# 3. Clone repository
git clone git@github.com:unicornsport/backend.git
cd backend

# 4. Create environment file
cp .env.example .env
# Edit .env with production values

# 5. Run migrations
docker-compose run --rm api ./server migrate up

# 6. Start services
docker-compose up -d

# 7. Configure Cloudflare DNS
# A record: api.unicornsport.africa → server IP
```

### Deploy New Version

```bash
# On server
cd /app/backend

# Pull latest
git pull origin main

# Pull new image
docker-compose pull

# Deploy with zero downtime
docker-compose up -d --no-deps api

# Check logs
docker-compose logs -f api
```

### Rollback

```bash
# If something goes wrong
docker-compose down
docker-compose pull api:previous-version
docker-compose up -d
```

---

## 🔐 Environment Variables

```bash
# .env.production

# Server
PORT=8080
GIN_MODE=release

# Database (single database)
DATABASE_URL=postgres://unicorn:STRONG_PASSWORD@db:5432/unicorn?sslmode=disable

# JWT
JWT_SECRET=generate-256-bit-random-string
JWT_ACCESS_EXPIRY=15m
JWT_REFRESH_EXPIRY=168h

# AWS S3
AWS_ACCESS_KEY_ID=AKIAXXXXXXXXXXXXXXXX
AWS_SECRET_ACCESS_KEY=xxxxx
AWS_S3_BUCKET=unicornsport-media
AWS_S3_DOCUMENTS_BUCKET=unicornsport-documents
AWS_REGION=eu-west-1
AWS_CLOUDFRONT_URL=https://cdn.unicornsport.africa

# Stripe
STRIPE_SECRET_KEY=sk_live_xxxxx
STRIPE_WEBHOOK_SECRET=whsec_xxxxx
STRIPE_SCOUT_PRICE_ID=price_xxxxx
STRIPE_PRO_PRICE_ID=price_xxxxx
STRIPE_CLUB_PRICE_ID=price_xxxxx

# URLs
FRONTEND_URL=https://unicornsport.africa
```

---

## 📊 Monitoring (MVP)

### Essential Monitoring (Free/Cheap)

| Tool | Purpose | Cost |
|------|---------|------|
| **UptimeRobot** | Uptime alerts | Free |
| **Sentry** | Error tracking | Free tier |
| **Docker logs** | Application logs | Free |
| **pg_stat_statements** | Slow queries | Built-in |

### Health Check Endpoint

```go
// GET /health
{
  "status": "healthy",
  "version": "1.0.0",
  "database": "connected",
  "uptime": "24h15m"
}
```

### Log Aggregation (Simple)

```bash
# View logs
docker-compose logs -f api

# Export logs
docker-compose logs api > logs/api-$(date +%Y%m%d).log
```

---

## 💾 Backup Strategy

### Database Backups

```bash
# backup.sh - Run daily via cron
#!/bin/bash

BACKUP_DIR="/backups"
DATE=$(date +%Y%m%d_%H%M%S)
BACKUP_FILE="unicorn_${DATE}.sql.gz"

# Dump and compress
docker-compose exec -T db pg_dump -U unicorn unicorn | gzip > ${BACKUP_DIR}/${BACKUP_FILE}

# Upload to AWS S3
aws s3 cp ${BACKUP_DIR}/${BACKUP_FILE} s3://unicornsport-backups/db/${BACKUP_FILE}

# Keep only last 7 days locally
find ${BACKUP_DIR} -name "*.sql.gz" -mtime +7 -delete
```

### Cron Setup

```bash
# crontab -e
0 3 * * * /app/scripts/backup.sh >> /var/log/backup.log 2>&1
```

---

## 🔮 Future: Kubernetes Migration

When you have 50k+ users, 5+ engineers, and need auto-scaling:

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                    FUTURE: AWS EKS (ELASTIC KUBERNETES SERVICE)             │
└─────────────────────────────────────────────────────────────────────────────┘

See original Kubernetes architecture (preserved below) for the full 
microservices deployment. Migrate when:

• Team grows to 5+ engineers
• Need independent service scaling
• Hit database bottlenecks requiring sharding
• Require 99.99% uptime SLA
```

```dockerfile
# Example: backend/services/auth-service/Dockerfile
# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o auth-service ./cmd/server

# Runtime stage
FROM alpine:3.19

RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app

COPY --from=builder /app/auth-service .
COPY --from=builder /app/migrations ./migrations

EXPOSE 8080
CMD ["./auth-service"]
```

### Image Registry

```
AWS ECR: 123456789.dkr.ecr.eu-west-1.amazonaws.com

Images:
├── unicornsport/api:v1.2.3
├── unicornsport/frontend:v1.2.3
└── (future microservices when needed)
```

---

## ⚙️ Kubernetes Resources

### Namespace Structure

```yaml
# namespaces.yaml
apiVersion: v1
kind: Namespace
metadata:
  name: unicornsport-prod
  labels:
    environment: production
---
apiVersion: v1
kind: Namespace
metadata:
  name: unicornsport-staging
  labels:
    environment: staging
```

### Deployment Example

```yaml
# auth-service/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: auth-service
  namespace: unicornsport-prod
  labels:
    app: auth-service
    version: v1
spec:
  replicas: 3
  selector:
    matchLabels:
      app: auth-service
  template:
    metadata:
      labels:
        app: auth-service
        version: v1
    spec:
      containers:
      - name: api
        image: 123456789.dkr.ecr.eu-west-1.amazonaws.com/unicornsport/api:v1.2.3
        ports:
        - containerPort: 8080
        env:
        - name: DATABASE_URL
          valueFrom:
            secretKeyRef:
              name: auth-secrets
              key: database-url
        - name: JWT_SECRET
          valueFrom:
            secretKeyRef:
              name: auth-secrets
              key: jwt-secret
        - name: NATS_URL
          value: "nats://nats-cluster:4222"
        resources:
          requests:
            cpu: "100m"
            memory: "128Mi"
          limits:
            cpu: "500m"
            memory: "512Mi"
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 10
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /ready
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
```

### Service Definition

```yaml
# auth-service/service.yaml
apiVersion: v1
kind: Service
metadata:
  name: auth-service
  namespace: unicornsport-prod
spec:
  selector:
    app: auth-service
  ports:
  - port: 80
    targetPort: 8080
  type: ClusterIP
```

### Horizontal Pod Autoscaler

```yaml
# auth-service/hpa.yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: auth-service-hpa
  namespace: unicornsport-prod
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: auth-service
  minReplicas: 3
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 70
  - type: Resource
    resource:
      name: memory
      target:
        type: Utilization
        averageUtilization: 80
```

### Ingress Configuration

```yaml
# ingress.yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: unicornsport-ingress
  namespace: unicornsport-prod
  annotations:
    kubernetes.io/ingress.class: nginx
    cert-manager.io/cluster-issuer: letsencrypt-prod
    nginx.ingress.kubernetes.io/rate-limit: "100"
    nginx.ingress.kubernetes.io/rate-limit-window: "1m"
spec:
  tls:
  - hosts:
    - api.unicornsport.africa
    - www.unicornsport.africa
    secretName: unicornsport-tls
  rules:
  - host: api.unicornsport.africa
    http:
      paths:
      - path: /api/v1/auth
        pathType: Prefix
        backend:
          service:
            name: auth-service
            port:
              number: 80
      - path: /api/v1/profiles
        pathType: Prefix
        backend:
          service:
            name: profile-service
            port:
              number: 80
      - path: /api/v1/videos
        pathType: Prefix
        backend:
          service:
            name: media-service
            port:
              number: 80
      - path: /api/v1/discover
        pathType: Prefix
        backend:
          service:
            name: discovery-service
            port:
              number: 80
  - host: www.unicornsport.africa
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: frontend
            port:
              number: 80
```

---

## 🔄 CI/CD Pipeline

### GitHub Actions Workflow

```yaml
# .github/workflows/deploy.yml
name: Build and Deploy

on:
  push:
    branches: [main]
    tags: ['v*']
  pull_request:
    branches: [main]

env:
  AWS_REGION: eu-west-1
  ECR_REGISTRY: 123456789.dkr.ecr.eu-west-1.amazonaws.com
  EKS_CLUSTER: unicornsport-eks

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.23'
      
      - name: Run tests
        run: |
          cd backend
          go test ./...

  build:
    needs: test
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Configure AWS credentials
        uses: aws-actions/configure-aws-credentials@v4
        with:
          aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
          aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
          aws-region: ${{ env.AWS_REGION }}
      
      - name: Login to Amazon ECR
        id: login-ecr
        uses: aws-actions/amazon-ecr-login@v2
      
      - name: Build and push
        run: |
          docker build -t ${{ env.ECR_REGISTRY }}/unicornsport/api:${{ github.sha }} \
            -f backend/Dockerfile backend
          docker push ${{ env.ECR_REGISTRY }}/unicornsport/api:${{ github.sha }}

  deploy-staging:
    needs: build
    if: github.event_name == 'push' && github.ref == 'refs/heads/main'
    runs-on: ubuntu-latest
    environment: staging
    steps:
      - uses: actions/checkout@v4
      
      - name: Configure AWS credentials
        uses: aws-actions/configure-aws-credentials@v4
        with:
          aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
          aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
          aws-region: ${{ env.AWS_REGION }}
      
      - name: Update kubeconfig
        run: aws eks update-kubeconfig --name ${{ env.EKS_CLUSTER }} --region ${{ env.AWS_REGION }}
      
      - name: Deploy to staging
        run: |
          kubectl set image deployment/api \
            api=${{ env.ECR_REGISTRY }}/unicornsport/api:${{ github.sha }} \
            -n unicornsport-staging

  deploy-production:
    needs: build
    if: startsWith(github.ref, 'refs/tags/v')
    runs-on: ubuntu-latest
    environment: production
    steps:
      - uses: actions/checkout@v4
      
      - name: Configure AWS credentials
        uses: aws-actions/configure-aws-credentials@v4
        with:
          aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
          aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
          aws-region: ${{ env.AWS_REGION }}
      
      - name: Update kubeconfig
        run: aws eks update-kubeconfig --name ${{ env.EKS_CLUSTER }} --region ${{ env.AWS_REGION }}
      
      - name: Deploy to production
        run: |
          kubectl set image deployment/api \
            api=${{ env.ECR_REGISTRY }}/unicornsport/api:${{ github.ref_name }} \
            -n unicornsport-prod
```

---

## 🗄️ Database Deployment

### AWS RDS PostgreSQL

```hcl
# terraform/database.tf
resource "aws_db_instance" "main" {
  identifier              = "unicornsport-db"
  engine                  = "postgres"
  engine_version          = "16"
  instance_class          = "db.t3.medium"
  allocated_storage       = 50
  storage_type            = "gp3"
  
  db_name                 = "unicornsport"
  username                = var.db_admin_username
  password                = var.db_admin_password
  
  backup_retention_period = 35
  multi_az               = true
  
  vpc_security_group_ids = [aws_security_group.db.id]
  db_subnet_group_name   = aws_db_subnet_group.main.name
  
  skip_final_snapshot    = false
  final_snapshot_identifier = "unicornsport-final-snapshot"
}

# Single database for modular monolith
# All tables in one database with module prefixes
```

### Database Migrations

```bash
# Run migrations via Kubernetes Job
kubectl apply -f - <<EOF
apiVersion: batch/v1
kind: Job
metadata:
  name: auth-migrate-$(date +%s)
  namespace: unicornsport-prod
spec:
  template:
    spec:
      containers:
      - name: migrate
        image: 123456789.dkr.ecr.eu-west-1.amazonaws.com/unicornsport/api:v1.2.3
        command: ["./server", "migrate", "up"]
        env:
        - name: DATABASE_URL
          valueFrom:
            secretKeyRef:
              name: auth-secrets
              key: database-url
      restartPolicy: Never
  backoffLimit: 3
EOF
```

---

## 📊 Monitoring & Observability

### AWS CloudWatch + Prometheus

```yaml
# monitoring/prometheus-config.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: prometheus-config
  namespace: monitoring
data:
  prometheus.yml: |
    global:
      scrape_interval: 15s
    
    scrape_configs:
      - job_name: 'kubernetes-pods'
        kubernetes_sd_configs:
          - role: pod
        relabel_configs:
          - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_scrape]
            action: keep
            regex: true
```

### Application Insights

```go
// pkg/telemetry/appinsights.go
package telemetry

import (
    "github.com/microsoft/ApplicationInsights-Go/appinsights"
)

func InitTelemetry(instrumentationKey string) appinsights.TelemetryClient {
    config := appinsights.NewTelemetryConfiguration(instrumentationKey)
    config.MaxBatchSize = 1000
    config.MaxBatchInterval = 10 * time.Second
    
    return appinsights.NewTelemetryClientFromConfig(config)
}
```

### Grafana Dashboards

Key dashboards:
- **Service Health**: Response times, error rates, throughput
- **Infrastructure**: CPU, memory, disk usage
- **Business Metrics**: Signups, videos uploaded, verifications
- **Security**: Login attempts, failed auth, rate limiting

---

## 🔄 Rollback Procedures

### Kubernetes Rollback

```bash
# View deployment history
kubectl rollout history deployment/auth-service -n unicornsport-prod

# Rollback to previous version
kubectl rollout undo deployment/auth-service -n unicornsport-prod

# Rollback to specific revision
kubectl rollout undo deployment/auth-service -n unicornsport-prod --to-revision=3
```

### Database Rollback

```bash
# Run down migration
kubectl exec -it auth-service-pod -n unicornsport-prod -- \
  ./migrate -path migrations -database "$DATABASE_URL" down 1
```

---

## 🔒 Secrets Management

### AWS Secrets Manager Integration

```yaml
# secrets/external-secrets.yaml
apiVersion: external-secrets.io/v1beta1
kind: ExternalSecret
metadata:
  name: api-secrets
  namespace: unicornsport-prod
spec:
  refreshInterval: 1h
  secretStoreRef:
    kind: ClusterSecretStore
    name: aws-secrets-manager
  target:
    name: api-secrets
  data:
  - secretKey: database-url
    remoteRef:
      key: unicornsport/database-url
  - secretKey: jwt-secret
    remoteRef:
      key: unicornsport/jwt-signing-key
```

---

## 📋 Deployment Checklist

### Pre-Deployment

- [ ] All tests passing
- [ ] Security scan completed
- [ ] Database migrations tested in staging
- [ ] Rollback procedure documented
- [ ] On-call engineer notified

### During Deployment

- [ ] Deploy to staging first
- [ ] Run smoke tests
- [ ] Check metrics for anomalies
- [ ] Deploy to production (canary if major change)
- [ ] Verify health endpoints
- [ ] Check error rates

### Post-Deployment

- [ ] Monitor for 30 minutes
- [ ] Verify key user journeys
- [ ] Check database performance
- [ ] Update deployment documentation
- [ ] Notify stakeholders

---

## 📞 On-Call Procedures

### Escalation Matrix

| Severity | Description | Response Time | Escalation |
|----------|-------------|---------------|------------|
| P1 | Service down | 15 minutes | CTO immediately |
| P2 | Major feature broken | 1 hour | Engineering lead |
| P3 | Minor issue | 4 hours | On-call engineer |
| P4 | Non-urgent | Next business day | Team queue |

### Runbooks Location

```
/docs/runbooks/
├── service-down.md
├── database-issues.md
├── high-cpu.md
├── rate-limiting.md
└── security-incident.md
```
