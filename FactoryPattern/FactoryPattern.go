package FactoryPattern

import (
	"math"
)

// In the factory pattern, a factory class or function creates objects of different classes depending on the input provided.
// The factory determines which class to instantiate at runtime based on the input, and returns an instance of that class.

func NewShape(shapeType string, params ...float64) Shape {
	switch shapeType {
	case "rectangle":
		return Rectangle{params[0], params[1]}
	case "circle":
		return Circle{params[0]}
	case "triangle":
		return Triangle{params[0], params[1]}
	default:
		return nil
	}
}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() float64 {
	return (t.base * t.height) / 2
}
