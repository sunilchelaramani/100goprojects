package main

import (
	"fmt"
	"log"

	"github.com/sethvargo/go-password/password"
)

func main() {

	fmt.Printf("Enter length of password: ")
	var length int
	fmt.Scanln(&length)

	// Generate a password that is 64 characters long with 10 digits, 10 symbols,
	// allowing upper and lower case letters, disallowing repeat characters.
	res, err := password.Generate(length, 2, 2, false, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
