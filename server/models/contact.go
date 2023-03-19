package models

import "gorm.io/gorm"

// Create a Contact struct
type Contact struct {
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
