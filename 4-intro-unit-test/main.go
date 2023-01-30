package main

import (
	"alta/salary/salary"
	"fmt"
)

/*
Program kalkulasi gaji

kalkulasi gaji diperoleh berdasarkan posisi karyawan, dengan ketentuan:
1. manager = 35 juta
2. senior = 25 juta
3. junior = 15 juta

kita diminta untuk membuat app perhitungan gaji karyawan

nama karyawan - posisi - gaji nya berapa

go test ./... -coverprofile=cover.out && go tool cover -html=cover.out
*/

/*
Pertama: jalankan go mod init, untuk generate file go.mod
go mod init namaproject

untuk menggunakan library testify:
go get -u github.com/stretchr/testify

untuk menjalankan unit test
go test ./...

go test ./... -coverprofile=cover.out && go tool cover -html=cover.out

*/
func main() {
	var nama1 = "Budi"
	var posisi1 = "manager"
	var gaji1 int = salary.HitungGaji(posisi1)
	// if posisi1 == "manager" {
	// 	gaji1 = 30000000
	// } else if posisi1 == "senior" {
	// 	gaji1 = 20000000
	// } else if posisi1 == "junior" {
	// 	gaji1 = 10000000
	// } else {
	// 	gaji1 = 0
	// }

	fmt.Printf("nama: %s, posisi: %s, gaji: %d\n", nama1, posisi1, gaji1)

	var nama2 = "Rudi"
	var posisi2 = "junior"
	var gaji2 int = salary.HitungGaji(posisi2)
	// if posisi2 == "manager" {
	// 	gaji2 = 35000000
	// } else if posisi2 == "senior" {
	// 	gaji2 = 25000000
	// } else if posisi2 == "junior" {
	// 	gaji2 = 15000000
	// } else {
	// 	gaji2 = 0
	// }

	fmt.Printf("nama: %s, posisi: %s, gaji: %d\n", nama2, posisi2, gaji2)

}
