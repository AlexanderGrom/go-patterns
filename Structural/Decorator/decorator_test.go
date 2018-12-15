package decorator

import (
	"testing"
)

func TestDecorator(t *testing.T) {

	expect := "<strong>I am component!</strong>"

	decorator := &ConcreteDecorator{&ConcreteComponent{}}

	result := decorator.Operation()

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
