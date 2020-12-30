package main

import (
	"fmt"
	"math"
)

// define interface
type Shape interface {
	area() float64
}

type Rectangle struct {
	height, width float64
}

type Circle struct {
	x, y, radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func getArea(s Shape) float64 {
	return Shape.area(s)
}

func main() {

	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{height: 5, width: 10}

	fmt.Printf("Circle area : %f\n", getArea(circle))
	fmt.Printf("Rectangle area : %f\n", getArea(rectangle))
}
