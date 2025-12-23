package payment

import (
	"ecommerce-go-api-gateway/models"
	"ecommerce-go-api-gateway/pkg/utils"
	"ecommerce-go-api-gateway/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	service services.PaymentService
}

func NewPaymentHandler(service services.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: service}
}

func (h *PaymentHandler) ProcessPayment(c *gin.Context) {
	var req models.CreatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	payment, err := h.service.ProcessPayment(req)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to process payment", err.Error())
		return
	}

	utils.SendSuccess(c, http.StatusOK, "Payment processed successfully", payment)
}
