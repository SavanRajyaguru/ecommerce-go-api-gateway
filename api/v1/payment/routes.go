package payment

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, handler *PaymentHandler) {
	routes := r.Group("/payments")
	{
		routes.POST("", handler.ProcessPayment)
	}
}
