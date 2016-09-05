package main

import "fmt"

func main() {
	// ask user to enter his/her name
	var userName,surname string
	//var surname string

	fmt.Print("Enter your name: ")
	fmt.Scanf("%s %s", &userName, &surname)

	fmt.Println("Hello", userName, surname)
}