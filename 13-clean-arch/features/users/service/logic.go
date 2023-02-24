package service

import (
	"be15/clean/features/users"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData users.UserDataInterface
	validate *validator.Validate
}

// Create implements users.UserServiceInterface
func (service *userService) Create(input users.Core) error {
	panic("unimplemented")
	// if input.Email == "" || input.Password == "" {
	// 	return errors.New("error")
	// }
}

// GetAll implements users.UserServiceInterface
func (service *userService) GetAll() ([]users.Core, error) {
	data, err := service.userData.SelectAll()
	return data, err
}

func New(repo users.UserDataInterface) users.UserServiceInterface {
	return &userService{
		userData: repo,
		validate: validator.New(),
	}
}
