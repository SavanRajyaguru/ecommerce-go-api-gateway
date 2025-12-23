package inventory

import (
	"ecommerce-go-api-gateway/models"
	"ecommerce-go-api-gateway/pkg/utils"
	"ecommerce-go-api-gateway/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InventoryHandler struct {
	service services.InventoryService
}

func NewInventoryHandler(service services.InventoryService) *InventoryHandler {
	return &InventoryHandler{service: service}
}

func (h *InventoryHandler) UpdateStock(c *gin.Context) {
	var req models.UpdateInventoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	err := h.service.UpdateStock(req)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to update stock", err.Error())
		return
	}

	utils.SendSuccess(c, http.StatusOK, "Stock updated successfully", nil)
}
