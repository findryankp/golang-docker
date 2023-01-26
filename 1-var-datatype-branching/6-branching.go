package main

import "fmt"

func main() {
	// var myAge = 30

	// if myAge == 5 {
	// 	fmt.Println("You're too young")
	// }

	// if myAge == 17 {
	// 	fmt.Println("So Sweet")
	// }

	// if myAge >= 17 && myAge < 30 {
	// 	fmt.Println("My Age is between 17 and 30")
	// }

	// dadAge := 9
	// if dadAge < myAge {
	// 	fmt.Println(dadAge)
	// }

	// number := 3
	// if number%2 == 0 {
	// 	fmt.Println("Genap")
	// } else {
	// 	fmt.Println("Ganjil")
	// }

	// date := 1
	// if date >= 25 {
	// 	fmt.Println("Baru gajian")
	// } else if date == 24 {
	// 	fmt.Println("akan gajian")
	// } else if date == 1 {
	// 	fmt.Println("awal bulan")
	// } else if date == 30 {
	// 	fmt.Println("akhir bulan")
	// } else {
	// 	fmt.Println("belum gajian")
	// }

	date := 0
	if date >= 1 && date <= 31 {
		if date >= 25 {
			fmt.Println("Sudah gajian")

			if date == 25 {
				fmt.Println("horray")
			}

		} else {
			fmt.Println("belum gajian")
		}
	} else if date == 0 {
		fmt.Println("gak ada cuy")
	} else {
		fmt.Println("tanggal salah")
	}

	if date == 0 {
		fmt.Println("kalau tanggal 0, maka ini tampil")
	}

}
