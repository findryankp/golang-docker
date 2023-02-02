package main

import "fmt"

func dominant(n int) int {
	var result int = 0
	for i := 0; i < n; i++ {
		result += 1
	}
	result = result + 10
	return result
}

func main() {
	fmt.Println(dominant(100))
}
