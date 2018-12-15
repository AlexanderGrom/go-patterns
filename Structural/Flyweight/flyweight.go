// Package flyweight is an example of the Flyweight Pattern.
package flyweight

// Flyweighter interface
type Flyweighter interface {
	GetName() string
	SetName(name string)
}

// FlyweightFactory implements a factory.
// If a suitable flyweighter is in pool, then returns it.
type FlyweightFactory struct {
	pool map[int]Flyweighter
}

// GetFlyweight creates or returns a suitable Flyweighter by state.
func (f *FlyweightFactory) GetFlyweight(state int) Flyweighter {
	if f.pool == nil {
		f.pool = make(map[int]Flyweighter)
	}
	if _, ok := f.pool[state]; !ok {
		f.pool[state] = &ConcreteFlyweight{state: state}
	}
	return f.pool[state]
}

// ConcreteFlyweight implements a Flyweighter interface.
type ConcreteFlyweight struct {
	state int
	name  string
}

// GetName returns name
func (f *ConcreteFlyweight) GetName() string {
	return "My name: " + f.name
}

// SetName sets a name
func (f *ConcreteFlyweight) SetName(name string) {
	f.name = name
}
