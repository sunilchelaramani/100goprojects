// example code for using the goroutine
package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Hello World")
	}()

	go func() {
		fmt.Println("Hello World2")
	}()

	// wait for the goroutine to finish
	// time.Sleep(time.Second * 2)

	time.Sleep(time.Second * 2)

	// output is not deterministic
}
