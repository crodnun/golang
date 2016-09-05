package main

import "fmt"

func variadicExample(numbers ... int) int {
	var max int

	for _, value := range numbers {
		if value > max {
			max = value
		}
	}

	return max
}

func isEven(number int) (int, bool) {
	var result int = number/2

	return result, (number%2==0)
}

func foo(numbers ... int) int {
	fmt.Println(numbers)

	return len(numbers)
}


func main() {
	var input int = 5

	fmt.Println(isEven(input))

	// receive here the result
	_, even := isEven(input)

	if(even) {
		fmt.Println(input, "Is an even number")
	} else {
		fmt.Println(input, "Is an odd number")
	}

	expr := func(number int) (int, bool) {
		var result int = number/2

		return result, (number%2==0)
	}

	_, even2 := expr(input)

	if(even2) {
		fmt.Println(input, "Is an even number")
	} else {
		fmt.Println(input, "Is an odd number")
	}

	var mySlice []int = []int{1, 4, 7, 29, 200, 78}

	fmt.Println("The max of", mySlice, "is", variadicExample(mySlice...))
	fmt.Println("The max of 1, 2, 4, 8, 90, 67, 34 is", variadicExample(1, 2, 4, 8, 90, 67, 34))
	fmt.Println("The max of <nothing> is", variadicExample())


	// show booleans usage
	fmt.Println((true && false) || (false && true) || !(false && false))

	foo(1,2,3)
}
