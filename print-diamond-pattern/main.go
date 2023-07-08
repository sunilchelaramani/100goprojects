package main

import "fmt"

func main() {

	var i, j, k, row int

	fmt.Print("Enter Diamond Star Pattern Rows = ")
	fmt.Scanln(&row)

	fmt.Println("**** Diamond Star Pattern ****")

	for i = 1; i <= row; i++ {
		for j = 1; j <= row-i; j++ {
			fmt.Printf(" ")
		}
		for k = 1; k <= i*2-1; k++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

	for i = row - 1; i > 0; i-- {
		for j = 1; j <= row-i; j++ {
			fmt.Printf(" ")
		}
		for k = 1; k <= i*2-1; k++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
