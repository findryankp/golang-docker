package main

import "fmt"

// pass by value
func fillCup(cup string) string {
	fmt.Println("address data awal", &cup)
	fmt.Println("data awal", cup)
	cup = "keisi"
	fmt.Println("data akhir", cup)
	return cup
}

// pass by reference
func fillCupPointer(cup *string) string {
	fmt.Println("data address awal", cup)
	fmt.Println("data value awal", *cup)
	*cup = "keisi"
	fmt.Println("data address akhir", cup)
	fmt.Println("data value akhir", *cup)
	return *cup
}

func main() {
	var cup string = "kosong"
	fmt.Println("address cup:", &cup)
	hasil1 := fillCup(cup)
	// hasil1 := fillCupPointer(&cup)
	fmt.Println("hasil1:", hasil1)
	fmt.Println("cup:", cup)
}
