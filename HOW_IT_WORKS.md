# How API Gateway Works - Technical Explanation

## Architecture Overview

Your API Gateway uses the **Reverse Proxy Pattern** - it receives requests on **one port (8080)** and forwards them to backend services on different ports.

---

## Request Flow

```
Client Request
    ↓
http://localhost:8080/api/v1/users/login
    ↓
API Gateway (Port 8080)
    ↓
Gin Router matches route → /api/v1/users/login
    ↓
UserHandler.Login() is called
    ↓
UserService.Login() makes HTTP request
    ↓
http://localhost:8081/login (User Service)
    ↓
User Service processes request
    ↓
Returns response to Gateway
    ↓
Gateway returns response to Client
```

---

## Code Flow Explanation

### 1. Main Entry Point (`cmd/api/main.go`)

```go
func main() {
    // Load config (reads config.yaml)
    cfg := config.LoadConfig()

    // Setup router with all routes
    r := api.SetupRouter(cfg)

    // Start server on port 8080
    srv := &http.Server{
        Addr:    cfg.Server.Port,  // ":8080"
        Handler: r,
    }

    srv.ListenAndServe()
}
```

**What happens**: Server starts listening on port 8080

---

### 2. Router Setup (`api/router.go`)

```go
func SetupRouter(cfg *config.Config) *gin.Engine {
    r := gin.New()

    // Create HTTP clients for all backend services
    serviceContainer := services.NewServiceContainer(cfg)

    // Create handlers that use these clients
    userHandler := user.NewUserHandler(serviceContainer.User)
    productHandler := product.NewProductHandler(serviceContainer.Product)
    // ... etc

    // Register routes
    v1 := r.Group("/api/v1")
    {
        user.RegisterRoutes(v1, userHandler)
        product.RegisterRoutes(v1, productHandler)
        // ... etc
    }

    return r
}
```

**What happens**:
- Creates HTTP clients that know backend service URLs
- Registers routes that map to handlers

---

### 3. Service Container (`services/interface.go`)

```go
func NewServiceContainer(cfg *config.Config) *ServiceContainer {
    client := resty.New()  // HTTP client

    return &ServiceContainer{
        User: NewUserService(
            cfg.Services.UserService,  // "http://localhost:8081"
            client
        ),
        Product: NewProductService(
            cfg.Services.ProductService,  // "http://localhost:8082"
            client
        ),
        // ... all services
    }
}
```

**What happens**: Creates service clients with backend URLs from config

---

### 4. Service Implementation (`services/user_service.go`)

```go
type userService struct {
    baseURL string           // "http://localhost:8081"
    client  *resty.Client    // HTTP client
}

func (s *userService) Login(req models.LoginRequest) (*models.LoginResponse, error) {
    // Make HTTP POST to backend service
    resp, err := s.client.R().
        SetBody(req).
        Post(s.baseURL + "/login")  // http://localhost:8081/login

    // Parse response
    var loginResp models.LoginResponse
    json.Unmarshal(resp.Body(), &loginResp)
    return &loginResp, nil
}
```

**What happens**: Makes actual HTTP call to backend service on port 8081

---

### 5. Handler (`api/v1/user/handler.go`)

```go
func (h *UserHandler) Login(c *gin.Context) {
    // Parse incoming request from client
    var req models.LoginRequest
    c.ShouldBindJSON(&req)

    // Call backend service
    resp, err := h.service.Login(req)

    // Return response to client
    utils.SendSuccess(c, http.StatusOK, "Login successful", resp)
}
```

**What happens**:
- Receives request from client on port 8080
- Calls backend service
- Returns response to client

---

### 6. Route Registration (`api/v1/user/routes.go`)

```go
func RegisterRoutes(r *gin.RouterGroup, handler *UserHandler) {
    routes := r.Group("/users")
    {
        routes.POST("/register", handler.Register)  // /api/v1/users/register
        routes.POST("/login", handler.Login)        // /api/v1/users/login
        routes.GET("/:id", handler.GetUser)         // /api/v1/users/:id
    }
}
```

**What happens**: Maps URL paths to handler functions

---

## Configuration (`config/config.yaml`)

```yaml
server:
  port: ":8080"          # Gateway listens here

services:
  user_service: "http://localhost:8081"        # Backend service URLs
  product_service: "http://localhost:8082"
  order_service: "http://localhost:8083"
  payment_service: "http://localhost:8084"
  inventory_service: "http://localhost:8085"
  notification_service: "http://localhost:8086"
  config_service: "http://localhost:8087"
```

**What happens**: Gateway knows where each backend service is running

---

## Complete Example: User Login

### Step 1: Client sends request

```bash
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"pass123"}'
```

### Step 2: Gin Router matches route

```
POST /api/v1/users/login → UserHandler.Login()
```

### Step 3: Handler processes request

```go
// api/v1/user/handler.go
func (h *UserHandler) Login(c *gin.Context) {
    var req models.LoginRequest
    c.ShouldBindJSON(&req)  // Parse {"email":"test@example.com",...}

    resp, err := h.service.Login(req)  // Call backend

    utils.SendSuccess(c, http.StatusOK, "Login successful", resp)
}
```

