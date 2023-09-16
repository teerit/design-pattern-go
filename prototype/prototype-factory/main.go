package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	Suit          int
	City, Country string
}

type Employee struct {
	Name   string
	Office Address
}

func (p *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	// Peek into structure
	// fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)
	result := Employee{}
	_ = d.Decode(&result)
	return &result
}

var mainOffice = Employee{
	"", Address{0, "123 East Dr", "London"}}
var auxOffice = Employee{
	"", Address{0, "66 Wst Dr", "London"}}

func newEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suit = suite
	return result
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&auxOffice, name, suite)
}

func main() {

	jane := NewAuxOfficeEmployee("Jane", 1)
	john := NewMainOfficeEmployee("John", 2)

	fmt.Println(jane)
	fmt.Println(john)

}
