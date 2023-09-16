package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/teste", func(c *fiber.Ctx) error {
		response := struct {
			Message string `json:"message"`
		}{
			Message: "ok",
		}
		return c.JSON(response)
	})
	app.Listen(":8080")

}
