// Паттерн Прототип (Prototype)
//

package prototype

// Тип Prototype, описывает интерфейс,
// который должен реализовать каждый конкретный продукт.
type Prototyper interface {
	clone() Prototyper // метод клонирования
	GetName() string   // каждый продукт имеет имя
}

// Тип конкретного продукта "A"
type ConcreteProductA struct {
	name string // Имя продукты
}

// Возвращает имя продукты
func (self *ConcreteProductA) GetName() string {
	return self.name
}

// Метод клонирования.
// Каждый объект должен реализовать сам,
// как он будет себя клонировать.
func (self *ConcreteProductA) clone() Prototyper {
	return &ConcreteProductA{self.name}
}

// Тип конкретного продукта "B"
type ConcreteProductB struct {
	name string // Имя продукты
}

// Возвращает имя продукты
func (self *ConcreteProductB) GetName() string {
	return self.name
}

// Метод клонирования.
// Каждый объект должен реализовать сам,
// как он будет себя клонировать.
func (self *ConcreteProductB) clone() Prototyper {
	return &ConcreteProductB{self.name}
}
