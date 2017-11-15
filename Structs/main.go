package main

import (
	"fmt"
)

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
		lastName:  "Kings",
		age:       34,
	}

	alex.contact = contactInfo{
		"dummy@dummy.com",
		23433,
	}

	alex.updateFirstName("Tymmy")

	alex.printPerson()
}

// the * represents the memory address, this is description of a type and only will work when get a address type person
func (p *person) updateFirstName(newFirstName string) {
	//Update the memory in this specific address and turns to value
	(*p).firstName = newFirstName
}

func (p person) printPerson() {
	fmt.Println(p)
}
