package handlers

import (
	"encoding/base64"
	"fiber-api/libs"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Handle the route with a parameter (`/:params`)
func SearchUserHandler(c *fiber.Ctx) error {
	var user libs.UserDB
	db := c.Locals("db").(*gorm.DB)
	username := c.FormValue("UserCode")
	password := c.FormValue("Password")

	// Tambahkan null character di antara setiap karakter password
	nullSeparatedPassword := strings.Join(strings.Split(password, ""), "\u0000")

	// Konversi ke base64
	encodedPassword := base64.StdEncoding.EncodeToString([]byte(nullSeparatedPassword))

	// Hilangkan dua karakter terakhir dan tambahkan wildcard
	encodedPasswordLike := encodedPassword[:len(encodedPassword)-2] + "%"

	result := db.Table("myuser").Where("UserCode = ? AND Password LIKE ?", username, encodedPasswordLike).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "User not found or password incorrect",
			})
		}

		log.Printf("Error finding user: %v", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	// User found, you can add more logic here (e.g., generate token)
	return c.JSON(fiber.Map{
		"message": "Login successful",
		"user":    username,
	})
}
