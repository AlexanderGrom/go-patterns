// Паттерн Компоновщик (Composite)
//

package composite

// Тип Component, описывает интерфейс для ветвей и листьев дерева
type Component interface {
	Add(child Component)
	Name() string
	Childs() []Component
	Print(prefix string) string
}

// Тип Directory, реализует ветви дерева
type Directory struct {
	name   string
	childs []Component
}

// Добавляет элемент в ветвь дерева
func (self *Directory) Add(child Component) {
	self.childs = append(self.childs, child)
}

// Возвращает имя компонента
func (self *Directory) Name() string {
	return self.name
}

// Возвращает дочерние элементы
func (self *Directory) Childs() []Component {
	return self.childs
}

// Возвращает листинг дерева
func (self *Directory) Print(prefix string) string {
	result := prefix + "/" + self.Name() + "\n"
	for _, val := range self.Childs() {
		result += val.Print(prefix + "/" + self.Name())
	}
	return result
}

// Тип File, реализует лист дерева
type File struct {
	name   string
}

// В лист нельзя добавить элемент
func (self *File) Add(child Component) {
}

// Возвращает имя компонента
func (self *File) Name() string {
	return self.name
}

// В случае с листом это вегда пустой срез
func (self *File) Childs() []Component {
	return []Component{}
}

func (self *File) Print(prefix string) string {
	return prefix + "/" + self.Name() + "\n"
}

func NewDirectory(name string) *Directory {
	return &Directory{
		name: name,
	}
}

func NewFile(name string) *File {
	return &File{
		name: name,
	}
}
