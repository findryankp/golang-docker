package main

import (
	"be15/clean/app/config"
	"be15/clean/app/database"
	_userData "be15/clean/features/users/data"
	_userService "be15/clean/features/users/service"
)

func main() {
	cfg := config.InitConfig()
	db := database.InitDBMySql(*cfg)
	// db := database.InitDBPosgres(*cfg)

	userData := _userData.New(db)
	userService := _userService.New(userData)

}
