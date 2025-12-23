package models

type Order struct {
	ID        uint        `json:"id"`
	UserID    uint        `json:"user_id"`
	Total     float64     `json:"total"`
	Status    string      `json:"status"`
	Items     []OrderItem `json:"items"`
	CreatedAt string      `json:"created_at"`
}

type OrderItem struct {
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type CreateOrderRequest struct {
	Items []OrderItemRequest `json:"items" binding:"required,dive"`
}

type OrderItemRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,gt=0"`
}
