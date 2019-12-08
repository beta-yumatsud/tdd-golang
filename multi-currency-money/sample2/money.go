package sample2

type Money struct {
	Amount   int
	currency string
}

var _ Expression = (*Money)(nil)

func NewMoney(amount int, currency string) Money {
	return Money{
		Amount:   amount,
		currency: currency,
	}
}

func Doller(amount int) Money {
	return Money{
		Amount:   amount,
		currency: "USD",
	}
}

func Franc(amount int) Money {
	return Money{
		Amount:   amount,
		currency: "CHF",
	}
}

func (m Money) Reduce(bank Bank, to string) Money {
	rate := bank.Rate(m.currency, to)
	return Money{
		Amount:   m.Amount / rate,
		currency: to,
	}
}

func (m *Money) Times(multiplier int) Expression {
	return &Money{
		Amount:   multiplier * m.Amount,
		currency: m.Currency(),
	}
}

func (m *Money) Plus(addend Expression) Expression {
	return NewSum(m, addend)
}

func (m *Money) Currency() string {
	return m.currency
}
