package main

import (
	"fmt"
)

const n string = "mencoba golang variable dan const"

func main() {

	m := "contoh vaiable "

	var x int
	x = 10

	y := 10

	var z int
	z = x * y

	u := z / 5

	s := z + u

	fmt.Println(n)
	fmt.Println(m, x, y, z)
	fmt.Println(u)
	fmt.Println(s)
}
