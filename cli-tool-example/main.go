package main

import (
	"fmt"
	"os"
)

func printUsage() {
	fmt.Println("Usage: go run main.go <url>")
	fmt.Println("Example: go run main.go https://golang.org")
	fmt.Println("Help: go run main.go help")
	fmt.Println("Version: go run main.go version")
}

func parseArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

func printArgs() {
	fmt.Println(os.Args[1])
}

func main() {
	parseArgs()
	if len(os.Args) > 1 && os.Args[1] == "help" {
		printUsage()
	}
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println("Version 1.0.0")
	}

	if len(os.Args) > 1 && os.Args[1] != "help" && os.Args[1] != "version" {
		printArgs()
	}
}
