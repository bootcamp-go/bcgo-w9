package main

import "fmt"

type Person struct {
	name    string
	age     int
	height  float64
	license bool
}

func (p Person) String() string {
	return fmt.Sprintf("Person Info: %s (%d years, %f height, %t license)", p.name, p.age, p.height, p.license)
}

func main() {
	// fmt is a package that provides formatting for input and output
	// It also provides functions for string manipulation
	//
	// placeholders:
	// %v - value in default format
	// %T - type of the value
	// %t - boolean
	// %d - decimal integer
	// %b - binary integer
	// %f - floating point number
	// %s - string
	// %q - double-quoted string
	// %p - pointer
	// %v - value in default format

	p := Person{
		name:    "John",
		age:     25,
		height:  1.75,
		license: true,
	}

	// Println prints to stdout
	fmt.Printf("Memory address: %p\n", &p)
	fmt.Printf("Person Info: %s (age: %d, height: %f, license: %t)\n", p.name, p.age, p.height, p.license)
	fmt.Printf("Person Info: %s\n", p)
}