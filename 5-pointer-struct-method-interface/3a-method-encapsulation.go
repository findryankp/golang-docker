package main

import "fmt"

type User struct {
	name string // Both non exported fields.
	age  int
}

type Mobil struct {
	Merk string
	CC   int
}

func (P User) GetName() string {
	return P.name + " amazing!"
}

func (P *User) IncreaseAge() {
	P.age = P.age + 1
}
func (mob Mobil) GetName() string {
	return mob.Merk + " mantabs"
}

func main() {
	PersonA := User{"John", 50}
	fmt.Printf("%v\n", PersonA)
	fmt.Println(PersonA.GetName())

	PersonA.IncreaseAge()
	fmt.Println(PersonA.age)

	MobilA := Mobil{"Toyota", 2000}
	fmt.Println(MobilA.GetName())
}
