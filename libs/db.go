package libs

import (
	"os"

	"github.com/joho/godotenv"
)

var UserDBConfig DatabaseConfig
var TodoDBConfig DatabaseConfig

func init() {
	godotenv.Load()

	UserDBConfig = DatabaseConfig{
		Host:     os.Getenv("USER_DB_HOST"),
		Port:     os.Getenv("USER_DB_PORT"),
		Username: os.Getenv("USER_DB_USERNAME"),
		Password: os.Getenv("USER_DB_PASSWORD"),
		Database: os.Getenv("USER_DB_NAME"),
	}

	// TodoDBConfig = DatabaseConfig{
	// 	Host:     os.Getenv("TODO_DB_HOST"),
	// 	Port:     os.Getenv("TODO_DB_PORT"),
	// 	Username: os.Getenv("TODO_DB_USERNAME"),
	// 	Password: os.Getenv("TODO_DB_PASSWORD"),
	// 	Database: os.Getenv("TODO_DB_NAME"),
	// }
}
