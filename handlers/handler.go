package handlers

import (
	"encoding/base64"
	"fiber-api/libs"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// Handle the route with a parameter (`/:params`)
func SearchUserHandler(c *fiber.Ctx) error {

	var user libs.UserDB
	db := c.Locals("db").(*gorm.DB)
	username := c.FormValue("UserCode")
	password := c.FormValue("Password")
	encodedPasswordLike := base64.StdEncoding.EncodeToString([]byte(strings.Join(strings.Split(password, ""), "\u0000")))[:len(base64.StdEncoding.EncodeToString([]byte(strings.Join(strings.Split(password, ""), "\u0000"))))-2] + "%"
	result := db.Table("myuser").Where("UserCode = ? AND Password LIKE ?", username, encodedPasswordLike).First(&user)
	claims := jwt.MapClaims{
		"name": username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("APP_TOKEN")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

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
		"token":   t,
	})
}
