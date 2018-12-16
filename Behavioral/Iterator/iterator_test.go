package iterator

import (
	"testing"
)

func TestIterator(t *testing.T) {

	shelf := new(BookShelf)

	books := []string{"A", "B", "C", "D", "E", "F"}

	for _, book := range books {
		shelf.Add(&Book{Name: book})
	}

	for iterator := shelf.Iterator(); iterator.Has(); iterator.Next() {
		index, value := iterator.Index(), iterator.Value().(*Book)
		if value.Name != books[index] {
			t.Errorf("Expect Book.Name to %s, but %s", books[index], value.Name)
		}
	}
}
