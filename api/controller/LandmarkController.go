package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/navkode/models"
	"gorm.io/gorm"
)

func CreateLandmark(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		newLandmark := new(models.Landmark)

		err := c.BodyParser(newLandmark)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Error parsing request body",
				"error":   err.Error(),
			})
		}

		err = db.Create(newLandmark).Error
		if err != nil {
			return c.Status(501).JSON(fiber.Map{
				"message": "Error inserting into database",
				"error":   err.Error(),
			})
		}

		return c.Status(201).JSON(fiber.Map{
			"message": "Landmark created succerssfully",
			"data":    newLandmark,
		})
	}
}

func CategoryLandmark(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		category := c.Params("category")

		var landmarks []models.Landmark

		err := db.Where("category = ?", category).Find(&landmarks).Error
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"message": "Error getting landmarks",
				"error":   err.Error(),
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "Successfully retrieved landmarksd",
			"data":    landmarks,
		})
	}
}

func GetLandmark(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		landmark := new(models.Landmark)

		err := db.Find(landmark, id).Error
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"message": "Error getting the landmark",
				"error":   err.Error(),
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "Successfully retrieved landmark",
			"data":    landmark,
		})

	}
}
