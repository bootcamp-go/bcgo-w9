package main

// VehicleRadio is the interface for vehicle radio
type VehicleRadio interface {
	PlayRadio()
}

// VehicleEngine is the interface for vehicle engine
type VehicleEngine interface {
	Move()
}

// _____________________________________________________________
// Car is an struct that implements VehicleEngine and VehicleRadio
type Car struct {
	// Name is the name of the car
	Name string

	// Engine is the engine of the car
	Engine VehicleEngine

	// Radio is the radio of the car
	Radio VehicleRadio
}

// Bus is an struct that implements VehicleEngine
type Bus struct {
	// Name is the name of the bus
	Name string

	// Engine is the engine of the bus
	Engine VehicleEngine
}