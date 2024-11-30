package repository

import (

	"github.com/faelr10/api-authorization-go/internal/infra/database/models"
	user_pkg "github.com/faelr10/api-authorization-go/pkg/interfaces/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Database *gorm.DB
}

func UserRepository(database *gorm.DB) user_pkg.IUserRepository {
	return &UserRepositoryImpl{
		Database: database,
	}
}

func (u *UserRepositoryImpl) InsertUser(params *user_pkg.NewUserRepositoryParams) (*models.User, error) {

	user := &models.User{
		ID:    uuid.New().String(),
		Name:  params.Name,
		Email: params.Email,
		Password: params.Password,
	}

	if err := u.Database.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepositoryImpl) GetUserById(where map[string]interface{}) (*models.User, error) {
	user := &models.User{}

	err := u.Database.Where(where).First(user).Error 
	 
	if err != nil {
		return nil, err
	}

	return user, nil
}
