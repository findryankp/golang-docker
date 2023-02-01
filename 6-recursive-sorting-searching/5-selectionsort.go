package main

import "fmt"

func selectionSort(elements []int) []int {
	// n := len(elements)
	for k := 0; k < len(elements); k++ {
		minimal := k
		for j := k + 1; j < len(elements); j++ {
			// < mengurutkan secara ASC
			// > mengurutkan secara DESC
			if elements[j] > elements[minimal] {
				minimal = j
			}
		}
		//swap a,b = b,a
		elements[k], elements[minimal] = elements[minimal], elements[k]
	}
	return elements
}

func main() {
	data := []int{3, 6, 8, 2, 4, 10, 15, 20, 7, 9}
	result := selectionSort(data)
	fmt.Println(result)
}
