package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// go get -u github.com/labstack/echo/v4

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	e := echo.New()
	// e.GET("/articles", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	e.GET("/articles", GetArticles)

	e.POST("/articles", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! POST")
	})

	e.GET("/users", GetUsersController)
	e.Logger.Fatal(e.Start(":80"))
}

func GetArticles(c echo.Context) error {
	// return c.String(http.StatusOK, "Hello, World! GET")
	return c.JSON(http.StatusOK, map[string]any{
		"message": "Data Berhasil dibaca",
		"status":  true,
	})
}

func GetUsersController(c echo.Context) error {
	var data = User{
		Name:  "Alta",
		Email: "alta@alterra.id",
	}
	return c.JSON(http.StatusOK, data)
}
