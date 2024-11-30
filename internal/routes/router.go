package routes

import (

	user_handler "github.com/faelr10/api-authorization-go/internal/handlers/user"
	"github.com/faelr10/api-authorization-go/internal/repository"
	user_usecase "github.com/faelr10/api-authorization-go/internal/usecase/user"
	user_handler_pkg "github.com/faelr10/api-authorization-go/pkg/interfaces/handle"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func routes(app *fiber.App, handler user_handler_pkg.IUserHandlers) {
	app.Post("/user", handler.NewUser)
	app.Get("/user/:id", handler.GetUserById)
}

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	user_repository := repository.UserRepository(db)
	user_usecase := user_usecase.UserUseCase(user_repository)
	user_handler := user_handler.UserHandler(user_usecase) 
	routes(app, user_handler)
}
