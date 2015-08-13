package flyweight

import (
	"testing"
)

func TestFlyweight(t *testing.T) {

	expect := "My name: Jeck"
	
	factory := &FlyweightFactory{}
	
	flyweight_1 := factory.GetFlyweight(1)
	flyweight_2 := factory.GetFlyweight(2)
	flyweight_3 := factory.GetFlyweight(3)
	
	flyweight_1.SetName("Jim")
	flyweight_2.SetName("Jeck")
	flyweight_3.SetName("Jill")
	
	flyweight_n := factory.GetFlyweight(2)
	
	result := flyweight_n.GetName()
	
	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}