package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// go get -u github.com/labstack/echo/v4

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Article struct {
	Id          int    `json:"id" form:"id"`
	Title       string `json:"title" form:"title"`
	Writer      string `json:"writer" form:"writer"`
	Description string `json:"description" form:"description"`
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

	// query param
	e.GET("/users", GetUsersController)
	// param
	e.GET("/users/:id", GetUserByIdController)
	// form value
	e.POST("/users", PostUserController)
	e.PUT("/users", PostUserController)
	// bind
	e.POST("/articles", PostArticleController)
	e.Logger.Fatal(e.Start(":80"))
}

func GetArticles(c echo.Context) error {
	// return c.String(http.StatusOK, "Hello, World! GET")
	return c.JSON(http.StatusOK, map[string]any{
		"message": "Data Berhasil dibaca",
		"status":  true,
	})
}

//query param
func GetUsersController(c echo.Context) error {
	nameQuery := c.QueryParam("name")
	pageQuery := c.QueryParam("page")

	var data = []User{
		{1, "Alta", "alta@alterra.id"},
		{2, "Alta", "alta@alterra.id"},
		{3, "Alta", "alta@alterra.id"},
		{4, "Alta", "alta@alterra.id"},
	}
	// return c.JSON(http.StatusOK, data)
	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
		"query":   nameQuery,
		"page":    pageQuery,
		"data":    data,
	})
}

//param
func GetUserByIdController(c echo.Context) error {
	id := c.Param("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "id harus angka",
		})
	}

	var data = User{
		Id:    idConv,
		Name:  "Alta",
		Email: "alta@alterra.id",
	}
	return c.JSON(http.StatusOK, data)
}

// request body dengan pendekatan form value
func PostUserController(c echo.Context) error {
	nameForm := c.FormValue("name")
	emailForm := c.FormValue("email")

	var data = User{
		Id:    1,
		Name:  nameForm,
		Email: emailForm,
	}
	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
		"data":    data,
	})
}

// request body dengan pendekatan bind
func PostArticleController(c echo.Context) error {
	// binding data
	articleInput := Article{}
	errBind := c.Bind(&articleInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}
	if articleInput.Title == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "title tidak boleh kosong",
		})
	}
	//  c.Request().FormFile("file")
	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
		"data":    articleInput,
	})
}
