package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}

func main() {
	john := Person{"John",
		&Address{"1234 London Rd", "London", "UK"}}

	// Problem with the copying
	// jane := john
	// jane.Name = "Jane" // OK
	// jane.Address.StreetAddress = "321 Baker St"

	// deep copying
	jane := john
	jane.Address = &Address{
		john.Address.StreetAddress,
		john.Address.City,
		john.Address.Country,
	}

	jane.Name = "Jane" // OK
	jane.Address.StreetAddress = "321 Baker St"

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
