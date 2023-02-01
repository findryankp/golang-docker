package main

import "fmt"

func main() {
	var secret interface{}

	secret = 10
	var number = secret.(int) * 10
	fmt.Println(secret, "multiplied by 10 is :", number)

}
