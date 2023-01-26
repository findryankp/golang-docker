package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Obtaining a slice from an array `array`
	// array[low:high]
	var primes = [5]int{2, 3, 5, 7, 11}

	// Creating a slice from the array
	var part_primes []int = primes[1:4]

	// part_primes = append(part_primes, 10000)
	// menambah data ke slice akan menambah data ke array juga

	fmt.Println(reflect.ValueOf(part_primes).Kind())
	fmt.Println(part_primes)

	// var data1 []int
	// fmt.Println(data1)
	// var data2 = []int{1, 4, 6, 9, 10}
	// fmt.Println(data2)
	data3 := []int{1, 4, 6, 9, 10}
	fmt.Println(data3)
	data3 = append(data3, 20)
	data3 = append(data3, 51)
	fmt.Println(data3)

	// var data4 = make([]int, 5, 10)
	// fmt.Println(data4)

	data5 := []string{"surabaya", "malang", "blitar", "kediri"}
	fmt.Println(data5)
	fmt.Println(len(data5))
	data5 = append(data5, "jogja")
	data5 = append(data5, "bali", "pasuruan")
	fmt.Println(data5)
	fmt.Println(data5[6])
	fmt.Println(len(data5))

}
