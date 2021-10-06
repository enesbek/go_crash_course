package main

import (
	"fmt"
	"strconv"
)

//Define person struct
type Person struct {
	//firstName string
	//lastName  string
	//city      string
	//age       int

	firstName, lastName, city string
	age                       int
}

//Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

//hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

func main() {

	//Init person using struct
	person1 := Person{firstName: "John", lastName: "Johny", city: "London", age: 23}
	fmt.Println(person1)

	//Alternative
	person2 := Person{"John", "Johny", "London", 23}
	fmt.Println(person2)
	fmt.Println(person2.firstName)

	person2.age++
	fmt.Println(person2.age)

	fmt.Println(person1.greet())
	person1.hasBirthday()
	fmt.Println(person1.greet())
}
