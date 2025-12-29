#!/bin/bash

# Script to stop all locally running microservices

set -e

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

BASE_DIR="/Users/yudizsolutionsltd/Documents/Project/GolangEcom"
PID_DIR="$BASE_DIR/pids"

echo -e "${YELLOW}Stopping all microservices...${NC}\n"

# Function to stop a service
stop_service() {
    local service_name=$1
    local pid_file="$PID_DIR/${service_name}.pid"

    if [ -f "$pid_file" ]; then
        pid=$(cat "$pid_file")
        if ps -p $pid > /dev/null 2>&1; then
            echo -e "${YELLOW}Stopping $service_name (PID: $pid)...${NC}"
            kill $pid
            rm "$pid_file"
            echo -e "${GREEN}✓ $service_name stopped${NC}"
        else
            echo -e "${RED}✗ $service_name not running (stale PID file)${NC}"
            rm "$pid_file"
        fi
    else
        echo -e "${RED}✗ $service_name not found (no PID file)${NC}"
    fi
}

# Stop all services
stop_service "api-gateway"
stop_service "config-service"
stop_service "user-service"
stop_service "product-service"
stop_service "order-service"
stop_service "payment-service"
stop_service "inventory-service"
stop_service "notification-service"

# Also kill any remaining Go processes on these ports (backup)
echo ""
echo -e "${YELLOW}Checking for any remaining processes on ports 8080-8087...${NC}"

for port in {8080..8087}; do
    pid=$(lsof -ti:$port 2>/dev/null || true)
    if [ ! -z "$pid" ]; then
        echo -e "${YELLOW}Killing process on port $port (PID: $pid)${NC}"
        kill $pid 2>/dev/null || true
    fi
done

echo ""
echo -e "${GREEN}✅ All services stopped!${NC}"
echo ""
