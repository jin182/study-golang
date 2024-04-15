package main

import "fmt"

func main() {
	name := [4] string {
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(name)

	a := name[0 : 2]
	b := name[1: 3]
	fmt.Println(a, b)

	b[0]= "XXX"
	fmt.Println(a, b)
	fmt.Println(name)
}