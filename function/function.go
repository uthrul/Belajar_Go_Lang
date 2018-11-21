package main

import (
	"fmt"
)

func main() {
	fmt.Println("function on Golang")

	x := 10
	y := 10

	z := add(x, y)
	fmt.Println(z)

	name := "sauthrully"
	result := hello(name)

	fmt.Println(result)

	xSwap := "hello"
	ySwap := "golang"

	resultX, resultY := swap(xSwap, ySwap)
	fmt.Println(resultY, resultX)

}

func add(x int, y int) int {
	return x * y
}

func hello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func swap(x string, y string) (string, string) {
	return y, x
}
