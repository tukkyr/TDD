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

type Sum struct {
	augend Money
	addend Money
}

func NewSum(augend, addend Money) (s *Sum) {
	s = &Sum{augend, addend}
	return
}

func (s *Sum) reduce(kind Kind) Money {
	amount := s.augend.getAmount() + s.addend.getAmount()
	return New(amount, kind)
}

type Expression interface {
	reduce(Kind) Money
}

type Money interface {
	Times(int) Money
	Equals(Money) bool
	getAmount() int
	currency() Kind
	Plus(Money) Expression
}

type money struct {
	amount int
	kind   Kind
}

func (m *money) Equals(obj Money) bool {
	return m.amount == obj.getAmount() && m.kind == obj.currency()
}

func (m *money) getAmount() int {
	return m.amount
}

func (m *money) currency() Kind {
	return m.kind
}

func (m *money) String() string {
	return fmt.Sprintf("%v<%v>", m.amount, m.kind)
}

func (m *money) Times(mul int) Money {
	return New(m.amount*mul, m.kind)
}

func (m *money) Plus(addend Money) Expression {
	return NewSum(m, addend)
}

func (m *money) reduce(kind Kind) Money {
	return m
}

func New(amount int, currency Kind) Money {
	m := &money{amount, currency}
	return m
}

type Bank struct {
}

func (b *Bank) reduce(source Expression, kind Kind) Money {
	switch v := source.(type) {
	case *money:
		return v.reduce(kind)
	case *Sum:
		return v.reduce(kind)
	default:
		return nil
	}
}
