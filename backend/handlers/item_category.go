package handlers

import (
	"net/http"
	"playtorium/services"

	"github.com/gin-gonic/gin"
)

type ItemCategoryHandler interface {
	GetAllItemCategories(c *gin.Context)
}

type itemCategoryHandler struct {
	itemCategoryService services.ItemCategoryService
}

func NewItemCategoryHandler(itemCategoryService services.ItemCategoryService) ItemCategoryHandler {
	return &itemCategoryHandler{itemCategoryService: itemCategoryService}
}

func (h *itemCategoryHandler) GetAllItemCategories(c *gin.Context) {
	categories, err := h.itemCategoryService.GetAllItemCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch item categories"})
		return
	}

	c.JSON(http.StatusOK, categories)
}
