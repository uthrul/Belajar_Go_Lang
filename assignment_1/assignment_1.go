package main

import "fmt"

func main() {
	var firstName string
	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)

	var lastName string
	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Your Name is: %s %s", firstName, lastName)
}
