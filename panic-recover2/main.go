package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	SecretKey string `json:"secret_key"`
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("An error occurred: %v\n", r)
			fmt.Println("Application terminated gracefully.")
		} else {
			fmt.Println("Application executed successfully.")
		}
	}()

	// Read configuration file.
	configData, err := ioutil.ReadFile("config.json")
	if err != nil {
		// If there's an error reading the file, panic and terminate the program.
		panic(fmt.Sprintf("Error reading configuration file: %v", err))
	}

	var config Config
	err = json.Unmarshal(configData, &config)
	if err != nil {
		// If there's an error unmarshaling the JSON data, panic and terminate the program.
		panic(fmt.Sprintf("Error parsing configuration data: %v", err))
	}

	if config.SecretKey == "" {
		// If the secret key is empty, panic and terminate the program.
		panic("Secret key is missing from configuration")
	}

	// Continue with the program execution.
	fmt.Println("Application started successfully.")
}
