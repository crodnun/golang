package main

import "fmt"

// This is an example of defer keyword and how it delays execution of a statement just till the end of the function
// code deferred executes before functions exits
func main() {
	//fmt.Println("World!")  // If we use this instead of next line we obtain different/wrong results
	defer fmt.Println("World!")
	fmt.Print("Hello ")
}
