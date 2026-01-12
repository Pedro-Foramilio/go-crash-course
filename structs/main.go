package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	age       int
	contact   contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		age:       30,
		contact: contactInfo{
			email:   "alex.anderson@example.com",
			zipCode: 12345,
		},
	}
	alex.updateFirstName("Alexander")
	alex.print()
}

func (p *person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
