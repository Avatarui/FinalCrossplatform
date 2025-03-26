package models

import "time"

type CartItem struct {
	CartItemID uint      `gorm:"primaryKey" json:"cart_item_id"`
	CartID     uint      `gorm:"not null" json:"cart_id"`
	ProductID  uint      `gorm:"not null" json:"product_id"`
	Quantity   int       `gorm:"not null" json:"quantity"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Product    Product   `gorm:"foreignKey:ProductID" json:"product"` // Relationship with Product
}
