package models

import "gorm.io/gorm"

// Account represents a customer account
type Account struct {
	gorm.Model
	Name           string    `gorm:"not null" json:"name"`
	Website        string    `json:"website"`
	Industry       string    `json:"industry"`
	Address        string    `gorm:"size:100" json:"address"`
	City           string    `gorm:"size:50" json:"city"`
	State          string    `gorm:"size:10" json:"state"`
	Postal         string    `gorm:"size:20" json:"postal"`
	Country        string    `gorm:"size:100" json:"country"`
	TotalEmployees int       `json:"employees"`
	TotalRevenue   float64   `json:"revenue"`
	Contacts       []Contact `json:"contacts,omitempty"`
	Notes          string    `gorm:"Type:longtext" json:"notes"`
}
