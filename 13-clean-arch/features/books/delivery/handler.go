package delivery

import (
	"be15/clean/features/books"
	"net/http"
	"strconv"

	"be15/clean/utils/helper"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	bookService books.BookServiceInterface
}

func New(srv books.BookServiceInterface) *BookHandler {
	return &BookHandler{
		bookService: srv,
	}
}

func (delivery *BookHandler) GetAllBook(c echo.Context) error {
	data, err := delivery.bookService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error read data"))
	}
	dataResponse := listCoreToResponse(data)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil membaca data user", dataResponse))
}

func (delivery *BookHandler) PostBook(c echo.Context) error {
	bookInput := BookRequest{}
	errBind := c.Bind(&bookInput) // menangkap inputan dari user melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}
	dataCore := requestToCore(bookInput)
	err := delivery.bookService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error: "+err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("berhasil insert data user"))
}

func (delivery *BookHandler) GetBooksByUserId(c echo.Context) error {
	idUser := c.Param("id")
	idUserConv, errConv := strconv.Atoi(idUser)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error convert userid"))
	}
	data, err := delivery.bookService.GetBooksByUserId(idUserConv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error read data"))
	}
	dataResponse := listCoreToResponse(data)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil membaca data user", dataResponse))
}
