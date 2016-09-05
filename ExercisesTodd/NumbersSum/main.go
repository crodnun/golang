package main

import "fmt"

func main() {
	var result int = 0

	// loop over 1 to 1000 and add the numbers that are 3 or 5 multiple
	for i:=1; i < 1000; i++ {
		var isThreeMultiple bool = false
		var isFiveMultiple bool = false

		isThreeMultiple = (i % 3 == 0)
		isFiveMultiple = (i % 5 == 0)

		if isThreeMultiple || isFiveMultiple {
			result += i

		}
	}

	// show the result obtained
	fmt.Println("Result is :", result)
}