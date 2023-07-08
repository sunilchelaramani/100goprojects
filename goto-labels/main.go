package main

import "fmt"

func main() {
	// Declare a label
outerLabel:

	for i := 1; i <= 3; i++ {
		fmt.Printf("Outer loop iteration: %d\n", i)

		for j := 1; j <= 3; j++ {
			fmt.Printf("    Inner loop iteration: %d\n", j)

			// Conditional check
			if i == 2 && j == 2 {
				// Jump to the label
				goto outerLabel
			}
		}
	}

	fmt.Println("Program never end")
}
