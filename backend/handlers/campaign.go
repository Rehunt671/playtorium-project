package handlers

import (
	"net/http"
	"strconv"

	"playtorium/services"

	"github.com/gin-gonic/gin"
)

type CampaignHandler interface {
	GetCampaignsByDiscountCategoryId(c *gin.Context)
}

type campaignHandler struct {
	campaignService services.CampaignService
}

func NewCampaignHandler(campaignService services.CampaignService) CampaignHandler {
	return &campaignHandler{campaignService: campaignService}
}

func (h *campaignHandler) GetCampaignsByDiscountCategoryId(c *gin.Context) {
	discountTypeIdStr := c.Query("discountCategoryId")
	if discountTypeIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "discountTypeId query parameter is required"})
		return
	}

	discountTypeId, err := strconv.Atoi(discountTypeIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "discountTypeId must be a valid integer"})
		return
	}

	campaigns, err := h.campaignService.GetCampaignsByDiscountCategoryId(c, discountTypeId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, campaigns)
}
