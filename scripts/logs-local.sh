#!/bin/bash

# Script to view logs from locally running microservices

BASE_DIR="/Users/yudizsolutionsltd/Documents/Project/GolangEcom"
LOG_DIR="$BASE_DIR/logs"

# Check if a specific service was requested
if [ $# -eq 1 ]; then
    service=$1
    log_file="$LOG_DIR/${service}.log"

    if [ -f "$log_file" ]; then
        echo "Tailing logs for $service..."
        echo "Press Ctrl+C to exit"
        echo "════════════════════════════════════════════════════════"
        tail -f "$log_file"
    else
        echo "Error: Log file not found for $service"
        echo "Available logs:"
        ls -1 "$LOG_DIR"/*.log 2>/dev/null | xargs -n 1 basename
        exit 1
    fi
else
    # Show all logs
    echo "Tailing ALL service logs..."
    echo "Press Ctrl+C to exit"
    echo "════════════════════════════════════════════════════════"
    tail -f "$LOG_DIR"/*.log
fi
