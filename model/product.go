package models

import "time"

type Product struct {
	ProductID     uint   `gorm:"primaryKey"`
	ProductName   string `gorm:"not null"`
	Description   string
	Price         float64 `gorm:"type:decimal(10,2);not null"`
	StockQuantity int     `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
