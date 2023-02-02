package main

import "fmt"

// n^2 --> n*n
func quadratic(n int) int {
	var result int = 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result += 1
		}
	}
	return result
}

// n*n*n --> n^3
func quadratic2(n int) int {
	var result int = 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				result += 1
			}
		}
	}
	return result
}

func main() {
	fmt.Println(quadratic2(5))
}
