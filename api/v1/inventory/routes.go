package inventory

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, handler *InventoryHandler) {
	routes := r.Group("/inventory")
	{
		routes.PUT("/stock", handler.UpdateStock)
	}
}
