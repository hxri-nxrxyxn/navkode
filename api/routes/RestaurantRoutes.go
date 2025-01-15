package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/navkode/controller"
	"gorm.io/gorm"
)

func RestaurantRoutes(db *gorm.DB, app *fiber.App) {

	api := app.Group("/api/v1")
	api.Post("/restaurant", controller.CreateRestaurant(db))
}
