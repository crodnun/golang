package main

import "fmt"

func main() {
	// ask user to enter 2 numbers
	var highNumber, lowNumber, remainder int

	// enter 2 numbers
	fmt.Print("Enter a high int number: ")
	fmt.Scanln(&highNumber)

	fmt.Print("Enter a low int number: ")
	fmt.Scanln(&lowNumber) // Use this to also read the newline

	// Calculate the remainder
	remainder = highNumber % lowNumber
	fmt.Printf("Remainder of %d/%d is %d", highNumber, lowNumber, remainder)
}