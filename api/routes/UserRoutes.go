package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/navkode/controller"
	"gorm.io/gorm"
)

func UserRoutes(db *gorm.DB, app *fiber.App) {
	api := app.Group("/api/v1")
	api.Post("/user", controller.CreateUser(db))
	api.Get("/user/:id", controller.GetUser(db))

	api.Post("/login", controller.Login(db))
}
