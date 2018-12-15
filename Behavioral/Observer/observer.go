// Package observer is an example of the Observer Pattern.
// Push model.
package observer

// Publisher interface.
type Publisher interface {
	Attach(observer Observer)
	Notify()
}

// Observer provides a subscriber interface.
type Observer interface {
	Update(state string)
}

// ConcretePublisher implements the Publisher interface.
type ConcretePublisher struct {
	observers []Observer
	State     string
}

// Attach a Observer
func (s *ConcretePublisher) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

// Notify sends notifications to subscribers.
// Push model.
func (s *ConcretePublisher) Notify() {
	for _, observer := range s.observers {
		observer.Update(s.State)
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
