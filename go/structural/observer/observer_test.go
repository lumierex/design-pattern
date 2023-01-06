package observer

import "testing"

func TestSubjectAndObserver(t *testing.T) {
	s := Subject{}

	o1 := Observer1{}
	o2 := Observer2{}


	s.Register(&o1)
	s.Register(&o2)
	

	s.Notify("hello")
	
}


func TestObserverRemove(t *testing.T) {
	s := Subject{}

	o1 := Observer1{}
	o2 := Observer2{}


	s.Register(&o1)
	s.Register(&o2)
	s.Remove(&o2)

	s.Notify("hello")
	
}

