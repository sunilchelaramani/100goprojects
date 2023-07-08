package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["m1"] = 7
	m["m2"] = 13

	fmt.Println("map:", m)

	v1 := m["m1"]
	fmt.Println("v1:", v1)

	v3 := m["m3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "m2")
	fmt.Println("map:", m)

	_, prs := m["m2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
