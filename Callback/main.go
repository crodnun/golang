package main

import "fmt"

func cbCaller(numbers []int, callback func(n int)) {
	for _, number := range numbers {
		callback(number)
	}
}

func filter(numbers []int, callback func(num int) bool) []int {
	var result []int

	for _, number := range numbers {
		if(callback(number)) {
			result = append(result, number)
		}
	}

	return result
}

// This is an example of using a callback
// we create an anonymous function and pass it as callback to cbCaller
// Then it is used there to manage the slice of integers supplied as first argument
func main() {
	//numbers := []int{1, 2, 3 , 4}

	cbCaller([]int{1, 2, 3 , 4}, func(n int) {
		fmt.Println("Number called from callback: ", n)
	})

	sliceOut := filter([]int{3, 4, 5,6}, func(n int) bool {
		return n > 4;
	})

	fmt.Println(sliceOut)
}
