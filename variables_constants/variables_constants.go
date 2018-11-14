package main

import "fmt"

func main() {
	// var - keyword
	var name string
	name = "sauthrully"
	fmt.Println("username : ", name)

	firstName := "saud"
	lastName := "maruli"
	fmt.Println("Full Name : ", firstName, lastName)

	var (
		a string = "tic"
		b int    = 4
		c bool   = true
	)
	fmt.Println(a, b, c)

	var d, e, f = "I am string ", 9.0, true
	fmt.Printf("d: %v, e: %2f, f: %v\n", d, e, f)

	d = " is now a new string "
	fmt.Println("d: ", d)

	const API_KEY = "dfbkj.abcd.1234"
	fmt.Printf(API_KEY)
}
