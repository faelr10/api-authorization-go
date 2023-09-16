package user_usecase

import (
	"fmt"

	user_pkg "github.com/faelr10/api-authorization-go/pkg/user"
)

type UserImpl struct{}

func NewUserUseCase() user_pkg.IUserUseCase {
	return &UserImpl{}
}

func (u *UserImpl) NewUser(input *user_pkg.INewUserParams) (*user_pkg.INewUserResponse, error) {
	fmt.Println(input)
	//verificar se já existe email cadastrado

	//verificar se password e confirmPassword conferem

	//fazer hash do password

	//salvar no banco de dados

	//retornar novo usuario sem senha e com id
	return &user_pkg.INewUserResponse{
		Id:    "uuid",
		Nome:  input.Nome,
		Email: input.Email,
	}, nil
}
