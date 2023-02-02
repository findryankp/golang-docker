package main

import "fmt"

//bruteforce
func findMaxMin(input []int) (max int, min int) {
	max = input[0]
	min = input[0]
	var count int
	for i := 1; i < len(input); i++ {
		count++
		fmt.Println("count:", count)
		if input[i] > max {
			max = input[i]
		}
		if input[i] < min {
			min = input[i]
		}
	}
	return
}

//dnc
func findMaxMinDivide(input []int) (max int, min int) {
	max = input[0]
	min = input[0]
	var kiri, kanan = 1, len(input) - 1
	var count int
	for kiri <= kanan {
		count++
		fmt.Println("count:", count)
		if input[kiri] > max {
			max = input[kiri]
		}
		if input[kanan] > max {
			max = input[kanan]
		}

		if input[kiri] < min {
			min = input[kiri]
		}
		if input[kanan] < min {
			min = input[kanan]
		}
		kiri = kiri + 1
		kanan = kanan - 1
	}
	return
}

func main() {
	// max, min := findMaxMin([]int{10, 7, 3, 5, 8, -2, 2, 9, 0, -1})
	max, min := findMaxMinDivide([]int{-3, 10, 7, 3, 5, 8, -2, 2, 9, 0, -1})
	fmt.Println(max, min)
}
