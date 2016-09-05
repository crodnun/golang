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

// main() func to test variadic calls

func main() {
	fmt.Println("Average is: ", average(123, 323, 33, 3323, 33, 45))
}


