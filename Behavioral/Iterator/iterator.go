// Package iterator is an example of the Iterator Pattern.
package iterator

// Iterator provides a iterator interface.
type Iterator interface {
	Index() int
	Value() interface{}
	Has() bool
	Next()
	Prev()
	Reset()
	End()
}

// Aggregate provides a collection interface.
type Aggregate interface {
	Iterator() Iterator
}

// BookIterator implements the Iterator interface.
type BookIterator struct {
	shelf    *BookShelf
	index    int
	internal int
}

// Index returns current index
func (i *BookIterator) Index() int {
	return i.index
}

// Value returns current value
func (i *BookIterator) Value() interface{} {
	return i.shelf.Books[i.index]
}

// Has implementation.
func (i *BookIterator) Has() bool {
	if i.internal < 0 || i.internal >= len(i.shelf.Books) {
		return false
	}
	return true
}

// Next goes to the next item.
func (i *BookIterator) Next() {
	i.internal++
	if i.Has() {
		i.index++
	}
}

// Prev goes to the previous item.
func (i *BookIterator) Prev() {
	i.internal--
	if i.Has() {
		i.index--
	}
}

// Reset resets iterator.
func (i *BookIterator) Reset() {
	i.index = 0
	i.internal = 0
}

// End goes to the last item.
func (i *BookIterator) End() {
	i.index = len(i.shelf.Books) - 1
	i.internal = i.index
}

// BookShelf implements the Aggregate interface.
type BookShelf struct {
	Books []*Book
}

// Iterator creates and returns the iterator over the collection.
func (b *BookShelf) Iterator() Iterator {
	return &BookIterator{shelf: b}
}

// Add adds an item to the collection.
func (b *BookShelf) Add(book *Book) {
	b.Books = append(b.Books, book)
}

// Book implements a item of the collection.
type Book struct {
	Name string
}
