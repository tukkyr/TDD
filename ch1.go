package tdd

import (
	"fmt"
)

type Kind string

const (
	USD = "USD"
	CHF = "CHF"
)

type Sum struct {
	augend Expression
	addend Expression
}

func NewSum(augend, addend Expression) (s *Sum) {
	s = &Sum{augend, addend}
	return
}

func (s *Sum) reduce(bank *Bank, kind Kind) Money {
	amount := s.augend.reduce(bank, kind).getAmount() + s.addend.reduce(bank, kind).getAmount()
	return New(amount, kind)
}

func (s *Sum) Plus(addend Expression) Expression {
	return NewSum(s, addend)
}

func (s *Sum) Times(multiplier int) Expression {
	return NewSum(s.augend.Times(multiplier), s.addend.Times(multiplier))
}

type Expression interface {
	reduce(*Bank, Kind) Money
	Plus(Expression) Expression
	Times(int) Expression
}

type Money interface {
	Times(int) Expression
	Equals(Money) bool
	getAmount() int
	currency() Kind
	Plus(Expression) Expression
	reduce(*Bank, Kind) Money
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

func (m *money) Times(mul int) Expression {
	return New(m.amount*mul, m.kind)
}

func (m *money) Plus(addend Expression) Expression {
	return NewSum(m, addend)
}

func (m *money) reduce(bank *Bank, kind Kind) Money {
	rate := bank.rate(m.kind, kind)
	return New(m.amount/rate, kind)
}

func New(amount int, currency Kind) Money {
	m := &money{amount, currency}
	return m
}

type RateMap map[Pair]int

type Bank struct {
	rates RateMap
}

func NewBank() (bank *Bank) {
	bank = &Bank{RateMap{}}
	bank.addRate(CHF, USD, 2)
	return
}

func (b *Bank) reduce(source Expression, kind Kind) Money {
	return source.reduce(b, kind)
}

func (b *Bank) addRate(from, to Kind, rate int) {
	b.rates[Pair{from, to}] = rate
	return
}

func (b *Bank) rate(from, to Kind) (rate int) {
	if from == to {
		rate = 1
	} else {
		rate = b.rates[Pair{from, to}]
	}
	return
}

type Pair struct {
	from Kind
	to   Kind
}
