# Development Mode Setup - Hot Reload for All Services

## What I've Created

I've set up **hot reload development mode** for ALL 8 services (7 backend + 1 gateway).

### Files Created (Per Service)

Each service now has:
- ✅ `Dockerfile.dev` - Development Docker image with Air
- ✅ `.air.toml` - Air configuration for hot reload

### Services Updated

1. ✅ ecommerce-go-user-service
2. ✅ ecommerce-go-product-service
3. ✅ ecommerce-go-order-service
4. ✅ ecommerce-go-payment-service
5. ✅ ecommerce-go-inventory-service
6. ✅ ecommerce-go-notification-service
7. ✅ ecommerce-go-config-service
8. ✅ ecommerce-go-api-gateway

---

## How It Works

### With Hot Reload (Development Mode)

**Before**:
```bash
# Edit code
vim user-service/main.go

# Stop service (Ctrl+C)
# Restart service
go run main.go

# Wait 2-3 seconds
```

**After (Hot Reload)**:
```bash
# Edit code
vim user-service/main.go

# Save file
# ✨ Air automatically detects change and reloads (1-2 sec)
# No manual restart needed!
```

---

## Two Ways to Use Development Mode

### Option 1: Manual Terminals (What You're Doing Now)

**No changes needed!** Continue running services manually:

```bash
# Terminal 1 - User Service
cd ecommerce-go-user-service
SERVER_PORT=:8081 go run main.go

# Terminal 2 - Product Service
cd ecommerce-go-product-service
SERVER_PORT=:8082 go run main.go

# ... and so on
```

**Development Files Created**: For future use if you want hot reload.

---

### Option 2: Docker with Hot Reload (New Option)

**Use hot reload** - code changes auto-reload!

#### Step 1: Start Development Environment

From the API Gateway directory:

```bash
cd ecommerce-go-api-gateway
make dev-up
```

This starts all 8 services in Docker with hot reload enabled.

#### Step 2: Edit Any Service

```bash
# Edit User Service
vim /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-user-service/main.go

# Save file (Cmd+S)
# ✨ Air detects change and auto-reloads in 1-2 seconds!
```

#### Step 3: View Logs

```bash
# All services
make dev-logs

# Specific service
docker-compose -f docker-compose.dev.yaml logs -f user-service
```

#### Step 4: Stop Everything

```bash
make dev-down
```

---

## Comparison

### Manual Terminals (Current Way)

```
✅ Pros:
- Simple and straightforward
- No Docker overhead
- Easy debugging
- See all logs in terminals

❌ Cons:
- 8 terminals to manage
- Manual restart after code changes (2-3 sec)
- Need to remember which port each service uses
```

### Docker Hot Reload (New Way)

```
✅ Pros:
- ONE command to start all services
- Auto-reload on code changes (1-2 sec)
- No manual restarts
- Easy to manage

❌ Cons:
- Requires Docker
- Slightly more resource usage
- Need to check logs via docker commands
```

---

## Using Hot Reload with Docker

### Daily Workflow

**Morning**:
```bash
cd ecommerce-go-api-gateway
make dev-up
```

**During Development**:
```bash
# Edit any service code
vim /path/to/service/handler.go

# Save file
# ✨ Automatically reloads!

# Check logs
make dev-logs
```

**Evening**:
```bash
make dev-down
```

---

### Example: Edit User Service

**Terminal 1** - Start services:
```bash
cd ecommerce-go-api-gateway
make dev-up
```

**Terminal 2** - Edit code:
```bash
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-user-service

# Edit handler
vim api/handler.go

# Save file
# Watch logs in Terminal 3
```

**Terminal 3** - Watch logs:
```bash
docker-compose -f docker-compose.dev.yaml logs -f user-service
```

You'll see:
```
user-service | watching .
user-service | main.go has changed
user-service | building...
user-service | running...
user-service | Server started on :8080
```

---

## File Structure (Per Service)

Each service now has:

```
ecommerce-go-user-service/
├── Dockerfile              # Production (optimized)
├── Dockerfile.dev          # Development (with Air) ✨ NEW
├── .air.toml              # Air configuration ✨ NEW
├── main.go
├── go.mod
└── ...
```

---

## How Air Works

