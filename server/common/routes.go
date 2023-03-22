package common

import (
	"fastcrm/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

func MapRoutes(app *fiber.App, db *gorm.DB) {

	// Default config
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New())

	app.Static("/", "./www")

	controller.RegisterLeads(app, db)

}
