// Паттерн Посредник (Mediator)
//

package mediator

// Тип Mediator, описывает интерфейс для просредника
type Mediator interface {
	ConnectColleagues()
	Send(msg string)
}

// Тип Colleague, описывает интерфейс процесса взаимодействия
// объектов-коллег с объектом типа Mediator;
type Colleague interface {
	setMediator(mediator Mediator)
}

// Тип ConcreteMediator, реализует посредника
type ConcreteMediator struct {
	*Farmer
	*Cannery
	*Shop
}

// Подключить коллег
func (self *ConcreteMediator) ConnectColleagues() {
	self.Farmer.setMediator(self)
	self.Cannery.setMediator(self)
	self.Shop.setMediator(self)
}

// Общение с коллегами
func (self *ConcreteMediator) Send(msg string) {
	if msg == "Farmer: Tomato complete..." {
		self.Cannery.money -= 15000.00
		self.Farmer.money += 15000.00
		self.Cannery.MakeKetchup(self.Farmer.GetTomato())
	} else if msg == "Cannery: Ketchup complete..." {
		self.Shop.money -= 30000.00
		self.Cannery.money += 30000.00
		self.Shop.SellKetchup(self.Cannery.GetKetchup())
	}
}

func NewMediator() *ConcreteMediator {
	mediator := &ConcreteMediator{}
	mediator.ConnectColleagues()
	mediator.Farmer.money = 7500.00
	mediator.Cannery.money = 15000.00
	mediator.Shop.money = 30000.00
	return mediator
}

// Тип Farmer, реализует коллегу Фермер
type Farmer struct {
	mediator Mediator
	tomato   int
	money    float64
}

// Установить посредника
func (self *Farmer) setMediator(mediator Mediator) {
	self.mediator = mediator
}

// Фермер выращивает помидоры
func (self *Farmer) GrowTomato(tomato int) {
	self.tomato = tomato
	self.money -= 7500.00
	self.mediator.Send("Farmer: Tomato complete...")
}

// Получить помидоры
func (self *Farmer) GetTomato() int {
	return self.tomato
}

// Тип Cannery, реализует коллегу Завод
type Cannery struct {
	mediator Mediator
	ketchup  int
	money    float64
}

// Установить посредника
func (self *Cannery) setMediator(mediator Mediator) {
	self.mediator = mediator
}

// Завод перерабатывает помидоры в кетчуп
func (self *Cannery) MakeKetchup(tomato int) {
	self.ketchup = tomato
	self.mediator.Send("Cannery: Ketchup complete...")
}

// Получить кетчуп
func (self *Cannery) GetKetchup() int {
	return self.ketchup
}

// Тип Shop, реализует коллегу Магазин
type Shop struct {
	mediator Mediator
	money    float64
}

// Установить посредника
func (self *Shop) setMediator(mediator Mediator) {
	self.mediator = mediator
}

// Магазин продает кетчуп
func (self *Shop) SellKetchup(ketchup int) {
	self.money = float64(ketchup) * 54.75
}

func (self *Shop) GetMoney() float64 {
	return self.money
}
