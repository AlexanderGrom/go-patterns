// Паттерн Хранитель (Memento)
//

package memento

// Тип Originator, реализует хозяина состояния
type Originator struct {
	State string
}

// Создает хранилище состояния
func (self *Originator) CreateMemento() *Memento {
	return &Memento{state: self.State}
}

// Восстанавливает состояние
func (self *Originator) SetMemento(memento *Memento) {
	self.State = memento.GetState()
}

// Тип Memento, реализует хранилище для состояния Originator
type Memento struct {
	state string
}

// Получить состояние
func (self *Memento) GetState() string {
	return self.state
}

// Тип Caretaker, получает и хранит объект-хранитель (Memento),
// пока он не понадобится хозяину.
type Caretaker struct {
	Memento *Memento
}
