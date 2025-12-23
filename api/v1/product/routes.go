package product

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, handler *ProductHandler) {
	routes := r.Group("/products")
	{
		routes.GET("", handler.ListProducts)
		routes.GET("/:id", handler.GetProduct)
		routes.POST("", handler.CreateProduct)
	}
}
