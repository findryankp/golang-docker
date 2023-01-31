package main

import "fmt"

type Karyawan struct {
	name   string
	salary int
}

func (e *Karyawan) changeName(newName string) {
	(*e).name = newName
}

func (e *Karyawan) increaseSalary(number int) {
	(*e).salary += number
}

func main() {
	e := Karyawan{
		name:   "Ross Geller",
		salary: 1200,
	}

	// e before name change
	fmt.Println("e before name change =", e)
	// create pointer to `e`
	ep := &e
	// change name
	ep.changeName("Monica Geller")
	ep.increaseSalary(10000)
	// e after name change
	fmt.Println("e after name change =", e)
}
