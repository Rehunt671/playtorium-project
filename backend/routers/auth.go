package routers

import (
	"playtorium/handlers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRouter(r *gin.RouterGroup, authHandler handlers.AuthHandler) {
	authRouteV1 := r.Group("/v1/auth")
	{
		authRouteV1.POST("/login", authHandler.Login)
		authRouteV1.POST("/register", authHandler.Register)
	}
}
