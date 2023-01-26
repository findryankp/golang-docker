package main

import "fmt"

func main() {
	/*
		tampilkan data orang dan salary yang diatas/sama dengan 20000
	*/
	var salary = map[string]int{}
	// fmt.Println(salary)
	salary["budi"] = 10000
	salary["rudi"] = 20000
	salary["eko"] = 30000
	salary["adi"] = 0

	var result = map[string]int{}
	for key, value := range salary {
		if value >= 20000 {
			result[key] = value
		}
	}

	fmt.Println(result)
}
