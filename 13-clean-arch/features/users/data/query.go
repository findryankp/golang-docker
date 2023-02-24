package data

import (
	"be15/clean/features/users"

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
	panic("unimplemented")
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
