package prototype

import (
	"testing"
)

func TestPrototype(t *testing.T) {

	productA := ConcreteProductA{"A"}
	productB := ConcreteProductA{"B"}
	
	cloneProductA := productA.Clone();
	cloneProductB := productB.Clone();

	if cloneProductA.GetName() != productA.GetName() {
		t.Error("Expect name \"A\" to equal, but not equal.")
	}
	
	if cloneProductB.GetName() != productB.GetName() {
		t.Error("Expect name \"B\" to equal, but not equal.")
	}

}