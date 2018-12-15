package observer

func ExampleObserver() {

	publisher := &ConcretePublisher{}

	publisher.Attach(&ConcreteObserver{})
	publisher.Attach(&ConcreteObserver{})
	publisher.Attach(&ConcreteObserver{})

	publisher.State = "New State..."

	publisher.Notify()
}
