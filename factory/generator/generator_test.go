package main

import "testing"

func TestNewEmployeeFactory(t *testing.T) {
	employee := NewEmployeeFactory("Software Engineer", 10000)
	testEmp := employee("Test")

	if testEmp.Name != "Test" {
		t.Errorf("Name should be equals")
	}

	if testEmp.Position != "Software Engineer" {
		t.Errorf("Position should be equals")
	}

	if testEmp.AnnualIncome != 10000 {
		t.Errorf("Income should be 10000")
	}
}

func TestNewEmployee2(t *testing.T) {
	empFactory := NewEmployeeFactory2("Software Engineer", 10000)
	testEmp := empFactory.Create("Test")

	if testEmp.Name != "Test" {
		t.Errorf("Name should be equals")
	}

	if testEmp.Position != "Software Engineer" {
		t.Errorf("Position should be equals")
	}

	if testEmp.AnnualIncome != 10000 {
		t.Errorf("Income should be 10000")
	}
}
