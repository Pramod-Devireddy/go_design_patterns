package ObserverPattern

// Subject is the interface that defines the methods that the observers
// need to implement.
type Subject interface {
	Attach(Observer)
	Detach(Observer)
	Notify()
}

// Observer is the interface that defines the methods that the subjects
// need to implement.
type Observer interface {
	Update(Subject)
}

// ConcreteSubject is a concrete implementation of the Subject interface.
type ConcreteSubject struct {
	observers []Observer
	state     string
}

// Attach adds an observer to the list of observers.
func (s *ConcreteSubject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

// Detach removes an observer from the list of observers.
func (s *ConcreteSubject) Detach(o Observer) {
	for i, observer := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

// Notify sends a notification to all the observers in the list.
func (s *ConcreteSubject) Notify() {
	for _, observer := range s.observers {
		observer.Update(s)
	}
}

// SetState sets the state of the subject and sends a notification to the observers.
func (s *ConcreteSubject) SetState(state string) {
	s.state = state
	s.Notify()
}

// ConcreteObserver is a concrete implementation of the Observer interface.
type ConcreteObserver struct {
	SubjectObj Subject
	State      string
}

// Update updates the state of the observer based on the state of the subject.
func (o *ConcreteObserver) Update(s Subject) {
	o.State = s.(*ConcreteSubject).state
}
