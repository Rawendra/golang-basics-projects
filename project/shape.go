package main

type Shape interface {
	getArea() float64
}

type Square struct {
	side float64
}
type Circle struct {
	radius float64
}

type Reactangle struct {
	height float64
	width  float64
}

func (s Square) getArea() float64 {
	return s.side * s.side
}

func (c Circle) getArea() float64 {
	return c.radius * c.radius * (22 / 7)
}
func (r Reactangle) getArea() float64 {
	return r.height * r.width
}
func getArea(s Shape) float64 {
	return s.getArea()
}
