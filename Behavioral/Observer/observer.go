// Package observer is an example of the Observer Pattern.
// Push model.
package observer

// Publisher interface.
type Publisher interface {
	Attach(observer Observer)
	SetState(state string)
	Notify()
}

// Observer provides a subscriber interface.
type Observer interface {
	Update(state string)
}

// ConcretePublisher implements the Publisher interface.
type ConcretePublisher struct {
	observers []Observer
	state     string
}

// NewPublisher is the Publisher constructor.
func NewPublisher() Publisher {
	return &ConcretePublisher{}
}

// Attach a Observer
func (s *ConcretePublisher) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

// SetState sets new state
func (s *ConcretePublisher) SetState(state string) {
	s.state = state
}

// Notify sends notifications to subscribers.
// Push model.
func (s *ConcretePublisher) Notify() {
	for _, observer := range s.observers {
		observer.Update(s.state)
	}
}

// ConcreteObserver implements the Observer interface.
type ConcreteObserver struct {
	state string
}

// Update set new state
func (s *ConcreteObserver) Update(state string) {
	s.state = state
}
