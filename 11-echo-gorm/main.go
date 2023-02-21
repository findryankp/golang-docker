package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// ID        uint `gorm:"primarykey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt DeletedAt `gorm:"index"`
	Name     string `json:"name" form:"name"`
	Email    string `gorm:"unique" json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
	Role     string `json:"role" form:"role"`
}

//buat struct book
// https://gorm.io/docs/has_many.html

var DB *gorm.DB

// database connection
func InitDB() {

	// declare struct config & variable connectionString
	connectionString := "root:qwerty123@tcp(127.0.0.1:3306)/db_book?charset=utf8&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
	//tambahkan auto migrate untuk book
}

func init() {
	fmt.Println("ini init")
}

func main() {
	fmt.Println("ini main")
	InitDB()
	InitialMigration()

	e := echo.New()
	// route
	e.POST("/users", CreateUserController)
	e.GET("/users", GetAllUserController)
	// PUT
	// e.PUT("/users/:userid", UpdateUserController)
	// DELETE
	// e.DELETE("/users/:id", DeleteUserController)

	// CRUD Book
	//e.POST("/books", CreateBookController)
	// e.GET("/books", GetAllBookController)
	// PUT
	// e.PUT("/books/:id", UpdateBookController)
	// DELETE
	// e.DELETE("/books/:id", DeleteBookController)

	// start server
	e.Logger.Fatal(e.Start(":80"))
}

//proses register
func CreateUserController(c echo.Context) error {
	userInput := User{}
	errBind := c.Bind(&userInput) // menangkap inputan dari user melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	tx := DB.Create(&userInput) // proses query insert
	if tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error insert data",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "berhasil insert data",
	})
}

// proses melihat semua data user
func GetAllUserController(c echo.Context) error {
	var users []User
	tx := DB.Find(&users)
	if tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error read data",
		})
	}
	return c.JSON(http.StatusOK, map[string]any{
		"message": "berhasil read data",
		"data":    users,
	})

}
