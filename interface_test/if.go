package my_interface

import (
	"log"
)

type Gender int

const (
	Female = iota
	Male
)

type Person interface {
	Name() string
	Title() string
}

func New(gender Gender, firstName, lastName string) (obj Person) {

	p := &person{firstName, lastName}

	switch gender {
	case Male:
		obj = &male{p}
	case Female:
		obj = &female{p}
	default:
		log.Printf("not support gender %v", gender)
		obj = nil
	}
	return
}

type person struct {
	firstName string
	lastName  string
}

func (p *person) Name() string {
	return p.firstName + " " + p.lastName
}

// typeだけ記述することがポイント
type female struct {
	*person
}

func (f *female) Title() string {
	return "Ms."
}

type male struct {
	*person
}

func (m *male) Title() string {
	return "Mr."
}

func FullName(p Person) string {
	return p.Title() + p.Name()
}
