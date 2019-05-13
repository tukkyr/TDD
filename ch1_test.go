package tdd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
			sum := tc.target.Plus(tc.in)
			bank := NewBank()
			if got := bank.reduce(sum, USD); *got.(*money) != *tc.want.(*money) {
				t.Errorf("%v is not %v", got, tc.want)
			}
		})
	}
}

func Test_PlusReturnSum(t *testing.T) {
	five := New(5, USD)
	result := five.Plus(five)
	sum := result.(*Sum)
	assert.Equal(t, *five.(*money), *sum.augend.(*money))
	assert.Equal(t, *five.(*money), *sum.addend.(*money))
}

func Test_ReduceMoney(t *testing.T) {
	bank := NewBank()
	result := bank.reduce(New(1, USD), USD)
	assert.Equal(t, *New(1, USD).(*money), *result.(*money))
}

func Test_ReduceMoneyDifferentDurrency(t *testing.T) {
	bank := NewBank()
	bank.addRate(CHF, USD, 2)
	result := bank.reduce(New(2, CHF), USD)
	assert.Equal(t, *New(1, USD).(*money), *result.(*money))
}

func Test_IdentityRate(t *testing.T) {
	assert.Equal(t, 1, NewBank().rate("USD", "USD"))
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

func Test_異なる種類の足し算(t *testing.T) {
	var fiveBucks, tenFrancs Expression = New(5, USD), New(10, CHF)
	bank := NewBank()
	bank.addRate(CHF, USD, 2)
	result := bank.reduce(fiveBucks.Plus(tenFrancs), "USD")
	assert.Equal(t, *New(10, USD).(*money), *result.(*money))
}

func Test_SumPlusMoney(t *testing.T) {
	var fiveBucks, tenFrancs Expression = New(5, USD), New(10, CHF)
	bank := NewBank()
	bank.addRate(CHF, USD, 2)
	sum := NewSum(fiveBucks, tenFrancs).Plus(New(5, USD))
	result := bank.reduce(sum, "USD")
	assert.Equal(t, *New(15, USD).(*money), *result.(*money))
}

func Test_SumTimes(t *testing.T) {
	var fiveBucks, tenFrancs Expression = New(5, USD), New(10, CHF)
	bank := NewBank()
	bank.addRate(CHF, USD, 2)
	sum := NewSum(fiveBucks, tenFrancs).Times(2)
	result := bank.reduce(sum, "USD")
	assert.Equal(t, *New(20, USD).(*money), *result.(*money))
}

// func Test_5ドル足す5ドルはMoney(t *testing.T) {
// 	sum := New(5, USD).Plus(New(5, USD))
// 	_, ok := sum.(Money)
// 	t.Logf("sum type is %T\n", sum)
// 	assert.True(t, ok)
// }
