package sample2

type Sum struct {
	augend Expression
	addend Expression
}

var _ Expression = (*Sum)(nil)

func NewSum(augend Expression, addend Expression) Expression {
	return &Sum{
		augend: augend,
		addend: addend,
	}
}

func (s *Sum) Reduce(bank Bank, to string) Money {
	//amount := s.addend.Amount + s.augend.Amount
	amount := s.augend.Reduce(bank, to).Amount + s.addend.Reduce(bank, to).Amount
	return NewMoney(amount, to)
}

func (s *Sum) Plus(added Expression) Expression {
	return NewSum(s, added)
}

func (s *Sum) Times(multiplier int) Expression {
	return NewSum(s.augend.Times(multiplier), s.addend.Times(multiplier))
}