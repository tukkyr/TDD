package tdd

type Dollar struct {
	amount int
}

func (s *Dollar) Times(mul int) *Dollar {
	return &Dollar{s.amount * mul}
}

func (s *Dollar) Equals(dollar *Dollar) bool {
	return s.amount == dollar.amount
}

type Franc struct {
	amount int
}

func (s *Franc) Times(mul int) *Franc {
	return &Franc{s.amount * mul}
}

func (s *Franc) Equals(dollar *Franc) bool {
	return s.amount == dollar.amount
}
