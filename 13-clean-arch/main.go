package main

import (
	"be15/clean/app/config"
	"be15/clean/app/database"
	_userData "be15/clean/features/users/data"
	_userHandler "be15/clean/features/users/delivery"
	_userService "be15/clean/features/users/service"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.InitConfig()
	db := database.InitDBMySql(*cfg)
	// db := database.InitDBPosgres(*cfg)

	e := echo.New()

	userData := _userData.New(db)
	userService := _userService.New(userData)
	userHandlerAPI := _userHandler.New(userService)

	e.GET("/users", userHandlerAPI.GetAllUser)

	e.Logger.Fatal(e.Start(":80"))
}
