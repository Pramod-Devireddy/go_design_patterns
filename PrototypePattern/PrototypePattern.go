package PrototypePattern

// Declare an interface for creating an object
type Prototype interface {
	Clone() Prototype
}

// Declare a struct for the object to be created
type ConcretePrototype struct {
	Name string
}

// Implement the Clone method for the ConcretePrototype struct
func (cp *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{
		Name: cp.Name,
	}
}
