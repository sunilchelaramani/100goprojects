package main

import "fmt"

func main() {

	// Here’s a basic example. Note that there is no parentheses around the three components of the for statement and the braces { } are always required.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	// range on arrays and slices provides both the index and value for each entry. Above we didn’t need the index, so we ignored it with the blank identifier _.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	// range on map iterates over key/value pairs.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	// range on map iterates over key/value pairs.
	for k := range kvs {
		fmt.Println("key:", k)
	}
	// range on strings iterates over Unicode code points. The first value is the starting byte index of the rune and the second the rune itself.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
