# Quick Start - One Page Reference

## 8 Commands to Run (One Per Terminal)

```bash
# TERMINAL 1 - Config Service
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-config-service
SERVER_PORT=:8087 go run main.go

# TERMINAL 2 - User Service
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-user-service
SERVER_PORT=:8081 go run main.go

# TERMINAL 3 - Product Service
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-product-service
SERVER_PORT=:8082 go run main.go

# TERMINAL 4 - Order Service
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-order-service
SERVER_PORT=:8083 go run main.go

# TERMINAL 5 - Payment Service
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-payment-service
SERVER_PORT=:8084 go run main.go

# TERMINAL 6 - Inventory Service
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-inventory-service
SERVER_PORT=:8085 go run main.go

# TERMINAL 7 - Notification Service
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-notification-service
SERVER_PORT=:8086 go run main.go

# TERMINAL 8 - API Gateway ⭐
cd /Users/yudizsolutionsltd/Documents/Project/GolangEcom/ecommerce-go-api-gateway
go run cmd/api/main.go
```

## Access Everything Through

**http://localhost:8080** ⭐

## Test

```bash
curl http://localhost:8080/health
curl http://localhost:8080/api/v1/products
curl http://localhost:8080/api/v1/users/1
```

## Stop

Press **Ctrl+C** in each terminal

---

**For detailed guide, see: MANUAL_SETUP.md**
