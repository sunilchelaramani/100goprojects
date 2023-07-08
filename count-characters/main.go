package main

func main() {

	// count characters in given string

	str := "Hello World"

	// convert string to byte array
	byteArray := []byte(str)

	// get length of byte array
	length := len(byteArray)

	// print length
	println(length)

	// print length of string
	println(len(str))

	// print length of string
	println(len("Hello World"))

	//count characters in given string without using len function

	str = "Hello World"

	// convert string to byte array
	byteArray = []byte(str)

	// declare counter variable
	var counter int

	// iterate over byte array
	for range byteArray {

		// increment counter
		counter++
	}

	// print counter
	println(counter)

	// count characters in given string without spaces

	str = "Hello World"

	// convert string to byte array
	byteArray = []byte(str)

	// declare counter variable
	counter = 0

	// iterate over byte array
	for _, char := range byteArray {

		// check if character is not space
		if char != ' ' {

			// increment counter
			counter++
		}
	}

	// print counter
	println(counter)

}
