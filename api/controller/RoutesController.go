package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/navkode/models"
	"gorm.io/gorm"
)

func CreateRoutes(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		newRoutes := new(models.Routes)

		err := c.BodyParser(newRoutes)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Error parsing request body",
				"error":   err.Error(),
			})
		}

		err = db.Create(newRoutes).Error
		if err != nil {
			return c.Status(501).JSON(fiber.Map{
				"message": "Error inserting into database",
				"error":   err.Error(),
			})
		}

		return c.Status(201).JSON(fiber.Map{
			"message": "Routes created succerssfully",
			"data":    newRoutes,
		})
	}
}

func UpdateRoutes(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		RoutesModel := new(models.Routes)

		// Parse the body contents into the Routes
		err := c.BodyParser(RoutesModel)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Error parsing request body",
				"error":   err.Error(),
			})
		}

		dbRoutes := new(models.Routes)
		err = db.First(dbRoutes, id).Error
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Error finding Routes",
				"error":   err.Error(),
			})
		}

		db.Save(RoutesModel)

		return c.Status(200).JSON(fiber.Map{
			"message": "Routes Updated successfully",
		})
	}
}
