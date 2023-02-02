package main

import "fmt"

func linearSearch(elements []int, cari int) int {
	var count int
	for i := 0; i < len(elements); i++ {
		count++
		fmt.Println("count:", count)
		if elements[i] == cari {
			return i
		}
	}
	return -1
}

func binarySearch(elements []int, cari int) int {
	var kiri int = 0
	var kanan int = len(elements) - 1
	for kiri <= kanan {
		var tengah = (kiri + kanan) / 2
		if cari < elements[tengah] {
			kanan = tengah - 1
		} else if cari > elements[tengah] {
			kiri = tengah + 1
		} else if cari == elements[tengah] {
			return tengah
		}
	}

	return -1
}

func main() {
	fmt.Println(binarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 15))
}
