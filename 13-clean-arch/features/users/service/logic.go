package service

import (
	"be15/clean/features/users"
	"errors"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData users.UserDataInterface
	validate *validator.Validate
}

// Login implements users.UserServiceInterface
func (service *userService) Login(email string, password string) (users.Core, string, error) {
	if email == "" || password == "" {
		return users.Core{}, "", errors.New("email dan password harus diisi")
	}
	data, token, err := service.userData.Login(email, password)
	return data, token, err
}

// Create implements users.UserServiceInterface
func (service *userService) Create(input users.Core) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}
	errInsert := service.userData.Insert(input)
	if errInsert != nil {
		return errInsert
	}
	return nil
}

// GetAll implements users.UserServiceInterface
func (service *userService) GetAll(page int) ([]users.Core, error) {
	/*
		limit = 10
		page = 1 --> data ke 1 - 10 --> offset = 0
		page = 2 --> data ke 11 - 20 --> offset = 10
		page = 3 --> data ke 21 - 30 --> offset = 20
	*/
	var limit = 2
	var offset = (page - 1) * limit

	data, err := service.userData.SelectAll(limit, offset)
	return data, err
}

func New(repo users.UserDataInterface) users.UserServiceInterface {
	return &userService{
		userData: repo,
		validate: validator.New(),
	}
}
