package main

import (
	"fiber-api/handlers"
	"fiber-api/libs"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db, err := libs.ConnectToDatabase()
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	app.Use(cors.New())
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	app.Get("/", handlers.HelloWorldHandler)
	app.Post("/login", handlers.SearchUserHandler)

	log.Fatal(app.Listen(":3000"))
}
