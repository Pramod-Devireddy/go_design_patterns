package main

import (
	"fmt"
	"go_design_patterns/FactoryPattern"
)

func main() {

	// Factory Pattern
	rect := FactoryPattern.NewShape("rectangle", 10, 20) // creates a new Rectangle with width 10 and height 20
	fmt.Println(rect.Area())                             // prints 200

	circle := FactoryPattern.NewShape("circle", 5) // creates a new Circle with radius 5
	fmt.Println(circle.Area())                     // prints 78.53981633974483

	triangle := FactoryPattern.NewShape("triangle", 10, 20) // creates a new Triangle with base 10 and height 20
	fmt.Println(triangle.Area())                            // prints 100

}
