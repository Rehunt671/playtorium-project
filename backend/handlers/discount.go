package handlers

import (
	"net/http"
	"playtorium/models"
	"playtorium/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DiscountHandler interface {
	GetDiscounts(c *gin.Context)
	CreateDiscount(c *gin.Context)
	UpdateDiscount(c *gin.Context)
	DeleteDiscount(c *gin.Context)
}

type discountHandlerImpl struct {
	discountService services.DiscountService
}

func NewDiscountHandler(discountService services.DiscountService) DiscountHandler {
	return &discountHandlerImpl{discountService: discountService}
}

func (h *discountHandlerImpl) GetDiscounts(c *gin.Context) {
	discounts, err := h.discountService.GetDiscounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, discounts)
}

func (h *discountHandlerImpl) CreateDiscount(c *gin.Context) {
	var discount models.Discount
	if err := c.ShouldBindJSON(&discount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.discountService.CreateDiscount(&discount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, discount)
}

func (h *discountHandlerImpl) UpdateDiscount(c *gin.Context) {
	var discount models.Discount
	if err := c.ShouldBindJSON(&discount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.discountService.UpdateDiscount(&discount); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Discount not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, discount)
}

func (h *discountHandlerImpl) DeleteDiscount(c *gin.Context) {
	discountIdStr := c.Param("id")
	discountId, err := strconv.Atoi(discountIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid discount ID"})
		return
	}

	if err := h.discountService.DeleteDiscount(uint(discountId)); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Discount not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Discount deleted successfully"})
}
