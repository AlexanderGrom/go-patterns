package chain_of_responsibility

import (
	"testing"
)

func TestChainOfResponsibility(t *testing.T) {

	expect := "Im handler 2"

	handlers := &ConcreteHandlerA{
		next: &ConcreteHandlerB{
			next: &ConcreteHandlerC{},
		},
	}

	result := handlers.SendRequest(2)

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
