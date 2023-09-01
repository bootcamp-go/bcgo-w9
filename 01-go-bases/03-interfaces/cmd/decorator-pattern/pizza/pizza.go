package main

import "fmt"

func main() {
	// default pizza
	pz := &PizzaBase{
		Price: 10,
	}
	// pizza with anchovy
	var pzA Pizza = &PizzaWithAnchovy{
		Pizza:        pz,
		AnchovyPrice: 2,
	}

	fmt.Println(pzA.GetPrice())
}

// ________________________________________________________________________________
// Pizza is the pizza interface
type Pizza interface {
	GetPrice() int
}

// ________________________________________________________________________________
// Implementations
// PizzaBase is the base pizza					-	#01
type PizzaBase struct {
	Price int
}
// GetPrice returns the price of the pizza
func (p *PizzaBase) GetPrice() int {
	return p.Price
}

// PizzaWithAnchovy is a pizza with anchovy		-	#02
type PizzaWithAnchovy struct {
	Pizza Pizza
	AnchovyPrice int
}
// GetPrice returns the price of the pizza with anchovy
func (p *PizzaWithAnchovy) GetPrice() int {
	return p.Pizza.GetPrice() + p.AnchovyPrice
}