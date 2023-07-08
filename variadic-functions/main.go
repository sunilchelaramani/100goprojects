package main

import "fmt"

// variadic function
func addNums(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// variadic function with fixed parameters
func formatString(text string, left, right rune) string {
	return fmt.Sprintf("%c%v%c", left, text, right)
}

func main() {
	// declare slice
	x := []int{1, 2, 3}
	y := []int{4, 5, 6}

	// append slices
	all := append(x, y...)

	// pass slice to variadic function
	result := addNums(all...)
	fmt.Println(result)

	// pass fixed parameters to variadic function
	text := formatString("format this", '(', ')')
	fmt.Println(text)
}
