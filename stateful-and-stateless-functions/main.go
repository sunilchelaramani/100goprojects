// example code explaining stateful and stateless functions

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// stateless function
func stateless() {
	fmt.Println("Stateless function")
}

// stateful function
func stateful() {
	fmt.Println("Stateful function")
}

func main() {
	// stateless function
	stateless()

	// stateful function
	stateful()

	// stateful function with random number
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Random number: ", rand.Intn(100))
}
