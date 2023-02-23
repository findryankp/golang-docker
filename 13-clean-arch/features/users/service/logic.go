package service

import "be15/clean/features/users"

type userService struct {
	userData users.UserDataInterface
}

// Create implements users.UserServiceInterface
func (service *userService) Create(input users.Core) error {
	panic("unimplemented")
}

// GetAll implements users.UserServiceInterface
func (service *userService) GetAll() ([]users.Core, error) {
	panic("unimplemented")
	// service.userData.SelectAll()
	// service.userData.Insert()
	// service.userData.Delete()
}

func New(repo users.UserDataInterface) users.UserServiceInterface {
	return &userService{
		userData: repo,
	}
}
