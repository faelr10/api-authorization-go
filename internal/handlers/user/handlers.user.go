package user_handler

import (
	"net/http"
	"github.com/faelr10/api-authorization-go/internal/handlers"
	user_pkg "github.com/faelr10/api-authorization-go/pkg/user"
	"github.com/gofiber/fiber/v2"
)

type UserHandlerImpl struct {
	NewUserUseCase user_pkg.IUserUseCase
}

func NewUserHandler(newUserCase user_pkg.IUserUseCase) *UserHandlerImpl {
	return &UserHandlerImpl{
		NewUserUseCase: newUserCase,
	}
}

func (u *UserHandlerImpl) NewUser(ctx *fiber.Ctx) error {
	var input user_pkg.INewUserParams
	if err := ctx.BodyParser(&input); err != nil {
		return handlers.NewErrorReponse(ctx, http.StatusBadRequest, err.Error())
	}

	success, err := u.NewUserUseCase.NewUser(&input)
	if err != nil {
		return handlers.NewErrorReponse(ctx, http.StatusBadRequest, err.Error())
	}
	return handlers.NewSuccessResponse(ctx, http.StatusCreated, "success", success)
}
