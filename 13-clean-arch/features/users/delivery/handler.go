package delivery

import (
	"be15/clean/features/users"
	"be15/clean/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService users.UserServiceInterface
}

func New(srv users.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: srv,
	}
}

func (delivery *UserHandler) GetAllUser(c echo.Context) error {
	data, err := delivery.userService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error read data"))
	}
	dataResponse := listCoreToResponse(data)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil membaca data user", dataResponse))
}
