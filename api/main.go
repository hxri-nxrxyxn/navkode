package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/navkode/database"
	"github.com/navkode/models"
)

func main() {

	db, err := database.NewConnection()
	if err != nil {
		log.Fatal("Could not load database")
	}
	
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	models.MigrateLandmark(db)
	

	app.Listen(":8080")
}
