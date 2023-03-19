package common

import "fastcrm/models"

func GormMigrate() {
	GetDb().AutoMigrate(&models.Lead{}, &models.Contact{}, &models.Account{})
}
