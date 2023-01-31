package main

import "fmt"

type Book struct {
	Judul    string
	Harga    int
	Penerbit string
	Penulis  string
}

func main() {

	var rumahbaca []Book
	var book1 = Book{
		Judul:    "One Piece",
		Harga:    50000,
		Penerbit: "Gramedia",
		Penulis:  "Oda",
	}

	rumahbaca = append(rumahbaca, book1)

	rumahbaca = append(rumahbaca, Book{
		Judul:    "Naruto",
		Harga:    70000,
		Penerbit: "Gramedia",
		Penulis:  "Masashi",
	})

	fmt.Println("banyaknya buku:", len(rumahbaca))
	for _, v := range rumahbaca {
		fmt.Println("Judul Buku", v.Judul)
	}

	var mappeminjam = map[string]Book{}
	mappeminjam["budi"] = Book{
		Judul:    "One Piece",
		Harga:    50000,
		Penerbit: "Gramedia",
		Penulis:  "Oda",
	}

	mappeminjam["rudi"] = book1

	for k, v := range mappeminjam {
		fmt.Printf("nama peminjam: %s, Judul Buku: %s", k, v.Judul)
	}
}
