// Паттерн Адаптер (Adapter)
//

package adapter

// Тип Target, описывает интерфейс с которым наша система хотела бы работать
type Target interface {
	Request() string
}

// Тип Adaptee, реализует ту систему, которую нужно адаптировать.
type Adaptee struct {
}

// Специфический Request
// Будем его адаптировать
func (self *Adaptee) SpecificRequest() string {
	return "Request"
}

// Тип Adapter, реализует адаптер
type Adapter struct {
	*Adaptee // Для адаптации используем встраивание
}

// Адаптирующий метод
func (self *Adapter) Request() string {
	return self.SpecificRequest()
}
