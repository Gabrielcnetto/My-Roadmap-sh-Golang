package main

import "fmt"

type Address struct {
	Street, City, State string
}
type country struct {
	Name string
}

type Person struct {
	Name string
	Address
	country
}

func main() {
	p := Person{
		Name:    "Alice",
		Address: Address{"123 Main St", "Anytown", "CA"},
		country: country{Name: "Brazil"},
	}
	fmt.Println(p.Street, p.country.Name)
}
