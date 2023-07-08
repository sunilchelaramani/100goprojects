//sample program to show how to use time package to find out how much time it takes to execute a program

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Hello, World!")
	fmt.Println("Time taken to execute the program: ", time.Since(start))
}
