package controllers

import (
	"FinalCrossplatform/database"
	models "FinalCrossplatform/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AddItemRequest struct {
	CartID    uint `json:"cart_id" binding:"required"`
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,gt=0"` // Quantity must be greater than 0
}

// AddItemResponse defines the structure of the response
type AddItemResponse struct {
	CartItemID uint `json:"cart_item_id"`
	CartID     uint `json:"cart_id"`
	ProductID  uint `json:"product_id"`
	Quantity   int  `json:"quantity"`
}

func AddItem(c *gin.Context) {
	// Define the request struct
	var req AddItemRequest

	// Bind JSON request body to the struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request data",
			"error":   err.Error(),
		})
		return
	}

	// Check if the cart exists
	var cart models.Cart
	if err := database.DB.First(&cart, req.CartID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Cart not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to check cart",
			"error":   err.Error(),
		})
		return
	}

	// Check if the product exists and has enough stock
	var product models.Product
	if err := database.DB.First(&product, req.ProductID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Product not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to check product",
			"error":   err.Error(),
		})
		return
	}

	// Check stock availability
	if product.StockQuantity < req.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Insufficient stock",
			"stock":   product.StockQuantity,
		})
		return
	}

	// Check if the item already exists in the cart
	var existingItem models.CartItem
	if err := database.DB.Where("cart_id = ? AND product_id = ?", req.CartID, req.ProductID).First(&existingItem).Error; err == nil {
		// Item exists, update quantity
		newQuantity := existingItem.Quantity + req.Quantity
		if product.StockQuantity < newQuantity {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Insufficient stock for updated quantity",
				"stock":   product.StockQuantity,
			})
			return
		}

		existingItem.Quantity = newQuantity
		if err := database.DB.Save(&existingItem).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to update cart item",
				"error":   err.Error(),
			})
			return
		}

		// Return success response
		c.JSON(http.StatusOK, gin.H{
			"message": "Item quantity updated successfully",
			"cart_item": AddItemResponse{
				CartItemID: existingItem.CartItemID,
				CartID:     existingItem.CartID,
				ProductID:  existingItem.ProductID,
				Quantity:   existingItem.Quantity,
			},
		})
		return
	}

	// Create new cart item if it doesn't exist
	newItem := models.CartItem{
		CartID:    req.CartID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	if err := database.DB.Create(&newItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to add item to cart",
			"error":   err.Error(),
		})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Item added to cart successfully",
		"cart_item": AddItemResponse{
			CartItemID: newItem.CartItemID,
			CartID:     newItem.CartID,
			ProductID:  newItem.ProductID,
			Quantity:   newItem.Quantity,
		},
	})
}
