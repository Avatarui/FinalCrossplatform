package models

type Customer struct {
	CustomerID  uint   `json:"customer_id" gorm:"column:customer_id;primaryKey"`
	FirstName   string `json:"first_name" gorm:"column:first_name"`
	LastName    string `json:"last_name" gorm:"column:last_name"`
	Email       string `json:"email" gorm:"column:email;unique"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number"`
	Address     string `json:"address" gorm:"column:address"`
	Password    string `json:"password" gorm:"column:password"`
	CreatedAt   string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   string `json:"updated_at" gorm:"column:updated_at"`
}

// บอก GORM ให้ใช้ตาราง "customer"
func (Customer) TableName() string {
	return "customer"
}
