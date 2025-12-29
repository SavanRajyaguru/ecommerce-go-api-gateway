# Adding New Routes - Complete Guide

## Scenario 1: Add Route ONLY to User Service (No Gateway)

**Use Case**: Testing, internal routes, or routes not needed through gateway

### Step 1: Add Route to User Service

In your User Service (`ecommerce-go-user-service`):

```go
// routes.go or handler.go
func (h *UserHandler) GetUserProfile(c *gin.Context) {
    // Your logic here
    c.JSON(200, gin.H{"profile": "user profile data"})
}

// Register route
router.GET("/profile", h.GetUserProfile)
```

### Step 2: That's It!

**No Gateway changes needed!**

### How to Access

```bash
# Direct access to User Service (port 8081)
curl http://localhost:8081/profile

# Gateway doesn't know about this route
curl http://localhost:8080/api/v1/users/profile  # ❌ 404 Not Found
```

---

## Scenario 2: Add Route Through Gateway (Recommended)

**Use Case**: Production routes, routes accessible to clients

You need to update BOTH services:

### Step 1: Add Route to User Service Backend

**File**: `ecommerce-go-user-service/handler.go` (or similar)

```go
// Add handler function
func (h *UserHandler) GetUserProfile(c *gin.Context) {
    userID := c.Param("id")

    // Your business logic
    profile := map[string]interface{}{
        "id": userID,
        "name": "John Doe",
        "email": "john@example.com",
    }

    c.JSON(200, gin.H{"profile": profile})
}

// Register route
router.GET("/users/:id/profile", h.GetUserProfile)
```

**Start/Restart User Service**:
```bash
cd ecommerce-go-user-service
SERVER_PORT=:8081 go run main.go
```

**Test Direct Access**:
```bash
curl http://localhost:8081/users/123/profile
```

---

### Step 2: Update API Gateway (5 Files)

#### 2a. Add Method to Service Interface

**File**: `ecommerce-go-api-gateway/services/interface.go`

```go
type UserService interface {
    Login(req models.LoginRequest) (*models.LoginResponse, error)
    Register(req models.CreateUserRequest) (*models.User, error)
    GetUser(id uint) (*models.User, error)
    GetUserProfile(id uint) (map[string]interface{}, error)  // ← NEW
}
```

#### 2b. Implement in User Service

**File**: `ecommerce-go-api-gateway/services/user_service.go`

```go
func (s *userService) GetUserProfile(id uint) (map[string]interface{}, error) {
    resp, err := s.client.R().
        Get(fmt.Sprintf("%s/users/%d/profile", s.baseURL, id))
        // Makes call to http://localhost:8081/users/123/profile

    if err != nil {
        return nil, err
    }
    if resp.IsError() {
        return nil, fmt.Errorf("user service error: %s", resp.String())
    }

    var profile map[string]interface{}
    if err := json.Unmarshal(resp.Body(), &profile); err != nil {
        return nil, err
    }
    return profile, nil
}
```

#### 2c. Add Handler Method

**File**: `ecommerce-go-api-gateway/api/v1/user/handler.go`

```go
func (h *UserHandler) GetUserProfile(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        utils.SendError(c, http.StatusBadRequest, "Invalid user ID", err.Error())
        return
    }

    profile, err := h.service.GetUserProfile(uint(id))
    if err != nil {
        utils.SendError(c, http.StatusNotFound, "Profile not found", err.Error())
        return
    }

    utils.SendSuccess(c, http.StatusOK, "User profile", profile)
}
```

#### 2d. Register Route

**File**: `ecommerce-go-api-gateway/api/v1/user/routes.go`

```go
func RegisterRoutes(r *gin.RouterGroup, handler *UserHandler) {
    routes := r.Group("/users")
    {
        routes.POST("/register", handler.Register)
        routes.POST("/login", handler.Login)
        routes.GET("/:id", handler.GetUser)
        routes.GET("/:id/profile", handler.GetUserProfile)  // ← NEW
    }
}
```

#### 2e. Add Model (Optional)

**File**: `ecommerce-go-api-gateway/models/user.go`

If you need a specific response type:

```go
type UserProfile struct {
    ID    uint   `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    Bio   string `json:"bio"`
}
```

---

### Step 3: Restart Gateway

```bash
cd ecommerce-go-api-gateway
go run cmd/api/main.go
```

---

### Step 4: Test Through Gateway

```bash
# Through Gateway (port 8080) ✅
curl http://localhost:8080/api/v1/users/123/profile

# Direct to User Service (port 8081) ✅
curl http://localhost:8081/users/123/profile
```

**Both work!**

---

## Quick Comparison

### Scenario 1: User Service Only
```
You add route → User Service only
Access: http://localhost:8081/profile
Gateway: Not aware of this route
Files changed: 1 (User Service)
```

### Scenario 2: Through Gateway
```
You add route → User Service + Gateway (5 files)
Access: http://localhost:8080/api/v1/users/123/profile
Gateway: Forwards to User Service
Files changed: 6 (User Service + 5 Gateway files)
```

---

## Step-by-Step Example: Add "Update User" Route

### Want to add: `PUT /api/v1/users/:id` (Update user)

#### Step 1: User Service Backend

**Terminal 1** (User Service):
```bash
cd ecommerce-go-user-service
```

Edit handler:
```go
func (h *UserHandler) UpdateUser(c *gin.Context) {
    var req UpdateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    // Update logic here
    c.JSON(200, gin.H{"message": "User updated"})
}

// Register
router.PUT("/users/:id", h.UpdateUser)
```

Restart:
```bash
SERVER_PORT=:8081 go run main.go
```

Test:
```bash
curl -X PUT http://localhost:8081/users/123 \
  -H "Content-Type: application/json" \
  -d '{"name":"Updated Name"}'
