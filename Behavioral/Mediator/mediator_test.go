package mediator

import (
	"testing"
)

func TestMediator(t *testing.T) {

	farmer := new(Farmer)
	cannery := new(Cannery)
	shop := new(Shop)

	farmer.AddMoney(7500.00)
	cannery.AddMoney(15000.00)
	shop.AddMoney(30000.00)

	СonnectСolleagues(farmer, cannery, shop)

	// A farmer grows a 1000kg tomato
	// and informs the mediator about the completion of his work.
	// Next, the mediator sends the tomatoes to the cannery.
	// After the cannery produces 1000 packs of ketchup,
	// he informs the mediator about his delivery to the store.
	farmer.GrowTomato(1000)

	expect := float64(54750)
	result := shop.GetMoney()

	if result != expect {
		t.Errorf("Expect result to equal %f, but %f.\n", expect, result)
	}
}
