package main

import (
	"fmt"
	"regexp"
)

func main() {
	phone1 := "1(234)5678901x1234"
	phone2 := "(+351) 282 43 50 50"
	phone3 := "90191919908"
	phone4 := "555-8909"
	phone5 := "001 6867684"
	phone6 := "001 6867684x1"
	phone7 := "1 (234) 567-8901"
	phone8 := "1-234-567-8901 ext1234"

	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)

	fmt.Printf("Pattern: %v\n", re.String()) // print pattern
	fmt.Printf("\nPhone: %v\t:%v\n", phone1, re.MatchString(phone1))
	fmt.Printf("Phone: %v\t:%v\n", phone2, re.MatchString(phone2))
	fmt.Printf("Phone: %v\t\t:%v\n", phone3, re.MatchString(phone3))
	fmt.Printf("Phone: %v\t\t\t:%v\n", phone4, re.MatchString(phone4))
	fmt.Printf("Phone: %v\t\t:%v\n", phone5, re.MatchString(phone5))
	fmt.Printf("Phone: %v\t\t:%v\n", phone6, re.MatchString(phone6))
	fmt.Printf("Phone: %v\t\t:%v\n", phone7, re.MatchString(phone7))
	fmt.Printf("Phone: %v\t:%v\n", phone8, re.MatchString(phone8))
}
