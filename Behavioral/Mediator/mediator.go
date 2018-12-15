// Package mediator is an example of the Mediator Pattern.
package mediator

// Mediator provides a mediator interface.
type Mediator interface {
	Notify(msg string)
}

// Тип ConcreteMediator, реализует посредника
type ConcreteMediator struct {
	*Farmer
	*Cannery
	*Shop
}

// Notify implementation.
func (m *ConcreteMediator) Notify(msg string) {
	if msg == "Farmer: Tomato complete..." {
		m.Cannery.AddMoney(-15000.00)
		m.Farmer.AddMoney(15000.00)
		m.Cannery.MakeKetchup(m.Farmer.GetTomato())
	} else if msg == "Cannery: Ketchup complete..." {
		m.Shop.AddMoney(-30000.00)
		m.Cannery.AddMoney(30000.00)
		m.Shop.SellKetchup(m.Cannery.GetKetchup())
	}
}

// СonnectСolleagues connects all colleagues.
func СonnectСolleagues(farmer *Farmer, cannery *Cannery, shop *Shop) {
	mediator := &ConcreteMediator{
		Farmer:  farmer,
		Cannery: cannery,
		Shop:    shop,
	}

	mediator.Farmer.SetMediator(mediator)
	mediator.Cannery.SetMediator(mediator)
	mediator.Shop.SetMediator(mediator)
}

// Farmer implements a Farmer colleague
type Farmer struct {
	mediator Mediator
	tomato   int
	money    float64
}

// SetMediator sets mediator.
func (f *Farmer) SetMediator(mediator Mediator) {
	f.mediator = mediator
}

// AddMoney adds money.
func (f *Farmer) AddMoney(m float64) {
	f.money += m
}

// GrowTomato implementation.
func (f *Farmer) GrowTomato(tomato int) {
	f.tomato = tomato
	f.money -= 7500.00
	f.mediator.Notify("Farmer: Tomato complete...")
}

// GetTomato returns tomatos.
func (f *Farmer) GetTomato() int {
	return f.tomato
}

// Cannery implements a Cannery colleague.
type Cannery struct {
	mediator Mediator
	ketchup  int
	money    float64
}

// SetMediator sets mediator.
func (c *Cannery) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

// AddMoney adds money.
func (c *Cannery) AddMoney(m float64) {
	c.money += m
}

// MakeKetchup implementation.
func (c *Cannery) MakeKetchup(tomato int) {
	c.ketchup = tomato
	c.mediator.Notify("Cannery: Ketchup complete...")
}

// GetKetchup returns ketchup.
func (c *Cannery) GetKetchup() int {
	return c.ketchup
}

// Shop implements a Shop colleague.
type Shop struct {
	mediator Mediator
	money    float64
}

// SetMediator sets mediator.
func (s *Shop) SetMediator(mediator Mediator) {
	s.mediator = mediator
}

// AddMoney adds money.
func (s *Shop) AddMoney(m float64) {
	s.money += m
}

// SellKetchup converts ketchup to money.
func (s *Shop) SellKetchup(ketchup int) {
	s.money = float64(ketchup) * 54.75
}

// GetMoney returns money.
func (s *Shop) GetMoney() float64 {
	return s.money
}
