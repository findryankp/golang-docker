package data

import (
	books "be15/clean/features/books"
	"be15/clean/features/users"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string
	UserID      uint
	Author      string
	Publisher   string
	PublishYear time.Time
	User        User
}

type User struct {
	gorm.Model
	// ID        uint `gorm:"primarykey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt DeletedAt `gorm:"index"`
	Name  string
	Email string `gorm:"unique"`
}

// get all books
// [
// 	{
// 		title: ....
// 		author: ...
// 		user_id: ...
// 		user: {
// 			name: ....
// 			email:....
// 		}
// 	},
// 	{
// 		title: ....
// 		author: ...
// 		user_id: ...
// 		user: {
// 			name: ....
// 			email:....
// 		}
// 	}
// ]

// mengubah dari struct core ke struct model
func CoreToModel(dataCore books.Core) Book {
	return Book{
		Title:       dataCore.Title,
		UserID:      dataCore.UserID,
		Author:      dataCore.Author,
		Publisher:   dataCore.Publisher,
		PublishYear: dataCore.PublishYear,
	}
}

//mengubah dari struct model ke struct core
func ModelToCore(dataModel Book) books.Core {
	return books.Core{
		Id:          dataModel.ID,
		Title:       dataModel.Title,
		UserID:      dataModel.UserID,
		Author:      dataModel.Author,
		Publisher:   dataModel.Publisher,
		PublishYear: dataModel.PublishYear,
		User: users.Core{
			Name:  dataModel.User.Name,
			Email: dataModel.User.Email,
		},
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}

// mengubah dari list struct model ke struct core
func ListModelToCore(dataModel []Book) []books.Core {
	var dataCore []books.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, ModelToCore(v))
	}
	return dataCore
}
