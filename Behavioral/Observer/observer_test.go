package observer

func ExampleObserver() {

	subject := &ConcreteSubject{}

	subject.Attach(&ConcreteObserver{})
	subject.Attach(&ConcreteObserver{})
	subject.Attach(&ConcreteObserver{})

	subject.State = "New State..."
	
	subject.Notify()
}