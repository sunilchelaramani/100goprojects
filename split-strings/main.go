//split strings into words

package main

import (
	"fmt"
	"strings"
)

func main() {
	//split string into words
	s := "This is a string"
	words := strings.Split(s, " ")
	fmt.Printf("Data Type is %T and Content of Slice is %v\n\n", words, words)

	fmt.Println("Now printing each word in a new line from the slice")
	for _, word := range words {
		fmt.Println(word)
	}
}
