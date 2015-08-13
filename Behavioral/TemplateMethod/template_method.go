// Паттерн Шаблонный метод (Template Method)
//
// На самом деле этот шаблон основывается на Абстрактном Классе и Наследовании.
// Но т.к. ничего этого в Go нет, будет применено агрегирование.
// Хотя по правилам, это паттерн уровня класса.

package template_method

// Тип QuotesInterface, описывает интерфейс установки разных кавычек
type QuotesInterface interface {
	Open() string
	Close() string
}

// Тип Quotes, рeализует Template Method
type Quotes struct {
	qt QuotesInterface
}

// Template Method
func (self *Quotes) Quotes(str string) string {
	return self.qt.Open() + str + self.qt.Close()
}

// Установка закрывающей кавычки
func NewQuotes(qt QuotesInterface) *Quotes {
	return &Quotes{qt}
}

// Тип FrenchQuotes, рeализует обрамление строки Французскими кавычками (Ёлочками)
type FrenchQuotes struct {
}

// Установка открывающей кавычки
func (self *FrenchQuotes) Open() string {
	return "«"
}

// Установка закрывающей кавычки
func (self *FrenchQuotes) Close() string {
	return "»"
}

// Тип GermanQuotes, рeализует обрамление строки Немецкими кавычками (Лапками)
type GermanQuotes struct {
}

// Установка открывающей кавычки
func (self *GermanQuotes) Open() string {
	return "„"
}

// Установка закрывающей кавычки
func (self *GermanQuotes) Close() string {
	return "“"
}
