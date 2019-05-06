package tdd

type Money interface {
	Times(int) Money
	Equals(Money) bool
}

type Dollar struct {
	amount int
}

func (s *Dollar) Times(mul int) Money {
	return &Dollar{s.amount * mul}
}

func (s *Dollar) Equals(money Money) bool {
	dollar, ok := money.(*Dollar)
	if !ok {
		return false
	}
	return s.amount == dollar.amount
}

type Franc struct {
	amount int
}

func (s *Franc) Times(mul int) Money {
	return &Franc{s.amount * mul}
}

func (s *Franc) Equals(money Money) bool {
	franc, ok := money.(*Franc)
	if !ok {
		return false
	}
	return s.amount == franc.amount
}
