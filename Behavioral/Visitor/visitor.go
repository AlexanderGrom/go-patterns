// Паттерн Посетитель (Visitor)
//

package visitor

// Тип Visitor, описывает интерфейс визитера
type Visitor interface {
	VisitSushiBar(p *SushiBar) string
	VisitPizzeria(p *Pizzeria) string
	VisitBurgerBar(p *BurgerBar) string
}

// Тип Place, описывает интерфейс элементов, которые визитер должен обойти
type Place interface {
	Accept(v Visitor) string
}

// Тип People, рeализует конкретного визитера
type People struct {
}

// Визит в суши бар
func (self *People) VisitSushiBar(p *SushiBar) string {
	return p.BuySushi()
}

// Визит в пиццирию
func (self *People) VisitPizzeria(p *Pizzeria) string {
	return p.BuyPizza()
}

// Визит в бургерную
func (self *People) VisitBurgerBar(p *BurgerBar) string {
	return p.BuyBurger()
}

// Тип City, рeализует структуру (коллекцию), в которой хранятся элементы для обхода
type City struct {
	places []Place //места посещения
}

// Добавляет заведение в коллекцию
func (self *City) Add(p Place) {
	self.places = append(self.places, p)
}

// Посещаем все заведения в городе
func (self *City) Accept(v Visitor) string {
	var result string
	for _, p := range self.places {
		result += p.Accept(v)
	}
	return result
}

// Тип SushiBar, рeализует елемент суши бар
type SushiBar struct {
}

func (self *SushiBar) Accept(v Visitor) string {
	return v.VisitSushiBar(self)
}

// Купить суши
func (self *SushiBar) BuySushi() string {
	return "Buy sushi..."
}

// Тип Pizzeria, рeализует елемент пицерия
type Pizzeria struct {
}

func (self *Pizzeria) Accept(v Visitor) string {
	return v.VisitPizzeria(self)
}

// Купить пиццу
func (self *Pizzeria) BuyPizza() string {
	return "Buy pizza..."
}

// Тип BurgerBar, рeализует елемент бургерная
type BurgerBar struct {
}

func (self *BurgerBar) Accept(v Visitor) string {
	return v.VisitBurgerBar(self)
}

// Купить бургер
func (self *BurgerBar) BuyBurger() string {
	return "Buy burger..."
}
