package AdapterPattern

// In this example, the Adaptee has a method called SpecificRequest that we want to use through the Target interface.
// To do this, we create an Adapter that has a reference to the Adaptee and implements the Target interface.
// The Adapter's Request method simply calls the Adaptee's SpecificRequest method and returns the result.

// The Client is the object that makes requests to the Target interface.
// In this example, it has a field called target of type Target and a method called SetTarget that sets this field.
// The Client also has a method called DoRequest that calls the Request method on the target object and returns the result.

// Target defines the interface that Client expects to use.
type Target interface {
	Request() string
}

// Client is the object that makes requests to the Target interface.
type Client struct {
	target Target
}

func (c *Client) SetTarget(target Target) {
	c.target = target
}

func (c *Client) DoRequest() string {
	return c.target.Request()
}

// Adaptee is the object that has the functionality that we want to adapt.
type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
	return "Specific Request"
}

// Adapter is the object that adapts the Adaptee to the Target interface.
type Adapter struct {
	adaptee *Adaptee
}

func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}
