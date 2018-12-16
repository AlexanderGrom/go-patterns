package abstract_factory

import (
	"testing"
)

func TestAbstractFactory(t *testing.T) {

	cocacolaFactory := NewCocaColaFactory()

	cocacolaWater := cocacolaFactory.CreateWater(2.5)
	cocacolaBottle := cocacolaFactory.CreateBottle(2.5)

	cocacolaBottle.PourWater(cocacolaWater)

	if cocacolaBottle.GetWaterVolume() != cocacolaBottle.GetBottleVolume() {
		t.Errorf("Expect volume to %.1fL, but %.1fL", cocacolaBottle.GetWaterVolume(), cocacolaBottle.GetBottleVolume())
	}
}
