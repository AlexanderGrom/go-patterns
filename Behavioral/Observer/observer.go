// Паттерн Наблюдатель (Observer)
//
// Описана модель проталкивания

package observer

// Тип Subject, описывает интерфейс издателя
type Subject interface {
	Attach(observer Observer)
	Notify()
}

// Тип Observer, описывает интерфейс подписчиков
type Observer interface {
	Update(state string)
}

// Тип ConcreteSubject, реализует издателя
type ConcreteSubject struct {
	observers []Observer
	State     string
}

// Добавляет подписчика
func (self *ConcreteSubject) Attach(observer Observer) {
	self.observers = append(self.observers, observer)
}

// Отправляет уведомления методом проталкивания
func (self *ConcreteSubject) Notify() {
	for _, observer := range self.observers {
		observer.Update(self.State)
	}
}

// Тип ConcreteObserver, реализует подписчика
type ConcreteObserver struct {
	state string
}

// Обновляет состояние подписчика
func (self *ConcreteObserver) Update(state string) {
	self.state = state
}
