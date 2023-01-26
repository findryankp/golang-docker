package main

import "fmt"

func main() {
	var data [3]int
	for i := 0; i < len(data); i++ {
		fmt.Println("Masukkan angka:")
		fmt.Scanln(&data[i])
	}

	fmt.Println(data)
}
