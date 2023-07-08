package main

import "fmt"

type SpamMail struct {
	amount int
}

const (
	Small = iota
	Medium
	Large
)

const (
	Ground = iota
	Air
)

type Beltsize int
type Shipping int

func (b Beltsize) String() string {
	return [...]string{"Small", "Medium", "Large"}[b]
}

func (s Shipping) String() string {
	return [...]string{"Ground", "Air"}[s]
}

type Conveyor interface {
	Convey() Beltsize
}

type Shipper interface {
	Ship() Shipping
}

type WarehouseAutomator interface {
	Conveyor
	Shipper
}

func (s *SpamMail) String() string {
	return "Spam Mail"
}

func (s *SpamMail) Ship() Shipping {
	return Air
}

func (s *SpamMail) Convey() Beltsize {
	return Medium
}

func automate(item WarehouseAutomator) {
	fmt.Printf("Convey %v on %v belt\n", item, item.Convey())
	fmt.Printf("Ship %v via %v\n", item, item.Ship())
}

func main() {
	mail := &SpamMail{amount: 100}
	automate(mail)
}
