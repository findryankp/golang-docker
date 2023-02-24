package users

import "time"

type Core struct {
	Id        uint
	Name      string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string
	Address   string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserDataInterface interface {
	SelectAll() ([]Core, error)
	Insert(input Core) error
	Delete(id int) error
}

type UserServiceInterface interface {
	GetAll() ([]Core, error)
	Create(input Core) error
}
