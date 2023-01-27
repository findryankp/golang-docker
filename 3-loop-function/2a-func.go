package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func greeting(hour int) {
	if hour < 12 {
		fmt.Println("Selamat Pagi")
	} else if hour < 18 {
		fmt.Println("Selamat Sore")
	} else {
		fmt.Println("Selamat Malam")
	}
}

func main() {
	hour := 15
	greeting(hour) // selamat sore
	greeting(10)   //selamat pagi
}
