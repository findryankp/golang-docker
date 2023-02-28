package router

import (
	_bookData "be15/clean/features/books/data"
	_bookHandler "be15/clean/features/books/delivery"
	_bookService "be15/clean/features/books/service"
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

	bookData := _bookData.New(db)
	bookService := _bookService.New(bookData)
	bookHandlerAPI := _bookHandler.New(bookService)

	e.POST("/login", userHandlerAPI.Login)
	e.GET("/users", userHandlerAPI.GetAllUser)
	e.POST("/users", userHandlerAPI.Register)

	e.POST("/books", bookHandlerAPI.PostBook)
	e.GET("/books", bookHandlerAPI.GetAllBook)
	e.GET("/users/:id/books", bookHandlerAPI.GetBooksByUserId)
}
