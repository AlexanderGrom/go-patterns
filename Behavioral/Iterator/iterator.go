// Паттерн Итератор (Iterator)
//

package iterator

// Тип Iterator, описывает общий интерфейс для итераторов.
type Iterator interface {
	Next() interface{}
	HasNext() bool
}

// Тип Aggregate, описывает общий интерфейс для коллекций.
type Aggregate interface {
	Iterator() Iterator
}

// Тип BookIterator (ConcreteIterator), реализует итератор по книжной полке
type BookIterator struct {
	shelf   *BookShelf
	current int
}

// Возвращает следующий элемент
func (self *BookIterator) Next() interface{} {
	book := self.shelf.Books[self.current]
	self.current++
	return book
}

// Проверяет доступность следующего элемента
func (self *BookIterator) HasNext() bool {
	if self.current >= len(self.shelf.Books) {
		return false
	}
	return true
}

// Тип BookShelf (ConcreteAggregate), реализует книжную полку (коллекцию элементов)
type BookShelf struct {
	Books []*Book
}

// Создает и возвращает итератор по коллекции
func (self *BookShelf) Iterator() Iterator {
	return &BookIterator{shelf: self}
}

// Добавляет элемент в коллекции
func (self *BookShelf) Add(book *Book) {
	self.Books = append(self.Books, book)
}

// Тип Book, реализует элемент коллекции
type Book struct {
	Name string
}
