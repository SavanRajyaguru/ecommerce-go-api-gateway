package api

import (
	"ecommerce-go-api-gateway/api/v1/inventory"
	"ecommerce-go-api-gateway/api/v1/middleware"
	"ecommerce-go-api-gateway/api/v1/notification"
	"ecommerce-go-api-gateway/api/v1/order"
	"ecommerce-go-api-gateway/api/v1/payment"
	"ecommerce-go-api-gateway/api/v1/product"
	"ecommerce-go-api-gateway/api/v1/user"
	"ecommerce-go-api-gateway/config"
	"ecommerce-go-api-gateway/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(middleware.Cors())

	// Initialize Service Container
	serviceContainer := services.NewServiceContainer(cfg)

	// Initialize Handlers
	userHandler := user.NewUserHandler(serviceContainer.User)
	productHandler := product.NewProductHandler(serviceContainer.Product)
	orderHandler := order.NewOrderHandler(serviceContainer.Order)
	paymentHandler := payment.NewPaymentHandler(serviceContainer.Payment)
	inventoryHandler := inventory.NewInventoryHandler(serviceContainer.Inventory)
	notificationHandler := notification.NewNotificationHandler(serviceContainer.Notification)

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API V1 Group
	v1 := r.Group("/api/v1")
	{
		user.RegisterRoutes(v1, userHandler)
		product.RegisterRoutes(v1, productHandler)
		order.RegisterRoutes(v1, orderHandler)
		payment.RegisterRoutes(v1, paymentHandler)
		inventory.RegisterRoutes(v1, inventoryHandler)
		notification.RegisterRoutes(v1, notificationHandler)
	}

	return r
}
