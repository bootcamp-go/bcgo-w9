package geometry

// NewRectangle returns a new rectangle
type ConfigRectangle struct {
	Width  float64
	Height float64
}

func NewRectangle(cfg *ConfigRectangle) (r *Rectangle) {
	r = &Rectangle{
		Width:  cfg.Width,
		Height: cfg.Height,
	}
	return
}

// Rectangle is an struct that implements Geometry interface
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns the area of the rectangle
func (r Rectangle) Area() (area float64) {
	area = r.Width * r.Height
	return
}

// Perimeter returns the perimeter of the rectangle
func (r Rectangle) Perimeter() (perimeter float64) {
	perimeter = 2 * (r.Width + r.Height)
	return
}