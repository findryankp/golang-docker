package books

import (
	"be15/clean/features/users"
	"time"
)

type Core struct {
	Id          uint
	Title       string
	UserID      uint
	Author      string
	Publisher   string
	PublishYear time.Time
	User        users.Core
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type BookDataInterface interface {
	SelectAll() ([]Core, error)
	SelectBooksByUserId(userid int) ([]Core, error)
	Insert(input Core) error
	Delete(id int) error
}

type BookServiceInterface interface {
	GetAll() ([]Core, error)
	GetBooksByUserId(userid int) ([]Core, error)
	Create(input Core) error
}
