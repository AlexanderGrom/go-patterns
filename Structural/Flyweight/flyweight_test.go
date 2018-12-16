package flyweight

import (
	"testing"
)

func TestFlyweight(t *testing.T) {

	expect := "My name: Jeck"

	factory := new(FlyweightFactory)

	flyweight1 := factory.GetFlyweight(1)
	flyweight2 := factory.GetFlyweight(2)
	flyweight3 := factory.GetFlyweight(3)

	flyweight1.SetName("Jim")
	flyweight2.SetName("Jeck")
	flyweight3.SetName("Jill")

	flyweightN := factory.GetFlyweight(2)

	result := flyweightN.GetName()

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
