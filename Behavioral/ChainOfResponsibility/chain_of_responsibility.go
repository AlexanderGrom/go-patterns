// Паттерн Цепочка ответственности (Chain Of Responsibility)
//

package chain_of_responsibility

// Тип Handler, описывает интерфейс обработчиков в цепочки
type Handler interface {
	SendRequest(message int) string
}

// Тип ConcreteHandlerA, реализует обработчик "A"
type ConcreteHandlerA struct {
	next Handler
}

func (self *ConcreteHandlerA) SendRequest(message int) (result string) {
	if message == 1 {
		result = "Im handler 1"
	} else if self.next != nil {
		result = self.next.SendRequest(message)
	}
	return
}

// Тип ConcreteHandlerB, реализует обработчик "B"
type ConcreteHandlerB struct {
	next Handler
}

func (self *ConcreteHandlerB) SendRequest(message int) (result string) {
	if message == 2 {
		result = "Im handler 2"
	} else if self.next != nil {
		result = self.next.SendRequest(message)
	}
	return
}

// Тип ConcreteHandlerC, реализует обработчик "C"
type ConcreteHandlerC struct {
	next Handler
}

func (self *ConcreteHandlerC) SendRequest(message int) (result string) {
	if message == 3 {
		result = "Im handler 3"
	} else if self.next != nil {
		result = self.next.SendRequest(message)
	}
	return
}
