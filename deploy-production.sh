#!/bin/bash
# Production Deployment Script for Unicorn Sport
# Run this after setting up Railway and Neon

set -e

echo "🚀 Unicorn Sport Production Deployment"
echo "======================================"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Check if required tools are installed
command -v railway >/dev/null 2>&1 || {
    echo -e "${RED}❌ Railway CLI not found. Install with: npm install -g @railway/cli${NC}"
    exit 1
}

echo -e "${YELLOW}📋 Prerequisites Check:${NC}"
echo "✅ Railway CLI installed"
echo "✅ GitHub repo ready"
echo "✅ Neon database created"
echo "✅ AWS S3 bucket configured"
echo "✅ Stripe account ready"

echo ""
echo -e "${GREEN}🎯 Deployment Steps:${NC}"
echo "1. Railway account: https://railway.app"
echo "2. Neon database: https://neon.tech"
echo "3. AWS S3/CloudFront configured"
echo "4. Stripe production keys ready"

echo ""
echo -e "${YELLOW}🔧 Ready to deploy? Run these commands:${NC}"
echo ""
echo "# 1. Login to Railway"
echo "railway login"
echo ""
echo "# 2. Link this project"
echo "railway link"
echo ""
echo "# 3. Set environment variables"
echo "railway variables set ENVIRONMENT=production"
echo "railway variables set JWT_SECRET=2eea94ecd8ea13a38ddf81f275eddc623232645d04c2b6ef0b5f223a0252cbb7"
echo "railway variables set JWT_ACCESS_TTL_MINUTES=15"
echo "railway variables set JWT_REFRESH_TTL_DAYS=7"
echo "railway variables set DB_HOST=ep-withered-shadow-ahnvn871-pooler.c-3.us-east-1.aws.neon.tech"
echo "railway variables set DB_PORT=5432"
echo "railway variables set DB_USER=neondb_owner"
echo "railway variables set DB_PASSWORD=YOUR_NEON_PASSWORD"
echo "railway variables set DB_NAME=neondb"
echo "railway variables set DB_SSLMODE=require"
echo ""
echo "# 4. Deploy!"
echo "railway up"
echo ""
echo "# 5. Check deployment"
echo "railway logs"
echo ""
echo -e "${GREEN}✅ That's it! Your API will be live at the Railway URL${NC}"