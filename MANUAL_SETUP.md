# Manual Setup Guide - Run in Separate Terminals

Simple guide to run all 8 services manually in separate terminals - **NO scripts needed!**

---

## Quick Reference

Open **8 terminals** and run these commands (one in each terminal):

```bash
# Terminal 1 - Config Service (Port 8087)
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-config-service
SERVER_PORT=:8087 go run main.go

# Terminal 2 - User Service (Port 8081)
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-user-service
SERVER_PORT=:8081 go run main.go

# Terminal 3 - Product Service (Port 8082)
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-product-service
SERVER_PORT=:8082 go run main.go

# Terminal 4 - Order Service (Port 8083)
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-order-service
SERVER_PORT=:8083 go run main.go

# Terminal 5 - Payment Service (Port 8084)
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-payment-service
SERVER_PORT=:8084 go run main.go

# Terminal 6 - Inventory Service (Port 8085)
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-inventory-service
SERVER_PORT=:8085 go run main.go

# Terminal 7 - Notification Service (Port 8086)
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-notification-service
SERVER_PORT=:8086 go run main.go

# Terminal 8 - API Gateway (Port 8080) ⭐
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-api-gateway
go run cmd/api/main.go
```

**Then access everything through**: http://localhost:8080

---

## Step-by-Step Guide

### Terminal 1: Config Service

```bash
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-config-service
SERVER_PORT=:8087 go run main.go
```

**You should see**:
```
Server started on port :8087
```

**Keep this terminal open!**

---

### Terminal 2: User Service

```bash
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-user-service
SERVER_PORT=:8081 go run main.go
```

**You should see**:
```
Server started on port :8081
```

**Keep this terminal open!**

---

### Terminal 3: Product Service

```bash
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-product-service
SERVER_PORT=:8082 go run main.go
```

**You should see**:
```
Server started on port :8082
```

**Keep this terminal open!**

---

### Terminal 4: Order Service

```bash
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-order-service
SERVER_PORT=:8083 go run main.go
```

**You should see**:
```
Server started on port :8083
```

**Keep this terminal open!**

---

### Terminal 5: Payment Service

```bash
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-payment-service
SERVER_PORT=:8084 go run main.go
```

**You should see**:
```
Server started on port :8084
```

**Keep this terminal open!**

---

### Terminal 6: Inventory Service

```bash
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-inventory-service
SERVER_PORT=:8085 go run main.go
```

**You should see**:
```
Server started on port :8085
```

**Keep this terminal open!**

---

### Terminal 7: Notification Service

```bash
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-notification-service
SERVER_PORT=:8086 go run main.go
```

**You should see**:
```
Server started on port :8086
```

**Keep this terminal open!**

---

### Terminal 8: API Gateway (Main Entry Point) ⭐

```bash
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-api-gateway
go run main.go
```

**You should see**:
```
Server started on port :8080
```

**This is your main entry point!**

---

## Testing Everything

Open a **new terminal** (Terminal 9) and test:

```bash
# Test API Gateway
curl http://localhost:8080/health

# Test User Service through Gateway
curl http://localhost:8080/api/v1/users/1

# Test Product Service through Gateway
curl http://localhost:8080/api/v1/products

# Test direct service access
curl http://localhost:8081/health  # User Service direct
curl http://localhost:8082/health  # Product Service direct
```

---

## Terminal Layout Suggestion

Organize your terminals like this:

```
┌─────────────────┬─────────────────┬─────────────────┐
│  Terminal 1     │  Terminal 2     │  Terminal 3     │
│  Config:8087    │  User:8081      │  Product:8082   │
├─────────────────┼─────────────────┼─────────────────┤
│  Terminal 4     │  Terminal 5     │  Terminal 6     │
│  Order:8083     │  Payment:8084   │  Inventory:8085 │
├─────────────────┼─────────────────┼─────────────────┤
│  Terminal 7     │  Terminal 8     │  Terminal 9     │
│  Notify:8086    │  Gateway:8080⭐ │  Testing        │
└─────────────────┴─────────────────┴─────────────────┘
```

---

## Making Code Changes

### To Change a Specific Service:

1. Go to that service's terminal
2. Press **Ctrl+C** to stop
3. Edit the code
4. Run `SERVER_PORT=:8081 go run main.go` again

