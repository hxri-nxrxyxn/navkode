package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/navkode/models"
	"gorm.io/gorm"
)

func CreateParking(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		newParking := new(models.Parking)

		err := c.BodyParser(newParking)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Error parsing request body",
				"error":   err.Error(),
			})
		}

		err = db.Create(newParking).Error
		if err != nil {
			return c.Status(501).JSON(fiber.Map{
				"message": "Error inserting into database",
				"error":   err.Error(),
			})
		}

		return c.Status(201).JSON(fiber.Map{
			"message": "Parking created succerssfully",
			"data":    newParking,
		})
	}
}
