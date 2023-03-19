package models

import "gorm.io/gorm"

// Create a Account struct
type Account struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Company   string
	JobTitle  string
	Industry  string
	Location  string
	Source    string
	Notes     string
}
