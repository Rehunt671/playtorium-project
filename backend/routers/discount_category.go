package routers

import (
	"playtorium/handlers"

	"github.com/gin-gonic/gin"
)

func SetupDiscountCategoryRouter(r *gin.RouterGroup, discountCategoryHandler handlers.DiscountCategoryHandler) {
	discountCategoryRouteV1 := r.Group("/v1/discount-categories")
	{
		discountCategoryRouteV1.GET("", discountCategoryHandler.GetAvailableDiscountCategories)
	}
}
