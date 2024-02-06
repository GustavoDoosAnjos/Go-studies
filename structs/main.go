package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	person := person{firstName: "Alex", lastName: "Anderson"}

	person.updateName("jimmy")
	fmt.Println(person)
}

func (p *person) updateName(newName string) {
	(*p).firstName = newName
}
