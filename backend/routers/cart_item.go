package routers

import (
	"playtorium/handlers"

	"github.com/gin-gonic/gin"
)

func SetupCartItemRouter(r *gin.RouterGroup, cartItemHandler handlers.CartItemHandler) {
	cartItemRouteV1 := r.Group("/v1/cartItems")
	{
		cartItemRouteV1.POST("", cartItemHandler.AddCartItem)
		cartItemRouteV1.PUT("/:id", cartItemHandler.UpdateCartItem)
		cartItemRouteV1.DELETE("/:id", cartItemHandler.RemoveCartItem)
	}
}
