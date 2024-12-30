package handlers

import (
	"net/http"
	"playtorium/services"

	"github.com/gin-gonic/gin"
)

type ItemHandler interface {
	GetAllItem(c *gin.Context)
}

type itemHandler struct {
	itemService services.ItemService
}

func NewItemHandler(itemService services.ItemService) ItemHandler {
	return &itemHandler{itemService: itemService}
}

func (h *itemHandler) GetAllItem(c *gin.Context) {
	items, err := h.itemService.GetAllItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch items"})
		return
	}

	c.JSON(http.StatusOK, items)
}
