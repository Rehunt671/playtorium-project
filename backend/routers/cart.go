package routers

import (
	"playtorium/handlers"

	"github.com/gin-gonic/gin"
)

func SetupCartRouter(r *gin.RouterGroup, cartHandler handlers.CartHandler) {
	cartRouteV1 := r.Group("/v1/carts")
	{
		cartRouteV1.GET("/:cart_id/cartItem", cartHandler.GetCartItemsByCartID)
		cartRouteV1.POST("/:cart_id/cartItem", cartHandler.AddCartItem)
		cartRouteV1.PUT("/:cart_id/cartItem/:cart_item_id", cartHandler.UpdateCartItem)
		cartRouteV1.DELETE("/:cart_id/cartItem/:cart_item_id", cartHandler.RemoveCartItem)
		cartRouteV1.GET("/user/:user_id", cartHandler.GetCartDetail)
	}
}
