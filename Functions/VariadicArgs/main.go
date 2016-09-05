package main

import "fmt"

func average(values ... float64) float64 {
	var total float64
	fmt.Println("Input values are", values);

	// calculate the average and store it in total
	for _, item := range(values) {
		//fmt.Println(index, item)
		total += item
	}

	return total/float64(len(values));
}

func sumInt(values ...int) int {
	var total int

	for _, value := range values {
		total += value
	}

	return total
}

// main() func to test variadic calls

func main() {
	var values []float64 = []float64{1, 2, 3, 4}
	var intValues []int = []int{1,2, 3, 100, 200, 23}

	fmt.Println("Average is: ", average(values...)) // we convert an slice into variadic arg expected
	fmt.Println("Sum is: ", sumInt(intValues...))
}


