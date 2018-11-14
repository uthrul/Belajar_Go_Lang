package main

import "fmt"

func main() {
	name := "sauth"
	age := 24
	weight := 70.5
	adult := false

	fmt.Printf("name : %T\t", name)
	fmt.Printf("age : %T\n", age)
	fmt.Printf("weight : %T\n", weight)
	fmt.Printf("adult : %T\n", adult)
}
