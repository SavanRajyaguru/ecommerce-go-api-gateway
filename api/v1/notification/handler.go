package notification

import (
	"ecommerce-go-api-gateway/models"
	"ecommerce-go-api-gateway/pkg/utils"
	"ecommerce-go-api-gateway/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	service services.NotificationService
}

func NewNotificationHandler(service services.NotificationService) *NotificationHandler {
	return &NotificationHandler{service: service}
}

func (h *NotificationHandler) SendNotification(c *gin.Context) {
	var req models.SendNotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	err := h.service.SendNotification(req)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to send notification", err.Error())
		return
	}

	utils.SendSuccess(c, http.StatusOK, "Notification sent successfully", nil)
}
