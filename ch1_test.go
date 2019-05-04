package tdd

import "testing"

func Test_5ドルに2をかけると10ドル(t *testing.T) {
	ten := &Dollar{10}
	five := &Dollar{5}
	five.Times(2)
	if five.amount != ten.amount {
		t.Errorf("$%v is not $%v", five.amount, ten.amount)
	}
}
