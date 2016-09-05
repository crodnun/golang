package main

import "fmt"

func main() {
	// loop over 1-100 and print Fizz when 3 multiple, Buzz when 5 multiple and FizzBuzz when both
	for i:=1; i <= 100; i++ {
		var isThreeMultiple bool = false
		var isFiveMultiple bool = false

		isThreeMultiple = (i % 3 == 0)
		isFiveMultiple = (i % 5 == 0)

		if isThreeMultiple && isFiveMultiple {
			fmt.Println("FizzBuzz")

		} else if isThreeMultiple {
			fmt.Println("Fizz")

		} else if isFiveMultiple {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}

	}
}