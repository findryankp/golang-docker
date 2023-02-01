package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []int{12, 15, 15, 19, 24, 31, 53, 59, 60}
	var cari = 20
	result := sort.SearchInts(data, cari)
	// fmt.Println(result)
	if data[result] == cari {
		fmt.Println("data ditemukan")
	} else {
		fmt.Println("data tidak ditemukan")
	}
}
