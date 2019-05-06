package tdd

import (
	"fmt"
	"testing"
)

func Test_Moneyの積(t *testing.T) {
	tests := []struct {
		target Money
		in     int
		want   Money
	}{
		{New(5, USD), 2, New(10, USD)},
		{New(5, USD), 3, New(15, USD)},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v*%v=%v", tc.target, tc.in, tc.want), func(t *testing.T) {
			product := tc.target.Times(tc.in).(*money)
			if product.getAmount() != tc.want.getAmount() {
				t.Errorf("%v is not %v", product, tc.want)
			}
		})
	}
}

func Test_Moneyどうしの足し算(t *testing.T) {
	tests := []struct {
		target Money
		in     Money
		want   Money
	}{
		{New(5, USD), New(5, USD), New(10, USD)},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v+%v=%v", tc.target, tc.in, tc.want), func(t *testing.T) {
			var sum Expression = tc.target.Plus(tc.in)
			bank := &Bank{}
			if got := bank.Reduce(sum, USD); *got.(*money) != *tc.want.(*money) {
				t.Errorf("%v is not %v", got, tc.want)
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
		{New(5, USD), New(5, USD), true},
		{New(5, USD), New(6, USD), false},
		{New(5, USD), New(5, CHF), false},
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
		factoryMethod func(int, Kind) Money
		want          Kind
	}{
		{New, USD},
		{New, CHF},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%vで有ることを期待する", tc.want), func(t *testing.T) {
			if got := tc.factoryMethod(1, tc.want).currency(); got != tc.want {
				t.Errorf("%v != %v", got, tc.want)
			}
		})
	}
}
