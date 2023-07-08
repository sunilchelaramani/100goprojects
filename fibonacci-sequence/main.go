// Write a program that generates the Fibonacci sequence up to a specified number of terms.

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please enter a number.")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fibonacci(n))
}

func fibonacci(n int) []int {
	sequence := make([]int, n)

	for i := 0; i < n; i++ {
		if i < 2 {
			sequence[i] = i
		} else {
			sequence[i] = sequence[i-1] + sequence[i-2]
		}
	}

	return sequence
}

// Output
// $ go run main.go 10
// [0 1 1 2 3 5 8 13 21 34]
