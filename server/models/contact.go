package models

import "gorm.io/gorm"

// Contact represents a person associated with an account
type Contact struct {
	gorm.Model
	FirstName    string  `gorm:"size:50;not null" json:"first"`
	LastName     string  `gorm:"size:50;not null" json:"last"`
	Email        string  `gorm:"size:100;uniqueIndex" json:"email"`
	Phone        string  `gorm:"size:20;not null" json:"phone"`
	MobilePhone  string  `gorm:"size:20;not null" json:"mphone"`
	Company      string  `gorm:"size:100" json:"company"`
	Title        string  `gorm:"size:100" json:"title"`
	Organization string  `gorm:"size:100" json:"org"`
	Address      string  `gorm:"size:100" json:"address"`
	City         string  `gorm:"size:50" json:"city"`
	State        string  `gorm:"size:10" json:"state"`
	Postal       string  `gorm:"size:20" json:"postal"`
	Country      string  `gorm:"size:100" json:"country"`
	AccountID    uint    `gorm:"not null" json:"account_id"`
	Account      Account `json:"-"`
	Notes        string  `gorm:"Type:longtext" json:"notes"`
}
