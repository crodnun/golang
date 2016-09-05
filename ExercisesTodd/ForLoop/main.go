package main

import "fmt"

func main() {
        ended := true
	i := 0
	// loop over infinite loop
	for ended {
		i++;

		if(i == 1256) {
			fmt.Println("Loop ended at", i, "!!")
			ended = false
		}

	}
}