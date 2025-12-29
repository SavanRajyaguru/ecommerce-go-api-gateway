# E-Commerce API Gateway

A high-performance API Gateway built with Go and Gin framework that serves as a unified entry point for a microservices-based e-commerce platform.

## Architecture

This API Gateway orchestrates communication between 6 backend microservices:

- **User Service** (Port 8081) - Authentication, registration, user management
- **Product Service** (Port 8082) - Product catalog and management
- **Order Service** (Port 8083) - Order creation and management
- **Payment Service** (Port 8084) - Payment processing
- **Inventory Service** (Port 8085) - Stock management
- **Notification Service** (Port 8086) - Email/SMS notifications

**API Gateway** runs on port **8080** and routes requests to appropriate services.

## Tech Stack

- **Language**: Go 1.23
- **Framework**: Gin Web Framework
- **HTTP Client**: Resty
- **Configuration**: Viper
- **Logging**: Zap (structured logging)
- **Containerization**: Docker & Docker Compose

## Prerequisites

- Go 1.23 or higher
- Docker and Docker Compose
- Make (optional, for using Makefile commands)

## Project Structure

```
ecommerce-go-api-gateway/
├── api/
│   ├── router.go              # Main router configuration
│   └── v1/                    # API version 1
│       ├── user/              # User endpoint handlers
│       ├── product/           # Product endpoint handlers
│       ├── order/             # Order endpoint handlers
│       ├── payment/           # Payment endpoint handlers
│       ├── inventory/         # Inventory endpoint handlers
│       ├── notification/      # Notification endpoint handlers
│       └── middleware/        # CORS middleware
├── cmd/api/main.go            # Application entry point
├── config/
│   ├── config.go              # Configuration loader
│   └── config.yaml            # Configuration file
├── models/                    # Data models for all domains
├── services/                  # Service layer (HTTP client wrappers)
├── pkg/
│   ├── logger/                # Logger initialization
│   └── utils/                 # Utility functions
├── Dockerfile                 # Docker build file
├── docker-compose.yaml        # Multi-service orchestration
└── Makefile                   # Build automation
```

## Quick Start

### Option 1: Using Docker Compose (Recommended for Testing)

This will start all 6 microservices plus the API Gateway.

```bash
# Start all services
make docker-up

# Or without Make
docker-compose up -d
```

The API Gateway will be available at `http://localhost:8080`

### Option 2: Running Locally

```bash
# Download dependencies
make deps

# Run the gateway (requires services to be running separately)
make run

# Or build and run
make build
./main
```

## Configuration

Configuration can be set via:

1. **config/config.yaml** - Default configuration file
2. **Environment variables** - Override YAML settings

### Environment Variables

```bash
# Server Configuration
SERVER_PORT=:8080
SERVER_MODE=debug              # or "release"

# Service URLs
SERVICES_USER_SERVICE=http://localhost:8081
SERVICES_PRODUCT_SERVICE=http://localhost:8082
SERVICES_ORDER_SERVICE=http://localhost:8083
SERVICES_PAYMENT_SERVICE=http://localhost:8084
SERVICES_INVENTORY_SERVICE=http://localhost:8085
SERVICES_NOTIFICATION_SERVICE=http://localhost:8086

# Logger
LOGGER_LEVEL=info              # debug, info, warn, error
```

Copy `.env.example` to `.env` and modify as needed:

```bash
cp .env.example .env
```

## API Endpoints

All endpoints are prefixed with `/api/v1`

### User Service
- `POST /api/v1/users/register` - Register new user
- `POST /api/v1/users/login` - User login
- `GET /api/v1/users/:id` - Get user details

### Product Service
- `GET /api/v1/products` - List all products
- `GET /api/v1/products/:id` - Get product by ID
- `POST /api/v1/products` - Create new product

### Order Service
- `POST /api/v1/orders` - Create new order
- `GET /api/v1/orders/:id` - Get order details

### Payment Service
- `POST /api/v1/payments` - Process payment

### Inventory Service
- `PUT /api/v1/inventory/stock` - Update stock levels

### Notification Service
- `POST /api/v1/notifications` - Send notification

### Health Check
- `GET /health` - Gateway health check

## Docker Commands

The Makefile provides convenient commands for Docker operations:

```bash
# Build all Docker images
make docker-build

# Start all services
make docker-up

# Start with logs visible
make docker-up-logs

# View logs from all services
make docker-logs

# View logs from specific service
make docker-logs-gateway
make docker-logs-user
make docker-logs-product

# Check service health
make docker-health

# Restart all services
make docker-restart

# Stop all services
make docker-down

# Clean up Docker resources
make docker-clean

# Rebuild everything from scratch
make docker-rebuild
```

## Development

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage
```

### Building

```bash
# Build binary
make build

# Clean build artifacts
make clean
```

### Project Requirements

All 6 backend services must be in the parent directory with these names:

```
parent-directory/
├── ecommerce-go-user-service/
├── ecommerce-go-product-service/
├── ecommerce-go-order-service/
├── ecommerce-go-payment-service/
├── ecommerce-go-inventory-service/
├── ecommerce-go-notification-service/
└── ecommerce-go-api-gateway/         # This project
```

Each service should have its own `Dockerfile`.

## Troubleshooting

### Services not communicating

Check that all services are on the same Docker network:

```bash
docker network ls
docker network inspect ecommerce-go-api-gateway_ecommerce-network
```

### Port conflicts

If ports 8080-8086 are already in use, you can modify them in `docker-compose.yaml`

### Service health checks failing

View logs to see what's wrong:

```bash
make docker-logs
```

Check if services have health check endpoints at `/health`

## Features

- RESTful API architecture
- Centralized routing and request forwarding
- CORS support for frontend integration
- Structured logging with Zap
- Graceful shutdown handling
- Health check endpoint
- Docker containerization
- Environment-based configuration

## Contributing

1. Create a feature branch
2. Make your changes
3. Run tests
4. Submit a pull request

## License

MIT
