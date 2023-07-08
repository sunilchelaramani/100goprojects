package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// MapValues takes a slice of values and a function that maps a value to a new value.
// This is example of how we can use constraints.Ordered to restrict the type of values
// that can be passed to the function.
// Example of how you can write Generic functions that can use constraints to accept various types.

func MapValues[T constraints.Ordered](values []T, mapFunc func(T) T) []T {
	var newValues []T
	for _, v := range values {
		newValues = append(newValues, mapFunc(v))

	}
	return newValues
}

func main() {
	values := []int{1, 2, 3, 4, 5}
	fmt.Println(MapValues(values, func(v int) int {
		return v * 2
	}))
}
