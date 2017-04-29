package factory_method

import (
	"testing"
)

func TestFactoryMethod(t *testing.T) {

	assert := []string{"A", "B", "C"}

	factory := new(ConcreteCreator)
	products := []Producter{
		factory.CreateProduct("A"),
		factory.CreateProduct("B"),
		factory.CreateProduct("C"),
	}

	for i, product := range products {
		if action := product.Use(); action != assert[i] {
			t.Errorf("Expect action to %s, but %s.\n", assert[i], action)
		}
	}

}