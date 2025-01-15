package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/navkode/controller"
	"gorm.io/gorm"
)

func RoutesRoutes(db *gorm.DB, app *fiber.App) {

	api := app.Group("/api/v1")
	api.Post("/routes", controller.CreateRoutes(db))
}
