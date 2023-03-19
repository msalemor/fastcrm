package models

import "gorm.io/gorm"

// Create a Lead struct
type Lead struct {
	gorm.Model
	FirstName    string
	LastName     string
	Email        string
	Phone        string
	Company      string
	JobTitle     string
	Organization string
	Industry     string
	Address      string
	Source       string
	Notes        string
}
