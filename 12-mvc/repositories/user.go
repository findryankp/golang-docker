package repositories

import (
	"be15/mvc/config"
	"be15/mvc/entities"
	"be15/mvc/models"
	"errors"
)

func GetAllUser() ([]entities.UserCore, error) {
	var users []models.User
	tx := config.DB.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// mapping dari struct model ke core/entities
	// var userCoreAll []entities.UserCore
	// for _, v := range users {
	// 	var userCore = entities.UserCore{
	// 		Id:        v.ID,
	// 		Name:      v.Name,
	// 		Email:     v.Email,
	// 		Password:  v.Password,
	// 		Address:   v.Address,
	// 		Role:      v.Role,
	// 		CreatedAt: v.CreatedAt,
	// 		UpdatedAt: v.UpdatedAt,
	// 	}
	// 	userCoreAll = append(userCoreAll, userCore)
	// }
	userCoreAll := models.ListModelToUserCore(users)

	return userCoreAll, nil
}

func CreateUser(input entities.UserCore) error {
	// proses mapping dari entities core ke gorm model
	// dataGorm := models.User{
	// 	Name:     input.Name,
	// 	Email:    input.Email,
	// 	Password: input.Password,
	// 	Address:  input.Address,
	// 	Role:     input.Role,
	// }
	dataGorm := models.UserCoreToModel(input)
	tx := config.DB.Create(&dataGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert error, row affected = 0")
	}
	return nil
}
