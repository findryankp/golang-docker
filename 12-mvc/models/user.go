package models

import (
	"be15/mvc/entities"

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
}

func UserCoreToModel(dataCore entities.UserCore) User {
	return User{
		Name:     dataCore.Name,
		Email:    dataCore.Email,
		Password: dataCore.Password,
		Address:  dataCore.Address,
		Role:     dataCore.Role,
	}
}

func ModelToUserCore(dataModel User) entities.UserCore {
	return entities.UserCore{
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

func ListModelToUserCore(dataModel []User) []entities.UserCore {
	var dataCore []entities.UserCore
	for _, v := range dataModel {
		dataCore = append(dataCore, ModelToUserCore(v))
	}
	return dataCore
}
