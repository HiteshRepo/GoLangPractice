package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int

	// firstName, lastName, city, gender string
	// age int
}

// Greeting method - value receiver
func (p Person) greet() string { // p acts like this keyword from other langs
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years old."
}

// hasBirthday method - pointer receiver
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried method - pointer receiver
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "f" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {

	person1 := Person{firstName: "Hitesh", lastName: "Pattanayak", city: "BBSR", gender: "m", age: 26}
	// person1 := Person{"Hitesh", "Pattanayak", "BBSR", "m", 26} - alternative

	fmt.Println(person1)           //{Hitesh Pattanayak BBSR m 26}
	fmt.Println(person1.firstName) // Hitesh
	person1.age++
	fmt.Println(person1.age) // 27
	person1.hasBirthday()
	fmt.Println(person1.greet()) // Hello, my name is Hitesh Pattanayak and I am 28 years old.
	person1.getMarried("Pattanayak2")
	fmt.Println(person1.greet()) // Hello, my name is Hitesh Pattanayak2 and I am 28 years old.

}
