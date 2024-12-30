package routers

import (
	"playtorium/handlers"

	"github.com/gin-gonic/gin"
)

func SetupCampaignRouter(r *gin.RouterGroup, campaignHandler handlers.CampaignHandler) {
	campaignRouteV1 := r.Group("/v1/campaigns")
	{
		campaignRouteV1.GET("", campaignHandler.GetCampaignsByDiscountCategoryId)
	}
}
