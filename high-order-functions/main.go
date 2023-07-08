package main

import "fmt"

// function to calculate area of circle
func calcArea(radius float64) float64 {
	return 3.14 * radius * radius
}

// function to calculate perimeter of circle
func calcPerimeter(radius float64) float64 {
	return 2 * 3.14 * radius
}

// function to calculate diameter of circle
func diameter(radius float64) float64 {
	return 2 * radius
}

// function to print result
func printResult(radius float64, calcFunction func(r float64) float64) {
	result := calcFunction(radius)
	fmt.Printf("Result is: %f\n", result)
	fmt.Printf("Thank you for using our program\n")
}

// function to return function based on query
func getFunction(query int) func(r float64) float64 {
	query_to_func := map[int]func(r float64) float64{
		1: calcArea,
		2: calcPerimeter,
		3: diameter,
	}
	return query_to_func[query]
}

func main() {

	// variable declaration
	var query int
	var radius float64

	// taking input from user
	fmt.Print("Enter the radius of circle: ")
	fmt.Scanf("%.2f", &radius)
	fmt.Printf("Enter \n 1 for area \n 2 for perimeter \n 3 for diameter \n")
	fmt.Scanf("%d", &query)

	// calling function to print result
	printResult(radius, getFunction(query))

}
