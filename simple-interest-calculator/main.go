package main

import "fmt"

func main() {

	var amount, InterestRate, time, simpleI float64

	fmt.Print("Enter the Principal or Total Amount = ")
	fmt.Scanln(&amount)

	fmt.Print("Enter the rate of Interest = ")
	fmt.Scanln(&InterestRate)

	fmt.Print("Enter the Total number of Years = ")
	fmt.Scanln(&time)

	simpleI = (amount * InterestRate * time) / 100

	fmt.Println("\nThe Simple Interest  = ", simpleI)
}
