package main

import (
	"fmt"
)

func main() {
	// var kabupaten [5]string
	// // var data string
	// fmt.Println(reflect.ValueOf(kabupaten).Kind())
	// // fmt.Println(reflect.ValueOf(data).Kind())

	// kabupaten[0] = "surabaya"
	// kabupaten[1] = "malang"
	// kabupaten[2] = "jogja"
	// kabupaten[3] = "bali"
	// kabupaten[4] = "batu"

	// fmt.Println(kabupaten)

	// var nilai [10]int
	// nilai[0] = 100
	// nilai[1] = 90
	// nilai[5] = 80
	// nilai[9] = 5
	// fmt.Println(nilai[1])
	// fmt.Println(len(nilai))

	// odd_numbers := [5]int{1, 3, 5, 7, 9}      // Intialized with values
	// var even_numbers [5]int = [5]int{1, 2, 4} // Partial assignment

	// fmt.Println(odd_numbers)
	// fmt.Println(even_numbers)

	// data1 := [5]int{1: 2, 2: 4, 4: 15}
	// fmt.Println(data1)

	primes := [...]int{2, 3, 3, 8, 15, 19}
	// fmt.Println(reflect.ValueOf(primes).Kind())
	// fmt.Println(len(primes))
	fmt.Println(primes)

	// technique 1
	for i := 0; i < len(primes); i++ {
		fmt.Println(i, "isinya:", primes[i])
	}

	// technique 2
	for index, value := range primes {
		fmt.Println(index, "=>", value)
	}
	for _, value := range primes {
		fmt.Println(value)
	}
	// technique 3
	index := 0
	for range primes {
		fmt.Println(primes[index])
		index++
	}

}
