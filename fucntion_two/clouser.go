package main

import (
	"fmt"
)

func main() {

	//clouser
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

	//Function as Argument
	f := func(v string) bool {
		return v == "golang"
	}

	result := match("golang", f)
	fmt.Println(result)

}

// clouser
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

//Function as Argument
func match(v string, f func(string) bool) bool {
	if f(v) {
		return true
	}
	return false
}
