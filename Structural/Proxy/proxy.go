// Паттерн Заместитель (Proxy)
//

package proxy

// Тип Subject, описывает интерфейс для реального объекта и его суррогата
type Subject interface {
	Send() string
}

// Тип Proxy, реализует объект суррогата
type Proxy struct {
	realSubject Subject
}

func (self *Proxy) Send() string {
	if self.realSubject == nil {
		self.realSubject = &RealSubject{}
	}
	return "<strong>" + self.realSubject.Send() + "</strong>"
}

// Тип RealSubject, реализует реальный объект
type RealSubject struct {
}

func (self *RealSubject) Send() string {
	return "I’ll be back!"
}
