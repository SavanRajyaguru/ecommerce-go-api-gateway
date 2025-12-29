.PHONY: help build run test clean docker-build docker-up docker-down docker-logs docker-restart deps dev-up dev-down dev-logs

# Default target
help:
	@echo "Available commands:"
	@echo ""
	@echo "⚡ Quick Local Development (RECOMMENDED):"
	@echo "  make local-start     - Start ALL 8 services locally (one command)"
	@echo "  make local-stop      - Stop all local services"
	@echo "  make local-status    - Check which services are running"
	@echo "  make local-logs      - View all logs in real-time"
	@echo "  make local-restart   - Restart all services"
	@echo ""
	@echo "Single Service Development:"
	@echo "  make build           - Build the API Gateway binary"
	@echo "  make run             - Run the API Gateway locally"
	@echo "  make test            - Run tests"
	@echo "  make clean           - Clean build artifacts"
	@echo "  make deps            - Download dependencies"
	@echo ""
	@echo "Production Docker (Optimized builds, no hot reload):"
	@echo "  make docker-build    - Build all Docker images for production"
	@echo "  make docker-up       - Start all services (production mode)"
	@echo "  make docker-down     - Stop all services"
	@echo "  make docker-logs     - View logs from all services"
	@echo "  make docker-restart  - Restart all services"
	@echo "  make docker-clean    - Stop and remove all containers, networks, and volumes"
	@echo ""
	@echo "Development Docker (Hot reload with Air):"
	@echo "  make dev-up          - Start all services in DEV mode with hot reload"
	@echo "  make dev-down        - Stop all dev services"
	@echo "  make dev-logs        - View dev logs"
	@echo "  make dev-rebuild     - Rebuild and restart dev services"
	@echo ""

# ============================================
# LOCAL DEVELOPMENT (No Docker - Simple!)
# ============================================

# Start all services locally
local-start:
	@bash scripts/start-all-local.sh

# Stop all local services
local-stop:
	@bash scripts/stop-all-local.sh

# Check status of local services
local-status:
	@bash scripts/status-local.sh

# View logs from all services
local-logs:
	@bash scripts/logs-local.sh

# View logs from specific service
local-logs-service:
	@echo "Usage: make local-logs-service SERVICE=user-service"
	@bash scripts/logs-local.sh $(SERVICE)

# Restart all local services
local-restart:
	@echo "Restarting all services..."
	@bash scripts/stop-all-local.sh
	@sleep 2
	@bash scripts/start-all-local.sh

# Clean log files
local-clean:
	@echo "Cleaning log files..."
	@rm -rf /Users/yudizsolutionsltd/Documents/Project/GolangEcom/logs/*.log
	@rm -rf /Users/yudizsolutionsltd/Documents/Project/GolangEcom/pids/*.pid
	@echo "✓ Logs cleaned"

# ============================================
# SINGLE SERVICE DEVELOPMENT
# ============================================

# Build the application
build:
	@echo "Building API Gateway..."
	@go build -o main ./cmd/api

# Run the application locally
run:
	@echo "Running API Gateway..."
	@go run ./cmd/api/main.go

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -f main
	@rm -f coverage.out coverage.html
	@go clean

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy

# Build Docker images
docker-build:
	@echo "Building Docker images..."
	@docker-compose build

# Start all services
docker-up:
	@echo "Starting all services..."
	@docker-compose up -d
	@echo "Services started! API Gateway available at http://localhost:8080"
	@echo "Use 'make docker-logs' to view logs"

# Start all services with logs
docker-up-logs:
	@echo "Starting all services with logs..."
	@docker-compose up

# Stop all services
docker-down:
	@echo "Stopping all services..."
	@docker-compose down

# View logs
docker-logs:
	@docker-compose logs -f

# View logs for specific service
docker-logs-gateway:
	@docker-compose logs -f api-gateway

docker-logs-user:
	@docker-compose logs -f user-service

docker-logs-product:
	@docker-compose logs -f product-service

docker-logs-order:
	@docker-compose logs -f order-service

docker-logs-payment:
	@docker-compose logs -f payment-service

docker-logs-inventory:
	@docker-compose logs -f inventory-service

docker-logs-notification:
	@docker-compose logs -f notification-service

# Restart all services
docker-restart:
	@echo "Restarting all services..."
	@docker-compose restart

# Restart specific service
docker-restart-gateway:
	@docker-compose restart api-gateway

# Clean Docker resources
docker-clean:
	@echo "Cleaning Docker resources..."
	@docker-compose down -v --remove-orphans
	@docker system prune -f

# Check service health
docker-health:
	@echo "Checking service health..."
	@docker-compose ps

# Rebuild and restart all services
docker-rebuild:
	@echo "Rebuilding and restarting all services..."
	@docker-compose down
	@docker-compose build --no-cache
	@docker-compose up -d
	@echo "Services rebuilt and restarted!"

# ============================================
# DEVELOPMENT MODE COMMANDS (with hot reload)
# ============================================

# Start development environment with hot reload
dev-up:
	@echo "Starting services in DEVELOPMENT mode with hot reload..."
	@docker-compose -f docker-compose.dev.yaml up -d
	@echo ""
	@echo "✅ Development environment started!"
	@echo "📝 Code changes will auto-reload (no rebuild needed)"
	@echo "🌐 API Gateway: http://localhost:8080"
	@echo "📊 View logs: make dev-logs"
	@echo ""

# Start dev with logs visible
dev-up-logs:
	@echo "Starting dev services with logs..."
	@docker-compose -f docker-compose.dev.yaml up

# Stop development services
dev-down:
	@echo "Stopping development services..."
	@docker-compose -f docker-compose.dev.yaml down

# View development logs
dev-logs:
	@docker-compose -f docker-compose.dev.yaml logs -f

# View logs for specific dev service
dev-logs-gateway:
	@docker-compose -f docker-compose.dev.yaml logs -f api-gateway

dev-logs-user:
	@docker-compose -f docker-compose.dev.yaml logs -f user-service

# Rebuild dev containers (only needed if Dockerfile.dev changes)
dev-rebuild:
	@echo "Rebuilding dev containers..."
	@docker-compose -f docker-compose.dev.yaml down
	@docker-compose -f docker-compose.dev.yaml build --no-cache
	@docker-compose -f docker-compose.dev.yaml up -d
	@echo "Dev services rebuilt!"

# Clean dev resources
dev-clean:
	@echo "Cleaning dev resources..."
	@docker-compose -f docker-compose.dev.yaml down -v --remove-orphans

# Check dev service status
dev-health:
	@docker-compose -f docker-compose.dev.yaml ps
