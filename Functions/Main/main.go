package main

import "fmt"

func myFirstFunction(name string) int {
	fmt.Println("String parameter is", name);
	return 0;
}

func myFirstFunction2Parameters(name, surname string ) int {
	fmt.Println("String parameter is", name, surname);
	return 0;
}


func myFuncWithMultipleReturn(name, surname string) (string, string, string) {
	return "Carlos", "Rodriguez", "Nunez"
}

func myFuncWithMultipleReturn2(name, surname string) (result string) {
	result = fmt.Sprint(name, surname)

	return
}

// main() func is the entry point of every program

func main() {
	fmt.Println("Hello world!!")

	myFirstFunction("carlos")
	myFirstFunction2Parameters("Carlos", "Rodriguez")
	fmt.Println(myFuncWithMultipleReturn("Carlos", "Rodriguez"))
	fmt.Println(myFuncWithMultipleReturn2("Carlos", "Rodriguez"))
}


