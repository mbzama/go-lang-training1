package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	age       int
	married   bool
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	//Create struct 1
	alex := person{"Alex", "Anderson", 35, true, contactInfo{"t@test.com", 121}}

	//Create struct 2
	john := person{
		firstName: "John",
		lastName:  "St",
		age:       50,
		contactInfo: contactInfo{
			email:   "john@test.com",
			zipCode: 950,
		},
	}

	//Create struct 3
	var alice person
	alice.firstName = "alice"
	alice.lastName = "mary"
	alice.contactInfo.email = "alice@gmail.com"
	alice.zipCode = 5000

	alex.print()
	john.print()
	alice.print()

	fmt.Println("----Updated values-----")
	alex.updateName("Alex3")
	alex.print()
}

func (p person) print() {
	fmt.Println(p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
