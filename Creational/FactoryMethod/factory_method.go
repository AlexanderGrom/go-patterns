// Паттерн Фабричный метод (FactoryMethod)
//

package factory_method

import (
	"log"
)

// Тип Creater, описывает интерфейс,
// который должна реализовать конкретная фабрика
// для производства продуктов.
type Creater interface {
	CreateProduct(action string) Producter // Параметризированный Фабричный Метод
	registerProduct(Producter)             // Регистрация созданого подукта (для полноты картины)
}

// Тип Producter, описывающий интерфейс продукта,
// который возвращает фабрика.
// Все продукты возвращаемые фабрикой должны придерживаться
// единого интерфейса.
type Producter interface {
	Use() string // каждый продукт должно быть можно использовать
}

// Тип конкретной фабрики по производству продуктов.
// Она имеет фабричный метод и производит продукты.
type ConcreteCreator struct {
	products []*Producter // произведенные продукты (для полноты картины)
}

// Параметризированный Фабричный Метод
func (self *ConcreteCreator) CreateProduct(action string) Producter {
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

	self.registerProduct(product)

	return product
}

// Регистрация созданого подукта на фабрике
func (self *ConcreteCreator) registerProduct(product Producter) {
	self.products = append(self.products, &product)
}

// Тип конкретного продукта "A", который создает фабрика
type ConcreteProductA struct {
	action string // Действие которое делает продукт
}

// Продукт можно использовать и получить действие которое он делает
func (self *ConcreteProductA) Use() string {
	return self.action
}

// Тип конкретного продукта "B", который создает фабрика
type ConcreteProductB struct {
	action string // Действие которое делает продукт
}

// Продукт можно использовать и получить действие которое он делает
func (self *ConcreteProductB) Use() string {
	return self.action
}

// Тип конкретного продукта "C", который создает фабрика
type ConcreteProductC struct {
	action string // Действие которое делает продукт
}

// Продукт можно использовать и получить действие которое он делает
func (self *ConcreteProductC) Use() string {
	return self.action
}
