package main

import (
	"fmt"
)

type Vector struct {
	X int
	Y int
}

type Player struct {
	ID   int
	Name string
}

func main() {

	p := Person{
		ID:   1,
		Name: "Go",
		Age:  13,
	}

	printPerson(p)

	var v Vector
	v.X = 10
	v.Y = 5

	fmt.Println(v)

	fmt.Println("X = ", v.X)
	fmt.Println("Y = ", v.Y)

	player1 := Player{ID: 1, Name: "Sauthrully"}

	fmt.Println(player1.ID)
	fmt.Println(player1.Name)
	fmt.Println("==============================")
}

type Person struct {
	ID   int
	Name string
	Age  int
}

func printPerson(p Person) {
	fmt.Println("==============================")
	fmt.Println("ID = ", p.ID)
	fmt.Println("Name = ", p.Name)
	fmt.Println("Age = ", p.Age)
	fmt.Println("==============================")
}
