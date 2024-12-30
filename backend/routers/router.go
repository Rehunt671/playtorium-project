package routers

import (
	"net/http"
	"playtorium/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, authHandler handlers.AuthHandler, cartHandler handlers.CartHandler, campaignHandler handlers.CampaignHandler, discountHandler handlers.DiscountHandler, discountCategoryHandler handlers.DiscountCategoryHandler, itemHandler handlers.ItemHandler, itemCategoryHandler handlers.ItemCategoryHandler, userHandler handlers.UserHandler) {
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the api v1",
		})
	})
	router := r.Group("/api")
	SetupAuthRouter(router, authHandler)
	SetupCampaignRouter(router, campaignHandler)
	SetupCartRouter(router, cartHandler)
	SetupDiscountRouter(router, discountHandler)
	SetupDiscountCategoryRouter(router, discountCategoryHandler)
	SetupItemRouter(router, itemHandler)
	SetupItemCategoryRouter(router, itemCategoryHandler)
	SetupUserRouter(router, userHandler)
}
