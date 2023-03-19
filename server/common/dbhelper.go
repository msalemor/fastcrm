package common

import "fastcrm/models"

func Migrate() {
	GetDb().AutoMigrate(&models.Lead{})
}
