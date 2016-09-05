package main

import "fmt"

func zero2(value *int) {
	*value = 0
}

func checkCleared(value int) {
	if value == 0  {
		fmt.Println("Value was cleared")
	} else {
		fmt.Println("Value was NOT cleared")
	}
}

func main() {
	a := 43

	var b *int = &a

	fmt.Println("B is the address ", b)
	fmt.Println("A's address is ", &a)
	fmt.Println("B dereferences ", *b)
	fmt.Println("A contains ", *b)

	zero(a)

	fmt.Println("A contains ", a)

	checkCleared(a)

	zero2(&a)

	fmt.Println("A contains ", a)

	checkCleared(a)
}

func zero(value int) {
	value = 0
}
