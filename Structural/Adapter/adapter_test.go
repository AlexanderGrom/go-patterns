package adapter

import (
	"testing"
)

func TestAdapter(t *testing.T) {

	adapter := NewAdapter(&Adaptee{})

	req := adapter.Request()

	if req != "Request" {
		t.Errorf("Expect volume to %s, but %s", "Request", req)
	}
}
