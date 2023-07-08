// example of using the go-flags package

package main

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
)

// Options struct
type Options struct {
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	Help    bool   `short:"h" long:"help" description:"Show this help message"`
}

// ParseOptions parses the command line options
func ParseOptions() *Options {
	var options Options
	parser := flags.NewParser(&options, flags.Default)
	parser.Usage = "[OPTIONS]"
	_, err := parser.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return &options
}

func main() {
	options := ParseOptions()
	fmt.Println("Verbose:", options.Verbose)
	fmt.Println("Help:", options.Help)
}

// Output:
// $ go run main.go -h
// Usage: main [OPTIONS]
//

// Options:
//   -v, --verbose    Show verbose debug information
//   -h, --help       Show this help message

// $ go run main.go -v
// Verbose: [true]
// Help: false
