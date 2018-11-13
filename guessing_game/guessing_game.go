package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var userInput int
	var secretNumber int

	rand.Seed(time.Now().Unix())
	secretNumber = rand.Intn(15)
	fmt.Println("your number: ", secretNumber)

	fmt.Print("Please enter Your number : ")
	fmt.Scan(&userInput)
	fmt.Println("your number: ", userInput)

	if userInput == secretNumber {
		fmt.Println("You Won")
	} else if userInput < secretNumber {
		fmt.Println("Too Low...")
	} else if userInput > secretNumber {
		fmt.Println("Too High..")
	}
}
