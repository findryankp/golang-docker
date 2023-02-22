package repositories

import (
	"be15/mvc/config"
	"be15/mvc/entities"
	"be15/mvc/models"
)

func GetAllUser() ([]entities.UserCore, error) {
	var users []models.User
	tx := config.DB.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// mapping dari struct model ke core/entities
	var userCoreAll []entities.UserCore
	for _, v := range users {
		var userCore = entities.UserCore{
			Id:        v.ID,
			Name:      v.Name,
			Email:     v.Email,
			Password:  v.Password,
			Address:   v.Address,
			Role:      v.Role,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		userCoreAll = append(userCoreAll, userCore)
	}

	return userCoreAll, nil
}
