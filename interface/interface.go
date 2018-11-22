package main

import (
	"fmt"
	"math"
)

func main() {
	circle := Circle{X: 1.0, Y: 1.0, Radius: 10.15}
	rectangle := Rectangle{width: 10, Height: 20}

	fmt.Println(getArea(circle))
	fmt.Println(getArea(rectangle))
}

type Shape interface {
	Area() float64
}

type Circle struct {
	X      float64
	Y      float64
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.Height
}

func getArea(s Shape) float64 {
	return s.Area()
}
