package data

import (
	"be15/clean/features/users"
	"errors"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

// Delete implements users.UserDataInterface
func (repo *userQuery) Delete(id int) error {
	panic("unimplemented")
}

// Create implements users.UserData
func (repo *userQuery) Insert(input users.Core) error {
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

// GetAll implements users.UserData
func (repo *userQuery) SelectAll() ([]users.Core, error) {
	var usersModel []User
	tx := repo.db.Find(&usersModel)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// mapping dari struct model ke core/entities
	userCoreAll := ListModelToCore(usersModel)

	return userCoreAll, nil
}

func New(db *gorm.DB) users.UserDataInterface {
	return &userQuery{
		db: db,
	}
}
