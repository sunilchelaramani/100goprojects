package main

import "fmt"

func main() {

	// for loop with init, condition and post
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for loop with single condition
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// for loop with no condition
	i = 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}

	// for loop with continue
	i = 0
	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i >= 5 {
			break
		}
	}

	// for loop with range
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}
}
