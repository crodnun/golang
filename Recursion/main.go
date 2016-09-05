package main

import "fmt"

func factorial(number int) int {
	if number == 0 {
		return 1
	}

	return number * factorial(number - 1)
}

func finobacci(n int) int {
	if n == 0 {
		return 0

	}else if  n == 1 {
		return 1

	} else {
		return finobacci(n - 1) + finobacci(n - 2)
	}
}

func main() {
	fmt.Println(factorial(4))

	for i :=0; i < 10; i++ {
		fmt.Println("Finobacci of ", i, ": ",finobacci(i))
	}
}
