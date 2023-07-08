// using method
package main

import "fmt"

type Product struct {
	brand string
	price float64
}

type ProductList []Product

func (pl ProductList) findProduct(min, max float64) []Product {
	var result []Product
	for _, product := range pl {
		if product.price >= min && product.price <= max {
			result = append(result, product)
		}
	}
	return result
}

func main() {
	products := ProductList{
		{brand: "HP", price: 1000.00},
		{brand: "Dell", price: 1200.00},
		{brand: "Lenovo", price: 1300.00},
		{brand: "Asus", price: 1100.00},
		{brand: "Acer", price: 1500.00},
	}

	result := products.findProduct(1000.00, 1300.00)

	fmt.Println(result)
}

/*
Using function
package main

import "fmt"

type Product struct {
	brand string
	price float64
}

func findProduct(products []Product, min, max float64) []Product {
	var result []Product
	for _, product := range products {
		if product.price >= min && product.price <= max {
			result = append(result, product)
		}
	}
	return result
}

func main() {
	products := []Product{
		{brand: "HP", price: 1000.00},
		{brand: "Dell", price: 1200.00},
		{brand: "Lenovo", price: 1300.00},
		{brand: "Asus", price: 1100.00},
		{brand: "Acer", price: 1500.00},
	}

	result := findProduct(products, 1000.00, 1300.00)

	fmt.Println(result)
}
*/
