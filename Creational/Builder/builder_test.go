package builder

import (
	"testing"
)

func TestBuilder(t *testing.T) {

	expect := 	"<header>Header</header>\n"+
				"<article>Content</article>\n"+
				"<footer>Footer</footer>\n"
				
	product := &Product{}
	
	director := Director{&ConcreteBuilder{product}}
	director.Construct()

	result := product.Show()
	
	if result != expect {
		t.Errorf("Expect result to %s, but %s", result, expect)
	}

}