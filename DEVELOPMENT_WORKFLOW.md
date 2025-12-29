# Development Workflow Guide

Complete guide for different development approaches and when to use each one.

## Table of Contents
- [Quick Comparison](#quick-comparison)
- [Option 1: Development Mode with Hot Reload (RECOMMENDED)](#option-1-development-mode-with-hot-reload-recommended)
- [Option 2: Local Development (No Docker)](#option-2-local-development-no-docker)
- [Option 3: Production Docker Mode](#option-3-production-docker-mode)
- [Switching Between Modes](#switching-between-modes)
- [Best Practices](#best-practices)

---

## Quick Comparison

| Mode | Code Changes | Build Time | Use Case |
|------|--------------|------------|----------|
| **Dev Mode (Hot Reload)** | Auto-reload | 0s ⚡ | Daily development |
| **Local (No Docker)** | Manual restart | 1-2s | Quick testing |
| **Production Docker** | Rebuild container | 60-180s | Testing production build |

---

## Option 1: Development Mode with Hot Reload (RECOMMENDED)

### What Is It?

- Uses **Docker volumes** to mount your code
- Uses **Air** for automatic reload when you save files
- **NO rebuild needed** when you change code
- All services run in Docker with networking

### Setup (One-Time)

You need to create `Dockerfile.dev` and `.air.toml` for each service.

#### Step 1: Create Dockerfile.dev in EACH service

Example for User Service:

```bash
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-user-service
```

Create `Dockerfile.dev`:

```dockerfile
# Dockerfile.dev - FOR DEVELOPMENT ONLY
FROM golang:1.23-alpine

# Install Air for hot reload
RUN go install github.com/air-verse/air@latest

# Set Go toolchain
ENV GOTOOLCHAIN=auto

WORKDIR /app

# Copy go.mod first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Source code will be mounted as volume (not copied)
# This allows hot reload when you edit files

# Expose port
EXPOSE 8080

# Run with Air (watches for changes)
CMD ["air", "-c", ".air.toml"]
```

Repeat for ALL 8 services (user, product, order, payment, inventory, notification, config, api-gateway).

#### Step 2: Copy .air.toml to each service

```bash
# Copy from API Gateway to all services
cp /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-api-gateway/.air.toml \
   /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-user-service/

cp /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-api-gateway/.air.toml \
   /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-product-service/

# Repeat for all services...
```

Or I can create a script to do this automatically.

#### Step 3: Start Development Mode

```bash
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-api-gateway

# Start all services in dev mode
make dev-up

# View logs
make dev-logs
```

### Daily Development Workflow

```bash
# 1. Start services (once per day)
make dev-up

# 2. Edit any code in any service
vim /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-user-service/main.go

# 3. Save the file
# ✅ Air automatically detects change and reloads!
# ⚡ Takes 1-2 seconds

# 4. Test your changes immediately
curl http://localhost:8081/health

# 5. When done for the day
make dev-down
```

### Advantages ✅

- **Instant feedback** - Save file, see changes in 1-2 seconds
- **No manual rebuilds** - Air handles everything
- **All services connected** - Full microservices testing
- **Logs visible** - See all output in real-time

### Disadvantages ❌

- **Initial setup** - Need to create Dockerfile.dev for each service
- **Larger containers** - Includes Go toolchain (development only)
- **Volume overhead** - Slight performance penalty on Mac/Windows

---

## Option 2: Local Development (No Docker)

### What Is It?

- Run services **directly** on your machine (no Docker)
- Use `go run` for instant startup
- Each service runs in a separate terminal

### Setup

You need to:
1. Have Go installed locally
2. Configure services to talk to each other via localhost

### Workflow

```bash
# Terminal 1 - User Service
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-user-service
SERVER_PORT=:8081 go run main.go

# Terminal 2 - Product Service
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-product-service
SERVER_PORT=:8082 go run main.go

# Terminal 3 - Order Service
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-order-service
SERVER_PORT=:8083 go run main.go

# ... and so on for all 8 services

# Terminal 8 - API Gateway
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-api-gateway
make run
```

### Making Changes

```bash
# 1. Stop the service (Ctrl+C in terminal)
# 2. Edit code
# 3. Restart: go run main.go
```

### Advantages ✅

- **Fastest startup** - 1-2 seconds
- **Easy debugging** - Use IDE debugger directly
- **No Docker overhead** - Native performance
- **Simple** - No containerization complexity

### Disadvantages ❌

- **8 terminals** - Hard to manage
- **Manual restarts** - Need to restart after every change
- **Environment differences** - Might work locally but fail in Docker
- **Complex configuration** - Need to manage 8 services manually

---

## Option 3: Production Docker Mode

### What Is It?

- Uses **optimized Dockerfiles** (multi-stage builds)
- Creates **small, production-ready images**
- No hot reload - must rebuild to see changes

### When to Use

- Before deploying to production
- Testing Docker builds
- CI/CD pipelines
- Final integration testing

### Workflow

```bash
# Build all images (5-10 minutes first time)
make docker-build

# Start services
make docker-up

# View logs
make docker-logs
```

### Making Changes

```bash
# 1. Edit code
vim /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-user-service/main.go

# 2. Rebuild ONLY the changed service (faster)
docker-compose build user-service

# 3. Restart the service
docker-compose restart user-service

# 4. Check logs
docker-compose logs -f user-service
```

Or rebuild everything:

```bash
# Rebuild all (slower)
make docker-rebuild
```

### Advantages ✅

- **Production-like** - Exact same environment as deployment
- **Optimized images** - Small size (~20MB Alpine)
- **Full testing** - Tests Docker networking, volumes, etc.

### Disadvantages ❌

- **Slow rebuilds** - 1-3 minutes per service
- **No hot reload** - Must manually rebuild
- **Resource intensive** - Uses more CPU/memory

---

## Switching Between Modes

### From Dev Mode to Production Mode

```bash
# Stop dev services
make dev-down

# Start production services
make docker-build
make docker-up
```

### From Production to Dev Mode

```bash
# Stop production services
make docker-down

# Start dev services
make dev-up
```

### From Docker to Local

```bash
# Stop all Docker services
make docker-down
make dev-down

# Run locally (8 terminals)
go run main.go  # In each service directory
```

---

## Best Practices

### Recommended Workflow

1. **Daily Development**: Use **Dev Mode** (Option 1)
   ```bash
   make dev-up
   # Edit code, auto-reload happens
   make dev-down
   ```

2. **Quick Testing**: Use **Local** (Option 2)
   ```bash
   go run main.go
   # Test specific feature
   # Ctrl+C when done
   ```

3. **Before Deployment**: Use **Production** (Option 3)
   ```bash
   make docker-build
   make docker-up
   # Run full integration tests
   make docker-down
   ```

### File Organization

Keep these files in **each service**:

```
ecommerce-go-user-service/
├── Dockerfile          # Production build
├── Dockerfile.dev      # Development build (with Air)
├── .air.toml          # Air configuration
├── .dockerignore      # Ignore tmp/, vendor/, etc.
├── main.go
└── ...
```

### .dockerignore

Create this in each service to speed up builds:

```
# .dockerignore
tmp/
vendor/
.git/
.idea/
*.log
coverage.out
main
```

### Git Configuration

Add to `.gitignore`:

```
# .gitignore
tmp/
main
*.log
```

But **commit** these files:
- `Dockerfile`
- `Dockerfile.dev`
- `.air.toml`
- `docker-compose.yaml`
- `docker-compose.dev.yaml`

---

## Commands Cheat Sheet

### Development Mode (Hot Reload)

```bash
make dev-up          # Start dev environment
make dev-down        # Stop dev environment
make dev-logs        # View logs
make dev-health      # Check status
make dev-rebuild     # Rebuild containers (rarely needed)
make dev-clean       # Remove everything
```

### Production Mode

```bash
make docker-build    # Build images
make docker-up       # Start services
make docker-down     # Stop services
make docker-logs     # View logs
make docker-health   # Check status
make docker-rebuild  # Rebuild everything
make docker-clean    # Remove everything
```

### Local Mode

```bash
make build           # Build binary
make run             # Run locally
make test            # Run tests
make clean           # Clean artifacts
```

---

## Troubleshooting

### Dev Mode Not Reloading

**Check if Air is running**:
```bash
docker-compose -f docker-compose.dev.yaml logs user-service
```

Look for: `watching .`

**Solution**: Rebuild dev containers
```bash
make dev-rebuild
```

### Changes Not Visible

**Check volume mount**:
```bash
docker inspect user-service-dev | grep Mounts -A 20
```

Should show your source directory mounted to `/app`

### Port Conflicts

**Check if services are already running**:
```bash
lsof -i :8080-8087
```

**Stop conflicting services**:
```bash
make docker-down
make dev-down
```

---

## Performance Tips

### Speed Up Dev Mode

1. **Use .dockerignore** - Exclude tmp/, vendor/
2. **Named volumes** - Faster than bind mounts for dependencies
3. **Limit services** - Only run services you're actively developing

### Speed Up Production Builds

1. **Layer caching** - Don't change go.mod often
2. **Multi-stage builds** - Already implemented
3. **BuildKit** - Enable in Docker settings

---

## FAQ

**Q: Can I use hot reload in production?**
A: No, use production Docker mode (optimized images).

**Q: Do I need to rebuild when I add a new dependency?**
A: In dev mode, yes - rebuild once to download new dependency.

**Q: Which mode uses less resources?**
A: Local mode (no Docker overhead).

**Q: Can I debug in dev mode?**
A: Yes, but local mode is easier for debugging.

**Q: How do I update just one service in dev mode?**
A: Just edit and save - Air auto-reloads. No rebuild needed!

---

**Recommendation**: Start with **Dev Mode** (Option 1) for daily work. Switch to **Production Mode** (Option 3) before merging to main branch.
