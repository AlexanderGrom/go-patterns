// Паттерн Стратегия (Strategy)
//

package strategy

// Тип StrategySort, описывает интерфейс стратегий (алгоритмов)
type StrategySort interface {
	Sort([]int)
}

// Тип BubbleSort, рeализует алгоритм сортировки пузырьком
type BubbleSort struct {
}

func (self *BubbleSort) Sort(s []int) {
	size := len(s)
	if size < 2 {
		return
	}
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if s[j] < s[j-1] {
				s[j], s[j-1] = s[j-1], s[j]
			}
		}
	}
}

// Тип InsertionSort, рeализует алгоритм сортировки вставками
type InsertionSort struct {
}

func (self *InsertionSort) Sort(s []int) {
	size := len(s)
	if size < 2 {
		return
	}
	for i := 1; i < size; i++ {
		var j int
		var buff int = s[i]
		for j = i - 1; j >= 0; j-- {
			if s[j] < buff {
				break
			}
			s[j+1] = s[j]
		}
		s[j+1] = buff
	}
}

// Тип Context, рeализует контекст выполнения той или иной стратегии
type Context struct {
	strategy StrategySort
}

// Подмена стратегии (алгоритма)
func (self *Context) Algorithm(a StrategySort) {
	self.strategy = a
}

// Сортировка в зависимости от выбраной стратегии (алгоритма)
func (self *Context) Sort(s []int) {
	self.strategy.Sort(s)
}
