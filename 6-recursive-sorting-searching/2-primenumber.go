package main

import (
	"fmt"
	"math"
)

// bil prima adalah sebuah bil yang hanya habis dibagi dirinya sendiri dan 1
// tidak ada bilangan yng bisa membagi habis bil tsb selain dirinya sendiri dan 1
func primeNumber(number int) bool {
	if number <= 1 {
		// fmt.Println("cek 1")
		return false
	}
	// var count int
	// for i := 1; i <= number; i++ {
	// 	if number%i == 0 {
	// 		count++
	// 	}
	// }
	// if count > 2 {
	// 	fmt.Println("cek false")
	// 	return false
	// }
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	// fmt.Println("cek true")
	return true
}

func checkPrime(number int) bool {
	if number < 2 {
		return false
	}
	sqrtNumber := int(math.Sqrt(float64(number)))
	for i := 2; i <= sqrtNumber; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkPrime(11))          //true
	fmt.Println(checkPrime(9))           //false
	fmt.Println(checkPrime(1))           //false
	fmt.Println(checkPrime(100))         //false
	fmt.Println(checkPrime(97))          //true
	fmt.Println(checkPrime(63018038201)) //true

}
