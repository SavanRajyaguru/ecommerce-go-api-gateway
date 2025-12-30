# Environment Variables Setup Guide

## Overview

All services now have `.env` files for configuration. These files contain environment variables that override the default config.yaml settings.

---

## Files Created

### API Gateway
```
ecommerce-go-api-gateway/.env
```
Contains:
- Server configuration (PORT, MODE)
- All backend service URLs
- Logger configuration
- Optional: CORS, JWT, Rate limiting

### Backend Services
```
ecommerce-go-user-service/.env
ecommerce-go-product-service/.env
ecommerce-go-order-service/.env
ecommerce-go-payment-service/.env
ecommerce-go-inventory-service/.env
ecommerce-go-notification-service/.env
ecommerce-go-config-service/.env
```
Each contains:
- Server configuration
- Database configuration (PostgreSQL, MySQL, MongoDB)
- Redis configuration
- Logger configuration
- Optional: JWT, Auth settings

---

## How Environment Variables Work

### Priority Order (Highest to Lowest)

1. **Environment Variables** (`.env` file or exported)
2. **config.yaml** file
3. **Default values** in code

### Example

**config.yaml**:
```yaml
server:
  port: ":8080"
```

**.env**:
```bash
SERVER_PORT=:9000
```

**Result**: Server runs on port **9000** (env var wins!)

---

## Using .env Files

### Option 1: Load .env Automatically (Recommended)

Install godotenv:

```bash
go get github.com/joho/godotenv
```

Update `main.go`:

```go
package main

import (
    "log"
    "github.com/joho/godotenv"
    // ... other imports
)

func main() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using config.yaml or defaults")
    }

    // Rest of your code
    cfg := config.LoadConfig()
    // ...
}
```

### Option 2: Manual Export (Current Way)

**Terminal 1 - User Service**:
```bash
cd ecommerce-go-user-service

# Export variables
export SERVER_PORT=:8081
export GIN_MODE=debug

# Run service
go run main.go
```

### Option 3: Load in One Line

```bash
cd ecommerce-go-user-service
env $(cat .env | xargs) go run main.go
```

---

## Environment Variables Reference

### API Gateway (.env)

```bash
# Server
SERVER_PORT=:8080                           # Gateway listen port
SERVER_MODE=debug                           # debug or release

# Backend Services (Local Development)
SERVICES_USER_SERVICE=http://localhost:8081
SERVICES_PRODUCT_SERVICE=http://localhost:8082
SERVICES_ORDER_SERVICE=http://localhost:8083
SERVICES_PAYMENT_SERVICE=http://localhost:8084
SERVICES_INVENTORY_SERVICE=http://localhost:8085
SERVICES_NOTIFICATION_SERVICE=http://localhost:8086
SERVICES_CONFIG_SERVICE=http://localhost:8087

# Backend Services (Docker)
# SERVICES_USER_SERVICE=http://user-service:8080
# SERVICES_PRODUCT_SERVICE=http://product-service:8080
# etc...

# Logger
LOGGER_LEVEL=info                           # debug, info, warn, error

# Optional
CORS_ALLOWED_ORIGINS=http://localhost:3000  # Frontend URLs
JWT_SECRET=your-secret-key
```

### Backend Services (.env)

```bash
# Server
SERVER_PORT=:8080                           # Service listen port
SERVER_MODE=debug                           # debug or release
GIN_MODE=debug                              # Gin framework mode

# Database (PostgreSQL example)
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USER=postgres
DATABASE_PASSWORD=your-password
DATABASE_NAME=user_service_db
DATABASE_SSL_MODE=disable

# Or single URL
DATABASE_URL=postgresql://user:pass@localhost:5432/dbname

# Redis (if using)
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

# Logger
LOGGER_LEVEL=info

# External Services
CONFIG_SERVICE_URL=http://localhost:8087

# Auth
JWT_SECRET=your-secret-key
JWT_EXPIRY=24h
```

---

## Configuration Examples

### Development Environment

**API Gateway (.env)**:
```bash
SERVER_PORT=:8080
SERVER_MODE=debug
SERVICES_USER_SERVICE=http://localhost:8081
SERVICES_PRODUCT_SERVICE=http://localhost:8082
# ... etc (localhost URLs)
LOGGER_LEVEL=debug
```

**User Service (.env)**:
```bash
SERVER_PORT=:8081
GIN_MODE=debug
DATABASE_URL=postgresql://postgres:password@localhost:5432/user_db
LOGGER_LEVEL=debug
```

### Docker Development

**API Gateway (.env)**:
```bash
SERVER_PORT=:8080
SERVER_MODE=debug
SERVICES_USER_SERVICE=http://user-service:8080
SERVICES_PRODUCT_SERVICE=http://product-service:8080
# ... etc (Docker service names)
LOGGER_LEVEL=info
```

**User Service (.env)**:
```bash
SERVER_PORT=:8080
GIN_MODE=debug
DATABASE_URL=postgresql://postgres:password@postgres:5432/user_db
LOGGER_LEVEL=info
```

