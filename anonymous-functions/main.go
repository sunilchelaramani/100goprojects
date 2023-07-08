package main

import "fmt"

// main function
func main() {

	// anonymous function
	x := func(a, b int) int {
		return a + b
	}(10, 20)

	// print the value of x
	fmt.Println(x)
}
