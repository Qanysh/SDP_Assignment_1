package main

import (
	"fmt"
)

type Observer interface {
	Update(message string)
}

type ConcreteObserver struct {
	name string
}

func NewConcreteObserver(name string) *ConcreteObserver {
	return &ConcreteObserver{name}
}

func (co *ConcreteObserver) Update(message string) {
	fmt.Printf("%s получил сообщение: %s\n", co.name, message)
}

type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

type ConcreteSubject struct {
	observers []Observer
	message   string
}

func NewConcreteSubject() *ConcreteSubject {
	return &ConcreteSubject{}
}

func (cs *ConcreteSubject) RegisterObserver(observer Observer) {
	cs.observers = append(cs.observers, observer)
}

func (cs *ConcreteSubject) RemoveObserver(observer Observer) {
	for i, o := range cs.observers {
		if o == observer {
			cs.observers = append(cs.observers[:i], cs.observers[i+1:]...)
			break
		}
	}
}

func (cs *ConcreteSubject) NotifyObservers() {
	for _, observer := range cs.observers {
		observer.Update(cs.message)
	}
}

func (cs *ConcreteSubject) SetMessage(message string) {
	cs.message = message
	cs.NotifyObservers()
}

func main() {
	observer1 := NewConcreteObserver("Наблюдатель 1")
	observer2 := NewConcreteObserver("Наблюдатель 2")
	observer3 := NewConcreteObserver("Наблюдатель 3")

	subject := NewConcreteSubject()

	subject.RegisterObserver(observer1)
	subject.RegisterObserver(observer2)
	subject.RegisterObserver(observer3)

	subject.SetMessage("Новое сообщение!")

	subject.RemoveObserver(observer2)
	subject.SetMessage("Измененное сообщение!")
}
