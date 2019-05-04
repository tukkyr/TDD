package tdd

import (
	"os"
	"testing"
)

var five *Dollar

func TestMain(m *testing.M) {
	five = &Dollar{5}
	os.Exit(m.Run())
}

func Test_5ドルに2をかけると10ドル(t *testing.T) {
	ten := &Dollar{10}
	product := five.Times(2)
	if product.amount != ten.amount {
		t.Errorf("$%v is not $%v", five.amount, ten.amount)
	}
}

func Test_5ドルに3をかけると15ドル(t *testing.T) {
	ten := &Dollar{15}
	product := five.Times(3)
	if product.amount != ten.amount {
		t.Errorf("$%v is not $%v", five.amount, ten.amount)
	}
}
