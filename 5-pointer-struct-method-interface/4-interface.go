package main

import "fmt"

type calculate interface {
	luas() int
	keliling() int
	Gambar(url string) (string, int)
}

type square struct {
	side int
}

type rectangle struct {
	panjang int
	lebar   int
}

// keliling implements calculate
func (rc rectangle) keliling() int {
	return 2 * (rc.panjang * rc.lebar)
}

// luas implements calculate
func (rc rectangle) luas() int {
	return rc.panjang * rc.lebar
}

func (rc rectangle) Gambar(url string) (string, int) {
	return "", 0
}

func (rc rectangle) besarSudut() int {
	return 90
}

func (sq square) luas() int {
	return sq.side * sq.side
}

func (sq square) keliling() int {
	return 4 * sq.side
}

// Gambar implements calculate
func (sq square) Gambar(url string) (string, int) {
	return "", 0
}

func main() {
	var kotak1 = square{
		side: 10,
	}
	var dimResult calculate
	dimResult = kotak1
	fmt.Println("Luas:", dimResult.luas())

	var dimResultRect calculate
	dimResultRect = rectangle{
		panjang: 10,
		lebar:   20,
	}
	fmt.Println("Luas Rect", dimResultRect.luas())
}
