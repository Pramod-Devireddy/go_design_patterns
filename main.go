package main

import (
	"fmt"
	"go_design_patterns/BuilderPattern"
	"go_design_patterns/FactoryPattern"
	"go_design_patterns/PrototypePattern"
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

	// Builder Pattern
	director := &BuilderPattern.Director{}
	builder1 := &BuilderPattern.ConcreteBuilder1{}
	builder2 := &BuilderPattern.ConcreteBuilder2{}

	director.SetBuilder(builder1)
	director.Construct()
	product1 := builder1.GetResult()

	director.SetBuilder(builder2)
	director.Construct()
	product2 := builder2.GetResult()

	fmt.Println(product1)
	fmt.Println(product2)

	// Prototype Pattern

	// Create an instance of the ConcretePrototype struct
	original := &PrototypePattern.ConcretePrototype{
		Name: "Prototype 1",
	}

	// Use the Clone method to create a copy of the original object
	copy := original.Clone()

	// Print the name of the original and copy objects
	fmt.Println("Original name:", original.Name)
	fmt.Println("Copy name:", copy.(*PrototypePattern.ConcretePrototype).Name)
}
