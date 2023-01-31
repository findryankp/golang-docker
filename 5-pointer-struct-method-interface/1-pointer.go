package main

import "fmt"

func main() {
	var a int = 100
	fmt.Println("value a", a)
	fmt.Println("address a", &a)

	var p *int = &a
	fmt.Println("value p", p)
	fmt.Println("value from address p", *p)
	*p = 200
	fmt.Println("value from address p", *p)
	fmt.Println("value a", a)

	var hasil = a + 1
	fmt.Println(hasil)

	var nama = "Budi"
	var namaAddress *string = &nama
	nama = "Rudi"
	fmt.Println(*namaAddress)
}
