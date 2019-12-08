package sample1

type Money interface {
	times(multiplier int) Money
	currency() string
}
