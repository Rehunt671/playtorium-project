package routers

import (
	"playtorium/handlers"

	"github.com/gin-gonic/gin"
)

func SetupDiscountRouter(r *gin.RouterGroup, discountHandler handlers.DiscountHandler) {
	discountRouteV1 := r.Group("/v1/discounts")
	{
		discountRouteV1.POST("", discountHandler.CreateDiscount)
		discountRouteV1.PUT("", discountHandler.UpdateDiscount)
		discountRouteV1.DELETE("/:id", discountHandler.DeleteDiscount)
	}
}
