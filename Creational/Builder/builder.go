// Package builder is an example of the Builder Pattern.
package builder

// Вuilder provides a builder interface.
type Вuilder interface {
	MakeHeader(str string)
	MakeContent(str string)
	MakeFooter(str string)
}

// Director implements a manager
type Director struct {
	builder Вuilder
}

// Construct tells the builder what to do and in what order.
func (d *Director) Construct() {
	d.builder.MakeHeader("Header")
	d.builder.MakeContent("Content")
	d.builder.MakeFooter("Footer")
}

// ConcreteBuilder implements Builder interface.
type ConcreteBuilder struct {
	product *Product
}

// MakeHeader builds a header of document..
func (b *ConcreteBuilder) MakeHeader(str string) {
	b.product.Header = "<header>" + str + "</header>\n"
}

// MakeContent builds a content of document.
func (b *ConcreteBuilder) MakeContent(str string) {
	b.product.Content = "<article>" + str + "</article>\n"
}

// MakeFooter builds a footer of document.
func (b *ConcreteBuilder) MakeFooter(str string) {
	b.product.Footer = "<footer>" + str + "</footer>\n"
}

// Product implementation.
type Product struct {
	Header  string
	Content string
	Footer  string
}

// Show returns product.
func (p *Product) Show() string {
	var result string
	result += p.Header
	result += p.Content
	result += p.Footer
	return result
}
