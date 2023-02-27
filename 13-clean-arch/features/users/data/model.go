package data

import (
	_bookModel "be15/clean/features/books/data"
	"be15/clean/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// ID        uint `gorm:"primarykey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt DeletedAt `gorm:"index"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Address  string
	Role     string
	Books    []_bookModel.Book `gorm:"foreignKey:UserID;references:ID"`
}

// get all users
// {
// 	name: ...
// 	email: ...
// 	books: [
// 		{
// 			title: ..
// 		}
// 	]
// }

// mengubah dari struct core ke struct model
func CoreToModel(dataCore users.Core) User {
	return User{
		Name:     dataCore.Name,
		Email:    dataCore.Email,
		Password: dataCore.Password,
		Address:  dataCore.Address,
		Role:     dataCore.Role,
	}
}

//mengubah dari struct model ke struct core
func ModelToCore(dataModel User) users.Core {
	return users.Core{
		Id:        dataModel.ID,
		Name:      dataModel.Name,
		Email:     dataModel.Email,
		Password:  dataModel.Password,
		Address:   dataModel.Address,
		Role:      dataModel.Role,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}

// mengubah dari list struct model ke struct core
func ListModelToCore(dataModel []User) []users.Core {
	var dataCore []users.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, ModelToCore(v))
	}
	return dataCore
}
