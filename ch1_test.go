package tdd

import (
	"fmt"
	"testing"
)

func Test_ドルの積(t *testing.T) {
	five := NewDollar(5)
	tests := []struct {
		in   int
		want Money
	}{
		{2, NewDollar(10)},
		{3, NewDollar(15)},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("$%v*%v=$%v", five.getAmount(), tc.in, tc.want.getAmount()), func(t *testing.T) {
			product := five.Times(tc.in).(*Dollar)
			if product.getAmount() != tc.want.getAmount() {
				t.Errorf("$%v is not $%v", product.getAmount(), tc.want.getAmount())
			}
		})
	}
}

func Test_フランの積(t *testing.T) {
	five := NewFranc(5)
	tests := []struct {
		in   int
		want Money
	}{
		{2, NewFranc(10)},
		{3, NewFranc(15)},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("$%v*%v=$%v", five.getAmount(), tc.in, tc.want.getAmount()), func(t *testing.T) {
			product := five.Times(tc.in).(*Franc)
			if product.getAmount() != tc.want.getAmount() {
				t.Errorf("$%v is not $%v", product.getAmount(), tc.want.getAmount())
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
		t.Run(fmt.Sprintf("%vと%vが%vことを期待する", tc.target, tc.in, message[tc.want]), func(t *testing.T) {
			if got := tc.target.Equals(tc.in); got != tc.want {
				t.Errorf("%vと%vが%v", tc.target, tc.in, message[!tc.want])
			}
		})
	}
}

func Test_Moneyの種類(t *testing.T) {
	tests := []struct {
		factoryMethod func(int) Money
		want          Kind
	}{
		{NewDollar, USD},
		{NewFranc, CHF},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%vで有ることを期待する", tc.want), func(t *testing.T) {
			if got := tc.factoryMethod(1).currency(); got != tc.want {
				t.Errorf("%v != %v", got, tc.want)
			}
		})
	}
}
