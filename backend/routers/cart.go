package routers

import (
	"playtorium/handlers"

	"github.com/gin-gonic/gin"
)

func SetupCartRouter(r *gin.RouterGroup, cartHandler handlers.CartHandler) {
	cartRouteV1 := r.Group("/v1/carts")
	{
		cartRouteV1.GET("/user/:user_id/", cartHandler.GetCartDetail)
	}
}
