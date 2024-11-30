package user_pkg

import "github.com/faelr10/api-authorization-go/internal/infra/database/models"

type IUserUseCase interface {
	NewUser(*INewUserParams) (*INewUserResponse, error)
	GetUserById(string) (*INewUserResponse, error)
}

type IUserRepository interface {
	InsertUser(*NewUserRepositoryParams) (*models.User, error)
	GetUserById(map[string]interface{}) (*models.User, error)
}

//_________________________________________________________________________________
//USE CASE
type INewUserParams struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}
type INewUserResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
//_________________________________________________________________________________
//REPOSITORY
type NewUserRepositoryParams struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
}