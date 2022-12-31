package CommandPattern

import "fmt"

// Command is an interface that defines the Execute method.
type Command interface {
	Execute()
}

// ConcreteCommandA is a concrete implementation of the Command interface.
// It has a reference to a Receiver object, which it passes the command to.
type ConcreteCommandA struct {
	Receiver *Receiver
}

// Execute is the method defined by the Command interface. It calls a method on the Receiver object.
func (c *ConcreteCommandA) Execute() {
	c.Receiver.ActionA()
}

// ConcreteCommandB is another concrete implementation of the Command interface.
// It also has a reference to a Receiver object, which it passes the command to.
type ConcreteCommandB struct {
	Receiver *Receiver
}

// Execute is the method defined by the Command interface. It calls a different method on the Receiver object.
func (c *ConcreteCommandB) Execute() {
	c.Receiver.ActionB()
}

// Invoker is the object that initiates the command. It has a reference to a Command object.
type Invoker struct {
	command Command
}

// SetCommand sets the Invoker's Command object.
func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

// ExecuteCommand executes the Command object's Execute method.
func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

// Receiver is the object that receives and executes the command.
type Receiver struct{}

// ActionA is a method that can be called by a ConcreteCommandA object.
func (r *Receiver) ActionA() {
	fmt.Println("Action A executed")
}

// ActionB is a method that can be called by a ConcreteCommandB object.
func (r *Receiver) ActionB() {
	fmt.Println("Action B executed")
}
