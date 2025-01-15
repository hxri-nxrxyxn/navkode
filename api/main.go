package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/navkode/database"
	"github.com/navkode/models"
	"github.com/navkode/routes"
)

func main() {

	db, err := database.NewConnection()
	if err != nil {
		log.Fatal("Could not load database")
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	models.MigrateUser(db)
	models.MigrateLandmark(db)
	models.MigrateParking(db)

	routes.LandmarkRoutes(db, app)
	routes.ParkingRoutes(db, app)
	routes.UserRoutes(db, app)

	app.Listen(":8080")
}
