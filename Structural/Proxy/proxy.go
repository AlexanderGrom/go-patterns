// Package proxy is an example of the Adapter Pattern.
package proxy

// Subject provides an interface for a real subject and its surrogate.
type Subject interface {
	Send() string
}

// Proxy implements a surrogate.
type Proxy struct {
	realSubject Subject
}

// Send sends a message
func (p *Proxy) Send() string {
	if p.realSubject == nil {
		p.realSubject = &RealSubject{}
	}
	return "<strong>" + p.realSubject.Send() + "</strong>"
}

// RealSubject implements a real subject
type RealSubject struct {
}

// Send sends a message
func (s *RealSubject) Send() string {
	return "I’ll be back!"
}
