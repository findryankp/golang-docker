package main

import (
	"fmt"
	"strings"
)

func main() {
	// // 1. len string
	// sentence := "Hello you"
	// lenSentence := len(sentence)
	// fmt.Println(lenSentence)

	// // 2. compare string
	// str1 := "abc"
	// str2 := "abc"
	// fmt.Println(str1 == str2)

	// // 3. Contains
	// str := "something"
	// substr := "thig"
	// res := strings.Contains(str, substr)
	// fmt.Println("res", res) // true

	// 4. substring
	value := "cat;dog"
	// Take substring from index 4 to length of string.
	// t;dog --> 2:7
	// dog --> 4:7
	// cat --> 0:3
	// substring := value[0:3]
	substring := value[1:3]
	// substring := value[4:len(value)]
	fmt.Println(substring)

	// 5. Replace
	var t string
	s := "this[things]I wo[uld lik[e to remove"
	t = strings.Replace(s, "[", "-", 2)
	// t = strings.Replace(t, "]", "+", -1)
	fmt.Printf("%s\n", t)

	// 6. Insert
	p := "green"
	batas := 4
	q := p[0:batas] + "HI" + p[batas:]
	fmt.Println(p, q)
	fmt.Println(string(p[0]))

}
