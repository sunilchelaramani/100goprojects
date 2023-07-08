// simple program demonstrating how to create a struct and a method
package main

import "fmt"

// Person is a struct
type Person struct {
	Name string
	Age  int
}

// Speak is a method
func (p Person) Speak() {
	fmt.Println("Hello from " + p.Name)
}

// main is the entry point for the program
func main() {
	// create a new Person
	flavio := Person{Age: 39, Name: "Flavio"}
	// call the Speak method
	flavio.Speak()
}
