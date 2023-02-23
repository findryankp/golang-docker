package controllers

import (
	"be15/mvc/entities"
	"be15/mvc/helper"
	"be15/mvc/middlewares"
	"be15/mvc/repositories"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllUserController(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	fmt.Println("idtoken:", idToken)

	result, err := repositories.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error read data"))
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

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil membaca data user", userRespose))
}

func CreateUserController(c echo.Context) error {
	userInput := entities.UserRequest{}
	errBind := c.Bind(&userInput) // menangkap inputan dari user melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	// mapping struct request ke struct entities core
	dataCore := entities.UserCore{
		Name:     userInput.Name,
		Email:    userInput.Email,
		Password: userInput.Password,
		Address:  userInput.Address,
		Role:     userInput.Role,
	}
	err := repositories.CreateUser(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error insert data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("berhasil insert data user"))
}