```

---

#### Step 2: API Gateway

**File 1**: `services/interface.go`
```go
type UserService interface {
    // ... existing methods
    UpdateUser(id uint, req models.UpdateUserRequest) (*models.User, error)  // NEW
}
```

**File 2**: `services/user_service.go`
```go
func (s *userService) UpdateUser(id uint, req models.UpdateUserRequest) (*models.User, error) {
    resp, err := s.client.R().
        SetBody(req).
        Put(fmt.Sprintf("%s/users/%d", s.baseURL, id))

    if err != nil {
        return nil, err
    }

    var user models.User
    json.Unmarshal(resp.Body(), &user)
    return &user, nil
}
```

**File 3**: `api/v1/user/handler.go`
```go
func (h *UserHandler) UpdateUser(c *gin.Context) {
    idStr := c.Param("id")
    id, _ := strconv.ParseUint(idStr, 10, 32)

    var req models.UpdateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        utils.SendError(c, 400, "Invalid request", err.Error())
        return
    }

    user, err := h.service.UpdateUser(uint(id), req)
    if err != nil {
        utils.SendError(c, 500, "Update failed", err.Error())
        return
    }

    utils.SendSuccess(c, 200, "User updated", user)
}
```

**File 4**: `api/v1/user/routes.go`
```go
func RegisterRoutes(r *gin.RouterGroup, handler *UserHandler) {
    routes := r.Group("/users")
    {
        routes.POST("/register", handler.Register)
        routes.POST("/login", handler.Login)
        routes.GET("/:id", handler.GetUser)
        routes.PUT("/:id", handler.UpdateUser)  // NEW
    }
}
```

**File 5**: `models/user.go`
```go
type UpdateUserRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

---

**Terminal 8** (Gateway):
```bash
cd ecommerce-go-api-gateway
go run cmd/api/main.go
```

Test:
```bash
curl -X PUT http://localhost:8080/api/v1/users/123 \
  -H "Content-Type: application/json" \
  -d '{"name":"Updated Name"}'
```

---

## Files to Change Checklist

When adding a new route through gateway:

- [ ] **User Service** (1 file)
  - [ ] Add handler function
  - [ ] Register route

- [ ] **API Gateway** (5 files)
  - [ ] `services/interface.go` - Add method to interface
  - [ ] `services/user_service.go` - Implement method
  - [ ] `api/v1/user/handler.go` - Add handler function
  - [ ] `api/v1/user/routes.go` - Register route
  - [ ] `models/user.go` - Add request/response models (if needed)

- [ ] **Restart Both Services**
  - [ ] Restart User Service (Terminal 2)
  - [ ] Restart Gateway (Terminal 8)

---

## Common Patterns

### GET Request (No Body)
```go
// Service
resp, err := s.client.R().Get(s.baseURL + "/endpoint")

// Handler
result, err := h.service.GetSomething()
```

### POST Request (With Body)
```go
// Service
resp, err := s.client.R().
    SetBody(req).
    Post(s.baseURL + "/endpoint")

// Handler
result, err := h.service.CreateSomething(req)
```

### PUT/PATCH Request
```go
// Service
resp, err := s.client.R().
    SetBody(req).
    Put(fmt.Sprintf("%s/endpoint/%d", s.baseURL, id))

// Handler
result, err := h.service.UpdateSomething(id, req)
```

### DELETE Request
```go
// Service
resp, err := s.client.R().
    Delete(fmt.Sprintf("%s/endpoint/%d", s.baseURL, id))

// Handler
err := h.service.DeleteSomething(id)
```

---

## Testing Workflow

### Test Backend Service First
```bash
# 1. Add route to User Service
# 2. Restart User Service
SERVER_PORT=:8081 go run main.go

# 3. Test directly
curl http://localhost:8081/new-endpoint

# If this works ✓, proceed to gateway
```

### Then Update Gateway
```bash
# 1. Update 5 gateway files
# 2. Restart Gateway
go run cmd/api/main.go

# 3. Test through gateway
curl http://localhost:8080/api/v1/users/new-endpoint
```

---

## Pro Tips

### 1. Test Backend First
Always test the backend service directly before adding to gateway:
```bash
curl http://localhost:8081/your-new-route
```

### 2. Use Postman Collections
Create a collection with:
- Direct service URLs (8081-8087)
- Gateway URLs (8080)

### 3. Keep Interfaces in Sync
Gateway interfaces should match backend routes.

### 4. Hot Reload During Development
When developing:
- Edit User Service → Ctrl+C → Restart (2 sec)
- Edit Gateway → Ctrl+C → Restart (2 sec)

### 5. Use Same Models
If possible, share model definitions between services.

---

## Summary

### Adding Route to User Service Only:
1. Add handler to User Service
2. Register route
3. Restart User Service
4. Access via `http://localhost:8081/route`

**Gateway changes**: NONE ❌

---

### Adding Route Through Gateway:
1. Add handler to User Service
2. Register route in User Service
3. Add method to Gateway interface
4. Implement in Gateway service
5. Add Gateway handler
6. Register in Gateway routes
7. Restart BOTH services
8. Access via `http://localhost:8080/api/v1/users/route`

**Gateway changes**: 5 files ✅

---

## Quick Reference

| Action | User Service Changes | Gateway Changes |
|--------|---------------------|-----------------|
| **New route for direct access only** | Add handler + route | None |
| **New route through gateway** | Add handler + route | 5 files (interface, service, handler, routes, models) |
| **Modify existing route logic** | Update handler | None (if interface unchanged) |
| **Change route path** | Update route | Update route in gateway |
| **Change request/response format** | Update handler | Update models + service |

---

The key point: **Gateway is just a proxy** - it needs to know about routes to forward them. If you don't add to gateway, the route only works on the backend service directly.
