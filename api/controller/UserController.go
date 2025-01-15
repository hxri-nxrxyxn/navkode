package controller

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/navkode/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		userModel := new(models.User)

		err := c.BodyParser(userModel)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Error parsing request body",
				"error":   err.Error(),
			})
		}

		if userModel.Email == "" || userModel.Password == "" {
			return c.Status(422).JSON(fiber.Map{
				"message": "Email and password are required",
			})
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userModel.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Error hashing password",
				"error":   err.Error(),
			})
		}
		userModel.Password = string(hashedPassword)

		err = db.Create(userModel).Error
		if err != nil {

			// Check if the error is a duplicate entry
			if strings.Contains(err.Error(), "(SQLSTATE 23505)") {
				return c.Status(409).JSON(fiber.Map{
					"message": "Email already taken",
				})
			}

			return c.Status(501).JSON(fiber.Map{
				"message": "Error inserting into database",
				"error":   err.Error(),
			})
		}

		return c.Status(201).JSON(fiber.Map{
			"message": "User created successfully",
		})
	}
}


func GetUser(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		userModel := new(models.User)
		err := db.First(userModel, id).Error

		if err != nil {

			if strings.Contains(err.Error(), "record not found") {
				return c.Status(404).JSON(fiber.Map{
					"message": "User not found",
					"error":   err.Error(),
				})
			}

			return c.Status(500).JSON(fiber.Map{
				"message": "Error retrieving user",
				"error":   err.Error(),
			})
		}

		userModel.Password = ""
		return c.Status(200).JSON(fiber.Map{
			"message": "Successfully retrieved user",
			"data":    userModel,
		})
	}
}

func Login(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		userModel := new(models.User)

		err := c.BodyParser(userModel)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Error parsing request body",
				"error":   err.Error(),
			})
		}

		if userModel.Email == "" || userModel.Password == "" {
			return c.Status(422).JSON(fiber.Map{
				"message": "Email and password are required",
			})
		}

		email := userModel.Email
		user := new(models.User)
		err = db.Where("email = ?", email).First(user).Error

		if err != nil {

			if strings.Contains(err.Error(), "record not found") {
				return c.Status(404).JSON(fiber.Map{
					"message": "User not found",
					"error":   err.Error(),
				})
			}
			
			return c.Status(500).JSON(fiber.Map{
				"message": "Error retrieving user",
				"error":   err.Error(),
			})
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userModel.Password))
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Password is wrong",
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "Successfully logged in",
		})
	}
}
