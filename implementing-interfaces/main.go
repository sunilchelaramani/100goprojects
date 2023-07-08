package main

import "fmt"

type catalog interface {
	shipping() float64
	tax() float64
}

type discount interface {
	offer() float64
}

type configurable struct {
	name       string
	price, qty float64
}

func (c *configurable) tax() float64 {
	return c.price * c.qty * 0.05
}

func (c *configurable) shipping() float64 {
	return c.qty * 5
}

func (c *configurable) offer() float64 {
	return c.price * 0.15
}

type download struct {
	name       string
	price, qty float64
}

func (d *download) tax() float64 {
	return d.price * d.qty * 0.10
}

type simple struct {
	name       string
	price, qty float64
}

func (s *simple) tax() float64 {
	return s.price * s.qty * 0.03
}

func (s *simple) shipping() float64 {
	return s.qty * 3
}

func (s *simple) offer() float64 {
	return s.price * 0.10
}

func main() {
	tshirt := configurable{}
	tshirt.price = 250
	tshirt.qty = 2
	fmt.Println("Configurable Product")
	fmt.Println("Shipping Charge: ", tshirt.shipping())
	fmt.Println("Tax: ", tshirt.tax())
	fmt.Println("Discount: ", tshirt.offer())

	mobile := simple{"Samsung S-7", 3000, 2}
	fmt.Println("\nSimple Product")
	fmt.Println(mobile.name)
	fmt.Println("Shipping Charge: ", mobile.shipping())
	fmt.Println("Tax: ", mobile.tax())
	fmt.Println("Discount: ", mobile.offer())

	book := download{"Python in 24 Hours", 50, 1}
	fmt.Println("\nDownloadable Product")
	fmt.Println(book.name)
	fmt.Println("Tax: ", book.tax())
}
