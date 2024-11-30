package user_usecase

import (
	"fmt"
	"log"

	user_pkg "github.com/faelr10/api-authorization-go/pkg/interfaces/user"
	"golang.org/x/crypto/bcrypt"
)

type UserImpl struct {
	Repository user_pkg.IUserRepository
}

func UserUseCase(repository user_pkg.IUserRepository) user_pkg.IUserUseCase {
	return &UserImpl{
		Repository: repository,
	}
}

func (u *UserImpl) NewUser(input *user_pkg.INewUserParams) (*user_pkg.INewUserResponse, error) {
	
	//verifyAlreadyExists
	verifyEmail, _ := u.Repository.GetUserById(map[string]interface{}{"email": input.Email})
	if verifyEmail != nil {
		return nil, fmt.Errorf("email já cadastrado")
	}

	//verifyPassword
	if input.Password != input.ConfirmPassword {
		return nil, fmt.Errorf("password e confirmPassword não conferem")
	}

	//hashPassword
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	user, err := u.Repository.InsertUser(&user_pkg.NewUserRepositoryParams{
		Name:  input.Name,
		Email: input.Email,
		Password: string(hash),
	})
	if err != nil {
		return nil, err
	}

	return &user_pkg.INewUserResponse{
		Id:    user.ID,
		Name:  input.Name,
		Email: input.Email,
	}, nil
}

func (u *UserImpl) GetUserById(id string) (*user_pkg.INewUserResponse, error) {

	user, err := u.Repository.GetUserById(map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return &user_pkg.INewUserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
