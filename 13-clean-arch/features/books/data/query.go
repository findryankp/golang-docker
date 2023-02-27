package data

import (
	"be15/clean/features/books"
	"errors"

	"gorm.io/gorm"
)

type bookQuery struct {
	db *gorm.DB
}

// SelectBooksByUserId implements books.BookDataInterface
func (repo *bookQuery) SelectBooksByUserId(userid int) ([]books.Core, error) {
	var booksModel []Book
	tx := repo.db.Where("user_id = ?", userid).Find(&booksModel)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// mapping dari struct model ke core/entities
	bookCoreAll := ListModelToCore(booksModel)

	return bookCoreAll, nil
}

// Delete implements books.BookDataInterface
func (*bookQuery) Delete(id int) error {
	panic("unimplemented")
}

// Insert implements books.BookDataInterface
func (repo *bookQuery) Insert(input books.Core) error {
	dataModel := CoreToModel(input)
	tx := repo.db.Create(&dataModel) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert error, row affected = 0")
	}
	return nil
}

// SelectAll implements books.BookDataInterface
func (repo *bookQuery) SelectAll() ([]books.Core, error) {
	var booksModel []Book
	tx := repo.db.Preload("User").Find(&booksModel)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// mapping dari struct model ke core/entities
	bookCoreAll := ListModelToCore(booksModel)

	return bookCoreAll, nil
}

func New(db *gorm.DB) books.BookDataInterface {
	return &bookQuery{
		db: db,
	}
}
