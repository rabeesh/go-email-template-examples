package main

import (
	"fmt"
	"time"
)

type Address struct {
	Firstname string
	LastName  string
}

type Item struct {
	Name     string
	Quantity int
	Total    float64
}

type Order struct {
	Billing Address
	Date    time.Time
	Number  string
	Items   []Item
}

func main() {

	tp := NewTplParser([]string{
		"orderNewCustomer",
		"orderNewAdmin",
	})

	order := Order{
		Date:    time.Now(),
		Billing: Address{"Rabeeshkumar", "A R"},
		Number:  "12320",
		Items:   []Item{Item{"P1", 2, 23.23}, Item{"P2", 8, 99.01}},
	}

	// get email template for customer
	customerEmailContent, err := tp.Parse("orderNewCustomer", order)
	if err != nil {
		panic(err)
	}
	fmt.Println(customerEmailContent)

	// get email template for admin
	adminEmailContent, err := tp.Parse("orderNewAdmin", order)
	if err != nil {
		panic(err)
	}
	fmt.Println(adminEmailContent)
}
