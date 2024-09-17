package cmd

import (
	"fiber-api/handlers"
	"fiber-api/libs"
	"log"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func RunServer() {
	userDB, err := libs.ConnectToDatabase(libs.UserDBConfig) // Use libs.userDBConfig directly
	if err != nil {
		panic(err)
	}

	// todoDB, err := libs.ConnectToDatabase(libs.TodoDBConfig) // Use libs.todoDBConfig directly
	// if err != nil {
	// 	panic(err)
	// }

	app := fiber.New()
	app.Use(cors.New())
	app.Use(swagger.New(libs.GetSwaggerConfig()))

	app.Post("/login", func(c *fiber.Ctx) error {
		c.Locals("db", userDB)
		return handlers.SearchUserHandler(c)
	})

	log.Fatal(app.Listen(":3000"))
}
