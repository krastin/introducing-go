package main

import "fmt"

type Person struct {
	Name string
}

type Android struct {
	Person
	Model string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func main() {
	joe := new(Person)
	joe.Name = "Joe"
	joe.Talk()

	iJoe := new(Android)
	iJoe.Name = "iJoe"
	iJoe.Talk()
	iJoe.Person.Talk()
}