package main

import (
	"belobetty-queue-manager/infra/http/fiber/git"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env:", err)
	}

	app := fiber.New()

	version := os.Getenv("VERSION")
	if version == "" {
		version = "/v1"
	}

	v1 := app.Group(version)

	git.SetRoutes(v1)

	port := os.Getenv("PORT")
	if version == "" {
		version = ":9001"
	}

	app.Listen(port)
}
