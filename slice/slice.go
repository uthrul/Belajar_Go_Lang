package main

import (
	"fmt"
)

func main() {
	// Slice(dinamic array)

	mySlice := []int{10, 20, 30, 40, 50}

	for i, v := range mySlice {
		fmt.Println(i, v)
	}

	mySlice_one := []string{"Go", "Rails", "Django", "React"}
	mySlice_one = append(mySlice_one, "ruby", "python")

	for _, v := range mySlice_one {
		fmt.Println(v)
	}

	//Slice dan Struct
	var booms []Boom
	booms = []Boom{Boom{ID: 1, Name: "Eden Hazard"}, Boom{ID: 2, Name: "Bp"}}
	booms = append(booms, Boom{ID: 3, Name: "Messi"})

	for _, x := range booms {
		fmt.Println(x)
	}
}

type Boom struct {
	ID   int
	Name string
}
