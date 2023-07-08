// example code for json encoding and decoding

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Person struct
type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := Person{"James", "Bond", 20}
	p2 := Person{"Miss", "Moneypenny", 18}

	people := []Person{p1, p2}

	fmt.Println(people)

	// encoding to json
	bs, err := json.Marshal(people)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(bs))

	// decoding from json
	var people2 []Person
	err = json.Unmarshal(bs, &people2)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(people2)

	//indenting json
	bs, err = json.MarshalIndent(people, "", "\t")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(bs))
}

// OUTPUT:
// [{James Bond 20} {Miss Moneypenny 18}]
// [{"First":"James","Last":"Bond","Age":20},{"First":"Miss","Last":"Moneypenny","Age":18}]
// [{James Bond 20} {Miss Moneypenny 18}]

// NOTES:
// 1. json.Marshal() returns a slice of bytes ([]byte) and an error
// 2. json.Unmarshal() takes a slice of bytes ([]byte) and a pointer to a data structure
//    and populates the data structure with the values from the slice of bytes
// 3. json.Marshal() only exports fields that start with a capital letter
// 4. json.Unmarshal() only populates fields that start with a capital letter
