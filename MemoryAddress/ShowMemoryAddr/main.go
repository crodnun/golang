package main

import "fmt"

func main() {
	a := 43
	fmt.Println("a's memory address is ", &a)
	fmt.Printf("a's memory address is %d\n", &a)
	fmt.Printf("a's memory address is %X\n", &a)
}
