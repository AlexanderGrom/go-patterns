package chain_of_responsibilities

import (
	"testing"
)

func TestChainOfResponsibilities(t *testing.T) {

	expect := "Im handler 2"
	
	handlers := &ConcreteHandlerA{
		next:&ConcreteHandlerB{
			next:&ConcreteHandlerC{},
		},
	}
	
	result := handlers.SendRequest(2)
	
	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}