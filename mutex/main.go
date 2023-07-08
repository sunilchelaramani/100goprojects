// example using mutex to lock and unlock a variable

package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func main() {
	wg.Add(2)
	go increment("Foo:")
	go increment("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func increment(s string) {
	for i := 0; i < 20; i++ {
		mutex.Lock()
		x := counter
		x++
		counter = x
		fmt.Println(s, i, "Counter:", counter)
		mutex.Unlock()
	}
	wg.Done()
}
