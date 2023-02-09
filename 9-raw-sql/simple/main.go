package main

/*
jalankan:
	go mod init namaproject

download mysql db driver:
	go get -u github.com/go-sql-driver/mysql


*/

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Guru struct {
	Id      string
	Nama    string
	Telepon string
	Email   string
}

func main() {
	// <username>:<password>@tcp(<hostname>:<port>)/<db_name>
	var connectionString = "root:qwerty123@tcp(127.0.0.1:3306)/db_be15"
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("error open connection", err.Error())
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	errPing := db.Ping()
	if errPing != nil {
		log.Fatal("error connect to db", errPing.Error())
	} else {
		fmt.Println("berhasil")
	}

	defer db.Close()

	// SELECT DATA
	// select * from guru
	rows, errSelect := db.Query("SELECT id, nama, telepon, email FROM guru")
	if errSelect != nil {
		log.Fatal("error query select", errSelect.Error())
	}

	for rows.Next() {
		var datarow Guru
		errScan := rows.Scan(&datarow.Id, &datarow.Nama, &datarow.Telepon, &datarow.Email)
		if errScan != nil {
			log.Fatal("error scan select", errScan.Error())
		}

		fmt.Println(datarow)
	}

}
