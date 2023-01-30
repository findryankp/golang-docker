package salary

func HitungGaji(posisi string) int {
	var gaji int
	if posisi == "manager" {
		gaji = 30000000
	} else if posisi == "senior" {
		gaji = 20000000
	} else if posisi == "junior" {
		gaji = 10000000
	} else {
		gaji = 0
	}

	return gaji
}
