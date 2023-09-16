package main

import (
	"github.com/faelr10/api-authorization-go/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.SetupRoutes(app)
	app.Listen(":8080")
}
