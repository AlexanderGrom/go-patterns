// Паттерн Декоратор (Decorator)
//

package decorator

// Тип Component, описывает интерфейс для декоратора и компонента
type Component interface {
	Operation() string
}

// Тип ConcreteComponent, реализует компонент
type ConcreteComponent struct {
}

// Операция, которую делает компонент сам по себе
func (self *ConcreteComponent) Operation() string {
	return "I am component!"
}

// Тип ConcreteDecorator, реализует декоратор
type ConcreteDecorator struct {
	component Component
}

// Декоратор оборачивает операцию компонента в тег <strong>
func (self *ConcreteDecorator) Operation() string {
	return "<strong>" + self.component.Operation() + "</strong>"
}
