package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"be15/clean/app/config"
)

func InitDBPosgres(cfg config.AppConfig) *gorm.DB {
	// declare struct config & variable connectionString
	// connectionString := "root:qwerty123@tcp(127.0.0.1:3306)/db_book?charset=utf8&parseTime=True&loc=Local"
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)

	fmt.Println("connect:", connectionString)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Println("error connect to DB:", err.Error())
		return nil
	}

	return db
}

func InitialMigrationPosgres(db *gorm.DB) {
	// db.AutoMigrate(&User{})
	//tambahkan auto migrate untuk book
}
