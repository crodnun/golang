package main

import "fmt"

func main() {
	var myMap map[string]int = make(map[string]int)

	myMap["A"] = 1
	myMap["B"] = 2

	fmt.Println(myMap)

	// Other different type of map initialization
	myMap2 := map[string]int{"fsdfa": 1, "fsdfasf": 4}

	fmt.Println(myMap2)

	// delete a given entry
	delete(myMap2, "fsdfa")

	fmt.Println(myMap2)
}