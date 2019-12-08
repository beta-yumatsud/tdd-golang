package sample2

type Expression interface {
	Plus(added Expression) Expression
	Times(multiplier int) Expression
	Reduce(bank Bank,to string) Money
}
