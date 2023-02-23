package controllers

import (
	"be15/mvc/entities"
	"be15/mvc/helper"
	"be15/mvc/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	loginInput := entities.AuthRequest{}
	errBind := c.Bind(&loginInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	dataUser, token, err := repositories.Login(loginInput.Email, loginInput.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("login failed"))
	}

	dataResponse := map[string]any{
		"token": token,
		"name":  dataUser.Name,
		"email": dataUser.Email,
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success login", dataResponse))
}
