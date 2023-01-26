package main

import "fmt"

func main() {
	var firstName string
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName)

	// fmt.Println("hello ", firstName, "Selamat pagi")
	fmt.Printf("hello %s, Selamat pagi\n", firstName)

	var bilangan1 int
	fmt.Println("Masukkan sebuah angka: ")
	fmt.Scanln(&bilangan1)
	fmt.Println("bilangannya adalah:", bilangan1)
	fmt.Printf("bilangannya adalah: %d", bilangan1)

	// fmt.Scanf("%s %s", &firstName, &lastName)
	// fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	// fmt.Sscanf(input, format, &f, &i, &s)
	// fmt.Println("From the string we read: ", f, i, s)
}
