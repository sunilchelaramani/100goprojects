package main

import "fmt"

func fib(n int, ch chan<- int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	ch := make(chan int, 10) // create a buffered channel with a capacity of 10
	go fib(10, ch)           // generate the first 10 Fibonacci numbers in a separate goroutine

	var y []int

	// read values from the channel until it's closed
	for x := range ch {
		y = append(y, x)
	}
	fmt.Println(y)
}
