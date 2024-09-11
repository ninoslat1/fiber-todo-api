package handlers

import (
	"encoding/base64"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Handle the root route (`/`)
func HelloWorldHandler(c *fiber.Ctx) error {
	return c.SendString("Hello from Home with Fiber!")
}

// Handle the route with a parameter (`/:params`)
func SearchUserHandler(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	username := c.FormValue("UserCode")
	password := c.FormValue("Password")
	encodedPassword := base64.StdEncoding.EncodeToString([]byte(password))
	result := db.Table("myuser").Where("UserCode = ? AND Password = ?", username, encodedPassword)

	if db == nil {
		panic("Database connection is nil")
	}

	if result.Error != nil {
		log.Printf("Error finding resident: %v", result.Error)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}

	return result
}
