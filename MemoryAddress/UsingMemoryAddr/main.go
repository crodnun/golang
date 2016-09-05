package main

import "fmt"

const metersPerYard float64 = 1.09361

func main() {
	var meters float64

	fmt.Print("Enter the number of meters: ")
	fmt.Scan(&meters)

	yards := meters * metersPerYard

	fmt.Printf("%f meters are %f yards\n", meters, yards)


}
