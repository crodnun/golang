package main

import "fmt"

// This function calculates the remainder of the input number after dividing it by each of the numbers from 1 to 20
// Returns true if has no remainder for all them or false when condition is not met and any remainder is != 0
func isTheExpectedNum(n int) bool {
	for i := 1; i <= 20; i++ {
		if n%i != 0 {
			// Stop here when remainder is not 0 for any in [1-20]
			return false
		}
	}

	// it is a multiple of all them (1-20) - true
	return true
}

func main() {
	// Initialize starting value, we can do safely this to safe time - lower values do not match the first 10
	var next int = 2520

	// Loop till match found
	for {
		// We do the checking in a separate function (there we go the 1-20 remainder calculation for input number)
		if(isTheExpectedNum(next)) {
			break
		}

		next++
	}

	// Show the result
	fmt.Println("The result is ", next)
}


/*
   2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
   What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
 */

/*
 If we run this program obtain:
 C:\Training\go\workspace\src\github.com\crodnun\EulerExercise>go run main.go

 The result is  232792560

 */