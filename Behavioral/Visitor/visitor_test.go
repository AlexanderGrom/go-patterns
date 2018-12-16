package visitor

import (
	"testing"
)

func TestVisitor(t *testing.T) {

	expect := "Buy sushi...Buy pizza...Buy burger..."

	city := new(City)

	city.Add(&SushiBar{})
	city.Add(&Pizzeria{})
	city.Add(&BurgerBar{})

	result := city.Accept(&People{})

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
