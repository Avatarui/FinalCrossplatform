package models

import "time"

type Cart struct {
	CartID     uint       `gorm:"primaryKey" json:"cart_id"`
	CustomerID uint       `gorm:"not null" json:"customer_id"`
	CartName   string     `json:"cart_name"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	Items      []CartItem `gorm:"foreignKey:CartID" json:"items"` // Define the relationship
}
