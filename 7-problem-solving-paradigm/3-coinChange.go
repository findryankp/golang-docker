package main

import (
	"fmt"
	"sort"
)

/*
coinValue = 10,25, 5, 1
uang = 42

42 = 1,1,1,1,1, ... 42x --> 42 pecahan
42 = 5,5,5,5,5,5,5,5,1,1 --> 10 pecahan
42 = 10,10,10,10,1,1 --> 6 pecahan
42 = 25,5,5,5,1,1  -> 6 pecahan
42 = 25, 1,1,1,1,1,1,1,  --> banyak
42 = 25, 10,5,1,1 --> 5 pecahan

15 = 10, 5
65 = 25,25,10,5
*/

func coinChange(uang int, coins []int) []int {
	sort.SliceStable(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})

	result := []int{}
	for _, pecahan := range coins {
		// if uang >= pecahan {
		for uang >= pecahan {
			result = append(result, pecahan)
			uang = uang - pecahan
		}
		// } else {
		// 	continue
		// }
	}
	return result
}

func main() {
	coinValue := []int{10, 25, 5, 1}
	uang := 65

	fmt.Println(coinChange(uang, coinValue))

}
