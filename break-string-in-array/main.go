package main

import (
	"fmt"
	"strings"
)

func main() {
	testString := "Australia is a country and continent surrounded by the Indian and Pacific oceans."
	testArray := strings.Fields(testString)
	fmt.Println("Array :", testArray)
	for _, v := range testArray {
		fmt.Println(v)
	}
}
