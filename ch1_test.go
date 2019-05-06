package tdd

import (
	"fmt"
	"testing"
)

func Test_ドルの積(t *testing.T) {
	five := NewDollar(5)
	tests := []struct {
		in   int
		want *Dollar
	}{
		{2, NewDollar(10)},
		{3, NewDollar(15)},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("$%v*%v=$%v", five.amount, tc.in, tc.want.amount), func(t *testing.T) {
			product := five.Times(tc.in).(*Dollar)
			if product.amount != tc.want.amount {
				t.Errorf("$%v is not $%v", product.amount, tc.want.amount)
			}
		})
	}
}

func Test_フランの積(t *testing.T) {
	five := NewFranc(5)
	tests := []struct {
		in   int
		want *Franc
	}{
		{2, NewFranc(10)},
		{3, NewFranc(15)},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("$%v*%v=$%v", five.amount, tc.in, tc.want.amount), func(t *testing.T) {
			product := five.Times(tc.in).(*Franc)
			if product.amount != tc.want.amount {
				t.Errorf("$%v is not $%v", product.amount, tc.want.amount)
			}
		})
	}
}

func Test_Moneyが等しいかどうか調べる(t *testing.T) {
	message := map[bool]string{true: "等しい", false: "等しくない"}
	tests := []struct {
		target Money
		in     Money
		want   bool
	}{
		{NewDollar(5), NewDollar(5), true},
		{NewDollar(5), NewDollar(6), false},
		{NewFranc(5), NewFranc(5), true},
		{NewFranc(5), NewFranc(6), false},
		{NewDollar(5), NewFranc(5), false},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%vことを期待する", message[tc.want]), func(t *testing.T) {
			if got := tc.target.Equals(tc.in); got != tc.want {
				t.Errorf("%v(%T)と%v(%T)が%v", tc.target.getAmount(), tc.target, tc.in.getAmount(), tc.in, message[!tc.want])
			}
		})
	}
}
