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

func New(amount int, currency Kind) Money {
	m := &money{amount, currency}
	return m
}

func (m *money) Times(mul int) Money {
	return New(m.amount*mul, m.kind)
}
