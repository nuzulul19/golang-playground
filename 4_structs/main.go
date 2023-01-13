package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// **1
	// contact   contactInfo
	contactInfo
}

func main() {
	// zul := person{"Nuzulul", "Aulia"}
	// zul := person{firstName: "Nuzulul", lastName: "Aulia"}
	// var zul person
	// zul.firstName = "Nuzulul"
	// zul.lastName = "Aulia"

	// **1
	// zul := person{
	// 	firstName: "Nuzulul",
	// 	lastName:  "Aulia",
	// 	contact: contactInfo{
	// 		email:   "nuzulul19@gmail.com",
	// 		zipCode: 40500,
	// 	},
	// }

	zul := person{
		firstName: "Nuzulul",
		lastName:  "Aulia",
		contactInfo: contactInfo{
			email:   "nuzulul19@gmail.com",
			zipCode: 40500,
		},
	}

	// zulPointer := &zul // &(ampersand operator) get the address (pointer) of var
	zul.updateName("Zul") // Go allows us to take the shorcuts w/o need to define the var of pointer with reciever
	zul.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName // *(star operator) get value of the address (pointer)
}
