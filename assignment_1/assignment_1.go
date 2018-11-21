package main

import (
	"fmt"
	"strings"
)

func main() {
	//Manipulasi String

	var firstName string
	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("lenght is : ", len(firstName))

	myFirstNameUpper := strings.ToUpper(firstName)
	fmt.Println(myFirstNameUpper)

	var lastName string
	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("lenght is : ", len(lastName))

	myLastNameUpper := strings.ToUpper(lastName)
	fmt.Println(myLastNameUpper)

	fmt.Printf("Your Name is: %s %s", firstName, lastName)

}
