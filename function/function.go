package main

import (
	"fmt"
)

func main() {
	fmt.Println("function on Golang")

	//fucntion two
	boom := func(x, y int) int {
		return x * y
	}
	fmt.Println("boom is = ", boom(10, 5))

	x := 10
	y := 10
	fmt.Println("add is = ", add(x, y))

	helloGo := func(name string) string {
		return fmt.Sprint("Hello Golang Mania ", name)
	}
	fmt.Println(helloGo(" sauth rully"))

	//fucntion one
	name := "sauthrully"
	result := hello(name)

	fmt.Println(result)

	xSwap := "hello"
	ySwap := "golang"

	resultX, resultY := swap(xSwap, ySwap)
	fmt.Println(resultY, resultX)

}

func add(x, y int) int {
	return x * y
}

func hello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func swap(x string, y string) (string, string) {
	return y, x
}