### .air.toml Configuration

```toml
[build]
  cmd = "go build -o ./tmp/main ./cmd/api"  # Build command
  bin = "./tmp/main"                        # Binary location
  delay = 1000                              # Wait 1 sec before rebuild
  include_ext = ["go", "yaml", "yml"]       # Watch these files
  exclude_dir = ["tmp", "vendor"]           # Ignore these
```

### Dockerfile.dev

```dockerfile
FROM golang:1.23-alpine

# Install Air
RUN go install github.com/air-verse/air@latest

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Source code mounted as volume (for hot reload)

# Run with Air
CMD ["air", "-c", ".air.toml"]
```

---

## Available Commands

### Development Mode (Docker Hot Reload)

```bash
make dev-up          # Start all services with hot reload
make dev-down        # Stop all services
make dev-logs        # View all logs
make dev-rebuild     # Rebuild containers (if Dockerfile.dev changes)
make dev-clean       # Clean up everything
make dev-health      # Check service status
```

### Production Mode (Docker)

```bash
make docker-build    # Build production images
make docker-up       # Start production services
make docker-down     # Stop services
make docker-logs     # View logs
```

### Local Manual Mode (What You Use Now)

```bash
# No special commands
# Just: go run main.go in each terminal
```

---

## When to Use Each Mode

### Use Manual Terminals When:
- Quick testing
- Debugging specific service
- Learning/experimenting
- Don't want Docker overhead

### Use Docker Hot Reload When:
- Daily development
- Working on multiple services
- Want auto-reload
- Need consistent environment

### Use Production Docker When:
- Testing deployment
- CI/CD pipeline
- Final integration tests
- Before production release

---

## Testing Hot Reload

### Test 1: Start Services

```bash
cd ecommerce-go-api-gateway
make dev-up
```

Wait for all services to start (check logs):
```bash
make dev-logs
```

### Test 2: Make a Change

```bash
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-user-service

# Edit a handler
echo '// Test comment' >> api/handler.go
```

### Test 3: Watch Auto-Reload

```bash
docker-compose -f docker-compose.dev.yaml logs -f user-service
```

You should see:
```
api/handler.go has changed
building...
running...
Server started
```

### Test 4: Test Endpoint

```bash
curl http://localhost:8080/api/v1/users/1
```

---

## Troubleshooting

### Issue 1: Air Not Reloading

**Check logs**:
```bash
docker-compose -f docker-compose.dev.yaml logs user-service
```

**Solution**: Make sure you're editing files in the mounted directory.

### Issue 2: Changes Not Visible

**Check if file is being watched**:
```bash
docker exec user-service-dev sh -c "ls -la /app"
```

**Solution**: Verify volume mount in docker-compose.dev.yaml.

### Issue 3: Build Errors

**View detailed logs**:
```bash
docker-compose -f docker-compose.dev.yaml logs user-service --tail=50
```

**Solution**: Fix syntax errors, Air will retry automatically.

---

## Configuration Files

### Dockerfile.dev (All Services)

Located in each service directory:
- `/ecommerce-go-user-service/Dockerfile.dev`
- `/ecommerce-go-product-service/Dockerfile.dev`
- ... and so on

### .air.toml (All Services)

Located in each service directory:
- `/ecommerce-go-user-service/.air.toml`
- `/ecommerce-go-product-service/.air.toml`
- ... and so on

### docker-compose.dev.yaml

Located in API Gateway:
- `/ecommerce-go-api-gateway/docker-compose.dev.yaml`

---

## Summary

You now have **3 development options**:

### Option 1: Manual Terminals (Current)
```bash
# 8 terminals
# go run main.go in each
# Manual restart after changes
```

### Option 2: Docker Hot Reload (New!)
```bash
make dev-up
# Edit code → Auto-reload ✨
make dev-down
```

### Option 3: Production Docker
```bash
make docker-build
make docker-up
# For final testing
```

---

## Recommendation

- **Daily Development**: Use Option 1 (Manual) or Option 2 (Hot Reload)
- **Before Deployment**: Use Option 3 (Production Docker)

**All files are ready!** You can start using hot reload whenever you want with:
```bash
make dev-up
```

Or continue using manual terminals - both work perfectly! 🚀
