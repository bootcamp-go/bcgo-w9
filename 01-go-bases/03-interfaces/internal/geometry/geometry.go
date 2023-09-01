package geometry

// Geometry interface for all geometry types
type Geometry interface {
	// Area returns the area of the geometry
	Area() (r float64)

	// Perimeter returns the perimeter of the geometry
	Perimeter() (r float64)
}