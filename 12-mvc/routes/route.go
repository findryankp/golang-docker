package routes

import (
	"be15/mvc/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	// route
	e.GET("/users", controllers.GetAllUserController)
	return e
}
