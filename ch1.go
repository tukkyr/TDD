package tdd

type Dollar struct {
	amount int
}

func (s *Dollar) Times(mul int) *Dollar {
	return &Dollar{s.amount * mul}
}
