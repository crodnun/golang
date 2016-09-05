package main

import "fmt"

func main() {
	if true {
		fmt.Println("This is executed")
	} else if false {
		fmt.Println("This is NOT executed")
	} else {
		fmt.Println("This is neither executed")
	}

	// Example of initialization statement
	if a := true; a {
		fmt.Println("If statement with initialization")
	}

}
