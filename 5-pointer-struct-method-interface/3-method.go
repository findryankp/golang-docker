package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
}

//function
func fullName(firstName string, lastName string) (fullName string) {
	fullName = firstName + " " + lastName
	return
}

//method
func (emp Employee) fullName() string {
	return emp.FirstName + " " + emp.LastName
}

func main() {
	e := Employee{
		FirstName: "Ross",
		LastName:  "Geller",
	}

	// fmt.Println(fullName(e.FirstName, e.LastName))
	fmt.Println(e.fullName())
}
