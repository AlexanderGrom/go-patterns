// Паттерн Приспособленец (Flyweight)
//

package flyweight

// Тип Flyweight, описывает общий интерфейс для приспособленцев
type Flyweighter interface {
	GetName() string
	SetName(name string)
}

// Тип FlyweightFactory, реализует фабрику для создания приспособленцев.
// Если подходящий приспособленец есть в пуле, то возвращает его
type FlyweightFactory struct {
	pool map[int]Flyweighter
}

// Создание или получение подходящего приспособленца по состоянию.
// В данном случает состояние это число, которое характеризует конкретного приспособленца
func (self *FlyweightFactory) GetFlyweight(state int) Flyweighter {
	if self.pool == nil {
		self.pool = make(map[int]Flyweighter)
	}
	if _, ok := self.pool[state]; !ok {
		self.pool[state] = &ConcreteFlyweight{state: state}
	}
	return self.pool[state]
}

// Тип ConcreteFlyweight, реализует приспособленца
type ConcreteFlyweight struct {
	state int
	name  string
}

func (self *ConcreteFlyweight) GetName() string {
	return "My name: " + self.name
}

func (self *ConcreteFlyweight) SetName(name string) {
	self.name = name
}
