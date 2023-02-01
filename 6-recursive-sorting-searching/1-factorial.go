package main

import "fmt"

/*
factorial
5! = 5*4*3*2*1 = 120
	= 5 * 4!
	= 5 * (5-1)!
4! = 4*3*2*1 = 24
	= 4 * 3!
	= 4 * (4-1)!
3! = 3*2*1 = 6
	= 3*2!

2! = 2*1!
1! = 1

*/

func factorialLoop(n int) int {
	var result int = 1
	for i := n; i >= 1; i-- {
		// 1 * 1 * 2 * 3 * 4 * 5
		// 1 * 5 * 4 * 3 * 2 * 1
		result = result * i
	}
	return result
}

//recursive
func factorial(n int) int {
	if n == 1 { //base case
		return 1
	} else { // recurrence relation
		return n * factorial(n-1)
	}
}

func main() {
	// hasil := factorialLoop(5) // 120
	hasil := factorial(1) // 120
	fmt.Println(hasil)
}
