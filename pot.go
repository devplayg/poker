package poker

type Pot interface {
	Total() int64
	Add(money int64)
}
