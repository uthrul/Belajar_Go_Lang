package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var userInput int
	var secretNumber int
	var quit bool = false

	rand.Seed(time.Now().Unix())
	secretNumber = rand.Intn(15)
	fmt.Println("your number: ", secretNumber)

	for quit != true {
		fmt.Print("Please enter Your number : ")
		fmt.Scan(&userInput)
		if userInput == secretNumber {
			fmt.Println("You Won")
			quit = true
		} else if userInput < secretNumber {
			fmt.Println("Too Low...")
		} else if userInput > secretNumber {
			fmt.Println("Too High..")
		}
	}
}
