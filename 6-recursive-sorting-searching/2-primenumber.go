package main

import "fmt"

// bil prima adalah sebuah bil yang hanya habis dibagi dirinya sendiri dan 1
func primeNumber(number int) bool {
	if number <= 1 {
		fmt.Println("cek 1")
		return false
	}
	var count int
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			count++
		}
	}
	if count > 2 {
		fmt.Println("cek false")
		return false
	}
	fmt.Println("cek true")
	return true
}

func main() {
	fmt.Println(primeNumber(11)) //true
	fmt.Println(primeNumber(9))  //false
	fmt.Println(primeNumber(1))  //false
}
