package DecoratorPattern

import "fmt"

// The decorator pattern is a design pattern that allows you to add new functionality to an existing object without modifying its structure.
// It does this by creating a wrapper object that contains the original object and adds the new functionality to it.

// Shape is an interface that defines a draw method
type Shape interface {
	Draw()
}

// Circle is a struct that implements the Shape interface
type Circle struct{}

// Draw is a method of the Circle struct that implements the Shape interface
func (c *Circle) Draw() {
	fmt.Println("Drawing Circle")
}

// Decorator is a struct that has a Shape field and a Draw method that calls the Draw method of the Shape field
type Decorator struct {
	Shape Shape
}

// Draw is a method of the Decorator struct that calls the Draw method of the Shape field
func (d *Decorator) Draw() {
	d.Shape.Draw()
}

// RedShapeDecorator is a struct that embeds the Decorator struct and has a SetRedBorder method that sets the border color of the decorated shape to red
type RedShapeDecorator struct {
	*Decorator
}

// SetRedBorder is a method of the RedShapeDecorator struct that sets the border color of the decorated shape to red
func (r *RedShapeDecorator) SetRedBorder() {
	fmt.Println("Setting red border")
}
