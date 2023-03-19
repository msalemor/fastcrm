package common

import (
	"fastcrm/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func MapLeads(app *fiber.App, db *gorm.DB) {

	// GET all leads
	app.Get("/api/leads", func(c *fiber.Ctx) error {
		var leads []models.Lead
		db.Find(&leads)
		return c.Status(fiber.StatusCreated).JSON(leads)
	})

	// GET a specific lead
	app.Get("/api/leads/:id", func(c *fiber.Ctx) error {
		var lead models.Lead
		db.First(&lead, c.Params("id"))
		if lead.ID == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Lead not found"})
		}
		return c.JSON(lead)
	})

	// POST a new lead
	app.Post("/api/leads", func(c *fiber.Ctx) error {
		lead := new(models.Lead)
		if err := c.BodyParser(lead); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Bad Request"})
		}
		db.Create(&lead)
		return c.JSON(lead)
	})

	// PUT update a specific lead
	app.Put("/api/leads/:id", func(c *fiber.Ctx) error {
		var lead models.Lead
		db.First(&lead, c.Params("id"))
		if lead.ID == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Lead not found"})
		}
		updatedLead := new(models.Lead)
		if err := c.BodyParser(updatedLead); err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"message": "Bad Request"})
		}
		db.Model(&lead).Updates(updatedLead)
		return c.JSON(updatedLead)
	})

	app.Delete("/api/leads/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		// Check if the lead exists
		lead := &models.Lead{}
		result := db.First(&lead, id)
		if result.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Lead not found",
			})
		}

		// Delete the lead from the database
		if err := db.Delete(&lead).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete lead",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Lead deleted successfully",
		})
	})

}
