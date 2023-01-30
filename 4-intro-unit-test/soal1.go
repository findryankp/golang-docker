package main

import "fmt"

/*
 * Complete the 'compareTriplets' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func compareTriplets(a []int32, b []int32) []int32 {
	var jumA int32
	var jumB int32
	var hasil []int32
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			jumA++
		} else if a[i] < b[i] {
			jumB++
		}
	}
	hasil = append(hasil, jumA, jumB)
	return hasil
}

func main() {
	a := []int32{5, 6, 7}
	b := []int32{3, 6, 10}
	fmt.Println(compareTriplets(a, b))
}
