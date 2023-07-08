package main

import "fmt"

type Expense struct {
	expense string
	amount  float64
	date    string
}

type ExpenseList []Expense

func (expenses ExpenseList) Total() float64 {
	total := 0.0
	for _, expense := range expenses {
		total += expense.amount
	}
	return total
}

func (e *Expense) getName() string {
	return e.expense
}

func main() {
	expenses := ExpenseList{
		Expense{"Grocery", 50.0, "2022-01-01"},
		Expense{"Gas", 30.0, "2022-01-02"},
		Expense{"Restaurant", 40.0, "2022-01-03"},
	}

	fmt.Println(expenses.Total())
	fmt.Println(expenses[0].getName())
}
