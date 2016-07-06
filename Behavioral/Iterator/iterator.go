// Паттерн Итератор (Iterator)
//

package iterator

// Тип Iterator, описывает общий интерфейс для итераторов.
type Iterator interface {
	Index() int
	Value() interface{}
	Has() bool
	Next()
	Prev()
	Reset()
	End()
}

// Тип Aggregate, описывает общий интерфейс для коллекций.
type Aggregate interface {
	Iterator() Iterator
}

// Тип BookIterator (ConcreteIterator), реализует итератор по книжной полке
type BookIterator struct {
	shelf    *BookShelf
	index    int
	internal int
}

// Возвращает текущий индекс
func (self *BookIterator) Index() int {
	return self.index
}

// Возвращает текущее значение
func (self *BookIterator) Value() interface{} {
	return self.shelf.Books[self.index]
}

// Проверет возможен ли переход к другому элементу
func (self *BookIterator) Has() bool {
	if self.internal < 0 || self.internal >= len(self.shelf.Books) {
		return false
	}
	return true
}

// Переходит к следующему элементу
func (self *BookIterator) Next() {
	self.internal++
	if self.Has() {
		self.index++
	}
}

// Переходит к предыдущему элементу
func (self *BookIterator) Prev() {
	self.internal--
	if self.Has() {
		self.index--
	}
}

// Сбрасывает итератор в начало
func (self *BookIterator) Reset() {
	self.index = 0
	self.internal = 0
}

// Прокручивает итератор в конец
func (self *BookIterator) End() {
	self.index = len(self.shelf.Books) - 1
	self.internal = self.index
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
