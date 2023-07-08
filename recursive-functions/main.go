package main

import "fmt"

// Recursive function
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// Main function
func main() {
	fmt.Println(factorial(5))
}
