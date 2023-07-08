// random number generator

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// generate a random number between 0 and 100
	randomNumber := rand.Intn(100)

	// print the random number
	fmt.Println(randomNumber)
}
