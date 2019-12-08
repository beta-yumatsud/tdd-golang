package sample1

type Franc struct {
	amount int
}

func NewFranc(amount int) Money {
	return &Franc{
		amount: amount,
	}
}

func (d *Franc) times(multiplier int) Money {
	return NewFranc(d.amount * multiplier)
}

func (d *Franc) currency() string {
	return "CHF"
}
