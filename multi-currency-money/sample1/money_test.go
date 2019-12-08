package sample1

import (
	"reflect"
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDoller(5)
	product := five.times(2)
	if !reflect.DeepEqual(product, NewDoller(10)) {
		t.Fatalf("amount is not 10: %+v", product)
	}
}

func TestEquality(t *testing.T) {
	five := NewDoller(5)
	if !reflect.DeepEqual(five, NewDoller(5)) {
		t.Fatal("same struct")
	}
	if reflect.DeepEqual(five, NewDoller(6)) {
		t.Fatal("both different")
	}
	five = NewFranc(5)
	if !reflect.DeepEqual(five, NewFranc(5)) {
		t.Fatal("same struct")
	}
	if reflect.DeepEqual(five, NewFranc(6)) {
		t.Fatal("both different")
	}
}

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	product := five.times(2)
	if !reflect.DeepEqual(product, NewFranc(10)) {
		t.Fatalf("amount is not 10: %+v", product)
	}
}

func TestDifferentCurrency(t *testing.T) {
	doll := NewDoller(3)
	flan := NewFranc(3)
	if reflect.DeepEqual(doll, flan) {
		t.Fatalf("bad communication")
	}
}

func TestCurrency(t *testing.T) {
	tests := []struct{
		name  string
		input Money
		want  string
	} {
		{
			name: "equal doll",
			input: NewDoller(2),
			want: "USD",
		},
		{
			name: "equal fran",
			input: NewFranc(2),
			want: "CHF",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.input.currency() != tt.want {
				t.Fatalf("not equal currency. input: %s, want: %s", tt.input.currency(), tt.want)
			}
		})
	}
}