### Production Environment

**API Gateway (.env)**:
```bash
SERVER_PORT=:8080
SERVER_MODE=release
SERVICES_USER_SERVICE=https://user-service.production.com
SERVICES_PRODUCT_SERVICE=https://product-service.production.com
# ... etc (production URLs)
LOGGER_LEVEL=warn
JWT_SECRET=${JWT_SECRET}  # From environment
```

**User Service (.env)**:
```bash
SERVER_PORT=:8080
SERVER_MODE=release
GIN_MODE=release
DATABASE_URL=${DATABASE_URL}  # From environment/secrets
REDIS_URL=${REDIS_URL}
LOGGER_LEVEL=warn
```

---

## Using .env with Different Modes

### Local Development (Manual Terminals)

Each service reads its own `.env`:

```bash
# Terminal 1 - User Service
cd ecommerce-go-user-service
# Reads: ecommerce-go-user-service/.env
go run main.go

# Terminal 2 - Product Service
cd ecommerce-go-product-service
# Reads: ecommerce-go-product-service/.env
go run main.go

# Terminal 8 - API Gateway
cd ecommerce-go-api-gateway
# Reads: ecommerce-go-api-gateway/.env
go run cmd/api/main.go
```

### Docker Development

Docker Compose can load .env files:

**docker-compose.dev.yaml**:
```yaml
services:
  user-service:
    env_file:
      - ../ecommerce-go-user-service/.env
    # OR
    environment:
      - SERVER_PORT=${SERVER_PORT:-:8080}
```

---

## Best Practices

### 1. Never Commit .env to Git

✅ `.env` files are in `.gitignore`
❌ Don't commit secrets, passwords, API keys

### 2. Use .env.example as Template

Each service has `.env.example`:
```bash
cp .env.example .env
# Edit .env with your values
```

### 3. Different .env for Different Environments

```
.env.development    # Local development
.env.staging        # Staging environment
.env.production     # Production (use secrets manager)
```

Load specific one:
```go
godotenv.Load(".env.development")
```

### 4. Use Secrets Manager in Production

Don't use .env files in production. Use:
- **AWS Secrets Manager**
- **HashiCorp Vault**
- **Kubernetes Secrets**
- **Environment variables from CI/CD**

### 5. Validate Environment Variables

Add validation in config.go:

```go
func LoadConfig() *Config {
    // Load .env
    godotenv.Load()

    // Load config
    viper.AutomaticEnv()
    // ...

    // Validate
    if config.Server.Port == "" {
        log.Fatal("SERVER_PORT is required")
    }

    return &config
}
```

---

## Troubleshooting

### Issue 1: Variables Not Loading

**Check**:
1. Is `.env` in the correct directory?
2. Are you loading it in main.go?
3. Are variable names correct (uppercase, underscores)?

```bash
# Debug: Print loaded vars
env | grep SERVER_PORT
```

### Issue 2: Wrong Values Being Used

**Check priority**:
1. Exported env vars override .env
2. .env overrides config.yaml
3. Check for typos in variable names

```bash
# Unset exported var if needed
unset SERVER_PORT
```

### Issue 3: .env Not Found

**Solution**:
```go
// Handle missing .env gracefully
if err := godotenv.Load(); err != nil {
    log.Println("No .env file found, using defaults")
}
```

---

## Quick Start

### 1. Install godotenv (Optional but Recommended)

```bash
# In each service
cd ecommerce-go-user-service
go get github.com/joho/godotenv
```

### 2. Update main.go

```go
import "github.com/joho/godotenv"

func main() {
    godotenv.Load()  // Load .env

    cfg := config.LoadConfig()
    // ... rest of code
}
```

### 3. Configure Your .env

```bash
cd ecommerce-go-user-service
vim .env
```

Set your values:
```bash
SERVER_PORT=:8081
DATABASE_URL=postgresql://localhost:5432/user_db
```

### 4. Run Service

```bash
go run main.go
```

Service uses .env values! ✅

---

## Security Checklist

- [ ] `.env` is in `.gitignore`
- [ ] No secrets committed to git
- [ ] Production uses secrets manager, not .env files
- [ ] .env.example has no real secrets
- [ ] Team members have their own .env (not shared)
- [ ] Rotate secrets regularly

---

## Summary

### Files Created

- ✅ API Gateway: `.env` (with all service URLs)
- ✅ All 7 Services: `.env` (with DB, Redis, etc.)
- ✅ All `.gitignore` updated (won't commit .env)

### Environment Variables

- ✅ Override config.yaml settings
- ✅ Can be loaded with godotenv
- ✅ Support local, Docker, production modes
- ✅ Secure (not committed to git)

### Next Steps

1. Review each `.env` file
2. Update with your values
3. Optional: Install godotenv and update main.go
4. Run services - they'll use .env values!

Your services are now ready for environment-based configuration! 🚀
