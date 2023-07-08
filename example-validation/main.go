// validator class using regular expressions
// to validate email and password

package main

import (
	"errors"
	"fmt"
	"regexp"
)

type Validator struct {
}

func (v *Validator) IsEmailValid(email string) error {
	// Regular expression for email validation
	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`

	// Check if the email matches the regex pattern
	match, _ := regexp.MatchString(emailRegex, email)
	if !match {
		return errors.New("invalid email format")
	}

	return nil
}

func (v *Validator) IsPasswordStrong(password string) error {
	// Check password strength criteria
	if len(password) < 8 {
		return errors.New("password should be at least 8 characters long")
	}

	// Add more password strength criteria as needed

	return nil
}

func main() {
	validator := &Validator{}

	var email string
	fmt.Print("Enter email: ")
	fmt.Scanln(&email)
	err := validator.IsEmailValid(email)
	if err != nil {
		fmt.Println("Email validation error:", err)
	}

	var password string
	fmt.Print("Enter password: ")
	fmt.Scanln(&password)
	err = validator.IsPasswordStrong(password)
	if err != nil {
		fmt.Println("Password validation error:", err)
	}
}
