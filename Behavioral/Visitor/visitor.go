// Package visitor is an example of the Visitor Pattern.
package visitor

// Visitor provides a visitor interface.
type Visitor interface {
	VisitSushiBar(p *SushiBar) string
	VisitPizzeria(p *Pizzeria) string
	VisitBurgerBar(p *BurgerBar) string
}

// Place provides an interface for place that the visitor should visit.
type Place interface {
	Accept(v Visitor) string
}

// People implements the Visitor interface.
type People struct {
}

// VisitSushiBar implements visit to SushiBar.
func (v *People) VisitSushiBar(p *SushiBar) string {
	return p.BuySushi()
}

// VisitPizzeria implements visit to Pizzeria.
func (v *People) VisitPizzeria(p *Pizzeria) string {
	return p.BuyPizza()
}

// VisitBurgerBar implements visit to BurgerBar.
func (v *People) VisitBurgerBar(p *BurgerBar) string {
	return p.BuyBurger()
}

// City implements a collection of places to visit.
type City struct {
	places []Place
}

// Add appends Place to the collection.
func (c *City) Add(p Place) {
	c.places = append(c.places, p)
}

// Accept implements a visit to all places in the city.
func (c *City) Accept(v Visitor) string {
	var result string
	for _, p := range c.places {
		result += p.Accept(v)
	}
	return result
}

// SushiBar implements the Place interface.
type SushiBar struct {
}

// Accept implementation.
func (s *SushiBar) Accept(v Visitor) string {
	return v.VisitSushiBar(s)
}

// BuySushi implementation.
func (s *SushiBar) BuySushi() string {
	return "Buy sushi..."
}

// Pizzeria implements the Place interface.
type Pizzeria struct {
}

// Accept implementation.
func (p *Pizzeria) Accept(v Visitor) string {
	return v.VisitPizzeria(p)
}

// BuyPizza implementation.
func (p *Pizzeria) BuyPizza() string {
	return "Buy pizza..."
}

// BurgerBar implements the Place interface.
type BurgerBar struct {
}

// Accept implementation.
func (b *BurgerBar) Accept(v Visitor) string {
	return v.VisitBurgerBar(b)
}

// BuyBurger implementation.
func (b *BurgerBar) BuyBurger() string {
	return "Buy burger..."
}
