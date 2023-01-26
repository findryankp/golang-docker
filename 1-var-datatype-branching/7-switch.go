package main

import "fmt"

func main() {
	// date := 25

	// switch date {
	// case 1:
	// 	fmt.Printf("Today is Monday")
	// 	fmt.Println("hello")
	// case 2:
	// 	fmt.Printf("Today is Tuesday")
	// case 25:
	// 	fmt.Println("horray")
	// default:
	// 	fmt.Printf("Invalid Date")
	// }

	// if date == 1 {
	// 	fmt.Printf("Today is Monday")
	// 	fmt.Println("hello")
	// } else if date == 2 {
	// 	fmt.Printf("Today is Tuesday")
	// } else if date == 25 {
	// 	fmt.Println("horray")
	// } else {
	// 	fmt.Printf("Invalid Date")
	// }

	// var today int = 2
	// switch {
	// case today == 1:
	// 	fmt.Printf("Today is Monday")
	// case today == 2:
	// 	fmt.Printf("Today is Tuesday")
	// default:
	// 	fmt.Printf("Invalid Date")
	// }

	value := 42
	switch value {
	case 100:
		fmt.Println(100)
		fallthrough
	case 42:
		fmt.Println(42)
		fallthrough
	case 1:
		fmt.Println(1)
		fallthrough
	default:
		fmt.Println("default")
	}

}
