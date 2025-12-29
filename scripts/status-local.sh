#!/bin/bash

# Script to check status of all locally running microservices

set -e

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

BASE_DIR="/Users/yudizsolutionsltd/Documents/Project/GolangEcom"
PID_DIR="$BASE_DIR/pids"

echo "════════════════════════════════════════════════════════"
echo "  Microservices Status"
echo "════════════════════════════════════════════════════════"
echo ""

# Function to check service status
check_service() {
    local service_name=$1
    local port=$2
    local pid_file="$PID_DIR/${service_name}.pid"

    printf "%-25s" "$service_name:"

    # Check PID file
    if [ -f "$pid_file" ]; then
        pid=$(cat "$pid_file")
        if ps -p $pid > /dev/null 2>&1; then
            # Check if port is listening
            if lsof -Pi :$port -sTCP:LISTEN -t >/dev/null 2>&1 ; then
                echo -e "${GREEN}✓ Running${NC} (PID: $pid, Port: $port)"

                # Test health endpoint
                health_status=$(curl -s -o /dev/null -w "%{http_code}" http://localhost:$port/health 2>/dev/null || echo "000")
                if [ "$health_status" = "200" ]; then
                    echo -e "                          ${GREEN}  Health check: OK${NC}"
                else
                    echo -e "                          ${YELLOW}  Health check: FAILED (HTTP $health_status)${NC}"
                fi
            else
                echo -e "${YELLOW}⚠ Running but not listening${NC} (PID: $pid)"
            fi
        else
            echo -e "${RED}✗ Not running${NC} (stale PID file)"
        fi
    else
        # No PID file, check if port is in use
        if lsof -Pi :$port -sTCP:LISTEN -t >/dev/null 2>&1 ; then
            pid=$(lsof -ti:$port)
            echo -e "${YELLOW}⚠ Running (no PID file)${NC} (PID: $pid, Port: $port)"
        else
            echo -e "${RED}✗ Not running${NC}"
        fi
    fi
}

# Check all services
check_service "Config Service" 8087
echo ""
check_service "User Service" 8081
echo ""
check_service "Product Service" 8082
echo ""
check_service "Order Service" 8083
echo ""
check_service "Payment Service" 8084
echo ""
check_service "Inventory Service" 8085
echo ""
check_service "Notification Service" 8086
echo ""
check_service "API Gateway" 8080
echo ""

echo "════════════════════════════════════════════════════════"

# Count running services
running_count=0
for port in 8080 8081 8082 8083 8084 8085 8086 8087; do
    if lsof -Pi :$port -sTCP:LISTEN -t >/dev/null 2>&1 ; then
        ((running_count++))
    fi
done

echo ""
if [ $running_count -eq 8 ]; then
    echo -e "${GREEN}✅ All services are running! ($running_count/8)${NC}"
else
    echo -e "${YELLOW}⚠ $running_count/8 services are running${NC}"
fi
echo ""
