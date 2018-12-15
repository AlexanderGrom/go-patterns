// Package adapter is an example of the Adapter Pattern.
package adapter

// Target provides an interface with which the system should work.
type Target interface {
	Request() string
}

// Adaptee implements system to be adapted.
type Adaptee struct {
}

// SpecificRequest implementation.
func (a *Adaptee) SpecificRequest() string {
	return "Request"
}

// Adapter implements Target interface and is an adapter.
type Adapter struct {
	*Adaptee
}

// Request is an adaptive method.
func (a *Adapter) Request() string {
	return a.SpecificRequest()
}
