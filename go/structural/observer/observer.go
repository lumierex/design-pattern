package observer

import "fmt"

type ISubject interface {
	Register(observer IObserver)
	Remove(observer IObserver)
	Notify(msg string)
}

type IObserver interface {
	Update(msg string)
}

//发布者
type Subject struct {
	observers []IObserver
}

// 新增订阅者
func (s *Subject) Register(observer IObserver) {
	s.observers = append(s.observers, observer)
}

// 移除订阅者
func (s *Subject) Remove(observer IObserver) {
	for i, v := range s.observers {
		if v == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
		}
	}
}

func (s *Subject) Notify(msg string) {
	for _, v := range s.observers {
		v.Update(msg)
	}
}

// 订阅者
type Observer1 struct{}

func (o *Observer1) Update(msg string) {
	fmt.Printf("%s observer1 update! \n", msg)
}

type Observer2 struct{}

func (o *Observer2) Update(msg string) {
	fmt.Printf("%s observer2 update! \n", msg)
}
