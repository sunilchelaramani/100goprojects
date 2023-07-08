package main

import "fmt"

func fib(n int, ch chan<- int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
}

func main() {
	ch := make(chan int) // create an unbuffered channel
	var n int = 10
	go fib(n, ch) // generate the first 10 Fibonacci numbers in a separate goroutine

	var y []int
	// read values from the channel until it's closed
	for i := 0; i < n; i++ {
		x := <-ch
		y = append(y, x)
	}
	fmt.Println(y)
}
