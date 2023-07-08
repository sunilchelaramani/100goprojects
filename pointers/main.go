package main

import "fmt"

func main() {
	var a int = 42
	var b *int = &a // pointer to a
	fmt.Println(a, b)
	fmt.Println(&a, b)
	fmt.Println(a, *b) // *b is the value stored at the address b
	a = 27
	fmt.Println(a, *b)
	*b = 14
	fmt.Println(a, *b)
}
