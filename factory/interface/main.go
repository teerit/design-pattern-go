package main

import "fmt"

type Person interface {
	SayHello()
}

type tiredPerson struct {
	name string
	age  int
}

type person struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s, I am %d year old.\n", p.name, p.age)
}

func (p *tiredPerson) SayHello() {
	fmt.Printf("Sorry, I am too tired!")
}

func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

func main() {
	// encapsulating some information using interface
	p := NewPerson("James", 34)
	p.SayHello()

	t := NewPerson("Test", 101)
	t.SayHello()
}