**Example - Updating User Service**:

```bash
# In Terminal 2 (User Service)
# Press Ctrl+C

# Edit code
vim main.go

# Restart
SERVER_PORT=:8081 go run main.go
```

**Other services keep running!**

---

## Service Port Reference

| Terminal | Service | Port | Command |
|----------|---------|------|---------|
| 1 | Config Service | 8087 | `SERVER_PORT=:8087 go run main.go` |
| 2 | User Service | 8081 | `SERVER_PORT=:8081 go run main.go` |
| 3 | Product Service | 8082 | `SERVER_PORT=:8082 go run main.go` |
| 4 | Order Service | 8083 | `SERVER_PORT=:8083 go run main.go` |
| 5 | Payment Service | 8084 | `SERVER_PORT=:8084 go run main.go` |
| 6 | Inventory Service | 8085 | `SERVER_PORT=:8085 go run main.go` |
| 7 | Notification Service | 8086 | `SERVER_PORT=:8086 go run main.go` |
| 8 | **API Gateway** ⭐ | 8080 | `go run cmd/api/main.go` |

---

## Troubleshooting

### Error: "address already in use"

**Problem**: Port is already taken

**Solution**:
```bash
# Find what's using the port (e.g., 8081)
lsof -i :8081

# Kill it
kill -9 <PID>

# Or restart that service's terminal
```

### Error: "connect: connection refused"

**Problem**: Backend service not running

**Solution**: Check all 7 backend services are running (Terminals 1-7)

### Changes not visible

**Problem**: Need to restart service

**Solution**:
- Go to service terminal
- Press Ctrl+C
- Run command again

---

## Tips

### 1. Use Terminal Tabs/Windows

Instead of 8 separate terminal windows, use **tabs**:

**iTerm2/Terminal**:
- Cmd+T to open new tab
- Cmd+Number to switch tabs

**VSCode**:
- Use split terminals
- Ctrl+` to toggle terminal

### 2. Name Your Terminals

**iTerm2**: Right-click tab → Edit Session → Set Title
**VSCode**: Right-click terminal → Rename

Example names:
- "Config:8087"
- "User:8081"
- "Gateway:8080"

### 3. Save Commands

Create a text file with all commands:

```bash
# save-commands.txt
Terminal 1: cd /Users/.../ecommerce-go-config-service && SERVER_PORT=:8087 go run main.go
Terminal 2: cd /Users/.../ecommerce-go-user-service && SERVER_PORT=:8081 go run main.go
...
```

Copy-paste when needed!

### 4. Use `cmd/api/main.go` if needed

Some services might have:
```bash
go run cmd/api/main.go
```

Check your service structure!

---

## Startup Checklist

- [ ] Terminal 1: Config Service running on :8087
- [ ] Terminal 2: User Service running on :8081
- [ ] Terminal 3: Product Service running on :8082
- [ ] Terminal 4: Order Service running on :8083
- [ ] Terminal 5: Payment Service running on :8084
- [ ] Terminal 6: Inventory Service running on :8085
- [ ] Terminal 7: Notification Service running on :8086
- [ ] Terminal 8: API Gateway running on :8080
- [ ] Test: `curl http://localhost:8080/health` works

---

## Stop Everything

To stop all services:

**In each terminal (1-8)**:
- Press **Ctrl+C**

That's it!

---

## Configuration

Your `config/config.yaml` should have:

```yaml
server:
  port: ":8080"
  mode: "debug"

services:
  user_service: "http://localhost:8081"
  product_service: "http://localhost:8082"
  order_service: "http://localhost:8083"
  payment_service: "http://localhost:8084"
  inventory_service: "http://localhost:8085"
  notification_service: "http://localhost:8086"
  config_service: "http://localhost:8087"

logger:
  level: "info"
```

**This is already set up!** ✅

---

## Summary

**Daily Workflow**:

1. Open 8 terminals
2. Run `go run main.go` in each (with SERVER_PORT for services)
3. Develop and edit code
4. Restart specific service when you make changes (Ctrl+C, then run again)
5. Test through http://localhost:8080
6. Press Ctrl+C in all terminals when done

**No scripts, no automation - just simple manual control!** 🎯
