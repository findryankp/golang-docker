package controllers

import (
	"be15/mvc/entities"
	"be15/mvc/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllUserController(c echo.Context) error {
	result, err := repositories.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error read data",
		})
	}
	var userRespose []entities.UserResponse
	for _, v := range result {
		userRespose = append(userRespose, entities.UserResponse{
			Id:      v.Id,
			Name:    v.Name,
			Address: v.Address,
			Role:    v.Role,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "berhasil read data",
		"data":    userRespose,
	})
}
