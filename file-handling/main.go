// file handling in golang

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// create a file
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// write to file
	file.WriteString("Hello World")
	file.Close()

	// read from file
	stream, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	readString := string(stream)
	fmt.Println(readString)
}
