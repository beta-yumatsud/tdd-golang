package sample2

type Pair struct {
	from string
	to   string
}

func NewPair(from string, to string) Pair {
	return Pair{from: from, to: to}
}

