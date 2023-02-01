package main

import "fmt"

/*
	3,6,8,2,4,10,15,20,7,9
	cari angka 4

	cek satu persatu angka dari depan
*/

func linearSearch(elements []int, cari int) int {
	var count int
	for i := 0; i < len(elements); i++ {
		count++
		fmt.Println("count:", count)
		if elements[i] == cari {
			return i
		}
	}
	return -1
}

func main() {
	data := []int{3, 6, 8, 2, 4, 10, 15, 20, 7, 9}
	var cari = 9
	hasil := linearSearch(data, cari)
	// fmt.Println(hasil)
	if hasil != -1 {
		fmt.Println("data ditemukan di index ke:", hasil)
	} else {
		fmt.Println("data tidak ditemukan")
	}
}
