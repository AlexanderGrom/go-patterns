package bridge

import (
	"testing"
)

func TestBridge(t *testing.T) {

	expect := "SssuuuuZzzuuuuKkiiiii"

	car := NewCar(&EngineSuzuki{})

	sound := car.Rase()

	if sound != expect {
		t.Errorf("Expect sound to %s, but %s", expect, sound)
	}
}
