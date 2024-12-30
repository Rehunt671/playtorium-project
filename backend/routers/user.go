package routers

import (
	"playtorium/handlers"

	"github.com/gin-gonic/gin"
)

func SetupUserRouter(r *gin.RouterGroup, userHandler handlers.UserHandler) {
	userRouteV1 := r.Group("/v1/users")
	{
		userRouteV1.GET("/me", userHandler.GetMe)
	}
}
