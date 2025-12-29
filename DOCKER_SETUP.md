# Docker Setup and Testing Guide

Complete guide for running and testing all microservices locally using Docker.

## Table of Contents
- [Overview](#overview)
- [Architecture](#architecture)
- [Prerequisites](#prerequisites)
- [Quick Start](#quick-start)
- [Step-by-Step Setup](#step-by-step-setup)
- [Testing Services](#testing-services)
- [Troubleshooting](#troubleshooting)
- [Docker Commands Reference](#docker-commands-reference)

---

## Overview

This e-commerce platform consists of **8 Docker containers**:
- 1 API Gateway (entry point)
- 7 Backend Microservices

All services are built with **Go** and the **Gin framework**, communicating via HTTP REST APIs.

---

## Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                     Client Applications                      │
│                  (Web, Mobile, Third-party)                  │
└───────────────────────────┬─────────────────────────────────┘
                            │
                            ▼
                  ┌─────────────────┐
                  │  API Gateway    │  Port 8080
                  │  (Entry Point)  │
                  └────────┬────────┘
                           │
        ┌──────────────────┼──────────────────┬──────────────┐
        │                  │                  │              │
        ▼                  ▼                  ▼              ▼
┌──────────────┐   ┌──────────────┐   ┌──────────────┐   ┌──────────────┐
│ User Service │   │Product Service│  │ Order Service│   │Payment Service│
│   Port 8081  │   │   Port 8082   │  │  Port 8083   │   │  Port 8084   │
└──────────────┘   └──────────────┘   └──────────────┘   └──────────────┘

        ▼                  ▼                  ▼
┌──────────────┐   ┌──────────────┐   ┌──────────────┐
│ Inventory    │   │Notification  │   │    Config    │
│   Service    │   │   Service    │   │   Service    │
│  Port 8085   │   │  Port 8086   │   │  Port 8087   │
└──────────────┘   └──────────────┘   └──────────────┘

All services connected via: ecommerce-network (Docker Bridge Network)
```

---

## Prerequisites

Before starting, ensure you have:

- **Docker** (v20.10 or higher)
- **Docker Compose** (v2.0 or higher)
- **Make** (optional, for convenience commands)
- **curl** (for testing endpoints)

**Verify installations**:
```bash
docker --version
docker-compose --version
make --version
curl --version
```

---

## Quick Start

```bash
# Navigate to API Gateway directory
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-api-gateway

# Build all Docker images
make docker-build

# Start all services
make docker-up

# Check status
make docker-health

# View logs
make docker-logs
```

**Access the API Gateway**: http://localhost:8080

---

## Step-by-Step Setup

### Step 1: Navigate to Project Directory

```bash
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-api-gateway
```

---

### Step 2: Clean Previous Setup (Optional)

If you've run the services before, clean up first:

```bash
# Using Makefile
make docker-clean

# Or manually
docker-compose down -v --remove-orphans
docker system prune -f
```

This removes:
- All containers
- All volumes
- Unused networks
- Dangling images

---

### Step 3: Build Docker Images

Build all 8 Docker images:

```bash
# Using Makefile (Recommended)
make docker-build

# Or manually
docker-compose build --no-cache
```

**Build time**: Approximately 5-10 minutes depending on your machine.

**Verify images were created**:
```bash
docker images | grep ecommerce
```

Expected output:
```
ecommerce-go-api-gateway-api-gateway              latest    ...
ecommerce-go-api-gateway-user-service            latest    ...
ecommerce-go-api-gateway-product-service         latest    ...
ecommerce-go-api-gateway-order-service           latest    ...
ecommerce-go-api-gateway-payment-service         latest    ...
ecommerce-go-api-gateway-inventory-service       latest    ...
ecommerce-go-api-gateway-notification-service    latest    ...
ecommerce-go-api-gateway-config-service          latest    ...
```

---

### Step 4: Start All Services

```bash
# Start in detached mode (background)
make docker-up

# Or start with logs visible (foreground)
make docker-up-logs

# Or manually
docker-compose up -d
```

**Startup sequence** (with health checks):
1. Backend services start first (8081-8087)
2. Each service waits for health check to pass
3. API Gateway starts last (depends on all services being healthy)

**Wait time**: 1-2 minutes for all services to be healthy.

---

### Step 5: Verify All Containers Are Running

```bash
# Check container status
make docker-health

# Or manually
docker-compose ps
```

**Expected output**:
```
NAME                    IMAGE              STATUS
api-gateway            ...                Up (healthy)
config-service         ...                Up (healthy)
inventory-service      ...                Up (healthy)
notification-service   ...                Up (healthy)
order-service          ...                Up (healthy)
payment-service        ...                Up (healthy)
product-service        ...                Up (healthy)
user-service           ...                Up (healthy)
```

All containers should show:
- **STATUS**: `Up (healthy)`
- **PORTS**: Correctly mapped

---

### Step 6: View Logs

**View all logs**:
```bash
make docker-logs
```

**View specific service logs**:
```bash
# API Gateway
make docker-logs-gateway

# User Service
make docker-logs-user

# Product Service
make docker-logs-product

# Other services
docker-compose logs -f order-service
docker-compose logs -f payment-service
docker-compose logs -f inventory-service
docker-compose logs -f notification-service
docker-compose logs -f config-service
```

**View last 100 lines**:
```bash
docker-compose logs --tail=100 api-gateway
```

**Follow logs in real-time**:
```bash
docker-compose logs -f api-gateway
```

---

## Testing Services

### Test 1: Health Check All Services

**Quick Test Script**:
```bash
#!/bin/bash
echo "Testing all services..."
echo "======================="

services=(
  "8080:API Gateway"
  "8081:User Service"
  "8082:Product Service"
  "8083:Order Service"
  "8084:Payment Service"
  "8085:Inventory Service"
  "8086:Notification Service"
  "8087:Config Service"
)

for service in "${services[@]}"; do
  port="${service%%:*}"
  name="${service##*:}"

  echo -n "Testing $name (port $port)... "
  response=$(curl -s -o /dev/null -w "%{http_code}" http://localhost:$port/health)

  if [ "$response" == "200" ]; then
    echo "✅ OK"
  else
    echo "❌ FAILED (HTTP $response)"
  fi
done

echo "======================="
```

**Save and run**:
```bash
# Save the script
cat > test-services.sh << 'EOF'
[paste script above]
EOF

# Make executable
chmod +x test-services.sh

# Run it
./test-services.sh
```

---

### Test 2: Individual Health Endpoints

```bash
# API Gateway
curl http://localhost:8080/health
echo ""

# User Service
curl http://localhost:8081/health
echo ""

# Product Service
curl http://localhost:8082/health
echo ""

# Order Service
curl http://localhost:8083/health
echo ""

# Payment Service
curl http://localhost:8084/health
echo ""

# Inventory Service
curl http://localhost:8085/health
echo ""

# Notification Service
curl http://localhost:8086/health
echo ""

# Config Service
curl http://localhost:8087/health
echo ""
```

**Expected Response**: HTTP 200 with JSON like:
```json
{"status":"ok"}
```

---

### Test 3: API Gateway Routing

Test that the gateway correctly routes to backend services:

**User Service Endpoints**:
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

# Get user by ID
curl http://localhost:8080/api/v1/users/1
```

**Product Service Endpoints**:
```bash
# Get all products
curl http://localhost:8080/api/v1/products

# Get product by ID
curl http://localhost:8080/api/v1/products/1

# Create product
curl -X POST http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test Product",
    "description": "A test product",
    "price": 99.99,
    "stock": 100
  }'
```

**Order Service Endpoints**:
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

# Get order by ID
curl http://localhost:8080/api/v1/orders/1
```

**Payment Service Endpoints**:
```bash
# Process payment
curl -X POST http://localhost:8080/api/v1/payments \
  -H "Content-Type: application/json" \
  -d '{
    "order_id": 1,
    "amount": 199.98,
    "payment_method": "credit_card"
  }'
```

**Inventory Service Endpoints**:
```bash
# Update stock
curl -X PUT http://localhost:8080/api/v1/inventory/stock \
  -H "Content-Type: application/json" \
  -d '{
    "product_id": 1,
    "quantity": 50
  }'
```

**Notification Service Endpoints**:
```bash
# Send notification
curl -X POST http://localhost:8080/api/v1/notifications \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "type": "email",
    "message": "Your order has been confirmed"
  }'
```

---

### Test 4: Network Connectivity

**Check Docker network**:
```bash
# List networks
docker network ls

# Inspect the ecommerce network
docker network inspect ecommerce-go-api-gateway_ecommerce-network
```

**Test inter-service communication**:
```bash
# From API Gateway to User Service
docker exec api-gateway wget -O- http://user-service:8080/health

# From API Gateway to Product Service
docker exec api-gateway wget -O- http://product-service:8080/health
```

---

### Test 5: Resource Monitoring

**Monitor CPU and Memory usage**:
```bash
# All containers
docker stats

# Specific containers
docker stats api-gateway user-service product-service
```

**Check container details**:
```bash
# Inspect a container
docker inspect api-gateway

# Check container logs
docker logs api-gateway

# Execute command inside container
docker exec -it api-gateway /bin/sh
```

---

## Troubleshooting

### Issue 1: Container Won't Start

**Symptoms**: Container keeps restarting or exits immediately.

**Diagnosis**:
```bash
# Check logs
docker-compose logs <service-name>

# Example
docker-compose logs user-service

# Check exit code
docker-compose ps
```

**Common Causes**:
- Missing `/health` endpoint
- Build errors
- Port conflicts
- Missing environment variables

**Solution**:
```bash
# Rebuild the service
docker-compose build --no-cache <service-name>

# Restart the service
docker-compose restart <service-name>
```

---

### Issue 2: Health Check Failing

**Symptoms**: Container status shows "unhealthy".

**Diagnosis**:
```bash
# Check if service is listening on port 8080
docker exec user-service netstat -tulpn | grep 8080

# Test health endpoint from inside container
docker exec user-service wget -O- http://localhost:8080/health

# Check health check logs
docker inspect user-service | grep -A 20 Health
```

**Solution**:
Ensure your service:
1. Has a `/health` endpoint
2. Returns HTTP 200 status
3. Listens on port specified by `SERVER_PORT` env var

---

### Issue 3: Port Already in Use

**Symptoms**: Error: "port is already allocated"

**Diagnosis**:
```bash
# Check what's using the ports (8080-8087)
lsof -i :8080
lsof -i :8081
# ... and so on

# Or check all at once
lsof -i :8080-8087
```

**Solution**:
```bash
# Option 1: Kill the process
kill -9 <PID>

# Option 2: Change port in docker-compose.yaml
# Edit the ports mapping, e.g., "9080:8080"
```

---

### Issue 4: Services Can't Communicate

**Symptoms**: Gateway returns "connection refused" errors.

**Diagnosis**:
```bash
# Check if services are on same network
docker network inspect ecommerce-go-api-gateway_ecommerce-network

# Test from gateway to service
docker exec api-gateway wget -O- http://user-service:8080/health
```

**Solution**:
```bash
# Recreate network
docker-compose down
docker network prune
docker-compose up -d
```

---

### Issue 5: "No Space Left on Device"

**Symptoms**: Build fails with disk space error.

**Diagnosis**:
```bash
# Check Docker disk usage
docker system df
```

**Solution**:
```bash
# Clean up unused resources
docker system prune -a

# Remove unused volumes
docker volume prune

# Remove unused images
docker image prune -a
```

---

### Issue 6: Slow Build Times

**Symptoms**: Building images takes very long.

**Solution**:
```bash
# Use BuildKit for faster builds
DOCKER_BUILDKIT=1 docker-compose build

# Or add to ~/.docker/config.json
{
  "features": {
    "buildkit": true
  }
}
```

---

### Issue 7: Container Logs Not Showing

**Diagnosis**:
```bash
# Check if container is running
docker-compose ps

# Check container ID
docker ps -a | grep user-service
```

**Solution**:
```bash
# View logs with timestamps
docker-compose logs -f --timestamps user-service

# Increase log detail
docker-compose logs --tail=500 user-service
```

---

## Docker Commands Reference

### Container Management

```bash
# Start all services
make docker-up
# Or: docker-compose up -d

# Start with logs
make docker-up-logs
# Or: docker-compose up

# Stop all services
make docker-down
# Or: docker-compose down

# Restart all services
make docker-restart
# Or: docker-compose restart

# Restart specific service
docker-compose restart user-service

# Stop specific service
docker-compose stop user-service

# Start specific service
docker-compose start user-service

# Remove all containers
docker-compose down -v
```

---

### Image Management

```bash
# Build all images
make docker-build
# Or: docker-compose build

# Build specific service
docker-compose build user-service

# Build without cache
docker-compose build --no-cache

# Rebuild everything
make docker-rebuild
# Or: docker-compose down && docker-compose build --no-cache && docker-compose up -d

# Remove all images
docker-compose down --rmi all

# List images
docker images | grep ecommerce
```

---

### Logs and Monitoring

```bash
# View all logs
make docker-logs
# Or: docker-compose logs -f

# View specific service logs
make docker-logs-gateway
make docker-logs-user
make docker-logs-product
# Or: docker-compose logs -f <service-name>

# View last N lines
docker-compose logs --tail=100 user-service

# View logs with timestamps
docker-compose logs -f --timestamps

# Monitor resources
docker stats

# Check container status
make docker-health
# Or: docker-compose ps
```

---

### Network Management

```bash
# List networks
docker network ls

# Inspect network
docker network inspect ecommerce-go-api-gateway_ecommerce-network

# Create network
docker network create ecommerce-network

# Remove unused networks
docker network prune
```

---

### Volume Management

```bash
# List volumes
docker volume ls

# Inspect volume
docker volume inspect <volume-name>

# Remove all volumes
docker volume prune

# Remove specific volume
docker volume rm <volume-name>
```

---

### Cleanup Commands

```bash
# Clean everything
make docker-clean
# Or: docker-compose down -v --remove-orphans && docker system prune -f

# Remove stopped containers
docker container prune

# Remove unused images
docker image prune -a

# Remove unused volumes
docker volume prune

# Remove unused networks
docker network prune

# Remove everything (CAUTION!)
docker system prune -a --volumes
```

---

### Debugging Commands

```bash
# Execute command in running container
docker exec -it api-gateway /bin/sh

# View container processes
docker top user-service

# Inspect container configuration
docker inspect user-service

# View container resource usage
docker stats user-service

# Copy files from container
docker cp api-gateway:/root/config/config.yaml ./config-backup.yaml

# Copy files to container
docker cp ./test.txt api-gateway:/root/
```

---

## Service Ports Reference

| Service | Container Port | Host Port | Health Check URL |
|---------|---------------|-----------|------------------|
| API Gateway | 8080 | 8080 | http://localhost:8080/health |
| User Service | 8080 | 8081 | http://localhost:8081/health |
| Product Service | 8080 | 8082 | http://localhost:8082/health |
| Order Service | 8080 | 8083 | http://localhost:8083/health |
| Payment Service | 8080 | 8084 | http://localhost:8084/health |
| Inventory Service | 8080 | 8085 | http://localhost:8085/health |
| Notification Service | 8080 | 8086 | http://localhost:8086/health |
| Config Service | 8080 | 8087 | http://localhost:8087/health |

---

## API Gateway Routes

All routes are prefixed with `/api/v1`

### User Service Routes
- `POST /api/v1/users/register` - Register new user
- `POST /api/v1/users/login` - User login
- `GET /api/v1/users/:id` - Get user by ID

### Product Service Routes
- `GET /api/v1/products` - List all products
- `GET /api/v1/products/:id` - Get product by ID
- `POST /api/v1/products` - Create new product

### Order Service Routes
- `POST /api/v1/orders` - Create new order
- `GET /api/v1/orders/:id` - Get order by ID

### Payment Service Routes
- `POST /api/v1/payments` - Process payment

### Inventory Service Routes
- `PUT /api/v1/inventory/stock` - Update stock levels

### Notification Service Routes
- `POST /api/v1/notifications` - Send notification

---

## Environment Variables

Configure services via environment variables in `docker-compose.yaml`:

```yaml
environment:
  - SERVER_PORT=:8080          # Port the service listens on
  - SERVER_MODE=debug          # debug or release
  - GIN_MODE=debug             # Gin framework mode
  - LOGGER_LEVEL=info          # Log level: debug, info, warn, error
```

For the API Gateway:
```yaml
environment:
  - SERVICES_USER_SERVICE=http://user-service:8080
  - SERVICES_PRODUCT_SERVICE=http://product-service:8080
  # ... and so on
```

---

## Best Practices

### Development Workflow

1. **Make changes to code**
2. **Rebuild specific service**:
   ```bash
   docker-compose build user-service
   ```
3. **Restart the service**:
   ```bash
   docker-compose restart user-service
   ```
4. **Check logs**:
   ```bash
   docker-compose logs -f user-service
   ```

### Production Considerations

- Change `GIN_MODE` to `release`
- Change `SERVER_MODE` to `release`
- Set `LOGGER_LEVEL` to `warn` or `error`
- Use proper secrets management
- Add database services
- Implement rate limiting
- Add monitoring (Prometheus, Grafana)
- Use reverse proxy (Nginx)

---

## Additional Resources

- [Docker Documentation](https://docs.docker.com/)
- [Docker Compose Documentation](https://docs.docker.com/compose/)
- [Gin Framework Documentation](https://gin-gonic.com/docs/)
- [Go Documentation](https://golang.org/doc/)

---

## Support

If you encounter issues:

1. Check the [Troubleshooting](#troubleshooting) section
2. View service logs: `docker-compose logs -f <service-name>`
3. Check container status: `docker-compose ps`
4. Inspect container: `docker inspect <container-name>`

---

**Last Updated**: December 29, 2025
**Version**: 1.0.0
