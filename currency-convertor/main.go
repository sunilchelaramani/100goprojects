package main

import "fmt"

func main() {

	var USD float64
	var INR float64

	USD = 74.5
	INR = 1

	productprice := 100.00
	result := productprice * USD / INR

	fmt.Printf("Your product amount in INR is %.2F\n", result)
}
