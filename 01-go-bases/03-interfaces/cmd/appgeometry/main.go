package main

import (
	"app/internal/geometry"
	"fmt"
)

func main() {
	// dependencies
	// -> geometry
	cfg := &geometry.ConfigRectangle{
		Width:  10,
		Height: 20,
	}
	geo := newGeometry(GeoRectangle, cfg)

	// -> geometry
	fmt.Println(geo.Area())
	fmt.Println(geo.Perimeter())
}

// newGeometry returns a new geometry
const (
	GeoRectangle = "rectangle"
	GeoCircle    = "circle"
)
func newGeometry(geoType string, cfg any) geometry.Geometry {
	switch geoType {
	case GeoRectangle:
		c := cfg.(*geometry.ConfigRectangle)
		return geometry.NewRectangle(c)
	case GeoCircle:
		c := cfg.(*geometry.ConfigCircle)
		return geometry.NewCircle(c)
	default:
		return nil
	}
}