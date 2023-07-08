package main

import "fmt"

type Person[T comparable, V int | string] map[T]V

func main() {

	// Create a new Person
	p := make(Person[string, int])

	// Add a new key-value pair
	p["john"] = 25

	// Add a new key-value pair
	p["dean"] = 30

	fmt.Printf("p = %+v\n", p)
}
