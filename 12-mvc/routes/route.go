package routes

import (
	"be15/mvc/controllers"
	"be15/mvc/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	// route
	e.GET("/users", controllers.GetAllUserController, middlewares.JWTMiddleware())
	e.POST("/users", controllers.CreateUserController)

	e.POST("/login", controllers.LoginController)
	return e
}
