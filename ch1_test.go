package tdd

import (
	"fmt"
	"testing"
)

func Test_ドルの積(t *testing.T) {
	five := &Dollar{5}
	tests := []struct {
		in   int
		want *Dollar
	}{
		{2, &Dollar{10}},
		{3, &Dollar{15}},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("$%v*%v=$%v", five.amount, tc.in, tc.want.amount), func(t *testing.T) {
			product := five.Times(tc.in)
			if product.amount != tc.want.amount {
				t.Errorf("$%v is not $%v", product.amount, tc.want.amount)
			}
		})
	}
}

func Test_フランの積(t *testing.T) {
	five := &Franc{5}
	tests := []struct {
		in   int
		want *Franc
	}{
		{2, &Franc{10}},
		{3, &Franc{15}},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("$%v*%v=$%v", five.amount, tc.in, tc.want.amount), func(t *testing.T) {
			product := five.Times(tc.in)
			if product.amount != tc.want.amount {
				t.Errorf("$%v is not $%v", product.amount, tc.want.amount)
			}
		})
	}
}

func Test_Dollarが等しいかどうか調べる(t *testing.T) {
	five := &Dollar{5}
	message := map[bool]string{true: "等しい", false: "等しくない"}
	tests := []struct {
		in   *Dollar
		want bool
	}{
		{&Dollar{5}, true},
		{&Dollar{6}, false},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%vことを期待する", message[tc.want]), func(t *testing.T) {
			if got := five.Equals(tc.in); got != tc.want {
				t.Errorf("%vと%vが%v", five, tc.in, message[!tc.want])
			}
		})
	}
}
