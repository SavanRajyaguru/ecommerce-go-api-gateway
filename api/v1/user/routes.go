package user

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, handler *UserHandler) {
	routes := r.Group("/users")
	{
		routes.POST("/register", handler.Register)
		routes.POST("/login", handler.Login)
		routes.GET("/:id", handler.GetUser)
	}
}
