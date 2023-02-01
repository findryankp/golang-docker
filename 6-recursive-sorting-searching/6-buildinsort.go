package main

import (
	"fmt"
	"sort"
)

func main() {

	// An example of sorting `int`s.
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	strs := []string{"ca", "af", "ad", "bc"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// We can also use `sort` to check if a slice is
	// already in sorted order.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

	int1 := []int{7, 2, 4}
	sort.Sort(sort.Reverse(sort.IntSlice(int1)))

	int2 := []int{7, 2, 4}
	/*
		< sort ASC
		> sort DESC
	*/
	sort.SliceStable(int2, func(i, j int) bool {
		return int2[i] > int2[j]
	})
	fmt.Println("Int2:   ", int2)
}
