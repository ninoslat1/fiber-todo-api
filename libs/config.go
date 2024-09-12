package libs

import (
	"github.com/gofiber/contrib/swagger"
)

func GetSwaggerConfig() swagger.Config {
	return swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "swagger",
		Title:    "Swagger API Docs",
	}
}
