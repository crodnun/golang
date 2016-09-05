package main

import "fmt"

// This example shows us how we can cope with anonymous functions and manage them as variables - functional programming concept -

func functionProvider() func() string {
	return func() string {
		return string("Hello I am also a function returned")
	}
}

func main() {
	// assign a function to a variable
	myFunction := func() {
		fmt.Println("Hello I am an anonymous function!!!")
	}

	myFunction()

	otherFunction := functionProvider()

	fmt.Println(otherFunction())

	fmt.Printf("My type is: %T\n", otherFunction)
	fmt.Printf("My type is: %T\n", myFunction)

}


