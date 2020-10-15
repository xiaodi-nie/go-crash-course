package main

import (
	"fmt" //print stuff
	"strconv"
)

//define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int

	//can also declare same type in a single line
	//firstName, lastName, city, gender string
	//age                               int
}

//greeting method (value receiver)
//value receiver methods are used mainly to access fields
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
//pointer receiver methods are used mainly to change some fields
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	//init a person using struct
	person1 := Person{firstName: "Samatha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	person2 := Person{"Bob", "Johnson", "New York", "m", 30}

	//fmt.Println(person1)
	//person1.age++
	//fmt.Println(person1.firstName)
	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	person1.getMarried("Williams")
	fmt.Println(person1.greet())

	person2.getMarried("Thomson")
	fmt.Println(person2.greet())

}
