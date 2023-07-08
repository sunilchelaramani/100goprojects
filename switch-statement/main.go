package main

import "fmt"

func main() {
	var num1 int = 99
	var num2 int = 100

	switch num1 {
	case 98, 99:
		fmt.Println("It's equal to 98 or 99")
	case 100:
		fmt.Println("It's equal to 100")
	default:
		fmt.Println("It's not equal to 98, 99 or 100")
	}

	switch {
	case num2 < 100:
		fmt.Println("It's less than 100")
	case num2 == 100:
		fmt.Println("It's equal to 100")
	case num2 > 100:
		fmt.Println("It's greater than 100")
	}
}
