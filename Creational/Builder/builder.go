// Паттерн Строитель (Builder)
//

package builder

// Тип Вuilder, описывает интерфейс строителя.
// Строитель должен ему соответвовать, что бы строить
// конкретный продукт.
type Вuilder interface {
	makeHeader(str string)
	makeContent(str string)
	makeFooter(str string)
}

// Тип Director, реализует руководителя строителем.
// Примимает строителя и руководит им!
type Director struct {
	builder Вuilder
}

// Конструктор, говорит строителю, что ему делать
// в нужной последовательности.
func (self *Director) Construct() {
	self.builder.makeHeader("Header")
	self.builder.makeContent("Content")
	self.builder.makeFooter("Footer")
}

// Конкретный строитель.
// Умеет строить части продукта.
type ConcreteBuilder struct {
	product *Product
}

// Строит шапку документа
func (self *ConcreteBuilder) makeHeader(str string) {
	self.product.Header = "<header>" + str + "</header>\n"
}

// Строит содержание документа
func (self *ConcreteBuilder) makeContent(str string) {
	self.product.Content = "<article>" + str + "</article>\n"
}

// Строит подвал документа
func (self *ConcreteBuilder) makeFooter(str string) {
	self.product.Footer = "<footer>" + str + "</footer>\n"
}

// Сложный продукт.
// Бывают реализации без него, когда строитель стразу строит,
// а не создает объект продукта и потом конструирует его.
type Product struct {
	Header  string
	Content string
	Footer  string
}

// Показать готовый (или нет) продукт
func (self *Product) Show() string {
	var result string
	result = self.Header
	result += self.Content
	result += self.Footer
	return result
}
