package geometry

// NewCircle returns a new circle
type ConfigCircle struct {
	Radius float64
}

func NewCircle(cfg *ConfigCircle) (c *Circle) {
	c = &Circle{
		Radius: cfg.Radius,
	}
	return
}

// Circle is an struct that implements Geometry interface
type Circle struct {
	Radius float64
}

// Area returns the area of the circle
func (c Circle) Area() (area float64) {
	area = 3.14 * c.Radius * c.Radius
	return
}

// Perimeter returns the perimeter of the circle
func (c Circle) Perimeter() (perimeter float64) {
	perimeter = 2 * 3.14 * c.Radius
	return
}