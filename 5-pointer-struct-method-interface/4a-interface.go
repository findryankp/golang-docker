package main

import "fmt"

func describe(input interface{}) {
	fmt.Printf("(%v, %T)\n", input, input)
}

func cetak(input any) {
	fmt.Printf("(%v, %T)\n", input, input)
}

func main() {
	var data interface{}
	data = 10
	describe(data)

	data = "halo"
	describe(data)

	data = true
	describe(data)

	fmt.Println()
	value1 := 10
	cetak(value1)
	value2 := "halo"
	cetak(value2)

	value3 := []int{1, 2, 3}
	cetak(value3)
	describe(value3)

	var datamap = map[string]any{}
	datamap["nama"] = "Budi"
	datamap["status"] = true
	datamap["age"] = 20
	datamap["angka"] = value3

	cetak(datamap)

}
