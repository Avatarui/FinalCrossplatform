package models

// "gorm.io/gorm"

type customer struct {
	ID        int
	Email     string
	Password  string
	FirstName string
	LastName  string
}
