package main

import (
	"strings"
)

func main() {
	reader := strings.NewReader("Let us catch up over a cup of coffee")
	// read string in chunks of 5 bytes, print ascii value of each character
	for {
		buffer := make([]byte, 5)
		bytesRead, err := reader.Read(buffer)
		if err != nil {
			println(err.Error())
			break
		}
		println(string(buffer[:bytesRead]))
	}
}
