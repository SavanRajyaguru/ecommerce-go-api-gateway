package services

import (
	"ecommerce-go-api-gateway/config"
	"ecommerce-go-api-gateway/models"

	"github.com/go-resty/resty/v2"
)

type UserService interface {
	Login(req models.LoginRequest) (*models.LoginResponse, error)
	Register(req models.CreateUserRequest) (*models.User, error)
	GetUser(id uint) (*models.User, error)
}

type ProductService interface {
	GetProduct(id uint) (*models.Product, error)
	ListProducts() ([]models.Product, error)
	CreateProduct(req models.CreateProductRequest) (*models.Product, error)
}

type OrderService interface {
	CreateOrder(req models.CreateOrderRequest) (*models.Order, error)
	GetOrder(id uint) (*models.Order, error)
}

type PaymentService interface {
	ProcessPayment(req models.CreatePaymentRequest) (*models.Payment, error)
}

type InventoryService interface {
	UpdateStock(req models.UpdateInventoryRequest) error
}

type NotificationService interface {
	SendNotification(req models.SendNotificationRequest) error
}

type ServiceContainer struct {
	User         UserService
	Product      ProductService
	Order        OrderService
	Payment      PaymentService
	Inventory    InventoryService
	Notification NotificationService
}

func NewServiceContainer(cfg *config.Config) *ServiceContainer {
	client := resty.New()
	return &ServiceContainer{
		User:         NewUserService(cfg.Services.UserService, client),
		Product:      NewProductService(cfg.Services.ProductService, client),
		Order:        NewOrderService(cfg.Services.OrderService, client),
		Payment:      NewPaymentService(cfg.Services.PaymentService, client),
		Inventory:    NewInventoryService(cfg.Services.InventoryService, client),
		Notification: NewNotificationService(cfg.Services.NotificationService, client),
	}
}
