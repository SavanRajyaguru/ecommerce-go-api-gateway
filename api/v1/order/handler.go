package order

import (
	"ecommerce-go-api-gateway/models"
	"ecommerce-go-api-gateway/pkg/utils"
	"ecommerce-go-api-gateway/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service services.OrderService
}

func NewOrderHandler(service services.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req models.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	order, err := h.service.CreateOrder(req)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to create order", err.Error())
		return
	}

	utils.SendSuccess(c, http.StatusCreated, "Order created successfully", order)
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid order ID", err.Error())
		return
	}

	order, err := h.service.GetOrder(uint(id))
	if err != nil {
		utils.SendError(c, http.StatusNotFound, "Order not found", err.Error())
		return
	}

	utils.SendSuccess(c, http.StatusOK, "Order details", order)
}
