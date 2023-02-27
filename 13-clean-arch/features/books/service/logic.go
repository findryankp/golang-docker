package service

import (
	"be15/clean/features/books"

	"github.com/go-playground/validator/v10"
)

type bookService struct {
	bookData books.BookDataInterface
	validate *validator.Validate
}

// GetBooksByUserId implements books.BookServiceInterface
func (service *bookService) GetBooksByUserId(userid int) ([]books.Core, error) {
	data, err := service.bookData.SelectBooksByUserId(userid)
	return data, err
}

// Create implements books.BookServiceInterface
func (service *bookService) Create(input books.Core) error {
	// errValidate := service.validate.Struct(input)
	// if errValidate != nil {
	// 	return errValidate
	// }
	errInsert := service.bookData.Insert(input)
	if errInsert != nil {
		return errInsert
	}
	return nil
}

// GetAll implements books.BookServiceInterface
func (service *bookService) GetAll() ([]books.Core, error) {
	data, err := service.bookData.SelectAll()
	return data, err
}

func New(repo books.BookDataInterface) books.BookServiceInterface {
	return &bookService{
		bookData: repo,
		validate: validator.New(),
	}
}
