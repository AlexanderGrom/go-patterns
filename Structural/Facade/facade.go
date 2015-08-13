// Паттерн Фасад (Facade)
//

package facade

import (
	"strings"
)

// Создает мужика
func NewMan() *Man {
	return &Man{
		house: &House{},
		tree:  &Tree{},
		child: &Child{},
	}
}

// Тип Man, реализует мужика и фасад
// В мужике заложено:
type Man struct {
	house *House
	tree  *Tree
	child *Child
}

// Мужик должен сделать
func (self *Man) Todo() string {
	result := []string{
		self.house.Build(),
		self.tree.Grow(),
		self.child.Born(),
	}
	return strings.Join(result, "\n")
}

// Тип House, реализует подсистему "Дом"
type House struct {
}

// Строительство дома
func (self *House) Build() string {
	return "Build house"
}

// Тип Tree, реализует подсистему "Дерево"
type Tree struct {
}

// Посадка дерева
func (self *Tree) Grow() string {
	return "Tree grow"
}

// Тип Child, реализует подсистему "Ребёнок"
type Child struct {
}

// Производство детей =)
func (self *Child) Born() string {
	return "Child born"
}
