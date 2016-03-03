// Абстрактная фабрика (Abstract Factory)
//

package abstract_factory

// Тип AbstractFactory, описывает интерфейс Абстрактной Фабрики
// по производству газировки.
// Конкретные фабрики должны его реализовать.
type AbstractFactory interface {
	createWater(volume float64) AbstractWater
	createBottle(volume float64) AbstractBottle
}

// Тип AbstractWater, описывает интерфейс напитка.
type AbstractWater interface {
	GetVolume() float64 // Возможность получения объема
}

// Тип AbstractBottle, описывает интерфейс бутылки.
type AbstractBottle interface {
	InteractWater(water AbstractWater)  // Бутылка взаимодействует с напитком
	GetBottleVolume() float64           // Возможность получения объема бутылки
	GetWaterVolume() float64            // Возможность получения объема напитка в бутылке
}

// Тип CocaColaFactory, реализует фабрику по производству CocaCola.
type CocaColaFactory struct {
}

// Создает воду
func (self *CocaColaFactory) createWater(volume float64) *CocaColaWater {
	return &CocaColaWater{volume: volume}
}

// Создает бутылку
func (self *CocaColaFactory) createBottle(volume float64) *CocaColaBottle {
	return &CocaColaBottle{volume: volume}
}

// Напиток CocaCola
type CocaColaWater struct {
	volume float64 // Объем созданого напитка
}

// Получить объем созданого напитка
func (self *CocaColaWater) GetVolume() float64 {
	return self.volume
}

// Бутылка CocaCola
type CocaColaBottle struct {
	water  *CocaColaWater // Бутылка должна содержать напиток
	volume float64        // Объем бутылки
}

// Наливаем напиток в бутылку
func (self *CocaColaBottle) InteractWater(water *CocaColaWater) {
	self.water = water
}

// Получить объем бутылки
func (self *CocaColaBottle) GetBottleVolume() float64 {
	return self.volume
}

// Получить объем налитого напитка
func (self *CocaColaBottle) GetWaterVolume() float64 {
	return self.water.GetVolume()
}
