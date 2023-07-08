package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type CustomData interface {
	constraints.Ordered | []byte | []rune
}

type User[T CustomData] struct {
	ID   int64
	Name string
	Data T
}

func main() {

	u := User[string]{ID: 1, Name: "John", Data: "1"}

	fmt.Println(u)

}