### Step 4: Service makes HTTP call to backend

```go
// services/user_service.go
func (s *userService) Login(req models.LoginRequest) (*models.LoginResponse, error) {
    resp, err := s.client.R().
        SetBody(req).
        Post(s.baseURL + "/login")
        // Makes HTTP POST to http://localhost:8081/login

    var loginResp models.LoginResponse
    json.Unmarshal(resp.Body(), &loginResp)
    return &loginResp, nil
}
```

### Step 5: Backend service responds

```
User Service (Port 8081) processes login
    ↓
Returns {"token":"xyz123","user":{...}}
    ↓
Gateway receives response
```

### Step 6: Gateway returns to client

```json
{
  "status": "success",
  "message": "Login successful",
  "data": {
    "token": "xyz123",
    "user": {...}
  }
}
```

---

## How It Routes to Different Services

### User Service (Port 8081)

```
POST http://localhost:8080/api/v1/users/register
    → Gateway forwards to → http://localhost:8081/register

POST http://localhost:8080/api/v1/users/login
    → Gateway forwards to → http://localhost:8081/login

GET http://localhost:8080/api/v1/users/1
    → Gateway forwards to → http://localhost:8081/users/1
```

### Product Service (Port 8082)

```
GET http://localhost:8080/api/v1/products
    → Gateway forwards to → http://localhost:8082/products

POST http://localhost:8080/api/v1/products
    → Gateway forwards to → http://localhost:8082/products

GET http://localhost:8080/api/v1/products/1
    → Gateway forwards to → http://localhost:8082/products/1
```

### Order Service (Port 8083)

```
POST http://localhost:8080/api/v1/orders
    → Gateway forwards to → http://localhost:8083/orders

GET http://localhost:8080/api/v1/orders/1
    → Gateway forwards to → http://localhost:8083/orders/1
```

---

## Key Components

### 1. **Gin Router**
- Listens on port 8080
- Matches incoming URLs to handlers
- Example: `/api/v1/users/login` → `UserHandler.Login()`

### 2. **Handlers**
- Receive requests from clients
- Parse request data
- Call appropriate service
- Return responses

### 3. **Services**
- HTTP clients for backend services
- Know backend service URLs (from config)
- Make HTTP requests to backends
- Parse and return responses

### 4. **Resty Client**
- HTTP library used to call backend services
- Handles request/response
- Used by all services

### 5. **Config**
- Stores backend service URLs
- Loaded at startup
- Used to initialize service clients

---

## Why Single Port Works

**Gateway acts as a proxy**:

1. **Client** connects to one port: `localhost:8080`
2. **Gateway** has routes for all services
3. Based on URL path, gateway knows which backend service to call
4. **Gateway** makes internal HTTP call to backend service
5. **Backend** processes and returns response
6. **Gateway** returns response to client

**Client only needs to know port 8080!**

---

## Advantages of This Architecture

### ✅ Single Entry Point
- Client only needs to know `http://localhost:8080`
- No need to track multiple service ports

### ✅ Service Discovery
- Backend service URLs configured in one place
- Easy to change backend ports/hosts

### ✅ Centralized Logic
- CORS, authentication, logging in one place
- Consistent error handling

### ✅ Load Balancing (Future)
- Can add multiple instances of same service
- Gateway can distribute requests

### ✅ Security
- Backend services can be on private network
- Only gateway exposed to clients

---

## Running Commands

Based on this architecture, you need:

### Terminal 1-7: Backend Services

Each backend service must:
1. Listen on its configured port (8081-8087)
2. Have the endpoints that gateway expects

```bash
# User Service
cd /path/to/user-service
SERVER_PORT=:8081 go run main.go
```

### Terminal 8: API Gateway

```bash
cd /path/to/api-gateway
go run cmd/api/main.go
```

Gateway reads `config.yaml` and knows to forward:
- `/api/v1/users/*` → `http://localhost:8081/*`
- `/api/v1/products/*` → `http://localhost:8082/*`
- etc.

---

## Summary

Your API Gateway is a **reverse proxy** that:

1. **Listens** on one port (8080)
2. **Routes** requests based on URL path
3. **Forwards** requests to backend services via HTTP
4. **Returns** backend responses to clients

**It's like a receptionist** - clients talk to the receptionist (port 8080), and the receptionist calls the right department (backend service) and relays the response back!

---

## Testing the Flow

```bash
# 1. Start User Service (Terminal 1)
cd /path/to/user-service
SERVER_PORT=:8081 go run main.go

# 2. Start Gateway (Terminal 2)
cd /path/to/api-gateway
go run cmd/api/main.go

# 3. Test (Terminal 3)
# Client calls gateway on port 8080
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"pass123"}'

# Behind the scenes:
# Gateway receives request on :8080
# Gateway forwards to http://localhost:8081/login
# User service processes on :8081
# Returns response to gateway
# Gateway returns to client
```

**Result**: Client only knows about port 8080! ✅
