package main

import "fmt"

func main() {
	// long declaration
	var salary = map[string]int{}
	fmt.Println(salary)
	salary["budi"] = 10000
	salary["rudi"] = 20000
	salary["eko"] = 30000
	salary["adi"] = 0
	fmt.Println("data", salary["adi"])
	fmt.Println("data", salary["dodi"])
	fmt.Println(len(salary))

	value, isExist := salary["dodi"]
	fmt.Println("valuenya", value, "isExist:", isExist)
	if isExist == true {
		fmt.Println("datanya ada lho")
	} else {
		fmt.Println("datanya tidak ada cuy")
	}

	// // long declaration with value
	// var salary_a = map[string]int{"umam": 1000, "iswanul": 2000}
	// fmt.Println(salary_a)

	// // short declaration
	// salary_b := map[string]int{}
	// fmt.Println(salary_b)

	// // using make
	// var salary_c = make(map[string]int)
	// salary_c["doe"] = 7000 // assign value
	// fmt.Println(salary_c)

	var data1 = map[int]string{}
	data1[15] = "lima belas"
	data1[20] = "dua puluh"
	data1[9] = "sembilan"
	data1[15] = "seratus"
	// fmt.Println(data1)
	// fmt.Println(data1)
	fmt.Println(data1)

	// var data2 = map[bool]string{}
	// data2[true] = "benar"
	// data2[false] = "salah"
	// data2["true"] = "salah"

}
