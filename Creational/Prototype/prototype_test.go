package prototype

import (
	"testing"
)

func TestPrototype(t *testing.T) {

	product := NewConcreteProduct("A")
	cloneProduct := product.Clone()

	if cloneProduct.GetName() != product.GetName() {
		t.Error("Expect name \"A\" to equal, but not equal.")
	}
}
