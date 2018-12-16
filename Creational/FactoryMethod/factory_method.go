// Package factory_method is an example of the Factory Method Pattern.
package factory_method

import (
	"log"
)

// Creater provides a factory interface.
type Creater interface {
	CreateProduct(action string) Producter // Factory Method
}

// Producter provides a product interface.
// All products returned by factory must provide a single interface.
type Producter interface {
	Use() string // Every product should be can be used
}

// ConcreteCreater implements Creater interface.
type ConcreteCreater struct {
}

// NewCreater is the ConcreteCreater constructor.
func NewCreater() Creater {
	return &ConcreteCreater{}
}

// CreateProduct is a Factory Method
func (p *ConcreteCreater) CreateProduct(action string) Producter {
	var product Producter

	switch action {
	case "A":
		product = &ConcreteProductA{action}
	case "B":
		product = &ConcreteProductB{action}
	case "C":
		product = &ConcreteProductC{action}
	default:
		log.Fatalln("Unknown Action")
	}

	return product
}

// ConcreteProductA implements product "A"
type ConcreteProductA struct {
	action string
}

// Use returns product action
func (p *ConcreteProductA) Use() string {
	return p.action
}

// ConcreteProductB implements product "B"
type ConcreteProductB struct {
	action string
}

// Use returns product action
func (p *ConcreteProductB) Use() string {
	return p.action
}

// ConcreteProductC implements product "C"
type ConcreteProductC struct {
	action string
}

// Use returns product action
func (p *ConcreteProductC) Use() string {
	return p.action
}
