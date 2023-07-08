package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		ch1 <- 1
		time.Sleep(1 * time.Second)
		wg.Done()
	}()

	go func() {
		ch2 <- 2
		time.Sleep(2 * time.Second)
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(ch1)
		close(ch2)
	}()

	receivedCh1 := false
	receivedCh2 := false

	for !(receivedCh1 && receivedCh2) {
		select {
		case <-ch1:
			fmt.Println("ch1")
			receivedCh1 = true
		case <-ch2:
			fmt.Println("ch2")
			receivedCh2 = true
		}
	}

	fmt.Println("Program ended")
}
