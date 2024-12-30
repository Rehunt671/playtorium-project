package routers

import (
	"playtorium/handlers"

	"github.com/gin-gonic/gin"
)

func SetupItemCategoryRouter(r *gin.RouterGroup, itemCategoryHandler handlers.ItemCategoryHandler) {
	itemRouteV1 := r.Group("/v1/item-categories")
	{
		itemRouteV1.GET("", itemCategoryHandler.GetAllItemCategories)
	}
}
