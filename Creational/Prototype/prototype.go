// Package prototype is an example of the Singleton Pattern.
package prototype

// Prototyper provides a cloning interface.
type Prototyper interface {
	Clone() Prototyper
	GetName() string
}

// ConcreteProductA implements product "A"
type ConcreteProductA struct {
	name string // Имя продукта
}

// GetName returns product name
func (p *ConcreteProductA) GetName() string {
	return p.name
}

// Clone returns a cloned object.
func (p *ConcreteProductA) Clone() Prototyper {
	return &ConcreteProductA{p.name}
}

// ConcreteProductB implements product "B"
type ConcreteProductB struct {
	name string // Имя продукта
}

// GetName returns product name
func (p *ConcreteProductB) GetName() string {
	return p.name
}

// Clone returns a cloned object.
func (p *ConcreteProductB) Clone() Prototyper {
	return &ConcreteProductB{p.name}
}
