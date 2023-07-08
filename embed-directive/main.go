package main

import (
	"embed"
)

var fileString string

var fileByte []byte

var folder embed.FS

func main() {

	print(fileString)
	print(string(fileByte))

	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
