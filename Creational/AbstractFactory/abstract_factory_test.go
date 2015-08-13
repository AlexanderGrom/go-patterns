package abstract_factory

import (
	"testing"
)

func TestAbstractFactory(t *testing.T) {

	cocacolaFactory := &CocaColaFactory{}
	
	cocacolaWater := cocacolaFactory.createWater(2.5)
	cocacolaBottle := cocacolaFactory.createBottle(2.5)
	
	cocacolaBottle.InteractWater(cocacolaWater)
	
	if cocacolaBottle.GetWaterVolume() != cocacolaBottle.GetBottleVolume() {
		t.Errorf("Expect volume to %.1fL, but %.1fL", cocacolaBottle.GetWaterVolume(), cocacolaBottle.GetBottleVolume())
	}

}