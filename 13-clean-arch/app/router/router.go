package router

import (
	_userData "be15/clean/features/users/data"
	_userHandler "be15/clean/features/users/delivery"
	_userService "be15/clean/features/users/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	userService := _userService.New(userData)
	userHandlerAPI := _userHandler.New(userService)

	e.GET("/users", userHandlerAPI.GetAllUser)
	e.POST("/users", userHandlerAPI.Register)
}
