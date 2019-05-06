package tdd

import (
	"fmt"
)

type Kind int

const (
	USD = iota
	CHF
)

func (k Kind) String() (s string) {
	switch k {
	case USD:
		s = "USD"
	case CHF:
		s = "CHF"
	default:
		s = "unknown"
	}
	return
}

type Money interface {
	Times(int) Money
	Equals(Money) bool
	getAmount() int
	currency() Kind
}

type money struct {
	amount int
	kind   Kind
}

func (s *money) Equals(obj Money) bool {
	return s.amount == obj.getAmount() && s.kind == obj.currency()
}

func (s *money) getAmount() int {
	return s.amount
}

func (s *money) currency() Kind {
	return s.kind
}

func (s *money) String() string {
	return fmt.Sprintf("%v<%v>", s.amount, s.kind)
}

type Dollar struct {
	*money
}

func NewDollar(amount int, currency Kind) Money {
	m := &money{amount, currency}
	return &Dollar{m}
}

func (s *Dollar) Times(mul int) Money {
	return NewDollar(s.amount*mul, USD)
}

type Franc struct {
	*money
}

func NewFranc(amount int, currency Kind) Money {
	m := &money{amount, currency}
	return &Franc{m}
}

func (s *Franc) Times(mul int) Money {
	return NewFranc(s.amount*mul, CHF)
}
