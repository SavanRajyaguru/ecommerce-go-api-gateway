package models

type Payment struct {
	ID        uint    `json:"id"`
	OrderID   uint    `json:"order_id"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
	Method    string  `json:"method"`
	CreatedAt string  `json:"created_at"`
}

type CreatePaymentRequest struct {
	OrderID uint    `json:"order_id" binding:"required"`
	Amount  float64 `json:"amount" binding:"required,gt=0"`
	Method  string  `json:"method" binding:"required,oneof=credit_card paypal"`
}
