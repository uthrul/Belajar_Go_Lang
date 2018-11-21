package main

import (
	"fmt"
)

func main() {
	nextVal := getNum()

	fmt.Println(nextVal())
	fmt.Println(nextVal())
	fmt.Println(nextVal())

	lv := love("sauthrully ")
	fmt.Println(lv("Golang"))
	fmt.Println(lv("ruby on rails"))
	fmt.Println(lv("Django"))
	fmt.Println(lv("JavaScript"))
	fmt.Println(lv("React Js"))
	fmt.Println(lv("React Native"))
}

func getNum() func() int {
	x := 0
	return func() int {
		x = x + 1
		return x
	}
}

func love(name string) func(string) string {
	return func(things string) string {
		return fmt.Sprint(name, "love ", things)
	}
}
