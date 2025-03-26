package controllers

import (
	"FinalCrossplatform/database"
	models "FinalCrossplatform/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CartResponse struct {
	CartID    uint           `json:"cart_id"`
	CartName  string         `json:"cart_name"`
	CreatedAt string         `json:"created_at"`
	Items     []ItemResponse `json:"items"`
}

// ItemResponse defines the structure of each item in the cart
type ItemResponse struct {
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	TotalPrice  float64 `json:"total_price"`
}

// GetCustomerCarts retrieves all carts for a customer
func GetCustomerCarts(c *gin.Context) {
	// Get customer_id from the URL path
	customerID := c.Param("id")

	// Validate customer_id (ensure it's a valid integer)
	var customer models.Customer
	if err := database.DB.First(&customer, customerID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Customer not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to check customer",
			"error":   err.Error(),
		})
		return
	}

	// Fetch all carts for the customer with their items and products
	var carts []models.Cart
	if err := database.DB.
		Preload("Items.Product"). // Preload CartItems and their associated Products
		Where("customer_id = ?", customerID).
		Find(&carts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch carts",
			"error":   err.Error(),
		})
		return
	}

	// Transform the data into the response format
	var cartResponses []CartResponse
	for _, cart := range carts {
		// Prepare the cart response
		cartResp := CartResponse{
			CartID:    cart.CartID,
			CartName:  cart.CartName,
			CreatedAt: cart.CreatedAt.Format("2006-01-02 15:04:05"), // Format timestamp
			Items:     []ItemResponse{},
		}

		// Add items to the cart response
		for _, item := range cart.Items { // This should now work
			itemResp := ItemResponse{
				ProductID:   item.ProductID,
				ProductName: item.Product.ProductName,
				Description: item.Product.Description,
				Price:       item.Product.Price,
				Quantity:    item.Quantity,
				TotalPrice:  item.Product.Price * float64(item.Quantity),
			}
			cartResp.Items = append(cartResp.Items, itemResp)
		}

		cartResponses = append(cartResponses, cartResp)
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Carts retrieved successfully",
		"carts":   cartResponses,
	})
}
