package main

import (
	"fmt"
	"go_design_patterns/FactoryPattern"
	"go_design_patterns/SingletonPattern"
)

func main() {

	// Factory Pattern
	rect := FactoryPattern.NewShape("rectangle", 10, 20) // creates a new Rectangle with width 10 and height 20
	fmt.Println(rect.Area())                             // prints 200

	circle := FactoryPattern.NewShape("circle", 5) // creates a new Circle with radius 5
	fmt.Println(circle.Area())                     // prints 78.53981633974483

	triangle := FactoryPattern.NewShape("triangle", 10, 20) // creates a new Triangle with base 10 and height 20
	fmt.Println(triangle.Area())                            // prints 100

	// Singleton Pattern
	s1 := SingletonPattern.GetInstance()
	s2 := SingletonPattern.GetInstance()

	if s1 == s2 {
		fmt.Println("s1 and s2 are the same instance")
	} else {
		fmt.Println("s1 and s2 are different instances")
	}

}