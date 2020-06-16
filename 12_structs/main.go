package main

import (
	"fmt"
	"strconv"
)

//NOTE Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

//NOTE Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years old."
}

//NOTE hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

//NOTE gotMarried method (pointer receiver)
func (p *Person) gotMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	//NOTE initialize person using struct
	person1 := Person{firstName: "Peyton", lastName: "Sonnefeld", city: "Boise", gender: "M", age: 30}
	person2 := Person{"Chelsea", "Dunn", "Boise", "F", 25}

	fmt.Println(person1, person2)

	//NOTE Get single field
	fmt.Println(person1.firstName)

	person1.age++
	fmt.Println(person1.age)

	person1.hasBirthday()
	person2.gotMarried("Dunn")
	person2.gotMarried("Sonnefeld")
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
