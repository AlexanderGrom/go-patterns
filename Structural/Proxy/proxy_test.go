package proxy

import (
	"testing"
)

func TestProxy(t *testing.T) {

	expect := "<strong>I’ll be back!</strong>"
	
	proxy := &Proxy{}
	
	result := proxy.Send()
	
	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}