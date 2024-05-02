package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email       string
	phoneNumber string
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// notice the pointer receiver, used to update the value of the struct in memory
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func main() {
	arielGuerrero := person{
		firstName: "Ariel",
		lastName:  "Guerrero",
		contactInfo: contactInfo{
			email:       "arielguerrero@email.com",
			phoneNumber: "123-456-7890",
		},
	}

	arielGuerrero.updateName("Vrixl")
	arielGuerrero.print()
}
