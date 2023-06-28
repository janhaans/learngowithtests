package geometry

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Triangle struct {
	Height float64
	Base   float64
}

func (t *Triangle) Area() float64 {
	return t.Height * t.Base * 0.5
}

func (t *Triangle) Perimeter() float64 {
	return 0.0
}
