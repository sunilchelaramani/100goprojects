package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	starttime := time.Now()

	c1 := make(chan string)
	c2 := make(chan string)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		c1 <- "one"
		wg.Done()
	}()

	go func() {
		c2 <- "two"
		wg.Done()
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
	wg.Wait()
	elaspedtime := time.Since(starttime)
	fmt.Println("Time taken to execute the program is ", elaspedtime)
}
