package main

import "fmt"

// main function
func main() {

	// defer statement
	defer fmt.Println("Bye")

	// print the value of x
	fmt.Println("Hello")

	// defer statement
	defer fmt.Println("World")
}
