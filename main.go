package main

import (
	"log"

	"github.com/faelr10/api-authorization-go/internal/infra/database"
	"github.com/faelr10/api-authorization-go/internal/routes"
	database_pkg "github.com/faelr10/api-authorization-go/pkg/infra"
	"github.com/gofiber/fiber/v2"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	db, err := database.CockroachDB().SetupDatabase(&database_pkg.Config{
		User:     "admin",
		Password: "admin",
		Host:     "localhost",
		Port:     5434,
		Database: "api-go-db",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("Erro ao configurar banco de dados: %v", err)
	}

	app := fiber.New()
	routes.SetupRoutes(app, db)
	app.Listen(":8080")
}
