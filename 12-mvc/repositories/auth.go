package repositories

import (
	"be15/mvc/config"
	"be15/mvc/entities"
	"be15/mvc/middlewares"
	"be15/mvc/models"
	"errors"
)

func Login(email, password string) (entities.UserCore, string, error) {
	var user models.User
	tx := config.DB.Where("email = ? AND password = ?", email, password).First(&user)
	if tx.Error != nil {
		return entities.UserCore{}, "", tx.Error
	}
	if tx.RowsAffected == 0 {
		return entities.UserCore{}, "", errors.New("login failed")

	}

	token, errToken := middlewares.CreateToken(int(user.ID))
	if errToken != nil {
		return entities.UserCore{}, "", errToken
	}

	dataCore := entities.UserCore{
		Id:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Role:    user.Role,
	}
	return dataCore, token, nil
}
