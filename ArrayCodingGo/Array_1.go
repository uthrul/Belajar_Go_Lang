package main

import "fmt"

func main() {
	friends := [6]string{"John", "Jane", "Joe", "Janet", "sauth", "maya"}
	fmt.Println("Length:", len(friends))
	fmt.Println("Capacity:", cap(friends))
	for i := 0; i < len(friends); i++ {
		fmt.Printf("Index: %v\tValue: %v\n", i, friends[i])
	}

	var arr [5]string
	arr[0] = "GoLang"
	arr[1] = "Rails"
	arr[2] = "React"
	arr[3] = "Django"
	arr[4] = "python"

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
