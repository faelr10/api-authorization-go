package routes

import (
	user_handler "github.com/faelr10/api-authorization-go/internal/handlers/user"
	user_usecase "github.com/faelr10/api-authorization-go/internal/usecase/user"
	user_handler_pkg "github.com/faelr10/api-authorization-go/pkg/handle"
	"github.com/gofiber/fiber/v2"
)

func routes(app *fiber.App, handler user_handler_pkg.IUserHandlers) {
	app.Post("/user", handler.NewUser)
	//app.Post("/transactionBRCode", handler.GetById)
}

func SetupRoutes(app *fiber.App) {
	user_usecase := user_usecase.NewUserUseCase()
	user_handler := user_handler.NewUserHandler(user_usecase) 
	routes(app, user_handler)
}
