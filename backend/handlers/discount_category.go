package handlers

import (
	"net/http"
	"playtorium/services"

	"github.com/gin-gonic/gin"
)

type DiscountCategoryHandler interface {
	GetAvailableDiscountCategories(c *gin.Context)
}

type discountCategoryHandlerImpl struct {
	discountCategoryService services.DiscountCategoryService
}

func NewDiscountCategoryHandler(discountCategoryService services.DiscountCategoryService) DiscountCategoryHandler {
	return &discountCategoryHandlerImpl{discountCategoryService: discountCategoryService}
}

func (h *discountCategoryHandlerImpl) GetAvailableDiscountCategories(c *gin.Context) {
	discountCategories, err := h.discountCategoryService.GetAvailableDiscountCategories(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, discountCategories)
}
