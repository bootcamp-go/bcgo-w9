package main

// This code breaks the Interface Segregation Principle
// because the Vehicle interface is too generic and
// the Bus struct does not need the PlayRadio method
// - it forces the Bus struct to implement a method
//
// "clients should not be forced to depend upon interfaces functionalities that they do not use."

// Vehicle is the interface for all vehicles
type Vehicle interface {
	// Move moves the vehicle
	Move()

	// PlayRadio plays the radio
	PlayRadio()
}

// _____________________________________________________________
// Car is an struct that implements Vehicle
type Car struct {
	// Name is the name of the car
	Name string
}

// Move moves the car
func (c *Car) Move() {
	println("Car is moving")
}

// PlayRadio plays the radio
func (c *Car) PlayRadio() {
	println("Car is playing the radio")
}

// _____________________________________________________________
// Bus is an struct that implements Vehicle
type Bus struct {
	// Name is the name of the bus
	Name string
}

// Move moves the bus
func (b *Bus) Move() {
	println("Bus is moving")
}

// PlayRadio plays the radio
func (b *Bus) PlayRadio() {}