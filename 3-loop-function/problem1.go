package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	// your code here
	var hasil string
	if len(a) > len(b) {
		check := strings.Contains(a, b)
		if check {
			hasil = b
		}
	} else {
		check := strings.Contains(b, a)
		if check {
			hasil = a
		}
	}
	return hasil
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
	fmt.Println(Compare("alterra", "alta"))   //
}
