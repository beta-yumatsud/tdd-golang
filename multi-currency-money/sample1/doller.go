package sample1

type Doller struct {
	amount int
}

func NewDoller(amount int) Money {
	return &Doller{
		amount: amount,
	}
}

func (d *Doller) times(multiplier int) Money {
	return NewDoller(d.amount * multiplier)
}

func (d *Doller) currency() string {
	return "USD"
}