package main

import "fmt"

func main() {
	namaFuncion()
	defer fmt.Println("satu")
	defer fmt.Println("dua")
	fmt.Println("tiga")

}

func namaFuncion() {
	fmt.Println("func 2")
	panic("error")
	defer fmt.Println("func 1")
	fmt.Println("func 3")
}
