package order

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, handler *OrderHandler) {
	routes := r.Group("/orders")
	{
		routes.POST("", handler.CreateOrder)
		routes.GET("/:id", handler.GetOrder)
	}
}
