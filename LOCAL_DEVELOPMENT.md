# Local Development Guide (No Docker)

Simple guide to run all 8 microservices locally without Docker - just like Express Gateway!

## Quick Start

```bash
# Start ALL services with ONE command
make local-start

# Access through API Gateway
curl http://localhost:8080/health

# Stop all services
make local-stop
```

That's it! 🎉

---

## How It Works

- **All 8 services** run in the background on your machine
- **No Docker** needed - just Go
- **API Gateway** on port 8080 routes to all services
- **Logs** saved to `/Users/yudizsolutionsltd/Documents/Project/GolangEcom/logs/`
- **PID files** saved to track running processes

---

## Daily Workflow

### Morning: Start Everything

```bash
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-api-gateway

# Start all 8 services
make local-start
```

**Output**:
```
Starting all microservices locally...

Starting config-service on port 8087...
✓ config-service started (PID: 12345)

Starting user-service on port 8081...
✓ user-service started (PID: 12346)

... (all services start)

✅ All services started successfully!

📊 Service Status:
  - API Gateway: http://localhost:8080 ⭐

🌐 Access everything through: http://localhost:8080
```

### During Development: Make Changes

```bash
# 1. Edit any service code
vim /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-user-service/main.go

# 2. Restart that service only
make local-restart

# OR restart everything
make local-restart
```

### Check What's Running

```bash
make local-status
```

**Output**:
```
════════════════════════════════════════
  Microservices Status
════════════════════════════════════════

Config Service:       ✓ Running (PID: 12345, Port: 8087)
                        Health check: OK
User Service:         ✓ Running (PID: 12346, Port: 8081)
                        Health check: OK
... (all services)

✅ All services are running! (8/8)
```

### View Logs

```bash
# View ALL logs in real-time
make local-logs

# View specific service logs
make local-logs-service SERVICE=user-service

# Or directly
tail -f /Users/yudizsolutionsltd/Documents/Project/GolangEcom/logs/user-service.log
```

### Evening: Stop Everything

```bash
make local-stop
```

**Output**:
```
Stopping all microservices...

Stopping api-gateway (PID: 12352)...
✓ api-gateway stopped

... (all services stop)

✅ All services stopped!
```

---

## Available Commands

| Command | Description |
|---------|-------------|
| `make local-start` | Start all 8 services in background |
| `make local-stop` | Stop all services |
| `make local-status` | Check which services are running |
| `make local-logs` | View all service logs |
| `make local-restart` | Restart all services |
| `make local-clean` | Clean log files and PIDs |

---

## Service Ports

| Service | Port | Direct URL | Through Gateway |
|---------|------|------------|----------------|
| **API Gateway** | 8080 | http://localhost:8080 | - |
| User Service | 8081 | http://localhost:8081 | http://localhost:8080/api/v1/users/* |
| Product Service | 8082 | http://localhost:8082 | http://localhost:8080/api/v1/products/* |
| Order Service | 8083 | http://localhost:8083 | http://localhost:8080/api/v1/orders/* |
| Payment Service | 8084 | http://localhost:8084 | http://localhost:8080/api/v1/payments/* |
| Inventory Service | 8085 | http://localhost:8085 | http://localhost:8080/api/v1/inventory/* |
| Notification Service | 8086 | http://localhost:8086 | http://localhost:8080/api/v1/notifications/* |
| Config Service | 8087 | http://localhost:8087 | - |

---

## Testing Through API Gateway

### Health Check

```bash
# Gateway health
curl http://localhost:8080/health

# Individual service (through gateway)
curl http://localhost:8081/health
```

### User Service

```bash
# Register user
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123"
  }'

# Login
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

### Product Service

```bash
# Get all products
curl http://localhost:8080/api/v1/products

# Create product
curl -X POST http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test Product",
    "price": 99.99,
    "stock": 100
  }'
```

### Order Service

```bash
# Create order
curl -X POST http://localhost:8080/api/v1/orders \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "items": [
      {"product_id": 1, "quantity": 2}
    ]
  }'
```

---

## Troubleshooting

### Issue 1: Port Already in Use

**Error**: `address already in use`

**Check**:
```bash
lsof -i :8080
```

**Solution**:
```bash
# Stop all services
make local-stop

