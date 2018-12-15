// Package decorator is an example of the Decorator Pattern.
package decorator

// Component provides an interface for a decorator and component.
type Component interface {
	Operation() string
}

// ConcreteComponent implements a component.
type ConcreteComponent struct {
}

// Operation implementation.
func (c *ConcreteComponent) Operation() string {
	return "I am component!"
}

// ConcreteDecorator implements a decorator.
type ConcreteDecorator struct {
	component Component
}

// Operation wraps operation of component
func (d *ConcreteDecorator) Operation() string {
	return "<strong>" + d.component.Operation() + "</strong>"
}
