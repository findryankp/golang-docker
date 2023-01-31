package main

import "fmt"

type Person struct {
	Firstname string
	Age       uint
	Weight    float32
}

type Siswa struct {
	Nama  string
	Umur  uint
	Kelas string
	NIS   string
}

type CV struct {
	Nama           string
	DataPendidikan []Pendidikan
}

type Pendidikan struct {
	Nama         string
	Level        string
	TahunMulai   uint
	TahunSelesai uint
}

func main() {
	//long decl
	var manusia1 Person
	manusia1.Firstname = "Budi"
	manusia1.Age = 20
	manusia1.Weight = 58.5
	fmt.Println(manusia1)
	fmt.Println(manusia1.Age)
	fmt.Println(manusia1.Firstname)
	fmt.Println(manusia1.Weight)
	fmt.Println(manusia1.Age + 1)

	var manusia1Nama = "budi"
	var manusia1Age = 20
	var manusia1Weight = 58.5
	fmt.Println(manusia1Nama, manusia1Age, manusia1Weight)

	var manusia2 Person
	manusia2.Firstname = "Rudi"
	manusia2.Age = 25
	manusia2.Weight = 60.0
	fmt.Println(manusia2)

	//long decl
	var manusia3 = Person{"Andi", 20, 50.5}
	fmt.Println(manusia3)

	// long decl
	var manusia4 = Person{
		Firstname: "Adi",
		Weight:    50.6,
	}
	fmt.Println(manusia4)

	//short declaration
	manusia5 := Person{
		Firstname: "Adi",
		Age:       20,
		Weight:    50.6,
	}
	fmt.Println(manusia5)

	var siswa1 = Siswa{
		Nama:  "Ani",
		Umur:  15,
		Kelas: "SMP IX",
		NIS:   "S0001",
	}
	var siswa2 = Siswa{
		Nama:  "Budi",
		Umur:  15,
		Kelas: "SMP IX",
		NIS:   "S0002",
	}
	fmt.Println(siswa1.Nama)
	fmt.Println(siswa2.Nama)

	var pendidikan1 Pendidikan
	pendidikan1.Nama = "SMA N 1 Jakarta"
	pendidikan1.Level = "SMA"
	pendidikan1.TahunMulai = 2018
	pendidikan1.TahunSelesai = 2020

	var pendidikan2 Pendidikan
	pendidikan2.Nama = "SMP 1 Jakarta"
	pendidikan2.Level = "SMP"
	pendidikan2.TahunMulai = 2018
	pendidikan2.TahunSelesai = 2020

	var cv1 CV
	cv1.Nama = "Andi"
	cv1.DataPendidikan = append(cv1.DataPendidikan, pendidikan2)
	cv1.DataPendidikan = append(cv1.DataPendidikan, pendidikan1)
	fmt.Println(cv1.Nama, cv1.DataPendidikan[0].Nama)
	fmt.Println(cv1.DataPendidikan)
	fmt.Println(cv1.DataPendidikan[0].TahunMulai)

	// tampilkan seluruh data pendidikan dari Andi
	for i := 0; i < len(cv1.DataPendidikan); i++ {
		fmt.Printf("Nama Sekolah: %s\nLevel: %s\n", cv1.DataPendidikan[i].Nama, cv1.DataPendidikan[i].Level)
	}

	for _, v := range cv1.DataPendidikan {
		fmt.Printf("Nama Sekolah: %s\nLevel: %s\n", v.Nama, v.Level)
	}
}
