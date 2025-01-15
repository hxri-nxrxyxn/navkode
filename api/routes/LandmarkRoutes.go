package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/navkode/controller"
	"gorm.io/gorm"
)

func LandmarkRoutes(db *gorm.DB, app *fiber.App) {

	api := app.Group("/api/v1")
	api.Post("/landmark", controller.CreateLandmark(db))
	api.Get("/landmark/:id", controller.GetLandmark(db))

	api.Get("/landmarkCat/:category", controller.CategoryLandmark(db))
}
