// Package composite is an example of the Composite Pattern.
package composite

// Component provides an interface for branches and leaves of a tree.
type Component interface {
	Add(child Component)
	Name() string
	Child() []Component
	Print(prefix string) string
}

// Directory implements branches of a tree
type Directory struct {
	name   string
	childs []Component
}

// Add appends an element to the tree branch.
func (d *Directory) Add(child Component) {
	d.childs = append(d.childs, child)
}

// Name returns name of the Component.
func (d *Directory) Name() string {
	return d.name
}

// Child returns child elements.
func (d *Directory) Child() []Component {
	return d.childs
}

// Print returns the branche in string representation.
func (d *Directory) Print(prefix string) string {
	result := prefix + "/" + d.Name() + "\n"
	for _, val := range d.Child() {
		result += val.Print(prefix + "/" + d.Name())
	}
	return result
}

// File implements a leaves of a tree
type File struct {
	name string
}

// Add implementation.
func (f *File) Add(child Component) {
}

// Name returns name of the Component.
func (f *File) Name() string {
	return f.name
}

// Child implementation.
func (f *File) Child() []Component {
	return []Component{}
}

// Print returns the leave in string representation.
func (f *File) Print(prefix string) string {
	return prefix + "/" + f.Name() + "\n"
}

// NewDirectory is constructor.
func NewDirectory(name string) *Directory {
	return &Directory{
		name: name,
	}
}

// NewFile is constructor.
func NewFile(name string) *File {
	return &File{
		name: name,
	}
}
