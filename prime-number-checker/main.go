// write a program to check if a given number is prime or not

package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	if isPrime(num) {
		fmt.Println("Prime")
	} else {
		fmt.Println("Not Prime")
	}
}

func isPrime(num int) bool {
	if num == 1 {
		return false
	}
	if num == 2 {
		return true
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// Enter a number: 11
// Prime
