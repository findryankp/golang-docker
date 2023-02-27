package main

import (
	"be15/clean/app/config"
	"be15/clean/app/database"
	"be15/clean/app/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	db := database.InitDBMySql(*cfg)
	// db := database.InitDBPosgres(*cfg)

	// auto migrate gorm
	database.InitialMigration(db)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	router.InitRouter(db, e)
	// userData := _userData.New(db)
	// userService := _userService.New(userData)
	// userHandlerAPI := _userHandler.New(userService)

	// e.GET("/users", userHandlerAPI.GetAllUser)
	// e.POST("/users", userHandlerAPI.Register)

	e.Logger.Fatal(e.Start(":80"))
}
