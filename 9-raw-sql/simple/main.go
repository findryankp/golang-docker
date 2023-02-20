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

	fmt.Println("Pilih Menu:\n1. Read Data\n2. Add Data")
	fmt.Println("Input pilihan anda:")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		// SELECT DATA
		// select * from guru
		rows, errSelect := db.Query("SELECT id, nama, telepon, email FROM guru")
		if errSelect != nil {
			log.Fatal("error query select", errSelect.Error())
		}

		var allGuru []Guru
		for rows.Next() {
			var datarow Guru
			errScan := rows.Scan(&datarow.Id, &datarow.Nama, &datarow.Telepon, &datarow.Email)
			if errScan != nil {
				log.Fatal("error scan select", errScan.Error())
			}
			allGuru = append(allGuru, datarow)
			// fmt.Println(datarow)
		}

		// fmt.Println(allGuru)
		for i, v := range allGuru {
			fmt.Println("i:", i, "v:", v.Id, v.Nama)
		}

	// CASE 2
	case 2:
		//INSERT DATA
		newGuru := Guru{}
		fmt.Println("Masukkan ID Guru:")
		fmt.Scanln(&newGuru.Id)
		fmt.Println("Masukkan Nama Guru:")
		fmt.Scanln(&newGuru.Nama)
		fmt.Println("Masukkan Telepon Guru:")
		fmt.Scanln(&newGuru.Telepon)
		fmt.Println("Masukkan Email Guru:")
		fmt.Scanln(&newGuru.Email)

		var query = "INSERT INTO guru(id, nama, telepon, email) VALUES(?,?,?,?)"
		statement, errPrepare := db.Prepare(query)
		if errPrepare != nil {
			log.Fatal("error prepare insert", errPrepare.Error())
		}

		result, errInsert := statement.Exec(newGuru.Id, newGuru.Nama, newGuru.Telepon, newGuru.Email)
		if errInsert != nil {
			log.Fatal("error exec insert", errInsert.Error())
		} else {
			row, _ := result.RowsAffected()
			if row > 0 {
				fmt.Println("proses berhasil dijalankan")
			} else {
				fmt.Println("proses gagal")
			}
		}

	// CASE 3
	case 3:
		fmt.Println("update")

	// CASE 4
	case 4:
		fmt.Println("delete")

	// CASE 5
	case 5:
		fmt.Println("get by id")

	// DEFAULT
	default:
		fmt.Println("Input tidak sesuai")

	}

}