# Or kill specific port
kill $(lsof -ti:8080)
```

### Issue 2: Service Won't Start

**Check logs**:
```bash
tail -f /Users/yudizsolutionsltd/Documents/Project/GolangEcom/logs/user-service.log
```

**Common issues**:
- Missing dependencies: `go mod download`
- Build errors: Check syntax
- Port conflict: Use `lsof -i :8081`

### Issue 3: Can't See Changes

Services run in background, so:

```bash
# Restart to see changes
make local-restart
```

### Issue 4: Services Running But Can't Access

**Check status**:
```bash
make local-status
```

**Check health**:
```bash
curl http://localhost:8080/health
curl http://localhost:8081/health
```

---

## Comparison with Docker

| Feature | Local (This Setup) | Docker |
|---------|-------------------|--------|
| **Startup** | 5-10 seconds ⚡ | 2-5 minutes |
| **Code changes** | Restart (2 sec) | Rebuild (1-3 min) |
| **Resource usage** | Low | High |
| **Debugging** | Easy | Complex |
| **Production match** | 90% | 100% |
| **Network** | localhost | Docker network |

**Recommendation**: Use local for development, Docker for final testing before deployment.

---

## Advanced Usage

### Run Specific Services Only

Edit `scripts/start-all-local.sh` and comment out services you don't need:

```bash
# start_service "notification-service" "$BASE_DIR/ecommerce-go-notification-service" 8086
```

### Change Ports

Edit the script to use different ports:

```bash
start_service "user-service" "$BASE_DIR/ecommerce-go-user-service" 9081
```

Then update `config/config.yaml`:

```yaml
services:
  user_service: "http://localhost:9081"
```

### Run in Terminals (Instead of Background)

If you prefer seeing logs in terminals:

```bash
# Terminal 1
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-user-service
SERVER_PORT=:8081 go run main.go

# Terminal 2
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-product-service
SERVER_PORT=:8082 go run main.go

# ... and so on
```

---

## Log Files Location

All logs are saved to:
```
/Users/yudizsolutionsltd/Documents/Project/GolangEcom/logs/
├── api-gateway.log
├── config-service.log
├── user-service.log
├── product-service.log
├── order-service.log
├── payment-service.log
├── inventory-service.log
└── notification-service.log
```

PID files are saved to:
```
/Users/yudizsolutionsltd/Documents/Project/GolangEcom/pids/
├── api-gateway.pid
├── user-service.pid
└── ...
```

---

## Tips & Tricks

### 1. Monitor All Services

```bash
# Watch status continuously
watch -n 2 make local-status

# Or
while true; do clear; make local-status; sleep 2; done
```

### 2. Quick Health Check All Services

```bash
for port in {8080..8087}; do
  echo -n "Port $port: "
  curl -s http://localhost:$port/health && echo " ✓" || echo " ✗"
done
```

### 3. Auto-Restart on Code Changes

Install `entr`:
```bash
brew install entr
```

Watch and auto-restart:
```bash
find . -name '*.go' | entr -r make local-restart
```

### 4. Clean Start

```bash
# Stop, clean, and start fresh
make local-stop
make local-clean
make local-start
```

---

## FAQ

**Q: Do I need to rebuild after code changes?**
A: No, just restart: `make local-restart`

**Q: Can I run some services in Docker and some locally?**
A: Yes, but you'll need to update port configurations.

**Q: How do I debug a service?**
A: Run it directly in a terminal or use Delve debugger:
```bash
cd /path/to/service
dlv debug main.go
```

**Q: Can I run this on Windows?**
A: The scripts are for Mac/Linux. For Windows, use Docker or WSL.

**Q: What if a service crashes?**
A: Check logs: `tail -f /Users/yudizsolutionsltd/Documents/Project/GolangEcom/logs/service-name.log`

---

## Summary

This local development setup gives you:

✅ **One command** to start everything: `make local-start`
✅ **Fast feedback** - restart in 2 seconds
✅ **Easy debugging** - native Go processes
✅ **Low resource usage** - no Docker overhead
✅ **Simple management** - clear status, logs, and controls

Perfect for daily development! 🚀
