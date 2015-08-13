package iterator

import (
	"testing"
)

func TestIterator(t *testing.T) {

	shelf := &BookShelf{}

	books := []string{"A", "B", "C", "D", "E", "F"}

	for _, book := range books {
		shelf.Add(&Book{name:book})
	}

	i := 0
	for iterator := shelf.Iterator(); iterator.HasNext(); {
		if elm := iterator.Next(); elm.(*Book).name != books[i] {
			t.Errorf("Expect Book.name to %s, but %s", books[i], elm.(*Book).name)
		}
		i++
	}
}