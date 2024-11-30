package user_handler

import (
	"net/http"

	"github.com/faelr10/api-authorization-go/internal/handlers"
	user_handler_pkg "github.com/faelr10/api-authorization-go/pkg/interfaces/handle"
	user_pkg "github.com/faelr10/api-authorization-go/pkg/interfaces/user"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

//definindo a struct do handler
type UserHandlerImpl struct {
	UserUseCase user_pkg.IUserUseCase
}

//função que cria um novo handler
func UserHandler(userCase user_pkg.IUserUseCase)  user_handler_pkg.IUserHandlers {
	return &UserHandlerImpl{
		UserUseCase: userCase,
	}
}

func (u *UserHandlerImpl) NewUser(ctx *fiber.Ctx) error {

	//Fazendo parse do body para struct desejada
	var input user_pkg.INewUserParams
	if err := ctx.BodyParser(&input); err != nil {
		return handlers.NewErrorReponse(ctx, http.StatusBadRequest, err.Error())
	}

	// Validando campos
	validate := validator.New()
	if err := validate.Struct(&input); err != nil {
		return handlers.NewErrorReponse(ctx, http.StatusBadRequest, "Validation failed: "+err.Error())
	}

	success, err := u.UserUseCase.NewUser(&input)
	if err != nil {
		return handlers.NewErrorReponse(ctx, http.StatusBadRequest, err.Error())
	}
	return handlers.NewSuccessResponse(ctx, http.StatusCreated, "success", success)
}

func (u *UserHandlerImpl) GetUserById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	success, err := u.UserUseCase.GetUserById(id)
	if err != nil {
		return handlers.NewErrorReponse(ctx, http.StatusBadRequest, err.Error())
	}
	return handlers.NewSuccessResponse(ctx, http.StatusOK, "success", success)
}