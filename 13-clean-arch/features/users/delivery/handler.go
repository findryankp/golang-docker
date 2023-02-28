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

func (delivery *UserHandler) Register(c echo.Context) error {
	userInput := UserRequest{}
	errBind := c.Bind(&userInput) // menangkap inputan dari user melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}
	dataCore := requestToCore(userInput)
	err := delivery.userService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error: "+err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("berhasil insert data user"))

}

func (delivery *UserHandler) Login(c echo.Context) error {
	loginInput := LoginRequest{}
	errBind := c.Bind(&loginInput) // menangkap inputan dari user melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}
	data, token, err := delivery.userService.Login(loginInput.Email, loginInput.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error: "+err.Error()))
	}

	dataResponse := map[string]any{
		"name":  data.Name,
		"token": token,
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil login", dataResponse))
}
