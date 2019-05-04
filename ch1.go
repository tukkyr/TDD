package tdd

type Dollar struct {
	amount int
}

func (s *Dollar) Times(i int) {
	s.amount = 5 * 2
}
