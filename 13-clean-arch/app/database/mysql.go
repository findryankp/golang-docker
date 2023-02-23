package database

func InitDB() {

	// declare struct config & variable connectionString
	connectionString := "root:qwerty123@tcp(127.0.0.1:3306)/db_book?charset=utf8&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}
