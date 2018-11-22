package main

import (
	"fmt"
)

func main() {

	//Basic Map
	var mapPerson map[int]string
	mapPerson = make(map[int]string)

	mapPerson[1] = "Go"
	mapPerson[2] = "Ruby"
	mapPerson[3] = "Python"
	mapPerson[4] = "Java"

	for k, v := range mapPerson {
		fmt.Println(k, v)
	}

	java, ok := mapPerson[4]
	if !ok {
		fmt.Println("not cool")
	} else {
		fmt.Println("yes is a ", java)
		fmt.Println("===========================")
	}

	mapBoom := make(map[int]Boom)
	mapBoom[1] = Boom{Id: 1, Name: "Golang"}
	mapBoom[2] = Boom{Id: 2, Name: "Ruby"}
	mapBoom[3] = Boom{Id: 3, Name: "Python"}

	for k, v := range mapBoom {
		fmt.Println(k, v.Name)
	}

}

type Boom struct {
	Id   int
	Name string
}
