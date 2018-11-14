package main

import "fmt"

func main() {
	friends := [6]string{"John", "Jane", "Joe", "Janet", "sauth", "maya"}
	fmt.Println("Length:", len(friends))
	fmt.Println("Capacity:", cap(friends))
	for i := 0; i < len(friends); i++ {
		fmt.Printf("Index: %v\tValue: %v\n", i, friends[i])
	}
}
