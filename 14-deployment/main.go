package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/users", HelloUsers)
	e.Logger.Fatal(e.Start(":80"))
}

func HelloUsers(c echo.Context) error {
	data := map[string]any{
		"name": "budi",
	}
	return c.JSON(http.StatusOK, data)
}
