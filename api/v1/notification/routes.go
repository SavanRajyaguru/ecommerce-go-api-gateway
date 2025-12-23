package notification

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, handler *NotificationHandler) {
	routes := r.Group("/notifications")
	{
		routes.POST("", handler.SendNotification)
	}
}
