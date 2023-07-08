package main

import (
	"fmt"
	"time"
)

func main() {

	starttime := time.Now()

	c1 := make(chan string)
	c2 := make(chan string)

	//wg := sync.WaitGroup{}
	//wg.Add(2)

	go func() {
		time.Sleep(7 * time.Second)
		// assuming that the first channel takes 7 seconds to execute
		c1 <- "one"
		// wg.Done()
	}()

	go func() {
		//time.Sleep(1 * time.Second)
		c2 <- "two"
		//wg.Done()
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("timeout")
			// break
			// You can either add non-blocking tasks or break the loop here
		default:
			fmt.Println("no channel operation ready")
			// this performs a non-blocking tasks
		}
	}
	//wg.Wait()
	elaspedtime := time.Since(starttime)
	fmt.Println("Time taken to execute the program is ", elaspedtime)
}
