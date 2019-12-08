package sample2

import (
	"reflect"
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := Doller(5)
	product := five.Times(2)
	want := Doller(10)
	if !reflect.DeepEqual(product, &want) {
		t.Fatalf("amount is not 10: %+v", product)
	}
}

func TestEquality(t *testing.T) {
	if !reflect.DeepEqual(Doller(5), Doller(5)) {
		t.Fatal("must equal")
	}
	if reflect.DeepEqual(Doller(5), Doller(6)) {
		t.Fatal("must not equal")
	}
}

func TestCurrency(t *testing.T) {
	tests := []struct{
		name  string
		input Money
		want  string
	} {
		{
			name:  "equal USD",
			input: Doller(5),
			want:  "USD",
		},
		{
			name:  "equal CHF",
			input: Franc(5),
			want:  "CHF",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.input.Currency() != tt.want {
				t.Fatal("must equal")
			}
		})
	}
}

func TestSimpleAddition(t *testing.T) {
	five := Doller(5)
	sum := five.Plus(&five)
	bank := NewBank()
	reduced := bank.Reduce(sum, "USD")
	if !reflect.DeepEqual(Doller(10), reduced) {
		t.Fatal("must equal")
	}
}

func TestPlusReturnsSum(t *testing.T) {
	five := Doller(5)
	result := five.Plus(&five)
	sum, _ := result.(*Sum)
	if !reflect.DeepEqual(&five, sum.addend) {
		t.Fatal("must equal")
	}
	if !reflect.DeepEqual(&five, sum.augend) {
		t.Fatal("must equal")
	}
}

func TestReduceSum(t *testing.T) {
	threeDoll := Doller(3)
	fourDoll := Doller(4)
	sum := NewSum(&threeDoll, &fourDoll)
	result := NewBank().Reduce(sum, "USD")
	if !reflect.DeepEqual(Doller(7), result) {
		t.Fatal("must equal")
	}
}

func TestReduceMoney(t *testing.T) {
	oneDoll := Doller(1)
	result := NewBank().Reduce(&oneDoll, "USD")
	if !reflect.DeepEqual(Doller(1), result) {
		t.Fatal("must equal")
	}
}

func TestReduceMoneyDifferentCurrency(t *testing.T) {
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	franc := Franc(2)
	result := bank.Reduce(&franc, "USD")
	if !reflect.DeepEqual(Doller(1), result) {
		t.Fatal("must equal")
	}
}

func TestIdentityRate(t *testing.T) {
	if NewBank().Rate("USD", "USD") != 1 {
		t.Fatal("must equal")
	}
}

func TestMixedAddition(t *testing.T) {
	fiveBucks := Doller(5)
	tenFrancs := Franc(10)
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(fiveBucks.Plus(&tenFrancs), "USD")
	if !reflect.DeepEqual(Doller(10), result) {
		t.Fatal("must equal")
	}
}

func TestSumPlusMoney(t *testing.T) {
	fiveBucks := Doller(5)
	tenFrancs := Franc(10)
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	sum := NewSum(&fiveBucks, &tenFrancs).Plus(&fiveBucks)
	result := bank.Reduce(sum, "USD")
	want := Doller(15)
	if !reflect.DeepEqual(want, result) {
		t.Fatalf("must equal. want: %+v, result: %+v", want, result)
	}
}

func TestSumTimes(t *testing.T) {
	fiveBucks := Doller(5)
	tenFrancs := Franc(10)
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	sum := NewSum(&fiveBucks, &tenFrancs).Times(2)
	result := bank.Reduce(sum, "USD")
	want := Doller(20)
	if !reflect.DeepEqual(want, result) {
		t.Fatalf("must equal. want: %+v, result: %+v", want, result)
	}
}
