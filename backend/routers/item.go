package routers

import (
	"playtorium/handlers"

	"github.com/gin-gonic/gin"
)

func SetupItemRouter(r *gin.RouterGroup, itemHandler handlers.ItemHandler) {
	itemRouteV1 := r.Group("/v1/items")
	{
		itemRouteV1.GET("", itemHandler.GetAllItem)
	}
}
