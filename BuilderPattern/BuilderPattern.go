package BuilderPattern

// The Builder pattern is a design pattern that separates the construction of a complex object from its representation.
// It allows for the construction of complex objects to be done in a step-by-step fashion,
// using the same construction process for different representations of the object.

// In the given code, we have defined a Builder interface with a Build method that takes a Product struct as an argument and returns a Product struct.
// The Product struct represents the complex object that we want to build.

// We have also defined two concrete implementations of the Builder interface, ConcreteBuilder1 and ConcreteBuilder2,
// each with their own implementation of the Build method.

// The Director struct has a Construct method that takes a Builder interface as an argument and calls the Build method on it,
// passing a Product struct as an argument. The Build method of the Builder interface is then responsible for building the
// Product struct in a specific way and returning it.

type Builder interface {
	SetPart1(string)
	SetPart2(string)
	SetPart3(string)
	GetResult() Product
}

type ConcreteBuilder1 struct {
	product Product
}

func (b *ConcreteBuilder1) SetPart1(p1 string) {
	b.product.part1 = p1
}

func (b *ConcreteBuilder1) SetPart2(p2 string) {
	b.product.part2 = p2
}

func (b *ConcreteBuilder1) SetPart3(p3 string) {
	b.product.part3 = p3
}

func (b *ConcreteBuilder1) GetResult() Product {
	return b.product
}

type ConcreteBuilder2 struct {
	product Product
}

func (b *ConcreteBuilder2) SetPart1(p1 string) {
	b.product.part1 = p1
}

func (b *ConcreteBuilder2) SetPart2(p2 string) {
	b.product.part2 = p2
}

func (b *ConcreteBuilder2) SetPart3(p3 string) {
	b.product.part3 = p3
}

func (b *ConcreteBuilder2) GetResult() Product {
	return b.product
}

type Product struct {
	part1 string
	part2 string
	part3 string
}

type Director struct {
	builder Builder
}

func (d *Director) SetBuilder(b Builder) {
	d.builder = b
}

func (d *Director) Construct() {
	d.builder.SetPart1("Part1")
	d.builder.SetPart2("Part2")
	d.builder.SetPart3("Part3")
}
