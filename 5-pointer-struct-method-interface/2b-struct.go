package main

import "fmt"

func main() {
	type Book struct {
		Judul    string
		Harga    int
		Penerbit string
		Penulis  string
	}

	var bookdata Book
	fmt.Println("Masukkan judul Buku:")
	fmt.Scanln(&bookdata.Judul)

	fmt.Println(bookdata)
}
