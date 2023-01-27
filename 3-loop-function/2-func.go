package main

import (
	"fmt"
)

func main() {
	// var jam = 21
	// if jam <= 12 {
	// 	CetakHello() // cetakhello di main
	// } else if jam == 21 {
	// 	service.CetakHello() //cetakhello di service
	// 	service.SapaNama()
	// } else {
	// 	CetakSelamatTidur()
	// }

	// var data1 = phi()
	// fmt.Println("return phi adalah", data1)

	var result1 = sapaOrang("alta")
	fmt.Println(result1)
	var result2 = sapaOrang("kharisma")
	fmt.Println(result2)
	var result3 = sapaOrang("adi")
	fmt.Println(result3)

	p := 10
	l := 5
	var hasil2 = LuasPersegiPanjang(p, l)
	fmt.Println("Luas 1 :", hasil2)
	var hasil3 = LuasPersegiPanjang(20, 7)
	fmt.Println("Luas 2 :", hasil3)

	fmt.Println(LuasPersegiPanjang(10, 7)) //70
	fmt.Println(LuasPersegiPanjang(10, 5)) //50
	fmt.Println(LuasPersegiPanjang(5, 10)) //50

}

func CetakHello() {
	fmt.Println("Hello Dunia")
}

func CetakSelamatTidur() {
	fmt.Println("Selamat tidur")
}

func phi() float64 {
	var nilai = 1000.1
	// var nilai = "hasil"
	return nilai
}

func sapaOrang(nama string) string {
	var hasil = "halo " + nama + " have a nice day!"
	return hasil
}

func LuasPersegiPanjang(panjang int, lebar int) int {
	// your code
	var hasil = panjang * lebar
	// fmt.Println(hasil)
	return hasil
}
