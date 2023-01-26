package main

import "fmt"

func main() {
	var bil1 int = 5
	var bil2 int = 3
	// fmt.Println(bil1 + bil2)
	hasil1 := bil1 + bil2
	fmt.Println(hasil1)
	hasil2 := bil1 - bil2
	fmt.Println(hasil2)
	hasil3 := bil1 * bil2
	fmt.Println(hasil3)
	hasil4 := bil1 / bil2
	fmt.Println(hasil4)
	hasil5 := bil2 % bil1
	fmt.Println(hasil5)

	bil3 := 10
	fmt.Println(bil3)
	bil3 = bil3 + 5
	bil3 = bil3 + 5
	fmt.Println(bil3)

	var nama1 string = "bambang"
	var nama2 string = "pamungkas"
	hasilnama := nama1 + " " + nama2 + "timnas"
	fmt.Println(hasilnama)

	var data uint64 = 10000000
	fmt.Println(data)

	var angka1 int = 5
	var angka2 int = 2

	// var angka3 float64 = float64(angka1 / angka2) //2.0
	var angka3 float64 = float64(angka1) / float64(angka2) // 2.5
	fmt.Println(angka3)

}
