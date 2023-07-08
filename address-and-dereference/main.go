package main

import "fmt"

func main() {

	x := 10

	ptr := &x

	fmt.Printf("x is of type %T and value is %v\n", ptr, *ptr)
	fmt.Printf("x is of type %T and value ie %v\n", &x, &x)
	fmt.Printf("x is of type %v\n", *(ptr))
}
