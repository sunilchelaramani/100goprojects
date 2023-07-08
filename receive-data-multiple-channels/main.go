package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 10; i < 20; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	for {
		select {
		case x, ok := <-ch1:
			if ok {
				fmt.Println("Received from ch1:", x)
			} else {
				fmt.Println("ch1 closed")
			}
		case x, ok := <-ch2:
			if ok {
				fmt.Println("Received from ch2:", x)
			} else {
				fmt.Println("ch2 closed")
			}
		}
	}
}
