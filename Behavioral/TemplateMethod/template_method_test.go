package template_method

import (
	"testing"
)

func TestTemplateMethod(t *testing.T) {

	expect := "«Test String»"

	qt := NewQuotes(&FrenchQuotes{})

	result := qt.Quotes("Test String")

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
