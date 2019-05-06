package tdd

type Money interface {
	Times(int) Money
	Equals(Money) bool
	getAmount() int
}

type money struct {
	amount int
}

func (s *money) Equals(obj Money) bool {
	return s.amount == obj.getAmount()
}

func (s *money) getAmount() int {
	return s.amount
}

type Dollar struct {
	*money
}

func NewDollar(amount int) *Dollar {
	m := &money{amount}
	return &Dollar{m}
}

func (s *Dollar) Times(mul int) Money {
	return NewDollar(s.amount * mul)
}

type Franc struct {
	*money
}

func NewFranc(amount int) *Franc {
	m := &money{amount}
	return &Franc{m}
}

func (s *Franc) Times(mul int) Money {
	return NewFranc(s.amount * mul)
}
