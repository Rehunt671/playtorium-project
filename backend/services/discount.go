package services

import (
	"log"
	"playtorium/models"

	"gorm.io/gorm"
)

type DiscountService interface {
	ApplyDiscounts(cartItems []models.CartItem) (float64, int, error)
}

type DiscountServiceImpl struct {
	db *gorm.DB
}

func NewDiscountService(db *gorm.DB) DiscountService {
	return &DiscountServiceImpl{db: db}
}

func (s *DiscountServiceImpl) ApplyDiscounts(cartItems []models.CartItem) (float64, int, error) {
	var total float64
	for _, cartItem := range cartItems {
		total += cartItem.Item.Price * float64(cartItem.Quantity)
	}

	initialTotal := total

	var discounts []models.Discount
	if err := s.db.Preload("FixedAmount").Preload("Percentage").Preload("PercentageCategory.ItemCategory").Preload("PointDiscount").Preload("Seasonal").Preload("DiscountCategory").Find(&discounts).Error; err != nil {
		return 0, 0, err
	}

	var fixedAmountDiscount *models.FixedAmountDiscount
	var percentageDiscount *models.PercentageDiscount
	var percentageCategoryDiscount *models.PercentageCategoryDiscount
	var pointDiscount *models.PointDiscount
	var seasonalDiscount *models.SeasonalDiscount

	for _, discount := range discounts {
		if discount.FixedAmount != nil {
			fixedAmountDiscount = discount.FixedAmount
		}
		if discount.Percentage != nil {
			percentageDiscount = discount.Percentage
		}
		if discount.PercentageCategory != nil {
			percentageCategoryDiscount = discount.PercentageCategory
		}
		if discount.PointDiscount != nil {
			pointDiscount = discount.PointDiscount
		}
		if discount.Seasonal != nil {
			seasonalDiscount = discount.Seasonal
		}
	}

	log.Printf("Initial total: %.2f", total)

	total = applyFixedAmountDiscount(total, fixedAmountDiscount)
	log.Printf("After fixed amount discount: %.2f", total)

	total = applyPercentageDiscount(total, initialTotal, percentageDiscount)
	log.Printf("After percentage discount: %.2f", total)

	total = applyCategoryDiscount(total, cartItems, percentageCategoryDiscount)
	log.Printf("After category discount: %.2f", total)

	total, maxPointsUsed := applyPointsDiscount(total, pointDiscount)
	log.Printf("After points discount: %.2f", total)

	total = applySeasonalDiscount(total, seasonalDiscount)
	log.Printf("After seasonal discount: %.2f", total)

	if total < 0 {
		total = 0
	}

	log.Printf("Final total: %.2f", total)

	return total, maxPointsUsed, nil
}

func applyFixedAmountDiscount(total float64, discount *models.FixedAmountDiscount) float64 {
	if discount != nil {
		total -= discount.Amount
	}
	if total < 0 {
		total = 0
	}
	return total
}

func applyPercentageDiscount(total float64, initialTotal float64, discount *models.PercentageDiscount) float64 {
	if discount != nil {
		total -= initialTotal * (discount.Percentage / 100)
	}
	if total < 0 {
		total = 0
	}
	return total
}

func applyCategoryDiscount(total float64, cartItems []models.CartItem, discount *models.PercentageCategoryDiscount) float64 {
	if discount != nil {
		for _, cartItem := range cartItems {
			if cartItem.Item.ItemCategoryID == discount.ItemCategoryID {
				total -= cartItem.Item.Price * float64(cartItem.Quantity) * (discount.Percentage / 100)
			}
		}
	}
	if total < 0 {
		total = 0
	}
	return total
}

func calculateMaxPointsUsed(total float64, discount *models.PointDiscount) int {
	if discount == nil {
		return 0
	}
	pointsDiscountAmount := float64(discount.Point)
	maxPointsAllowed := total * 0.2
	if pointsDiscountAmount > maxPointsAllowed {
		pointsDiscountAmount = maxPointsAllowed
	}
	return int(pointsDiscountAmount)
}

func applyPointsDiscount(total float64, discount *models.PointDiscount) (float64, int) {
	maxPointsUsed := calculateMaxPointsUsed(total, discount)
	if discount != nil {
		total -= float64(maxPointsUsed)
	}
	if total < 0 {
		total = 0
	}
	return total, maxPointsUsed
}

func applySeasonalDiscount(total float64, discount *models.SeasonalDiscount) float64 {
	if discount != nil {
		discountAmount := float64(int(total/discount.EveryXTHB)) * discount.DiscountYTHB
		total -= discountAmount
	}
	if total < 0 {
		total = 0
	}
	return total
}
