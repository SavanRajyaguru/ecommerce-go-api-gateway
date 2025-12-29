#!/bin/bash

# Script to start all microservices locally without Docker
# Each service runs in background with logs saved to logs/ directory

set -e

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

BASE_DIR="/Users/yudizsolutionsltd/Documents/Project/GolangEcom"
LOG_DIR="$BASE_DIR/logs"
PID_DIR="$BASE_DIR/pids"

# Create directories if they don't exist
mkdir -p "$LOG_DIR"
mkdir -p "$PID_DIR"

echo -e "${GREEN}Starting all microservices locally...${NC}\n"

# Function to start a service
start_service() {
    local service_name=$1
    local service_dir=$2
    local port=$3

    echo -e "${YELLOW}Starting $service_name on port $port...${NC}"

    cd "$service_dir"

    # Start service in background
    SERVER_PORT=:$port nohup go run main.go > "$LOG_DIR/${service_name}.log" 2>&1 &

    # Save PID
    echo $! > "$PID_DIR/${service_name}.pid"

    echo -e "${GREEN}✓ $service_name started (PID: $!)${NC}"
    echo "  Log: $LOG_DIR/${service_name}.log"
    echo ""
}

# Start all services
echo "════════════════════════════════════════════════════════"
echo "  Starting Backend Services"
echo "════════════════════════════════════════════════════════"
echo ""

start_service "config-service" "$BASE_DIR/ecommerce-go-config-service" 8087
sleep 2

start_service "user-service" "$BASE_DIR/ecommerce-go-user-service" 8081
sleep 1

start_service "product-service" "$BASE_DIR/ecommerce-go-product-service" 8082
sleep 1

start_service "order-service" "$BASE_DIR/ecommerce-go-order-service" 8083
sleep 1

start_service "payment-service" "$BASE_DIR/ecommerce-go-payment-service" 8084
sleep 1

start_service "inventory-service" "$BASE_DIR/ecommerce-go-inventory-service" 8085
sleep 1

start_service "notification-service" "$BASE_DIR/ecommerce-go-notification-service" 8086
sleep 2

echo "════════════════════════════════════════════════════════"
echo "  Starting API Gateway"
echo "════════════════════════════════════════════════════════"
echo ""

start_service "api-gateway" "$BASE_DIR/ecommerce-go-api-gateway" 8080
sleep 2

echo ""
echo "════════════════════════════════════════════════════════"
echo -e "${GREEN}✅ All services started successfully!${NC}"
echo "════════════════════════════════════════════════════════"
echo ""
echo "📊 Service Status:"
echo "  - Config Service:       http://localhost:8087"
echo "  - User Service:         http://localhost:8081"
echo "  - Product Service:      http://localhost:8082"
echo "  - Order Service:        http://localhost:8083"
echo "  - Payment Service:      http://localhost:8084"
echo "  - Inventory Service:    http://localhost:8085"
echo "  - Notification Service: http://localhost:8086"
echo "  - API Gateway:          http://localhost:8080 ⭐"
echo ""
echo "🌐 Access everything through: http://localhost:8080"
echo ""
echo "📝 Useful commands:"
echo "  - View logs:     cd $BASE_DIR && tail -f logs/*.log"
echo "  - Stop all:      bash scripts/stop-all-local.sh"
echo "  - Check status:  bash scripts/status-local.sh"
echo ""
